<!-- frontend/src/views/AppelsPage.vue -->

<template>
  <div class="h-screen flex flex-col bg-stone-50 dark:bg-stone-950">
    
    <!-- Header -->
    <AppelsHeader
      :stats="stats"
      :is-creating="isCreating || isEditing"
      @start-new="startNew"
      @cancel-new="cancelEdit"
      @generate-report="handleReport"
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
        
        <!-- Vue par défaut -->
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
          ref="editorRef"
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
          @delete="handleDelete"
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
import { useAppelReports } from '../composables/useAppelReports'


// On extrait les fonctions de ton composable
const { generateWeekly, generateMonthly } = useAppelReports()

// Voici la fonction "handleReport" qui manquait
const handleReport = (type) => {
  console.log("Demande de rapport reçue :", type)
  
  // 'appels' est la variable qui contient ta liste d'appels dans AppelPage.vue
  if (type === 'weekly') {
    generateWeekly(appels.value)
  } else if (type === 'monthly') {
    generateMonthly(appels.value)
  }
}

// Pour le rapport par client (à déclencher depuis l'éditeur ou la liste)
const reportForCurrentClient = () => {
    if (selectedAppel.value?.client_id) {
        const name = `${selectedAppel.value.prospect_prenom} ${selectedAppel.value.prospect_nom}`;
        generateByClient(appels.value, selectedAppel.value.client_id, name);
    }
}

const {
  appels,
  stats,
  loading,
  error,
  loadAppels,
  loadStats,
  loadAppelById,
  createAppel,
  updateAppel,
  deleteAppel
} = useAppels()

const selectedAppelId = ref(null)
const selectedAppel = ref(null)
const isCreating = ref(false)
const isSaving = ref(false)
const clients = ref([])
const users = ref([])
const editorRef = ref(null)

const isEditing = computed(() => !!selectedAppel.value?.id)
const canSave = computed(() => !isSaving.value)

// ── ACTIONS ──────────────────────────────────────────────────────────────────

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
  isCreating.value = false
  selectedAppel.value = null
  selectedAppelId.value = null
}

/**
 * Reçoit les données validées du FormKit (@submit de AppelsEditor)
 * C'est ici que tout s'enregistre réellement
 */
const saveAppel = async (formData) => {
  isSaving.value = true
  try {
    const appelData = {
      date_appel: new Date().toISOString().split('T')[0],
      heure_appel: new Date().toLocaleTimeString('fr-CA', { hour: '2-digit', minute: '2-digit' }),
      ...formData,
      // S'assurer que client_id est un int ou null
      client_id: formData.client_id ? parseInt(formData.client_id) : null,
      assigne_a: formData.assigne_a ? parseInt(formData.assigne_a) : null,
    }

    if (selectedAppel.value?.id) {
      // Mise à jour
      await updateAppel(selectedAppel.value.id, appelData)
    } else {
      // Création
      await createAppel(appelData)
    }

    isCreating.value = false
    selectedAppel.value = null
    selectedAppelId.value = null
  } catch (err) {
    console.error('❌ Erreur sauvegarde appel:', err)
  } finally {
    isSaving.value = false
  }
}

/**
 * Déclenche la soumission du FormKit depuis le bouton Sauvegarder du footer
 * FormKit gère la validation avant d'émettre @submit
 */
const submitForm = () => {
  const formEl = document.querySelector('.formkit-form')
  if (formEl) {
    formEl.dispatchEvent(new Event('submit', { bubbles: true, cancelable: true }))
  }
}

const handleDelete = async () => {
  if (!selectedAppel.value?.id) return
  if (!confirm('Voulez-vous vraiment supprimer cet appel ?')) return

  try {
    await deleteAppel(selectedAppel.value.id)
    cancelEdit()
  } catch (err) {
    console.error('❌ Erreur suppression:', err)
  }
}

// ── CHARGEMENT ────────────────────────────────────────────────────────────────

const loadClients = async () => {
  try {
    const { GetClients } = await import('../../wailsjs/go/main/App')
    clients.value = await GetClients() || []
  } catch (err) {
    console.error('Erreur chargement clients:', err)
    clients.value = []
  }
}

const loadUsers = async () => {
  try {
    // On essaie, mais on ne laisse pas l'erreur tuer la page
    const { GetUsers } = await import('../../wailsjs/go/main/App')
    if (typeof GetUsers === 'function') {
        users.value = await GetUsers() || []
    }
  } catch (err) {
    // On simule un utilisateur par défaut pour que l'interface reste vivante
    users.value = [{ id: 1, full_name: "Intervenant (Défaut)" }]
    console.warn("GetUsers absent du backend. Interface débloquée avec user générique.")
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