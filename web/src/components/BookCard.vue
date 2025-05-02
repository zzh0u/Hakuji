<template>
  <div class="book-list">
    <!-- <div class="book-card" v-for="bookItem in recentBooks" :key="bookItem.id"> -->
    <div
      class="book-card"
      v-for="bookItem in recentBooks"
      :key="bookItem.id"
      @click="goToDetail(bookItem.id)"
      style="cursor: pointer"
    >
      <img :src="bookItem.cover" alt="书籍封面" class="cover" />
      <!-- <div class="content">
        <h3 class="title">{{ bookItem.title }}</h3>
        <p class="author">{{ bookItem.author }}</p>
        <p class="description">{{ bookItem.description }}</p>
      </div> -->
    </div>
    <div style="margin-top: 2rem">
      <input type="file" @change="onFileChange" />
      <button @click="handleUpload" :disabled="!selectedFile">上传书籍</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import booksData from '../lib.json'
import { uploadBook } from '../api/api'
import { useRouter } from 'vue-router'

const recentBooks = ref(booksData)
const selectedFile = ref(null)

function onFileChange(e) {
  selectedFile.value = e.target.files[0]
}

const router = useRouter()

function goToDetail(id) {
  router.push({ name: 'book-detail', params: { id } })
}

async function handleUpload() {
  if (!selectedFile.value) return
  try {
    const res = await uploadBook(selectedFile.value)
    alert(res.data.message || '上传成功')
    selectedFile.value = null
  } catch (err) {
    alert(err.response?.data?.error || '上传失败')
  }
}
</script>

<style scoped>
.book-list {
  max-width: 70rem;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 24px;
  /* border: 1px solid #eee;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1); */
}

.book-card {
  /* border: 1px solid #eee; */
  padding: 0;
  aspect-ratio: 3/4;
  display: block;
  transition: transform 0.2s ease-in-out;
  background: white;
}

.book-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translate(-2px);
}

.cover {
  width: 200px;
  height: 266px;
  object-fit: cover;
}
</style>
