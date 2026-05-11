<template>
  <main class="landing-page">
    <header class="topbar">
      <router-link class="brand" to="/platform">
        <span class="brand-mark">RP</span>
        <span>RayPilot</span>
      </router-link>
      <nav class="topnav">
        <a href="#network">网络</a>
        <a href="#control">控制台</a>
        <a href="#operate">运维</a>
        <router-link to="/login">登录</router-link>
      </nav>
    </header>

    <section class="hero-shell">
      <div class="hero-copy">
        <p class="eyebrow">EDGE CONTROL / SUBSCRIPTION ROUTER</p>
        <h1>RayPilot</h1>
        <p class="summary">
          面向多出口节点、订阅分发、流量池计费和中转链路的统一控制平台。
        </p>
        <div class="cta-row">
          <router-link class="primary-link" to="/login">进入用户控制台</router-link>
          <router-link class="secondary-link" to="/">查看销售首页</router-link>
        </div>
      </div>

      <div class="terminal-card" aria-label="RayPilot status panel">
        <div class="terminal-head">
          <span></span>
          <span></span>
          <span></span>
          <strong>raypilot@edge:~</strong>
        </div>
        <pre>
╭─ network manifest
│  protocol: VLESS / Reality / XHTTP
│  relay:    TCP pass-through
│  billing:  normal + residential pools
│
├─ active lanes
│  HK relay  ──► US exit
│  direct    ──► multi transport
│
╰─ status: synchronized</pre>
      </div>
    </section>

    <section class="status-grid" id="network">
      <article v-for="item in statusItems" :key="item.label" class="status-item">
        <span>{{ item.code }}</span>
        <strong>{{ item.value }}</strong>
        <small>{{ item.label }}</small>
      </article>
    </section>

    <section class="content-grid" id="control">
      <article class="panel wide">
        <div class="panel-label">01 / CONTROL PLANE</div>
        <h2>把节点服务器当成可编排资源</h2>
        <p>
          出口 IP、传输模式、节点分组、中转后端和用户同步都在一个后台完成，
          适合持续扩容、迁移控制面和维护多地区线路。
        </p>
        <div class="line-list">
          <span>multi_exit agent</span>
          <span>node host inventory</span>
          <span>fallback center urls</span>
        </div>
      </article>

      <article class="panel">
        <div class="panel-label">02 / SUBSCRIPTION</div>
        <h2>订阅按能力生成</h2>
        <p>按套餐、流量池、节点状态和线路模式过滤，给用户下发可用线路。</p>
      </article>

      <article class="panel">
        <div class="panel-label">03 / LEDGER</div>
        <h2>双流量池计费</h2>
        <p>普通流量和家宽流量分池扣费，倍率实时生效，账本保留真实流量和扣费流量。</p>
      </article>
    </section>

    <section class="ops-section" id="operate">
      <div>
        <p class="eyebrow">OPERATIONS</p>
        <h2>为真实节点运维准备</h2>
      </div>
      <div class="ops-list">
        <div v-for="item in opsItems" :key="item.title" class="ops-row">
          <span>{{ item.no }}</span>
          <strong>{{ item.title }}</strong>
          <small>{{ item.text }}</small>
        </div>
      </div>
    </section>

    <footer class="landing-footer">
      <span>RayPilot / control the route, audit the flow.</span>
      <div>
        <router-link to="/register">创建账号</router-link>
        <router-link to="/login">用户登录</router-link>
        <router-link to="/">销售首页</router-link>
      </div>
    </footer>
  </main>
</template>

<script setup>
const statusItems = [
  { code: 'N01', value: 'multi-ip', label: '多出口识别' },
  { code: 'N02', value: 'xhttp', label: '多传输订阅' },
  { code: 'N03', value: 'relay', label: '中转链路' },
  { code: 'N04', value: 'ledger', label: '用户级账本' },
]

const opsItems = [
  { no: 'A', title: '一键部署', text: 'SSH 部署 agent、Xray/HAProxy、中心地址和节点记录。' },
  { no: 'B', title: '日志中心', text: '运行日志、部署日志、操作日志按时间检索。' },
  { no: 'C', title: '迁移容灾', text: 'agent 多中心入口 + SSH 兜底修复中心地址。' },
]
</script>

