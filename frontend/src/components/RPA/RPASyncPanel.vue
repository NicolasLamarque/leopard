<template>
  <div
    @click.self="handleClose"
    class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4"
  >
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-2xl max-w-2xl w-full">
      
      <!-- Header -->
      <div class="px-6 py-4 border-b dark:border-gray-700 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <RefreshCw :size="24" :class="syncing ? 'animate-spin text-blue-600' : 'text-gray-600'" />
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
            Synchronisation MSSS
          </h2>
        </div>
        <button
          @click="handleClose"
          :disabled="syncing"
          class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 disabled:opacity-50"
        >
          <X :size="24" />
        </button>
      </div>

      <!-- Body -->
      <div class="p-6 space-y-6">
        
        <!-- Statut de synchronisation -->
        <div v-if="!synced && !error" class="text-center">
          <div v-if="!syncing" class="space-y-4">
            <Database :size="64" class="mx-auto text-blue-500" />
            <div>
              <p class="text-lg font-medium text-gray-900 dark:text-white">
                Synchroniser avec le registre MSSS
              </p>
              <p class="text-sm text-gray-500 dark:text-gray-400 mt-2">
                Cette opération va récupérer la liste complète des RPA du Québec et mettre à jour votre base de données.
              </p>
            </div>

            <div class="bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 rounded-lg p-4">
              <h3 class="text-sm font-semibold text-blue-900 dark:text-blue-300 mb-2">
                Ce qui sera synchronisé:
              </h3>
              <ul class="text-sm text-blue-800 dark:text-blue-400 space-y-1">
                <li>✓ Nouveaux RPA ajoutés</li>
                <li>✓ Informations mises à jour (nom, adresse, etc.)</li>
                <li>✓ RPA fermés marqués automatiquement</li>
                <li>✓ Historique des changements conservé</li>
              </ul>
            </div>

            <p class="text-xs text-gray-500 dark:text-gray-400">
              ⏱️ Temps estimé: 30-60 secondes
            </p>
          </div>

          <!-- En cours de sync -->
          <div v-else class="space-y-6">
            <div class="inline-block animate-spin rounded-full h-16 w-16 border-4 border-blue-500 border-t-transparent"></div>
            <div>
              <p class="text-lg font-medium text-gray-900 dark:text-white">
                Synchronisation en cours...
              </p>
              <p class="text-sm text-gray-500 dark:text-gray-400 mt-2">
                {{ syncMessage }}
              </p>
            </div>

            <!-- Barre de progression (approximative) -->
            <div class="max-w-md mx-auto">
              <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2">
                <div
                  class="bg-blue-600 h-2 rounded-full transition-all duration-500"
                  :style="{ width: progress + '%' }"
                ></div>
              </div>
              <p class="text-xs text-gray-500 dark:text-gray-400 mt-2 text-center">
                {{ progress }}%
              </p>
            </div>
          </div>
        </div>

        <!-- Résultats de la synchronisation -->
        <div v-if="synced && stats" class="space-y-4">
          <div class="text-center">
            <CheckCircle :size="64" class="mx-auto text-green-500 mb-4" />
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
              Synchronisation terminée!
            </h3>
          </div>

          <!-- Statistiques -->
          <div class="grid grid-cols-2 gap-4">
            <div class="bg-blue-50 dark:bg-blue-900/20 rounded-lg p-4 text-center">
              <p class="text-3xl font-bold text-blue-600 dark:text-blue-400">
                {{ stats.total_scraped }}
              </p>
              <p class="text-sm text-blue-800 dark:text-blue-300 mt-1">Total récupéré</p>
            </div>

            <div class="bg-green-50 dark:bg-green-900/20 rounded-lg p-4 text-center">
              <p class="text-3xl font-bold text-green-600 dark:text-green-400">
                {{ stats.nouveaux }}
              </p>
              <p class="text-sm text-green-800 dark:text-green-300 mt-1">Nouveaux</p>
            </div>

            <div class="bg-yellow-50 dark:bg-yellow-900/20 rounded-lg p-4 text-center">
              <p class="text-3xl font-bold text-yellow-600 dark:text-yellow-400">
                {{ stats.mis_a_jour }}
              </p>
              <p class="text-sm text-yellow-800 dark:text-yellow-300 mt-1">Mis à jour</p>
            </div>

            <div class="bg-red-50 dark:bg-red-900/20 rounded-lg p-4 text-center">
              <p class="text-3xl font-bold text-red-600 dark:text-red-400">
                {{ stats.fermes }}
              </p>
              <p class="text-sm text-red-800 dark:text-red-300 mt-1">Fermés</p>
            </div>
          </div>

          <!-- Erreurs -->
          <div v-if="stats.erreurs && stats.erreurs.length > 0" class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
            <h4 class="text-sm font-semibold text-red-800 dark:text-red-300 mb-2 flex items-center gap-2">
              <AlertCircle :size="18" />
              Erreurs rencontrées ({{ stats.erreurs.length }})
            </h4>
            <ul class="text-sm text-red-700 dark:text-red-400 space-y-1 max-h-32 overflow-y-auto">
              <li v-for="(err, i) in stats.erreurs.slice(0, 5)" :key="i" class="text-xs">
                • {{ err }}
              </li>
              <li v-if="stats.erreurs.length > 5" class="text-xs italic">
                ... et {{ stats.erreurs.length - 5 }} autres erreurs
              </li>
            </ul>
          </div>

          <p class="text-xs text-gray-500 dark:text-gray-400 text-center">
            Synchronisé le {{ formatDate(stats.date_sync) }}
          </p>
        </div>

        <!-- Erreur globale -->
        <div v-if="error" class="text-center space-y-4">
          <XCircle :size="64" class="mx-auto text-red-500" />
          <div>
            <h3 class="text-xl font-semibold text-red-900 dark:text-red-300">
              Erreur de synchronisation
            </h3>
            <p class="text-sm text-red-700 dark:text-red-400 mt-2">
              {{ error }}
            </p>
          </div>
        </div>

      </div>

      <!-- Footer -->
      <div class="px-6 py-4 border-t dark:border-gray-700 flex justify-end gap-3 bg-gray-50 dark:bg-gray-900">
        <button
          v-if="!syncing && !synced"
          @click="handleClose"
          class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
        >
          Annuler
        </button>

        <button
          v-if="!syncing && !synced"
          @click="handleSync"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
        >
          <RefreshCw :size="18" />
          Lancer la synchronisation
        </button>

        <button
          v-if="synced || error"
          @click="handleClose"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
        >
          Fermer
        </button>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { RefreshCw, X, Database, CheckCircle, XCircle, AlertCircle } from 'lucide-vue-next'
