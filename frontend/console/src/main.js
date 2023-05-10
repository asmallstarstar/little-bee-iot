import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/main.css'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as Icons from '@element-plus/icons-vue'
import vueI18n from '@/lang/index'
import piniaPersist from 'pinia-plugin-persist'
import './permission'

const app = createApp(App)
const pinia = createPinia()
pinia.use(() => ({ router }))
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)
app.use(ElementPlus)
app.use(vueI18n)

Object.keys(Icons).forEach(key => {
    app.component(key, Icons[key])
})

app.mount('#app')
