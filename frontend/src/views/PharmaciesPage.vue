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

            <!-- Bouton Import CSV -->
            <button
              @click="handleImportCSV"
              :disabled="importLoading"
              class="flex items-center gap-2 px-4 py-2 bg-teal-600 hover:bg-teal-700 disabled:opacity-50 disabled:cursor-not-allowed text-white rounded-lg transition-colors text-sm font-medium"
              title="Importer des pharmacies depuis un fichier CSV"
            >
              <Loader2 v-if="importLoading" :size="16" class="animate-spin" />
              <Upload v-else :size="16" />
              {{ importLoading ? 'Import...' : 'Importer CSV' }}
            </button>

            <!-- Bouton Ajouter -->
            <button
              @click="handleNewPharmacie"
              class="flex items-center gap-2 px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg transition-colors text-sm font-medium"
            >
              <Plus :size="18" />
              Ajouter
            </button>
          </div>
        </div>

        <!-- Notification import -->
        <transition name="fade">
          <div
            v-if="importNotif"
            :class="[
              'mt-3 px-4 py-2.5 rounded-lg text-sm font-medium flex items-center gap-2',
              importNotif.type === 'success'
                ? 'bg-emerald-50 dark:bg-emerald-900/20 text-emerald-700 dark:text-emerald-400 border border-emerald-200 dark:border-emerald-800'
                : 'bg-rose-50 dark:bg-rose-900/20 text-rose-700 dark:text-rose-400 border border-rose-200 dark:border-rose-800'
            ]"
          >
            <CheckCircle v-if="importNotif.type === 'success'" :size="16" />
            <AlertCircle v-else :size="16" />
            {{ importNotif.message }}
          </div>
        </transition>
      </div>
    </div>

    <!-- Contenu principal -->
    <div class="max-w-7xl mx-auto p-6">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">

        <!-- Sidebar Filtres -->
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

    <!-- Modal détails / édition -->
    <PharmacieDetailsModal
      v-if="selectedPharmacie"
      :pharmacie="selectedPharmacie"
      @close="selectedPharmacie = null"
      @save="handleSavePharmacie"
      @delete="handleDeleteFromModal"
      @view-client="handleViewClient"
    />

    <!-- Modal nouvelle pharmacie -->
    <PharmacieDetailsModal
      v-if="showNewPharmacie"
      :pharmacie="newPharmacieTemplate"
      mode="create"
      @close="showNewPharmacie = false"
      @save="handleCreatePharmacie"
    />

  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { Pill, Plus, Upload, Loader2, CheckCircle, AlertCircle } from 'lucide-vue-next'
import PharmacieList from '../components/Pharmacies/PharmaciesList.vue'
import PharmacieDetailsModal from '../components/Pharmacies/Pharmaciesdetailsmodal.vue'
import PharmacieFilters from '../components/Pharmacies/Pharmaciesfilters.vue'
import {
  GetAllPharmacies,
  SavePharmacie,
  DeletePharmacie,
  ImportPharmaciesCSV
} from '../../wailsjs/go/main/App'

// ─── État ────────────────────────────────────────────────────────────────────
const pharmacies    = ref([])
const allPharmacies = ref([])
const loading       = ref(false)
const selectedPharmacie = ref(null)
const showNewPharmacie  = ref(false)
const importLoading = ref(false)
const importNotif   = ref(null)

// Template vide pour "nouvelle pharmacie" — évite le crash pharmacie=null
const newPharmacieTemplate = {
  ID: 0,
  NomOrganisation: '',
  Banniere: '',
  Adresse: '',
  Ville: '',
  Region: '',
  Tel: '',
  Fax: '',
  DimancheOuvert: 0,  HeureOuvertureDimanche: '', HeureFermetureDimanche: '',
  LundiOuvert: 0,     HeureOuvertureLundi: '',    HeureFermetureLundi: '',
  MardiOuvert: 0,     HeureOuvertureMardi: '',    HeureFermetureMardi: '',
  MercrediOuvert: 0,  HeureOuvertureMercredi: '', HeureFermetureMercredi: '',
  JeudiOuvert: 0,     HeureOuvertureJeudi: '',    HeureFermetureJeudi: '',
  VendrediOuvert: 0,  HeureOuvertureVendredi: '', HeureFermetureVendredi: '',
  SamediOuvert: 0,    HeureOuvertureSamedi: '',   HeureFermetureSamedi: '',
  DateMaj: '',
  note: ''
}

const filters = ref({
  nom: '',
  ville: '',
  region: '',
  banniere: '',
  ouvertDimanche: false
})

