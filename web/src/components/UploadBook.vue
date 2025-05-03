<template>
  <div class="upload-book">
    <input type="file" @change="onFileChange" />
    <button @click="handleUpload" :disabled="!selectedFile">上传书籍</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { uploadBook } from '../api/api'

const emit = defineEmits(['uploaded'])
const selectedFile = ref(null)

function onFileChange(e) {
  selectedFile.value = e.target.files[0]
}

async function handleUpload() {
  if (!selectedFile.value) return
  try {
    const res = await uploadBook(selectedFile.value)
    alert(res.data.message || '上传成功')
    selectedFile.value = null
    emit('uploaded')
  } catch (err) {
    alert(err.response?.data?.error || '上传失败')
  }
}
</script>

<style scoped>
.upload-book {
  margin-top: 2rem;
}
button[disabled] {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
