<template>
  <el-container class="user-layout">
    <el-header class="header">
      <div class="header-left">
        <router-link to="/" class="logo">
          <span>RP</span>
          RayPilot
        </router-link>
      </div>
      <div class="header-right">
        <router-link to="/dashboard">首页</router-link>
        <router-link to="/plans">套餐</router-link>
        <router-link to="/subscription">我的订阅</router-link>
        <router-link to="/redeem">兑换码</router-link>
        <router-link to="/profile">个人中心</router-link>
        <el-button text @click="handleLogout">退出</el-button>
      </div>
    </el-header>
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup>
// 用户侧布局。顶部导航栏 + 主内容区。
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus/es/components/message/index.mjs'
import { authApi } from '@/api'

const router = useRouter()
const userStore = useUserStore()

async function handleLogout() {
  try {
    await authApi.logout()
  } catch {
    // 忽略错误
  }
  userStore.logout()
  ElMessage.success('已退出登录')
  router.push('/login')
}
</script>

<style scoped>
.user-layout {
  min-height: 100vh;
  background:
    radial-gradient(circle at 20% 8%, rgba(66, 245, 255, 0.16), transparent 28rem),
    radial-gradient(circle at 88% 24%, rgba(255, 61, 242, 0.12), transparent 26rem),
    var(--rp-bg);
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
  border-bottom: 1px solid rgba(66, 245, 255, 0.14);
  background: rgba(8, 15, 28, 0.76);
  backdrop-filter: blur(16px);
  box-shadow: 0 10px 32px rgba(0, 0, 0, 0.22);
}
.header-left .logo {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  font-size: 20px;
  font-weight: 800;
  color: var(--rp-text);
  text-decoration: none;
}
.header-left .logo span {
  width: 34px;
  height: 34px;
  display: grid;
  place-items: center;
  color: #061019;
  font-size: 13px;
  border-radius: 8px;
  background: linear-gradient(135deg, var(--rp-cyan), var(--rp-pink));
  box-shadow: 0 0 18px rgba(66, 245, 255, 0.32);
}
.header-right {
  display: flex;
  align-items: center;
  gap: 4px;
}
.header-right a, .header-right .el-button {
  margin-left: 16px;
  color: var(--rp-muted);
  text-decoration: none;
  font-size: 14px;
}
.header-right a.router-link-active {
  color: var(--rp-cyan);
}
.el-main {
  padding: 22px;
}
</style>
