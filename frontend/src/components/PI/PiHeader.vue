<script setup>
import { 
  ClipboardList, ShieldCheck, FileDown, Plus, X, 
  Edit3, Save, ArrowLeft, MoreVertical 
} from 'lucide-vue-next'

const props = defineProps({
  client: {
    type: Object,
    required: true
  },
  selectedPi: Object,
  currentView: {
    type: String,
    default: 'list'
  },
  isExporting: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['start-new', 'cancel-editing', 'export-pdf', 'close'])

const getStatusBadge = () => {
  if (!props.selectedPi) return null
  
  if (props.selectedPi.verrouille) {
    return {
      label: 'Scellé & Actif',
      icon: ShieldCheck,
      class: 'bg-emerald-500/90 text-white border-emerald-400/30'
    }
  }
  
  const statusMap = {
    'brouillon': {
      label: 'Brouillon',
      class: 'bg-amber-500/90 text-white border-amber-400/30'
    },
    'actif': {
      label: 'Actif',
      class: 'bg-teal-500/90 text-white border-teal-400/30'
    },
    'termine': {
      label: 'Terminé',
      class: 'bg-slate-500/90 text-white border-slate-400/30'
    },
    'archive': {
      label: 'Archivé',
      class: 'bg-slate-400/90 text-white border-slate-300/30'
    }
  }
  
  return statusMap[props.selectedPi.statut] || statusMap.brouillon
}

const statusBadge = computed(() => getStatusBadge())
</script>

<template>
  <div class="px-6 py-4 border-b border-slate-200 dark:border-slate-800 
              bg-gradient-to-r from-rose-500/95 via-pink-500/95 to-purple-500/95 
              backdrop-blur-sm shadow-lg">
    <div class="flex items-center justify-between">
      
      <!-- Left Section: Title & Client Info -->
      <div class="flex items-center gap-4">
        <div class="p-2.5 bg-white/20 rounded-xl text-white backdrop-blur-md shadow-sm">
          <ClipboardList :size="24" :stroke-width="2.5" />
        </div>
        
        <div>
          <h2 class="text-xl font-bold text-white flex items-center gap-2">
            Plan d'Intervention
            <span v-if="currentView === 'create'" class="text-sm font-medium opacity-75">
              / Nouveau
            </span>
            <span v-else-if="currentView === 'edit'" class="text-sm font-medium opacity-75">
              / Édition
            </span>
          </h2>
          <div class="flex items-center gap-3 mt-0.5">
            <p class="text-xs text-white/90 font-medium">
              {{ client?.prenom }} {{ client?.nom }}
            </p>
            <span class="text-white/40">•</span>
            <p class="text-xs text-white/75 font-mono">
              {{ client?.no_dossier_leopard }}
            </p>
          </div>
        </div>
      </div>
      
      <!-- Right Section: Actions & Status -->
      <div class="flex items-center gap-3">
        
        <!-- Status Badge -->
        <div 
          v-if="statusBadge && currentView === 'view'"
          :class="['flex items-center gap-2 px-4 py-2 rounded-xl border shadow-sm', statusBadge.class]"
        >
          <component v-if="statusBadge.icon" :is="statusBadge.icon" :size="16" :stroke-width="2.5" />
          <span class="text-xs font-bold uppercase tracking-wide">{{ statusBadge.label }}</span>
        </div>

        <!-- Action Buttons -->
        <template v-if="currentView === 'list'">
          <!-- Export PDF -->
          <button 
            v-if="selectedPi"
            @click="$emit('export-pdf')"
            :disabled="isExporting"
            class="bg-white/20 hover:bg-white/30 text-white px-4 py-2 rounded-xl 
                   text-sm font-semibold flex items-center gap-2 transition-all 
                   disabled:opacity-50 disabled:cursor-not-allowed backdrop-blur-sm
                   border border-white/20 shadow-sm"
          >
            <FileDown :size="18" :stroke-width="2" />
            <span>{{ isExporting ? 'Export...' : 'PDF' }}</span>
          </button>

          <!-- New PI -->
          <button 
            @click="$emit('start-new')" 
            class="bg-white hover:bg-white/95 text-rose-600 px-5 py-2 rounded-xl 
                   text-sm font-bold flex items-center gap-2 shadow-lg transition-all
                   hover:shadow-xl hover:scale-[1.02]"
          >
            <Plus :size="20" :stroke-width="2.5" />
            <span>Nouveau PI</span>
          </button>
        </template>

        <template v-else>
          <!-- Cancel/Back Button -->
          <button 
            @click="$emit('cancel-editing')" 
            class="bg-white/20 hover:bg-white/30 text-white px-4 py-2 rounded-xl 
                   text-sm font-semibold flex items-center gap-2 transition-all
                   backdrop-blur-sm border border-white/20"
          >
            <ArrowLeft :size="18" :stroke-width="2" />
            <span>Retour</span>
          </button>
        </template>

        <!-- Divider -->
        <div class="h-8 w-px bg-white/20"></div>
        
        <!-- Close Button -->
        <button 
          @click="$emit('close')" 
          class="p-2 hover:bg-white/20 rounded-xl transition-colors backdrop-blur-sm"
          title="Fermer"
        >
          <X :size="22" class="text-white" :stroke-width="2" />
        </button>
      </div>
    </div>
  </div>
</template>