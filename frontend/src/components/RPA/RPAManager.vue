<template>
  <div class="space-y-6">
    <!-- En-tête avec actions -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
            Gestion des RPA
          </h2>
          <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
            Synchronisation avec le registre du MSSS
          </p>
        </div>
        
        <button
          @click="syncRPA"
          :disabled="syncing"
          class="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
        >
          <RefreshCw :size="18" :class="{ 'animate-spin': syncing }" />
          {{ syncing ? 'Synchronisation...' : 'Synchroniser' }}
        </button>
      </div>

      <!-- Statistiques -->
      <div v-if="stats" class="grid grid-cols-4 gap-4 mt-4">
        <div class="bg-blue-50 dark:bg-blue-900/20 rounded-lg p-4">
          <p class="text-sm text-blue-600 dark:text-blue-400 font-medium">Total récupéré</p>
          <p class="text-2xl font-bold text-blue-700 dark:text-blue-300">{{ stats.total_scraped }}</p>
        </div>
        <div class="bg-green-50 dark:bg-green-900/20 rounded-lg p-4">
          <p class="text-sm text-green-600 dark:text-green-400 font-medium">Nouveaux</p>
          <p class="text-2xl font-bold text-green-700 dark:text-green-300">{{ stats.nouveaux }}</p>
        </div>
        <div class="bg-yellow-50 dark:bg-yellow-900/20 rounded-lg p-4">
          <p class="text-sm text-yellow-600 dark:text-yellow-400 font-medium">Mis à jour</p>
          <p class="text-2xl font-bold text-yellow-700 dark:text-yellow-300">{{ stats.mis_a_jour }}</p>
        </div>
        <div class="bg-red-50 dark:bg-red-900/20 rounded-lg p-4">
          <p class="text-sm text-red-600 dark:text-red-400 font-medium">Fermés</p>
          <p class="text-2xl font-bold text-red-700 dark:text-red-300">{{ stats.fermes }}</p>
        </div>
      </div>

      <!-- Erreurs -->
      <div v-if="stats && stats.erreurs && stats.erreurs.length > 0" class="mt-4 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
        <h3 class="text-sm font-semibold text-red-800 dark:text-red-300 mb-2">Erreurs lors de la synchronisation:</h3>
        <ul class="text-sm text-red-700 dark:text-red-400 space-y-1">
          <li v-for="(err, i) in stats.erreurs" :key="i">• {{ err }}</li>
        </ul>
      </div>

      <p v-if="lastSync" class="text-xs text-gray-500 dark:text-gray-400 mt-4">
        Dernière synchronisation: {{ formatDate(lastSync) }}
      </p>
    </div>

    <!-- Recherche et filtres -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Rechercher un RPA</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Municipalité
          </label>
          <input
            v-model="filters.municipalite"
            @input="searchRPA"
            type="text"
            placeholder="Ex: Montréal"
            class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Nom du RPA
          </label>
          <input
            v-model="filters.nom"
            @input="searchRPA"
            type="text"
            placeholder="Ex: Résidence du Parc"
            class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Statut
          </label>
          <select
            v-model="filters.statut"
            @change="searchRPA"
            class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
          >
            <option value="">Tous</option>
            <option value="actif">Actif</option>
            <option value="ferme">Fermé</option>
            <option value="suspendu">Suspendu</option>
          </select>
        </div>
      </div>

      <button
        @click="clearFilters"
        class="mt-4 text-sm text-blue-600 dark:text-blue-400 hover:underline"
      >
        Effacer les filtres
      </button>
    </div>

    <!-- Liste des RPA -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow overflow-hidden">
      <div class="px-6 py-4 border-b dark:border-gray-700">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
          Résultats ({{ rpas.length }})
        </h3>
      </div>

      <div v-if="loading" class="p-8 text-center">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-blue-500 border-t-transparent"></div>
        <p class="mt-2 text-gray-500 dark:text-gray-400">Chargement...</p>
      </div>

      <div v-else-if="rpas.length === 0" class="p-8 text-center text-gray-500 dark:text-gray-400">
        <Building2 :size="48" class="mx-auto mb-2 opacity-50" />
        <p>Aucun RPA trouvé</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50 dark:bg-gray-900">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                Statut
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                Nom
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                Municipalité
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                Région
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                Registre
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
            <tr v-for="rpa in rpas" :key="rpa.id" class="hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors">
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    rpa.statut === 'actif' ? 'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400' :
                    rpa.statut === 'ferme' ? 'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400' :
                    'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400'
                  ]"
                >
                  {{ rpa.statut }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900 dark:text-white">
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
                <button
                  @click="viewDetails(rpa)"
                  class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300"
                >
                  <Eye :size="18" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal détails -->
    <div
      v-if="selectedRPA"
      @click.self="selectedRPA = null"
      class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4"
    >
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b dark:border-gray-700 flex items-center justify-between">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
            Détails du RPA
          </h3>
          <button
            @click="selectedRPA = null"
            class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300"
          >
            <X :size="24" />
          </button>
        </div>

        <div class="p-6 space-y-4">
          <div>
            <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Nom</label>
            <p class="text-gray-900 dark:text-white font-medium">{{ selectedRPA.titre }}</p>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Municipalité</label>
              <p class="text-gray-900 dark:text-white">{{ selectedRPA.municipalite }}</p>
            </div>
            <div>
              <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Région</label>
              <p class="text-gray-900 dark:text-white">{{ selectedRPA.region }}</p>
            </div>
          </div>

          <div>
            <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Adresse</label>
            <p class="text-gray-900 dark:text-white">{{ selectedRPA.adresse || 'Non disponible' }}</p>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Registre</label>
              <p class="text-gray-900 dark:text-white font-mono">{{ selectedRPA.registre }}</p>
            </div>
            <div>
              <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Statut</label>
              <span
                :class="[
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                  selectedRPA.statut === 'actif' ? 'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400' :
                  selectedRPA.statut === 'ferme' ? 'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400' :
                  'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400'
                ]"
              >
                {{ selectedRPA.statut }}
              </span>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Date d'ajout</label>
              <p class="text-gray-900 dark:text-white text-sm">{{ formatDate(selectedRPA.date_ajout) }}</p>
            </div>
            <div>
              <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Dernière vérification</label>
              <p class="text-gray-900 dark:text-white text-sm">{{ formatDate(selectedRPA.derniere_verification) }}</p>
            </div>
          </div>

          <div v-if="selectedRPA.notes">
            <label class="text-sm font-medium text-gray-500 dark:text-gray-400">Notes</label>
            <p class="text-gray-900 dark:text-white text-sm">{{ selectedRPA.notes }}</p>
          </div>
        </div>

        <div class="p-6 border-t dark:border-gray-700 flex justify-end gap-3">
          <button
            @click="selectedRPA = null"
            class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
          >
            Fermer
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { RefreshCw, Building2, Eye, X } from 'lucide-vue-next'
import { SearchResidences, SyncRPA } from '../../../wailsjs/go/main/App'


