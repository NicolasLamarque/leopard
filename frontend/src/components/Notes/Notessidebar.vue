<template>
  <div class="w-80 border-r dark:border-gray-800 bg-slate-50 dark:bg-gray-900/20 flex flex-col">
    <!-- Header du sidebar -->
    <div class="p-4 border-b dark:border-gray-800 bg-slate-100/50 dark:bg-gray-900/50">
      <div class="flex items-center justify-between mb-3">
        <span class="text-[10px] font-bold text-gray-400 uppercase tracking-widest">
          Historique des notes
        </span>
        <button 
          v-if="notes.length > 0"
          @click="$emit('toggle-select-all')"
          class="text-xs text-teal-600 dark:text-teal-400 hover:underline font-medium"
        >
          {{ allSelected ? 'Désélectionner' : 'Tout sélectionner' }}
        </button>
      </div>

      <!-- Barre de recherche -->
      <div class="relative">
        <Search :size="14" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          type="text"
          :value="searchQuery"
          @input="$emit('update:search-query', $event.target.value)"
          placeholder="Rechercher..."
          class="w-full pl-9 pr-3 py-2 bg-white dark:bg-gray-900 border border-gray-300 dark:border-gray-800 rounded-lg text-sm outline-none focus:ring-2 focus:ring-teal-500"
        />
      </div>
    </div>

    <!-- Filtres par type -->
    <div class="p-3 border-b dark:border-gray-800 bg-white dark:bg-gray-900/30">
      <div class="flex flex-wrap gap-2">
        <button
          @click="$emit('filter-type', null)"
          :class="[
            'px-3 py-1 rounded-lg text-xs font-medium transition-all',
            !filterType 
              ? 'bg-teal-100 dark:bg-teal-900/50 text-teal-700 dark:text-teal-300 border-2 border-teal-300 dark:border-teal-700'
              : 'bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-400 hover:bg-gray-200 dark:hover:bg-gray-700'
          ]"
        >
          Toutes
        </button>
        <button
          v-for="type in noteTypes"
          :key="type"
          @click="$emit('filter-type', type)"
          :class="[
            'px-3 py-1 rounded-lg text-xs font-medium transition-all',
            filterType === type
              ? 'bg-teal-100 dark:bg-teal-900/50 text-teal-700 dark:text-teal-300 border-2 border-teal-300 dark:border-teal-700'
              : 'bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-400 hover:bg-gray-200 dark:hover:bg-gray-700'
          ]"
        >
          {{ type }}
        </button>
      </div>
    </div>

    <!-- Stats -->
    <div class="p-3 border-b dark:border-gray-800 bg-gradient-to-r from-slate-50 to-teal-50 dark:from-slate-900/20 dark:to-teal-900/20">
      <div class="grid grid-cols-3 gap-2 text-center">
        <div>
          <div class="text-lg font-bold text-slate-700 dark:text-slate-300">{{ totalNotes }}</div>
          <div class="text-[9px] text-gray-500 uppercase">Total</div>
        </div>
        <div>
          <div class="text-lg font-bold text-teal-600 dark:text-teal-400">{{ selectedNotes.length }}</div>
          <div class="text-[9px] text-gray-500 uppercase">Sélection</div>
        </div>
        <div>
          <div class="text-lg font-bold text-sky-600 dark:text-sky-400">{{ draftsCount }}</div>
          <div class="text-[9px] text-gray-500 uppercase">Brouillons</div>
        </div>
      </div>
    </div>

    <!-- Liste des notes -->
    <div class="flex-1 overflow-y-auto p-4 space-y-3">
      <div v-if="filteredNotes.length === 0" class="text-center py-8 text-gray-400 text-sm">
        <ClipboardList :size="32" class="mx-auto mb-2 opacity-30" />
        <p>Aucune note</p>
      </div>

      <div 
        v-for="note in filteredNotes" 
        :key="note.id"
        :class="[
          'p-3 rounded-xl border dark:border-gray-800 bg-white dark:bg-gray-900 cursor-pointer transition-all group relative',
          selectedNote?.id === note.id 
            ? 'ring-2 ring-teal-500 border-teal-200' 
            : 'hover:ring-2 hover:ring-teal-300'
        ]"
      >
        <div class="flex items-start gap-3">
          <!-- Checkbox -->
          <input 
            type="checkbox"
            :checked="selectedNotes.includes(note.id)"
            @change="$emit('toggle-note', note.id)"
            @click.stop
            class="mt-1 w-4 h-4 text-teal-600 rounded border-gray-300 focus:ring-teal-500 cursor-pointer"
          />
          
         <div @click="$emit('view-note', note)" class="flex-1 min-w-0">
  <div class="flex justify-between items-start mb-1">
    <span :class="[
      'text-[10px] font-bold uppercase px-2 py-0.5 rounded',
      getTypeColor(note.type_note)
    ]">
      {{ note.type_note }}
    </span>
    <span class="text-[10px] text-gray-400">{{ formatDate(note.date_note) }}</span>
  </div>
  
  <p class="text-sm font-bold text-gray-800 dark:text-gray-200 truncate mb-1">
    {{ note.titre || 'Sans titre' }} </p>

  <div v-if="note.note_liee_id" class="flex items-center gap-1 mb-2">
    <Link2 :size="10" class="text-sky-500" />
    <span class="text-[9px] text-sky-600 dark:text-sky-400">
      {{ note.type_lien || 'Lié' }} à: {{ note.note_liee_titre || 'Note #' + note.note_liee_id }}
    </span>
  </div>
  
  <div class="flex items-center justify-between">
    <span class="text-[9px] text-gray-400 italic truncate">
      {{ note.signature_nom?.split(',')[0] || 'Non signé' }}
    </span>
    <div class="flex items-center gap-1">
      <Lock v-if="note.verrouille === true || note.verrouille === 1" :size="12" class="text-green-500" />
      <FileEdit v-else :size="12" class="text-orange-500" />
    </div>
  </div>
