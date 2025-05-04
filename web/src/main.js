// import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createLogto } from '@logto/vue'
import App from './App.vue'
import router from './router'

const config = {
  endpoint: 'https://3za1zw.logto.app/',
  appId: '418j3uxd3r0pgfcmprfam',
}

const app = createApp(App)

app.use(router)
app.use(createPinia())
app.use(createLogto, config)

app.mount('#app')
