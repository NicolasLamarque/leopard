<!-- frontend/src/pages/PharmaciesPage.vue -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <!-- Header -->
    <div class="bg-white dark:bg-gray-800 border-b dark:border-gray-700 sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-6 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <Pill :size="32" class="text-emerald-600 dark:text-emerald-400" />
            <div>
              <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
                Gestion des Pharmacies
              </h1>
              <p class="text-sm text-gray-500 dark:text-gray-400">
                Répertoire des pharmacies du Québec
              </p>
            </div>
          </div>

          <div class="flex items-center gap-3">
            <!-- Stats rapides -->
            <div class="flex items-center gap-4 mr-4">
              <div class="text-center">
                <p class="text-2xl font-bold text-emerald-600 dark:text-emerald-400">{{ stats.total }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Total</p>
              </div>
              <div class="text-center">
                <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ stats.ouvertes }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">Ouvertes</p>
              </div>
            </div>

            <button
              @click="showNewPharmacie = true"
              class="flex items-center gap-2 px-4 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors"
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
          <PharmacieFilters 
            v-model:filters="filters"
            @search="handleSearch"
            @reset="handleResetFilters"
          />
        </div>

        <!-- Liste des pharmacies -->
        <div class="lg:col-span-3">
          <PharmacieList
            :pharmacies="pharmacies"
            :loading="loading"
            @select="handleSelectPharmacie"
            @delete="handleDeletePharmacie"
          />
        </div>

      </div>
    </div>

    <!-- Modal détails/édition pharmacie -->
    <PharmacieDetailsModal
      v-if="selectedPharmacie"
      :pharmacie="selectedPharmacie"
      @close="selectedPharmacie = null"
      @save="handleSavePharmacie"
      @delete="handleDeletePharmacie"
      @view-client="handleViewClient"
    />

    <!-- Modal nouvelle pharmacie -->
    <PharmacieDetailsModal
      v-if="showNewPharmacie"
      :pharmacie="null"
      mode="create"
      @close="showNewPharmacie = false"
      @save="handleCreatePharmacie"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { Pill, Plus } from 'lucide-vue-next'
import PharmacieList from '../components/Pharmacies/PharmaciesList.vue'
import PharmacieDetailsModal from '../components/Pharmacies/Pharmaciesdetailsmodal.vue'
import PharmacieFilters from '../components/Pharmacies/Pharmaciesfilters.vue'
import { 
  GetAllPharmacies, 
  SavePharmacie, 
  DeletePharmacie 
} from '../../wailsjs/go/main/App'

const pharmacies = ref([])
const loading = ref(false)
const selectedPharmacie = ref(null)
const showNewPharmacie = ref(false)
const allPharmacies = ref([])

const filters = ref({
  nom: '',
  ville: '',
  region: '',
  banniere: '',
  ouvertDimanche: false
})

// Stats
const stats = computed(() => {
  const data = pharmacies.value || []
  const aujourdhui = new Date().getDay() // 0 = Dimanche, 1 = Lundi, etc.
  
  let ouvertes = 0
  data.forEach(p => {
    switch(aujourdhui) {
      case 0: if (p.DimancheOuvert) ouvertes++; break;
      case 1: if (p.LundiOuvert) ouvertes++; break;
      case 2: if (p.MardiOuvert) ouvertes++; break;
      case 3: if (p.MercrediOuvert) ouvertes++; break;
      case 4: if (p.JeudiOuvert) ouvertes++; break;
      case 5: if (p.VendrediOuvert) ouvertes++; break;
      case 6: if (p.SamediOuvert) ouvertes++; break;
    }
  })

  return {
    total: data.length,
    ouvertes: ouvertes
  }
})

