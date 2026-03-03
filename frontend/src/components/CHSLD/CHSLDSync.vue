<template>
  <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm p-6">
    
    <!-- En-tête -->
    <div class="flex items-center gap-3 mb-6">
      <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg">
        <RefreshCw :size="20" class="text-blue-600 dark:text-blue-400" />
      </div>
      <div>
        <h3 class="font-semibold text-gray-900 dark:text-white">Synchronisation MSSS</h3>
        <p class="text-sm text-gray-500 dark:text-gray-400">
          Importer les CHSLD depuis le répertoire officiel du MSSS
        </p>
      </div>
    </div>

    <!-- Compteur actuel -->
    <div class="flex items-center gap-4 mb-6 p-4 bg-gray-50 dark:bg-gray-700/50 rounded-lg">
      <Database :size="18" class="text-gray-400 shrink-0" />
      <div>
        <span class="text-sm text-gray-500 dark:text-gray-400">CHSLD en base : </span>
        <span class="font-bold text-gray-900 dark:text-white">{{ chsldCount }}</span>
      </div>
      <button
        @click="refreshCount"
        class="ml-auto p-1 text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 transition-colors"
        title="Rafraîchir"
      >
        <RefreshCw :size="14" :class="{ 'animate-spin': loadingCount }" />
      </button>
    </div>

    <!-- Zone de progression (visible uniquement pendant le sync) -->
    <Transition name="slide-down">
      <div v-if="syncing || lastResult" class="mb-6">
        
        <!-- Barre de progression -->
        <div v-if="syncing" class="mb-4">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-gray-700 dark:text-gray-300">
              {{ progress.region || 'Initialisation...' }}
            </span>
            <span class="text-sm text-gray-500 dark:text-gray-400">
              {{ progress.current }}/{{ progress.total }}
            </span>
          </div>
          
          <!-- Barre -->
          <div class="h-2 bg-gray-200 dark:bg-gray-600 rounded-full overflow-hidden">
            <div
              class="h-full bg-blue-500 rounded-full transition-all duration-300"
              :style="{ width: progressPercent + '%' }"
            />
          </div>

          <!-- Message courant -->
          <p class="mt-2 text-xs text-gray-500 dark:text-gray-400 truncate">
            {{ progress.message }}
          </p>

          <!-- Compteurs live -->
          <div class="flex gap-4 mt-3">
            <span class="text-xs text-green-600 dark:text-green-400">
              ✅ {{ progress.inserted }} insérés
            </span>
            <span class="text-xs text-gray-400">
              ⏭️ {{ progress.skipped }} déjà présents
            </span>
            <span v-if="progress.errors > 0" class="text-xs text-red-500">
              ❌ {{ progress.errors }} erreurs
            </span>
          </div>
        </div>

        <!-- Résultat final -->
        <div
          v-if="lastResult && !syncing"
          class="p-4 rounded-lg"
          :class="lastResult.errors > 0
            ? 'bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-700'
            : 'bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-700'"
        >
          <div class="flex items-start gap-3">
            <CheckCircle
              v-if="lastResult.errors === 0"
              :size="18"
              class="text-green-600 dark:text-green-400 shrink-0 mt-0.5"
            />
            <AlertTriangle
              v-else
              :size="18"
              class="text-yellow-600 dark:text-yellow-400 shrink-0 mt-0.5"
            />
            <div class="text-sm">
              <p class="font-medium text-gray-900 dark:text-white mb-1">Sync terminé</p>
              <p class="text-gray-600 dark:text-gray-300">{{ lastResult.message }}</p>
            </div>
          </div>
        </div>

      </div>
    </Transition>

    <!-- Log console (déroulant) -->
    <Transition name="slide-down">
      <div v-if="syncing && logs.length > 0" class="mb-6">
        <div
          ref="logBox"
          class="h-32 overflow-y-auto bg-gray-900 rounded-lg p-3 font-mono text-xs text-gray-300 space-y-0.5"
        >
          <div v-for="(log, i) in logs" :key="i" class="leading-relaxed">
            <span class="text-gray-500">{{ log.time }}</span>
            <span :class="log.color"> {{ log.text }}</span>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Boutons -->
    <div class="flex items-center gap-3">
      
      <!-- Bouton principal Sync -->
      <button
        @click="startSync"
        :disabled="syncing"
        class="flex items-center gap-2 px-5 py-2.5 rounded-lg font-medium transition-all"
        :class="syncing
          ? 'bg-gray-100 dark:bg-gray-700 text-gray-400 cursor-not-allowed'
          : 'bg-blue-600 hover:bg-blue-700 text-white shadow-sm hover:shadow-md'"
      >
        <RefreshCw :size="16" :class="{ 'animate-spin': syncing }" />
        {{ syncing ? 'Synchronisation...' : (chsldCount > 0 ? 'Mettre à jour' : 'Importer depuis MSSS') }}
      </button>

      <!-- Bouton annuler (placeholder — le crawl Go ne supporte pas encore cancel) -->
      <button
        v-if="syncing"
        @click="cancelSync"
        class="px-4 py-2.5 rounded-lg text-sm text-gray-600 dark:text-gray-300 border border-gray-300 dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
      >
        Annuler
      </button>

      <!-- Effacer le log résultat -->
      <button
        v-if="lastResult && !syncing"
        @click="lastResult = null; logs = []"
        class="px-4 py-2.5 rounded-lg text-sm text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors"
      >
        Effacer
      </button>

    </div>

    <!-- Avertissement réseau -->
    <p class="mt-4 text-xs text-gray-400 dark:text-gray-500 flex items-center gap-1.5">
      <Globe :size="12" />
      Requiert une connexion Internet — site MSSS : m02.pub.msss.rtss.qc.ca
    </p>

  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onUnmounted } from 'vue'
