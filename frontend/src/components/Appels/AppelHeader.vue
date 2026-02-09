<!-- frontend/src/components/appels/AppelsHeader.vue -->

<template>
  <div class="px-6 py-4 border-b border-stone-200 dark:border-stone-800 flex justify-between items-center bg-gradient-to-r from-stone-100 via-pink-50 to-stone-100 dark:from-stone-900 dark:via-stone-800 dark:to-stone-900">
    <div class="flex items-center gap-3">
      <div class="p-2.5 bg-white/60 dark:bg-stone-800/60 rounded-xl text-stone-700 dark:text-stone-300 backdrop-blur-sm shadow-sm">
        <Phone :size="22" />
      </div>
      <div>
        <h2 class="text-xl font-semibold text-stone-800 dark:text-stone-200">Gestion des Appels</h2>
        <p class="text-xs text-stone-500 dark:text-stone-400 font-light">
          Suivi des demandes et consultations
        </p>
      </div>
    </div>
    
    <div class="flex items-center gap-4">
      <!-- Stats rapides -->
      <div class="flex items-center gap-4 px-4 py-2 bg-white/50 dark:bg-stone-800/50 rounded-xl backdrop-blur-sm border border-stone-200 dark:border-stone-700">
        <div class="text-center">
          <div class="text-lg font-semibold text-stone-700 dark:text-stone-300">{{ stats.nouveaux }}</div>
          <div class="text-[9px] text-stone-500 dark:text-stone-400 uppercase tracking-wide">Nouveaux</div>
        </div>
        <div class="w-px h-8 bg-stone-300 dark:bg-stone-700"></div>
        <div class="text-center">
          <div class="text-lg font-semibold text-pink-600 dark:text-pink-400">{{ stats.urgents }}</div>
          <div class="text-[9px] text-stone-500 dark:text-stone-400 uppercase tracking-wide">Urgents</div>
        </div>
        <div class="w-px h-8 bg-stone-300 dark:bg-stone-700"></div>
        <div class="text-center">
          <div class="text-lg font-semibold text-stone-600 dark:text-stone-400">{{ stats.total }}</div>
          <div class="text-[9px] text-stone-500 dark:text-stone-400 uppercase tracking-wide">Total</div>
        </div>
      </div>

      <!-- Bouton Nouvel Appel -->
      <button 
        v-if="!isCreating" 
        @click="$emit('start-new')" 
        class="bg-gradient-to-r from-pink-500/85 to-rose-500/45 hover:from-pink-600/30 hover:to-rose-600/55 text-white px-5 py-2.5 rounded-xl text-sm font-medium flex items-center gap-2 shadow-md shadow-pink-500/20 transition-all duration-200 hover:shadow-lg hover:shadow-pink-500/30"
      >
        <Plus :size="18" /> Nouvel Appel
      </button>
      
      <!-- Bouton Retour à la liste -->
      <button 
        v-else
        @click="$emit('cancel-new')" 
        class="bg-white dark:bg-stone-800 text-stone-700 dark:text-stone-300 px-5 py-2.5 rounded-xl text-sm font-medium flex items-center gap-2 shadow-md border border-stone-200 dark:border-stone-700 transition-all duration-200 hover:bg-stone-50 dark:hover:bg-stone-700"
      >
        <ArrowLeft :size="18" /> Retour
      </button>
    </div>
  </div>
</template>

<script setup>
import { Phone, Plus, ArrowLeft } from 'lucide-vue-next'

defineProps({
  stats: {
    type: Object,
    default: () => ({ nouveaux: 0, urgents: 0, total: 0 })
  },
  isCreating: {
    type: Boolean,
    default: false
  }
})

defineEmits(['start-new', 'cancel-new'])
</script>