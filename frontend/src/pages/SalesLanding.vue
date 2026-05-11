<template>
  <main class="sales-page">
    <header class="sales-nav">
      <router-link class="brand" to="/">
        <span class="brand-mark">RP</span>
        <span>{{ cfg.brand }}</span>
      </router-link>
      <nav>
        <template v-for="link in navLinks" :key="`${link.label}-${link.to}`">
          <router-link v-if="isRouteLink(safeLink(link).to)" :to="safeLink(link).to">{{ link.label }}</router-link>
          <a v-else :href="safeLink(link).to">{{ link.label }}</a>
        </template>
      </nav>
    </header>

    <section class="hero">
      <div class="hero-copy">
        <div class="badge-line">
          <span v-for="badge in badges" :key="badge">{{ badge }}</span>
        </div>
        <h1>{{ cfg.title }}</h1>
        <p>{{ cfg.subtitle }}</p>
        <div class="hero-actions">
          <router-link v-if="isRouteLink(cfg.primary_cta?.to)" class="primary-btn" :to="cfg.primary_cta.to">
            {{ cfg.primary_cta.label }}
          </router-link>
          <a v-else class="primary-btn" :href="safeHref(cfg.primary_cta?.to)">{{ cfg.primary_cta?.label }}</a>
          <router-link v-if="isRouteLink(cfg.secondary_cta?.to)" class="secondary-btn" :to="cfg.secondary_cta.to">
            {{ cfg.secondary_cta.label }}
          </router-link>
          <a v-else class="secondary-btn" :href="safeHref(cfg.secondary_cta?.to)">{{ cfg.secondary_cta?.label }}</a>
        </div>
        <div class="trust-row">
          <span v-for="tag in trustTags" :key="tag">{{ tag }}</span>
        </div>
      </div>

      <div class="connect-card">
        <div class="connect-top">
          <span class="pulse"></span>
          <strong>节点状态</strong>
          <small>ONLINE</small>
        </div>
        <div class="route-map">
          <div class="node-point origin">
            <span>YOU</span>
          </div>
          <div class="route-line"></div>
          <div class="node-stack">
            <div v-for="node in heroNodes" :key="node.name" class="node-row">
              <span>{{ node.flag }}</span>
              <div>
                <strong>{{ node.name }}</strong>
                <small>{{ node.desc }}</small>
              </div>
              <em>{{ node.latency }}</em>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="selling-points" id="nodes">
      <article v-for="item in sellingPoints" :key="item.title">
        <span>{{ item.no }}</span>
        <h2>{{ item.title }}</h2>
        <p>{{ item.text }}</p>
      </article>
    </section>

    <section class="plans-section" id="plans">
      <div class="section-head">
        <p>VPN TRAFFIC PLANS</p>
        <h2>选择适合你的流量套餐</h2>
      </div>
      <div class="plan-grid">
        <article v-for="plan in plans" :key="plan.name" class="plan-card" :class="{ featured: plan.featured }">
          <div class="plan-head">
            <span>{{ plan.tag }}</span>
            <strong>{{ plan.name }}</strong>
          </div>
          <div class="price">
            <span>{{ plan.price }}</span>
            <small>{{ plan.unit }}</small>
          </div>
          <ul>
            <li v-for="feature in plan.features" :key="feature">{{ feature }}</li>
          </ul>
          <router-link to="/register">{{ plan.action }}</router-link>
        </article>
      </div>
    </section>

    <section class="use-cases">
      <div class="section-head">
        <p>USE CASES</p>
        <h2>为高频使用场景准备</h2>
      </div>
      <div class="case-grid">
        <div v-for="item in useCases" :key="item.title" class="case-card">
          <strong>{{ item.title }}</strong>
          <span>{{ item.text }}</span>
        </div>
      </div>
    </section>

    <section class="faq-section" id="faq">
      <div>
        <p>FAQ</p>
        <h2>购买前你可能关心的事</h2>
      </div>
      <div class="faq-list">
        <details v-for="item in faqs" :key="item.q">
          <summary>{{ item.q }}</summary>
          <p>{{ item.a }}</p>
        </details>
      </div>
    </section>

    <section class="final-cta">
      <h2>{{ cfg.final_cta.title }}</h2>
      <p>{{ cfg.final_cta.text }}</p>
      <router-link v-if="isRouteLink(cfg.final_cta.button_to)" class="primary-btn" :to="cfg.final_cta.button_to">
        {{ cfg.final_cta.button_label }}
      </router-link>
      <a v-else class="primary-btn" :href="safeHref(cfg.final_cta.button_to)">{{ cfg.final_cta.button_label }}</a>
    </section>

    <footer class="sales-footer">
      <span>{{ cfg.footer_text }}</span>
      <div>
        <template v-for="link in footerLinks" :key="`${link.label}-${link.to}`">
          <router-link v-if="isRouteLink(safeLink(link).to)" :to="safeLink(link).to">{{ link.label }}</router-link>
          <a v-else :href="safeLink(link).to">{{ link.label }}</a>
        </template>
      </div>
    </footer>
  </main>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { siteApi } from '@/api'

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
    { tag: 'STARTER', name: '轻量流量', price: '按套餐', unit: '灵活开通', action: '开始使用', features: ['适合临时访问和轻量使用', '标准节点订阅', '用户中心自助查看'] },
    { tag: 'POPULAR', name: '高速节点', price: '推荐', unit: '日常主力', action: '选择推荐', featured: true, features: ['适合 AI、办公和影音访问', '多线路订阅', '支持流量池计费'] },
    { tag: 'PRO', name: '大流量套餐', price: '长期', unit: '高频使用', action: '开通套餐', features: ['适合多设备和长期使用', '更多流量额度', '可配合兑换码续费'] },
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

