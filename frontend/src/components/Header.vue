<template>
  <header class="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 shadow-sm transition-colors sticky top-0 z-50">
    <div class="max-w-[1400px] mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        
        <!-- Logo et titre -->
        <div class="flex items-center gap-3">
          <LogoLeopard class="w-10 h-10 flex-shrink-0" />
          <div class="hidden sm:block">
            <h1 class="text-xl font-extrabold text-gray-900 dark:text-white tracking-tight">LEOPARD</h1>
            <p class="text-xs text-gray-500 dark:text-gray-400 uppercase tracking-wider">SGBD Psychosocial</p>
          </div>
        </div>

        <!-- Navigation principale (TOUJOURS VISIBLE sur desktop) -->
        <nav class="hidden md:flex items-center gap-1">
          <button
            v-for="item in navItems"
            :key="item.id"
            @click="handleNavigation(item.id)"
            :class="[
              'flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200',
              isActiveRoute(item.id)
                ? 'bg-blue-600 text-white shadow-md'
                : 'text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700',
            ]"
          >
            <component :is="item.icon" :size="18" />
            <span>{{ item.label }}</span>
          </button>
        </nav>

        <!-- Actions rapides (desktop) -->
        <div class="hidden md:flex items-center gap-2">
          
          <!-- Recherche -->
          <button 
            @click="toggleSearch"
            :class="[
              'p-2.5 rounded-lg transition-colors',
              showSearch ? 'bg-blue-50 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400' : 'text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700'
            ]"
            title="Rechercher"
          >
            <Search :size="20" />
          </button>

          <!-- Notifications -->
          <button 
            @click="toggleNotifications"
            class="relative p-2.5 text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-colors"
            title="Notifications"
          >
            <Bell :size="20" />
            <span v-if="notificationCount > 0" 
                  class="absolute top-1.5 right-1.5 w-2 h-2 bg-red-500 rounded-full animate-pulse">
            </span>
          </button>

          <!-- Paramètres -->
          <button
            @click="handleNavigation('settings')"
            :class="[
              'p-2.5 rounded-lg transition-colors',
              isActiveRoute('settings') 
                ? 'bg-blue-50 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400' 
                : 'text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700'
            ]"
            title="Paramètres"
          >
            <Settings :size="20" />
          </button>

          <!-- Séparateur -->
          <div class="w-px h-8 bg-gray-200 dark:bg-gray-700 mx-2"></div>

          <!-- Profil utilisateur -->
          <div class="flex items-center gap-3">
            <div class="text-right hidden lg:block">
              <p class="text-sm font-semibold text-gray-900 dark:text-white">{{ user.fullName }}</p>
              <p class="text-xs text-gray-500 dark:text-gray-400">{{ getRoleLabel(user.role) }}</p>
            </div>
            <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-pink-500 rounded-full flex items-center justify-center text-white font-bold shadow-md">
              {{ getInitials(user.fullName) }}
            </div>
          </div>

          <!-- Déconnexion -->
          <button
            @click="handleLogout" 
            class="flex items-center gap-2 px-3 py-2 text-gray-600 dark:text-gray-400 hover:text-red-600 dark:hover:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
            title="Déconnexion"
          >
            <LogOut :size="18" />
          </button>
        </div>

        <!-- Bouton menu mobile -->
        <button 
          @click="mobileMenuOpen = !mobileMenuOpen" 
          class="md:hidden p-2 text-gray-600 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg"
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
        <div v-if="mobileMenuOpen" class="md:hidden py-4 border-t border-gray-200 dark:border-gray-700">
          
          <!-- Navigation mobile -->
          <nav class="flex flex-col gap-1">
            <button
              v-for="item in navItems"
              :key="item.id"
              @click="handleNavigation(item.id)"
              :class="[
                'flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium transition-all',
                isActiveRoute(item.id)
                  ? 'bg-blue-600 text-white'
                  : 'text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700',
              ]"
            >
              <component :is="item.icon" :size="20" />
              <span>{{ item.label }}</span>
            </button>
          </nav>
          
          <!-- Actions mobile -->
          <div class="mt-4 pt-4 border-t border-gray-200 dark:border-gray-700 flex flex-col gap-2">
            <button
              @click="toggleSearch"
              class="flex items-center gap-3 px-4 py-3 text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg"
            >
              <Search :size="20" />
              <span class="text-sm font-medium">Rechercher</span>
            </button>
            
            <button
              @click="handleLogout"
              class="flex items-center gap-3 px-4 py-3 text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg"
            >
              <LogOut :size="20" />
              <span class="text-sm font-medium">Déconnexion</span>
            </button>
          </div>
        </div>
      </transition>
    </div>

    <!-- Barre de recherche globale -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0 -translate-y-4"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 -translate-y-4"
    >
      <div v-if="showSearch" class="border-t border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-900/50">
        <div class="max-w-[1400px] mx-auto px-4 py-4">
          <div class="relative">
            <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" :size="20" />
            <input
              ref="searchInput"
              v-model="searchQuery"
              type="text"
              placeholder="Rechercher un client, une note, un médecin..."
              class="w-full pl-11 pr-4 py-3 rounded-lg border border-gray-300 dark:border-gray-600
                     bg-white dark:bg-gray-800 text-gray-900 dark:text-white
                     focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              @keydown.esc="showSearch = false"
            />
          </div>
        </div>
      </div>
    </transition>
  </header>
