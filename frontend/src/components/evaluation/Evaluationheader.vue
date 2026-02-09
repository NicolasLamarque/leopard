<template>
  <div class="px-6 py-4 border-b dark:border-gray-800 flex justify-between items-center bg-gradient-to-r from-slate-700 to-slate-600">
    <div class="flex items-center gap-3">
      <div class="p-2 bg-white/20 rounded-lg text-white backdrop-blur-sm">
        <Shield :size="24" />
      </div>
      <div>
        <h2 class="text-xl font-bold text-white">Évaluation du Fonctionnement Social</h2>
        <p class="text-xs text-slate-100 uppercase font-medium tracking-wider">
          {{ client?.prenom }} {{ client?.nom }} - {{ client?.no_dossier_leopard }}
        </p>
      </div>
    </div>
    
    <div class="flex items-center gap-4">
      <!-- Badge statut si évaluation verrouillée -->
      <div v-if="selectedEvaluation?.verrouille" class="flex items-center gap-2 bg-emerald-500/20 text-emerald-300 px-4 py-2 rounded-lg border border-emerald-500/30">
        <ShieldCheck :size="16" />
        <span class="text-xs font-bold">Verrouillée</span>
      </div>

      <!-- Bouton Export PDF -->
      <button 
        v-if="!isCreating && selectedEvaluation"
        @click="$emit('export-pdf')"
        :disabled="isExporting"
        class="bg-teal-600 hover:bg-teal-700 text-white px-5 py-2 rounded-xl text-sm font-bold flex items-center gap-2 transition-all disabled:opacity-50"
      >
        <FileDown :size="18" />
        <span v-if="!isExporting">Export PDF</span>
        <span v-else">Génération...</span>
      </button>

      <!-- Bouton Nouvelle Évaluation -->
      <button 
        v-if="!isCreating" 
        @click="$emit('start-new')" 
        class="bg-white hover:bg-white/90 text-slate-700 px-5 py-2 rounded-xl text-sm font-bold flex items-center gap-2 shadow-lg transition-all"
      >
        <Plus :size="18" /> Nouvelle Évaluation
      </button>
      
      <!-- Bouton Voir historique (si en mode création) -->
      <button 
        v-else 
        @click="$emit('cancel-creation')" 
        class="text-white hover:text-slate-100 text-sm font-bold"
      >
        Voir historique
      </button>

      <div class="h-6 w-px bg-white/20 mx-2"></div>
      
      <!-- Bouton Fermer -->
      <button @click="$emit('close')" class="p-2 hover:bg-white/20 rounded-full transition-colors">
        <X :size="22" class="text-white" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { Shield, ShieldCheck, FileDown, Plus, X } from 'lucide-vue-next'

defineProps({
  client: Object,
  selectedEvaluation: Object,
  isCreating: Boolean,
  isExporting: Boolean
})

defineEmits(['start-new', 'cancel-creation', 'export-pdf', 'close'])
</script>