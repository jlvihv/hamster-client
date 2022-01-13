import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import './assets/css/app.scss'
import { VueClipboard } from '@soerenmartius/vue3-clipboard'

// import wails runtime
import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
  createApp(App)
      .use(Antd)
      .use(store)
      .use(router)
      .use(VueClipboard)
      .mount('#app')
})
