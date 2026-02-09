import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Login from "../views/Login.vue";
import Dashboard from "../views/Dashboard.vue";
import HomePage from "../views/homePage.vue";
import ClientPage from "../views/ClientPage.vue";
import ClientDetailsPage from "../views/ClientDetailsPage.vue";
import SettingPage from "../views/SettingPage.vue";
import RPAPage from "../views/etablissements/RPAPage.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "login",
    component: Login,
  },
  {
    path: "/app",
    component: Dashboard,
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
        component: () => import("../views/MedecinPage.vue"),
      },

      // 📋 Autres sections (lazy loading)
      {
        path: "notaires",
        name: "notaires",
        component: () => import("../views/NotairePage.vue"),
      },
      {
        path: "intervenants",
        name: "intervenants",
        component: () => import("../views/IntervenantPage.vue"),
      },

      // 🏢 Établissements (page avec onglets RPA/CHSLD/RI)
      {
        path: "etablissements",
        name: "etablissements",
        component: () =>
          import("../views/etablissements/EtablissementPage.vue"),
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

      {
        path: "rapports",
        name: "rapports",
        component: () => import("../views/RapportPage.vue"),
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

export default router;
