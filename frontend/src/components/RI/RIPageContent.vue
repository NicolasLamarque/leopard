<template>
  <div class="p-6">
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
      
      <!-- Sidebar - Filtres -->
      <div class="lg:col-span-1">
        <RIFilters 
          v-model:filters="filters"
          @search="handleSearch"
          @reset="handleResetFilters"
        />
      </div>

      <!-- Liste des RPA -->
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

            <div class="flex items-center gap-3">
              <button
                @click="showSyncPanel = true"
                class="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
              >
                <RefreshCw :size="18" />
                Synchroniser
              </button>

              <button
                @click="showNewRI = true"
                class="flex items-center gap-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
              >
                <Plus :size="18" />
                Ajouter
              </button>
            </div>
          </div>
        </div>

        <RPAList
          :ri="ris"
          :loading="loading"
          @select="handleSelectRI"
          @delete="handleDeleteRI"
        />
      </div>

    </div>

    <!-- Modal détails/édition RPA -->
    <RPADetails
      v-if="selectedRI"
      :rpa="selectedRI"
      :mode="detailsMode"
      @close="selectedRI = null"
      @save="handleSaveRI"
      @delete="handleDeleteRI"
    />

    <!-- Panel de synchronisation -->
    <RPASyncPanel
      v-if="showSyncPanel"
      @close="showSyncPanel = false"
      @synced="handleSynced"
    />

    <!-- Modal nouveau RI -->
    <RIDetails
      v-if="showNewRI"
      :rpa="null"
      mode="create"
      @close="showNewRI = false"
      @save="handleCreateRI"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { RefreshCw, Plus } from 'lucide-vue-next'
import RPAList from './RPAList.vue'
import RPADetails from './RIDetails.vue'
import RPAFilters from './RIFilters.vue'
import RPASyncPanel from './RISyncPanel.vue'
import { SearchResidences, DeleteResidence } from '../../../wailsjs/go/main/App'

const ris = ref([])
const loading = ref(false)
const selectedRI = ref(null)
const detailsMode = ref('view')
const showSyncPanel = ref(false)
const showNewRI = ref(false)
const allRIs = ref([])

const filters = ref({
  municipalite: '',
  nom: '',
  region: '',
  registre: '',
  statut: ''
})

const stats = computed(() => {
  const data = ris.value || []
  return {
    total: data.length,
    actifs: data.filter(r => r.statut === 'actif').length,
    fermes: data.filter(r => r.statut === 'ferme').length,
    suspendus: data.filter(r => r.statut === 'suspendu').length
  }
})

const handleSearch = () => {
  console.log('🔍 Filtrage avec:', filters.value)
  
  rpas.value = allRpas.value.filter(rpa => {
    if (filters.value.nom && !ri.titre.toLowerCase().includes(filters.value.nom.toLowerCase())) {
      return false
    }
    
    if (filters.value.municipalite && filters.value.municipalite.length >= 4) {
      if (!ri.municipalite.toLowerCase().includes(filters.value.municipalite.toLowerCase())) {
        return false
      }
    }
    
    if (filters.value.region && ri.region !== filters.value.region) {
      return false
    }
    
    if (filters.value.registre && !ri.registre.includes(filters.value.registre)) {
      return false
    }
    
    if (filters.value.statut && ri.statut !== filters.value.statut) {
      return false
    }
    
    return true
  })
  
  console.log('📊 Résultats filtrés:', ris.value.length)
}

const handleResetFilters = () => {
  filters.value = {
    municipalite: '',
    nom: '',
    region: '',
    registre: '',
    statut: ''
  }
  ris.value = [...allRIs.value]
}


const handleSelectRI = (ri) => {
  selectedRI.value = ri
  detailsMode.value = 'view'
}

const handleSaveRI = async (updatedRI) => {
  try {
    await UpdateRI(updatedRI)
    selectedRI.value = null
    await loadAllRIs()
    alert('✅ Ressource intermédiaire mise à jour!')
  } catch (err) {
    console.error('Erreur sauvegarde RI:', err)
    alert('Erreur lors de la sauvegarde')
  }
}

const handleCreateRI = async (newRI) => {
  try {
    await CreateRI(newRI)
    showNewRI.value = false
    await loadAllRIs()
    alert('✅ Ressource intermédiaire créée!')
  } catch (err) {
    console.error('Erreur création RI:', err)
    alert('Erreur lors de la création')
  }
}

const handleDeleteRI = async (ri) => {
  if (!confirm(`Supprimer "${ri.titre_rpa}" ?`)) return

  try {
    await DeleteRPA(ri.id)
    selectedRI.value = null
    await loadAllRIs()
    alert('✅ Ressource intermédiaire supprimée!')
  } catch (err) {
    console.error('Erreur suppression RI:', err)
    alert('Erreur lors de la suppression')
  }
}

const handleSynced = async () => {
  showSyncPanel.value = false
  await loadAllRIs()
}

const loadAllRIs = async () => {
  loading.value = true
  try {
    allIs.value = await SearchResidences('', '', '')
    ris.value = [...allIs.value]
    console.log('✅ RIs chargés:', allIs.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement RI:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadAllRIs() 
})
</script>