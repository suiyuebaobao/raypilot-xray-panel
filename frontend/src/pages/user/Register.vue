<template>
  <div class="login-page cyber-auth">
    <div class="auth-orbit"></div>
    <el-card class="login-card cyber-card">
      <div class="brand-lockup">
        <span class="brand-chip">RP</span>
        <div>
          <h2>用户注册</h2>
          <p>创建 RayPilot 订阅身份</p>
        </div>
      </div>
      <el-form :model="form" :rules="rules" ref="formRef" @submit.prevent="handleRegister">
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="用户名" prefix-icon="User" size="large" />
        </el-form-item>
        <el-form-item prop="email">
          <el-input v-model="form.email" placeholder="邮箱（可选）" prefix-icon="Message" size="large" />
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" placeholder="密码" prefix-icon="Lock" show-password size="large" />
        </el-form-item>
        <el-form-item prop="confirmPassword">
          <el-input v-model="form.confirmPassword" type="password" placeholder="确认密码" prefix-icon="Lock" show-password size="large" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister" :loading="loading" size="large" style="width: 100%">
            注册
          </el-button>
        </el-form-item>
      </el-form>
      <div class="link-row">
        <router-link to="/login">已有账号？去登录</router-link>
      </div>
    </el-card>
  </div>
</template>

<script setup>
// 用户注册页。
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { authApi } from '@/api'
import { ElMessage } from 'element-plus/es/components/message/index.mjs'

const router = useRouter()
const formRef = ref(null)
const loading = ref(false)

const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
})

const validateConfirmPassword = (rule, value, callback) => {
  if (value !== form.password) {
    callback(new Error('两次密码不一致'))
  } else {
    callback()
  }
}

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 32, message: '用户名长度 3-32 个字符', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少 6 个字符', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' },
  ],
}

async function handleRegister() {
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    await authApi.register({
      username: form.username,
      email: form.email || undefined,
      password: form.password,
    })
    ElMessage.success('注册成功，请登录')
    router.push('/login')
  } catch (err) {
    ElMessage.error(err.message || '注册失败')
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
    radial-gradient(circle at 22% 22%, rgba(66, 245, 255, 0.24), transparent 30rem),
    radial-gradient(circle at 80% 70%, rgba(255, 61, 242, 0.2), transparent 28rem),
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
  width: 520px;
  height: 520px;
  border: 1px solid rgba(66, 245, 255, 0.14);
  border-radius: 50%;
  box-shadow: 0 0 80px rgba(66, 245, 255, 0.12), inset 0 0 80px rgba(255, 61, 242, 0.08);
}
.auth-orbit::before,
.auth-orbit::after {
  content: "";
  position: absolute;
  inset: 58px;
  border: 1px dashed rgba(255, 61, 242, 0.18);
  border-radius: 50%;
}
.auth-orbit::after {
  inset: 126px;
  border-color: rgba(66, 245, 255, 0.2);
}
.login-card {
  width: min(430px, calc(100vw - 32px));
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
