<template>
  <div v-if="brouillons.length > 0" class="border-t dark:border-gray-800 bg-gradient-to-r from-orange-50 to-amber-50 dark:from-orange-900/10 dark:to-amber-900/10 p-4">
    <div class="max-w-7xl mx-auto">
      <div class="flex items-center justify-between mb-3">
        <h3 class="text-sm font-bold text-orange-700 dark:text-orange-400 uppercase tracking-wider flex items-center gap-2">
          <MousePointerClick :size="16" />
          Brouillons en cours ({{ brouillons.length }})
        </h3>
      </div>
      
      <div class="flex gap-3 overflow-x-auto pb-2">
        <div 
          v-for="draft in brouillons" 
          :key="draft.id"
          class="flex-shrink-0 w-72 bg-white dark:bg-gray-900 rounded-xl border-2 border-orange-200 dark:border-orange-800 p-4 hover:shadow-lg transition-all group"
        >
          <div class="flex items-start justify-between mb-2">
            <div class="flex-1">
              <div class="text-xs text-orange-600 dark:text-orange-400 font-bold mb-1">
                {{ formatDate(draft.created_at) }}
              </div>
              <p class="text-sm font-semibold text-gray-900 dark:text-white line-clamp-2">
                {{ draft.objet_evaluation || 'Sans titre' }}
              </p>
            </div>
            <div class="flex items-center gap-1 ml-2">
              <span class="w-2 h-2 rounded-full bg-orange-500 animate-pulse"></span>
            </div>
          </div>
          
          <div class="text-xs text-gray-500 dark:text-gray-400 mb-3 line-clamp-2">
            {{ draft.contexte_evaluation || 'Pas de contexte' }}
          </div>

          <div class="flex gap-2">
            <button
              @click="$emit('view-draft', draft)"
              class="flex-1 px-3 py-2 bg-orange-100 hover:bg-orange-200 dark:bg-orange-900/30 dark:hover:bg-orange-900/50 text-orange-700 dark:text-orange-300 rounded-lg text-xs font-bold transition-all"
            >
              Reprendre
            </button>
            <button
              @click="$emit('delete-draft', draft.id)"
              class="px-3 py-2 bg-red-100 hover:bg-red-200 dark:bg-red-900/30 dark:hover:bg-red-900/50 text-red-700 dark:text-red-300 rounded-lg transition-all"
            >
              <Trash2 :size="16" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { MousePointerClick, Trash2 } from 'lucide-vue-next'

defineProps({
  brouillons: Array
})

defineEmits(['view-draft', 'delete-draft'])

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA')
}
</script>