<template>
  <Transition name="modal-fade">
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-md" @click="closeModal"></div>

      <div class="relative bg-white dark:bg-gray-950 w-full max-w-6xl h-[90vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-gray-200 dark:border-gray-800">
        
        <!-- Header -->
        <div class="px-6 py-4 border-b dark:border-gray-800 flex justify-between items-center bg-gray-50/50 dark:bg-gray-900/50">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-blue-600 rounded-lg text-white">
              <ClipboardList :size="20" />
            </div>
            <div>
              <h2 class="text-xl font-bold dark:text-white">Dossier Clinique</h2>
              <p class="text-xs text-gray-500 uppercase font-medium tracking-wider">{{ client?.prenom }} {{ client?.nom }}</p>
            </div>
          </div>
          <div class="flex items-center gap-4">
            <!-- Bouton Export PDF -->
            <button 
              v-if="selectedNotes.length > 0"
              @click="exportToPDF"
              :disabled="isExporting"
              class="bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white px-5 py-2 rounded-xl text-sm font-bold flex items-center gap-2 shadow-lg shadow-green-500/20 transition-all"
            >
              <FileDown :size="18" />
              <span v-if="!isExporting">Exporter PDF ({{ selectedNotes.length }})</span>
              <span v-else>Génération...</span>
            </button>

            <button v-if="!isCreating" @click="startNewNote" class="bg-blue-600 hover:bg-blue-700 text-white px-5 py-2 rounded-xl text-sm font-bold flex items-center gap-2 shadow-lg shadow-blue-500/20 transition-all">
              <Plus :size="18" /> Nouvelle Note
            </button>
            <button v-else @click="cancelCreation" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 text-sm font-bold">Voir historique</button>
            <div class="h-6 w-px bg-gray-200 dark:bg-gray-800 mx-2"></div>
            <button @click="closeModal" class="p-2 hover:bg-gray-100 dark:hover:bg-gray-800 dark:text-gray-400 rounded-full transition-colors"><X :size="22" /></button>
          </div>
        </div>

        <div class="flex-1 flex overflow-hidden">
          
          <!-- Sidebar - Liste des notes avec checkboxes -->
          <div class="w-80 border-r dark:border-gray-800 bg-gray-50/30 dark:bg-gray-900/20 flex flex-col">
            <div class="p-4 border-b dark:border-gray-800 bg-gray-50/50 dark:bg-gray-900/50 flex items-center justify-between">
              <span class="text-[10px] font-bold text-gray-400 uppercase tracking-widest">Historique des notes</span>
              <button 
                v-if="notes.length > 0"
                @click="toggleSelectAll"
                class="text-xs text-blue-600 dark:text-blue-400 hover:underline font-medium"
              >
                {{ selectedNotes.length === notes.length ? 'Désélectionner' : 'Tout sélectionner' }}
              </button>
            </div>
            <div class="flex-1 overflow-y-auto p-4 space-y-3">
              <div v-for="note in notes" :key="note.id" 
                class="p-3 rounded-xl border dark:border-gray-800 bg-white dark:bg-gray-900 hover:ring-2 hover:ring-blue-500 cursor-pointer transition-all group">
                
                <div class="flex items-start gap-3">
                  <!-- Checkbox -->
                  <input 
                    type="checkbox"
                    :checked="selectedNotes.includes(note.id)"
                    @change="toggleNoteSelection(note.id)"
                    @click.stop
                    class="mt-1 w-4 h-4 text-blue-600 rounded border-gray-300 focus:ring-blue-500 cursor-pointer"
                  />
                  
                  <!-- Contenu note -->
                  <div @click="viewNote(note)" class="flex-1 min-w-0">
                    <div class="flex justify-between text-[10px] mb-1">
                      <span class="font-bold text-blue-600 uppercase">{{ note.type_note }}</span>
                      <span class="text-gray-400">{{ formatDate(note.date_note) }}</span>
                    </div>
                    <p class="text-sm font-bold text-gray-800 dark:text-gray-200 truncate">{{ note.sujet }}</p>
                    <div class="mt-2 flex items-center justify-between">
                      <span class="text-[9px] text-gray-400 italic">Signé: {{ note.signature_nom.split(',')[0] }}</span>
                      <Lock :size="12" class="text-green-500" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Zone principale - Formulaire ou Vue détail -->
          <div class="flex-1 flex flex-col bg-white dark:bg-gray-950">
            
            <!-- MODE CRÉATION -->
            <div v-if="isCreating" class="flex-1 flex flex-col overflow-hidden">
              
              <!-- Métadonnées de la note -->
              <div class="p-6 space-y-4 border-b dark:border-gray-800 bg-gray-50/30 dark:bg-gray-900/30">
                <div class="grid grid-cols-3 gap-4">
                  <!-- Type de note -->
                  <div class="space-y-2">
                    <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Type de note *</label>
                    <select 
                      v-model="formData.type_note" 
                      :class="[
                        'w-full bg-white dark:bg-gray-900 border rounded-xl p-3 text-sm dark:text-white outline-none transition-all',
                        errors.type_note ? 'border-red-500 ring-2 ring-red-200 dark:ring-red-900' : 'border-gray-300 dark:border-gray-800 focus:ring-2 focus:ring-blue-500'
                      ]"
                    >
                      <option value="">Sélectionner...</option>
                      <option value="Suivi">Suivi</option>
                      <option value="Ouverture">Ouverture</option>
                      <option value="Plan">Plan d'intervention</option>
                      <option value="Fermeture">Fermeture</option>
                      <option value="Incident">Incident</option>
                    </select>
                    <span v-if="errors.type_note" class="text-xs text-red-600 dark:text-red-400">{{ errors.type_note }}</span>
                  </div>

                  <!-- Date d'intervention -->
                  <div class="space-y-2">
                    <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Date intervention</label>
                    <input 
                      type="date" 
                      v-model="formData.date_intervention" 
                      class="w-full bg-white dark:bg-gray-900 border border-gray-300 dark:border-gray-800 rounded-xl p-3 text-sm dark:text-white outline-none focus:ring-2 focus:ring-blue-500"
                    >
                  </div>

                  <!-- Mode d'intervention -->
                  <div class="space-y-2">
                    <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Mode</label>
                    <select 
                      v-model="formData.mode_intervention" 
                      class="w-full bg-white dark:bg-gray-900 border border-gray-300 dark:border-gray-800 rounded-xl p-3 text-sm dark:text-white outline-none focus:ring-2 focus:ring-blue-500"
                    >
                      <option value="En personne">En personne</option>
                      <option value="Téléphone">Téléphone</option>
                      <option value="Visioconférence">Visioconférence</option>
                      <option value="Courriel">Courriel</option>
                      <option value="Domicile">Domicile</option>
                    </select>
                  </div>
                </div>

                <!-- Sujet -->
                <div class="space-y-2">
                  <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Sujet de la note *</label>
                  <input 
                    v-model="formData.sujet" 
                    :class="[
                      'w-full bg-white dark:bg-gray-900 border rounded-xl p-3 text-lg font-semibold dark:text-white placeholder-gray-300 dark:placeholder-gray-700 outline-none transition-all',
                      errors.sujet ? 'border-red-500 ring-2 ring-red-200 dark:ring-red-900' : 'border-gray-300 dark:border-gray-800 focus:ring-2 focus:ring-blue-500'
                    ]"
                    placeholder="Ex: Évaluation initiale, Suivi hebdomadaire..."
                  />
                  <span v-if="errors.sujet" class="text-xs text-red-600 dark:text-red-400">{{ errors.sujet }}</span>
                </div>
              </div>

              <!-- Zone d'édition -->
              <div class="flex-1 flex flex-col overflow-hidden">
                
                <!-- Barre d'outils -->
                <div class="px-6 py-3 border-b dark:border-gray-800 flex items-center justify-between bg-white dark:bg-gray-900">
                  <div class="flex items-center gap-4 text-xs text-gray-500 dark:text-gray-400">
                    <div class="flex items-center gap-2">
                      <Type :size="14" />
                      <span>{{ wordCount }} mots</span>
                    </div>
                    <div class="w-px h-4 bg-gray-300 dark:bg-gray-700"></div>
                    <div class="flex items-center gap-2">
                      <FileText :size="14" />
                      <span>{{ charCount }} caractères</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center gap-2">
                    <button 
                      @click="insertTemplate"
                      class="px-3 py-1.5 text-xs bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors flex items-center gap-1.5"
                    >
                      <FileType :size="14" />
                      Insérer template
                    </button>
                    <button 
                      @click="clearContent"
                      v-if="formData.contenu"
                      class="px-3 py-1.5 text-xs text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
                    >
                      Effacer
                    </button>
                  </div>
                </div>

                <!-- Textarea principal -->
                <div class="flex-1 overflow-y-auto p-6">
                  <textarea 
                    v-model="formData.contenu" 
                    ref="textareaRef"
                    :class="[
                      'w-full h-full bg-transparent border-none p-0 text-base font-mono leading-relaxed dark:text-gray-200 placeholder-gray-300 dark:placeholder-gray-700 focus:ring-0 resize-none outline-none',
                      errors.contenu ? 'text-red-600 dark:text-red-400' : 'text-gray-800'
                    ]"
                    placeholder="Rédigez votre note ici...

