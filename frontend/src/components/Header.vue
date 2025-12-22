<template>
  <header class="bg-white border-b border-gray-200 shadow-sm">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <div class="flex items-center gap-3">
          <LogoLeopard class="w-10 h-10" />
          <div class="hidden sm:block">
            <h1 class="text-xl font-extrabold text-gray-900 tracking-tight">LEOPARD</h1>
            <p class="text-xs text-gray-500 uppercase tracking-wider">SGBD Psychosocial</p>
          </div>
        </div>

        <nav class="hidden md:flex items-center gap-1">
          <button
            v-for="item in navItems"
            :key="item.id"
            @click="handleNavigation(item.id)"
            :class="[
              'flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200',
              isActiveRoute(item.id)
                ? 'bg-blue-50 text-blue-700 shadow-sm'
                : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900',
            ]"
          >
            <component :is="item.icon" :size="18" />
            <span>{{ item.label }}</span>
          </button>
        </nav>

        <div class="hidden md:flex items-center gap-3">
          <button 
            @click="handleSearch"
            class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors"
            title="Rechercher"
          >
            <Search :size="20" />
          </button>
          
          <button 
            @click="handleNotifications"
            class="relative p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors"
            title="Notifications"
          >
            <Bell :size="20" />
            <span v-if="notificationCount > 0" class="absolute top-1 right-1 w-2 h-2 bg-red-500 rounded-full"></span>
          </button>

          <button
            @click="handleNavigation('settings')"
            :class="[
              'p-2 rounded-lg transition-colors',
              isActiveRoute('settings') 
                ? 'text-blue-700 bg-blue-50' 
                : 'text-gray-500 hover:text-gray-700 hover:bg-gray-100'
            ]"
            title="Paramètres"
          >
            <Settings :size="20" />
          </button>

          <div class="w-px h-8 bg-gray-200"></div>

          <div class="flex items-center gap-3">
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ user.fullName }}</p>
              <p class="text-xs text-gray-500">{{ getRoleLabel(user.role) }}</p>
            </div>
            <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-pink-500 rounded-full flex items-center justify-center text-white font-bold">
              {{ getInitials(user.fullName) }}
            </div>
          </div>

          <button
            @click="handleLogout" 
            class="flex items-center gap-2 px-3 py-2 text-gray-600 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
            title="Déconnexion"
          >
            <LogOut :size="18" />
            <span class="text-sm font-medium">Quitter</span>
          </button>
        </div>

        <button 
          @click="mobileMenuOpen = !mobileMenuOpen" 
          class="md:hidden p-2 text-gray-600 hover:bg-gray-100 rounded-lg"
        >
          <X v-if="mobileMenuOpen" :size="24" />
          <Menu v-else :size="24" />
        </button>
      </div>

      <!-- Menu Mobile -->
      <transition
        enter-active-class="transition ease-out duration-200"
        enter-from-class="opacity-0 -translate-y-2"
        enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition ease-in duration-150"
        leave-from-class="opacity-100 translate-y-0"
        leave-to-class="opacity-0 -translate-y-2"
      >
        <div v-if="mobileMenuOpen" class="md:hidden py-4 border-t border-gray-200">
          <nav class="flex flex-col gap-2">
            <button
              v-for="item in navItems"
              :key="item.id"
              @click="handleNavigation(item.id)"
              :class="[
                'flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-all',
                isActiveRoute(item.id)
                  ? 'bg-blue-50 text-blue-700'
                  : 'text-gray-600 hover:bg-gray-50',
              ]"
            >
              <component :is="item.icon" :size="20" />
              <span>{{ item.label }}</span>
            </button>
          </nav>
          
          <div class="mt-4 pt-4 border-t border-gray-200">
            <button
              @click="handleLogout"
              class="flex items-center gap-3 px-4 py-3 w-full text-left text-gray-600 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
            >
              <LogOut :size="20" />
              <span class="text-sm font-medium">Déconnexion</span>
            </button>
          </div>
        </div>
      </transition>
    </div>
  </header>
