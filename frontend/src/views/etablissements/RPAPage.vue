<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <!-- Header -->
    <div class="bg-white dark:bg-gray-800 border-b dark:border-gray-700 sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-6 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <Building2 :size="32" class="text-blue-600 dark:text-blue-400" />
            <div>
              <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
                Gestion des RPA
              </h1>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                Résidences pour Personnes Âgées du Québec
              </p>
            </div>
          </div>

          <div class="flex items-center gap-3">
            <!-- Stats rapides -->
            <div class="flex items-center gap-4 mr-4">
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
    </div>

    <!-- Contenu principal -->
    <div class="max-w-7xl mx-auto p-6">
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
          <RPAList
            :rpas="rpas"
            :loading="loading"
            @select="handleSelectRPA"
            @delete="handleDeleteRPA"
          />
        </div>

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
import { Building2, RefreshCw, Plus } from 'lucide-vue-next'
import RPAList from '../../components/RPA/RPAList.vue'
import RPADetails from '../../components/RPA/RPADetails.vue'
import RPAFilters from '../../components/RPA/RPAFilters.vue'
import RPASyncPanel from '../../components/RPA/RPASyncPanel.vue'
import { SearchResidences, DeleteResidence } from '../../../wailsjs/go/main/App' // ⬅️ AJOUTÉ

const rpas = ref([])
const loading = ref(false)
const selectedRPA = ref(null)
const detailsMode = ref('view') // 'view', 'edit', 'create'
const showSyncPanel = ref(false)
const showNewRPA = ref(false)
const allRpas = ref([])

const filters = ref({
  municipalite: '',
  nom: '',
  region: '',
  registre: '',      // ⬅️ AJOUTÉ
  statut: ''
})

// On s'assure que rpas.value n'est jamais null pour ne pas faire planter .length
const stats = computed(() => {
  const data = rpas.value || [] // Si c'est null, on prend un tableau vide []
  return {
    total: data.length,
    actifs: data.filter(r => r.statut === 'actif').length,
    fermes: data.filter(r => r.statut === 'ferme').length,
    suspendus: data.filter(r => r.statut === 'suspendu').length
  }
})

// Rechercher les RPA
// Fonction de filtrage local
const handleSearch = () => {
  console.log('🔍 Filtrage avec:', filters.value)
  
  rpas.value = allRpas.value.filter(rpa => {
    // Filtre nom
    if (filters.value.nom && !rpa.titre.toLowerCase().includes(filters.value.nom.toLowerCase())) {
      return false
    }
    
    // Filtre municipalité (min 4 caractères)
    if (filters.value.municipalite && filters.value.municipalite.length >= 4) {
      if (!rpa.municipalite.toLowerCase().includes(filters.value.municipalite.toLowerCase())) {
        return false
      }
    }
    
    // Filtre région
    if (filters.value.region && rpa.region !== filters.value.region) {
      return false
    }
    
    // Filtre registre
    if (filters.value.registre && !rpa.registre.includes(filters.value.registre)) {
      return false
    }
    
    // Filtre statut
    if (filters.value.statut && rpa.statut !== filters.value.statut) {
      return false
    }
    
    return true
  })
  
  console.log('📊 Résultats filtrés:', rpas.value.length)
}


// Réinitialiser les filtres
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
// Sélectionner un RPA
const handleSelectRPA = (rpa) => {
  selectedRPA.value = rpa
  detailsMode.value = 'view'
}

// Sauvegarder les modifications
const handleSaveRPA = async (updatedRPA) => {
  try {
    await UpdateRPA(updatedRPA)
    selectedRPA.value = null
    await handleSearch() // Recharger la liste
    alert('✅ RPA mis à jour!')
  } catch (err) {
    console.error('Erreur sauvegarde RPA:', err)
    alert('Erreur lors de la sauvegarde')
  }
}

// Créer un nouveau RPA
const handleCreateRPA = async (newRPA) => {
  try {
    await CreateRPA(newRPA)
    showNewRPA.value = false
    await handleSearch()
    alert('✅ RPA créé!')
  } catch (err) {
    console.error('Erreur création RPA:', err)
    alert('Erreur lors de la création')
  }
}

// Supprimer un RPA
const handleDeleteRPA = async (rpa) => {
  if (!confirm(`Supprimer "${rpa.titre_rpa}" ?`)) return
  
  try {
    await DeleteRPA(rpa.id)
    selectedRPA.value = null
    await handleSearch()
    alert('✅ RPA supprimé!')
  } catch (err) {
    console.error('Erreur suppression RPA:', err)
    alert('Erreur lors de la suppression')
  }
}

// Après synchronisation
const handleSynced = async () => {
  showSyncPanel.value = false
  await handleSearch()
}

// Fonction de chargement initial
const loadAllRPAs = async () => {
  loading.value = true
  try {
    // Charge TOUT sans filtre
    allRpas.value = await SearchResidences('', '', '')
    rpas.value = [...allRpas.value]
    console.log('✅ RPA chargés:', allRpas.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement RPA:', err)
  } finally {
    loading.value = false
  }
}

// Charger au démarrage
onMounted(() => {
  loadAllRPAs() 
})
</script>