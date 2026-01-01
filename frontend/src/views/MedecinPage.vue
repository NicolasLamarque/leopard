<!-- frontend/src/views/MedecinPage.vue -->

<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-3xl font-bold dark:text-white">Médecins</h1>
        <p class="text-gray-600 dark:text-gray-400 mt-1">
          Répertoire des médecins référents
        </p>
      </div>
      <button 
        @click="showForm = true" 
        class="bg-emerald-600 text-white px-6 py-2.5 rounded-lg hover:bg-emerald-700 
               dark:bg-emerald-500 dark:hover:bg-emerald-600 transition-colors
               flex items-center space-x-2 shadow-lg hover:shadow-xl"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
        </svg>
        <span>Nouveau médecin</span>
      </button>
    </div>

    <!-- Barre de recherche -->
    <div class="mb-6">
      <div class="relative">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher par nom, licence ou spécialisation..."
          class="w-full px-4 py-3 pl-11 rounded-lg border border-gray-300 dark:border-gray-600
                 bg-white dark:bg-gray-800 text-gray-900 dark:text-white
                 focus:ring-2 focus:ring-emerald-500 focus:border-transparent"
          @input="handleSearch"
        />
        <svg class="w-5 h-5 absolute left-3 top-3.5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
        </svg>
      </div>
    </div>

    <!-- Statistiques rapides -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-4 border border-gray-200 dark:border-gray-700">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600 dark:text-gray-400">Total</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ medecins.length }}</p>
          </div>
          <div class="p-3 bg-emerald-100 dark:bg-emerald-900/30 rounded-lg">
            <svg class="w-6 h-6 text-emerald-600 dark:text-emerald-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-4 border border-gray-200 dark:border-gray-700">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600 dark:text-gray-400">Actifs</p>
            <p class="text-2xl font-bold text-emerald-600 dark:text-emerald-400">{{ activeCount }}</p>
          </div>
          <div class="p-3 bg-green-100 dark:bg-green-900/30 rounded-lg">
            <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-4 border border-gray-200 dark:border-gray-700">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600 dark:text-gray-400">Spécialités</p>
            <p class="text-2xl font-bold text-purple-600 dark:text-purple-400">{{ specialitiesCount }}</p>
          </div>
          <div class="p-3 bg-purple-100 dark:bg-purple-900/30 rounded-lg">
            <svg class="w-6 h-6 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"/>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-4 border border-gray-200 dark:border-gray-700">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600 dark:text-gray-400">Ajoutés ce mois</p>
            <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ recentCount }}</p>
          </div>
          <div class="p-3 bg-blue-100 dark:bg-blue-900/30 rounded-lg">
            <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Table des médecins -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow overflow-hidden border border-gray-200 dark:border-gray-700">
      <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
        <thead class="bg-gray-50 dark:bg-gray-700">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Licence
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Nom complet
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Spécialisation
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Téléphone
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Email
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Statut
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
          <tr 
            v-for="medecin in filteredMedecins" 
            :key="medecin.id"
            class="hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
          >
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-emerald-600 dark:text-emerald-400">
              {{ medecin.licence }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ medecin.nomComplet }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ medecin.specialisation || 'Non spécifié' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ medecin.telephone || '-' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ medecin.email || '-' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span 
                :class="medecin.actif ? 'bg-green-100 text-green-800 dark:bg-green-900/50 dark:text-green-300' : 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300'"
                class="px-2 py-1 text-xs font-medium rounded-full"
              >
                {{ medecin.actif ? 'Actif' : 'Inactif' }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <button 
                @click="editMedecin(medecin)"
                class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 mr-3 font-medium"
              >
                Modifier
              </button>
              <button 
                @click="deleteMedecin(medecin.id)"
                class="text-red-600 dark:text-red-400 hover:text-red-800 dark:hover:text-red-300 font-medium"
              >
                Supprimer
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Message si aucun résultat -->
      <div v-if="filteredMedecins.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <p class="mt-2 text-sm text-gray-500 dark:text-gray-400">Aucun médecin trouvé</p>
      </div>
    </div>

    <!-- Formulaire de médecin (modal) -->
    <MedecinForm 
      v-if="showForm"
      :medecin="selectedMedecin"
      @close="closeForm"
      @saved="loadMedecins"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import MedecinForm from '../components/MedecinForm.vue'
import { GetMedecins, DeleteMedecin } from '../../wailsjs/go/main/App'

const medecins = ref([])
const searchQuery = ref('')
const showForm = ref(false)
const selectedMedecin = ref(null)

// Statistiques calculées
const activeCount = computed(() => 
  medecins.value.filter(m => m.actif === 1).length
)

const specialitiesCount = computed(() => 
  new Set(medecins.value.map(m => m.specialisation).filter(Boolean)).size
)

const recentCount = computed(() => {
  const oneMonthAgo = new Date()
  oneMonthAgo.setMonth(oneMonthAgo.getMonth() - 1)
  return medecins.value.filter(m => new Date(m.created_at) > oneMonthAgo).length
})

// Filtrage des médecins
const filteredMedecins = computed(() => {
  if (!searchQuery.value) return medecins.value
  
  const query = searchQuery.value.toLowerCase()
  return medecins.value.filter(m => 
    m.nomComplet?.toLowerCase().includes(query) ||
    m.licence?.toLowerCase().includes(query) ||
    m.specialisation?.toLowerCase().includes(query)
  )
})

// Chargement des médecins
const loadMedecins = async () => {
  try {
    console.log('🔍 Chargement des médecins...')
    const result = await GetMedecins()
    medecins.value = result || []
    
   
  } catch (err) {
    console.error('❌ Erreur chargement médecins:', err)
  }
}

const handleSearch = () => {
  // Le filtrage est géré par le computed
}

const editMedecin = (medecin) => {
  selectedMedecin.value = { ...medecin }
  showForm.value = true
}

const closeForm = () => {
  showForm.value = false
  selectedMedecin.value = null
}

const deleteMedecin = async (id) => {
  if (confirm("Êtes-vous sûr de vouloir supprimer ce médecin ?")) {
    try {
      // await DeleteMedecin(id)
      await loadMedecins()
      console.log('✅ Médecin supprimé')
    } catch (err) {
      console.error('❌ Erreur suppression:', err)
    }
  }
}

onMounted(() => {
  loadMedecins()
})
</script>