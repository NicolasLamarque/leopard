<!-- frontend/src/components/appels/AppelsSidebar.vue -->

<template>
  <div class="w-80 border-r border-stone-200 dark:border-stone-800 bg-stone-50/50 dark:bg-stone-900/30 flex flex-col backdrop-blur-sm">
    
    <!-- Header du sidebar -->
    <div class="p-4 border-b border-stone-200 dark:border-stone-800 bg-white/40 dark:bg-stone-900/40">
      <div class="flex items-center justify-between mb-3">
        <span class="text-[10px] font-semibold text-stone-500 dark:text-stone-400 uppercase tracking-widest">
          Liste des appels
        </span>
        <span class="text-xs text-stone-600 dark:text-stone-400 font-medium">
          {{ filteredAppels.length }}
        </span>
      </div>

      <!-- Barre de recherche -->
      <div class="relative">
        <Search :size="14" class="absolute left-3 top-1/2 -translate-y-1/2 text-stone-400" />
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Rechercher..."
          class="w-full pl-9 pr-3 py-2.5 bg-white dark:bg-stone-900 border border-stone-300 dark:border-stone-700 rounded-xl text-sm outline-none focus:ring-2 focus:ring-pink-400 dark:focus:ring-pink-500 focus:border-transparent transition-all placeholder:text-stone-400"
        />
      </div>
    </div>

    <!-- Filtres -->
    <div class="p-3 border-b border-stone-200 dark:border-stone-800 bg-white/30 dark:bg-stone-900/30">
      <div class="flex flex-wrap gap-2">
        <button
          @click="filterStatut = null"
          :class="[
            'px-3 py-1.5 rounded-lg text-xs font-medium transition-all',
            !filterStatut 
              ? 'bg-pink-100 dark:bg-pink-900/40 text-pink-700 dark:text-pink-300 border border-pink-300 dark:border-pink-700/50'
              : 'bg-stone-100 dark:bg-stone-800 text-stone-600 dark:text-stone-400 hover:bg-stone-200 dark:hover:bg-stone-700'
          ]"
        >
          Tous
        </button>
        <button
          v-for="statut in statuts"
          :key="statut.value"
          @click="filterStatut = statut.value"
          :class="[
            'px-3 py-1.5 rounded-lg text-xs font-medium transition-all',
            filterStatut === statut.value
              ? 'bg-pink-100 dark:bg-pink-900/40 text-pink-700 dark:text-pink-300 border border-pink-300 dark:border-pink-700/50'
              : 'bg-stone-100 dark:bg-stone-800 text-stone-600 dark:text-stone-400 hover:bg-stone-200 dark:hover:bg-stone-700'
          ]"
        >
          {{ statut.label }}
        </button>
      </div>
    </div>

    <!-- Liste des appels -->
    <div class="flex-1 overflow-y-auto p-3 space-y-2">
      <div
        v-for="appel in filteredAppels"
        :key="appel.id"
        @click="$emit('select-appel', appel.id)"
        :class="[
          'p-4 rounded-xl cursor-pointer transition-all duration-200 border',
          selectedAppelId === appel.id
            ? 'bg-pink-50 dark:bg-pink-900/20 border-pink-300 dark:border-pink-700/50 shadow-sm'
            : 'bg-white dark:bg-stone-800/50 border-stone-200 dark:border-stone-700/50 hover:bg-stone-50 dark:hover:bg-stone-800/70 hover:shadow-sm'
        ]"
      >
        <!-- Date et heure -->
        <div class="flex items-center justify-between mb-2">
          <div class="flex items-center gap-2 text-xs text-stone-500 dark:text-stone-400">
            <Calendar :size="12" />
            <span>{{ formatDate(appel.date_appel) }}</span>
            <span v-if="appel.heure_appel">• {{ appel.heure_appel }}</span>
          </div>
          <span 
            :class="[
              'px-2 py-0.5 rounded-full text-[10px] font-medium',
              getPrioriteClass(appel.priorite)
            ]"
          >
            {{ appel.priorite }}
          </span>
        </div>

        <!-- Nom -->
        <div class="font-medium text-stone-800 dark:text-stone-200 mb-1">
          {{ getDisplayName(appel) }}
        </div>

        <!-- Type demande -->
        <div class="text-xs text-stone-600 dark:text-stone-400 mb-2">
          {{ getTypeDemande(appel.type_demande) }}
        </div>

        <!-- Statut -->
        <div class="flex items-center justify-between">
          <span 
            :class="[
              'px-2 py-1 rounded-lg text-[10px] font-medium',
              getStatutClass(appel.statut)
            ]"
          >
            {{ getStatutLabel(appel.statut) }}
          </span>
          <Phone :size="14" class="text-stone-400" />
        </div>
      </div>

      <!-- Vide -->
      <div v-if="filteredAppels.length === 0" class="text-center py-12 text-stone-400 dark:text-stone-500 text-sm">
        Aucun appel
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Search, Calendar, Phone } from 'lucide-vue-next'

