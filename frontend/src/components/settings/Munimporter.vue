<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6">
    <div>
      <h2 class="text-xl font-semibold text-gray-900 dark:text-white flex items-center gap-2">
        <MapPin :size="20" class="text-teal-600" />
        Import des données géographiques
      </h2>
      <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
        Importe les municipalités et arrondissements depuis les CSV du gouvernement du Québec.<br>
        Les fichiers doivent être dans le dossier <code class="bg-gray-100 dark:bg-gray-700 px-1 rounded text-xs">data/</code> de l'application.
      </p>
    </div>

    <!-- MUN.csv -->
    <div class="border dark:border-gray-700 rounded-lg p-4 space-y-3">
      <div class="flex items-center justify-between">
        <div>
          <p class="font-medium text-gray-900 dark:text-white">Municipalités</p>
          <p class="text-xs text-gray-500 dark:text-gray-400 font-mono">data/MUN.csv</p>
        </div>
        <div class="flex items-center gap-2">
          <span v-if="munResult !== null" class="text-xs px-2 py-1 rounded-full font-medium"
            :class="munResult.error 
              ? 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300'
              : 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300'">
            {{ munResult.error ? munResult.error : `${munResult.count} importées` }}
          </span>
          <button
            @click="importMun"
            :disabled="loadingMun"
            class="flex items-center gap-1.5 px-4 py-2 rounded-lg text-sm font-medium bg-teal-600 hover:bg-teal-700 text-white disabled:opacity-50 disabled:cursor-not-allowed transition-all"
          >
            <Loader2 v-if="loadingMun" :size="14" class="animate-spin" />
            <Upload v-else :size="14" />
            {{ loadingMun ? 'Import...' : 'Importer' }}
          </button>
        </div>
      </div>
    </div>

    <!-- ARR.csv -->
    <div class="border dark:border-gray-700 rounded-lg p-4 space-y-3">
      <div class="flex items-center justify-between">
        <div>
          <p class="font-medium text-gray-900 dark:text-white">Arrondissements / Secteurs</p>
          <p class="text-xs text-gray-500 dark:text-gray-400 font-mono">data/ARR.csv</p>
        </div>
        <div class="flex items-center gap-2">
          <span v-if="arrResult !== null" class="text-xs px-2 py-1 rounded-full font-medium"
            :class="arrResult.error
              ? 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300'
              : 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300'">
            {{ arrResult.error ? arrResult.error : `${arrResult.count} importés` }}
          </span>
          <button
            @click="importArr"
            :disabled="loadingArr"
            class="flex items-center gap-1.5 px-4 py-2 rounded-lg text-sm font-medium bg-teal-600 hover:bg-teal-700 text-white disabled:opacity-50 disabled:cursor-not-allowed transition-all"
          >
            <Loader2 v-if="loadingArr" :size="14" class="animate-spin" />
            <Upload v-else :size="14" />
            {{ loadingArr ? 'Import...' : 'Importer' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Note INSERT OR IGNORE -->
    <p class="text-xs text-gray-400 dark:text-gray-500 flex items-start gap-1.5">
      <Info :size="12" class="mt-0.5 shrink-0" />
      Les entrées déjà existantes sont ignorées (INSERT OR IGNORE). 
      Tu peux relancer l'import sans risque de doublons.
    </p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { MapPin, Upload, Loader2, Info } from 'lucide-vue-next'
import { ImportMunicipalitesCSV, ImportArrondissementsCSV } from '../../../wailsjs/go/main/App'

const loadingMun = ref(false)
const loadingArr = ref(false)
const munResult  = ref(null)
const arrResult  = ref(null)

const importMun = async () => {
  loadingMun.value = true
  munResult.value  = null
  try {
    const count = await ImportMunicipalitesCSV('data/MUN.csv')
    munResult.value = { count }
  } catch (e) {
    munResult.value = { error: e.toString() }
  } finally {
    loadingMun.value = false
  }
}

const importArr = async () => {
  loadingArr.value = true
  arrResult.value  = null
  try {
    const count = await ImportArrondissementsCSV('data/ARR.csv')
    arrResult.value = { count }
  } catch (e) {
    arrResult.value = { error: e.toString() }
  } finally {
    loadingArr.value = false
  }
}
</script>