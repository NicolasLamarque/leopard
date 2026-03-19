import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Login from "../views/Login.vue";
import Dashboard from "../views/Dashboard.vue";
import HomePage from "../views/homePage.vue";
import ClientPage from "../views/ClientPage.vue";
import ClientDetailsPage from "../views/ClientDetailsPage.vue";
import SettingPage from "../views/SettingPage.vue";
import RPAPage from "../views/etablissements/RPAPage.vue";
import PharmaciesPage from "@/views/PharmaciesPage.vue";
import MedecinPage from "../views/MedecinPage.vue";
import NotairePage from "../views/NotairePage.vue";
import IntervenantPage from "../views/IntervenantPage.vue";
import EtablissementPage from "../views/etablissements/EtablissementPage.vue";
import RevenusPage from "../views/RevenusPage.vue";
import ServicePage from "../views/ServicePage.vue";
import { R } from "vue-router/dist/router-CWoNjPRp.mjs";
import { useAuthStore } from '@/stores/auth'
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "login",
    component: Login,
  },
  {
    path: "/app",
    component: Dashboard,
    props: true,
    children: [
      // 🏠 Page d'accueil par défaut
      {
        path: "",
        name: "home",
        component: HomePage,
      },

      // 👥 Gestion des clients
      {
        path: "clients",
        name: "clients",
        component: ClientPage,
      },
      {
        path: "clients/:id",
        name: "client-details",
        component: ClientDetailsPage,
      },

      // 🩺 Gestion des médecins
      {
        path: "medecins",
        name: "medecins",
        component: MedecinPage,
      },

      // 📋 Autres sections (lazy loading)
      {
        path: "notaires",
        name: "notaires",
        component: NotairePage,
      },
      {
        path: "intervenants",
        name: "intervenants",
        component: IntervenantPage,
      },

      // 🏢 Établissements (page avec onglets RPA/CHSLD/RI)
      {
        path: "etablissements",
        name: "etablissements",
        component: EtablissementPage,
      },

      // 🏠 RPA (page dédiée - garde-la si tu l'utilises encore)
      {
        path: "rpa",
        name: "rpa",
        component: RPAPage,
      },

      // 🏥 CHSLD (nouvelle page dédiée)
      {
        path: "chsld",
        name: "chsld",
        component: () => import("../views/etablissements/CHSLDPage.vue"),
      },
      // 🏘️ pharmacies
      {
        path: "pharmacies",
        name: "pharmacies",
        component: PharmaciesPage,
      },

      {
        path: "rapports",
        name: "rapports",
        component: ServicePage,
      },

         {
        path: "admin",
        name: "admin",
        component: RevenusPage,
      },
      {
        path: "appels",
        name: "appels",
        component: () => import("../views/AppelPage.vue"),
      },

      // ⚙️ Paramètres
      {
        path: "settings",
        name: "settings",
        component: SettingPage,
      },
    ],
  },

  // Redirection pour les routes inconnues
  { path: "/:pathMatch(.*)*", redirect: "/" },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});


router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.name !== 'login' && !authStore.user) {
    next({ name: 'login' })  // ← pas connecté = retour login
  } else {
    next()  // ← connecté = laisse passer
  }
})
export default router;