Conseil : Utilisez le template pour structurer votre contenu de manière professionnelle."
                    @keydown.ctrl.s.prevent="handleSave"
                    @keydown.meta.s.prevent="handleSave"
                  ></textarea>
                  <span v-if="errors.contenu" class="text-sm text-red-600 dark:text-red-400 mt-2 block">{{ errors.contenu }}</span>
                </div>
              </div>
            </div>

            <!-- MODE VISUALISATION -->
            <div v-else-if="selectedNote" class="flex-1 flex flex-col overflow-hidden">
              <div class="flex-1 overflow-y-auto p-8">
                <div class="max-w-4xl mx-auto space-y-6">
                  
                  <!-- En-tête de la note -->
                  <div class="border-b dark:border-gray-800 pb-6">
                    <div class="flex items-start justify-between mb-4">
                      <div>
                        <span class="inline-block px-3 py-1 text-xs font-bold bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300 rounded-full mb-3">
                          {{ selectedNote.type_note }}
                        </span>
                        <h3 class="text-2xl font-bold text-gray-900 dark:text-white">{{ selectedNote.sujet }}</h3>
                      </div>
                      <Lock :size="20" class="text-green-500" />
                    </div>
                    
                    <div class="grid grid-cols-2 gap-4 text-sm">
                      <div>
                        <span class="text-gray-500 dark:text-gray-400">Date de création</span>
                        <p class="font-medium dark:text-gray-200">{{ formatDateTime(selectedNote.date_note) }}</p>
                      </div>
                      <div>
                        <span class="text-gray-500 dark:text-gray-400">Signé par</span>
                        <p class="font-medium dark:text-gray-200">{{ selectedNote.signature_nom }}</p>
                      </div>
                    </div>
                  </div>

                  <!-- Contenu de la note -->
                  <div class="prose dark:prose-invert max-w-none">
                    <pre class="whitespace-pre-wrap font-mono text-sm leading-relaxed text-gray-800 dark:text-gray-200 bg-gray-50 dark:bg-gray-900/50 p-6 rounded-xl border dark:border-gray-800">{{ selectedNote.contenu }}</pre>
                  </div>
                </div>
              </div>
            </div>

            <!-- ÉTAT VIDE -->
            <div v-else class="flex-1 flex items-center justify-center text-gray-400 flex-col gap-4">
               <div class="p-6 bg-gray-50 dark:bg-gray-900 rounded-full">
                 <MousePointerClick :size="40" class="opacity-20" />
               </div>
               <p class="text-sm font-medium">Sélectionnez une note ou créez-en une nouvelle</p>
            </div>
          </div>
        </div>

        <!-- Footer avec actions (seulement en mode création) -->
        <div v-if="isCreating" class="px-8 py-4 bg-gray-50 dark:bg-gray-900 border-t dark:border-gray-800 flex items-center justify-between">
          <div class="flex gap-8 items-center text-[11px]">
            <div class="flex flex-col">
              <span class="font-bold text-gray-400 uppercase tracking-tighter leading-none mb-1">Date de création</span>
              <span class="dark:text-gray-300">{{ formatDate(new Date()) }}</span>
            </div>
            <div class="flex flex-col">
              <span class="font-bold text-gray-400 uppercase tracking-tighter leading-none mb-1">Signature Prévue</span>
              <span class="dark:text-blue-400 font-bold italic">Signature automatique active</span>
            </div>
            <div class="flex items-center gap-2 px-3 py-1 bg-orange-100 dark:bg-orange-900/30 text-orange-600 rounded-full">
               <span class="w-1.5 h-1.5 rounded-full bg-orange-500 animate-pulse"></span>
               <span class="font-bold uppercase tracking-tighter text-[9px]">Brouillon non scellé</span>
            </div>
          </div>

          <div class="flex gap-4">
            <button 
              @click="cancelCreation" 
              class="px-4 py-2 text-sm font-bold text-gray-500 hover:text-gray-700 dark:hover:text-gray-300 transition-colors"
            >
              Annuler
            </button>
            <button 
              @click="handleSave" 
              :disabled="isSaving"
              class="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white px-8 py-2.5 rounded-xl font-bold shadow-xl transition-all active:scale-95 flex items-center gap-2"
            >
              <Check :size="18" /> 
              <span v-if="!isSaving">Enregistrer & Verrouiller</span>
              <span v-else>Enregistrement...</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { X, ClipboardList, Plus, Check, Lock, MousePointerClick, FileText, Type, FileType, FileDown } from 'lucide-vue-next'