// ─── Stats ───────────────────────────────────────────────────────────────────
const stats = computed(() => {
  const data = pharmacies.value || []
  const aujourd_hui = new Date().getDay()
  const joursKeys = ['DimancheOuvert','LundiOuvert','MardiOuvert','MercrediOuvert','JeudiOuvert','VendrediOuvert','SamediOuvert']
  const key = joursKeys[aujourd_hui]

  return {
    total: data.length,
    ouvertes: data.filter(p => p[key] && p[key] !== 0).length
  }
})

// ─── Chargement ──────────────────────────────────────────────────────────────
const loadAllPharmacies = async () => {
  loading.value = true
  try {
    const result = await GetAllPharmacies()
    allPharmacies.value = result || []
    pharmacies.value    = result || []
    console.log('✅ Pharmacies chargées:', pharmacies.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement pharmacies:', err)
  } finally {
    loading.value = false
  }
}

// ─── Filtrage ─────────────────────────────────────────────────────────────────
const handleSearch = () => {
  pharmacies.value = allPharmacies.value.filter(pharma => {
    if (filters.value.nom?.trim()) {
      const term = filters.value.nom.toLowerCase().trim()
      if (!(pharma.NomOrganisation || '').toLowerCase().includes(term)) return false
    }
    if (filters.value.ville?.trim().length >= 3) {
      const term = filters.value.ville.toLowerCase().trim()
      if (!(pharma.Ville || '').toLowerCase().includes(term)) return false
    }
    if (filters.value.region?.trim()) {
      const fr = filters.value.region.trim().toLowerCase()
      const pr = (pharma.Region || '').trim().toLowerCase()
      if (!pr.includes(fr) && !fr.includes(pr)) return false
    }
    if (filters.value.banniere?.trim()) {
      const fb = filters.value.banniere.toLowerCase().trim()
      const pb = (pharma.Banniere || '').toLowerCase().trim()
      if (pb !== fb) return false
    }
    if (filters.value.ouvertDimanche) {
      if (!pharma.DimancheOuvert || pharma.DimancheOuvert === 0) return false
    }
    return true
  })
}

const handleResetFilters = () => {
  filters.value = { nom: '', ville: '', region: '', banniere: '', ouvertDimanche: false }
  pharmacies.value = [...allPharmacies.value]
}

// ─── Import CSV ───────────────────────────────────────────────────────────────
const handleImportCSV = async () => {
  importLoading.value = true
  importNotif.value   = null
  try {
    const result = await ImportPharmaciesCSV()
    if (!result) {
      // L'utilisateur a annulé le dialog
      return
    }
    importNotif.value = { type: 'success', message: `✅ ${result.message}` }
    await loadAllPharmacies()

    // Effacer la notif après 5 secondes
    setTimeout(() => { importNotif.value = null }, 5000)
  } catch (err) {
    console.error('❌ Erreur import CSV:', err)
    importNotif.value = { type: 'error', message: `Erreur lors de l'import : ${err}` }
  } finally {
    importLoading.value = false
  }
}

// ─── CRUD ─────────────────────────────────────────────────────────────────────
const handleSelectPharmacie = (pharmacie) => {
  selectedPharmacie.value = pharmacie
}

const handleNewPharmacie = () => {
  showNewPharmacie.value = true
}

const handleSavePharmacie = async (updatedPharmacie) => {
  try {
    await SavePharmacie(updatedPharmacie)
    selectedPharmacie.value = null
    await loadAllPharmacies()
  } catch (err) {
    console.error('❌ Erreur sauvegarde pharmacie:', err)
  }
}

const handleCreatePharmacie = async (newPharmacie) => {
  try {
    await SavePharmacie(newPharmacie)
    showNewPharmacie.value = false
    await loadAllPharmacies()
  } catch (err) {
    console.error('❌ Erreur création pharmacie:', err)
  }
}

// Suppression depuis la liste (card)
const handleDeletePharmacie = async (pharmacie) => {
  if (!confirm(`Supprimer "${pharmacie.NomOrganisation}" ?\n\nCette action est irréversible.`)) return
  try {
    await DeletePharmacie(pharmacie.ID)
    selectedPharmacie.value = null
    await loadAllPharmacies()
  } catch (err) {
    console.error('❌ Erreur suppression:', err)
  }
}

// Suppression depuis le modal (le modal a déjà son propre confirm)
const handleDeleteFromModal = async (pharmacie) => {
  try {
    await DeletePharmacie(pharmacie.ID)
    selectedPharmacie.value = null
    await loadAllPharmacies()
  } catch (err) {
    console.error('❌ Erreur suppression:', err)
  }
}

const handleViewClient = (clientId) => {
  console.log('Voir client:', clientId)
  // router.push({ name: 'client', params: { id: clientId } })
}

// ─── Init ─────────────────────────────────────────────────────────────────────
onMounted(() => {
  loadAllPharmacies()
})
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>