const props = defineProps({
  appels: {
    type: Array,
    default: () => []
  },
  selectedAppelId: {
    type: Number,
    default: null
  }
})

defineEmits(['select-appel'])

const searchQuery = ref('')
const filterStatut = ref(null)

const statuts = [
  { value: 'nouveau', label: 'Nouveau' },
  { value: 'en_attente', label: 'En attente' },
  { value: 'rdv_planifie', label: 'RDV' },
  { value: 'complete', label: 'Complété' }
]

const filteredAppels = computed(() => {
  let result = props.appels

  // Filtre par statut
  if (filterStatut.value) {
    result = result.filter(a => a.statut === filterStatut.value)
  }

  // Filtre par recherche
  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    result = result.filter(a => 
      a.appelant_nom?.toLowerCase().includes(search) ||
      a.appelant_prenom?.toLowerCase().includes(search) ||
      a.prospect_nom?.toLowerCase().includes(search) ||
      a.prospect_prenom?.toLowerCase().includes(search) ||
      a.appelant_telephone?.includes(search)
    )
  }

  return result
})

const getDisplayName = (appel) => {
  if (appel.prospect_nom || appel.prospect_prenom) {
    return `${appel.prospect_prenom || ''} ${appel.prospect_nom || ''}`.trim()
  }
  return `${appel.appelant_prenom || ''} ${appel.appelant_nom || ''}`.trim() || 'Sans nom'
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fr-CA', { day: 'numeric', month: 'short' })
}

const getTypeDemande = (type) => {
  const types = {
    'evaluation': 'Évaluation',
    'mandat_protection': 'Mandat protection',
    'consultation': 'Consultation',
    'information': 'Information',
    'autre': 'Autre'
  }
  return types[type] || type
}

const getPrioriteClass = (priorite) => {
  const classes = {
    'urgente': 'bg-red-100 text-red-700 dark:bg-red-900/40 dark:text-red-300',
    'haute': 'bg-orange-100 text-orange-700 dark:bg-orange-900/40 dark:text-orange-300',
    'normale': 'bg-stone-100 text-stone-600 dark:bg-stone-800 dark:text-stone-400',
    'basse': 'bg-stone-100 text-stone-500 dark:bg-stone-800 dark:text-stone-500'
  }
  return classes[priorite] || classes.normale
}

const getStatutClass = (statut) => {
  const classes = {
    'nouveau': 'bg-sky-100 text-sky-700 dark:bg-sky-900/40 dark:text-sky-300',
    'en_attente': 'bg-amber-100 text-amber-700 dark:bg-amber-900/40 dark:text-amber-300',
    'rdv_planifie': 'bg-purple-100 text-purple-700 dark:bg-purple-900/40 dark:text-purple-300',
    'complete': 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/40 dark:text-emerald-300',
    'annule': 'bg-stone-200 text-stone-600 dark:bg-stone-700 dark:text-stone-400'
  }
  return classes[statut] || classes.nouveau
}

const getStatutLabel = (statut) => {
  const labels = {
    'nouveau': 'Nouveau',
    'en_attente': 'En attente',
    'rdv_planifie': 'RDV planifié',
    'complete': 'Complété',
    'annule': 'Annulé'
  }
  return labels[statut] || statut
}
</script>