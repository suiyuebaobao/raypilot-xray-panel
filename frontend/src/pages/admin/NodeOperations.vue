<template>
  <div class="node-operations">
    <div class="ops-hero">
      <div>
        <p class="eyebrow">NODE OPERATIONS CENTER</p>
        <h2>节点运营中心</h2>
        <p>按健康度、延迟、运行负载与流量消耗判断节点状态。</p>
      </div>
      <el-button type="primary" :loading="loading" @click="fetchSummary">刷新扫描结果</el-button>
    </div>

    <el-row :gutter="16" class="status-grid">
      <el-col :xs="12" :sm="8" :md="4">
        <div class="status-card healthy">
          <strong>{{ counters.healthy }}</strong>
          <span>健康</span>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="4">
        <div class="status-card degraded">
          <strong>{{ counters.degraded }}</strong>
          <span>异常</span>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="4">
        <div class="status-card down">
          <strong>{{ counters.down }}</strong>
          <span>离线</span>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="4">
        <div class="status-card unchecked">
          <strong>{{ counters.unchecked }}</strong>
          <span>未检查</span>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="4">
        <div class="status-card disabled">
          <strong>{{ counters.disabled }}</strong>
          <span>禁用</span>
        </div>
      </el-col>
      <el-col :xs="12" :sm="8" :md="4">
        <div class="status-card total">
          <strong>{{ counters.total }}</strong>
          <span>总线路</span>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="18" class="rank-grid">
      <el-col :xs="24" :lg="12">
        <el-card>
          <template #header>今日流量排行</template>
          <traffic-rank :items="trafficRank.today" />
        </el-card>
      </el-col>
      <el-col :xs="24" :lg="12">
        <el-card>
          <template #header>本月流量排行</template>
          <traffic-rank :items="trafficRank.month" />
        </el-card>
      </el-col>
    </el-row>

    <el-card class="nodes-card">
      <template #header>
        <div class="card-header">
          <span>节点健康矩阵</span>
          <small>最近生成：{{ summary.generated_at ? formatDate(summary.generated_at) : '-' }}</small>
        </div>
      </template>
      <div class="ops-table-scroll">
        <el-table :data="nodes" border :fit="false" scrollbar-always-on style="width: 100%" v-loading="loading">
          <el-table-column label="节点" min-width="210">
            <template #default="{ row }">
              <div class="node-name">
                <strong>{{ row.node.name }}</strong>
                <small>#{{ row.node.id }} · {{ row.node.host }}:{{ row.node.port }}</small>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="健康" width="150">
            <template #default="{ row }">
              <div class="health-cell">
                <el-tag :type="statusTag(row.health?.status)" effect="dark">{{ row.health_text || '未检查' }}</el-tag>
                <span>{{ row.health?.health_score ?? '-' }} 分</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="故障原因" min-width="240">
            <template #default="{ row }">
              <div class="reason-cell">
                <strong>{{ row.health?.reason_code || 'unchecked' }}</strong>
                <small>{{ row.health?.reason_message || '等待 worker 生成健康检查记录' }}</small>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="延迟" width="100">
            <template #default="{ row }">{{ row.health?.tcp_latency_ms ? `${row.health.tcp_latency_ms} ms` : '-' }}</template>
          </el-table-column>
          <el-table-column label="心跳" width="180">
            <template #default="{ row }">{{ row.node.last_heartbeat_at ? formatDate(row.node.last_heartbeat_at) : '-' }}</template>
          </el-table-column>
          <el-table-column label="流量状态" min-width="210">
            <template #default="{ row }">
              <div class="traffic-cell">
                <el-tag size="small" :type="row.health?.traffic_ok ? 'success' : 'warning'" effect="plain">{{ row.traffic_text }}</el-tag>
                <small>{{ row.node.last_traffic_success_at ? formatDate(row.node.last_traffic_success_at) : '暂无成功上报' }}</small>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="运行负载" min-width="260">
            <template #default="{ row }">
              <div v-if="row.metric" class="metric-grid">
                <span>CPU {{ formatPercent(row.metric.cpu_usage_percent) }}</span>
                <span>MEM {{ formatPercent(row.metric.memory_usage_percent) }}</span>
                <span>DISK {{ formatPercent(row.metric.disk_usage_percent) }}</span>
                <span>LOAD {{ formatNumber(row.metric.load1) }}</span>
                <span>TCP {{ row.metric.tcp_connections }}</span>
                <span>Xray {{ row.metric.xray_running ? '运行' : '停止' }}</span>
              </div>
              <span v-else>-</span>
            </template>
          </el-table-column>
          <el-table-column label="检查时间" width="180">
            <template #default="{ row }">{{ row.health?.checked_at ? formatDate(row.health.checked_at) : '-' }}</template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { computed, defineComponent, h, onMounted, onUnmounted, ref } from 'vue'
