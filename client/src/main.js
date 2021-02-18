import { createApp } from 'vue'
import App from './App.vue'
import router from './router/main' 
import store from './store/main'
import './index.css'

createApp(App).use(router).use(store).mount('#app')