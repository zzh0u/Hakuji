<script setup>
import '@fortawesome/fontawesome-free/css/all.min.css'
import { ref } from 'vue'
import UserPopup from './UserPopup.vue'

const isLoggedIn = ref(false)
const showUserPopup = ref(false)
const showMenu = ref(false)

const openSearchLine = () => {
  // Todo:搜索功能实现
  console.log('打开搜索')
}

const toggleUserPopup = () => {
  showUserPopup.value = !showUserPopup.value
}

const toggleMenu = () => {
  showMenu.value = !showMenu.value
}
</script>

<template>
  <header class="header-container">
    <div class="logo">
      <a href="/">
        <img src="/favicon.ico" />
        <h1>hakuji library</h1>
      </a>
    </div>

    <!-- 顶部右侧导航栏 -->
    <nav class="nav">
      <!-- 搜索按钮 -->
      <button class="naviSearch" @click="openSearchLine()">
        <i class="fas fa-search"></i>
      </button>

      <!-- 用户按钮 -->
      <div class="user-button" @click="toggleUserPopup">
        <i class="fas fa-user-circle"></i>
      </div>

      <!-- 用户信息悬浮窗组件 -->
      <UserPopup
        :is-logged-in="isLoggedIn"
        :show-user-popup="showUserPopup"
        @close="toggleUserPopup"
      />

      <!-- 菜单按钮 -->
      <div class="menu-button" @click="toggleMenu">
        <i class="fas fa-bars"></i>
      </div>
    </nav>
  </header>
</template>

<style lang="scss" scoped>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
  background-image: linear-gradient(to right, #87ceeb, #1e90ff);

  .logo {
    margin: 0 10px 0;
    display: flex;
    align-items: center;

    a {
      display: flex;
      align-items: center;
      gap: 8px;
      text-decoration: none;
    }
  }

  .nav {
    display: flex;
    align-items: center;
    gap: 2rem;
    padding-right: 2rem;
    height: 64px;
    margin: 0 10px 0;

    > a:first-child {
      margin-right: auto;
    }
    > *:not(a:first-child) {
      display: flex;
      align-items: center;
    }
  }
}

.naviSearch {
  padding: 8px 12px;
  background: none;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
}

.fas {
  font-size: 1.5rem;
  color: #fff;
  cursor: pointer;
  transition: all 0.2s ease-in-out;

  &:hover {
    opacity: 0.6;
  }
}

.user-button,
.menu-button {
  position: relative;
  cursor: pointer;
  transition: all 0.2s ease-in-out;

  &:hover {
    opacity: 0.6;
  }
}
</style>
