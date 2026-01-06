<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6 sticky top-24">
    <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
      <Filter :size="20" />
      Recherche avancée
    </h3>

    <!-- 1. Nom de la résidence -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Nom de la résidence
      </label>
      <input
        v-model="localFilters.nom"
        @input="debouncedSearch"
        type="text"
        placeholder="Ex: Manoir, Villa..."
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
      />
    </div>

    <!-- 2. Municipalité -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Municipalité
        <span class="text-xs text-gray-400">(min 4 caractères)</span>
      </label>
      <input
        v-model="localFilters.municipalite"
        @input="debouncedSearch"
        type="text"
        placeholder="Ex: Montréal, Québec..."
        minlength="4"
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
      />
      <p v-if="localFilters.municipalite.length > 0 && localFilters.municipalite.length < 4" class="text-xs text-orange-500 mt-1">
        ⚠️ Minimum 4 caractères requis
      </p>
    </div>

    <!-- 3. Région sociosanitaire -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Région sociosanitaire
      </label>
      <select
        v-model="localFilters.region"
        @change="handleSearch"
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
      >
        <option value="">Toutes les régions</option>
        <option value="Abitibi-Témiscamingue">01 - Abitibi-Témiscamingue</option>
        <option value="Bas-Saint-Laurent">02 - Bas-Saint-Laurent</option>
        <option value="Capitale-Nationale">03 - Capitale-Nationale</option>
        <option value="Centre-du-Québec">04 - Centre-du-Québec</option>
        <option value="Chaudière-Appalaches">05 - Chaudière-Appalaches</option>
        <option value="Côte-Nord">06 - Côte-Nord</option>
        <option value="Estrie">07 - Estrie</option>
        <option value="Gaspésie—Îles-de-la-Madeleine">08 - Gaspésie—Îles-de-la-Madeleine</option>
        <option value="Lanaudière">09 - Lanaudière</option>
        <option value="Laurentides">10 - Laurentides</option>
        <option value="Laval">11 - Laval</option>
        <option value="Mauricie">12 - Mauricie</option>
        <option value="Montérégie">13 - Montérégie</option>
        <option value="Montréal">14 - Montréal</option>
        <option value="Nord-du-Québec">15 - Nord-du-Québec</option>
        <option value="Outaouais">16 - Outaouais</option>
        <option value="Saguenay—Lac-Saint-Jean">17 - Saguenay—Lac-Saint-Jean</option>
      </select>
    </div>

    <!-- 4. Numéro au registre -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Numéro au registre
      </label>
      <input
        v-model="localFilters.registre"
        @input="debouncedSearch"
        type="text"
        placeholder="Ex: 6540"
        class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 font-mono"
      />
    </div>

    <!-- 5. Statut (TON FILTRE PERSO) -->
    <div>
      <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
        Statut
        <span class="text-xs text-gray-400">(filtrage local)</span>
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
    // Ne rechercher que si municipalité vide OU >= 4 caractères
    if (localFilters.value.municipalite === '' || localFilters.value.municipalite.length >= 4) {
      handleSearch()
    }
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
    registre: '',
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
  if (localFilters.value.registre) count++
  if (localFilters.value.statut) count++
  return count
})

// Sync avec le parent
watch(() => props.filters, (newFilters) => {
  localFilters.value = { ...newFilters }
}, { deep: true })
</script>