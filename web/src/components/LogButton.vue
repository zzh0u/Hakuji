<script>
import { useLogto } from '@logto/vue'
import { ref } from 'vue'

export default {
  setup() {
    const { signIn, signOut, isAuthenticated } = useLogto()
    const error = ref(null)

    const onClickSignIn = async () => {
      try {
        await signIn(import.meta.env.VITE_LOGTO_CALLBACK || 'http://localhost:5173/callback')
        error.value = null
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
      error,
      onClickSignIn,
      onClickSignOut,
    }
  },
}
</script>

<template>
  <div>
    <button v-if="!isAuthenticated" @click="onClickSignIn" class="auth-button">登陆</button>
    <button v-else @click="onClickSignOut" class="auth-button">登出</button>
  </div>
</template>