</template>

<script setup>
import { ref, nextTick } from "vue";
import { useRouter, useRoute } from 'vue-router';
import { Home, Users, FileText, Settings, LogOut, Menu, X, Folder, Bell, Search } from "lucide-vue-next";
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

// États
const mobileMenuOpen = ref(false);
const showSearch = ref(false);
const showNotifications = ref(false);
const searchQuery = ref('');
const searchInput = ref(null);
const notificationCount = ref(0);

// Navigation
const navItems = [
  { id: "home", label: "Accueil", icon: Home },
  { id: "clients", label: "Clients", icon: Users },
  { id: "medecins", label: "Médecins", icon: FileText },
  { id: "dossiers", label: "Dossiers", icon: Folder, IsAction: true },
];

const isActiveRoute = (routeId) => {
  if (routeId === 'home') {
    return route.path === '/app' || route.path === '/app/' || route.name === 'home';
  }
  return route.path.includes(`/app/${routeId}`) || route.name === routeId;
};

const handleNavigation = async (viewId) => {
  mobileMenuOpen.value = false;
  
  try {
    if (viewId === 'dossiers') {
      await openClientFolder();
      return;
    }
    
    const targetPath = viewId === 'home' ? '/app' : `/app/${viewId}`;
    
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
        console.log('📂 Dossier ouvert avec succès:', result.path);
      }
      emit('openFolder', 'clients');
    }
  } catch (error) {
    console.error('❌ Erreur ouverture dossier:', error);
  }
};

// Recherche
const toggleSearch = async () => {
  showSearch.value = !showSearch.value;
  if (showSearch.value) {
    await nextTick();
    searchInput.value?.focus();
  }
  emit('search', searchQuery.value);
};

// Notifications
const toggleNotifications = () => {
  showNotifications.value = !showNotifications.value;
  emit('notifications');
};

// Déconnexion
const handleLogout = async () => {
  try {
    const confirmed = confirm("Êtes-vous sûr de vouloir vous déconnecter ?");
    if (!confirmed) return;
    
    emit("logout");
    await router.push('/');
    
    if (window.go?.main?.App?.Logout) {
      await window.go.main.App.Logout();
    }
  } catch (error) {
    console.error("❌ Erreur lors de la déconnexion:", error);
  }
};

// Utilitaires
const getRoleLabel = (role) => {
  const labels = { 
    admin: "Administrateur", 
    psychologue: "Psychologue", 
    travailleur_social: "Intervenant", 
    lecteur: "Consultation" 
  };
  return labels[role] || role;
};

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