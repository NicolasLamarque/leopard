// frontend/src/composables/useAppels.js
import { ref } from 'vue'
import { GetAllAppels, GetStatsAppels, CreateAppel, GetAppelByID } from '../../wailsjs/go/main/App'

const appels = ref([])
const stats = ref({ aujourdhui: 0, enAttente: 0, urgents: 0, rdvPlanifies: 0, total: 0 })
const loading = ref(false)
const error = ref(null)

export function useAppels() {
  
  const loadAppels = async () => {
    loading.value = true
    error.value = null
    try {
      appels.value = await GetAllAppels() || []
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const loadStats = async () => {
    try {
      stats.value = await GetStatsAppels() || { aujourdhui: 0, enAttente: 0, urgents: 0, rdvPlanifies: 0, total: 0 }
    } catch (err) {
      console.error(err)
    }
  }

  const loadAppelById = async (id) => {
    loading.value = true
    try {
      return await GetAppelByID(id)
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  const createAppel = async (appelData) => {
    loading.value = true
    try {
      await CreateAppel(appelData)
      await loadAppels()
      await loadStats()
    } catch (err) {
      error.value = err.message
      throw err
    } finally {
      loading.value = false
    }
  }

  const getStatutBadge = (statut) => {
    const badges = {
      'nouveau': { class: 'bg-sky-100 text-sky-700 dark:bg-sky-900/30 dark:text-sky-300', label: 'Nouveau' },
      'en_attente': { class: 'bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-300', label: 'En attente' },
      'rdv_planifie': { class: 'bg-purple-100 text-purple-700 dark:bg-purple-900/30 dark:text-purple-300', label: 'RDV planifié' },
      'complete': { class: 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300', label: 'Complété' },
      'refuse': { class: 'bg-gray-100 text-gray-700 dark:bg-gray-700/30 dark:text-gray-400', label: 'Refusé' }
    }
    return badges[statut] || badges['nouveau']
  }

  const getPrioriteBadge = (priorite) => {
    const badges = {
      'urgente': { class: 'bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-300', label: 'Urgent', icon: '🚨' },
      'haute': { class: 'bg-orange-100 text-orange-700 dark:bg-orange-900/30 dark:text-orange-300', label: 'Haute', icon: '⚡' },
      'normale': { class: 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300', label: 'Normale', icon: '📋' },
      'basse': { class: 'bg-gray-100 text-gray-700 dark:bg-gray-700/30 dark:text-gray-400', label: 'Basse', icon: '📝' }
    }
    return badges[priorite] || badges['normale']
  }

  const formatDate = (dateString) => {
    if (!dateString) return 'N/A'
    return new Intl.DateTimeFormat('fr-CA', { year: 'numeric', month: 'long', day: 'numeric' }).format(new Date(dateString))
  }

  return {
    appels, stats, loading, error,
    loadAppels, loadStats, loadAppelById, createAppel,
    getStatutBadge, getPrioriteBadge, formatDate
  }
}