<style scoped>
.landing-page {
  min-height: 100vh;
  padding: 28px clamp(18px, 5vw, 72px) 34px;
  color: #dff8ff;
  background:
    linear-gradient(rgba(66, 245, 255, 0.045) 1px, transparent 1px),
    linear-gradient(90deg, rgba(66, 245, 255, 0.035) 1px, transparent 1px),
    radial-gradient(circle at 18% 18%, rgba(66, 245, 255, 0.18), transparent 32rem),
    radial-gradient(circle at 82% 34%, rgba(255, 61, 242, 0.16), transparent 30rem),
    #050814;
  background-size: 34px 34px, 34px 34px, auto, auto;
}
.topbar,
.hero-shell,
.status-grid,
.content-grid,
.ops-section,
.landing-footer {
  width: min(1120px, 100%);
  margin: 0 auto;
}
.topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 36px;
}
.brand,
.topnav,
.cta-row,
.landing-footer div {
  display: flex;
  align-items: center;
}
.brand {
  gap: 10px;
  color: #f2fbff;
  font-weight: 800;
  text-decoration: none;
}
.brand-mark {
  width: 38px;
  height: 38px;
  display: grid;
  place-items: center;
  color: #061019;
  border-radius: 8px;
  background: linear-gradient(135deg, var(--rp-cyan), var(--rp-pink));
}
.topnav {
  gap: 22px;
  font-size: 13px;
}
.topnav a,
.landing-footer a {
  color: #9fb6cb;
  text-decoration: none;
}
.topnav a:hover,
.landing-footer a:hover {
  color: var(--rp-cyan);
}
.hero-shell {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(320px, 460px);
  gap: 42px;
  align-items: center;
  min-height: 520px;
}
.eyebrow,
.panel-label {
  margin: 0 0 14px;
  color: var(--rp-cyan);
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 12px;
  letter-spacing: 0;
}
.hero-copy h1 {
  margin: 0;
  font-size: clamp(64px, 11vw, 136px);
  line-height: 0.88;
  letter-spacing: 0;
}
.summary {
  max-width: 560px;
  margin: 28px 0 0;
  color: #b7cadc;
  font-size: 18px;
  line-height: 1.8;
}
.cta-row {
  gap: 12px;
  margin-top: 32px;
}
.primary-link,
.secondary-link {
  min-height: 42px;
  display: inline-flex;
  align-items: center;
  padding: 0 18px;
  border-radius: 6px;
  font-weight: 700;
  text-decoration: none;
}
.primary-link {
  color: #061019;
  background: linear-gradient(135deg, var(--rp-cyan), #6f7dff);
  box-shadow: 0 0 26px rgba(66, 245, 255, 0.26);
}
.secondary-link {
  color: #dff8ff;
  border: 1px solid rgba(92, 241, 255, 0.28);
  background: rgba(10, 19, 34, 0.72);
}
.terminal-card,
.panel,
.status-item,
.ops-section {
  border: 1px solid rgba(92, 241, 255, 0.2);
  background: rgba(9, 16, 30, 0.74);
  box-shadow: 0 18px 52px rgba(0, 0, 0, 0.34), inset 0 1px 0 rgba(255, 255, 255, 0.04);
  backdrop-filter: blur(14px);
}
.terminal-card {
  overflow: hidden;
  border-radius: 8px;
}
.terminal-head {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 14px;
  border-bottom: 1px solid rgba(92, 241, 255, 0.16);
  color: #8fa7bd;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 12px;
}
.terminal-head span {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: var(--rp-cyan);
}
.terminal-head span:nth-child(2) {
  background: var(--rp-amber);
}
.terminal-head span:nth-child(3) {
  background: var(--rp-pink);
}
.terminal-head strong {
  margin-left: auto;
  font-weight: 600;
}
.terminal-card pre {
  margin: 0;
  padding: 22px;
  overflow: auto;
  color: #c7f8ff;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 13px;
  line-height: 1.7;
}
.status-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 12px;
  margin-top: 10px;
}
.status-item {
  padding: 16px;
  border-radius: 8px;
}
.status-item span,
.ops-row span {
  color: var(--rp-cyan);
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 12px;
}
.status-item strong {
  display: block;
  margin-top: 16px;
  color: #f2fbff;
  font-size: 22px;
}
.status-item small {
  display: block;
  margin-top: 6px;
  color: #8fa7bd;
}
.content-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 14px;
  margin-top: 56px;
}
.panel {
  min-height: 230px;
  padding: 24px;
  border-radius: 8px;
}
.panel.wide {
  grid-column: span 2;
}
.panel h2,
.ops-section h2 {
  margin: 0;
  color: #f2fbff;
  font-size: 28px;
}
.panel p {
  margin: 16px 0 0;
  color: #aec3d5;
  line-height: 1.8;
}
.line-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 22px;
}
.line-list span {
  padding: 6px 10px;
  border: 1px solid rgba(92, 241, 255, 0.22);
  border-radius: 999px;
  color: #bdf8ff;
  font-size: 12px;
}
.ops-section {
  display: grid;
  grid-template-columns: 0.8fr 1.2fr;
  gap: 28px;
  margin-top: 14px;
  padding: 28px;
  border-radius: 8px;
}
.ops-list {
  display: grid;
  gap: 10px;
}
.ops-row {
  display: grid;
  grid-template-columns: 32px 150px 1fr;
  gap: 12px;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px dashed rgba(92, 241, 255, 0.16);
}
.ops-row:last-child {
  border-bottom: 0;
}
.ops-row strong {
  color: #f2fbff;
}
.ops-row small {
  color: #9fb6cb;
  line-height: 1.6;
}
.landing-footer {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  margin-top: 48px;
  padding-top: 20px;
  border-top: 1px solid rgba(92, 241, 255, 0.16);
  color: #8fa7bd;
  font-size: 13px;
}
.landing-footer div {
  gap: 16px;
}
@media (max-width: 900px) {
  .topbar,
  .landing-footer {
    align-items: flex-start;
    flex-direction: column;
  }
  .topnav {
    flex-wrap: wrap;
  }
  .hero-shell,
  .content-grid,
  .ops-section {
    grid-template-columns: 1fr;
  }
  .hero-shell {
    min-height: 0;
  }
  .panel.wide {
    grid-column: auto;
  }
  .status-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
  .ops-row {
    grid-template-columns: 28px 1fr;
  }
  .ops-row small {
    grid-column: 2;
  }
}
@media (max-width: 560px) {
  .landing-page {
    padding-inline: 16px;
  }
  .status-grid {
    grid-template-columns: 1fr;
  }
  .cta-row {
    align-items: stretch;
    flex-direction: column;
  }
  .primary-link,
  .secondary-link {
    justify-content: center;
  }
}
</style>
