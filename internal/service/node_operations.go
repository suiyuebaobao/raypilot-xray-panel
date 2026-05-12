package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"suiyue/internal/model"
	"suiyue/internal/repository"

	"gorm.io/gorm"
)

const (
	nodeHeartbeatStaleAfter = 90 * time.Second
	nodeTrafficStaleAfter   = 10 * time.Minute
	nodeLoadMetricStale     = 5 * time.Minute
	defaultTCPProbeTimeout  = 3 * time.Second
)

// RuntimeMetricInput 是 node-agent 心跳上报的运行指标。
type RuntimeMetricInput struct {
	CPUUsagePercent    float64   `json:"cpu_usage_percent"`
	MemoryUsagePercent float64   `json:"memory_usage_percent"`
	DiskUsagePercent   float64   `json:"disk_usage_percent"`
	Load1              float64   `json:"load1"`
	Load5              float64   `json:"load5"`
	Load15             float64   `json:"load15"`
	TCPConnections     uint32    `json:"tcp_connections"`
	XrayRunning        *bool     `json:"xray_running"`
	XrayUptimeSeconds  *uint64   `json:"xray_uptime_seconds"`
	ObservedAt         time.Time `json:"observed_at"`
}

// NodeOperationsService 提供节点运营中心能力。
type NodeOperationsService struct {
	nodeRepo        *repository.NodeRepository
	metricRepo      *repository.NodeRuntimeMetricRepository
	healthRepo      *repository.NodeHealthCheckRepository
	usageLedgerRepo *repository.UsageLedgerRepository
	tcpProbe        func(ctx context.Context, node model.Node) (bool, *int)
	now             func() time.Time
}

// NewNodeOperationsService 创建节点运营中心服务。
func NewNodeOperationsService(
	nodeRepo *repository.NodeRepository,
	metricRepo *repository.NodeRuntimeMetricRepository,
	healthRepo *repository.NodeHealthCheckRepository,
	usageLedgerRepo *repository.UsageLedgerRepository,
) *NodeOperationsService {
	svc := &NodeOperationsService{
		nodeRepo:        nodeRepo,
		metricRepo:      metricRepo,
		healthRepo:      healthRepo,
		usageLedgerRepo: usageLedgerRepo,
		now:             time.Now,
	}
	svc.tcpProbe = svc.defaultTCPProbe
	return svc
}

// SetTCPProbeForTest 替换 TCP 探测器，仅供测试使用。
func (s *NodeOperationsService) SetTCPProbeForTest(fn func(ctx context.Context, node model.Node) (bool, *int)) {
	s.tcpProbe = fn
}

// SetNowForTest 替换当前时间函数，仅供测试使用。
func (s *NodeOperationsService) SetNowForTest(fn func() time.Time) {
	s.now = fn
}

// RecordNodeMetric 记录单个逻辑节点运行指标。
func (s *NodeOperationsService) RecordNodeMetric(ctx context.Context, nodeID uint64, input RuntimeMetricInput) error {
	if s == nil || s.metricRepo == nil {
		return nil
	}
	nodeIDCopy := nodeID
	return s.metricRepo.Create(ctx, metricFromInput(&nodeIDCopy, nil, input, s.now()))
}

// RecordNodeHostMetric 记录物理服务器运行指标，并复制到其下逻辑节点。
func (s *NodeOperationsService) RecordNodeHostMetric(ctx context.Context, nodeHostID uint64, nodes []model.Node, input RuntimeMetricInput) error {
	if s == nil || s.metricRepo == nil {
		return nil
	}
	hostIDCopy := nodeHostID
	if err := s.metricRepo.Create(ctx, metricFromInput(nil, &hostIDCopy, input, s.now())); err != nil {
		return err
	}
	for _, node := range nodes {
		nodeIDCopy := node.ID
		hostIDCopy := nodeHostID
		if err := s.metricRepo.Create(ctx, metricFromInput(&nodeIDCopy, &hostIDCopy, input, s.now())); err != nil {
			return err
		}
	}
	return nil
}