import { GetClientNotes, CreateNote, GetNoteByID } from '../../wailsjs/go/main/App'
import { jsPDF } from 'jspdf'
import 'jspdf-autotable'
import logoPath from '../assets/images/LeopardLogo.png'

const props = defineProps({
  isOpen: Boolean,
  client: Object
})

const emit = defineEmits(['close'])

const notes = ref([])
const selectedNote = ref(null)
const selectedNotes = ref([])
const isCreating = ref(false)
const isSaving = ref(false)
const isExporting = ref(false)
const textareaRef = ref(null)

const formData = ref({
  sujet: '',
  type_note: '',
  contenu: '',
  date_intervention: new Date().toISOString().split('T')[0],
  mode_intervention: 'En personne',
  type_intervention: ''
})

const errors = ref({
  type_note: '',
  sujet: '',
  contenu: ''
})

// Template clinique de base
const getClinicalTemplate = () => `DATE: ${new Date().toISOString().split('T')[0]}
PERSONNES PRÉSENTES : 
LIEU : 
BUT : 
OBJECTIFS : 
MOYENS : 
FAITS OU OBS : 
INTERVENTIONS : 
RÉSULTATS : 
ENTENTES : `

// Compteurs
const charCount = computed(() => formData.value.contenu.length)
const wordCount = computed(() => {
  const text = formData.value.contenu.trim()
  return text ? text.split(/\s+/).length : 0
})

