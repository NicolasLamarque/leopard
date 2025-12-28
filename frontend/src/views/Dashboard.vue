<template>
  <div class="dashboard dark:bg-gray-900 min-h-screen transition-colors">
    <AppHeader 
      :user="user" 
      :active-view="route.name"
      @navigate="handleNavigation"
    />
    <main class="content p-6">
      <router-view /> 
    </main>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'  // ✅ AJOUTE CET IMPORT
import { useRoute, useRouter } from 'vue-router'
import AppHeader from '../components/Header.vue'
import { useDarkMode } from '@/composables/useDarkMode'
import { GetSettings } from '../../wailsjs/go/main/App'

const props = defineProps({
  user: { type: Object, required: true }
})

const route = useRoute()
const router = useRouter()

// ✅ AJOUTE: Récupère setTheme depuis le composable
const { setTheme } = useDarkMode()

const handleNavigation = (viewId) => {
  // IMPORTANT: Ajoute /app/ devant le chemin
  router.push(`/app/${viewId}`)
}

// ✅ Charger le thème au montage du Dashboard
onMounted(async () => {
  try {
    const settings = await GetSettings()
    setTheme(settings.theme || 'light')
    console.log('🎨 Thème chargé:', settings.theme)
  } catch (err) {
    console.error('❌ Erreur chargement thème:', err)
  }
})
</script>

<style scoped>
.dashboard {
  min-height: 100vh;
}

.content {
  max-width: 1400px;
  margin: 0 auto;
}
</style>