import { createApp } from 'vue'
import { createRouter, createWebHashHistory} from 'vue-router'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import IndianFood from './components/IndianFood.vue'
import IndexPrincipal from './components/Index.vue'
import MexicanFood from './components/MexicanFood.vue'
import ChineseFood from './components/ChineseFood.vue'
loadFonts()

const routes =[
    {
      path: "/",
      name: "home",
      component: IndexPrincipal
    },
    {
      path: "/indian",
      name: "idian",
      component: IndianFood
    },
    {
      path: "/mexican",
      name: "mexican",
      component: MexicanFood
    },
    {
      path: "/chinese",
      name: "chinese",
      component: ChineseFood
    },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

const app = createApp(App)
app.use(router)
app.use(vuetify)
app.mount('#app')