// Sélection multiple
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

// Export PDF
const exportToPDF = async () => {
  if (selectedNotes.value.length === 0) return
  
  isExporting.value = true
  
  try {
    // Charger toutes les notes sélectionnées
    const fullNotes = await Promise.all(
      selectedNotes.value.map(id => GetNoteByID(id))
    )
    
    // Trier par date
    fullNotes.sort((a, b) => new Date(a.date_note) - new Date(b.date_note))
    
    // Créer le PDF
    const doc = new jsPDF()
    let yPosition = 20
    
    // Logo (en haut à gauche)
    const img = new Image()
    img.src = logoPath
    await new Promise((resolve) => {
      img.onload = () => {
        doc.addImage(img, 'PNG', 15, 10, 30, 30)
        resolve()
      }
    })
    
    // Section identification client (à droite du logo)
    doc.setFontSize(10)
    doc.setFont('helvetica', 'bold')
    doc.text('INFORMATIONS CLIENT', 50, 15)
    
    doc.setFont('helvetica', 'normal')
    doc.setFontSize(9)
    yPosition = 20
    
    const clientInfo = [
      `Nom: ${props.client?.nom || '-'} ${props.client?.prenom || '-'}`,
      `Date de naissance: ${props.client?.date_naissance || '-'}`,
      `No Dossier Leopard: ${props.client?.no_dossier_leopard || '-'}`,
      `No RAMQ: ${props.client?.numero_assurance_maladie || '-'}`,
      `Téléphone: ${props.client?.telephone || '-'}`,
      `Cellulaire: ${props.client?.cellulaire || '-'}`,
      `Email: ${props.client?.email || '-'}`,
      `Adresse: ${props.client?.adresse || '-'}`,
      `${props.client?.code_postal || ''} ${props.client?.ville || ''}, ${props.client?.pays || ''}`
    ]
    
    clientInfo.forEach(line => {
      doc.text(line, 50, yPosition)
      yPosition += 5
    })
    
    yPosition += 10
    
    // Ligne de séparation
    doc.setDrawColor(200, 200, 200)
    doc.line(15, yPosition, 195, yPosition)
    yPosition += 10
    
    // Notes
    fullNotes.forEach((note, index) => {
      // Vérifier si on a assez d'espace, sinon nouvelle page
      if (yPosition > 250) {
        doc.addPage()
        yPosition = 20
      }
      
      // En-tête de la note
      doc.setFillColor(240, 248, 255)
      doc.rect(15, yPosition - 5, 180, 8, 'F')
      
      doc.setFontSize(11)
      doc.setFont('helvetica', 'bold')
      doc.text(`NOTE ${index + 1} - ${note.type_note}`, 17, yPosition)
      yPosition += 10
      
      doc.setFontSize(9)
      doc.setFont('helvetica', 'normal')
      doc.text(`Sujet: ${note.sujet}`, 17, yPosition)
      yPosition += 5
      doc.text(`Date de création: ${formatDateTime(note.date_note)}`, 17, yPosition)
      yPosition += 5
      doc.text(`Signé par: ${note.signature_nom}`, 17, yPosition)
      yPosition += 5
      
      if (note.date_intervention) {
        doc.text(`Date d'intervention: ${formatDate(note.date_intervention)}`, 17, yPosition)
        yPosition += 5
      }
      
      yPosition += 3
      
      // Contenu
      doc.setFont('courier', 'normal')
      doc.setFontSize(8)
      const splitContent = doc.splitTextToSize(note.contenu, 175)
      
      splitContent.forEach(line => {
        if (yPosition > 280) {
          doc.addPage()
          yPosition = 20
        }
        doc.text(line, 17, yPosition)
        yPosition += 4
      })
      
      yPosition += 10
      
      // Séparateur entre notes
      if (index < fullNotes.length - 1) {
        doc.setDrawColor(200, 200, 200)
        doc.line(15, yPosition, 195, yPosition)
        yPosition += 10
      }
    })
    
    // Pied de page sur chaque page
    const pageCount = doc.internal.getNumberOfPages()
    for (let i = 1; i <= pageCount; i++) {
      doc.setPage(i)
      doc.setFontSize(8)
      doc.setFont('helvetica', 'italic')
      doc.text(
        `Imprimé le ${new Date().toLocaleString('fr-CA')} par User #${props.client?.created_by || 'N/A'}`,
        15,
        285
      )
      doc.text(`Page ${i} / ${pageCount}`, 180, 285)
    }
    
    // Sauvegarder
    const fileName = `Dossier_${props.client?.nom}_${props.client?.prenom}_${new Date().toISOString().split('T')[0]}.pdf`
    doc.save(fileName)
    
    alert(`✅ PDF généré avec succès: ${fileName}`)
    
  } catch (err) {
    console.error('❌ Erreur génération PDF:', err)
    alert('❌ Erreur lors de la génération du PDF: ' + err)
  } finally {
    isExporting.value = false
  }
}