</template>

<script setup>
import { ref } from "vue";
import { useRouter, useRoute } from 'vue-router';
import { Users, FileText, Settings, LogOut, Menu, X, Folder, Bell, Search } from "lucide-vue-next";
import LogoLeopard from "./LeopardLogo.vue";

const router = useRouter();
const route = useRoute();

const props = defineProps({
  user: { 
    type: Object, 
    required: true, 
    default: () => ({ fullName: "Utilisateur", role: "lecteur" }) 
  }
});

const emit = defineEmits(["logout", "search", "notifications", "openFolder"]);

const mobileMenuOpen = ref(false);
const notificationCount = ref(0);

const navItems = [
  { id: "clients", label: "Clients", icon: Users },
  { id: "dossiers", label: "Dossiers", icon: Folder,IsAction: true },
  { id: "notes", label: "Notes", icon: FileText },
];

/**
 * Vérifie si une route est active
 */
const isActiveRoute = (routeId) => {
  return route.path.includes(`/app/${routeId}`) || route.name === routeId;
};

const handleNavigation = async (viewId) => {
  mobileMenuOpen.value = false;
  
  try {
    // 1. Détection de l'action spéciale
    if (viewId === 'dossiers') {
      await openClientFolder();
      return; // <--- TRÈS IMPORTANT : On s'arrête ici ! 
              // L'explorateur s'ouvre, mais on ne change pas de page dans l'app.
    }
    
    // 2. Navigation normale (uniquement si ce n'est pas 'dossiers')
    const targetPath = `/app/${viewId}`;
    
    if (route.path !== targetPath) {
      await router.push(targetPath);
    }
  } catch (error) {
    console.error(`❌ Erreur de navigation vers ${viewId}:`, error);
  }
};

const openClientFolder = async () => {
  try {
    if (window.go?.main?.App?.OpenMainClientsFolder) {
      const result = await window.go.main.App.OpenMainClientsFolder();
      
      if (result.success) {
        console.log('📁 Dossier ouvert avec succès:', result.path);
      } else {
        console.error('❌ Erreur:', result.error);
      }
      
      emit('openFolder', 'clients');
    }
  } catch (error) {
    console.error('❌ Erreur ouverture dossier:', error);
  }
};

/**
 * Gestion de la recherche
 */
const handleSearch = () => {
  emit('search');
  console.log('🔍 Recherche ouverte');
};

/**
 * Gestion des notifications
 */
const handleNotifications = () => {
  emit('notifications');
  console.log('🔔 Notifications ouvertes');
};

/**
 * Gestion de la déconnexion avec confirmation
 */
const handleLogout = async () => {
  try {
    console.log("🚪 Déconnexion en cours...");
    
    // Confirmation optionnelle
    const confirmed = confirm("Êtes-vous sûr de vouloir vous déconnecter ?");
    if (!confirmed) return;
    
    // Émettre l'événement logout vers le parent
    emit("logout");
    
    // Navigation vers la page de connexion
    await router.push('/');
    
    // Optionnel : Appeler la fonction Logout du backend
    if (window.go && window.go.main && window.go.main.App) {
      await window.go.main.App.Logout();
    }
  } catch (error) {
    console.error("❌ Erreur lors de la déconnexion:", error);
  }
};

/**
 * Obtenir le label du rôle
 */
const getRoleLabel = (role) => {
  const labels = { 
    admin: "Administrateur", 
    psychologue: "Psychologue", 
    travailleur_social: "Intervenant", 
    lecteur: "Consultation" 
  };
  return labels[role] || role;
};

/**
 * Obtenir les initiales du nom
 */
const getInitials = (name) => {
  if (!name) return "U";
  return name
    .split(" ")
    .map((n) => n[0])
    .join("")
    .toUpperCase()
    .slice(0, 2);
};
</script>

<style scoped>
/* Animations pour les transitions */
</style>