<script setup>
import { computed } from 'vue'
// On importe les vrais objets de composants depuis ta librairie installée
import { 
  Users, 
  Stethoscope, 
  Briefcase, 
  Hospital, 
  Building, 
  BarChart, 
  UserPlus, 
  FileText, 
  Search,
  ChevronRight 
} from 'lucide-vue-next'

const props = defineProps({
  title: { type: String, required: true },
  description: { type: String, required: true },
  icon: { type: String, required: true },
  color: { type: String, default: 'blue' },
  count: { type: Number }
})

defineEmits(['click'])

/**
 * MAPPING DES ICÔNES
 * On lie la string reçue en prop à l'objet composant Lucide.
 */
const iconComponent = computed(() => {
  const icons = {
    'users': Users,
    'stethoscope': Stethoscope,
    'briefcase': Briefcase,
    'hospital': Hospital,
    'building': Building,
    'chart': BarChart,
    'user-plus': UserPlus,
    'file-text': FileText,
    'search': Search
  }
  // Retourne Users par défaut si la clé n'existe pas pour éviter un crash
  return icons[props.icon] || Users
})

/**
 * CONFIGURATION DES STYLES
 * Au lieu de 5 computed séparés, on centralise pour garantir 
 * la cohérence visuelle entre le badge, l'icône et le hover.
 */
const colorStyles = computed(() => {
  const configs = {
    blue: {
      gradient: 'from-blue-500 to-blue-600',
      bg: 'bg-blue-100 dark:bg-blue-900/30',
      text: 'text-blue-600 dark:text-blue-400',
      badge: 'bg-blue-100 text-blue-700 dark:bg-blue-900/50 dark:text-blue-300'
    },
    emerald: {
      gradient: 'from-emerald-500 to-emerald-600',
      bg: 'bg-emerald-100 dark:bg-emerald-900/30',
      text: 'text-emerald-600 dark:text-emerald-400',
      badge: 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/50 dark:text-emerald-300'
    },
    amber: {
      gradient: 'from-amber-500 to-amber-600',
      bg: 'bg-amber-100 dark:bg-amber-900/30',
      text: 'text-amber-600 dark:text-amber-400',
      badge: 'bg-amber-100 text-amber-700 dark:bg-amber-900/50 dark:text-amber-300'
    },
    purple: {
      gradient: 'from-purple-500 to-purple-600',
      bg: 'bg-purple-100 dark:bg-purple-900/30',
      text: 'text-purple-600 dark:text-purple-400',
      badge: 'bg-purple-100 text-purple-700 dark:bg-purple-900/50 dark:text-purple-300'
    },
    indigo: {
      gradient: 'from-indigo-500 to-indigo-600',
      bg: 'bg-indigo-100 dark:bg-indigo-900/30',
      text: 'text-indigo-600 dark:text-indigo-400',
      badge: 'bg-indigo-100 text-indigo-700 dark:bg-indigo-900/50 dark:text-indigo-300'
    },
    rose: {
      gradient: 'from-rose-500 to-rose-600',
      bg: 'bg-rose-100 dark:bg-rose-900/30',
      text: 'text-rose-600 dark:text-rose-400',
      badge: 'bg-rose-100 text-rose-700 dark:bg-rose-900/50 dark:text-rose-300'
    }
  }
  return configs[props.color] || configs.blue
})
</script>

<template>
  <div 
    @click="$emit('click')"
    class="group relative bg-white dark:bg-gray-800 rounded-xl shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer overflow-hidden border border-gray-200 dark:border-gray-700 hover:scale-105"
  >
    <div 
      class="absolute inset-0 opacity-0 group-hover:opacity-10 transition-opacity duration-300 bg-gradient-to-br"
      :class="colorStyles.gradient"
    ></div>
    
    <div class="relative p-6">
      <div class="flex items-start justify-between mb-4">
        <div :class="colorStyles.bg" class="p-3 rounded-lg">
          <component 
            :is="iconComponent" 
            :class="colorStyles.text" 
            class="w-8 h-8" 
          />
        </div>
        
        <span 
          v-if="count !== undefined" 
          :class="colorStyles.badge" 
          class="px-3 py-1 rounded-full text-sm font-semibold"
        >
          {{ count }}
        </span>
      </div>

      <h3 class="text-xl font-bold text-gray-900 dark:text-white mb-2 group-hover:text-opacity-80 transition-colors">
        {{ title }}
      </h3>
      <p class="text-gray-600 dark:text-gray-400 text-sm">
        {{ description }}
      </p>

      <div 
        class="mt-4 flex items-center text-sm font-medium group-hover:translate-x-2 transition-transform duration-300" 
        :class="colorStyles.text"
      >
        <span>Accéder</span>
        <ChevronRight class="w-4 h-4 ml-2" />
      </div>
    </div>
  </div>
</template>