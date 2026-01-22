<template>
  <div class="chsld-filters bg-white dark:bg-gray-800 rounded-lg shadow-sm p-4 sticky top-20">
    <div class="flex items-center justify-between mb-4">
      <h3 class="font-semibold text-gray-900 dark:text-white flex items-center gap-2">
        <Filter :size="18" />
        Filtres
      </h3>
      <button
        @click="handleReset"
        class="text-sm text-blue-600 dark:text-blue-400 hover:underline"
      >
        Réinitialiser
      </button>
    </div>

    <div class="space-y-4">
      <!-- Nom du CHSLD -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
          Nom du CHSLD
        </label>
        <input
          v-model="localFilters.nom"
          @input="emitFilters"
          type="text"
          placeholder="Rechercher..."
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white text-sm"
        />
      </div>

      <!-- Municipalité -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
          Municipalité
        </label>
        <input
          v-model="localFilters.municipalite"
          @input="emitFilters"
          type="text"
          placeholder="Min. 4 caractères..."
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white text-sm"
        />
        <p v-if="localFilters.municipalite && localFilters.municipalite.length < 4" 
           class="text-xs text-amber-600 dark:text-amber-400 mt-1">
          Minimum 4 caractères requis
        </p>
      </div>

      <!-- Région -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
          Région
        </label>
        <select
          v-model="localFilters.region"
          @change="emitFilters"
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white text-sm"
        >
          <option value="">Toutes les régions</option>
          <option value="01">01 - Bas-Saint-Laurent</option>
          <option value="02">02 - Saguenay–Lac-Saint-Jean</option>
          <option value="03">03 - Capitale-Nationale</option>
          <option value="04">04 - Mauricie</option>
          <option value="05">05 - Estrie</option>
          <option value="06">06 - Montréal</option>
          <option value="07">07 - Outaouais</option>
          <option value="08">08 - Abitibi-Témiscamingue</option>
          <option value="09">09 - Côte-Nord</option>
          <option value="10">10 - Nord-du-Québec</option>
          <option value="11">11 - Gaspésie–Îles-de-la-Madeleine</option>
          <option value="12">12 - Chaudière-Appalaches</option>
          <option value="13">13 - Laval</option>
          <option value="14">14 - Lanaudière</option>
          <option value="15">15 - Laurentides</option>
          <option value="16">16 - Montérégie</option>
          <option value="17">17 - Centre-du-Québec</option>
        </select>
      </div>

      <!-- Type de CHSLD -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
          Type
        </label>
        <select
          v-model="localFilters.type"
          @change="emitFilters"
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white text-sm"
        >
          <option value="">Tous les types</option>
          <option value="Public">Public</option>
          <option value="Privé conventionné">Privé conventionné</option>
          <option value="Privé non conventionné">Privé non conventionné</option>
        </select>
      </div>

      <!-- Statut -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
          Statut
        </label>
        <select
          v-model="localFilters.statut"
          @change="emitFilters"
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white text-sm"
        >
          <option value="">Tous les statuts</option>
          <option value="Actif">Actif</option>
          <option value="Fermé">Fermé</option>
          <option value="Suspendu">Suspendu</option>
          <option value="En rénovation">En rénovation</option>
        </select>
      </div>

      <!-- Capacité minimale -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
          Capacité minimale (lits)
        </label>
        <input
          v-model.number="localFilters.capaciteMin"
          @input="emitFilters"
          type="number"
          min="0"
          placeholder="0"
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white text-sm"
        />
      </div>
    </div>

    <!-- Bouton rechercher -->
    <button
      @click="handleSearch"
      class="w-full mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center justify-center gap-2"
    >
      <Search :size="18" />
      Rechercher
    </button>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { Filter, Search } from 'lucide-vue-next'

const props = defineProps({
  filters: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['update:filters', 'search', 'reset'])

const localFilters = ref({ ...props.filters })

// Synchroniser avec les props
watch(() => props.filters, (newVal) => {
  localFilters.value = { ...newVal }
}, { deep: true })

// Émettre les changements de filtres
const emitFilters = () => {
  emit('update:filters', { ...localFilters.value })
}

// Rechercher
const handleSearch = () => {
  emit('search')
}

// Réinitialiser
const handleReset = () => {
  localFilters.value = {
    nom: '',
    municipalite: '',
    region: '',
    type: '',
    statut: '',
    capaciteMin: 0
  }
  emit('update:filters', { ...localFilters.value })
  emit('reset')
}
</script>

<style scoped>

</style>