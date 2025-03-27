// 修改前
// import Vue from 'vue'

// 修改后
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// 修改实例创建方式
const app = createApp(App)
app.use(router)
app.mount('#app')