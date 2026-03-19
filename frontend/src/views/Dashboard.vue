<template>
  <div class="dashboard min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-200">
    <AppHeader 
      :user="authStore.user" 
      @logout="handleLogout"
      @search="handleSearch"
      @notifications="handleNotifications"
      @openFolder="handleFolderOpen"
    />
    
    <main class="main-content transition-all duration-200">
      <div class="content-wrapper">
        <router-view v-slot="{ Component }">
          <transition
            name="fade-slide"
            mode="out-in"
          >
            <component :is="Component" />
          </transition>
        </router-view>
      </div>
    </main>

    <!-- Footer optionnel -->
    <footer class="bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 py-4 mt-auto">
      <div class="max-w-[1400px] mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex flex-col sm:flex-row justify-between items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
          <p>&copy; 2025 Leopard SGBD. Tous droits réservés.</p>
          <p>Version 1.0.0</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import AppHeader from '../components/Header.vue'
import { useDarkMode } from '@/composables/useDarkMode'
import { GetSettings } from '../../wailsjs/go/main/App'
import { useAuthStore } from '@/stores/auth'       // ← AJOUTE

const authStore = useAuthStore()   

const router = useRouter()
const { setTheme } = useDarkMode()

// Charger les paramètres utilisateur au montage
onMounted(async () => {
  try {
    const settings = await GetSettings()
    setTheme(settings.theme || 'light')
    console.log('🎨 Thème chargé:', settings.theme)
  } catch (err) {
    console.error('❌ Erreur chargement paramètres:', err)
  }
})

const handleLogout = async () => {
  try {
    authStore.clearUser()
    sessionStorage.clear()
    localStorage.clear()
    await router.push({ name: 'login' })
    window.location.reload() // ← remet tout à zéro proprement
  } catch (err) {
    console.error('❌ Erreur déconnexion:', err)
  }
}
// Gestion de la recherche
const handleSearch = (query) => {
  console.log('🔍 Recherche:', query)
  // TODO: Implémenter la recherche globale
}

// Gestion des notifications
const handleNotifications = () => {
  console.log('🔔 Notifications ouvertes')
  // TODO: Implémenter le système de notifications
}

// Gestion de l'ouverture de dossier
const handleFolderOpen = (type) => {
  console.log('📂 Dossier ouvert:', type)
}
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  min-height: calc(100vh - 64px - 60px); /* viewport - header - footer */
}

.content-wrapper {
  max-width: 1400px;
  margin: 0 auto;
  padding: 1.5rem;
}

/* Responsive */
@media (max-width: 640px) {
  .content-wrapper {
    padding: 1rem;
  }
}

/* Transitions entre pages */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.2s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateX(-10px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateX(10px);
}
</style>