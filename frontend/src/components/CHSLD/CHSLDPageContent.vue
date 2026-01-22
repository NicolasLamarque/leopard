<template>
  <div class="p-6">
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
      
      <!-- Sidebar - Filtres -->
      <div class="lg:col-span-1">
        <CHSLDFilters 
          v-model:filters="filters"
          @search="handleSearch"
          @reset="handleResetFilters"
        />
      </div>

      <!-- Liste des CHSLD -->
      <div class="lg:col-span-3">
        <!-- Actions et stats -->
        <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm p-4 mb-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div class="text-center">
                <p class="text-2xl font-bold text-green-600 dark:text-green-400">{{ stats.actifs }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Actifs</p>
              </div>
              <div class="text-center">
                <p class="text-2xl font-bold text-red-600 dark:text-red-400">{{ stats.fermes }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Fermés</p>
              </div>
              <div class="text-center">
                <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ stats.total }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Total</p>
              </div>
            </div>

            <button
              @click="showNewCHSLD = true"
              class="flex items-center gap-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
            >
              <Plus :size="18" />
              Ajouter
            </button>
          </div>
        </div>

        <CHSLDList
          :chslds="chslds"
          :loading="loading"
          @select="handleSelectCHSLD"
          @delete="handleDeleteCHSLD"
        />
      </div>

    </div>

    <!-- Modal détails/édition CHSLD -->
    <CHSLDDetails
      v-if="selectedCHSLD"
      :chsld="selectedCHSLD"
      :mode="detailsMode"
      @close="selectedCHSLD = null"
      @save="handleSaveCHSLD"
      @delete="handleDeleteCHSLD"
    />

    <!-- Modal nouveau CHSLD -->
    <CHSLDDetails
      v-if="showNewCHSLD"
      :chsld="null"
      mode="create"
      @close="showNewCHSLD = false"
      @save="handleCreateCHSLD"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { Plus } from 'lucide-vue-next'
import CHSLDList from '../CHSLD/CHSLDList.vue'
import CHSLDDetails from '../CHSLD/CHSLDDetails.vue'
import CHSLDFilters from '../CHSLD/CHSLDFilters.vue'
import { SearchCHSLD, CreateCHSLD, UpdateCHSLD, DeleteCHSLD } from '../../../wailsjs/go/main/App'

const chslds = ref([])
const loading = ref(false)
const selectedCHSLD = ref(null)
const detailsMode = ref('view')
const showNewCHSLD = ref(false)
const allChslds = ref([])

const filters = ref({
  nom: '',
  municipalite: '',
  region: '',
  type: '',
  statut: '',
  capaciteMin: 0
})

const stats = computed(() => {
  const data = chslds.value || []
  return {
    total: data.length,
    actifs: data.filter(c => c.Statut?.toLowerCase().includes('actif')).length,
    fermes: data.filter(c => c.Statut?.toLowerCase().includes('fermé')).length
  }
})

const handleSearch = () => {
  chslds.value = allChslds.value.filter(chsld => {
    if (filters.value.nom && !chsld.TitreCHSLD?.toLowerCase().includes(filters.value.nom.toLowerCase())) {
      return false
    }
    if (filters.value.municipalite && filters.value.municipalite.length >= 4) {
      if (!chsld.Municipalite?.toLowerCase().includes(filters.value.municipalite.toLowerCase())) {
        return false
      }
    }
    if (filters.value.region && chsld.Region !== filters.value.region) {
      return false
    }
    if (filters.value.type && chsld.TypeCHSLD !== filters.value.type) {
      return false
    }
    if (filters.value.statut && chsld.Statut !== filters.value.statut) {
      return false
    }
    if (filters.value.capaciteMin > 0 && (chsld.Capacite || 0) < filters.value.capaciteMin) {
      return false
    }
    return true
  })
}

const handleResetFilters = () => {
  filters.value = {
    nom: '',
    municipalite: '',
    region: '',
    type: '',
    statut: '',
    capaciteMin: 0
  }
  chslds.value = [...allChslds.value]
}

const handleSelectCHSLD = (chsld) => {
  selectedCHSLD.value = chsld
  detailsMode.value = 'view'
}

const handleSaveCHSLD = async (updatedCHSLD) => {
  try {
    await UpdateCHSLD(updatedCHSLD)
    selectedCHSLD.value = null
    await loadAllCHSLDs()
    alert('✅ CHSLD mis à jour!')
  } catch (err) {
    console.error('Erreur sauvegarde CHSLD:', err)
    alert('❌ Erreur lors de la sauvegarde')
  }
}

const handleCreateCHSLD = async (newCHSLD) => {
  try {
    await CreateCHSLD(newCHSLD)
    showNewCHSLD.value = false
    await loadAllCHSLDs()
    alert('✅ CHSLD créé!')
  } catch (err) {
    console.error('Erreur création CHSLD:', err)
    alert('❌ Erreur lors de la création')
  }
}

const handleDeleteCHSLD = async (chsld) => {
  if (!confirm(`Supprimer "${chsld.TitreCHSLD}" ?`)) return
  
  try {
    await DeleteCHSLD(chsld.id)
    selectedCHSLD.value = null
    await loadAllCHSLDs()
    alert('✅ CHSLD supprimé!')
  } catch (err) {
    console.error('Erreur suppression CHSLD:', err)
    alert('❌ Erreur lors de la suppression')
  }
}

const loadAllCHSLDs = async () => {
  loading.value = true
  try {
    allChslds.value = await SearchCHSLD('', '', '')
    chslds.value = [...allChslds.value]
    console.log('✅ CHSLD chargés:', allChslds.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement CHSLD:', err)
    allChslds.value = []
    chslds.value = []
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadAllCHSLDs()
})
</script>