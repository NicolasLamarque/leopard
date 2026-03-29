<template>
  <div class="px-6 py-4 border-b border-stone-200 dark:border-stone-800 flex justify-between items-center bg-gradient-to-r from-stone-100 via-pink-50 to-stone-100 dark:from-stone-900 dark:via-stone-800 dark:to-stone-900">
    
    <div class="flex items-center gap-3">
      <div class="p-2.5 bg-white/60 dark:bg-stone-800/60 rounded-xl text-stone-700 dark:text-stone-300 backdrop-blur-sm shadow-sm">
        <Phone :size="22" />
      </div>
      <div>
        <h2 class="text-xl font-semibold text-stone-800 dark:text-stone-200">Gestion des Appels</h2>
        <p class="text-xs text-stone-500 dark:text-stone-400 font-light">Suivi des demandes et consultations</p>
      </div>
    </div>

    <div class="flex items-center gap-4">
      
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

      <div class="flex items-center gap-2">
        <div v-if="!isCreating" class="relative group">
          <button 
            type="button"
            class="bg-white dark:bg-stone-800 text-stone-700 dark:text-stone-300 px-4 py-2.5 rounded-xl text-sm font-medium flex items-center gap-2 shadow-sm border border-stone-200 dark:border-stone-700 transition-all hover:bg-stone-50"
          >
            <FileText :size="18" class="text-pink-500" />
            <span>Rapports</span>
            <ChevronDown :size="14" class="text-stone-400 group-hover:rotate-180 transition-transform" />
          </button>
          
          <div class="absolute right-0 pt-2 w-52 bg-white dark:bg-stone-800 border border-stone-200 dark:border-stone-700 rounded-xl shadow-xl hidden group-hover:block z-[100] overflow-hidden animate-in fade-in slide-in-from-top-1">
            <div class="p-1.5 flex flex-col">
              <button @click.stop="$emit('generate-report', 'weekly')" class="flex items-center gap-2 w-full text-left px-3 py-2 text-xs font-medium text-stone-700 dark:text-stone-300 hover:bg-pink-50 dark:hover:bg-pink-900/20 hover:text-pink-600 rounded-lg transition-colors">
                <Calendar :size="14" /> Rapport Hebdomadaire
              </button>
              <button @click.stop="$emit('generate-report', 'monthly')" class="flex items-center gap-2 w-full text-left px-3 py-2 text-xs font-medium text-stone-700 dark:text-stone-300 hover:bg-pink-50 dark:hover:bg-pink-900/20 hover:text-pink-600 rounded-lg transition-colors">
                <BarChart3 :size="14" /> Rapport Mensuel
              </button>
            </div>
          </div>
        </div>

        <button
          v-if="!isCreating"
          @click="$emit('start-new')"
          class="bg-gradient-to-r from-pink-500 to-rose-500 hover:from-pink-600 hover:to-rose-600 text-white px-5 py-2.5 rounded-xl text-sm font-medium flex items-center gap-2 shadow-md shadow-pink-500/20 transition-all"
        >
          <Plus :size="18" /> Nouvel Appel
        </button>

        <button
          v-else
          @click="$emit('cancel-new')"
          class="bg-white dark:bg-stone-800 text-stone-700 dark:text-stone-300 px-5 py-2.5 rounded-xl text-sm font-medium flex items-center gap-2 shadow-md border border-stone-200 dark:border-stone-700 hover:bg-stone-50 transition-all"
        >
          <ArrowLeft :size="18" /> Retour
        </button>
      </div>

    </div>
  </div>
</template>

<script setup>
import { 
  Phone, 
  Plus, 
  ArrowLeft, 
  FileText, 
  ChevronDown, 
  Calendar, 
  BarChart3 
} from "lucide-vue-next";

defineProps({
  stats: {
    type: Object,
    default: () => ({ nouveaux: 0, urgents: 0, total: 0 }),
  },
  isCreating: {
    type: Boolean,
    default: false,
  },
});

defineEmits(["start-new", "cancel-new", "generate-report"]);
</script>