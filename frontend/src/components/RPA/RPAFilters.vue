<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6 sticky top-24">
    <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
      <Filter :size="20" />
      Filtres
    </h3>

    <!-- Recherche rapide -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Recherche rapide
      </label>
      <input
        v-model="localFilters.nom"
        @input="debouncedSearch"
        type="text"
        placeholder="Nom du RPA..."
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
      />
    </div>

    <!-- Municipalité -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Municipalité
      </label>
      <input
        v-model="localFilters.municipalite"
        @input="debouncedSearch"
        type="text"
        placeholder="Ex: Montréal, Québec..."
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
      />
    </div>

    <!-- Région -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Région
      </label>
      <select
        v-model="localFilters.region"
        @change="handleSearch"
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
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

    <!-- Statut -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Statut
      </label>
      <div class="space-y-2">
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="radio"
            v-model="localFilters.statut"
            value=""
            @change="handleSearch"
            class="text-blue-600 focus:ring-2 focus:ring-blue-500"
          />
          <span class="text-sm text-gray-700 dark:text-gray-300">Tous</span>
        </label>
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="radio"
            v-model="localFilters.statut"
            value="actif"
            @change="handleSearch"
            class="text-blue-600 focus:ring-2 focus:ring-blue-500"
          />
          <span class="flex items-center gap-2 text-sm">
            <span class="w-3 h-3 bg-green-500 rounded-full"></span>
            Actifs
          </span>
        </label>
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="radio"
            v-model="localFilters.statut"
            value="ferme"
            @change="handleSearch"
            class="text-blue-600 focus:ring-2 focus:ring-blue-500"
          />
          <span class="flex items-center gap-2 text-sm">
            <span class="w-3 h-3 bg-red-500 rounded-full"></span>
            Fermés
          </span>
        </label>
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="radio"
            v-model="localFilters.statut"
            value="suspendu"
            @change="handleSearch"
            class="text-blue-600 focus:ring-2 focus:ring-blue-500"
          />
          <span class="flex items-center gap-2 text-sm">
            <span class="w-3 h-3 bg-yellow-500 rounded-full"></span>
            Suspendus
          </span>
        </label>
      </div>
    </div>

    <!-- Boutons d'action -->
    <div class="pt-4 space-y-2">
      <button
        @click="handleReset"
        class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors text-sm"
      >
        Réinitialiser
      </button>
    </div>

    <!-- Compteur de résultats -->
    <div class="pt-4 border-t dark:border-gray-700">
      <p class="text-sm text-gray-500 dark:text-gray-400 text-center">
        Filtres actifs: <span class="font-semibold text-blue-600 dark:text-blue-400">{{ activeFiltersCount }}</span>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Filter } from 'lucide-vue-next'

const props = defineProps({
  filters: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['update:filters', 'search', 'reset'])

const localFilters = ref({ ...props.filters })

// Debounce pour la recherche
let debounceTimer = null
const debouncedSearch = () => {
  clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    handleSearch()
  }, 500)
}

const handleSearch = () => {
  emit('update:filters', localFilters.value)
  emit('search')
}

const handleReset = () => {
  localFilters.value = {
    municipalite: '',
    nom: '',
    region: '',
    statut: ''
  }
  emit('update:filters', localFilters.value)
  emit('reset')
}

// Compteur de filtres actifs
const activeFiltersCount = computed(() => {
  let count = 0
  if (localFilters.value.municipalite) count++
  if (localFilters.value.nom) count++
  if (localFilters.value.region) count++
  if (localFilters.value.statut) count++
  return count
})

// Sync avec le parent
watch(() => props.filters, (newFilters) => {
  localFilters.value = { ...newFilters }
}, { deep: true })
</script>