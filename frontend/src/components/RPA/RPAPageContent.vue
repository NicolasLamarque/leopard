<template>
  <div class="p-6">
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
      
      <!-- Sidebar - Filtres -->
      <div class="lg:col-span-1">
        <RPAFilters 
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
                @click="showNewRPA = true"
                class="flex items-center gap-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
              >
                <Plus :size="18" />
                Ajouter
              </button>
            </div>
          </div>
        </div>

        <RPAList
          :rpas="rpas"
          :loading="loading"
          @select="handleSelectRPA"
          @delete="handleDeleteRPA"
        />
      </div>

    </div>

    <!-- Modal détails/édition RPA -->
    <RPADetails
      v-if="selectedRPA"
      :rpa="selectedRPA"
      :mode="detailsMode"
      @close="selectedRPA = null"
      @save="handleSaveRPA"
      @delete="handleDeleteRPA"
    />

    <!-- Panel de synchronisation -->
    <RPASyncPanel
      v-if="showSyncPanel"
      @close="showSyncPanel = false"
      @synced="handleSynced"
    />

    <!-- Modal nouveau RPA -->
    <RPADetails
      v-if="showNewRPA"
      :rpa="null"
      mode="create"
      @close="showNewRPA = false"
      @save="handleCreateRPA"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { RefreshCw, Plus } from 'lucide-vue-next'
import RPAList from './RPAList.vue'
import RPADetails from './RPADetails.vue'
import RPAFilters from './RPAFilters.vue'
import RPASyncPanel from './RPASyncPanel.vue'
import { SearchResidences, DeleteResidence } from '../../../wailsjs/go/main/App'

const rpas = ref([])
const loading = ref(false)
const selectedRPA = ref(null)
const detailsMode = ref('view')
const showSyncPanel = ref(false)
const showNewRPA = ref(false)
const allRpas = ref([])

const filters = ref({
  municipalite: '',
  nom: '',
  region: '',
  registre: '',
  statut: ''
})

const stats = computed(() => {
  const data = rpas.value || []
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
    if (filters.value.nom && !rpa.titre.toLowerCase().includes(filters.value.nom.toLowerCase())) {
      return false
    }
    
    if (filters.value.municipalite && filters.value.municipalite.length >= 4) {
      if (!rpa.municipalite.toLowerCase().includes(filters.value.municipalite.toLowerCase())) {
        return false
      }
    }
    
    if (filters.value.region && rpa.region !== filters.value.region) {
      return false
    }
    
    if (filters.value.registre && !rpa.registre.includes(filters.value.registre)) {
      return false
    }
    
    if (filters.value.statut && rpa.statut !== filters.value.statut) {
      return false
    }
    
    return true
  })
  
  console.log('📊 Résultats filtrés:', rpas.value.length)
}

const handleResetFilters = () => {
  filters.value = {
    municipalite: '',
    nom: '',
    region: '',
    registre: '',
    statut: ''
  }
  rpas.value = [...allRpas.value]
}

const handleSelectRPA = (rpa) => {
  selectedRPA.value = rpa
  detailsMode.value = 'view'
}

const handleSaveRPA = async (updatedRPA) => {
  try {
    await UpdateRPA(updatedRPA)
    selectedRPA.value = null
    await loadAllRPAs()
    alert('✅ RPA mis à jour!')
  } catch (err) {
    console.error('Erreur sauvegarde RPA:', err)
    alert('Erreur lors de la sauvegarde')
  }
}

const handleCreateRPA = async (newRPA) => {
  try {
    await CreateRPA(newRPA)
    showNewRPA.value = false
    await loadAllRPAs()
    alert('✅ RPA créé!')
  } catch (err) {
    console.error('Erreur création RPA:', err)
    alert('Erreur lors de la création')
  }
}

const handleDeleteRPA = async (rpa) => {
  if (!confirm(`Supprimer "${rpa.titre_rpa}" ?`)) return
  
  try {
    await DeleteRPA(rpa.id)
    selectedRPA.value = null
    await loadAllRPAs()
    alert('✅ RPA supprimé!')
  } catch (err) {
    console.error('Erreur suppression RPA:', err)
    alert('Erreur lors de la suppression')
  }
}

const handleSynced = async () => {
  showSyncPanel.value = false
  await loadAllRPAs()
}

const loadAllRPAs = async () => {
  loading.value = true
  try {
    allRpas.value = await SearchResidences('', '', '')
    rpas.value = [...allRpas.value]
    console.log('✅ RPA chargés:', allRpas.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement RPA:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadAllRPAs() 
})
</script>