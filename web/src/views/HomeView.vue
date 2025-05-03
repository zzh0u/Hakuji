<template>
  <n-message-provider>
    <n-config-provider :theme-overrides="themeOverrides">
      <div class="home-container">
        <AppHeader />
        <SearchBox />
        <!-- <LogButton /> -->
        <BookCard />
      </div>
    </n-config-provider>
  </n-message-provider>
</template>

<script setup>
import { NConfigProvider, NMessageProvider } from 'naive-ui'
import AppHeader from '../components/AppHeader.vue'
import BookCard from '../components/BookCard.vue'
import SearchBox from '../components/SearchBox.vue'
// import LogButton from '../components/LogButton.vue'

const themeOverrides = {
  common: {
    primaryColor: '#1e90ff',
    primaryColorHover: '#4da6ff',
    primaryColorPressed: '#0077e6',
    primaryColorSuppl: '#4da6ff',

    infoColor: '#1e90ff',
    infoColorHover: '#4da6ff',
    infoColorPressed: '#0077e6',

    successColor: '#52c41a',
    warningColor: '#faad14',
    errorColor: '#f5222d',
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
