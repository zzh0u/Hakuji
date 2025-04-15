// import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createLogto } from '@logto/vue'
import App from './App.vue'

const config = {
  endpoint: 'https://3za1zw.logto.app/',
  appId: 'ja904h6clsvvl0vllpgyq',
}

const app = createApp(App)

app.use(createPinia())
app.use(createLogto, config)

app.mount('#app')
