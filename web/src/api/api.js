import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:5173', // 替换为你的后端 API 地址
  timeout: 5000, // 超时时间
})

// 封装 GET 请求
export const getUser = (userId) => api.get(`/user/${userId}`)

// 封装 POST 请求
export const login = (username, password) => api.post('/login', { username, password })

// 上传书籍文件
export const uploadBook = (file) => {
  const formData = new FormData();
  formData.append('file', file);
  return api.post('/api/upload_book', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
}

// 下载书籍文件
export const downloadBook = (bookName) => {
  return api.get('/api/download_book', {
    params: { book_name: bookName }
  });
}
export default api
