import { createApp } from 'vue'
import App from './App.vue'
import router from './router/main' // <---
import './index.css'

createApp(App).use(router).mount('#app')