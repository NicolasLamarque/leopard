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
import { useRoute, useRouter } from 'vue-router'
import AppHeader from '../components/Header.vue'


const props = defineProps({
  user: { type: Object, required: true }
})

const route = useRoute()
const router = useRouter()

const handleNavigation = (viewId) => {
  // IMPORTANT: Ajoute /app/ devant le chemin
  router.push(`/app/${viewId}`)
}
</script>

<style scoped>
.dashboard {
  min-height: 100vh;
  background: #f5f7fa;
}
.content {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
}
</style>