const config = ref(defaultConfig)
const cfg = computed(() => ({ ...defaultConfig, ...config.value }))
const navLinks = computed(() => cfg.value.nav_links || defaultConfig.nav_links)
const badges = computed(() => cfg.value.badges || defaultConfig.badges)
const trustTags = computed(() => cfg.value.trust_tags || defaultConfig.trust_tags)
const heroNodes = computed(() => cfg.value.hero_nodes || defaultConfig.hero_nodes)
const sellingPoints = computed(() => cfg.value.selling_points || defaultConfig.selling_points)
const plans = computed(() => cfg.value.plans || defaultConfig.plans)
const useCases = computed(() => cfg.value.use_cases || defaultConfig.use_cases)
const faqs = computed(() => cfg.value.faqs || defaultConfig.faqs)
const footerLinks = computed(() => cfg.value.final_cta?.footer_links || defaultConfig.final_cta.footer_links)

function isRouteLink(to) {
  return typeof to === 'string' && to.startsWith('/')
}

function safeHref(to) {
  const value = String(to || '').trim()
  if (!value) return '#'
  if (value.startsWith('/') || value.startsWith('#')) return value
  try {
    const url = new URL(value)
    return ['http:', 'https:'].includes(url.protocol) ? value : '#'
  } catch {
    return '#'
  }
}

function safeLink(link) {
  return { ...link, to: safeHref(link?.to) }
}

async function fetchConfig() {
  try {
    const res = await siteApi.salesLanding()
    config.value = { ...defaultConfig, ...(res.data || {}) }
  } catch {
    config.value = defaultConfig
  }
}

onMounted(fetchConfig)
</script>

