<!-- frontend/src/components/Pharmacies/PharmacieFilters.vue -->
<template>
  <div class="bg-white dark:bg-gray-800 rounded-xl border border-gray-200 dark:border-gray-700 p-5 space-y-4 sticky top-24">
    <div class="flex items-center justify-between mb-4">
      <h3 class="font-bold text-gray-900 dark:text-white flex items-center gap-2">
        <Filter :size="20" />
        Filtres
      </h3>
      <button
        @click="$emit('reset')"
        class="text-xs text-slate-600 dark:text-sky-800 hover:underline"
      >
        Réinitialiser
      </button>
    </div>

    <!-- Nom de la pharmacie -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Nom de la pharmacie
      </label>
      <div class="relative">
        <Search :size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          v-model="localFilters.nom"
          type="text"
          placeholder="Rechercher..."
          @input="debouncedSearch"
          class="w-full pl-9 pr-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
        />
      </div>
    </div>

    <!-- Ville -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Ville
      </label>
      <div class="relative">
        <MapPin :size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          v-model="localFilters.ville"
          type="text"
          placeholder="Montréal, Québec..."
          @input="debouncedSearch"
          class="w-full pl-9 pr-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
        />
      </div>
      <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
        Min. 3 caractères
      </p>
    </div>

    <!-- Région -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Région
      </label>
      <!-- Région -->
<select
  v-model="localFilters.region"
  @change="handleSelectChange"  
  class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
>
        <option value="">Toutes les régions</option>
        <option value="Abitibi-Témiscamingue">Abitibi-Témiscamingue</option>
        <option value="Bas-Saint-Laurent">Bas-Saint-Laurent</option>
        <option value="Capitale-Nationale">Capitale-Nationale</option>
        <option value="Centre-du-Québec">Centre-du-Québec</option>
        <option value="Chaudière-Appalaches">Chaudière-Appalaches</option>
        <option value="Côte-Nord">Côte-Nord</option>
        <option value="Estrie">Estrie</option>
        <option value="Gaspésie–Îles-de-la-Madeleine">Gaspésie–Îles-de-la-Madeleine</option>
        <option value="Lanaudière">Lanaudière</option>
        <option value="Laurentides">Laurentides</option>
        <option value="Laval">Laval</option>
        <option value="Mauricie">Mauricie</option>
        <option value="Montérégie">Montérégie</option>
        <option value="Montréal">Montréal</option>
        <option value="Nord-du-Québec">Nord-du-Québec</option>
        <option value="Outaouais">Outaouais</option>
        <option value="Saguenay–Lac-Saint-Jean">Saguenay–Lac-Saint-Jean</option>
      </select>
    </div>

    <!-- Bannière -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Bannière
      </label>
      <select
        v-model="localFilters.banniere"
        @change="handleSelectChange"
        class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
      >
        <option value="">Toutes les bannières</option>
        <option value="Jean Coutu">Jean Coutu</option>
        <option value="Pharmaprix">Pharmaprix</option>
        <option value="Familiprix">Familiprix</option>
        <option value="Uniprix">Uniprix</option>
        <option value="Proxim">Proxim</option>
        <option value="Brunet">Brunet</option>
        <option value="Indépendante">Indépendante</option>
      </select>
    </div>

    <!-- Ouvert dimanche -->
    <div class="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700/50 rounded-lg">
      <label class="text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-2">
        <Calendar :size="16" />
        Ouvert le dimanche
      </label>
      <input
        v-model="localFilters.ouvertDimanche"
        type="checkbox"
        @change="$emit('search')"
        class="w-5 h-5 text-emerald-600 border-gray-300 dark:border-gray-600 rounded focus:ring-emerald-500"
      />
    </div>

    <!-- Bouton de recherche -->
    <button
      @click="$emit('search')"
      class="w-full px-4 py-2.5 bg-slate-600 hover:bg-sky-800 text-white font-semibold rounded-lg transition-colors flex items-center justify-center gap-2"
    >
      <Search :size="18" />
      Rechercher
    </button>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { Filter, Search, MapPin, Calendar } from 'lucide-vue-next'

const props = defineProps({
  filters: { type: Object, required: true }
})

const emit = defineEmits(['update:filters', 'search', 'reset'])

const localFilters = ref({ ...props.filters })



const handleSelectChange = () => {
  emit('update:filters', localFilters.value)
  emit('search')
}

// 1. On utilise UN SEUL nom pour la variable de temps
let debounceTimer = null

// 2. Ta fonction debouncedSearch
const debouncedSearch = () => {
  // On annule le délai précédent en utilisant le BON nom
  clearTimeout(debounceTimer) 
  
  // On lance le nouveau délai
  debounceTimer = setTimeout(() => {
    // On met à jour le parent et on lance la recherche
    emit('update:filters', localFilters.value)
    emit('search')
  }, 300)
}

// 3. Synchronisation (pour quand tu cliques sur "Réinitialiser" dans le parent)
watch(() => props.filters, (newVal) => {
  localFilters.value = { ...newVal }
}, { deep: true })

</script>