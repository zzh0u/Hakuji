<script setup>
import { ref } from 'vue'
import LogButton from './LogButton.vue'

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
      <div class="user-avatar">
        <i class="fas fa-user-circle"></i>
      </div>
      <div class="user-info">
        <h3>用户名</h3>
        <p>邮箱</p>
      </div>
      <button class="edit-btn">编辑信息</button>
    </div>

    <div v-else class="logged-out-content">
      <div class="default-avatar">
        <i class="fas fa-user-circle"></i>
      </div>
      <!-- <button class="login-btn">登陆</button> -->
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
    background: none;
    border: none;
    cursor: pointer;
    color: #666;
  }

  &:hover {
    color: #333;
  }
}

.logged-in-content,
.logged-out-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;

  .user-avatar,
  .default-avatar {
    font-size: 3rem;
    color: #1e90ff;
  }

  .user-info {
    text-align: center;
    h3 {
      margin: 0;
      color: #333;
    }
    p {
      margin: 4px 0 0;
      font-size: 0.9rem;
      color: #666;
    }
  }

  .edit-btn {
    width: 100%;
    padding: 8px;
    background: #1e90ff;
    border: none;
    border-radius: 4px;
    color: white;
    cursor: pointer;
    transition: all 0.2s ease-in-out;
  }
}

.login-btn {
  padding: 8px 16px;
  background: #1e90ff;
  border: none;
  border-radius: 4px;
  color: white;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
}
</style>
