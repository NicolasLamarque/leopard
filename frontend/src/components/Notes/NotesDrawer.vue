<template>
  <Transition name="modal-fade">
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-md" @click="handleClose"></div>

      <div class="relative bg-white dark:bg-gray-950 w-full max-w-7xl h-[92vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-gray-200 dark:border-gray-800">
        
        <!-- Header -->
        <NotesHeader
          :client="client"
          :selected-note="selectedNote"
          :selected-notes="selectedNotes"
          :is-creating="isCreating"
          :is-draft="isDraft"
          :is-exporting="isExporting"
          @start-new="startNewNote"
          @cancel-creation="cancelCreation"
          @export-pdf="exportToPDF"
          @close="handleClose"
        />

        <div class="flex-1 flex overflow-hidden">
          
          <!-- Sidebar -->
          <NotesSidebar
            :notes="notes"
            :selected-note="selectedNote"
            :selected-notes="selectedNotes"
            :search-query="searchQuery"
            :filter-type="filterType"
            @view-note="viewNote"
            @toggle-note="toggleNoteSelection"
            @toggle-select-all="toggleSelectAll"
            @filter-type="(type) => filterType = type"
            @update:search-query="(query) => searchQuery = query"
          />

          <!-- Zone principale -->
          <div class="flex-1 flex flex-col bg-white dark:bg-gray-950">
            <!-- MODE CRÉATION / ÉDITION -->
            <NotesEditor
              v-if="isCreating"
              v-model="formData"
              :errors="errors"
              :is-saving="isSaving"
              :is-locked="false"
              :available-notes="verrouilledNotes"
              @save-draft="handleSaveDraft"
              @finalize="handleFinalize"
            />

            <!-- MODE VISUALISATION -->
            <NotesViewer
              v-else-if="selectedNote"
              :note="selectedNote"
              @create-correction="createCorrection"
              @create-addendum="createAddendum"
              @export-single="exportSingleNote"
              @view-linked-note="viewLinkedNote"
            />

            <!-- MODE VIDE -->
            <div v-else class="flex-1 flex items-center justify-center">
              <div class="text-center max-w-md">
                <div class="p-6 bg-slate-100 dark:bg-slate-900/30 rounded-full inline-flex mb-4">
                  <ClipboardList :size="64" class="text-slate-400 dark:text-slate-600" />
                </div>
                <h3 class="text-xl font-bold text-gray-700 dark:text-gray-300 mb-2">
                  Aucune note sélectionnée
                </h3>
                <p class="text-gray-500 dark:text-gray-400 mb-6">
                  Sélectionnez une note dans la liste ou créez-en une nouvelle
                </p>
                <button
                  @click="startNewNote"
                  class="px-6 py-3 bg-gradient-to-r from-teal-600 to-teal-700 hover:from-teal-700 hover:to-teal-800 text-white rounded-xl font-bold shadow-lg shadow-teal-500/30 transition-all flex items-center gap-2 mx-auto"
                >
                  <Plus :size="20" />
                  Créer une nouvelle note
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer avec brouillons -->
        <NotesFooter
          :brouillons="brouillons"
          @view-draft="viewDraft"
          @delete-draft="deleteDraft"
          @delete-all-drafts="deleteAllDrafts"
          @finalize-draft="finalizeDraft"
        />

       
      </div>
    </div>
  </Transition>

  <!-- Modale de confirmation -->
  <ConfirmModal />
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useToast } from '../../composables/useToast.js'
import { ClipboardList, Plus } from 'lucide-vue-next'
import NotesHeader from './NotesHeader.vue'
import NotesSidebar from './NotesSidebar.vue'
import NotesEditor from './NotesEditor.vue'
import NotesViewer from './NotesViewer.vue'
import NotesFooter from './NotesFooter.vue'
import { useNotesPDF } from '../../composables/useNotePDF.js'
// Centralisation de tous les imports Wails du module App
import { 
  ExportNotesToPDF,
  CreateNote, 
  UpdateNoteDraft, 
  DeleteNote, 
  LockNote, 
  GetClientNotes, 
  GetNoteByID,
  CheckNoteLieeExists,
  ExportNotesToPDF as ExportPDFBackend 
} from '../../../wailsjs/go/main/App'

