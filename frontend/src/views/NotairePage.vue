<!-- frontend/src/views/NotairePage.vue -->
<template>
  <div class="min-h-screen bg-slate-50 dark:bg-slate-950 p-6 lg:p-10 transition-colors duration-300">
    
    <!-- Header -->
    <div class="max-w-7xl mx-auto mb-10 flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div class="space-y-1">
        <div class="flex items-center gap-3">
          <div class="p-2.5 bg-slate-700 dark:bg-slate-600 rounded-lg text-white shadow-lg">
            <Scale class="w-7 h-7" />
          </div>
          <div>
            <h1 class="text-3xl font-bold text-slate-900 dark:text-white tracking-tight">
              Répertoire des Notaires
            </h1>
            <p class="text-xs text-slate-500 dark:text-slate-400 font-bold uppercase tracking-[0.2em]">
              Registre professionnel
            </p>
          </div>
        </div>
      </div>
      
      <button 
        @click="openCreateForm" 
        class="flex items-center gap-2 bg-slate-900 dark:bg-slate-700 hover:bg-slate-800 dark:hover:bg-slate-600 text-white px-6 py-3 rounded-xl font-bold text-sm transition-all shadow-sm active:scale-95"
      >
        <UserPlus class="w-4 h-4" />
        <span>Nouveau Notaire</span>
      </button>
    </div>

    <!-- Barre de recherche -->
    <div class="max-w-7xl mx-auto mb-8">
      <div class="relative group">
        <div class="absolute inset-y-0 left-0 pl-5 flex items-center pointer-events-none">
          <Search class="w-5 h-5 text-slate-400 group-focus-within:text-slate-600 transition-colors" />
        </div>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher par nom, ville, secteur..."
          class="w-full pl-14 pr-32 py-4 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 rounded-2xl border border-slate-200 dark:border-slate-800 focus:border-slate-500 dark:focus:border-slate-500 focus:ring-4 focus:ring-slate-500/10 outline-none font-medium transition-all shadow-sm"
          @input="handleSearch"
        />
        <div v-if="searchQuery" class="absolute right-4 top-1/2 -translate-y-1/2 px-3 py-1 bg-slate-100 dark:bg-slate-800 text-slate-500 dark:text-slate-400 rounded-lg text-[10px] font-bold uppercase tracking-wider border border-slate-200 dark:border-slate-700">
          {{ filteredNotaires.length }} résultat{{ filteredNotaires.length > 1 ? 's' : '' }}
        </div>
      </div>
    </div>

    <!-- Stats -->
    <div class="max-w-7xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
      <div v-for="(stat, index) in [
        { label: 'Total Notaires', val: notaires.length, icon: Users, color: 'text-slate-600' },
        { label: 'Villes', val: citiesCount, icon: MapPin, color: 'text-slate-600' },
        { label: 'Secteurs', val: secteursCount, icon: Award, color: 'text-slate-600' },
        { label: 'Nouveaux (Mois)', val: recentCount, icon: TrendingUp, color: 'text-slate-600' }
      ]" :key="index" class="bg-white dark:bg-slate-900 rounded-2xl p-5 border border-slate-200 dark:border-slate-800 shadow-sm">
        <div class="flex items-center gap-3">
          <component :is="stat.icon" class="w-4 h-4" :class="stat.color" />
          <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">{{ stat.label }}</p>
        </div>
        <p class="text-2xl font-bold text-slate-900 dark:text-white mt-2">{{ stat.val }}</p>
      </div>
    </div>

    <!-- Table -->
    <div class="max-w-7xl mx-auto bg-white dark:bg-slate-900 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full border-separate border-spacing-0">
          <thead>
            <tr class="bg-slate-50 dark:bg-slate-800/50">
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Notaire</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Secteur</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Localisation</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Contact</th>
              <th class="px-6 py-4 text-center text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr 
              v-for="notaire in filteredNotaires" 
              :key="notaire.id"
              class="group hover:bg-slate-50/80 dark:hover:bg-slate-800/40 transition-colors"
            >
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-10 h-10 rounded-lg bg-slate-100 dark:bg-slate-800 flex items-center justify-center font-semibold text-sm text-slate-600 dark:text-slate-400">
                    {{ notaire.nom.charAt(0) }}{{ notaire.prenom.charAt(0) }}
                  </div>
                  <div>
                    <div class="font-bold text-slate-900 dark:text-slate-200">{{ notaire.prenom }} {{ notaire.nom }}</div>
                    <div class="text-[10px] text-slate-500 uppercase">{{ notaire.civilite || 'Me' }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div v-if="notaire.secteur_activite" class="flex items-center gap-2 text-sm text-slate-600 dark:text-slate-400">
                  <Award class="w-3.5 h-3.5 text-slate-400" />
                  {{ notaire.secteur_activite }}
                </div>
                <span v-else class="text-sm text-slate-400">—</span>
              </td>
              <td class="px-6 py-4">
                <div v-if="notaire.ville" class="flex items-center gap-2 text-sm text-slate-600 dark:text-slate-400">
                  <MapPin class="w-3.5 h-3.5 text-slate-400" />
                  {{ notaire.ville }}
                </div>
                <span v-else class="text-sm text-slate-400">—</span>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div v-if="notaire.telephone" class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400">
                    <Phone class="w-3 h-3" /> {{ notaire.telephone }}
                  </div>
                  <div v-if="notaire.email" class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400">
                    <Mail class="w-3 h-3" /> {{ notaire.email }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center justify-center gap-2">
                  <button @click="viewDetails(notaire)" class="p-2 text-slate-400 hover:text-slate-700 dark:hover:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg transition-all" title="Détails">
                    <Eye class="w-4 h-4" />
                  </button>
                  <button @click="editNotaire(notaire)" class="p-2 text-slate-400 hover:text-amber-600 dark:hover:text-amber-400 hover:bg-amber-50 dark:hover:bg-amber-500/10 rounded-lg transition-all" title="Modifier">
                    <Edit class="w-4 h-4" />
                  </button>
                  <button @click="deleteNotaire(notaire.id)" class="p-2 text-slate-400 hover:text-red-600 dark:hover:text-red-400 hover:bg-red-50 dark:hover:bg-red-500/10 rounded-lg transition-all" title="Supprimer">
                    <Trash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="filteredNotaires.length === 0" class="text-center py-20 bg-slate-50/50 dark:bg-slate-900/50">
        <UserX class="w-12 h-12 text-slate-300 dark:text-slate-700 mx-auto mb-4" />
        <p class="text-slate-500 dark:text-slate-400 font-medium">Aucun notaire trouvé</p>
      </div>
    </div>

    <!-- Modals -->
    <NotaireDetailsModal 
      v-if="showDetails" 
      :notaire="selectedNotaire" 
      @close="closeDetails" 
      @edit="editFromDetails" 
      @delete="deleteFromDetails"
      @view-client="handleViewClient"
    />
    
    <NotaireAddModal 
      v-if="showAddModal"
      :is-open="showAddModal"
      :notaire="selectedNotaire"
      @close="closeAddModal" 
      @saved="handleSaved" 
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  Scale, Users, Award, TrendingUp, Search, MapPin,
  Eye, Edit, Trash2, Phone, Mail, UserPlus, UserX
} from 'lucide-vue-next'
import NotaireDetailsModal from '../components/Notaires/NotaireDetailsModal.vue'
import NotaireAddModal from '../components/Notaires/NotaireAddModal.vue'

const notaires = ref([])
const searchQuery = ref('')
const showAddModal = ref(false)
const showDetails = ref(false)
const selectedNotaire = ref(null)


// Stats
const citiesCount = computed(() => 
  new Set(notaires.value.map(n => n.ville).filter(Boolean)).size
)

const secteursCount = computed(() => 
  new Set(notaires.value.map(n => n.secteur_activite).filter(Boolean)).size
)

const recentCount = computed(() => {
  const oneMonthAgo = new Date()
  oneMonthAgo.setMonth(oneMonthAgo.getMonth() - 1)
  return notaires.value.filter(n => new Date(n.created_at) > oneMonthAgo).length
})

// Filtrage
const filteredNotaires = computed(() => {
  if (!searchQuery.value) return notaires.value
  
  const query = searchQuery.value.toLowerCase()
  return notaires.value.filter(n => 
    `${n.prenom} ${n.nom}`.toLowerCase().includes(query) ||
    n.ville?.toLowerCase().includes(query) ||
    n.secteur_activite?.toLowerCase().includes(query) ||
    n.email?.toLowerCase().includes(query)
  )
})

const loadNotaires = async () => {
  try {
    notaires.value = await window.go.main.App.GetAllNotaires() || []
    console.log('✅ Notaires chargés:', notaires.value.length)
  } catch (err) {
    console.error('❌ Erreur:', err)
  }
}

const handleSearch = () => {
  // Filtrage géré par computed
}

const viewDetails = (notaire) => {
  selectedNotaire.value = notaire
  showDetails.value = true
}

const closeDetails = () => {
  showDetails.value = false
  selectedNotaire.value = null
}

const openCreateForm = () => {
  selectedNotaire.value = null
  showAddModal.value = true
}

const editNotaire = (notaire) => {
  selectedNotaire.value = { ...notaire }
  showAddModal.value = true
}

const editFromDetails = (notaire) => {
  closeDetails()
  editNotaire(notaire)
}

const closeAddModal = () => {
  showAddModal.value = false
  selectedNotaire.value = null
}

const handleSaved = async () => {
  await loadNotaires()
  closeAddModal()
}

const deleteNotaire = async (id) => {
  if (confirm("⚠️ Supprimer ce notaire définitivement ?")) {
    try {
      await window.go.main.App.DeleteNotaire(id)
      await loadNotaires()
      console.log('✅ Notaire supprimé')
    } catch (err) {
      console.error('❌ Erreur:', err)
    }
  }
}

const deleteFromDetails = async (id) => {
  closeDetails()
  await deleteNotaire(id)
}

const handleViewClient = (clientId) => {
  console.log('Navigation vers client:', clientId)
  // TODO: implémenter navigation
}

onMounted(() => {
  loadNotaires()
})
</script>