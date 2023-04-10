import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { globalRegister } from './global'
import 'element-plus/theme-chalk/src/message.scss'

const app = createApp(App)
app.use(globalRegister)
app.use(router)
app.use(createPinia())
app.mount('#app')
