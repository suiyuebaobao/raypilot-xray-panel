<template>
  <div class="login-page cyber-auth">
    <div class="auth-orbit"></div>
    <el-card class="login-card cyber-card">
      <div class="brand-lockup">
        <span class="brand-chip">RP</span>
        <div>
          <h2>用户登录</h2>
          <p>接入 RayPilot 订阅网络</p>
        </div>
      </div>
      <el-form :model="form" :rules="rules" ref="formRef" @submit.prevent="handleLogin">
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="用户名" prefix-icon="User" size="large" />
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" placeholder="密码" prefix-icon="Lock" show-password size="large" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin" :loading="loading" size="large" style="width: 100%">
            登录
          </el-button>
        </el-form-item>
      </el-form>
      <div class="link-row">
        <router-link to="/register">还没有账号？去注册</router-link>
      </div>
    </el-card>
  </div>
</template>

<script setup>
// 用户登录页。登录成功后保存 Access Token 到 Pinia。
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { authApi } from '@/api'
import { ElMessage } from 'element-plus/es/components/message/index.mjs'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const formRef = ref(null)
const loading = ref(false)

const form = reactive({
  username: '',
  password: '',
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}

async function handleLogin() {
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    const res = await authApi.login(form)
    // Alova 的 responded.onSuccess 已返回后端的 { success, message, code, data }
    userStore.setLogin(res.data.accessToken, res.data.user)
    ElMessage.success('登录成功')

    // 跳转回来源页面或首页
    const redirect = route.query.redirect || '/dashboard'
    router.push(redirect)
  } catch (err) {
    ElMessage.error(err.message || '登录失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  overflow: hidden;
  background:
    linear-gradient(rgba(66, 245, 255, 0.045) 1px, transparent 1px),
    linear-gradient(90deg, rgba(66, 245, 255, 0.04) 1px, transparent 1px),
    radial-gradient(circle at 24% 22%, rgba(66, 245, 255, 0.24), transparent 30rem),
    radial-gradient(circle at 78% 72%, rgba(255, 61, 242, 0.2), transparent 28rem),
    #050814;
  background-size: 34px 34px, 34px 34px, auto, auto;
}
.login-page::after {
  content: "";
  position: absolute;
  inset: 0;
  pointer-events: none;
  background: linear-gradient(transparent 50%, rgba(66, 245, 255, 0.025) 50%);
  background-size: 100% 4px;
}
.auth-orbit {
  position: absolute;
  width: 480px;
  height: 480px;
  border: 1px solid rgba(66, 245, 255, 0.14);
  border-radius: 50%;
  box-shadow: 0 0 80px rgba(66, 245, 255, 0.12), inset 0 0 80px rgba(255, 61, 242, 0.08);
}
.auth-orbit::before,
.auth-orbit::after {
  content: "";
  position: absolute;
  inset: 54px;
  border: 1px dashed rgba(255, 61, 242, 0.18);
  border-radius: 50%;
}
.auth-orbit::after {
  inset: 116px;
  border-color: rgba(66, 245, 255, 0.2);
}
.login-card {
  width: min(420px, calc(100vw - 32px));
  z-index: 1;
}
.brand-lockup {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 24px;
}
.brand-chip {
  width: 46px;
  height: 46px;
  display: grid;
  place-items: center;
  color: #061019;
  font-weight: 900;
  border-radius: 10px;
  background: linear-gradient(135deg, var(--rp-cyan), var(--rp-pink));
  box-shadow: 0 0 24px rgba(66, 245, 255, 0.36);
}
.login-card h2 {
  margin: 0;
  color: var(--rp-text);
}
.login-card p {
  margin: 4px 0 0;
  color: var(--rp-muted);
  font-size: 13px;
}
.link-row {
  text-align: center;
  margin-top: 12px;
}
.link-row a {
  color: #409eff;
  text-decoration: none;
}
</style>