import { RefreshCw, Database, CheckCircle, AlertTriangle, Globe } from 'lucide-vue-next'
import { SyncCHSLDFromMSSS, GetCHSLDCount } from '../../../wailsjs/go/main/App'
import { EventsOn, EventsOff } from '../../../wailsjs/runtime/runtime'

// ─── État ───────────────────────────────────────────────────────────────────
const syncing      = ref(false)
const loadingCount = ref(false)
const chsldCount   = ref(0)
const lastResult   = ref(null)
const logs         = ref([])
const logBox       = ref(null)

const progress = ref({
  region:   '',
  current:  0,
  total:    18,
  message:  '',
  inserted: 0,
  skipped:  0,
  errors:   0,
})

// ─── Computed ────────────────────────────────────────────────────────────────
const progressPercent = computed(() => {
  if (!progress.value.total) return 0
  return Math.round((progress.value.current / progress.value.total) * 100)
})

// ─── Lifecycle ───────────────────────────────────────────────────────────────
onMounted(async () => {
  await refreshCount()

  // Écouter les événements de progression émis par Go (Wails v2)
  EventsOn('chsld:sync:progress', handleProgress)
})

onUnmounted(() => {
  EventsOff('chsld:sync:progress')
})

// ─── Handlers ────────────────────────────────────────────────────────────────
function handleProgress(p) {
  progress.value = { ...progress.value, ...p }

  if (p.message) {
    const now = new Date()
    const time = `${String(now.getHours()).padStart(2,'0')}:${String(now.getMinutes()).padStart(2,'0')}:${String(now.getSeconds()).padStart(2,'0')}`
    let color = 'text-gray-300'
    if (p.message.startsWith('✅')) color = 'text-green-400'
    if (p.message.startsWith('❌')) color = 'text-red-400'
    if (p.message.startsWith('⏭️')) color = 'text-gray-500'

    logs.value.push({ time, text: p.message, color })

    // Auto-scroll du log
    nextTick(() => {
      if (logBox.value) {
        logBox.value.scrollTop = logBox.value.scrollHeight
      }
    })
  }
}

async function startSync() {
  if (syncing.value) return

  syncing.value  = true
  lastResult.value = null
  logs.value     = []
  progress.value = { region: '', current: 0, total: 18, message: '', inserted: 0, skipped: 0, errors: 0 }

  try {
    // SyncCHSLDFromMSSS est bloquant côté Go mais les events arrivent en temps réel
    const result = await SyncCHSLDFromMSSS()
    lastResult.value = result
    await refreshCount()
  } catch (err) {
    console.error('Erreur sync MSSS:', err)
    lastResult.value = {
      inserted: 0,
      skipped:  0,
      errors:   1,
      message:  `Erreur : ${err}`,
    }
  } finally {
    syncing.value = false
  }
}

function cancelSync() {
  // TODO : implémenter annulation via context.WithCancel en Go
  // Pour l'instant on affiche juste un message
  console.warn('Annulation non encore supportée — le sync Go continue en arrière-plan')
  syncing.value = false
}

async function refreshCount() {
  loadingCount.value = true
  try {
    chsldCount.value = await GetCHSLDCount()
  } catch {
    chsldCount.value = 0
  } finally {
    loadingCount.value = false
  }
}
</script>

<style scoped>
.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease;
}
.slide-down-enter-from,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
</style>