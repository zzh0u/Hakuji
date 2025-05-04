<template>
  <div class="upload-book">
    <button @click="showModal = true" class="upload-btn">上传书籍</button>

    <n-modal v-model:show="showModal" style="width: 80%; max-width: 800px;">
      <n-card title="上传书籍" :bordered="false" size="huge" role="dialog" aria-modal="true">
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          label-placement="left"
          label-width="100px"
          require-mark-placement="right-hanging"
        >
          <n-form-item label="书籍文件" path="file" required>
            <n-upload
              v-model:file-list="fileList"
              @change="handleFileChange"
              :max="1"
              :accept="'.pdf,.epub,.mobi'"
            >
              <n-button>选择文件</n-button>
            </n-upload>
          </n-form-item>
          <n-form-item label="ISBN" path="isbn" required>
            <n-input v-model:value="formValue.isbn" placeholder="请输入ISBN" />
          </n-form-item>
          <n-form-item label="书名" path="title" required>
            <n-input v-model:value="formValue.title" placeholder="自动使用文件名" readonly />
          </n-form-item>
          <n-form-item label="作者" path="author" required>
            <n-input v-model:value="formValue.author" placeholder="请输入作者" />
          </n-form-item>
          <n-form-item label="封面图片" path="coverFile">
            <n-upload
              v-model:file-list="coverFileList"
              @change="handleCoverChange"
              :max="1"
              :accept="'image/*'"
              list-type="image-card"
            >
              选择封面
            </n-upload>
          </n-form-item>
          <n-form-item label="出版社" path="publisher">
            <n-input v-model:value="formValue.publisher" placeholder="请输入出版社" />
          </n-form-item>
          <n-form-item label="出版日期" path="publishedDate">
            <n-date-picker v-model:value="formValue.publishedDate" type="date" />
          </n-form-item>
          <n-form-item label="分类" path="category">
            <n-input v-model:value="formValue.category" placeholder="请输入分类" />
          </n-form-item>
          <n-form-item label="内容简介" path="contentSummary">
            <n-input
              v-model:value="formValue.contentSummary"
              type="textarea"
              placeholder="请输入内容简介"
              :autosize="{ minRows: 3, maxRows: 5 }"
            />
          </n-form-item>
        </n-form>
        <template #footer>
          <div style="display: flex; justify-content: flex-end; gap: 12px;">
            <n-button @click="showModal = false">取消</n-button>
            <n-button type="primary" @click="handleUpload" :disabled="!formValue.file || !formValue.isbn || !formValue.author">
              上传
            </n-button>
          </div>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { uploadBook } from '../api/api'
import {
  NModal,
  NCard,
  NForm,
  NFormItem,
  NInput,
  NButton,
  NUpload,
  NDatePicker,
  useMessage
} from 'naive-ui'

const emit = defineEmits(['uploaded'])
const message = useMessage()
const showModal = ref(false)
const fileList = ref([])
const coverFileList = ref([])
const formRef = ref(null)

const formValue = reactive({
  file: null,
  coverFile: null,
  isbn: '',
  title: '',
  author: '',
  coverURL: '',
  publisher: '',
  publishedDate: null,
  category: '',
  contentSummary: ''
})

const rules = {
  isbn: {
    required: true,
    message: '请输入ISBN',
    trigger: 'blur'
  },
  author: {
    required: true,
    message: '请输入作者',
    trigger: 'blur'
  },
  file: {
    required: true,
    message: '请选择文件',
    trigger: 'change',
    validator: (rule, value) => !!formValue.file
  }
}

function handleFileChange(options) {
  if (options.fileList.length > 0) {
    formValue.file = options.fileList[0].file
    // 使用文件名（包括后缀）作为书名
    formValue.title = options.fileList[0].file.name
  } else {
    formValue.file = null
    formValue.title = ''
  }
}

function handleCoverChange(options) {
  if (options.fileList.length > 0) {
    formValue.coverFile = options.fileList[0].file
  } else {
    formValue.coverFile = null
  }
}

async function handleUpload() {
  formRef.value?.validate(async (errors) => {
    if (errors) {
      message.error('请填写必填项')
      return
    }

    if (!formValue.file) {
      message.error('请选择文件')
      return
    }

    try {
      const formData = new FormData()
      formData.append('file', formValue.file)
      formData.append('coverFile', formValue.coverFile || '')

      // 添加书籍信息
      formData.append('isbn', formValue.isbn)
      formData.append('title', formValue.title)
      formData.append('author', formValue.author)
      formData.append('publisher', formValue.publisher || '')
      formData.append('publishedDate', formValue.publishedDate ? new Date(formValue.publishedDate).toISOString() : '')
      formData.append('category', formValue.category || '')
      formData.append('contentSummary', formValue.contentSummary || '')

      const res = await uploadBook(formData)
      message.success(res.data.message || '上传成功')
      showModal.value = false
      fileList.value = []
      coverFileList.value = []
      resetForm()
      emit('uploaded')
    } catch (err) {
      message.error(err.response?.data?.error || '上传失败')
    }
  })
}

function resetForm() {
  Object.keys(formValue).forEach(key => {
    if (key === 'publishedDate') {
      formValue[key] = null
    } else {
      formValue[key] = ''
    }
  })
  formValue.file = null
  formValue.coverFile = null
}
</script>

<style scoped>
.upload-book {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 2rem;
}

.upload-btn {
  padding: 0.75rem 1.5rem;
  background-color: #1e90ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.3s;
}

.upload-btn:hover {
  background-color: #0077e6;
}

button[disabled] {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