import { useConfirm } from '../../composables/useConfirm.js'
const { confirm } = useConfirm()

const props = defineProps({
  isOpen: Boolean,
  client: Object
})

const emit = defineEmits(['close'])
// 2. INITIALISE LE COMPOSABLE ICI (Pas dans la fonction !)
const { generateSingleNotePDF, generateNotesWithSummaryPDF, loadLogoAsBase64 } = useNotesPDF()
const { toast } = useToast()
// États principaux
const notes = ref([])
const selectedNote = ref(null)
const selectedNotes = ref([])

const isCreating = ref(false)
const isSaving = ref(false)
const isExporting = ref(false)
const errors = ref({})
const searchQuery = ref('')
const filterType = ref(null)

// Formulaire de création/édition
const formData = ref({
  titre: '',
  type_note: 'Suivi',
  contenu: '',
  date_intervention: new Date().toISOString().split('T')[0],
  mode_intervention: 'En personne',
  type_intervention: '',
  note_liee_id: null,
  type_lien: null
})

// Computed
const brouillons = computed(() => notes.value.filter(n => !n.verrouille))
const verrouilledNotes = computed(() => notes.value.filter(n => n.verrouille))
const isDraft = computed(() => !selectedNote.value?.verrouille)

// --- GESTION DES NOTES ---

