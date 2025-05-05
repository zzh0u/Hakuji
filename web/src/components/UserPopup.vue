<script setup>
// import { ref } from 'vue'
import LogButton from './LogButton.vue'
import UploadBook from './UploadBook.vue'

const props = defineProps({
  isLoggedIn: {
    type: Boolean,
    default: false,
  },
  showUserPopup: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['close'])

const closePopup = () => {
  emit('close')
}
</script>

<template>
  <div v-if="showUserPopup" class="user-popup">
    <button class="close-btn" @click="closePopup">
      <i class="fas fa-times"></i>
    </button>

    <div v-if="isLoggedIn" class="logged-in-content">
      <div class="default-avatar">
        <i class="fas fa-user-circle"></i>
      </div>
      <LogButton />
      <UploadBook @uploaded="() => window.location.reload()" />
    </div>

    <div v-else class="logged-out-content">
      <div class="default-avatar">
        <i class="fas fa-user-circle"></i>
      </div>
      <LogButton />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.user-popup {
  position: absolute;
  top: 64px;
  right: 75px;
  width: 320px;
  height: 435px;
  background-color: #f3eeee;
  // border: #1e90ff solid 1px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 1000;

  .close-btn {
    position: absolute;
    top: 10px;
    right: 10px;
    font-size: 1.5rem;
    background: none;
    border: none;
    cursor: pointer;
    color: #666;
  }
}

.logged-in-content,
.logged-out-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 80%;
  gap: 20px;
  padding: 100px;

  .default-avatar {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    font-size: 6rem;
    color: #1e90ff;
  }
}
</style>