</div>
        </div>

        <!-- Badge brouillon si non verrouillé -->
        <div v-if="!note.verrouille" class="absolute -top-1 -right-1">
          <span class="inline-flex items-center gap-1 bg-orange-500 text-white text-[8px] px-1.5 py-0.5 rounded-full font-bold">
            <FileEdit :size="8" />
            BROUILLON
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Search, ClipboardList, Lock, FileEdit, Link2 } from 'lucide-vue-next'

const props = defineProps({
  notes: { type: Array, default: () => [] },
  selectedNote: Object,
  selectedNotes: { type: Array, default: () => [] },
  searchQuery: { type: String, default: '' },
  filterType: String,
  noteTypes: { type: Array, default: () => ['Suivi', 'Ouverture', 'Plan', 'Fermeture', 'Incident', 'Correction'] }
})

defineEmits(['view-note', 'toggle-note', 'toggle-select-all', 'filter-type', 'update:search-query'])

const totalNotes = computed(() => props.notes.length)
const draftsCount = computed(() => props.notes.filter(n => !n.verrouille).length)
const allSelected = computed(() => 
  props.notes.length > 0 && props.selectedNotes.length === props.notes.length
)

const filteredNotes = computed(() => {
  let result = props.notes

  // Filtre par type
  if (props.filterType) {
    result = result.filter(n => n.type_note === props.filterType)
  }

  // Filtre par recherche
  if (props.searchQuery) {
    const query = props.searchQuery.toLowerCase()
    result = result.filter(n => 
      n.titre?.toLowerCase().includes(query) ||
      n.contenu?.toLowerCase().includes(query) ||
      n.signature_nom?.toLowerCase().includes(query)
    )
  }

  return result
})

const getTypeColor = (type) => {
  const colors = {
    'Suivi': 'bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300',
    'Ouverture': 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300',
    'Plan': 'bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300',
    'Fermeture': 'bg-gray-100 dark:bg-gray-900/30 text-gray-700 dark:text-gray-300',
    'Incident': 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300',
    'Correction': 'bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300',
    'Addendum': 'bg-sky-100 dark:bg-sky-900/30 text-sky-700 dark:text-sky-300'
  }
  return colors[type] || 'bg-teal-100 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300'
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA')
}
</script>