import { adminApi } from '@/api'
import { ElMessage } from 'element-plus/es/components/message/index.mjs'

const summary = ref({
  counters: {},
  nodes: [],
  traffic_rank: { today: [], month: [] },
})
const loading = ref(false)
let timer = null

const counters = computed(() => ({
  total: summary.value.counters?.total ?? 0,
  healthy: summary.value.counters?.healthy ?? 0,
  degraded: summary.value.counters?.degraded ?? 0,
  down: summary.value.counters?.down ?? 0,
  unchecked: summary.value.counters?.unchecked ?? 0,
  disabled: summary.value.counters?.disabled ?? 0,
}))
const nodes = computed(() => summary.value.nodes || [])
const trafficRank = computed(() => summary.value.traffic_rank || { today: [], month: [] })

const TrafficRank = defineComponent({
  name: 'TrafficRank',
  props: {
    items: { type: Array, default: () => [] },
  },
  setup(props) {
    return () => {
      if (!props.items.length) {
        return h('div', { class: 'empty-rank' }, '暂无流量数据')
      }
      return h('div', { class: 'rank-list' }, props.items.map((item, index) => h('div', { class: 'rank-row', key: item.node_id || index }, [
        h('span', { class: 'rank-index' }, String(index + 1).padStart(2, '0')),
        h('div', { class: 'rank-main' }, [
          h('strong', item.node_name || `节点 #${item.node_id}`),
          h('small', `${poolLabel(item.traffic_pool)} · ${item.active_user_count || 0} 用户`),
        ]),
        h('div', { class: 'rank-traffic' }, [
          h('strong', formatBytes(item.billed_total)),
          h('small', `真实 ${formatBytes(item.total)}`),
        ]),
      ])))
    }
  },
})

async function fetchSummary() {
  loading.value = true
  try {
    const res = await adminApi.nodeOperations.summary()
    summary.value = res.data || summary.value
  } catch (err) {
    ElMessage.error('获取节点运营数据失败：' + (err.message || '未知错误'))
  } finally {
    loading.value = false
  }
}

function statusTag(status) {
  if (status === 'healthy') return 'success'
  if (status === 'degraded') return 'warning'
  if (status === 'down') return 'danger'
  if (status === 'disabled') return 'info'
  return 'info'
}

function poolLabel(pool) {
  return pool === 'residential' ? '家宽' : '普通'
}

function formatDate(value) {
  if (!value) return '-'
  return new Date(value).toLocaleString('zh-CN', { hour12: false })
}

function formatNumber(value) {
  const n = Number(value || 0)
  return n.toFixed(2)
}

function formatPercent(value) {
  return `${formatNumber(value)}%`
}

function formatBytes(value) {
  const n = Number(value || 0)
  if (n >= 1024 ** 4) return `${(n / 1024 ** 4).toFixed(2)} TB`
  if (n >= 1024 ** 3) return `${(n / 1024 ** 3).toFixed(2)} GB`
  if (n >= 1024 ** 2) return `${(n / 1024 ** 2).toFixed(2)} MB`
  if (n >= 1024) return `${(n / 1024).toFixed(2)} KB`
  return `${n} B`
}

