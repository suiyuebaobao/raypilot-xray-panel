<template>
  <div class="admin-sales-landing">
    <div class="page-header">
      <div>
        <h2>销售首页</h2>
        <p>编辑公网首页的销售文案、套餐展示、节点卖点和页脚链接。</p>
      </div>
      <div class="header-actions">
        <el-button @click="openPreview">预览首页</el-button>
        <el-button @click="fetchConfig" :loading="loading">刷新</el-button>
        <el-button type="primary" @click="handleSave" :loading="saving">保存配置</el-button>
      </div>
    </div>

    <el-form v-loading="loading" class="config-form" label-position="top">
      <el-card>
        <template #header>首屏文案</template>
        <div class="form-grid two">
          <el-form-item label="品牌名称">
            <el-input v-model="form.brand" maxlength="32" />
          </el-form-item>
          <el-form-item label="首页标题">
            <el-input v-model="form.title" maxlength="32" />
          </el-form-item>
        </div>
        <el-form-item label="副标题">
          <el-input v-model="form.subtitle" type="textarea" :rows="3" maxlength="180" show-word-limit />
        </el-form-item>
        <div class="form-grid two">
          <el-form-item label="主按钮文案">
            <el-input v-model="form.primary_cta.label" />
          </el-form-item>
          <el-form-item label="主按钮链接">
            <el-input v-model="form.primary_cta.to" placeholder="/register" />
          </el-form-item>
          <el-form-item label="副按钮文案">
            <el-input v-model="form.secondary_cta.label" />
          </el-form-item>
          <el-form-item label="副按钮链接">
            <el-input v-model="form.secondary_cta.to" placeholder="/login" />
          </el-form-item>
        </div>
      </el-card>

      <el-tabs v-model="activeTab" class="edit-tabs">
        <el-tab-pane label="导航与标签" name="nav">
          <div class="section-grid">
            <el-card>
              <template #header>
                <div class="card-head">
                  <span>顶部导航</span>
                  <el-button size="small" @click="addLink(form.nav_links)">新增</el-button>
                </div>
              </template>
              <div v-for="(link, index) in form.nav_links" :key="`nav-${index}`" class="row-editor link-row">
                <el-input v-model="link.label" placeholder="名称" />
                <el-input v-model="link.to" placeholder="/login 或 #plans" />
                <el-button type="danger" text @click="removeItem(form.nav_links, index)">删除</el-button>
              </div>
            </el-card>

            <el-card>
              <template #header>
                <div class="card-head">
                  <span>首屏徽标</span>
                  <el-button size="small" @click="addText(form.badges)">新增</el-button>
                </div>
              </template>
              <div v-for="(_, index) in form.badges" :key="`badge-${index}`" class="row-editor text-row">
                <el-input v-model="form.badges[index]" placeholder="例如：高速节点" />
                <el-button type="danger" text @click="removeItem(form.badges, index)">删除</el-button>
              </div>
            </el-card>

            <el-card>
              <template #header>
                <div class="card-head">
                  <span>信任标签</span>
                  <el-button size="small" @click="addText(form.trust_tags)">新增</el-button>
                </div>
              </template>
              <div v-for="(_, index) in form.trust_tags" :key="`trust-${index}`" class="row-editor text-row">
                <el-input v-model="form.trust_tags[index]" placeholder="例如：Clash / Mihomo 订阅" />
                <el-button type="danger" text @click="removeItem(form.trust_tags, index)">删除</el-button>
              </div>
            </el-card>
          </div>
        </el-tab-pane>

        <el-tab-pane label="节点展示" name="nodes">
          <el-card>
            <template #header>
              <div class="card-head">
                <span>首屏节点状态</span>
                <el-button size="small" @click="addHeroNode">新增节点</el-button>
              </div>
            </template>
            <div v-for="(node, index) in form.hero_nodes" :key="`hero-node-${index}`" class="block-editor">
              <div class="form-grid four">
                <el-form-item label="标识">
                  <el-input v-model="node.flag" placeholder="HK" />
                </el-form-item>
                <el-form-item label="名称">
                  <el-input v-model="node.name" placeholder="香港入口" />
                </el-form-item>
                <el-form-item label="描述">
                  <el-input v-model="node.desc" placeholder="低延迟中转" />
                </el-form-item>
                <el-form-item label="延迟">
                  <el-input v-model="node.latency" placeholder="35ms" />
                </el-form-item>
              </div>
              <el-button type="danger" text @click="removeItem(form.hero_nodes, index)">删除节点</el-button>
            </div>
          </el-card>
        </el-tab-pane>

        <el-tab-pane label="卖点与场景" name="points">
          <div class="section-grid">
            <el-card>
              <template #header>
                <div class="card-head">
                  <span>核心卖点</span>
                  <el-button size="small" @click="addItem(form.selling_points)">新增</el-button>
                </div>
              </template>
              <div v-for="(item, index) in form.selling_points" :key="`point-${index}`" class="block-editor">
                <div class="form-grid two">
                  <el-form-item label="编号">
                    <el-input v-model="item.no" placeholder="01" />
                  </el-form-item>
                  <el-form-item label="标题">
                    <el-input v-model="item.title" />
                  </el-form-item>
                </div>
                <el-form-item label="描述">
                  <el-input v-model="item.text" type="textarea" :rows="2" />
                </el-form-item>
                <el-button type="danger" text @click="removeItem(form.selling_points, index)">删除卖点</el-button>
              </div>
            </el-card>

            <el-card>
              <template #header>
                <div class="card-head">
                  <span>使用场景</span>
                  <el-button size="small" @click="addItem(form.use_cases)">新增</el-button>
                </div>
              </template>
              <div v-for="(item, index) in form.use_cases" :key="`case-${index}`" class="block-editor">
                <el-form-item label="标题">
                  <el-input v-model="item.title" />
                </el-form-item>
                <el-form-item label="描述">
                  <el-input v-model="item.text" type="textarea" :rows="2" />
                </el-form-item>
                <el-button type="danger" text @click="removeItem(form.use_cases, index)">删除场景</el-button>
              </div>
            </el-card>
          </div>
        </el-tab-pane>

        <el-tab-pane label="套餐卡片" name="plans">
          <el-card>
            <template #header>
              <div class="card-head">
                <span>首页套餐展示</span>
                <el-button size="small" @click="addPlan">新增套餐卡片</el-button>
              </div>
            </template>
            <div v-for="(plan, index) in form.plans" :key="`plan-${index}`" class="block-editor plan-editor">
              <div class="form-grid five">
                <el-form-item label="标签">
                  <el-input v-model="plan.tag" placeholder="POPULAR" />
                </el-form-item>
                <el-form-item label="名称">
                  <el-input v-model="plan.name" />
                </el-form-item>
                <el-form-item label="价格展示">
                  <el-input v-model="plan.price" />
                </el-form-item>
                <el-form-item label="单位说明">
                  <el-input v-model="plan.unit" />
                </el-form-item>
                <el-form-item label="按钮文案">
                  <el-input v-model="plan.action" />
                </el-form-item>
              </div>
              <el-form-item label="推荐卡片">
                <el-switch v-model="plan.featured" active-text="高亮" inactive-text="普通" />
              </el-form-item>
              <div class="feature-editor">
                <div class="card-head compact">
                  <span>套餐权益</span>
                  <el-button size="small" @click="addText(plan.features)">新增权益</el-button>
                </div>
                <div v-for="(_, featureIndex) in plan.features" :key="`plan-${index}-feature-${featureIndex}`" class="row-editor text-row">
                  <el-input v-model="plan.features[featureIndex]" placeholder="权益说明" />
                  <el-button type="danger" text @click="removeItem(plan.features, featureIndex)">删除</el-button>
                </div>
              </div>
              <el-button type="danger" text @click="removeItem(form.plans, index)">删除套餐卡片</el-button>
            </div>
          </el-card>
        </el-tab-pane>

        <el-tab-pane label="FAQ 与页脚" name="footer">
          <div class="section-grid">
            <el-card>
              <template #header>
                <div class="card-head">
                  <span>FAQ</span>
                  <el-button size="small" @click="addFAQ">新增问题</el-button>
                </div>
              </template>
              <div v-for="(faq, index) in form.faqs" :key="`faq-${index}`" class="block-editor">
                <el-form-item label="问题">
                  <el-input v-model="faq.q" />
                </el-form-item>
                <el-form-item label="答案">
                  <el-input v-model="faq.a" type="textarea" :rows="2" />
                </el-form-item>
                <el-button type="danger" text @click="removeItem(form.faqs, index)">删除 FAQ</el-button>
              </div>
            </el-card>

            <el-card>
              <template #header>底部转化区</template>
              <el-form-item label="标题">
                <el-input v-model="form.final_cta.title" />
              </el-form-item>
              <el-form-item label="说明">
                <el-input v-model="form.final_cta.text" type="textarea" :rows="2" />
              </el-form-item>
              <div class="form-grid two">
                <el-form-item label="按钮文案">
                  <el-input v-model="form.final_cta.button_label" />
                </el-form-item>
                <el-form-item label="按钮链接">
                  <el-input v-model="form.final_cta.button_to" />
                </el-form-item>
              </div>
              <el-form-item label="页脚文案">
                <el-input v-model="form.footer_text" />
              </el-form-item>
              <div class="card-head compact">
                <span>页脚链接</span>
                <el-button size="small" @click="addLink(form.final_cta.footer_links)">新增链接</el-button>
              </div>
              <div v-for="(link, index) in form.final_cta.footer_links" :key="`footer-link-${index}`" class="row-editor link-row">
                <el-input v-model="link.label" placeholder="名称" />
                <el-input v-model="link.to" placeholder="/login 或 /platform" />
                <el-button type="danger" text @click="removeItem(form.final_cta.footer_links, index)">删除</el-button>
              </div>
            </el-card>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-form>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import { adminApi } from '@/api'
