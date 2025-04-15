<script setup>
const openSearchLine = () => {
  // Todo:搜索功能实现
  console.log('打开搜索')
}
</script>

<template>
  <header class="header container">
    <div class="logo">
      <a href="/">
        <img src="/favicon.ico" />
        <h1>hakuji library</h1>
      </a>
    </div>
    <nav class="nav">
      <button class="naviSearch" @click="openSearchLine()">search(换成搜索 logo)</button>

      <div @click="openHeadUser()" class="user-button">
        <!-- 换成默认用户头像 -->
        <img src="/favicon.ico" alt="logo" />
        <span>用户</span>
      </div>

      <div>
        <!-- 选项菜单 -->
        <img src="/favicon.ico" alt="logo" />
        <span>选项</span>
      </div>
    </nav>

    <!-- 用户信息悬浮窗 -->
  </header>
</template>

<style lang="scss" scoped>
.header {
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

    img {
      max-height: 48px;
    }

    h1 {
      margin: 0;
      color: #fff;
      font-size: 1.5rem;
      font-weight: 800;
    }
  }

  .nav {
    display: flex;
    align-items: center;
    gap: 1.5rem;
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
  padding: 8px, 12px;
  background: #f5f5f5;
  border-radius: 4px;
  border: 1px solid #ddd;
  cursor: pointer;
}
</style>
