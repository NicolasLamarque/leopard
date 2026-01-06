<!-- EtablissementPage.vue  -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <!-- Header avec navigation par onglets -->
    <div class="bg-white dark:bg-gray-800 border-b dark:border-gray-700 sticky top-0 z-20">
      <div class="max-w-7xl mx-auto px-6">
        <!-- Titre principal -->
        <div class="py-4 border-b border-gray-200 dark:border-gray-700">
          <div class="flex items-center gap-3">
            <Building :size="32" class="text-blue-600 dark:text-blue-400" />
            <div>
              <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
                Gestion des Établissements
              </h1>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                Résidences, CHSLD et Résidences Intermédiaires
              </p>
            </div>
          </div>
        </div>

        <!-- Onglets de navigation -->
        <div class="flex gap-2 pt-2">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'flex items-center gap-2 px-6 py-3 font-medium transition-all',
              activeTab === tab.id
                ? 'text-blue-600 dark:text-blue-400 border-b-2 border-blue-600 dark:border-blue-400'
                : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white'
            ]"
          >
            <component :is="tab.icon" :size="20" />
            <span>{{ tab.label }}</span>
            <span
              v-if="tab.count !== undefined"
              :class="[
                'px-2 py-0.5 text-xs font-semibold rounded-full',
                activeTab === tab.id
                  ? 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400'
                  : 'bg-gray-100 text-gray-600 dark:bg-gray-700 dark:text-gray-400'
              ]"
            >
              {{ tab.count }}
            </span>
          </button>
        </div>
      </div>
    </div>

    <!-- Contenu des onglets -->
    <div class="max-w-7xl mx-auto">
      <!-- Onglet RPA -->
      <div v-show="activeTab === 'rpa'" class="animate-fade-in">
        <RPAPageContent />
      </div>

      <!-- Onglet CHSLD -->
      <div v-show="activeTab === 'chsld'" class="animate-fade-in">
        <CHSLDPageContent />
      </div>

      <!-- Onglet RI -->
      <div v-show="activeTab === 'ri'" class="animate-fade-in">
        <RIPageContent />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Building, Home, Hospital, Building2 } from 'lucide-vue-next'
import RPAPageContent from '../../components/RPA/RPAPageContent.vue'
import CHSLDPageContent from '../../components/CHSLD/CHSLDPageContent.vue'
import RIPageContent from '../../components/RPA/RPAPageContent.vue'

const activeTab = ref('rpa')

// Configuration des onglets
const tabs = computed(() => [
  {
    id: 'rpa',
    label: 'RPA',
    icon: Home,
    count: undefined // Sera chargé dynamiquement
  },
  {
    id: 'chsld',
    label: 'CHSLD',
    icon: Hospital,
    count: undefined
  },
  {
    id: 'ri',
    label: 'Résidences Intermédiaires',
    icon: Building2,
    count: undefined
  }
])
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.3s ease-in;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>