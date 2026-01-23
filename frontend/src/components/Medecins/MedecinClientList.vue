<!-- frontend/src/components/Medecins/MedecinClientsList.vue -->
<template>
  <section class="border-t border-slate-200 dark:border-slate-700 pt-6">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-lg font-bold text-slate-900 dark:text-white flex items-center gap-2">
        <Users :size="20" class="text-slate-600 dark:text-slate-400" />
        Patients en commun
        <span v-if="!loading" class="text-sm font-normal text-slate-500 dark:text-slate-400">
          ({{ clients.length }})
        </span>
      </h3>
      
      <!-- Recherche -->
      <div v-if="clients.length > 0" class="relative">
        <Search :size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher un patient..."
          class="pl-9 pr-4 py-2 text-sm border border-slate-200 dark:border-slate-700 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-slate-500"
        />
      </div>
    </div>

    <!-- Chargement -->
    <div v-if="loading" class="flex items-center justify-center py-8">
      <Loader2 :size="24" class="animate-spin text-slate-400" />
      <span class="ml-3 text-sm text-slate-600 dark:text-slate-400">Chargement des patients...</span>
    </div>

    <!-- Erreur -->
    <div v-else-if="error" class="p-4 bg-rose-50 dark:bg-rose-900/20 rounded-lg border border-rose-200 dark:border-rose-800">
      <p class="text-sm text-rose-700 dark:text-rose-300">{{ error }}</p>
    </div>

    <!-- État vide -->
    <div v-else-if="clients.length === 0" class="text-center py-8">
      <div class="inline-flex items-center justify-center w-16 h-16 bg-slate-100 dark:bg-slate-800 rounded-full mb-4">
        <Users :size="32" class="text-slate-400" />
      </div>
      <p class="text-sm font-medium text-slate-700 dark:text-slate-300 mb-1">
        Aucun patient en commun
      </p>
      <p class="text-xs text-slate-500 dark:text-slate-400">
        Ce médecin n'a pas encore de patients assignés
      </p>
    </div>

    <!-- Liste des patients -->
    <div v-else class="space-y-2 max-h-[400px] overflow-y-auto pr-2">
      <div
        v-for="client in filteredClients"
        :key="client.id"
        class="group flex items-center gap-4 p-3 rounded-lg transition-all duration-200 hover:bg-slate-50 dark:hover:bg-slate-800/50 border border-transparent hover:border-slate-200 dark:hover:border-slate-700"
      >
        <!-- Avatar -->
        <div class="flex-shrink-0 w-10 h-10 bg-slate-600 rounded-full flex items-center justify-center">
          <span class="text-sm font-semibold text-white">
            {{ getInitials(client.nom, client.prenom) }}
          </span>
        </div>

        <!-- Info patient -->
        <div class="flex-1 min-w-0">
          <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
            {{ client.prenom }} {{ client.nom }}
          </p>
          <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400">
            <span v-if="client.no_dossier_leopard" class="font-mono">
              # {{ client.no_dossier_leopard }}
            </span>
            <span v-if="client.date_naissance" class="flex items-center gap-1">
              <Calendar :size="12" />
              {{ formatDateNaissance(client.date_naissance) }}
            </span>
          </div>
        </div>

        <!-- Actions -->
        <button
          @click="$emit('view-client', client.id)"
          class="opacity-0 group-hover:opacity-100 px-3 py-1.5 bg-slate-700 hover:bg-slate-800 text-white rounded-lg text-xs font-medium transition-all duration-200 flex items-center gap-1.5"
        >
          <FileText :size="14" />
          Voir dossier
        </button>
      </div>

      <!-- Message si recherche vide -->
      <div v-if="filteredClients.length === 0 && searchQuery" class="text-center py-6">
        <p class="text-sm text-slate-600 dark:text-slate-400">
          Aucun patient trouvé pour "{{ searchQuery }}"
        </p>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Users, Search, Loader2, FileText, Calendar } from 'lucide-vue-next'
import medecinService from '@/utils/medecinService'

const props = defineProps({
  medecinLicence: { type: String, required: true }
})

const emit = defineEmits(['view-client'])

const clients = ref([])
const loading = ref(false)
const error = ref(null)
const searchQuery = ref('')

const filteredClients = computed(() => {
  if (!searchQuery.value) return clients.value
  
  const query = searchQuery.value.toLowerCase()
  return clients.value.filter(client => {
    const fullName = `${client.prenom} ${client.nom}`.toLowerCase()
    const dossier = client.no_dossier_leopard?.toLowerCase() || ''
    return fullName.includes(query) || dossier.includes(query)
  })
})

const getInitials = (nom, prenom) => {
  const n = nom?.charAt(0)?.toUpperCase() || ''
  const p = prenom?.charAt(0)?.toUpperCase() || ''
  return `${p}${n}`
}

const formatDateNaissance = (date) => {
  if (!date) return ''
  try {
    const d = new Date(date)
    const age = new Date().getFullYear() - d.getFullYear()
    return `${age} ans`
  } catch {
    return ''
  }
}

const loadClients = async () => {
  loading.value = true
  error.value = null
  
  try {
    clients.value = await medecinService.getClients(props.medecinLicence)
  } catch (err) {
    error.value = 'Impossible de charger les patients'
    console.error('Erreur chargement clients:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadClients()
})
</script>

<style scoped>
/* Scrollbar personnalisée */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.dark .overflow-y-auto::-webkit-scrollbar-thumb {
  background: #475569;
}
</style>