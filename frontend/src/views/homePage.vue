<!-- frontend/src/views/homePage.vue -->

<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50/30 to-purple-50/20 dark:from-gray-900 dark:via-slate-900/80 dark:to-gray-800 p-8 transition-colors duration-500">
    
    <!-- En-tête de bienvenue -->
    <div class="max-w-7xl mx-auto mb-10">
      <h1 class="text-4xl font-semibold bg-gradient-to-r from-gray-800 to-gray-600 dark:from-gray-100 dark:to-gray-300 bg-clip-text text-transparent mb-2">
        Bienvenue dans Leopard
      </h1>
      <p class="text-sm text-gray-600 dark:text-slate-400 font-light">
        Gestion complète de votre pratique privée
      </p>
    </div>

    <!-- Grille de navigation principale -->
    <div class="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5 mb-10">
      
      <!-- Card Clients -->
      <NavCard
        title="Clients"
        description="Gérer les dossiers clients et consultations"
        icon="users"
        color="blue"
        :count="stats.clients"
        @click="navigate('clients')"
      />

      <!-- Card Médecins -->
      <NavCard
        title="Médecins"
        description="Répertoire des médecins référents"
        icon="stethoscope"
        color="emerald"
        :count="stats.medecins"
        @click="navigate('medecins')"
      />

      <!-- Card Notaires -->
      <NavCard
        title="Notaires"
        description="Contacts et informations notariales"
        icon="briefcase"
        color="amber"
        :count="stats.notaires"
        @click="navigate('notaires')"
      />

      <!-- Card Professionnels de santé -->
 <NavCard
        title="Intervenants"
        description="Gestion du réseau social et partenaires"
        icon="user-cog"
        color="emerald"
        :count="stats.intervenants"
        @click="navigate('intervenants')"
      />
      <!-- Card Établissements -->
      <NavCard
        title="Établissements"
        description="RPA, CHSLD et résidences"
        icon="building"
        color="indigo"
        :count="stats.etablissements"
        @click="navigate('etablissements')"
      />

      <!-- Card Appels (NOUVEAU) -->
      <NavCard
  title="Appels téléphoniques"
  description="Gestion des demandes et suivis"
  icon="phone"
  color="cyan"
  :count="stats.appels"
  @click="openAppels" 
/>

      <!-- Card Rapports -->
      <NavCard
        title="Rapports"
        description="Statistiques et analyses"
        icon="chart"
        color="rose"
        @click="navigate('rapports')"
      />
      
    </div>

    <!-- Section Actions rapides -->
    <div class="max-w-7xl mx-auto mb-10">
      <h2 class="text-xl font-semibold text-gray-800 dark:text-gray-200 mb-5">
        Actions rapides
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <QuickAction
          title="Nouveau client"
          icon="user-plus"
          @click="quickAction('new-client')"
        />
        <QuickAction
          title="Nouvelle note"
          icon="file-text"
          @click="quickAction('new-note')"
        />
        <QuickAction
          title="Rechercher"
          icon="search"
          @click="quickAction('search')"
        />
      </div>
    </div>

    <!-- Activité récente -->
    <div class="max-w-7xl mx-auto">
      <div class="bg-white/70 dark:bg-gray-800/50 backdrop-blur-sm rounded-xl shadow-md border border-gray-100 dark:border-gray-700/50 p-6">
        <h2 class="text-xl font-semibold text-gray-800 dark:text-gray-200 mb-4">
          Activité récente
        </h2>
        
        <div v-if="recentActivities.length > 0" class="space-y-3">
          <ActivityItem
            v-for="activity in recentActivities"
            :key="activity.id"
            :activity="activity"
          />
        </div>
        
        <div v-else class="text-center py-12 text-gray-400 dark:text-gray-500 text-sm">
          Aucune activité récente
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import NavCard from '../components/NavCard.vue'
import QuickAction from '../components/QuickAction.vue'
import ActivityItem from '../components/ActivityItem.vue'

import { useAppels } from '../composables/useAppels'

const router = useRouter()
const { loadStats: loadAppelsStats, stats: appelsStats } = useAppels()

const stats = ref({
  clients: 0,
  medecins: 0,
  notaires: 0,
  professionnels: 0,
   intervenants: 0,
  etablissements: 0,
  appels: 0
})

const recentActivities = ref([])

// Modifie la fonction pour les appels
const openAppels = () => {
  router.push('/app/appels')
}
const navigate = (section) => {
  router.push(`/app/${section}`)
}

const quickAction = (action) => {
  switch(action) {
    case 'new-client':
      router.push('/app/clients?action=new')
      break
    case 'new-note':
      // Implémenter la logique de création rapide
      break
    case 'search':
      // Implémenter la recherche globale
      break
  }
}

const handleAppelCreated = async () => {
  showNouvelAppelModal.value = false
  // Recharger les stats après création
  await loadAppelsStats()
  stats.value.appels = appelsStats.value.total || 0
}

// Charger les statistiques au montage
onMounted(async () => {
  // Charger les stats des appels
  await loadAppelsStats()
  stats.value.appels = appelsStats.value.total || 0
  
  // TODO: Appeler les autres fonctions Wails pour récupérer les stats
  // const clientCount = await GetClientsCount()
  // stats.value.clients = clientCount
})
</script>