// Validation
const validateForm = () => {
  errors.value = {}
  let isValid = true

  if (!formData.value.type_note) {
    errors.value.type_note = 'Le type de note est requis'
    isValid = false
  }

  if (!formData.value.sujet || formData.value.sujet.trim().length === 0) {
    errors.value.sujet = 'Le sujet est requis'
    isValid = false
  }

  if (!formData.value.contenu || formData.value.contenu.trim().length === 0) {
    errors.value.contenu = 'Le contenu ne peut pas être vide'
    isValid = false
  }

  return isValid
}

// Actions
const startNewNote = () => {
  formData.value = {
    sujet: '',
    type_note: 'Suivi',
    contenu: getClinicalTemplate(),
    date_intervention: new Date().toISOString().split('T')[0],
    mode_intervention: 'En personne',
    type_intervention: ''
  }
  errors.value = {}
  selectedNote.value = null
  isCreating.value = true
}

const cancelCreation = () => {
  if (formData.value.contenu && formData.value.contenu !== getClinicalTemplate()) {
    if (!confirm('Voulez-vous vraiment annuler ? Les modifications seront perdues.')) {
      return
    }
  }
  isCreating.value = false
  selectedNote.value = null
}

const insertTemplate = () => {
  if (confirm('Voulez-vous insérer le template ? Cela écrasera le contenu actuel.')) {
    formData.value.contenu = getClinicalTemplate()
    textareaRef.value?.focus()
  }
}

