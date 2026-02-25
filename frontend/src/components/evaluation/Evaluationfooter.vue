<template>
  <div v-if="brouillons.length > 0" class="shrink-0 border-t dark:border-gray-800 bg-gradient-to-r from-orange-50/80 to-amber-50/80 dark:from-orange-900/10 dark:to-amber-900/10 p-4">
    <div class="flex items-center gap-2 mb-3">
      <span class="w-2 h-2 rounded-full bg-orange-400 animate-pulse" />
      <h3 class="text-xs font-bold text-orange-700 dark:text-orange-400 uppercase tracking-wider">
        Brouillons en cours — {{ brouillons.length }}
      </h3>
    </div>

    <div class="flex gap-3 overflow-x-auto pb-1">
      <div
        v-for="draft in brouillons"
        :key="draft.id"
        class="flex-shrink-0 w-64 bg-white dark:bg-gray-900 rounded-xl border-2 border-orange-200 dark:border-orange-800/50 p-3.5 hover:border-orange-400 hover:shadow-md transition-all"
      >
        <!-- Type + Date -->
        <div class="flex items-center justify-between mb-2">
          <span :class="['text-[9px] font-bold px-1.5 py-0.5 rounded-full uppercase', typeClass(draft.type_evaluation)]">
            {{ typeLabel(draft.type_evaluation) }}
          </span>
          <span class="text-[10px] text-gray-400">{{ formatDate(draft.created_at) }}</span>
        </div>

        <!-- Nom Léopard -->
        <p v-if="draft.nom_leopard" class="text-[9px] font-mono text-teal-500 mb-1 truncate">
          {{ draft.nom_leopard }}
        </p>

        <!-- Contenu -->
        <p class="text-xs font-semibold text-gray-900 dark:text-white line-clamp-2 mb-1 leading-snug">
          {{ draft.objet_evaluation || draft.motif_reference || 'Brouillon sans titre' }}
        </p>
        <p class="text-[10px] text-gray-400 line-clamp-1">
          {{ draft.contexte_evaluation || '—' }}
        </p>

        <!-- Actions -->
        <div class="flex gap-2 mt-3">
          <button
            @click="$emit('view-draft', draft)"
            class="flex-1 py-1.5 bg-orange-100 hover:bg-orange-200 dark:bg-orange-900/30 dark:hover:bg-orange-900/50 text-orange-700 dark:text-orange-300 rounded-lg text-[11px] font-bold transition-all"
          >
            Reprendre →
          </button>
          <button
            @click="$emit('delete-draft', draft.id)"
            class="p-1.5 bg-red-50 hover:bg-red-100 dark:bg-red-900/20 dark:hover:bg-red-900/30 text-red-500 rounded-lg transition-all"
            title="Supprimer"
          >
            <Trash2 :size="14" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Trash2 } from 'lucide-vue-next'

defineProps({ brouillons: Array })
defineEmits(['view-draft', 'delete-draft'])

const typeLabel = (t) => {
  const m = { tutelle: 'Tutelle', mandat: 'Mandat', conseil_tutelle: 'Conseil', suivi_psychosocial: 'Suivi' }
  return m[t] || 'Éval.'
}
const typeClass = (t) => {
  const m = {
    tutelle:            'bg-blue-100 text-blue-700',
    mandat:             'bg-purple-100 text-purple-700',
    conseil_tutelle:    'bg-amber-100 text-amber-700',
    suivi_psychosocial: 'bg-teal-100 text-teal-700'
  }
  return m[t] || 'bg-gray-100 text-gray-600'
}
const formatDate = (d) => d ? new Date(d).toLocaleDateString('fr-CA', { month: 'short', day: 'numeric' }) : '-'
</script>

<style scoped>
.animate-pulse { animation: pulse 2s cubic-bezier(0.4,0,0.6,1) infinite; }
@keyframes pulse { 0%,100%{opacity:1}50%{opacity:.5} }
</style>