func metricFromInput(nodeID *uint64, nodeHostID *uint64, input RuntimeMetricInput, fallback time.Time) *model.NodeRuntimeMetric {
	observedAt := input.ObservedAt
	if observedAt.IsZero() {
		observedAt = fallback
	}
	xrayRunning := false
	if input.XrayRunning != nil {
		xrayRunning = *input.XrayRunning
	}
	return &model.NodeRuntimeMetric{
		NodeID:             nodeID,
		NodeHostID:         nodeHostID,
		CPUUsagePercent:    sanitizePercent(input.CPUUsagePercent),
		MemoryUsagePercent: sanitizePercent(input.MemoryUsagePercent),
		DiskUsagePercent:   sanitizePercent(input.DiskUsagePercent),
		Load1:              sanitizeFloat(input.Load1),
		Load5:              sanitizeFloat(input.Load5),
		Load15:             sanitizeFloat(input.Load15),
		TCPConnections:     input.TCPConnections,
		XrayRunning:        xrayRunning,
		XrayUptimeSeconds:  input.XrayUptimeSeconds,
		ObservedAt:         observedAt,
	}
}

func sanitizePercent(value float64) float64 {
	value = sanitizeFloat(value)
	if value > 100 {
		return 100
	}
	return value
}

func sanitizeFloat(value float64) float64 {
	if value < 0 {
		return 0
	}
	return value
}

// CheckAllNodes 检查所有节点并写入健康检查记录。
func (s *NodeOperationsService) CheckAllNodes(ctx context.Context) (int, error) {
	nodes, err := s.nodeRepo.List(ctx)
	if err != nil {
		return 0, err
	}
	if len(nodes) == 0 {
		return 0, nil
	}
	nodeIDs := make([]uint64, 0, len(nodes))
	for _, node := range nodes {
		nodeIDs = append(nodeIDs, node.ID)
	}
	metrics, err := s.metricRepo.LatestByNodeIDs(ctx, nodeIDs)
	if err != nil {
		return 0, err
	}
	for _, node := range nodes {
		metric, hasMetric := metrics[node.ID]
		if err := s.CheckNodeHealth(ctx, node, metric, hasMetric); err != nil {
			return 0, err
		}
	}
	return len(nodes), nil
}

// CheckNodeHealth 检查单个节点并写入健康结果。
func (s *NodeOperationsService) CheckNodeHealth(ctx context.Context, node model.Node, metric model.NodeRuntimeMetric, hasMetric bool) error {
	check := s.BuildHealthCheck(ctx, node, metric, hasMetric)
	return s.healthRepo.Create(ctx, check)
}

// BuildHealthCheck 生成单个节点健康检查结果。
func (s *NodeOperationsService) BuildHealthCheck(ctx context.Context, node model.Node, metric model.NodeRuntimeMetric, hasMetric bool) *model.NodeHealthCheck {
	now := s.now()
	if !node.IsEnabled {
		return &model.NodeHealthCheck{
			NodeID:        node.ID,
			Status:        model.NodeHealthDisabled,
			HealthScore:   0,
			ReasonCode:    "node_disabled",
			ReasonMessage: "节点已在后台禁用",
			TCPReachable:  false,
			HeartbeatOK:   false,
			TrafficOK:     false,
			LoadOK:        true,
			CheckedAt:     now,
		}
	}

	heartbeatOK := node.LastHeartbeatAt != nil && now.Sub(*node.LastHeartbeatAt) <= nodeHeartbeatStaleAfter
	tcpReachable, latency := s.tcpProbe(ctx, node)
	trafficOK := nodeTrafficOK(node, now)
	loadOK := nodeLoadOK(metric, hasMetric, now)
	xrayOK := true
	if hasMetric && now.Sub(metric.ObservedAt) <= nodeLoadMetricStale {
		xrayOK = metric.XrayRunning
	}

	score := 100
	if !heartbeatOK {
		score -= 40
	}
	if !tcpReachable {
		score -= 35
	}
	if !xrayOK {
		score -= 20
	}
	if !trafficOK {
		score -= 15
	}
	if !loadOK {
		score -= 10
	}
	if score < 0 {
		score = 0
	}

	status := model.NodeHealthHealthy
	if score < 50 || (!heartbeatOK && !tcpReachable) {
		status = model.NodeHealthDown
	} else if score < 90 {
		status = model.NodeHealthDegraded
	}
	reasonCode, reasonMessage := classifyNodeHealthReason(node, metric, hasMetric, now, heartbeatOK, tcpReachable, trafficOK, loadOK, xrayOK)

	return &model.NodeHealthCheck{
		NodeID:        node.ID,
		Status:        status,
		HealthScore:   score,
		ReasonCode:    reasonCode,
		ReasonMessage: reasonMessage,
		TCPLatencyMS:  latency,
		TCPReachable:  tcpReachable,
		HeartbeatOK:   heartbeatOK,
		TrafficOK:     trafficOK,
		LoadOK:        loadOK,
		CheckedAt:     now,
	}
}