const syncing = ref(false)
const loading = ref(false)
const stats = ref(null)
const lastSync = ref(null)
const rpas = ref([])
const selectedRPA = ref(null)

const filters = ref({
  municipalite: '',
  nom: '',
  statut: ''
})

// Synchroniser avec le MSSS
const syncRPA = async () => {
  syncing.value = true
  try {
    // Appel de la fonction exportée par Wails
    const result = await SyncRPA()
    stats.value = result
    lastSync.value = result.date_sync
    
    // Recharger la liste avec les filtres actuels
    await searchRPA()
    
    alert(`✅ Synchronisation terminée!\n\nNouveau: ${result.nouveaux}\nMis à jour: ${result.mis_a_jour}\nFermés: ${result.fermes}`)
  } catch (err) {
    console.error('Erreur sync:', err)
    alert('❌ Erreur lors de la synchronisation: ' + err)
  } finally {
    syncing.value = false
  }
}

// Rechercher des RPA
// Fonction locale qui gère le chargement et l'assignation
const searchRPA = async () => {
  loading.value = true
  try {
    // On appelle Go ET on assigne le résultat à la ref 'rpas'
    const data = await SearchResidences(
      filters.value.municipalite,
      filters.value.nom,
      filters.value.statut
    )
    rpas.value = data || [] // Force un tableau vide si data est null
  } catch (err) {
    console.error('Erreur recherche:', err)
    rpas.value = [] 
  } finally {
    loading.value = false
  }
}


// Effacer les filtres
const clearFilters = () => {
  filters.value = {
    municipalite: '',
    nom: '',
    statut: ''
  }
  searchRPA()
}

const viewDetails = (rpa) => {
  selectedRPA.value = rpa
}

const formatDate = (date) => {
  if (!date) return 'N/A'
  return new Date(date).toLocaleString('fr-CA', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// Charger la liste au démarrage
onMounted(() => {
  searchRPA()
})
</script>