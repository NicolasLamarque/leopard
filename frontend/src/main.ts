// ========================================
// frontend/src/main.js
// ========================================
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { plugin, defaultConfig } from '@formkit/vue'
import App from './App.vue'
import router from './router' // 1. Importe ton fichier router/index.ts

const app = createApp(App)

app.use(createPinia())
app.use(router) 
app.use(plugin, defaultConfig)

app.mount('#app')