func nodeTrafficOK(node model.Node, now time.Time) bool {
	if node.TrafficErrorCount > 0 || node.LastTrafficErrorAt != nil {
		return false
	}
	if node.LastTrafficSuccessAt == nil {
		return true
	}
	return now.Sub(*node.LastTrafficSuccessAt) <= nodeTrafficStaleAfter
}

func nodeLoadOK(metric model.NodeRuntimeMetric, hasMetric bool, now time.Time) bool {
	if !hasMetric || now.Sub(metric.ObservedAt) > nodeLoadMetricStale {
		return true
	}
	return metric.CPUUsagePercent < 90 && metric.MemoryUsagePercent < 90 && metric.DiskUsagePercent < 95 && metric.Load1 < 32
}

func classifyNodeHealthReason(
	node model.Node,
	metric model.NodeRuntimeMetric,
	hasMetric bool,
	now time.Time,
	heartbeatOK bool,
	tcpReachable bool,
	trafficOK bool,
	loadOK bool,
	xrayOK bool,
) (string, string) {
	if !heartbeatOK && !tcpReachable {
		return "server_unreachable", "节点服务器心跳过期且端口不可达"
	}
	if !heartbeatOK {
		return "agent_offline", "节点端口可达但 node-agent 心跳过期"
	}
	if !tcpReachable {
		return "xray_or_firewall", "node-agent 在线但节点端口不可达"
	}
	if !xrayOK {
		return "xray_not_running", "node-agent 上报 Xray 进程未运行"
	}
	if !trafficOK {
		if node.TrafficErrorCount > 0 || node.LastTrafficErrorAt != nil {
			return "traffic_report_error", "最近一次流量上报处理失败"
		}
		return "traffic_report_stale", "节点流量上报成功时间过旧"
	}
	if !loadOK {
		return "node_overloaded", "节点 CPU、内存、磁盘或负载超过阈值"
	}
	if hasMetric && now.Sub(metric.ObservedAt) > nodeLoadMetricStale {
		return "metric_stale", "运行指标过旧，仅按心跳和端口判断"
	}
	return "healthy", "节点当前状态正常"
}

func (s *NodeOperationsService) defaultTCPProbe(ctx context.Context, node model.Node) (bool, *int) {
	host := strings.TrimSpace(node.Host)
	if host == "" || node.Port == 0 {
		return false, nil
	}
	dialCtx, cancel := context.WithTimeout(ctx, defaultTCPProbeTimeout)
	defer cancel()
	start := time.Now()
	conn, err := (&net.Dialer{}).DialContext(dialCtx, "tcp", net.JoinHostPort(host, fmt.Sprintf("%d", node.Port)))
	if err != nil {
		return false, nil
	}
	_ = conn.Close()
	ms := int(time.Since(start).Milliseconds())
	if ms < 1 {
		ms = 1
	}
	return true, &ms
}

// NodeOperationsSummary 是节点运营中心汇总响应。
type NodeOperationsSummary struct {
	Counters    NodeOperationsCounters `json:"counters"`
	Nodes       []NodeOperationsNode   `json:"nodes"`
	TrafficRank NodeTrafficRankSummary `json:"traffic_rank"`
	GeneratedAt time.Time              `json:"generated_at"`
}

// NodeOperationsCounters 是健康状态计数。
type NodeOperationsCounters struct {
	Total     int `json:"total"`
	Healthy   int `json:"healthy"`
	Degraded  int `json:"degraded"`
	Down      int `json:"down"`
	Disabled  int `json:"disabled"`
	Unchecked int `json:"unchecked"`
}

// NodeOperationsNode 是运营中心节点行。
type NodeOperationsNode struct {
	Node        model.Node               `json:"node"`
	Health      *model.NodeHealthCheck   `json:"health,omitempty"`
	Metric      *model.NodeRuntimeMetric `json:"metric,omitempty"`
	HealthText  string                   `json:"health_text"`
	TrafficText string                   `json:"traffic_text"`
}

