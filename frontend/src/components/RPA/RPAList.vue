<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow">
    <!-- En-tête -->
    <div class="px-6 py-4 border-b dark:border-gray-700 flex items-center justify-between">
      <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
        Liste des RPA ({{ rpas?.length || 0 }})
      </h2>
      <div class="flex items-center gap-2">
        <button
          @click="viewMode = 'grid'"
          :class="viewMode === 'grid' ? 'bg-blue-100 dark:bg-blue-900/30 text-blue-600' : 'text-gray-400'"
          class="p-2 rounded hover:bg-gray-100 dark:hover:bg-gray-700"
        >
          <LayoutGrid :size="18" />
        </button>
        <button
          @click="viewMode = 'table'"
          :class="viewMode === 'table' ? 'bg-blue-100 dark:bg-blue-900/30 text-blue-600' : 'text-gray-400'"
          class="p-2 rounded hover:bg-gray-100 dark:hover:bg-gray-700"
        >
          <List :size="18" />
        </button>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="p-12 text-center">
      <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-blue-500 border-t-transparent"></div>
      <p class="mt-4 text-gray-500 dark:text-gray-400">Chargement...</p>
    </div>

    <!-- Vide -->
    <div v-else-if="!rpas || rpas.length === 0" class="p-12 text-center">
      <Building2 :size="64" class="mx-auto text-gray-300 dark:text-gray-600 mb-4" />
      <p class="text-gray-500 dark:text-gray-400 text-lg">Aucun RPA trouvé</p>
      <p class="text-sm text-gray-400 dark:text-gray-500 mt-2">
        Essayez de modifier vos filtres ou lancez une synchronisation
      </p>
    </div>

    <!-- Vue Grille -->
    <div v-else-if="viewMode === 'grid'" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-6">
      <div
        v-for="rpa in rpas"
        :key="rpa.id"
        @click="$emit('select', rpa)"
        class="border dark:border-gray-700 rounded-lg p-4 hover:shadow-lg hover:border-blue-500 dark:hover:border-blue-400 cursor-pointer transition-all group"
      >
        <div class="flex items-start justify-between mb-3">
          <span
            :class="getStatutClass(rpa.statut)"
            class="px-2 py-1 rounded-full text-xs font-medium"
          >
            {{ rpa.statut }}
          </span>
          <button
            @click.stop="$emit('delete', rpa)"
            class="opacity-0 group-hover:opacity-100 text-red-500 hover:text-red-700 transition-opacity"
          >
            <Trash2 :size="16" />
          </button>
        </div>

        <h3 class="font-semibold text-gray-900 dark:text-white mb-2 line-clamp-2">
          {{ rpa.titre }}
        </h3>

        <div class="space-y-1 text-sm text-gray-600 dark:text-gray-400">
          <p class="flex items-center gap-2">
            <MapPin :size="14" class="text-gray-400" />
            {{ rpa.municipalite }}
          </p>
          <p v-if="rpa.adresse" class="flex items-center gap-2 line-clamp-1">
            <Home :size="14" class="text-gray-400" />
            {{ rpa.adresse }}
          </p>
          <p class="flex items-center gap-2 text-xs text-gray-500">
            <Hash :size="14" class="text-gray-400" />
            {{ rpa.registre }}
          </p>
        </div>
      </div>
    </div>

    <!-- Vue Tableau -->
    <div v-else class="overflow-x-auto">
      <table class="w-full">
        <thead class="bg-gray-50 dark:bg-gray-900">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">
              Statut
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">
              Nom
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">
              Municipalité
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">
              Région
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">
              Registre
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
          <tr
            v-for="rpa in rpas"
            :key="rpa.id"
            @click="$emit('select', rpa)"
            class="hover:bg-gray-50 dark:hover:bg-gray-700/50 cursor-pointer transition-colors"
          >
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatutClass(rpa.statut)" class="px-2.5 py-0.5 rounded-full text-xs font-medium">
                {{ rpa.statut }}
              </span>
            </td>
            <td class="px-6 py-4">
              <div class="font-medium text-gray-900 dark:text-white">
                {{ rpa.titre }}
              </div>
              <div v-if="rpa.adresse" class="text-xs text-gray-500 dark:text-gray-400 mt-1">
                {{ rpa.adresse }}
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-white">
              {{ rpa.municipalite }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
              {{ rpa.region }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-mono text-gray-500 dark:text-gray-400">
              {{ rpa.registre }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <div class="flex items-center gap-2">
                <button
                  @click.stop="$emit('select', rpa)"
                  class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300"
                  title="Voir détails"
                >
                  <Eye :size="18" />
                </button>
                <button
                  @click.stop="$emit('delete', rpa)"
                  class="text-red-600 dark:text-red-400 hover:text-red-800 dark:hover:text-red-300"
                  title="Supprimer"
                >
                  <Trash2 :size="18" />
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Building2, Eye, Trash2, MapPin, Home, Hash, LayoutGrid, List } from 'lucide-vue-next'

defineProps({
  rpas: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  }
})

defineEmits(['select', 'delete'])

const viewMode = ref('table') // 'table' ou 'grid'

const getStatutClass = (statut) => {
  const classes = {
    'actif': 'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400',
    'ferme': 'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400',
    'suspendu': 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400'
  }
  return classes[statut] || 'bg-gray-100 text-gray-800 dark:bg-gray-900/30 dark:text-gray-400'
}
</script>