<!--
这个文件是 Vue 组件的模板文件，用于显示登录回调页面。
首先，我们导入了 handleSignInCallback 函数和 router 对象。handleSignInCallback 函数用于处理登录回调，router 对象用于导航到其他页面。
然后，我们使用 handleSignInCallback 函数处理登录回调，并将返回的 isLoading 和 error 对象解构出来。
接下来，我们使用 v-if 指令根据 isLoading 和 error 对象的状态来显示不同的内容。如果 isLoading 为 true，我们显示一个加载动画和一个提示信息；如果 error 不为 null，我们显示一个错误信息。
-->
<script setup>
import { useHandleSignInCallback } from '@logto/vue';
import router from '@/router';

const { isLoading } = useHandleSignInCallback(()=>{
  router.push('/');
});
</script>

<template>
  <div>
    <p v-if="isLoading">重定向中...</p>
  </div>
</template>

<style scoped>
</style>