// NodeTrafficRankSummary 是流量排行汇总。
type NodeTrafficRankSummary struct {
	Today []repository.NodeUsageTotal `json:"today"`
	Month []repository.NodeUsageTotal `json:"month"`
}

// Summary 返回节点运营中心汇总。
func (s *NodeOperationsService) Summary(ctx context.Context) (*NodeOperationsSummary, error) {
	nodes, err := s.nodeRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	nodeIDs := make([]uint64, 0, len(nodes))
	for _, node := range nodes {
		nodeIDs = append(nodeIDs, node.ID)
	}
	metrics, err := s.metricRepo.LatestByNodeIDs(ctx, nodeIDs)
	if err != nil {
		return nil, err
	}
	healthChecks, err := s.healthRepo.LatestByNodeIDs(ctx, nodeIDs)
	if err != nil {
		return nil, err
	}

	result := &NodeOperationsSummary{
		GeneratedAt: s.now(),
		Nodes:       make([]NodeOperationsNode, 0, len(nodes)),
	}
	for _, node := range nodes {
		item := NodeOperationsNode{Node: node}
		if check, ok := healthChecks[node.ID]; ok {
			checkCopy := check
			item.Health = &checkCopy
			item.HealthText = nodeHealthText(check.Status)
			incrementHealthCounter(&result.Counters, check.Status)
		} else {
			item.HealthText = nodeHealthText(model.NodeHealthUnchecked)
			result.Counters.Unchecked++
		}
		if metric, ok := metrics[node.ID]; ok {
			metricCopy := metric
			item.Metric = &metricCopy
		}
		item.TrafficText = trafficStatusText(node)
		result.Nodes = append(result.Nodes, item)
	}
	result.Counters.Total = len(nodes)
	sort.Slice(result.Nodes, func(i, j int) bool {
		left := statusSortWeight(result.Nodes[i].Health)
		right := statusSortWeight(result.Nodes[j].Health)
		if left != right {
			return left < right
		}
		return result.Nodes[i].Node.ID < result.Nodes[j].Node.ID
	})

	if s.usageLedgerRepo != nil {
		now := s.now()
		todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		tomorrowStart := todayStart.AddDate(0, 0, 1)
		monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		nextMonthStart := monthStart.AddDate(0, 1, 0)
		result.TrafficRank.Today, err = s.usageLedgerRepo.SumByNode(ctx, todayStart, tomorrowStart, 20)
		if err != nil {
			return nil, err
		}
		result.TrafficRank.Month, err = s.usageLedgerRepo.SumByNode(ctx, monthStart, nextMonthStart, 20)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// RecentChecks 查询单节点近期健康检查记录。
func (s *NodeOperationsService) RecentChecks(ctx context.Context, nodeID uint64, limit int) ([]model.NodeHealthCheck, error) {
	if _, err := s.nodeRepo.FindByID(ctx, nodeID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return s.healthRepo.RecentByNode(ctx, nodeID, limit)
}

func incrementHealthCounter(counters *NodeOperationsCounters, status string) {
	switch status {
	case model.NodeHealthHealthy:
		counters.Healthy++
	case model.NodeHealthDegraded:
		counters.Degraded++
	case model.NodeHealthDown:
		counters.Down++
	case model.NodeHealthDisabled:
		counters.Disabled++
	default:
		counters.Unchecked++
	}
}

func statusSortWeight(check *model.NodeHealthCheck) int {
	if check == nil {
		return 4
	}
	switch check.Status {
	case model.NodeHealthDown:
		return 0
	case model.NodeHealthDegraded:
		return 1
	case model.NodeHealthHealthy:
		return 2
	case model.NodeHealthDisabled:
		return 3
	default:
		return 4
	}
}

func nodeHealthText(status string) string {
	switch status {
	case model.NodeHealthHealthy:
		return "健康"
	case model.NodeHealthDegraded:
		return "异常"
	case model.NodeHealthDown:
		return "离线"
	case model.NodeHealthDisabled:
		return "禁用"
	default:
		return "未检查"
	}
}

func trafficStatusText(node model.Node) string {
	if node.TrafficErrorCount > 0 || node.LastTrafficErrorAt != nil {
		return "流量上报异常"
	}
	if node.LastTrafficSuccessAt == nil {
		return "暂无成功上报"
	}
	return "流量上报正常"
}