import { ElMessage } from 'element-plus/es/components/message/index.mjs'

const defaultConfig = {
  brand: 'RayPilot',
  nav_links: [
    { label: '套餐', to: '#plans' },
    { label: '节点', to: '#nodes' },
    { label: '说明', to: '#faq' },
    { label: '登录', to: '/login' },
  ],
  badges: ['高速节点', '稳定订阅', '按量流量'],
  title: '高速 VPN 节点',
  subtitle: '面向 AI、游戏、跨境办公和日常网络访问，提供多地区出口、专属订阅链接和清晰的流量管理。',
  primary_cta: { label: '立即开通', to: '/register' },
  secondary_cta: { label: '已有账号登录', to: '/login' },
  trust_tags: ['VLESS Reality', 'XHTTP 可选', 'Clash / Mihomo 订阅'],
  hero_nodes: [
    { flag: 'HK', name: '香港入口', desc: '低延迟中转', latency: '35ms' },
    { flag: 'US', name: '美国出口', desc: 'AI / 海外服务', latency: '128ms' },
    { flag: 'SG', name: '新加坡备用', desc: '亚洲优化', latency: '68ms' },
  ],
  selling_points: [
    { no: '01', title: '多地区高速节点', text: '按地区和线路能力下发订阅，支持直连与中转线路，减少单点不稳定带来的影响。' },
    { no: '02', title: '流量清晰可查', text: '用户中心展示套餐、剩余流量和订阅链接，用多少、剩多少一目了然。' },
    { no: '03', title: '客户端导入简单', text: '支持 Clash / Mihomo 等常见客户端订阅格式，复制订阅链接即可导入使用。' },
  ],
  plans: [
    { tag: 'STARTER', name: '轻量流量', price: '按套餐', unit: '灵活开通', action: '开始使用', featured: false, features: ['适合临时访问和轻量使用', '标准节点订阅', '用户中心自助查看'] },
    { tag: 'POPULAR', name: '高速节点', price: '推荐', unit: '日常主力', action: '选择推荐', featured: true, features: ['适合 AI、办公和影音访问', '多线路订阅', '支持流量池计费'] },
    { tag: 'PRO', name: '大流量套餐', price: '长期', unit: '高频使用', action: '开通套餐', featured: false, features: ['适合多设备和长期使用', '更多流量额度', '可配合兑换码续费'] },
  ],
  use_cases: [
    { title: 'AI 工具访问', text: '为海外 AI 服务准备稳定出口线路。' },
    { title: '游戏加速', text: '选择低延迟地区节点，减少跨境链路波动。' },
    { title: '跨境办公', text: '让资料查询、远程协作和海外服务访问更顺畅。' },
    { title: '多设备订阅', text: '同一账号管理套餐和订阅链接，使用更方便。' },
  ],
  faqs: [
    { q: '购买后怎么使用？', a: '注册并开通套餐后，在用户中心复制订阅链接，导入 Clash Verge Rev、Mihomo 等客户端即可使用。' },
    { q: '流量怎么计算？', a: '系统会按套餐规则统计已用流量和剩余流量，不同套餐可能有不同的计费倍率。' },
    { q: '支持哪些节点模式？', a: '当前系统支持 VLESS Reality、TCP、XHTTP 和中转线路，具体以下发订阅为准。' },
  ],
  final_cta: {
    title: '现在开通 RayPilot 节点服务',
    text: '注册账号后进入用户中心，选择套餐或兑换码开通订阅。',
    button_label: '创建账号',
    button_to: '/register',
    footer_links: [
      { label: '用户登录', to: '/login' },
      { label: '平台能力', to: '/platform' },
    ],
  },
  footer_text: 'RayPilot VPN 节点与流量服务',
}

