// todo:将config放到文件中读取
import { createLogto, LogtoConfig } from '@logto/vue'
import { createApp } from 'vue'
import App from './App.vue'

const config: LogtoConfig = {
  endpoint: 'https://3za1zw.logto.app/',
  appId: 'ja904h6clsvvl0vllpgyq',
};

const app = createApp(App)

app.use(createLogto, config)
app.mount('#app')
