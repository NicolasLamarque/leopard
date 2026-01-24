<template>
  <div class="relative">
    <!-- Header -->
    <div class="flex items-center justify-between mb-4 px-1">
      <h4 class="text-sm font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wide flex items-center gap-2">
        <div class="p-1.5 bg-teal-100 dark:bg-teal-900/30 rounded-lg">
          <Users :size="14" class="text-teal-600 dark:text-teal-400" />
        </div>
        <span>Clients rattachés</span>
        <span class="ml-1 px-2 py-0.5 bg-teal-600 text-white rounded-full text-xs font-semibold">
          {{ clients.length }}
        </span>
      </h4>
      
      <div v-if="loading" class="animate-spin text-teal-600 dark:text-teal-400">
        <Loader2 :size="18" />
      </div>
    </div>

    <!-- Liste des clients -->
    <div v-if="clients.length > 0" class="space-y-2">
      <div 
        v-for="client in clients" :key="client.id"
        class="group flex items-center justify-between p-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl hover:border-teal-400 dark:hover:border-teal-600 transition-all hover:shadow-sm cursor-pointer"
        @click="handleNavigate(client.id)"
      >
        <div class="flex items-center gap-3">
          <!-- Avatar -->
          <div class="w-11 h-11 rounded-xl bg-teal-50 dark:bg-teal-900/20 flex items-center justify-center text-sm font-semibold text-teal-700 dark:text-teal-400 border border-teal-200 dark:border-teal-800">
            {{ getInitials(client.nom, client.prenom) }}
          </div>
          
          <div>
            <p class="font-semibold text-sm text-slate-900 dark:text-white group-hover:text-teal-600 dark:group-hover:text-teal-400 transition-colors">
              {{ client.prenom }} {{ client.nom }}
            </p>
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mt-0.5">
              <span class="font-mono bg-slate-100 dark:bg-slate-800 px-2 py-0.5 rounded">
                {{ client.no_dossier_leopard || 'SANS-DOSSIER' }}
              </span>
              <span>•</span>
              <span>{{ calculateAge(client.date_naissance) }}</span>
            </div>
          </div>
        </div>

        <button class="p-2 rounded-lg bg-slate-100 dark:bg-slate-800 text-slate-400 group-hover:bg-teal-600 group-hover:text-white transition-all">
          <ExternalLink :size="16" />
        </button>
      </div>
    </div>

    <!-- Empty state -->
    <div v-else-if="!loading" class="flex flex-col items-center justify-center py-12 px-4 border-2 border-dashed border-slate-200 dark:border-slate-700 rounded-xl bg-white dark:bg-slate-900">
      <div class="p-4 bg-slate-100 dark:bg-slate-800 rounded-2xl mb-3">
        <UserPlus :size="32" class="text-slate-400 dark:text-slate-600" />
      </div>
      <p class="text-sm font-medium text-slate-600 dark:text-slate-400">
        Aucun client lié à ce notaire
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { Users, Loader2, ExternalLink, UserPlus } from 'lucide-vue-next'

const props = defineProps({
  notaireId: { type: Number, required: true }
})

const clients = ref([])
const loading = ref(false)

const loadLinkedClients = async () => {
  if (!props.notaireId) return
  loading.value = true
  try {
    clients.value = await window.go.main.App.GetClientsByNotaire(props.notaireId)
  } catch (err) {
    console.error("Erreur lors du chargement des clients liés:", err)
  } finally {
    loading.value = false
  }
}

const getInitials = (nom, prenom) => {
  return `${prenom?.charAt(0) || ''}${nom?.charAt(0) || ''}`.toUpperCase()
}

const calculateAge = (dob) => {
  if (!dob) return 'Âge inconnu'
  const birthDate = new Date(dob)
  const today = new Date()
  let age = today.getFullYear() - birthDate.getFullYear()
  const m = today.getMonth() - birthDate.getMonth()
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--
  }
  return `${age} ans`
}

const handleNavigate = (clientId) => {
  console.log("Navigation vers le client ID:", clientId)
}

watch(() => props.notaireId, () => {
  loadLinkedClients()
})

onMounted(loadLinkedClients)
</script>