const clearContent = () => {
  if (confirm('Voulez-vous vraiment effacer tout le contenu ?')) {
    formData.value.contenu = ''
    textareaRef.value?.focus()
  }
}

const viewNote = async (noteListItem) => {
  try {
    console.log('🔍 Chargement note complète:', noteListItem.id)
    const fullNote = await GetNoteByID(noteListItem.id)
    selectedNote.value = fullNote
    isCreating.value = false
    console.log('✅ Note chargée:', fullNote.sujet)
  } catch (err) {
    console.error('❌ Erreur chargement note:', err)
    alert('Erreur lors du chargement de la note')
  }
}

const loadNotes = async () => {
  try {
    console.log('🔍 Chargement notes pour client:', props.client?.id)
    const result = await GetClientNotes(props.client?.id)
    notes.value = result || []  // ✅ Si null, on met un tableau vide
    console.log('✅ Notes chargées:', notes.value.length)
  } catch (err) {
    console.error('❌ Erreur chargement notes:', err)
    notes.value = []  // ✅ En cas d'erreur aussi, tableau vide
    alert('Erreur chargement: ' + err)
  }
}

const handleSave = async () => {
  if (!validateForm()) {
    alert('⚠️ Veuillez remplir tous les champs obligatoires')
    return
  }

  isSaving.value = true
  try {
    await CreateNote({
      client_id: props.client?.id,
      date_intervention: formData.value.date_intervention,
      mode_intervention: formData.value.mode_intervention,
      type_intervention: formData.value.type_intervention,
      type_note: formData.value.type_note,
      sujet: formData.value.sujet,
      contenu: formData.value.contenu
    })
    
    alert('✅ Note créée et verrouillée avec succès!')
    isCreating.value = false
    selectedNote.value = null
    await loadNotes()
    
    // Reset form
    formData.value = {
      sujet: '',
      type_note: '',
      contenu: '',
      date_intervention: new Date().toISOString().split('T')[0],
      mode_intervention: 'En personne',
      type_intervention: ''
    }
  } catch (err) {
    console.error('Erreur sauvegarde:', err)
    alert('❌ Erreur: ' + err)
  } finally {
    isSaving.value = false
  }
}

const closeModal = () => {
  if (isCreating.value && formData.value.contenu) {
    if (!confirm('Voulez-vous vraiment fermer ? Les modifications seront perdues.')) {
      return
    }
  }
  isCreating.value = false
  selectedNote.value = null
  selectedNotes.value = []
  emit('close')
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA')
}

const formatDateTime = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleString('fr-CA', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

watch(() => props.isOpen, (newVal) => {
  if (newVal && props.client?.id) {
    loadNotes()
    selectedNote.value = null
    selectedNotes.value = []
    isCreating.value = false
  }
})
</script>

<style scoped>
.modal-fade-enter-active, .modal-fade-leave-active { 
  transition: opacity 0.3s ease; 
}
.modal-fade-enter-from, .modal-fade-leave-to { 
  opacity: 0; 
}

::-webkit-scrollbar { 
  width: 4px; 
  height: 4px;
}
::-webkit-scrollbar-thumb { 
  background: #3b82f6; 
  border-radius: 10px; 
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>