// Filtrage local
const handleSearch = () => {
  console.log('🔍 Filtrage pharmacies:', filters.value)
  
  pharmacies.value = allPharmacies.value.filter(pharma => {
    // Filtre nom - gère null/undefined et recherche insensible à la casse
    if (filters.value.nom && filters.value.nom.trim()) {
      const searchTerm = filters.value.nom.toLowerCase().trim()
      const pharmacieName = (pharma.NomOrganisation || '').toLowerCase()
      if (!pharmacieName.includes(searchTerm)) {
        return false
      }
    }
    
    // Filtre ville - min 3 caractères, gère null/undefined
    if (filters.value.ville && filters.value.ville.trim().length >= 3) {
      const searchTerm = filters.value.ville.toLowerCase().trim()
      const ville = (pharma.Ville || '').toLowerCase()
      if (!ville.includes(searchTerm)) {
        return false
      }
    }
    
    // Filtre région - comparaison exacte mais flexible (trim + casse)
    if (filters.value.region && filters.value.region.trim()) {
      const filterRegion = filters.value.region.trim()
      const pharmaRegion = (pharma.Region || '').trim()
      
      // Gère "Mauricie et Centre-du-Québec" vs "Mauricie" vs "Centre-du-Québec"
      const regionMatch = pharmaRegion.toLowerCase().includes(filterRegion.toLowerCase()) ||
                         filterRegion.toLowerCase().includes(pharmaRegion.toLowerCase())
      
      if (!regionMatch) {
        return false
      }
    }
    
    // Filtre bannière - gère null, espaces, et casse
    if (filters.value.banniere && filters.value.banniere.trim()) {
      const filterBanniere = filters.value.banniere.toLowerCase().trim()
      const pharmaBanniere = (pharma.Banniere || '').toLowerCase().trim()
      
      if (pharmaBanniere !== filterBanniere) {
        return false
      }
    }
    
    // Filtre ouvert dimanche - gère int (1/0) et boolean (true/false)
    if (filters.value.ouvertDimanche) {
      if (!pharma.DimancheOuvert || pharma.DimancheOuvert === 0) {
        return false
      }
    }
    
    return true
  })
  
  console.log('📊 Résultats filtrés:', pharmacies.value.length, '/', allPharmacies.value.length)
}
// Réinitialiser les filtres
const handleResetFilters = () => {
  filters.value = {
    nom: '',
    ville: '',
    region: '',
    banniere: '',
    ouvertDimanche: false
  }
  pharmacies.value = [...allPharmacies.value]
}

// Sélectionner une pharmacie
const handleSelectPharmacie = (pharmacie) => {
  selectedPharmacie.value = pharmacie
}

// Sauvegarder les modifications
// Sauvegarder les modifications
const handleSavePharmacie = async (updatedPharmacie) => {
  try {
    // Utilise la fonction SavePharmacie que tu as importée (qui gère Create et Update)
    await SavePharmacie(updatedPharmacie) 
    selectedPharmacie.value = null
    await loadAllPharmacies()
    // Optionnel: utilise une notification plus moderne que alert si tu as le temps
  } catch (err) {
    console.error('Erreur sauvegarde pharmacie:', err)
  }
}

// Créer une nouvelle pharmacie
const handleCreatePharmacie = async (newPharmacie) => {
  try {
    // Même chose ici, on utilise SavePharmacie
    await SavePharmacie(newPharmacie)
    showNewPharmacie.value = false
    await loadAllPharmacies()
  } catch (err) {
    console.error('Erreur création pharmacie:', err)
  }
}

// Supprimer une pharmacie
const handleDeletePharmacie = async (pharmacie) => {
  if (!confirm(`Supprimer "${pharmacie.NomOrganisation}" ?`)) return
  
  try {
    await DeletePharmacie(pharmacie.ID)
    selectedPharmacie.value = null
    await loadAllPharmacies()
    alert('✅ Pharmacie supprimée!')
  } catch (err) {
    console.error('Erreur suppression pharmacie:', err)
    alert('Erreur lors de la suppression')
  }
}

// Voir le dossier d'un client
const handleViewClient = (clientId) => {
  // Rediriger vers la page du client
  console.log('Voir client:', clientId)
  // router.push({ name: 'client', params: { id: clientId } })
}
const loadAllPharmacies = async () => {
  try {
    const result = await GetAllPharmacies()  // ✅ Utilise l'import
    allPharmacies.value = result || []       // ✅ Stock dans allPharmacies
    pharmacies.value = result || []          // ✅ Et aussi dans pharmacies
    console.log('✅ Pharmacies chargées:', pharmacies.value.length)
    console.log('📦 Première pharmacie:', JSON.stringify(pharmacies.value[0], null, 2))
  } catch (err) {
    console.error('❌ Erreur:', err)
  }
}

const formatPhone = (phone) => {
  if (!phone) return ''
  
  // Enlever tout sauf les chiffres
  const cleaned = phone.replace(/\D/g, '')
  
  // Si 10 chiffres : format (819) 999-0011
  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }
  
  // Si 11 chiffres commençant par 1 : format 1 (819) 999-0011
  if (cleaned.length === 11 && cleaned.startsWith('1')) {
    return `1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`
  }
  
  // Sinon retourner tel quel
  return phone
}

// Charger au démarrage
onMounted(() => {
  loadAllPharmacies()
})
</script>