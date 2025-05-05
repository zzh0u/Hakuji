<script>
import router from '@/router'
import { useLogto } from '@logto/vue'
import { ref, watchEffect } from 'vue'

export default {
  setup() {
    const { signIn, signOut, isAuthenticated, fetchUserInfo } = useLogto()
    const error = ref(null)
    const userInfo = ref(null)

    watchEffect(async () => {
      if (isAuthenticated.value) {
        userInfo.value = await fetchUserInfo()
      } else {
        userInfo.value = null
      }
    })
    const onClickSignIn = async () => {
      try {
        await signIn(import.meta.env.VITE_LOGTO_CALLBACK || 'http://localhost:5173/callback')
        router.push('/')
      } catch (err) {
        error.value = err.message
      }
    }

    const onClickSignOut = async () => {
      try {
        await signOut(import.meta.env.VITE_LOGTO_CALLBACK || 'http://localhost:5173')
        error.value = null
      } catch (err) {
        error.value = err.message
      }
    }
    return {
      isAuthenticated,
      userInfo,
      error,
      onClickSignIn,
      onClickSignOut,
    }
  },
}
</script>

<template>
  <div v-if="error" class="error-message">
    {{ error }}
  </div>
  <div v-if="isAuthenticated && userInfo">
    <!-- <h2>用户信息</h2> -->
    <!-- <p>ID:{{ userInfo.sub }}</p> -->
    <p>用户名：{{ userInfo.username }}</p>
    <!-- <p>邮箱:{{ userInfo.email }}</p> -->
    <button @click="onClickSignOut" class="auth-button">登出</button>
  </div>
  <div v-else>
    <button @click="onClickSignIn" class="auth-button">登录</button>
  </div>
</template>

<style>
.auth-button {
  padding: 10px 20px;
  background-color: #1e90ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>
