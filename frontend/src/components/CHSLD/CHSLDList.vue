<template>
  <div class="chsld-list space-y-4">
    <!-- Barre de recherche et actions -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm p-4 mb-4">
      <div class="flex items-center justify-between">
        <div class="flex-1 mr-4">
          <input
            type="text"
            v-model="searchQuery"
            @input="handleSearch"
            placeholder="Rechercher un CHSLD..."
            class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white"
          />
        </div>
        <div class="flex items-center gap-2">
          <select
            v-model="sortBy"
            @change="handleSort"
            class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg dark:bg-gray-700 dark:text-white"
          >
            <option value="TitreCHSLD">Nom</option>
            <option value="Municipalite">Municipalité</option>
            <option value="Capacite">Capacité</option>
            <option value="Statut">Statut</option>
            <option value="Region">Région</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-12">
      <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      <p class="mt-4 text-gray-600 dark:text-gray-400">Chargement...</p>
    </div>

    <!-- Empty state -->
    <div v-else-if="chslds.length === 0" class="bg-white dark:bg-gray-800 rounded-lg shadow-sm p-12 text-center">
      <Building2 :size="48" class="mx-auto text-gray-400 mb-4" />
      <p class="text-gray-600 dark:text-gray-400">Aucun CHSLD trouvé</p>
    </div>

    <!-- Liste des CHSLD -->
    <div v-else class="grid grid-cols-1 gap-4">
      <div
        v-for="chsld in chslds"
        :key="chsld.id"
        @click="$emit('select', chsld)"
        class="bg-white dark:bg-gray-800 rounded-lg shadow-sm hover:shadow-md transition-all cursor-pointer border border-gray-200 dark:border-gray-700 hover:border-blue-500"
      >
        <div class="p-4">
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <!-- Titre et statut -->
              <div class="flex items-center gap-3 mb-2">
                <Building2 :size="24" class="text-blue-600 dark:text-blue-400" />
                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                  {{ chsld.TitreCHSLD }}
                </h3>
                <span
                  v-if="chsld.Statut"
                  :class="getStatutClass(chsld.Statut)"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ chsld.Statut }}
                </span>
              </div>

              <!-- Informations principales -->
              <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-3">
                <div class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
                  <MapPin :size="16" />
                  <span>{{ chsld.Municipalite || 'N/A' }}</span>
                </div>

                <div class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
                  <Map :size="16" />
                  <span>{{ chsld.Region || 'N/A' }}</span>
                </div>

                <div class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
                  <Users :size="16" />
                  <span>{{ chsld.Capacite || 0 }} lits</span>
                </div>

                <div class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
                  <Building :size="16" />
                  <span>{{ chsld.TypeCHSLD || 'N/A' }}</span>
                </div>
              </div>

              <!-- Adresse -->
              <div v-if="chsld.Adresse" class="mt-3 flex items-start gap-2 text-sm text-gray-500 dark:text-gray-400">
                <Home :size="16" class="mt-0.5 flex-shrink-0" />
                <span>{{ chsld.Adresse }}</span>
              </div>

              <!-- Téléphone -->
              <div v-if="chsld.Telephone" class="mt-2 flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
                <Phone :size="16" />
                <span>{{ chsld.Telephone }}</span>
              </div>

              <!-- Propriétaire -->
              <div v-if="chsld.Proprietaire" class="mt-2 flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
                <Briefcase :size="16" />
                <span>{{ chsld.Proprietaire }}</span>
              </div>
            </div>

            <!-- Bouton supprimer -->
            <button
              @click.stop="$emit('delete', chsld)"
              class="ml-4 p-2 text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
              title="Supprimer"
            >
              <Trash2 :size="18" />
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="mt-6 flex justify-center gap-2">
      <button
        v-for="page in totalPages"
        :key="page"
        @click="currentPage = page"
        :class="[
          'px-4 py-2 rounded-lg transition-colors',
          currentPage === page
            ? 'bg-blue-600 text-white'
            : 'bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'
        ]"
      >
        {{ page }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Building2, MapPin, Map, Users, Building, Home, Phone, Briefcase, Trash2 } from 'lucide-vue-next'

const props = defineProps({
  chslds: {
    type: Array,
    required: true
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['select', 'delete', 'search', 'sort'])

const searchQuery = ref('')
const sortBy = ref('TitreCHSLD')
const currentPage = ref(1)
const itemsPerPage = 10

const totalPages = computed(() => {
  return Math.ceil(props.chslds.length / itemsPerPage)
})

const handleSearch = () => {
  emit('search', searchQuery.value)
}

const handleSort = () => {
  emit('sort', sortBy.value)
}

const getStatutClass = (statut) => {
  const statutLower = statut?.toLowerCase() || ''
  
  if (statutLower.includes('actif') || statutLower.includes('ouvert')) {
    return 'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400'
  }
  if (statutLower.includes('fermé') || statutLower.includes('ferme')) {
    return 'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400'
  }
  if (statutLower.includes('suspendu')) {
    return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400'
  }
  
  return 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-400'
}
</script>

<style scoped></style>