
<!-- frontend/src/views/homePage.vue -->

<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-gray-900 dark:to-gray-800 p-8 transition-colors">
    <!-- En-tête de bienvenue -->
    <div class="max-w-7xl mx-auto mb-12">
      <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-2">
        Bienvenue dans Leopard
      </h1>
      <p class="text-sm text-gray-600 dark:text-sky-300">
        Gestion complète de votre pratique privée
      </p>
    </div>

    <!-- Grille de navigation principale -->
    <div class="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
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
        title="Professionnels"
        description="Réseau de professionnels de la santé"
        icon="hospital"
        color="purple"
        :count="stats.professionnels"
        @click="navigate('professionnels')"
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

      <!-- Card Statistiques -->
      <NavCard
        title="Rapports"
        description="Statistiques et analyses"
        icon="chart"
        color="rose"
        @click="navigate('rapports')"
      />
    </div>

    <!-- Section Actions rapides -->
    <div class="max-w-7xl mx-auto mt-12">
      <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-6">
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
    <div class="max-w-7xl mx-auto mt-12">
      <div class="bg-white dark:bg-gray-800 rounded-xl shadow-lg p-6">
        <h2 class="text-xl font-bold text-gray-900 dark:text-white mb-4">
          Activité récente
        </h2>
        <div class="space-y-3">
          <ActivityItem
            v-for="activity in recentActivities"
            :key="activity.id"
            :activity="activity"
          />
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

const router = useRouter()

const stats = ref({
  clients: 0,
  medecins: 0,
  notaires: 0,
  professionnels: 0,
  etablissements: 0,
})

const recentActivities = ref([

])

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

// Charger les statistiques au montage
onMounted(async () => {
  // TODO: Appeler les fonctions Wails pour récupérer les stats
  // const clientCount = await GetClientsCount()
  // stats.value.clients = clientCount
})
</script>