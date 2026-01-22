<!-- frontend/src/views/MedecinPage.vue - VERSION ULTRA PREMIUM 🔥 -->
<template>
  <div class="min-h-screen bg-slate-50 dark:bg-slate-950 p-6 lg:p-10 transition-colors duration-300">
    
    <div class="max-w-7xl mx-auto mb-10 flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div class="space-y-1">
        <div class="flex items-center gap-3">
          <div class="p-2.5 bg-teal-600 dark:bg-teal-500 rounded-lg text-white shadow-lg shadow-teal-500/20">
            <Stethoscope class="w-7 h-7" />
          </div>
          <div>
            <h1 class="text-3xl font-bold text-slate-900 dark:text-white tracking-tight">
              Répertoire Médical
            </h1>
            <p class="text-xs text-slate-500 dark:text-slate-400 font-bold uppercase tracking-[0.2em]">
              Registre du personnel autorisé
            </p>
          </div>
        </div>
      </div>
      
      <button 
        @click="openCreateForm" 
        class="flex items-center gap-2 bg-slate-900 dark:bg-teal-600 hover:bg-slate-800 dark:hover:bg-teal-500 text-white px-6 py-3 rounded-xl font-bold text-sm transition-all shadow-sm active:scale-95"
      >
        <UserPlus class="w-4 h-4" />
        <span>Nouveau Médecin</span>
      </button>
    </div>

    <div class="max-w-7xl mx-auto mb-8">
      <div class="relative group">
        <div class="absolute inset-y-0 left-0 pl-5 flex items-center pointer-events-none">
          <Search class="w-5 h-5 text-slate-400 group-focus-within:text-teal-500 transition-colors" />
        </div>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher par nom, licence, spécialisation..."
          class="w-full pl-14 pr-32 py-4 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 rounded-2xl border border-slate-200 dark:border-slate-800 focus:border-teal-500 dark:focus:border-teal-500 focus:ring-4 focus:ring-teal-500/10 outline-none font-medium transition-all shadow-sm"
          @input="handleSearch"
        />
        <div v-if="searchQuery" class="absolute right-4 top-1/2 -translate-y-1/2 px-3 py-1 bg-slate-100 dark:bg-slate-800 text-slate-500 dark:text-slate-400 rounded-lg text-[10px] font-bold uppercase tracking-wider border border-slate-200 dark:border-slate-700">
          {{ filteredMedecins.length }} résultat{{ filteredMedecins.length > 1 ? 's' : '' }}
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
      <div v-for="(stat, index) in [
        { label: 'Total Médecins', val: medecins.length, icon: Users, color: 'text-slate-600' },
        { label: 'Praticiens Actifs', val: activeCount, icon: CheckCircle, color: 'text-teal-600' },
        { label: 'Spécialités', val: specialitiesCount, icon: Award, color: 'text-indigo-600' },
        { label: 'Nouveaux (Mois)', val: recentCount, icon: TrendingUp, color: 'text-blue-600' }
      ]" :key="index" class="bg-white dark:bg-slate-900 rounded-2xl p-5 border border-slate-200 dark:border-slate-800 shadow-sm">
        <div class="flex items-center gap-3">
          <component :is="stat.icon" class="w-4 h-4" :class="stat.color" />
          <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest">{{ stat.label }}</p>
        </div>
        <p class="text-2xl font-bold text-slate-900 dark:text-white mt-2">{{ stat.val }}</p>
      </div>
    </div>

    <div class="max-w-7xl mx-auto bg-white dark:bg-slate-900 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full border-separate border-spacing-0">
          <thead>
            <tr class="bg-slate-50 dark:bg-slate-800/50">
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Licence</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Praticien</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Spécialité</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Contact</th>
              <th class="px-6 py-4 text-left text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Statut</th>
              <th class="px-6 py-4 text-center text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr 
              v-for="medecin in filteredMedecins" 
              :key="medecin.id"
              class="group hover:bg-slate-50/80 dark:hover:bg-slate-800/40 transition-colors"
            >
              <td class="px-6 py-4">
                <span class="font-mono text-xs font-bold text-teal-600 dark:text-teal-400 bg-teal-50 dark:bg-teal-500/10 px-2 py-1 rounded">
                  {{ medecin.licence }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="font-bold text-slate-900 dark:text-slate-200">{{ medecin.nomComplet }}</div>
                <div class="text-[10px] text-slate-500 uppercase">{{ medecin.civilite || 'Docteur' }}</div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-2 text-sm text-slate-600 dark:text-slate-400">
                  <Award class="w-3.5 h-3.5 text-slate-400" />
                  {{ medecin.specialisation || 'Généraliste' }}
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="space-y-1">
                  <div v-if="medecin.telephone" class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400">
                    <Phone class="w-3 h-3" /> {{ medecin.telephone }}
                  </div>
                  <div v-if="medecin.email" class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400">
                    <Mail class="w-3 h-3" /> {{ medecin.email }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span 
                  :class="medecin.actif ? 'text-teal-600 dark:text-teal-400 bg-teal-50 dark:bg-teal-500/10 border-teal-100 dark:border-teal-500/20' : 'text-slate-400 bg-slate-50 dark:bg-slate-800 border-slate-100 dark:border-slate-700'"
                  class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-[10px] font-bold border"
                >
                  <span :class="medecin.actif ? 'bg-teal-500 animate-pulse' : 'bg-slate-400'" class="w-1.5 h-1.5 rounded-full"></span>
                  {{ medecin.actif ? 'ACTIF' : 'INACTIF' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center justify-center gap-2">
                  <button @click="viewDetails(medecin)" class="p-2 text-slate-400 hover:text-teal-600 dark:hover:text-teal-400 hover:bg-teal-50 dark:hover:bg-teal-500/10 rounded-lg transition-all" title="Détails">
                    <Eye class="w-4 h-4" />
                  </button>
                  <button @click="editMedecin(medecin)" class="p-2 text-slate-400 hover:text-amber-600 dark:hover:text-amber-400 hover:bg-amber-50 dark:hover:bg-amber-500/10 rounded-lg transition-all" title="Modifier">
                    <Edit class="w-4 h-4" />
                  </button>
                  <button @click="deleteMedecin(medecin.id)" class="p-2 text-slate-400 hover:text-red-600 dark:hover:text-red-400 hover:bg-red-50 dark:hover:bg-red-500/10 rounded-lg transition-all" title="Supprimer">
                    <Trash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="filteredMedecins.length === 0" class="text-center py-20 bg-slate-50/50 dark:bg-slate-900/50">
        <UserX class="w-12 h-12 text-slate-300 dark:text-slate-700 mx-auto mb-4" />
        <p class="text-slate-500 dark:text-slate-400 font-medium">Aucun médecin trouvé dans le registre</p>
      </div>
    </div>

    <MedecinDetailsModal v-if="showDetails" :medecin="selectedMedecin" @close="closeDetails" @edit="editFromDetails" @delete="deleteFromDetails" />
    <MedecinForm v-if="showForm" :medecin="selectedMedecin" @close="closeForm" @saved="handleSaved" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  Stethoscope, Users, CheckCircle, Award, TrendingUp, Search, 
  Eye, Edit, Trash2, Phone, Mail, FileText, UserPlus, UserX
} from 'lucide-vue-next'
import MedecinForm from '../components/Medecins/MedecinForm.vue'
import MedecinDetailsModal from '../components/Medecins/MedecinDetailsModal.vue'
import { GetMedecins, DeleteMedecin, CreateMedecin, UpdateMedecin } from '../../wailsjs/go/main/App'

const medecins = ref([])
const searchQuery = ref('')
const showForm = ref(false)
const showDetails = ref(false)
const selectedMedecin = ref(null)

// Statistiques
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

// Filtrage
const filteredMedecins = computed(() => {
  if (!searchQuery.value) return medecins.value
  
  const query = searchQuery.value.toLowerCase()
  return medecins.value.filter(m => 
    m.nomComplet?.toLowerCase().includes(query) ||
    m.licence?.toLowerCase().includes(query) ||
    m.specialisation?.toLowerCase().includes(query) ||
    m.email?.toLowerCase().includes(query)
  )
})

const loadMedecins = async () => {
  try {
    const result = await GetMedecins()
    medecins.value = result || []
    console.log('✅ Médecins chargés:', medecins.value.length)
  } catch (err) {
    console.error('❌ Erreur:', err)
  }
}

const handleSearch = () => {
  // Filtrage géré par computed
}

// Voir détails
const viewDetails = (medecin) => {
  selectedMedecin.value = medecin
  showDetails.value = true
}

const closeDetails = () => {
  showDetails.value = false
  selectedMedecin.value = null
}

// Créer
const openCreateForm = () => {
  selectedMedecin.value = null
  showForm.value = true
}

// Modifier
const editMedecin = (medecin) => {
  selectedMedecin.value = { ...medecin }
  showForm.value = true
}

const editFromDetails = (medecin) => {
  closeDetails()
  editMedecin(medecin)
}

const closeForm = () => {
  showForm.value = false
  selectedMedecin.value = null
}

// Sauvegarder
const handleSaved = async (data) => {
  try {
    if (selectedMedecin.value?.id) {
      // Update
      await UpdateMedecin({ id: selectedMedecin.value.id, ...data })
      console.log('✅ Médecin mis à jour')
    } else {
      // Create
      await CreateMedecin(data)
      console.log('✅ Médecin créé')
    }
    await loadMedecins()
    closeForm()
  } catch (err) {
    console.error('❌ Erreur sauvegarde:', err)
    alert('Erreur: ' + err)
  }
}

// Supprimer
const deleteMedecin = async (id) => {
  if (confirm("⚠️ Supprimer ce médecin définitivement ?")) {
    try {
      await DeleteMedecin(id)
      await loadMedecins()
      console.log('✅ Médecin supprimé')
    } catch (err) {
      console.error('❌ Erreur:', err)
    }
  }
}

const deleteFromDetails = async (id) => {
  closeDetails()
  await deleteMedecin(id)
}

onMounted(() => {
  loadMedecins()
})
</script>

<style scoped>
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>