package service_test

import (
	"context"
	"testing"
	"time"

	"suiyue/internal/model"
	"suiyue/internal/repository"
	"suiyue/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupNodeOperationsTest(t *testing.T) (*gorm.DB, *service.NodeOperationsService) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(
		&model.NodeHost{},
		&model.Node{},
		&model.NodeRuntimeMetric{},
		&model.NodeHealthCheck{},
		&model.UsageLedger{},
	))
	svc := service.NewNodeOperationsService(
		repository.NewNodeRepository(db),
		repository.NewNodeRuntimeMetricRepository(db),
		repository.NewNodeHealthCheckRepository(db),
		repository.NewUsageLedgerRepository(db),
	)
	return db, svc
}

func TestNodeOperations_BuildHealthCheckHealthy(t *testing.T) {
	_, svc := setupNodeOperationsTest(t)
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	svc.SetNowForTest(func() time.Time { return now })
	svc.SetTCPProbeForTest(func(ctx context.Context, node model.Node) (bool, *int) {
		latency := 35
		return true, &latency
	})
	heartbeatAt := now.Add(-10 * time.Second)
	trafficAt := now.Add(-time.Minute)

	check := svc.BuildHealthCheck(context.Background(), model.Node{
		ID:                   1,
		Name:                 "node-healthy",
		Host:                 "example.com",
		Port:                 443,
		LastHeartbeatAt:      &heartbeatAt,
		LastTrafficSuccessAt: &trafficAt,
		IsEnabled:            true,
	}, model.NodeRuntimeMetric{
		NodeID:             uint64Ptr(1),
		CPUUsagePercent:    10,
		MemoryUsagePercent: 20,
		DiskUsagePercent:   30,
		Load1:              1,
		XrayRunning:        true,
		ObservedAt:         now.Add(-10 * time.Second),
	}, true)

	assert.Equal(t, model.NodeHealthHealthy, check.Status)
	assert.Equal(t, 100, check.HealthScore)
	assert.Equal(t, "healthy", check.ReasonCode)
	assert.True(t, check.TCPReachable)
	assert.Equal(t, 35, *check.TCPLatencyMS)
}

func TestNodeOperations_BuildHealthCheckAgentOffline(t *testing.T) {
	_, svc := setupNodeOperationsTest(t)
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	svc.SetNowForTest(func() time.Time { return now })
	svc.SetTCPProbeForTest(func(ctx context.Context, node model.Node) (bool, *int) {
		latency := 20
		return true, &latency
	})
	heartbeatAt := now.Add(-5 * time.Minute)

	check := svc.BuildHealthCheck(context.Background(), model.Node{
		ID:              2,
		Host:            "example.com",
		Port:            443,
		LastHeartbeatAt: &heartbeatAt,
		IsEnabled:       true,
	}, model.NodeRuntimeMetric{}, false)

	assert.Equal(t, model.NodeHealthDegraded, check.Status)
	assert.Equal(t, 60, check.HealthScore)
	assert.Equal(t, "agent_offline", check.ReasonCode)
}

func TestNodeOperations_BuildHealthCheckRecentReportWithOldSuccessIsHealthy(t *testing.T) {
	_, svc := setupNodeOperationsTest(t)
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	svc.SetNowForTest(func() time.Time { return now })
	svc.SetTCPProbeForTest(func(ctx context.Context, node model.Node) (bool, *int) {
		latency := 20
		return true, &latency
	})
	heartbeatAt := now.Add(-10 * time.Second)
	reportAt := now.Add(-30 * time.Second)
	successAt := now.Add(-2 * time.Hour)

	check := svc.BuildHealthCheck(context.Background(), model.Node{
		ID:                   4,
		Host:                 "example.com",
		Port:                 443,
		LastHeartbeatAt:      &heartbeatAt,
		LastTrafficReportAt:  &reportAt,
		LastTrafficSuccessAt: &successAt,
		IsEnabled:            true,
	}, model.NodeRuntimeMetric{
		NodeID:             uint64Ptr(4),
		CPUUsagePercent:    10,
		MemoryUsagePercent: 20,
		DiskUsagePercent:   30,
		Load1:              1,
		XrayRunning:        true,
		ObservedAt:         now.Add(-10 * time.Second),
	}, true)

	assert.Equal(t, model.NodeHealthHealthy, check.Status)
	assert.Equal(t, 100, check.HealthScore)
	assert.Equal(t, "healthy", check.ReasonCode)
	assert.True(t, check.TrafficOK)
}

func TestNodeOperations_BuildHealthCheckServerUnreachable(t *testing.T) {
	_, svc := setupNodeOperationsTest(t)
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	svc.SetNowForTest(func() time.Time { return now })
	svc.SetTCPProbeForTest(func(ctx context.Context, node model.Node) (bool, *int) {
		return false, nil
	})
	heartbeatAt := now.Add(-5 * time.Minute)

	check := svc.BuildHealthCheck(context.Background(), model.Node{
		ID:              3,
		Host:            "example.com",
		Port:            443,
		LastHeartbeatAt: &heartbeatAt,
		IsEnabled:       true,
	}, model.NodeRuntimeMetric{}, false)

	assert.Equal(t, model.NodeHealthDown, check.Status)
	assert.Equal(t, 25, check.HealthScore)
	assert.Equal(t, "server_unreachable", check.ReasonCode)
}

func TestNodeOperations_RecordNodeHostMetricCopiesToNodes(t *testing.T) {
	db, svc := setupNodeOperationsTest(t)
	host := &model.NodeHost{Name: "host", SSHHost: "example.com", AgentBaseURL: "http://agent", AgentTokenHash: "hash", IsEnabled: true}
	require.NoError(t, db.Create(host).Error)
	require.NoError(t, db.Create(&model.Node{ID: 10, Name: "n1", Host: "example.com", Port: 443, NodeHostID: &host.ID, IsEnabled: true}).Error)
	require.NoError(t, db.Create(&model.Node{ID: 11, Name: "n2", Host: "example.com", Port: 8443, NodeHostID: &host.ID, IsEnabled: true}).Error)
	nodes := []model.Node{}
	require.NoError(t, db.Where("node_host_id = ?", host.ID).Find(&nodes).Error)
	xrayRunning := true

	err := svc.RecordNodeHostMetric(context.Background(), host.ID, nodes, service.RuntimeMetricInput{
		CPUUsagePercent:    12,
		MemoryUsagePercent: 34,
		XrayRunning:        &xrayRunning,
		ObservedAt:         time.Now(),
	})
	require.NoError(t, err)

	var count int64
	require.NoError(t, db.Model(&model.NodeRuntimeMetric{}).Count(&count).Error)
	assert.Equal(t, int64(3), count)
	var nodeMetric model.NodeRuntimeMetric
	require.NoError(t, db.Where("node_id = ?", 10).First(&nodeMetric).Error)
	assert.Equal(t, 12.0, nodeMetric.CPUUsagePercent)
	assert.True(t, nodeMetric.XrayRunning)
}

func uint64Ptr(v uint64) *uint64 {
	return &v
}
