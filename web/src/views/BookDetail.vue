<template>
  <div v-if="book">
    <img
      :src="book.cover"
      alt="封面"
      style="width: 200px; height: 266px; object-fit: cover; margin-bottom: 1rem"
    />
    <h2>{{ book.title || '无书名' }}</h2>
    <p><strong>作者：</strong>{{ book.author }}</p>
    <p><strong>分类：</strong>{{ book.category }}</p>
    <p><strong>出版社：</strong>{{ book.publisher }}</p>
    <p><strong>ISBN：</strong>{{ book.isbn }}</p>
    <p><strong>简介：</strong>{{ book.description }}</p>
  </div>
  <div v-else>
    <p>未找到该书籍。</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import booksData from '../lib.json'

const route = useRoute()
const bookId = Number(route.params.id)
const book = ref(null)

book.value = booksData.find((b) => b.id === bookId)
</script>

<style scoped>
div {
  max-width: 600px;
  margin: 2rem auto;
  background: #fff;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.07);
}
h2 {
  margin-bottom: 1rem;
}
p {
  margin: 0.5rem 0;
}
</style>