<style scoped>
.sales-page {
  min-height: 100vh;
  padding: 24px clamp(16px, 5vw, 72px) 34px;
  color: #edfaff;
  background:
    linear-gradient(rgba(66, 245, 255, 0.04) 1px, transparent 1px),
    linear-gradient(90deg, rgba(66, 245, 255, 0.035) 1px, transparent 1px),
    radial-gradient(circle at 16% 12%, rgba(66, 245, 255, 0.2), transparent 34rem),
    radial-gradient(circle at 82% 28%, rgba(255, 61, 242, 0.16), transparent 30rem),
    #050814;
  background-size: 34px 34px, 34px 34px, auto, auto;
}
.sales-nav,
.hero,
.selling-points,
.plans-section,
.use-cases,
.faq-section,
.final-cta,
.sales-footer {
  width: min(1160px, 100%);
  margin: 0 auto;
}
.sales-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 34px;
}
.brand,
.sales-nav nav,
.hero-actions,
.trust-row,
.sales-footer div {
  display: flex;
  align-items: center;
}
.brand {
  gap: 10px;
  color: #f6fdff;
  font-weight: 900;
  text-decoration: none;
}
.brand-mark {
  width: 40px;
  height: 40px;
  display: grid;
  place-items: center;
  color: #061019;
  border-radius: 8px;
  background: linear-gradient(135deg, var(--rp-cyan), var(--rp-pink));
  box-shadow: 0 0 24px rgba(66, 245, 255, 0.28);
}
.sales-nav nav {
  gap: 22px;
  font-size: 14px;
}
.sales-nav a,
.sales-footer a {
  color: #a9bed0;
  text-decoration: none;
}
.sales-nav a:hover,
.sales-footer a:hover {
  color: var(--rp-cyan);
}
.hero {
  display: grid;
  grid-template-columns: minmax(0, 1.05fr) minmax(330px, 0.95fr);
  gap: 48px;
  align-items: center;
  min-height: 620px;
}
.badge-line,
.trust-row {
  flex-wrap: wrap;
  gap: 8px;
}
.badge-line span,
.trust-row span {
  border: 1px solid rgba(92, 241, 255, 0.24);
  border-radius: 999px;
  color: #c4f8ff;
  background: rgba(9, 18, 33, 0.72);
}
.badge-line span {
  padding: 7px 11px;
  font-size: 12px;
}
.hero h1 {
  max-width: 780px;
  margin: 24px 0 0;
  font-size: clamp(48px, 7.4vw, 96px);
  line-height: 1.02;
  letter-spacing: 0;
}
.hero-copy p {
  max-width: 650px;
  margin: 26px 0 0;
  color: #b8cada;
  font-size: 18px;
  line-height: 1.8;
}
.hero-actions {
  gap: 12px;
  margin-top: 34px;
}
.primary-btn,
.secondary-btn,
.plan-card a {
  min-height: 44px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 20px;
  border-radius: 6px;
  font-weight: 800;
  text-decoration: none;
}
.primary-btn,
.plan-card a {
  color: #061019;
  background: linear-gradient(135deg, var(--rp-cyan), #6f7dff);
  box-shadow: 0 0 26px rgba(66, 245, 255, 0.28);
}
.secondary-btn {
  color: #e6f9ff;
  border: 1px solid rgba(92, 241, 255, 0.26);
  background: rgba(9, 18, 33, 0.72);
}
.trust-row {
  margin-top: 20px;
}
.trust-row span {
  padding: 6px 10px;
  font-size: 12px;
}
.connect-card,
.selling-points article,
.plan-card,
.case-card,
.faq-section,
.final-cta {
  border: 1px solid rgba(92, 241, 255, 0.2);
  border-radius: 8px;
  background: rgba(9, 16, 30, 0.76);
  box-shadow: 0 18px 52px rgba(0, 0, 0, 0.34), inset 0 1px 0 rgba(255, 255, 255, 0.04);
  backdrop-filter: blur(14px);
}
.connect-card {
  padding: 20px;
}
.connect-top {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-bottom: 18px;
  border-bottom: 1px solid rgba(92, 241, 255, 0.14);
}
.connect-top small {
  margin-left: auto;
  color: var(--rp-success);
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
}
.pulse {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: var(--rp-success);
  box-shadow: 0 0 18px var(--rp-success);
}
.route-map {
  position: relative;
  display: grid;
  grid-template-columns: 88px 1fr;
  gap: 22px;
  align-items: center;
  min-height: 360px;
}
.node-point {
  width: 76px;
  height: 76px;
  display: grid;
  place-items: center;
  border: 1px solid rgba(66, 245, 255, 0.38);
  border-radius: 50%;
  color: var(--rp-cyan);
  font-weight: 900;
  box-shadow: 0 0 34px rgba(66, 245, 255, 0.18);
}
.route-line {
  position: absolute;
  left: 78px;
  right: 34px;
  top: 50%;
  height: 1px;
  background: linear-gradient(90deg, var(--rp-cyan), rgba(255, 61, 242, 0.2));
}
.node-stack {
  display: grid;
  gap: 12px;
}
.node-row {
  position: relative;
  display: grid;
  grid-template-columns: 44px 1fr auto;
  gap: 12px;
  align-items: center;
  padding: 14px;
  border: 1px solid rgba(92, 241, 255, 0.17);
  border-radius: 8px;
  background: rgba(5, 12, 24, 0.78);
}
.node-row span {
  color: var(--rp-cyan);
  font-weight: 900;
}
.node-row strong,
.plan-head strong,
.case-card strong {
  display: block;
  color: #f5fcff;
}
.node-row small,
.case-card span {
  color: #94aabd;
}
.node-row em {
  color: var(--rp-success);
  font-style: normal;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
}
.selling-points {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 14px;
}
.selling-points article,
.plan-card,
.case-card {
  padding: 22px;
}
.selling-points span,
.section-head p,
.faq-section > div > p {
  margin: 0;
  color: var(--rp-cyan);
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 12px;
}
.selling-points h2 {
  margin: 18px 0 0;
  font-size: 22px;
}
.selling-points p,
.faq-section p,
.final-cta p {
  color: #a9bed0;
  line-height: 1.75;
}
.plans-section,
.use-cases,
.faq-section,
.final-cta {
  margin-top: 64px;
}
.section-head h2,
.faq-section h2,
.final-cta h2 {
  margin: 10px 0 0;
  color: #f5fcff;
  font-size: clamp(30px, 4vw, 46px);
}
.plan-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 14px;
  margin-top: 24px;
}
.plan-card {
  position: relative;
  min-height: 390px;
  display: flex;
  flex-direction: column;
}
.plan-card.featured {
  border-color: rgba(66, 245, 255, 0.48);
  box-shadow: 0 0 42px rgba(66, 245, 255, 0.14), 0 18px 52px rgba(0, 0, 0, 0.34);
}
.plan-head span {
  color: var(--rp-cyan);
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 12px;
}
.plan-head strong {
  margin-top: 8px;
  font-size: 24px;
}
.price {
  margin-top: 28px;
}
.price span {
  color: #f5fcff;
  font-size: 38px;
  font-weight: 900;
}
.price small {
  margin-left: 8px;
  color: #8fa7bd;
}
.plan-card ul {
  display: grid;
  gap: 10px;
  margin: 24px 0;
  padding: 0;
  list-style: none;
  color: #bdd0df;
}
.plan-card li::before {
  content: ">";
  margin-right: 8px;
  color: var(--rp-cyan);
}
.plan-card a {
  margin-top: auto;
}
.case-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 14px;
  margin-top: 24px;
}
.case-card {
  min-height: 150px;
}
.case-card span {
  display: block;
  margin-top: 14px;
  line-height: 1.7;
}
.faq-section {
  display: grid;
  grid-template-columns: 0.8fr 1.2fr;
  gap: 28px;
  padding: 28px;
}
.faq-list {
  display: grid;
  gap: 10px;
}
details {
  border-bottom: 1px dashed rgba(92, 241, 255, 0.16);
  padding-bottom: 14px;
}
summary {
  cursor: pointer;
  color: #f5fcff;
  font-weight: 800;
}
.final-cta {
  padding: 34px;
  text-align: center;
}
.sales-footer {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  margin-top: 48px;
  padding-top: 20px;
  border-top: 1px solid rgba(92, 241, 255, 0.16);
  color: #8fa7bd;
  font-size: 13px;
}
.sales-footer div {
  gap: 16px;
}
@media (max-width: 920px) {
  .sales-nav,
  .sales-footer {
    align-items: flex-start;
    flex-direction: column;
  }
  .sales-nav nav {
    flex-wrap: wrap;
  }
  .hero,
  .selling-points,
  .plan-grid,
  .case-grid,
  .faq-section {
    grid-template-columns: 1fr;
  }
  .hero {
    min-height: 0;
  }
  .route-map {
    min-height: 300px;
  }
}
@media (max-width: 560px) {
  .sales-page {
    padding-inline: 16px;
  }
  .hero-actions {
    align-items: stretch;
    flex-direction: column;
  }
  .hero-copy p {
    font-size: 16px;
  }
  .connect-card {
    padding: 14px;
  }
  .route-map {
    grid-template-columns: 1fr;
  }
  .route-line {
    display: none;
  }
  .node-point {
    width: 62px;
    height: 62px;
  }
  .node-row {
    grid-template-columns: 38px 1fr;
  }
  .node-row em {
    grid-column: 2;
  }
}
</style>
