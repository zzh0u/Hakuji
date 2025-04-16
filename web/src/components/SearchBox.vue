<!--
  搜索框光标和placeholder位置不一样
  搜索按钮颜色不对
-->
<template>
  <div class="search-box">
    <n-input
      v-model:value="searchQuery"
      type="text"
      placeholder="按书名、作者、ISBN、出版社等搜索..."
      clearable
      @keyup.enter="handleSearch"
      style="width: 70rem"
      class="search-input"
    >
    </n-input>
    <n-button type="primary" @click="handleSearch" class="search-button"> 搜索 </n-button>
  </div>
</template>

<script>
import { NInput, NButton } from 'naive-ui'

export default {
  name: 'SearchBox',
  components: {
    NInput,
    NButton,
  },
  data() {
    return {
      searchQuery: '',
    }
  },
  methods: {
    handleSearch() {
      this.$emit('search', this.searchQuery)
    },
  },
}
</script>

<style scoped>
.search-box {
  display: flex;
  width: 100%;
  max-width: 70rem;
  height: 25rem;
  margin: auto;
  /* margin-bottom: 20px; */
  border-radius: 5px 0 0 5px;
  justify-content: center;
  align-items: center;
}

.search-input {
  padding: 10px;
  margin-right: 0;
  font-size: 18px;
  height: 3rem;
  border-radius: 5px 0 0 5px;
}

.search-button {
  height: 3rem;
  width: 12rem;
  padding: 10px 20px;
  font-size: 18px;
  border-radius: 0 5px 5px 0;
  cursor: pointer;
  hover {
    background-color: #4caf50;
  }
}
</style>