const activeTab = ref('nav')
const loading = ref(false)
const saving = ref(false)
const form = reactive(cloneConfig(defaultConfig))

function cloneConfig(value) {
  return JSON.parse(JSON.stringify(value))
}

function assignConfig(value) {
  const next = {
    ...cloneConfig(defaultConfig),
    ...(value || {}),
    primary_cta: { ...defaultConfig.primary_cta, ...(value?.primary_cta || {}) },
    secondary_cta: { ...defaultConfig.secondary_cta, ...(value?.secondary_cta || {}) },
    final_cta: {
      ...cloneConfig(defaultConfig.final_cta),
      ...(value?.final_cta || {}),
      footer_links: Array.isArray(value?.final_cta?.footer_links)
        ? value.final_cta.footer_links
        : cloneConfig(defaultConfig.final_cta.footer_links),
    },
  }
  Object.assign(form, cloneConfig(next))
}

function payload() {
  return cloneConfig(form)
}

function addText(list) {
  list.push('')
}

function addLink(list) {
  list.push({ label: '', to: '' })
}

function addHeroNode() {
  form.hero_nodes.push({ flag: '', name: '', desc: '', latency: '' })
}

function addItem(list) {
  list.push({ no: '', title: '', text: '' })
}

function addPlan() {
  form.plans.push({ tag: '', name: '', price: '', unit: '', action: '开始使用', featured: false, features: [] })
}

