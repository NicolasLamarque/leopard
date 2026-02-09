<!-- frontend/src/views/AppelsPage.vue -->

<template>
  <div class="h-screen flex flex-col bg-stone-50 dark:bg-stone-950">
    
    <!-- Header -->
    <AppelsHeader
      :stats="stats"
      :is-creating="isCreating || isEditing"
      @start-new="startNew"
      @cancel-new="cancelEdit"
    />

    <!-- Conteneur principal -->
    <div class="flex-1 flex overflow-hidden">
      
      <!-- Sidebar -->
      <AppelsSidebar
        :appels="appels"
        :selected-appel-id="selectedAppelId"
        @select-appel="selectAppel"
      />

      <!-- Zone principale -->
      <div class="flex-1 flex flex-col">
        
        <!-- Vue par défaut (aucun appel sélectionné) -->
        <div v-if="!isCreating && !selectedAppel" class="flex-1 flex items-center justify-center bg-white dark:bg-stone-900">
          <div class="text-center">
            <div class="w-20 h-20 mx-auto mb-6 rounded-full bg-gradient-to-br from-pink-100 to-rose-100 dark:from-pink-900/30 dark:to-rose-900/30 flex items-center justify-center">
              <Phone :size="32" class="text-pink-500/30 dark:text-pink-400/45" />
            </div>
            <h3 class="text-xl font-semibold text-stone-700 dark:text-stone-300 mb-2">
              Gestion des appels
            </h3>
            <p class="text-sm text-stone-500 dark:text-stone-400 mb-6">
              Sélectionnez un appel ou créez-en un nouveau
            </p>
            <button
              @click="startNew"
              class="px-5 py-2.5 bg-gradient-to-r from-pink-500/85 to-rose-500/45 hover:from-pink-600/30 hover:to-rose-600/55 text-white rounded-xl text-sm font-medium shadow-md shadow-pink-500/20 hover:shadow-lg hover:shadow-pink-500/30 transition-all duration-200 flex items-center gap-2 mx-auto"
            >
              <Plus :size="18" />
              Nouvel appel
            </button>
          </div>
        </div>

        <!-- Éditeur -->
        <AppelsEditor
          v-else
          :appel="selectedAppel"
          :clients="clients"
          :users="users"
          @submit="saveAppel"
        />

        <!-- Footer -->
        <AppelsFooter
          v-if="isCreating || isEditing"
          :appel="selectedAppel"
          :is-saving="isSaving"
          :can-save="canSave"
          @save="submitForm"
          @cancel="cancelEdit"
          @delete="deleteAppel"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Phone, Plus } from 'lucide-vue-next'
import AppelsHeader from '../components/appels/AppelHeader.vue'
import AppelsSidebar from '../components/appels/AppelSidebar.vue'
import AppelsEditor from '../components/Appels/Appeleditor.vue'
import AppelsFooter from '../components/Appels/Appelfooter.vue'
import { useAppels } from '../composables/useAppels'

// Composable
const {
  appels,
  stats,
  loading,
  error,
  loadAppels,
  loadStats,
  loadAppelById,
  createAppel
} = useAppels()

// État local
const selectedAppelId = ref(null)
const selectedAppel = ref(null)
const isCreating = ref(false)
const isSaving = ref(false)
const clients = ref([])
const users = ref([])
const formRef = ref(null)

const isEditing = computed(() => !!selectedAppel.value?.id)
const canSave = computed(() => !isSaving.value)

// ========== ACTIONS ==========

const startNew = () => {
  selectedAppel.value = {
    appelant_nom: '',
    appelant_prenom: '',
    appelant_telephone: '',
    appelant_relation: 'lui_meme',
    client_id: null,
    prospect_nom: '',
    prospect_prenom: '',
    prospect_telephone: '',
    type_demande: 'evaluation',
    motif_appel: '',
    priorite: 'normale',
    statut: 'nouveau',
    notes_internes: '',
    rdv_date: '',
    rdv_heure: '',
    rdv_lieu: '',
    assigne_a: null
  }
  selectedAppelId.value = null
  isCreating.value = true
}

const selectAppel = async (id) => {
  selectedAppelId.value = id
  isCreating.value = false
  selectedAppel.value = await loadAppelById(id)
}

const cancelEdit = () => {
  if (isCreating.value || isEditing.value) {
    isCreating.value = false
    selectedAppel.value = null
    selectedAppelId.value = null
  } else {
    router.push('/app') 
  }
}

const handleFormSubmit = async (formData) => {
  isSaving.value = true
  try {
    const appelData = {
      date_appel: new Date().toISOString().split('T')[0],
      heure_appel: new Date().toLocaleTimeString('fr-CA', { hour: '2-digit', minute: '2-digit' }),
      ...formData
    }
    await createAppel(appelData)
    await loadAppels()
    await loadStats()
    isCreating.value = false
    selectedAppel.value = null
  } catch (err) {
    console.error('Erreur sauvegarde:', err)
  } finally {
    isSaving.value = false
  }
}

const deleteAppel = async () => {
  if (!selectedAppel.value?.id) return
  if (!confirm('Êtes-vous sûr de vouloir supprimer cet appel ?')) return
  
  try {
    // await deleteAppel(selectedAppel.value.id)
    await loadAppels()
    await loadStats()
    cancelEdit()
  } catch (err) {
    console.error('Erreur suppression:', err)
  }
}

// ========== CHARGEMENT ==========

const loadClients = async () => {
  try {
    // TODO: Importer GetClients depuis wailsjs
    // const { GetClients } = await import('../../wailsjs/go/main/App')
    // clients.value = await GetClients()
    clients.value = [] // Temporaire
  } catch (err) {
    console.error('Erreur chargement clients:', err)
  }
}

const loadUsers = async () => {
  try {
    // TODO: Implémenter GetUsers dans app.go si nécessaire
    users.value = [] // Temporaire
  } catch (err) {
    console.error('Erreur chargement users:', err)
  }
}

onMounted(async () => {
  await Promise.all([
    loadAppels(),
    loadStats(),
    loadClients(),
    loadUsers()
  ])
})
</script>