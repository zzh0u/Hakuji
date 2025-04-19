<template>
  <n-config-provider :theme-overrides="themeOverrides">
    <div class="home-container">
      <AppHeader />
      <SearchBox />
      <!-- <LogButton /> -->
      <BookCard />
    </div>
  </n-config-provider>
</template>

<script setup lang="ts">
import AppHeader from '../components/AppHeader.vue'
import BookCard from '../components/BookCard.vue'
import SearchBox from '../components/SearchBox.vue'
// import LogButton from '../components/LogButton.vue'

const themeOverrides = {

  common: {
    primaryColor: '#000000',
  },
  Button: {
    textColor: '#999999',
  },
  Select: {
    peers: {
      InternalSelection: {
        textColor: '#FF0000',
      },
    },
  },
}
</script>

<style scoped>
.home-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
</style>
