<template>
  <div class="home-container">
    <AppHeader />
    <SearchBox />
    <!-- <LogButton /> -->
    <BookCard />
  </div>
</template>

<script setup lang="ts">
import AppHeader from '../components/AppHeader.vue'
import BookCard from '../components/BookCard.vue'
import SearchBox from '../components/SearchBox.vue'
// import LogButton from '../components/LogButton.vue'
</script>

<style scoped>
.home-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
</style>
