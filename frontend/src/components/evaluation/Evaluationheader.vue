<template>
  <div class="px-6 py-4 border-b dark:border-gray-800 flex justify-between items-center bg-gradient-to-r from-slate-800 via-slate-700 to-slate-600 shrink-0">
    <div class="flex items-center gap-3">
      <div class="p-2 bg-white/15 rounded-lg text-white backdrop-blur-sm">
        <Shield :size="22" />
      </div>
      <div>
        <div class="flex items-center gap-2">
          <h2 class="text-lg font-bold text-white leading-tight">Évaluation du Fonctionnement Social</h2>
          <!-- Badge type -->
          <span v-if="isCreating && typeEvaluation" :class="typeBadgeClass" class="text-[10px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wider whitespace-nowrap">
            {{ typeBadgeLabel }}
          </span>
        </div>
        <p class="text-xs text-slate-300 font-medium tracking-wide">
          {{ client?.prenom }} {{ client?.nom }}
          <span class="text-teal-400 font-mono ml-1">{{ client?.no_dossier_leopard }}</span>
        </p>
      </div>
    </div>

    <div class="flex items-center gap-2">

      <!-- Badge verrouillé -->
      <div v-if="selectedEvaluation?.verrouille && !isCreating"
           class="flex items-center gap-1.5 bg-emerald-500/20 text-emerald-300 px-3 py-1.5 rounded-lg border border-emerald-500/30">
        <ShieldCheck :size="14" />
        <span class="text-xs font-bold">Scellée</span>
      </div>

      <!-- Nom Léopard si sélectionnée -->
      <div v-if="selectedEvaluation?.nom_leopard && !isCreating"
           class="hidden lg:flex items-center gap-1.5 bg-slate-900/60 text-teal-400 px-3 py-1.5 rounded-lg border border-slate-600 font-mono">
        <Tag :size="12" />
        <span class="text-xs">{{ selectedEvaluation.nom_leopard }}</span>
      </div>

      <!-- Export PDF -->
      <button
        v-if="!isCreating && selectedEvaluation"
        @click="$emit('export-pdf')"
        :disabled="isExporting"
        class="flex items-center gap-2 bg-teal-600 hover:bg-teal-500 text-white px-4 py-2 rounded-xl text-sm font-bold transition-all disabled:opacity-50 shadow-lg"
      >
        <Loader2 v-if="isExporting" :size="16" class="animate-spin" />
        <FileDown v-else :size="16" />
        <span class="hidden sm:inline">{{ isExporting ? 'Génération...' : 'Export PDF' }}</span>
      </button>

      <!-- Nouvelle évaluation -->
      <button
        v-if="!isCreating"
        @click="$emit('start-new')"
        class="flex items-center gap-2 bg-white hover:bg-white/90 text-slate-800 px-4 py-2 rounded-xl text-sm font-bold shadow-lg transition-all"
      >
        <Plus :size="16" />
        <span>Nouvelle évaluation</span>
      </button>

      <!-- Voir historique (si création) -->
      <button
        v-else
        @click="$emit('cancel-creation')"
        class="flex items-center gap-2 text-white/70 hover:text-white text-sm font-medium transition-colors border border-white/20 hover:border-white/40 px-3 py-2 rounded-lg"
      >
        <ChevronLeft :size="16" />
        Historique
      </button>

      <div class="w-px h-6 bg-white/20 mx-1" />

      <button @click="$emit('close')" class="p-2 hover:bg-white/10 rounded-full transition-colors">
        <X :size="20" class="text-white/80 hover:text-white" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Shield, ShieldCheck, FileDown, Plus, X, Loader2, ChevronLeft, Tag } from 'lucide-vue-next'

const props = defineProps({
  client: Object,
  selectedEvaluation: Object,
  isCreating: Boolean,
  isExporting: Boolean,
  typeEvaluation: String
})

defineEmits(['start-new', 'cancel-creation', 'export-pdf', 'close'])

const typeMeta = {
  tutelle:           { label: 'Tutelle',           class: 'bg-blue-500/30 text-blue-200 border border-blue-400/30' },
  mandat:            { label: 'Mandat',             class: 'bg-purple-500/30 text-purple-200 border border-purple-400/30' },
  conseil_tutelle:   { label: 'Conseil tutelle',   class: 'bg-amber-500/30 text-amber-200 border border-amber-400/30' },
  suivi_psychosocial:{ label: 'Suivi psychosocial',class: 'bg-teal-500/30 text-teal-200 border border-teal-400/30' }
}

const typeBadgeLabel = computed(() => typeMeta[props.typeEvaluation]?.label || '')
const typeBadgeClass = computed(() => typeMeta[props.typeEvaluation]?.class || '')
</script>

<style scoped>
.animate-spin { animation: spin 1s linear infinite; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
</style>