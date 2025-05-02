<template>
  <AppHeader />
  <div v-if="book" class="book-detail-header">
    <button class="back-btn" @click="goBack">
      <GoBackButton />
    </button>
    <span class="book-title">{{ book.title }}</span>
  </div>
  <div v-if="book" class="book-detail-container">
    <div class="cover-section">
      <img :src="book.cover" alt="封面" class="cover-img" />
    </div>
    <div class="info-section">
      <h2>{{ book.title }}</h2>
      <p><strong>作者：</strong>{{ book.author }}</p>
      <p><strong>分类：</strong>{{ book.category }}</p>
      <p><strong>出版社：</strong>{{ book.publisher }}</p>
      <p><strong>ISBN：</strong>{{ book.isbn }}</p>
      <p><strong>简介：</strong>{{ book.description }}</p>
    </div>
  </div>
  <div v-else>
    <p>未找到该书籍。</p>
  </div>
</template>

<script setup>
import AppHeader from '../components/AppHeader.vue'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import booksData from '../lib.json'
import GoBackButton from '@/components/GoBackButton.vue'

const route = useRoute()
const router = useRouter()
const bookId = Number(route.params.id)
const book = ref(null)

book.value = booksData.find((b) => b.id === bookId)

function goBack() {
  router.back()
}
</script>

<style scoped>
.book-detail-header {
  display: flex;
  align-items: center;
  max-width: 70rem;
  margin: 1.5rem auto 0.5rem auto;
}

.book-title {
  font-size: 1.2rem;
  font-weight: 600;
  color: #333;
}
.book-detail-container {
  max-width: 70rem;
  margin: 2rem auto;
  background: #fff;
  padding: 2rem;
  border-radius: 12px;
  display: flex;
  flex-direction: row;
  gap: 2rem;
}
.cover-section {
  flex: 1 1 20%;
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
.cover-img {
  width: 100%;
  max-width: 200px;
  height: 266px;
  object-fit: cover;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}
.info-section {
  flex: 1 1 80%;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}
h2 {
  margin-bottom: 1rem;
}
p {
  margin: 0.5rem 0;
}
</style>