function addFAQ() {
  form.faqs.push({ q: '', a: '' })
}

function removeItem(list, index) {
  list.splice(index, 1)
}

function openPreview() {
  window.open('/', '_blank', 'noopener')
}

async function fetchConfig() {
  loading.value = true
  try {
    const res = await adminApi.site.salesLanding()
    assignConfig(res.data)
  } catch (err) {
    ElMessage.error(err.message || '获取销售首页配置失败')
  } finally {
    loading.value = false
  }
}

async function handleSave() {
  saving.value = true
  try {
    const res = await adminApi.site.updateSalesLanding(payload())
    assignConfig(res.data)
    ElMessage.success('销售首页配置已保存')
  } catch (err) {
    ElMessage.error(err.message || '保存销售首页配置失败')
  } finally {
    saving.value = false
  }
}

onMounted(fetchConfig)
</script>

<style scoped>
.admin-sales-landing {
  padding: 20px;
}
.page-header,
.card-head,
.header-actions,
.row-editor {
  display: flex;
  align-items: center;
}
.page-header {
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 20px;
}
.page-header h2 {
  margin: 0;
}
.page-header p {
  margin: 8px 0 0;
  color: var(--rp-muted);
}
.header-actions,
.card-head {
  gap: 10px;
}
.header-actions {
  flex-wrap: wrap;
  justify-content: flex-end;
}
.card-head {
  justify-content: space-between;
  width: 100%;
}
.card-head.compact {
  margin: 8px 0 12px;
}
.config-form {
  display: grid;
  gap: 18px;
}
.edit-tabs {
  margin-top: 4px;
}
.section-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
}
.form-grid {
  display: grid;
  gap: 14px;
}
.form-grid.two {
  grid-template-columns: repeat(2, minmax(0, 1fr));
}
.form-grid.four {
  grid-template-columns: 0.6fr 1fr 1fr 0.7fr;
}
.form-grid.five {
  grid-template-columns: repeat(5, minmax(0, 1fr));
}
.row-editor {
  gap: 10px;
  margin-bottom: 10px;
}
.link-row {
  grid-template-columns: 0.45fr 1fr auto;
}
.text-row {
  grid-template-columns: 1fr auto;
}
.block-editor {
  margin-bottom: 14px;
  padding: 14px;
  border: 1px solid rgba(92, 241, 255, 0.14);
  border-radius: 8px;
  background: rgba(5, 12, 24, 0.42);
}
.block-editor:last-child,
.row-editor:last-child {
  margin-bottom: 0;
}
.plan-editor {
  margin-bottom: 18px;
}
.feature-editor {
  margin: 10px 0;
  padding: 12px;
  border: 1px dashed rgba(92, 241, 255, 0.16);
  border-radius: 8px;
}
@media (max-width: 1100px) {
  .section-grid,
  .form-grid.two,
  .form-grid.four,
  .form-grid.five {
    grid-template-columns: 1fr;
  }
  .page-header {
    align-items: flex-start;
    flex-direction: column;
  }
}
</style>
