import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Login from '../views/Login.vue' // Importe ton Login
import Dashboard from '../views/Dashboard.vue'
import ClientPage from '../views/ClientPage.vue'
import ClientDetailsPage from '../views/ClientDetailsPage.vue'
import SettingPage from '../views/SettingPage.vue'

const routes: Array<RouteRecordRaw> = [
  { 
    path: '/', 
    name: 'login', 
    component: Login 
  },
  { 
    path: '/app',
    component: Dashboard, // Le Dashboard devient l'enveloppe
    children: [
      { path: 'clients', name: 'clients', component: ClientPage },
      { path: 'clients/:id', name: 'client-details', component: ClientDetailsPage },
      { path: 'settings', name: 'settings', component: SettingPage },
    ]
  },
  // Si on tape n'importe quoi, on retourne au login
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router