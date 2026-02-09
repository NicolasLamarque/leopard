<template>   
 <div class="min-h-screen bg-gray-50 dark:bg-gray-950 transition-colors p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <!-- Header -->
<div class="flex justify-between items-center mb-6">
  <div>
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
      Dossier de {{ store.currentClient?.prenom }} {{ store.currentClient?.nom }}
    </h1>
    <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
      No Dossier: {{ store.currentClient?.no_dossier_leopard || 'Non généré' }}
    </p>
  </div>
  
  <!-- Boutons d'action -->
  <div class="flex items-center gap-3">

    <button 
    @click="isPIOpen = true"
    class="flex items-center gap-2 bg-teal-600 hover:bg-teal-700 text-white px-5 py-2.5 rounded-lg shadow-sm transition-all duration-200"
  >
    <FilePlus :size="18" />
    <span class="font-medium">Nouveau PI</span>
  </button>
    <!-- Bouton Notes -->
    <button 
      @click="showNotes = true"
      class="flex items-center gap-2 bg-blue-600 hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 text-white px-5 py-2.5 rounded-lg shadow-sm transition-all duration-200"
    >
      <FileText :size="18" />
      <span class="font-medium">Notes</span>
    </button>

    <!-- Bouton Fermer -->
    <button 
      @click="router.push('/app/clients')" 
      class="flex items-center gap-2 bg-gray-200 dark:bg-gray-800 hover:bg-gray-300 dark:hover:bg-gray-700 text-gray-700 dark:text-gray-300 px-5 py-2.5 rounded-lg shadow-sm transition-all duration-200 border border-gray-300 dark:border-gray-700"
    >
      <ArrowLeft :size="18" />
      <span class="font-medium">Fermer</span>
    </button>
  </div>
</div>

      <!-- Loading State -->
      <div v-if="isLoading" class="flex items-center justify-center py-12">
        <div class="flex flex-col items-center gap-3">
          <Loader2 :size="32" class="animate-spin text-blue-600 dark:text-blue-400" />
          <p class="text-gray-600 dark:text-gray-400">Chargement du dossier...</p>
        </div>
      </div>

      <!-- Client Form -->



      
      <ClientDetailsForm 
        v-else-if="store.currentClient" 
        :clientData="store.currentClient" 
        @save="onSave"
        @folderCreated="onFolderCreated"
        @show-notes="showNotes = true"
        @show-evaluations="isEvalOpen = true"
      />

      <!-- Error State -->
      <div v-else class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl p-6">
        <div class="flex items-start gap-3">
          <AlertCircle :size="24" class="text-red-600 dark:text-red-400 flex-shrink-0 mt-0.5" />
          <div>
            <h3 class="text-lg font-semibold text-red-900 dark:text-red-100 mb-1">
              Client introuvable
            </h3>
            <p class="text-sm text-red-700 dark:text-red-300">
              Impossible de charger les informations du client. Veuillez réessayer.
            </p>
            <button 
              @click="router.push('/app/clients')"
              class="mt-4 px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-colors"
            >
              Retour à la liste
            </button>
          </div>
        </div>
      </div>
    </div>
  <!-- Drawer Notes -->
    <NotesDrawer
  :is-open="showNotes"
  :client="store.currentClient"
  @close="showNotes = false"
/>



<!-- Drawer Evaluations -->
<EvaluationView
  :is-open="isEvalOpen"
  :client="store.currentClient"
  @close="isEvalOpen = false"
/>

<div v-if="isPIOpen" class="fixed inset-0 z-[100] bg-black/50 backdrop-blur-sm flex items-center justify-center p-4">
      <div class="bg-white dark:bg-gray-900 w-full max-w-6xl max-h-[95vh] overflow-y-auto rounded-2xl shadow-2xl">
        <PIViewer
          :client="store.currentClient"
          @close="isPIOpen = false"
        />
      </div>
    </div>
<!-- Client Details Form -->
<!-- ClientDetailsForm 
  :clientData="store.currentClient" 
  @save="onSave"
  @folderCreated="onFolderCreated"
  @show-notes="showNotes = true"
/-->
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useClientStore } from '../stores/clientStore'
import { UpdateClient } from '../../wailsjs/go/main/App'
import ClientDetailsForm from '../components/clients/ClientDetailsForms.vue'
import { ArrowLeft, Loader2, AlertCircle, FileText, FilePlus, } from 'lucide-vue-next'
import NotesDrawer from '../components/Notes/NotesDrawer.vue'  
import {GetAllContactsByClientID, CreateContact, UpdateContact, DeleteContact} from '../../wailsjs/go/main/App'
import { createClientFolder } from '@/utils/clientFolderManager'
import EvaluationView from '../components/evaluation/EvaluationView.vue'

import PIViewer from '../components/pi/PIViewer.vue' // Ton composant PI

const isPIOpen = ref(false) // Pour contrôler l'ouverture
const showNotes = ref(false)

const route = useRoute()
const store = useClientStore()
const router = useRouter()
const isLoading = ref(true)
const isEvalOpen = ref(false)

const onSave = async (data) => {
  try {
    await UpdateClient(data)
    
    // Notification de succès moderne
    showNotification('✅ Modifications enregistrées avec succès!')
    
    // Rafraîchir les données du client
    const id = parseInt(route.params.id)
    await store.selectClient(id)
  } catch (err) {
    console.error('❌ Erreur sauvegarde:', err)
    showNotification('❌ Erreur lors de la sauvegarde', 'error')
  }
}

const onFolderCreated = (folderData) => {
  console.log('📁 Dossier créé:', folderData)
  showNotification(`📁 Dossier créé: ${folderData.leopardNumber}`)
}

// Fonction de notification simple
const showNotification = (message, type = 'success') => {
  // Tu peux remplacer par un système de toast plus sophistiqué
  alert(message)
}

onMounted(async () => {
  try {
    isLoading.value = true
    const id = parseInt(route.params.id)
    await store.selectClient(id)
  } catch (err) {
    console.error('❌ Erreur chargement client:', err)
  } finally {
    isLoading.value = false
  }
})
</script>

<style scoped>
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>