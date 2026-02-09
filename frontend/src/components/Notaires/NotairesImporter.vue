<!-- frontend/src/components/Medecin/MedecinImporter.vue -->
<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
    <div class="mb-6">
      <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">
        Import de Notaires
      </h2>
      <p class="text-sm text-gray-600 dark:text-gray-400">
        Importez une liste de Notaires depuis un fichier Excel (.xlsx)
      </p>
    </div>

    <!-- Zone de sélection de fichier -->
   <!-- REMPLACEZ toute la section "Zone de sélection de fichier" -->
<div class="mb-6">
  <div class="border-2 border-dashed rounded-lg p-8 text-center border-gray-300 dark:border-gray-600 hover:border-blue-400 transition-all">
    <Upload :size="48" class="mx-auto mb-4 text-gray-400 dark:text-gray-500" />
    
    <div v-if="!selectedFilePath">
      <p class="text-sm font-medium text-gray-900 dark:text-white mb-2">
        Sélectionnez un fichier Excel
      </p>
      <p class="text-xs text-gray-500 dark:text-gray-400 mb-4">
        Cliquez pour ouvrir le sélecteur de fichier
      </p>
      <button 
        @click="selectFile"
        class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
      >
        <FileSpreadsheet :size="16" />
        Choisir un fichier
      </button>
    </div>

    <div v-else class="space-y-3">
      <div class="flex items-center justify-center gap-3 text-green-600 dark:text-green-400">
        <CheckCircle :size="24" />
        <span class="font-medium">{{ getFileName(selectedFilePath) }}</span>
      </div>
      <p class="text-xs text-gray-500 dark:text-gray-400 break-all">{{ selectedFilePath }}</p>
      <button 
        @click="clearFile"
        class="text-sm text-red-600 hover:text-red-700 dark:text-red-400"
      >
        Changer de fichier
      </button>
    </div>
  </div>
</div>

    <!-- Options d'import -->
    <div class="mb-6 p-4 bg-gray-50 dark:bg-gray-900/50 rounded-lg">
      <label class="flex items-center gap-3">
        <input 
          v-model="updateExisting"
          type="checkbox" 
          class="w-5 h-5 text-blue-600 rounded focus:ring-2 focus:ring-blue-500"
        />
        <div>
          <p class="font-medium text-gray-900 dark:text-white">Mettre à jour les notaires existants</p>
          <p class="text-xs text-gray-500 dark:text-gray-400">Si un notaire existe déjà (même numéro de licence), ses informations seront mises à jour</p>
        </div>
      </label>
    </div>

    <!-- Bouton d'import -->
    <div class="flex items-center gap-4">
      <button 
        @click="importNotaires"
        :disabled="!selectedFilePath || importing"
        class="flex-1 flex items-center justify-center gap-2 px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors font-medium"
      >
        <Upload v-if="!importing" :size="20" />
        <Loader2 v-else :size="20" class="animate-spin" />
        {{ importing ? 'Import en cours...' : 'Démarrer l\'import' }}
      </button>

      <button 
        v-if="importResult"
        @click="clearResult"
        class="px-4 py-3 text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white transition-colors"
      >
        Effacer
      </button>
    </div>

    <!-- Résultat de l'import -->
    <div 
      v-if="importResult" 
      class="mt-6 p-4 rounded-lg animate-slide-down"
      :class="importResult.success ? 'bg-green-50 dark:bg-green-900/20' : 'bg-red-50 dark:bg-red-900/20'"
    >
      <div class="flex items-start gap-3">
        <CheckCircle 
          v-if="importResult.success" 
          :size="24" 
          class="text-green-600 dark:text-green-400 flex-shrink-0"
        />
        <AlertCircle 
          v-else 
          :size="24" 
          class="text-red-600 dark:text-red-400 flex-shrink-0"
        />
        
        <div class="flex-1">
          <p 
            class="font-medium mb-2"
            :class="importResult.success ? 'text-green-900 dark:text-green-300' : 'text-red-900 dark:text-red-300'"
          >
            {{ importResult.message }}
          </p>

          <!-- Statistiques détaillées -->
          <div v-if="importResult.stats" class="mt-3 space-y-2">
            <div class="flex items-center gap-2 text-sm">
              <span class="text-green-600 dark:text-green-400">✓ {{ importResult.stats.nouveaux }} nouveaux</span>
              <span v-if="updateExisting" class="text-blue-600 dark:text-blue-400">↻ {{ importResult.stats.mis_a_jour }} mis à jour</span>
              <span v-if="importResult.stats.errors?.length" class="text-red-600 dark:text-red-400">✗ {{ importResult.stats.errors.length }} erreurs</span>
            </div>

            <!-- Liste des erreurs -->
            <div v-if="importResult.stats.errors?.length" class="mt-3">
              <button 
                @click="showErrors = !showErrors"
                class="text-sm text-red-600 dark:text-red-400 hover:underline"
              >
                {{ showErrors ? '▼' : '▶' }} Voir les erreurs ({{ importResult.stats.errors.length }})
              </button>
              
              <div v-if="showErrors" class="mt-2 max-h-64 overflow-y-auto">
                <div class="space-y-1 text-xs font-mono bg-white dark:bg-gray-800 p-3 rounded">
                  <p v-for="(error, i) in importResult.stats.errors" :key="i" class="text-red-600 dark:text-red-400">
                    {{ error }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Upload, FileSpreadsheet, CheckCircle, AlertCircle, Loader2 } from 'lucide-vue-next'

const selectedFilePath = ref(null)
const importing = ref(false)
const updateExisting = ref(false)
const importResult = ref(null)
const showErrors = ref(false)

// ✅ VERSION SIMPLE comme votre clientFolderManager
const selectFile = async () => {
  try {
    console.log('🔍 Sélection fichier...')
    
    // ✅ Appel direct comme vous faites pour les clients
    const filePath = await window.go.main.App.SelectExcelFile()
    
    if (filePath) {
      console.log('✅ Fichier sélectionné:', filePath)
      selectedFilePath.value = filePath
    } else {
      console.log('ℹ️ Sélection annulée')
    }
  } catch (error) {
    console.error('❌ Erreur:', error)
    alert('Erreur: ' + error)
  }
}

const getFileName = (path) => {
  if (!path) return ''
  return path.split(/[/\\]/).pop()
}

const clearFile = () => {
  selectedFilePath.value = null
}

const clearResult = () => {
  importResult.value = null
  showErrors.value = false
}

const importNotaires = async () => {
  if (!selectedFilePath.value) {
    alert('Veuillez sélectionner un fichier')
    return
  }
  
  importing.value = true
  importResult.value = null
  
  try {
    console.log('🚀 Import:', selectedFilePath.value)
    
    let result
    if (updateExisting.value) {
      // ✅ Appel direct comme vos autres fonctions
      result = await window.go.main.App.ImportNotairesWithUpdate(selectedFilePath.value)
      
      importResult.value = {
        success: true,
        message: `✅ Import terminé !`,
        stats: result
      }
    } else {
      result = await window.go.main.App.ImportNotaires(selectedFilePath.value)
      
      importResult.value = {
        success: true,
        message: result
      }
    }
    
    clearFile()
  } catch (error) {
    console.error('❌ Erreur:', error)
    importResult.value = {
      success: false,
      message: `❌ ${error}`
    }
  } finally {
    importing.value = false
  }
}
</script>

<style scoped>
@keyframes slide-down {
  from {
    transform: translateY(-10px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

.animate-slide-down {
  animation: slide-down 0.3s ease-out;
}
</style>