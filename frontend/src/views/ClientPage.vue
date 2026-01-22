<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-950 p-6">
    <div class="max-w-7xl mx-auto">
      
      <!-- Header avec stats -->
      <div class="mb-8">
        <div class="flex items-center justify-between mb-6">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">Dossiers Clients</h1>
            <p class="text-gray-600 dark:text-gray-400">Une gestion complète de vos dossiers clients</p>
          </div>
          <button 
            @click="showForm = true"
            class="flex items-center gap-2 bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white px-6 py-3 rounded-xl font-semibold shadow-lg hover:shadow-xl transition-all"
          >
            <UserPlus :size="20" />
            <span>Nouveau dossier</span>
          </button>
        </div>

        <!-- Statistiques rapides -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
          <div class="bg-white dark:bg-gray-900 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-800">
            <div class="flex items-center gap-3">
              <div class="p-3 bg-blue-50 dark:bg-blue-900/30 rounded-lg">
                <Users :size="24" class="text-blue-600 dark:text-blue-400" />
              </div>
              <div>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.total }}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400">Total clients</p>
              </div>
            </div>
          </div>

          <div class="bg-white dark:bg-gray-900 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-800">
            <div class="flex items-center gap-3">
              <div class="p-3 bg-green-50 dark:bg-green-900/30 rounded-lg">
                <UserCheck :size="24" class="text-teal-600 dark:text-teal-400" />
              </div>
              <div>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.actifs }}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400">Actifs</p>
              </div>
            </div>
          </div>

          <div class="bg-white dark:bg-gray-900 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-800">
            <div class="flex items-center gap-3">
              <div class="p-3 bg-purple-50 dark:bg-purple-900/30 rounded-lg">
                <Calendar :size="24" class="text-purple-600 dark:text-purple-400" />
              </div>
              <div>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.nouveaux }}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400">Ce mois</p>
              </div>
            </div>
          </div>

          <div class="bg-white dark:bg-gray-900 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-800">
            <div class="flex items-center gap-3">
              <div class="p-3 bg-orange-50 dark:bg-orange-900/30 rounded-lg">
                <Archive :size="24" class="text-orange-600 dark:text-orange-400" />
              </div>
              <div>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">{{ stats.archives }}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400">Archivés</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Barre de recherche et filtres -->
        <div class="bg-white dark:bg-gray-900 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-800">
          <div class="flex flex-col lg:flex-row gap-4">
            <!-- Recherche -->
            <div class="flex-1 relative">
              <Search :size="20" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input 
                v-model="searchQuery"
                type="text"
                placeholder="Rechercher par nom, prénom, N° dossier..."
                class="w-full pl-10 pr-4 py-2.5 border-2 border-gray-200 dark:border-gray-700 rounded-lg bg-gray-50 dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
              />
            </div>

            <!-- Filtres -->
            <div class="flex gap-2">
              <select 
                v-model="filterStatut"
                class="px-4 py-2.5 border-2 border-gray-200 dark:border-gray-700 rounded-lg bg-gray-50 dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 outline-none"
              >
                <option value="tous">Tous les statuts</option>
                <option value="actif">Actifs</option>
                <option value="inactif">Inactifs</option>
                <option value="dcd">Décédés</option>
              </select>

              <button 
                @click="toggleView"
                class="px-4 py-2.5 border-2 border-gray-200 dark:border-gray-700 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
              >
                <component :is="viewMode === 'table' ? Grid3x3 : List" :size="20" class="text-gray-700 dark:text-gray-300" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Vue Tableau -->
      <div v-if="viewMode === 'table'" class="bg-white dark:bg-gray-900 rounded-xl shadow-sm border border-gray-200 dark:border-gray-800 overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-800">
            <thead class="bg-gray-50 dark:bg-gray-800/50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Client
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  N° Dossier
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Coordonnées
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Statut
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white dark:bg-gray-900 divide-y divide-gray-200 dark:divide-gray-800">
              <tr 
                v-for="client in filteredClients" 
                :key="client.id"
                class="hover:bg-gray-50 dark:hover:bg-gray-800/50 transition-colors cursor-pointer"
                @click="viewClient(client.id)"
              >
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-full bg-gradient-to-br from-blue-500 to-purple-500 flex items-center justify-center text-white font-semibold">
                      {{ getInitials(client) }}
                    </div>
                    <div>
                      <p class="text-sm font-medium text-gray-900 dark:text-white">
                        {{ client.prenom }} {{ client.nom }}
                      </p>
                      <p class="text-xs text-gray-500 dark:text-gray-400">
                        {{ client.date_naissance || 'Date non renseignée' }}
                      </p>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="font-mono text-sm text-gray-900 dark:text-white">
                    {{ client.no_dossier_leopard || 'N/A' }}
                  </span>
                </td>
                <td class="px-6 py-4">
                  <div class="text-sm">
                    <p v-if="client.telephone" class="text-gray-900 dark:text-white flex items-center gap-1">
                      <Phone :size="14" class="text-gray-400" />
                      {{ client.telephone }}
                    </p>
                    <p v-if="client.email" class="text-gray-500 dark:text-gray-400 flex items-center gap-1 mt-1">
                      <Mail :size="14" class="text-gray-400" />
                      {{ client.email }}
                    </p>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span 
                    :class="getStatusClass(client)"
                    class="px-2.5 py-1 rounded-full text-xs font-medium"
                  >
                    {{ getStatusText(client) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm">
                  <div class="flex items-center gap-2">
                    <button 
                      @click.stop="viewClient(client.id)"
                      class="p-2 hover:bg-blue-50 dark:hover:bg-blue-900/30 rounded-lg transition-colors group"
                      title="Voir le dossier"
                    >
                      <Eye :size="18" class="text-blue-600 dark:text-blue-400" />
                    </button>
                    <button 
                      @click.stop="deleteClient(client.id)"
                      class="p-2 hover:bg-red-50 dark:hover:bg-red-900/30 rounded-lg transition-colors group"
                      title="Supprimer"
                    >
                      <Trash2 :size="18" class="text-red-600 dark:text-red-400" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Vue Cartes -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div 
          v-for="client in filteredClients" 
          :key="client.id"
          @click="viewClient(client.id)"
          class="bg-white dark:bg-gray-900 rounded-xl p-5 shadow-sm border border-gray-200 dark:border-gray-800 hover:shadow-lg transition-all cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="flex items-center gap-3">
              <div class="w-12 h-12 rounded-full bg-gradient-to-br from-blue-500 to-purple-500 flex items-center justify-center text-white font-bold text-lg">
                {{ getInitials(client) }}
              </div>
              <div>
                <h3 class="font-semibold text-gray-900 dark:text-white group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors">
                  {{ client.prenom }} {{ client.nom }}
                </h3>
                <p class="text-xs text-gray-500 dark:text-gray-400 font-mono">
                  {{ client.no_dossier_leopard || 'N/A' }}
                </p>
              </div>
            </div>
            <span 
              :class="getStatusClass(client)"
              class="px-2 py-1 rounded-full text-xs font-medium"
            >
              {{ getStatusText(client) }}
            </span>
          </div>

          <div class="space-y-2 text-sm">
            <p v-if="client.telephone" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <Phone :size="14" />
              {{ client.telephone }}
            </p>
            <p v-if="client.email" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <Mail :size="14" />
              {{ client.email }}
            </p>
            <p v-if="client.ville" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <MapPin :size="14" />
              {{ client.ville }}
            </p>
          </div>
        </div>
      </div>

      <!-- Message si aucun résultat -->
      <div v-if="filteredClients.length === 0" class="bg-white dark:bg-gray-900 rounded-xl p-12 text-center shadow-sm border border-gray-200 dark:border-gray-800">
        <UserX :size="48" class="mx-auto text-gray-400 mb-4" />
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">Aucun client trouvé</h3>
        <p class="text-gray-600 dark:text-gray-400 mb-4">Essayez de modifier vos critères de recherche</p>
        <button 
          @click="resetFilters"
          class="text-blue-600 dark:text-blue-400 hover:underline"
        >
          Réinitialiser les filtres
        </button>
      </div>

    </div>

    <!-- Modal création client -->
    <ClientForm 
      v-if="showForm" 
      @close="showForm = false" 
      @created="loadClients" 
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { GetClients, DeleteClient } from '../../wailsjs/go/main/App'
import ClientForm from '../components/clients/ClientForm.vue'
import { useClientStore } from '../stores/clientStore'
import { 
  Users, UserPlus, UserCheck, Calendar, Archive, Search, 
  Grid3x3, List, Eye, Trash2, Phone, Mail, MapPin, UserX
} from 'lucide-vue-next'

const router = useRouter()
const store = useClientStore()

const clients = ref([])
const showForm = ref(false)
const searchQuery = ref('')
const filterStatut = ref('tous')
const viewMode = ref('table') // 'table' ou 'cards'

const stats = computed(() => {
  const total = clients.value.length
  const actifs = clients.value.filter(c => c.actif === 1 && c.dcd === 0).length
  const archives = clients.value.filter(c => c.actif === 0).length
  
  // Clients créés ce mois
  const now = new Date()
  const currentMonth = now.getMonth()
  const currentYear = now.getFullYear()
  const nouveaux = clients.value.filter(c => {
    const createdAt = new Date(c.created_at)
    return createdAt.getMonth() === currentMonth && createdAt.getFullYear() === currentYear
  }).length

  return { total, actifs, archives, nouveaux }
})

const filteredClients = computed(() => {
  let filtered = clients.value

  // Filtre par recherche
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(c => 
      c.nom.toLowerCase().includes(query) ||
      c.prenom.toLowerCase().includes(query) ||
      (c.no_dossier_leopard && c.no_dossier_leopard.toLowerCase().includes(query))
    )
  }

  // Filtre par statut
  if (filterStatut.value !== 'tous') {
    if (filterStatut.value === 'actif') {
      filtered = filtered.filter(c => c.actif === 1 && c.dcd === 0)
    } else if (filterStatut.value === 'inactif') {
      filtered = filtered.filter(c => c.actif === 0)
    } else if (filterStatut.value === 'dcd') {
      filtered = filtered.filter(c => c.dcd === 1)
    }
  }

  return filtered
})

const getInitials = (client) => {
  return `${client.prenom[0]}${client.nom[0]}`.toUpperCase()
}

const getStatusClass = (client) => {
  if (client.dcd === 1) return 'bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300'
  if (client.actif === 1) return 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300'
  return 'bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300'
}

const getStatusText = (client) => {
  if (client.dcd === 1) return 'Décédé'
  if (client.actif === 1) return 'Actif'
  return 'Archivé'
}

const toggleView = () => {
  viewMode.value = viewMode.value === 'table' ? 'cards' : 'table'
}

const resetFilters = () => {
  searchQuery.value = ''
  filterStatut.value = 'tous'
}

const loadClients = async () => {
  try {
    const result = await GetClients()
    clients.value = result || []
  } catch (err) {
    console.error('Erreur chargement clients:', err)
  }
}

const viewClient = async (id) => {
  store.clearClient()
  router.push(`/app/clients/${id}`)
}

const deleteClient = async (id) => {
  if (confirm("Êtes-vous sûr de vouloir supprimer ce dossier ? Cette action est irréversible.")) {
    try {
      await DeleteClient(id)
      await loadClients()
    } catch (err) {
      console.error('Erreur suppression:', err)
    }
  }
}

onMounted(() => {
  loadClients()
})
</script>