import { SyncRPA } from '../../../wailsjs/go/main/App'

const emit = defineEmits(['close', 'synced'])

const syncing = ref(false)
const synced = ref(false)
const error = ref(null)
const stats = ref(null)
const progress = ref(0)
const syncMessage = ref('Connexion au serveur MSSS...')

const handleSync = async () => {
  syncing.value = true
  error.value = null
  progress.value = 0
  
  // Animation de progression
  const progressInterval = setInterval(() => {
    if (progress.value < 90) {
      progress.value += Math.random() * 15
      
      if (progress.value < 30) {
        syncMessage.value = 'Récupération de la liste des RPA...'
      } else if (progress.value < 60) {
        syncMessage.value = 'Analyse des données...'
      } else {
        syncMessage.value = 'Mise à jour de la base de données...'
      }
    }
  }, 800)

  try {
    const result = await SyncRPA()
    
    clearInterval(progressInterval)
    progress.value = 100
    syncMessage.value = 'Synchronisation terminée!'
    
    stats.value = result
    synced.value = true
    
    //setTimeout(() => {
      //emit('synced', result)
    //}, 2000)
  } catch (err) {
    clearInterval(progressInterval)
    console.error('Erreur synchronisation:', err)
    error.value = err.message || 'Une erreur est survenue lors de la synchronisation'
  } finally {
    syncing.value = false
  }
}

const handleClose = () => {
  if (!syncing.value) {
    emit('close')
  }
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
</script>