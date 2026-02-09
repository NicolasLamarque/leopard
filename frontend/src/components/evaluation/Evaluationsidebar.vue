<template>
  <div class="w-72 border-r dark:border-gray-800 bg-slate-50 dark:bg-gray-900/20 flex flex-col">
    <div class="p-4 border-b dark:border-gray-800 bg-slate-100/50 dark:bg-gray-900/50">
      <span class="text-[10px] font-bold text-gray-400 uppercase tracking-widest">
        {{ isCreating ? 'Sections du formulaire' : 'Évaluations finalisées' }}
      </span>
    </div>

    <!-- MODE CRÉATION: Navigation sections -->
    <div v-if="isCreating" class="flex-1 overflow-y-auto p-4 space-y-2">
      <button
        v-for="section in sections"
        :key="section.id"
        @click="$emit('section-change', section.id)"
        :class="[
          'w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all text-left',
          activeSection === section.id
            ? 'bg-teal-50 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300 border-2 border-teal-200 dark:border-teal-800 shadow-sm'
            : 'hover:bg-slate-100 dark:hover:bg-gray-800 text-gray-700 dark:text-gray-300'
        ]"
      >
        <component :is="section.icon" :size="18" />
        <span class="flex-1 font-medium text-sm">{{ section.label }}</span>
      </button>
    </div>

    <!-- MODE VISUALISATION: Liste des évaluations finalisées -->
    <div v-else class="flex-1 overflow-y-auto p-4 space-y-3">
      <div 
        v-for="ev in evaluations" 
        :key="ev.id"
        @click="$emit('view-evaluation', ev)"
        :class="[
          'p-3 rounded-xl border dark:border-gray-800 bg-white dark:bg-gray-900 cursor-pointer transition-all group',
          selectedEvaluation?.id === ev.id 
            ? 'ring-2 ring-teal-500 border-teal-200' 
            : 'hover:ring-2 hover:ring-teal-300'
        ]"
      >
        <div class="flex justify-between text-[10px] mb-1">
          <span class="font-bold text-teal-600 uppercase">Évaluation</span>
          <span class="text-gray-400">{{ formatDate(ev.created_at) }}</span>
        </div>
        <p class="text-sm font-bold text-gray-800 dark:text-gray-200 truncate mb-2">
          {{ ev.objet_evaluation || 'Sans titre' }}
        </p>
        <div class="flex items-center justify-between">
          <span class="text-[9px] text-gray-400 italic truncate">
            {{ ev.signature_nom?.split(',')[0] || 'Non signé' }}
          </span>
          <Lock v-if="ev.verrouille" :size="12" class="text-green-500 flex-shrink-0" />
        </div>
      </div>

      <!-- Message si aucune évaluation -->
      <div v-if="evaluations.length === 0" class="text-center py-8 text-gray-400 text-sm">
        <Lock :size="32" class="mx-auto mb-2 opacity-30" />
        <p>Aucune évaluation finalisée</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Lock } from 'lucide-vue-next'

defineProps({
  evaluations: Array,
  selectedEvaluation: Object,
  isCreating: Boolean,
  activeSection: String,
  sections: Array
})

defineEmits(['section-change', 'view-evaluation'])

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA')
}
</script>