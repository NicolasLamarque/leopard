import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
import HomePage from '../views/homePage.vue'
import ClientPage from '../views/ClientPage.vue'
import ClientDetailsPage from '../views/ClientDetailsPage.vue'
import SettingPage from '../views/SettingPage.vue'
import RPAPage from '../views/RPAPage.vue'

const routes: Array<RouteRecordRaw> = [
  { 
    path: '/', 
    name: 'login', 
    component: Login 
  },
  { 
    path: '/app',
    component: Dashboard,
    children: [
      // 🏠 Page d'accueil par défaut
      { 
        path: '', 
        name: 'home', 
        component: HomePage 
      },
      
      // 👥 Gestion des clients
      { 
        path: 'clients', 
        name: 'clients', 
        component: ClientPage 
      },
      { 
        path: 'clients/:id', 
        name: 'client-details', 
        component: ClientDetailsPage 
      },
      
      // 🩺 Gestion des médecins
      { 
        path: 'medecins', 
        name: 'medecins', 
        component: () => import('../views/MedecinPage.vue')
      },
      
      // 📋 Autres sections (lazy loading)
      { 
        path: 'notaires', 
        name: 'notaires', 
        component: () => import('../views/NotairePage.vue')
      },
      { 
        path: 'professionnels', 
        name: 'professionnels', 
        component: () => import('../views/ProfessionnelPage.vue')
      },
      { 
        path: 'etablissements', 
        name: 'etablissements', 
        component: () => import('../views/EtablissementPage.vue')
      },
      // 🤖 Gestion RPA
      {
        path: 'rpa',
        name: 'rpa',
        component: RPAPage
      },
      { 
        path: 'rapports', 
        name: 'rapports', 
        component: () => import('../views/RapportPage.vue')
      },
      
      // ⚙️ Paramètres
      { 
        path: 'settings', 
        name: 'settings', 
        component: SettingPage 
      }
    ]
  },
  
  // Redirection pour les routes inconnues
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router