const loadNotes = async () => {
  try {
    console.log('🔍 Chargement notes pour client:', props.client?.id)
    const result = await GetClientNotes(props.client?.id)
    notes.value = result || []
    console.log('✅ Notes chargées:', notes.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement notes:', err)
    notes.value = []
    toast.error('Erreur de chargement', err?.toString())
  }
}

const viewNote = async (noteListItem) => {
  try {
    console.log('🔍 Chargement note complète:', noteListItem.id)
    const fullNote = await GetNoteByID(noteListItem.id)
    selectedNote.value = fullNote
    isCreating.value = false
    console.log('✅ Note chargée:', fullNote.titre)
  } catch (err) {
    console.error('❌ Erreur chargement note:', err)
    toast.error('Impossible de charger la note')
  }
}

const viewLinkedNote = async (noteId) => {
  try {
    const fullNote = await GetNoteByID(noteId)
    selectedNote.value = fullNote
    isCreating.value = false
  } catch (err) {
    console.error('❌ Erreur chargement note liée:', err)
    toast.error('Impossible de charger la note liée')
  }
}

// --- CRÉATION / ÉDITION ---

const startNewNote = () => {
  formData.value = {
    titre: '',
    type_note: 'Suivi',
    contenu: getDefaultTemplate(),
    date_intervention: new Date().toISOString().split('T')[0],
    mode_intervention: 'En personne',
    type_intervention: '',
    note_liee_id: null,
    type_lien: null
  }
  errors.value = {}
  selectedNote.value = null
  isCreating.value = true
}

const createCorrection = async () => {
  if (!selectedNote.value) return

  if (!selectedNote.value.verrouille) {
    toast.warning('Note non finalisée', 'Finalisez ce brouillon avant de créer une correction.')
    return
  }

  // Vérifier qu'une correction n'existe pas déjà
  const exists = await CheckNoteLieeExists(selectedNote.value.id, 'Correction')
  if (exists) {
    toast.warning('Correction déjà existante', 'Une correction existe déjà pour cette note.')
    return
  }
  
  formData.value = {
    titre: `Correction - ${selectedNote.value.titre}`,
    type_note: 'Correction',
    contenu: `CORRECTION APPORTÉE À LA NOTE #${selectedNote.value.id}\n\n` +
             `Note originale: ${selectedNote.value.titre}\n` +
             `Date de la note originale: ${new Date(selectedNote.value.date_note).toLocaleDateString('fr-CA')}\n\n` +
             `ÉLÉMENTS À CORRIGER\n` +
             `[Préciser les éléments inexacts ou incomplets de la note originale]\n\n` +
             `CORRECTIONS APPORTÉES\n` +
             `[Décrire les corrections apportées]\n\n` +
             `JUSTIFICATION\n` +
             `[Expliquer la raison de la correction]`,
    date_intervention: new Date().toISOString().split('T')[0],
    mode_intervention: selectedNote.value.mode_intervention || 'En personne',
    type_intervention: '',
    note_liee_id: selectedNote.value.id,
    type_lien: 'Correction'
  }
  errors.value = {}
  selectedNote.value = null
  isCreating.value = true
}

const createAddendum = async () => {
  if (!selectedNote.value) return

  if (!selectedNote.value.verrouille) {
    toast.warning('Note non finalisée', 'Finalisez ce brouillon avant de créer un addendum.')
    return
  }

  // Vérifier qu'un addendum n'existe pas déjà
  const exists = await CheckNoteLieeExists(selectedNote.value.id, 'Addendum')
  if (exists) {
    toast.warning('Addendum déjà existant', 'Un addendum existe déjà pour cette note.')
    return
  }
  
  formData.value = {
    titre: `Addendum - ${selectedNote.value.titre}`,
    type_note: 'Addendum',
    contenu: `ADDENDUM À LA NOTE #${selectedNote.value.id}\n\n` +
             `Note originale: ${selectedNote.value.titre}\n` +
             `Date de la note originale: ${new Date(selectedNote.value.date_note).toLocaleDateString('fr-CA')}\n\n` +
             `INFORMATIONS COMPLÉMENTAIRES\n` +
             `[Ajouter les informations supplémentaires pertinentes]\n\n` +
             `CONTEXTE DE L'AJOUT\n` +
             `[Expliquer pourquoi ces informations sont ajoutées maintenant]`,
    date_intervention: new Date().toISOString().split('T')[0],
    mode_intervention: selectedNote.value.mode_intervention || 'En personne',
    type_intervention: '',
    note_liee_id: selectedNote.value.id,
    type_lien: 'Addendum'
  }
  errors.value = {}
  selectedNote.value = null
  isCreating.value = true
}

const cancelCreation = async () => {
  if (formData.value.contenu && formData.value.contenu !== getDefaultTemplate()) {
    const ok = await confirm({
      title: 'Annuler les modifications ?',
      message: 'Les modifications non sauvegardées seront perdues.',
      level: 'warning',
      confirmLabel: 'Oui, annuler'
    })
    if (!ok) return
  }
  toast.info('Édition annulée')
  isCreating.value = false
  selectedNote.value = null
}

// --- VALIDATION ---
const validateForm = () => {
  errors.value = {}
  let isValid = true

  if (!formData.value.type_note) {
    errors.value.type_note = 'Le type de note est requis'
    isValid = false
  }

  if (!formData.value.titre || formData.value.titre.trim().length === 0) {
    errors.value.titre = 'Le sujet de la note est requis'
    isValid = false
  }

  if (!formData.value.contenu || formData.value.contenu.trim().length < 20) {
    errors.value.contenu = 'Le contenu doit contenir au moins 20 caractères'
    isValid = false
  }

  return isValid
}

// --- SAUVEGARDE / FINALISATION ---

const handleSaveDraft = async () => {
  if (!validateForm()) return

  isSaving.value = true
  try {
    // 🆕 LOGIQUE UPDATE vs CREATE
    if (selectedNote.value?.id && !selectedNote.value.verrouille) {
      // CAS 1: On édite un brouillon existant → UPDATE
      console.log('📝 Mise à jour du brouillon #', selectedNote.value.id)
      
      await UpdateNoteDraft(selectedNote.value.id, {
        client_id: props.client?.id,
        date_intervention: formData.value.date_intervention,
        mode_intervention: formData.value.mode_intervention,
        type_intervention: formData.value.type_intervention,
        type_note: formData.value.type_note,
        titre: formData.value.titre,
        contenu: formData.value.contenu,
        note_liee_id: formData.value.note_liee_id,
        type_lien: formData.value.type_lien,
        verrouille: false
      })

      // Recharger la note mise à jour
      const updatedNote = await GetNoteByID(selectedNote.value.id)
      selectedNote.value = updatedNote
      
      console.log('✅ Brouillon mis à jour avec succès')
    } else {
      // CAS 2: Nouvelle note → CREATE
      console.log('✨ Création d\'une nouvelle note')
      
      const noteId = await CreateNote({
        client_id: props.client?.id,
        date_intervention: formData.value.date_intervention,
        mode_intervention: formData.value.mode_intervention,
        type_intervention: formData.value.type_intervention,
        type_note: formData.value.type_note,
        titre: formData.value.titre,
        contenu: formData.value.contenu,
        note_liee_id: formData.value.note_liee_id,
        type_lien: formData.value.type_lien,
        verrouille: false
      })

      // Charger la note nouvellement créée
      const newNote = await GetNoteByID(noteId)
      selectedNote.value = newNote
      
      console.log('✅ Note créée avec ID:', noteId)
    }

    // Rafraîchir la liste
    await loadNotes()
    
    toast.success('Brouillon sauvegardé', formData.value.titre || 'Note sans titre')
    
  } catch (err) {
    console.error('❌ Erreur sauvegarde:', err)
    toast.error('Erreur de sauvegarde', err?.toString())
  } finally {
    isSaving.value = false
  }
}

const handleFinalize = async () => {
  if (!validateForm()) return
  const ok = await confirm({
    title: 'Finaliser et signer ?',
    message: 'La note sera verrouillée définitivement.',
    detail: 'Cette action est irréversible.',
    level: 'danger',
    confirmLabel: 'Finaliser et signer'
  })
  if (!ok) return
  isSaving.value = true
  try {
    let noteIdToLock

    // Si on édite un brouillon existant
    if (selectedNote.value?.id && !selectedNote.value.verrouille) {
      // D'abord mettre à jour les données
      await UpdateNoteDraft(selectedNote.value.id, {
        client_id: props.client?.id,
        date_intervention: formData.value.date_intervention,
        mode_intervention: formData.value.mode_intervention,
        type_intervention: formData.value.type_intervention,
        type_note: formData.value.type_note,
        titre: formData.value.titre,
        contenu: formData.value.contenu,
        note_liee_id: formData.value.note_liee_id,
        type_lien: formData.value.type_lien,
        verrouille: false
      })
      noteIdToLock = selectedNote.value.id
    } else {
      // Créer une nouvelle note
      noteIdToLock = await CreateNote({
        client_id: props.client?.id,
        date_intervention: formData.value.date_intervention,
        mode_intervention: formData.value.mode_intervention,
        type_intervention: formData.value.type_intervention,
        type_note: formData.value.type_note,
        titre: formData.value.titre,
        contenu: formData.value.contenu,
        note_liee_id: formData.value.note_liee_id,
        type_lien: formData.value.type_lien,
        verrouille: false
      })
    }

    // Maintenant verrouiller la note
    await LockNote(noteIdToLock)
    
    await loadNotes()
    
    // Charger la note verrouillée pour l'afficher
    const lockedNote = await GetNoteByID(noteIdToLock)
    selectedNote.value = lockedNote
    
    isCreating.value = false
    const titreFinalise = formData.value.titre || 'Note sans titre'
    resetForm()
    toast.success('Note finalisée et signée', titreFinalise)
    
  } catch (err) {
    console.error('❌ Erreur finalisation:', err)
    toast.error('Erreur de finalisation', err?.toString())
  } finally {
    isSaving.value = false
  }
}

const finalizeDraft = async (draftId) => {
  const ok = await confirm({
    title: 'Finaliser ce brouillon ?',
    message: 'La note sera verrouillée définitivement.',
    detail: 'Cette action est irréversible.',
    level: 'danger',
    confirmLabel: 'Finaliser et signer'
  })
  if (!ok) return
  try {

    await LockNote(draftId)
    await loadNotes()
    
    toast.success('Note finalisée et signée')
  } catch (err) {
    console.error('Erreur finalisation brouillon:', err)
    toast.error('Erreur de finalisation', err?.toString())
  }
}

// --- BROUILLONS ---

const viewDraft = async (draft) => {
  try {
    console.log('🔍 Chargement du brouillon via ID:', draft.id)
    const fullNote = await GetNoteByID(draft.id)
    
    // Remplir le formulaire avec les données du brouillon
    formData.value = {
      titre: fullNote.titre || '',
      type_note: fullNote.type_note || 'Suivi',
      contenu: fullNote.contenu || '',
      date_intervention: fullNote.date_intervention 
        ? new Date(fullNote.date_intervention).toISOString().split('T')[0]
        : new Date().toISOString().split('T')[0],
      mode_intervention: fullNote.mode_intervention || 'En personne',
      type_intervention: fullNote.type_intervention || '',
      note_liee_id: fullNote.note_liee_id || null,
      type_lien: fullNote.type_lien || null
    }

    // Définir la note sélectionnée
    selectedNote.value = fullNote
    
    // Passer en mode édition
    isCreating.value = true
    
    console.log('✅ Brouillon chargé pour édition')
  } catch (err) {
    console.error('❌ Erreur chargement brouillon:', err)
    toast.error('Impossible de charger ce brouillon')
  }
}

const deleteDraft = async (draftId) => {
  const ok = await confirm({
    title: 'Supprimer ce brouillon ?',
    message: 'Ce brouillon sera supprimé définitivement.',
    level: 'danger',
    confirmLabel: 'Supprimer'
  })
  if (!ok) return
  try {
    await DeleteNote(draftId)
    await loadNotes()
    
    // Si on était en train d'éditer ce brouillon, fermer l'éditeur
    if (selectedNote.value?.id === draftId) {
      selectedNote.value = null
      isCreating.value = false
    }
    
    toast.success('Brouillon supprimé')
  } catch (err) {
    console.error('Erreur suppression brouillon:', err)
    toast.error('Erreur de suppression', err?.toString())
  }
}

const deleteAllDrafts = async () => {
  const ok = await confirm({
    title: 'Supprimer tous les brouillons ?',
    message: `${brouillons.value.length} note(s) seront supprimées définitivement.`,
    detail: 'Cette action est irréversible.',
    level: 'danger',
    confirmLabel: 'Tout supprimer'
  })
  if (!ok) return
  try {
    for (const draft of brouillons.value) {
      await DeleteNote(draft.id)
    }
    
    await loadNotes()
    selectedNote.value = null
    isCreating.value = false
    
    toast.success('Brouillons supprimés', `${brouillons.value.length} note(s) supprimée(s)`)
  } catch (err) {
    console.error('Erreur suppression brouillons:', err)
    toast.error('Erreur de suppression', err?.toString())
  }
}

// --- SÉLECTION MULTIPLE ---

const toggleNoteSelection = (noteId) => {
  const index = selectedNotes.value.indexOf(noteId)
  if (index > -1) {
    selectedNotes.value.splice(index, 1)
  } else {
    selectedNotes.value.push(noteId)
  }
}

const toggleSelectAll = () => {
  if (selectedNotes.value.length === notes.value.length) {
    selectedNotes.value = []
  } else {
    selectedNotes.value = notes.value.map(n => n.id)
  }
}

// --- EXPORT PDF ---

// Export d'une seule note (bouton dans NotesViewer)
const exportSingleNote = async () => {
  if (!selectedNote.value) return
  isExporting.value = true
  try {
    const logo = await loadLogoAsBase64()
    const doc = await generateSingleNotePDF(selectedNote.value, logo)
    const pdfBase64 = doc.output('datauristring')
    await ExportNotesToPDF(
      props.client.no_dossier_leopard,
      [selectedNote.value.id],
      pdfBase64
    )
    toast.success('PDF exporté', `${selectedNote.value.titre}`)
  } catch (err) {
    toast.error("Erreur d'exportation PDF", err?.toString())
  } finally {
    isExporting.value = false
  }
}

const exportToPDF = async () => {
  if (selectedNotes.value.length === 0) {
    toast.warning('Aucune note sélectionnée', 'Cochez au moins une note dans la liste.')
    return
  }

  isExporting.value = true
  try {
    const logo = await loadLogoAsBase64()
    
    // On crée la liste des notes à exporter SEULEMENT ICI
    // Et on utilise GetNoteByID pour avoir le contenu complet (le texte)
    console.log("📥 Récupération des contenus complets pour l'export...");
    const notesCompletes = await Promise.all(
      selectedNotes.value.map(id => GetNoteByID(id))
    )

    let doc
    if (notesCompletes.length === 1) {
      doc = await generateSingleNotePDF(notesCompletes[0], logo)
    } else {
      const leopardNo = props.client?.no_dossier_leopard || 'INCONNU'
      doc = await generateNotesWithSummaryPDF(notesCompletes, leopardNo, logo)
    }

    const pdfBase64 = doc.output('datauristring')
    
    // Envoi au backend Go
    const filePath = await ExportNotesToPDF(
      props.client.no_dossier_leopard, 
      selectedNotes.value, 
      pdfBase64
    )
    
    toast.success('PDF exporté avec succès', `${notesCompletes.length} note(s) — ${props.client?.prenom} ${props.client?.nom}`)
    
  } catch (err) {
    console.error("Erreur d'exportation:", err)
    toast.error('Erreur d\'exportation PDF', err?.toString())
  } finally {
    isExporting.value = false
  }
}

// --- UTILITAIRES ---

const getDefaultTemplate = () => {
  return `CONTEXTE D'INTERVENTION
[Décrire brièvement le contexte et la raison de l'intervention]

OBSERVATIONS ET DONNÉES RECUEILLIES
[Observations objectives, données factuelles, résultats d'évaluation]

ANALYSE CLINIQUE
[Interprétation professionnelle basée sur le raisonnement clinique]

INTERVENTIONS RÉALISÉES
[Actions posées durant la rencontre]

PLAN DE SUIVI
[Prochaines étapes, recommandations, objectifs]

CONCLUSION
[Synthèse de la rencontre et état de la situation]`
}

const resetForm = () => {
  formData.value = {
    titre: '', 
    type_note: 'Suivi',
    contenu: '',
    date_intervention: new Date().toISOString().split('T')[0],
    mode_intervention: 'En personne',
    type_intervention: '',
    note_liee_id: null,
    type_lien: null
  }
}

const handleClose = async () => {
  if (isCreating.value && formData.value.contenu) {
    const ok = await confirm({
      title: 'Fermer sans sauvegarder ?',
      message: 'Les modifications en cours seront perdues.',
      level: 'warning',
      confirmLabel: 'Fermer quand même'
    })
    if (!ok) return
  }
  isCreating.value = false
  selectedNote.value = null
  selectedNotes.value = []
  emit('close')
}

// --- WATCHERS ---

watch(() => props.isOpen, (newVal) => {
  if (newVal && props.client?.id) {
    loadNotes()
    selectedNotes.value = []
  }
})
</script>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active { 
  transition: opacity 0.3s ease; 
}

.modal-fade-enter-from,
.modal-fade-leave-to { 
  opacity: 0; 
}

::-webkit-scrollbar { 
  width: 6px; 
  height: 6px;
}

::-webkit-scrollbar-thumb { 
  background: linear-gradient(to bottom, #0d9488, #14b8a6);
  border-radius: 10px; 
}

::-webkit-scrollbar-track {
  background: transparent;
}
</style>