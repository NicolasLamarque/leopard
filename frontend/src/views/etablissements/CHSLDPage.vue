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
                Gestion des CHSLD
              </h1>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                Centres d'hébergement et de soins de longue durée
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
              @click="showNewCHSLD = true"
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
          <CHSLDFilters 
            v-model:filters="filters"
            @search="handleSearch"
            @reset="handleResetFilters"
          />
        </div>

        <!-- Liste des CHSLD -->
        <div class="lg:col-span-3">
          <CHSLDList
            :chslds="chslds"
            :loading="loading"
            @select="handleSelectCHSLD"
            @delete="handleDeleteCHSLD"
          />
        </div>

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
import { Building2, Plus } from 'lucide-vue-next'
import CHSLDList from '../../components/CHSLD/CHSLDList.vue'
import CHSLDDetails from '../../components/CHSLD/CHSLDDetails.vue'
import CHSLDFilters from '../../components/CHSLD/CHSLDFilters.vue'
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

// Statistiques
const stats = computed(() => {
  const data = chslds.value || []
  return {
    total: data.length,
    actifs: data.filter(c => c.Statut?.toLowerCase().includes('actif')).length,
    fermes: data.filter(c => c.Statut?.toLowerCase().includes('fermé')).length,
    suspendus: data.filter(c => c.Statut?.toLowerCase().includes('suspendu')).length
  }
})

// Fonction de filtrage local
const handleSearch = () => {
  console.log('🔍 Filtrage CHSLD avec:', filters.value)
  
  chslds.value = allChslds.value.filter(chsld => {
    // Filtre nom
    if (filters.value.nom && !chsld.TitreCHSLD?.toLowerCase().includes(filters.value.nom.toLowerCase())) {
      return false
    }
    
    // Filtre municipalité (min 4 caractères)
    if (filters.value.municipalite && filters.value.municipalite.length >= 4) {
      if (!chsld.Municipalite?.toLowerCase().includes(filters.value.municipalite.toLowerCase())) {
        return false
      }
    }
    
    // Filtre région
    if (filters.value.region && chsld.Region !== filters.value.region) {
      return false
    }
    
    // Filtre type
    if (filters.value.type && chsld.TypeCHSLD !== filters.value.type) {
      return false
    }
    
    // Filtre statut
    if (filters.value.statut && chsld.Statut !== filters.value.statut) {
      return false
    }
    
    // Filtre capacité minimale
    if (filters.value.capaciteMin > 0 && (chsld.Capacite || 0) < filters.value.capaciteMin) {
      return false
    }
    
    return true
  })
  
  console.log('📊 Résultats filtrés:', chslds.value.length)
}

// Réinitialiser les filtres
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

// Sélectionner un CHSLD
const handleSelectCHSLD = (chsld) => {
  selectedCHSLD.value = chsld
  detailsMode.value = 'view'
}

// Sauvegarder les modifications
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

// Créer un nouveau CHSLD
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

// Supprimer un CHSLD
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

// Fonction de chargement initial
const loadAllCHSLDs = async () => {
  loading.value = true
  try {
    // Charge TOUT sans filtre
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

// Charger au démarrage
onMounted(() => {
  loadAllCHSLDs()
})
</script>