onMounted(() => {
  fetchSummary()
  timer = setInterval(fetchSummary, 30000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})
</script>

<style scoped>
.node-operations {
  display: flex;
  flex-direction: column;
  gap: 18px;
}
.ops-hero {
  position: relative;
  overflow: hidden;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
  padding: 24px;
  border: 1px solid rgba(66, 245, 255, 0.18);
  border-radius: 14px;
  background:
    radial-gradient(circle at 84% 16%, rgba(255, 61, 242, 0.18), transparent 22rem),
    linear-gradient(135deg, rgba(10, 24, 42, 0.96), rgba(5, 10, 22, 0.92));
  box-shadow: 0 18px 48px rgba(0, 0, 0, 0.28);
}
.ops-hero::before {
  content: "";
  position: absolute;
  inset: 0;
  pointer-events: none;
  background: linear-gradient(90deg, transparent, rgba(66, 245, 255, 0.12), transparent);
  transform: translateX(-60%);
  animation: scan 4s linear infinite;
}
.eyebrow {
  margin: 0 0 6px;
  color: var(--rp-cyan);
  font-size: 12px;
  letter-spacing: 0.22em;
}
.ops-hero h2 {
  margin: 0;
  color: var(--rp-text);
  font-size: 28px;
}
.ops-hero p:last-child {
  margin: 8px 0 0;
  color: var(--rp-muted);
}
.status-grid {
  row-gap: 16px;
}
.status-card {
  min-height: 104px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 6px;
  padding: 18px;
  border: 1px solid rgba(66, 245, 255, 0.14);
  border-radius: 12px;
  background: rgba(7, 14, 28, 0.72);
}
.status-card strong {
  color: var(--rp-text);
  font-size: 34px;
  line-height: 1;
}
.status-card span {
  color: var(--rp-muted);
}
.status-card.healthy strong {
  color: var(--rp-success);
}
.status-card.degraded strong {
  color: var(--rp-warning);
}
.status-card.down strong {
  color: var(--rp-danger);
}
.status-card.unchecked strong,
.status-card.disabled strong,
.status-card.total strong {
  color: var(--rp-cyan);
}
.rank-grid {
  row-gap: 18px;
}
.rank-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.rank-row {
  display: grid;
  grid-template-columns: 42px 1fr auto;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border: 1px solid rgba(66, 245, 255, 0.12);
  border-radius: 10px;
  background: rgba(4, 10, 22, 0.45);
}
.rank-index {
  color: var(--rp-cyan);
  font-weight: 800;
  letter-spacing: 0.08em;
}
.rank-main,
.rank-traffic,
.node-name,
.reason-cell,
.traffic-cell {
  display: flex;
  flex-direction: column;
  gap: 4px;
}
.rank-main strong,
.node-name strong,
.reason-cell strong {
  color: var(--rp-text);
}
.rank-main small,
.rank-traffic small,
.node-name small,
.reason-cell small,
.traffic-cell small,
.card-header small {
  color: var(--rp-muted);
}
.rank-traffic {
  align-items: flex-end;
}
.rank-traffic strong {
  color: var(--rp-cyan);
}
.empty-rank {
  padding: 28px;
  color: var(--rp-muted);
  text-align: center;
}
.nodes-card {
  overflow: hidden;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}
.ops-table-scroll {
  width: 100%;
  overflow-x: auto;
}
.health-cell {
  display: flex;
  align-items: center;
  gap: 10px;
}
.health-cell span {
  color: var(--rp-muted);
}
.metric-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(70px, 1fr));
  gap: 6px;
  color: var(--rp-muted);
  font-size: 12px;
}
@keyframes scan {
  0% {
    transform: translateX(-80%);
  }
  100% {
    transform: translateX(80%);
  }
}
@media (max-width: 768px) {
  .ops-hero {
    flex-direction: column;
    align-items: flex-start;
  }
  .rank-row {
    grid-template-columns: 34px 1fr;
  }
  .rank-traffic {
    grid-column: 2;
    align-items: flex-start;
  }
}
</style>
