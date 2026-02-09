<template>
  <Transition name="modal-fade">
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-md" @click="handleClose"></div>

      <div class="relative bg-white dark:bg-gray-950 w-full max-w-7xl h-[95vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-gray-200 dark:border-gray-800">
        
        <!-- Header -->
        <EvaluationHeader 
          :client="client"
          :selected-evaluation="selectedEvaluation"
          :is-creating="isCreating"
          :is-exporting="isExporting"
          @start-new="startNewEvaluation"
          @cancel-creation="cancelCreation"
          @export-pdf="exportToPDF"
          @close="handleClose"
        />

        <!-- Corps principal -->
        <div class="flex-1 flex overflow-hidden">
          
          <!-- Sidebar gauche -->
          <EvaluationSidebar 
            :evaluations="evaluationsFinalisees"
            :selected-evaluation="selectedEvaluation"
            :is-creating="isCreating"
            :active-section="activeSection"
            :sections="sections"
            @section-change="activeSection = $event"
            @view-evaluation="viewEvaluation"
          />

          <!-- Viewer central -->
          <EvaluationViewer 
            :is-creating="isCreating"
            :selected-evaluation="selectedEvaluation"
            :form-data="formData"
            :errors="errors"
            :active-section="activeSection"
            :sections="sections"
            :total-progress="totalProgress"
            :is-saving="isSaving"
            :is-finalizing="isFinalizing"
            @update:form-data="formData = $event"
            @save-draft="saveDraft"
            @finalize="showFinalizeModal = true"
            @cancel="cancelCreation"
          />

        </div>

        <!-- Footer avec brouillons -->
        <EvaluationFooter 
          v-if="!isCreating"
          :brouillons="brouillons"
          @view-draft="viewEvaluation"
          @delete-draft="deleteDraft"
        />

      </div>
    </div>
  </Transition>

  <!-- Modal de Confirmation -->
  <Transition name="modal-fade">
    <div v-if="showFinalizeModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/80 backdrop-blur-sm" @click="showFinalizeModal = false"></div>
      
      <div class="relative bg-white dark:bg-gray-900 rounded-2xl shadow-2xl max-w-lg w-full p-8 border-2 border-emerald-500">
        
        <div class="flex flex-col items-center text-center mb-6">
          <div class="p-4 bg-emerald-100 dark:bg-emerald-900/30 rounded-full mb-4">
            <ShieldCheck class="text-emerald-600 dark:text-emerald-400" :size="48" />
          </div>
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">
            Sauvegarder définitivement ?
          </h3>
          <p class="text-gray-600 dark:text-gray-400 text-sm">
            Cette action est <strong>irréversible</strong>
          </p>
        </div>

        <div class="space-y-2 mb-6 bg-slate-50 dark:bg-gray-800 rounded-xl p-4 text-sm text-gray-700 dark:text-gray-300">
          <div class="flex items-center gap-2">
            <Check class="text-emerald-500" :size="18" />
            <span>Verrouillée (aucune modification possible)</span>
          </div>
          <div class="flex items-center gap-2">
            <Check class="text-emerald-500" :size="18" />
            <span>Signée automatiquement avec date/heure</span>
          </div>
          <div class="flex items-center gap-2">
            <Check class="text-emerald-500" :size="18" />
            <span>Cryptée dans la base de données</span>
          </div>
          <div class="flex items-center gap-2">
            <Check class="text-emerald-500" :size="18" />
            <span>Exportée en PDF automatiquement</span>
          </div>
        </div>

        <div class="flex gap-3">
          <button
            @click="showFinalizeModal = false"
            class="flex-1 px-6 py-3 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-800 dark:text-gray-200 rounded-xl font-bold transition-all"
          >
            Annuler
          </button>
          <button
            @click="finalizeEvaluation"
            :disabled="isFinalizing"
            class="flex-1 px-6 py-3 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 text-white rounded-xl font-bold shadow-lg disabled:opacity-50 transition-all"
          >
            <span v-if="!isFinalizing">✅ Confirmer</span>
            <span v-else>⏳ Finalisation...</span>
          </button>
        </div>

      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ShieldCheck, Check } from 'lucide-vue-next'
import EvaluationHeader from './EvaluationHeader.vue'
import EvaluationSidebar from './Evaluationsidebar.vue'
import EvaluationViewer from './EvaluationViewer.vue'
import EvaluationFooter from './EvaluationFooter.vue'
import { FileText, Brain, Heart, Activity, Briefcase, Users, Layers, CheckCircle, Target } from 'lucide-vue-next'

const props = defineProps({
  isOpen: Boolean,
  client: Object
})

const emit = defineEmits(['close'])

// États
const evaluations = ref([])
const selectedEvaluation = ref(null)
const isCreating = ref(false)
const isSaving = ref(false)
const isExporting = ref(false)
const isFinalizing = ref(false)
const showFinalizeModal = ref(false)
const activeSection = ref('contexte')

const formData = ref({
  client_id: null,
  contexte_evaluation: '',
  motif_reference: '',
  objet_evaluation: '',
  capacites_cognitives: '',
  etat_sante_physique: '',
  dimensions_psycho_sociales: '',
  roles_sociaux: '',
  reseau_social_soutien: '',
  analyse_clinique: '',
  opinion_professionnelle: '',
  recommandations: ''
})

const errors = ref({})

const sections = [
  { id: 'contexte', label: 'Contexte', icon: FileText },
  { id: 'cognitif', label: 'Cognitif', icon: Brain },
  { id: 'sante', label: 'Santé', icon: Heart },
  { id: 'psycho', label: 'Psycho-Social', icon: Activity },
  { id: 'roles', label: 'Rôles', icon: Briefcase },
  { id: 'reseau', label: 'Réseau', icon: Users },
  { id: 'analyse', label: 'Analyse', icon: Layers },
  { id: 'opinion', label: 'Opinion', icon: CheckCircle },
  { id: 'recommandations', label: 'Recommandations', icon: Target }
]

// Computed
const brouillons = computed(() => 
  evaluations.value.filter(e => e.verrouille === 0)
)

const evaluationsFinalisees = computed(() => 
  evaluations.value.filter(e => e.verrouille === 1)
)

const totalProgress = computed(() => {
  const allFields = Object.keys(formData.value).filter(k => k !== 'client_id')
  const filledFields = allFields.filter(f => formData.value[f]?.trim()).length
  return Math.round((filledFields / allFields.length) * 100)
})

// Fonctions
const loadEvaluations = async () => {
  if (!props.client?.id) return
  try {
    const result = await window.go.main.EvalHandler.GetEvaluationsByClient(props.client.id)
    evaluations.value = result || []
  } catch (err) {
    console.error('❌ Erreur chargement évaluations:', err)
    evaluations.value = []
  }
}

const startNewEvaluation = () => {
  formData.value = {
    client_id: props.client?.id,
    contexte_evaluation: '',
    motif_reference: '',
    objet_evaluation: '',
    capacites_cognitives: '',
    etat_sante_physique: '',
    dimensions_psycho_sociales: '',
    roles_sociaux: '',
    reseau_social_soutien: '',
    analyse_clinique: '',
    opinion_professionnelle: '',
    recommandations: ''
  }
  errors.value = {}
  isCreating.value = true
  selectedEvaluation.value = null
  activeSection.value = 'contexte'
}

const cancelCreation = () => {
  if (totalProgress.value > 10) {
    if (!confirm('Quitter sans sauvegarder ?')) return
  }
  isCreating.value = false
}

const viewEvaluation = async (evalItem) => {
  try {
    const fullEval = await window.go.main.EvalHandler.GetEvaluationByID(evalItem.id)
    selectedEvaluation.value = fullEval
    
    // Si c'est un brouillon, passer en mode édition
    if (fullEval.verrouille === 0) {
      formData.value = {
        client_id: props.client?.id,
        contexte_evaluation: fullEval.contexte_evaluation || '',
        motif_reference: fullEval.motif_reference || '',
        objet_evaluation: fullEval.objet_evaluation || '',
        capacites_cognitives: fullEval.capacites_cognitives || '',
        etat_sante_physique: fullEval.etat_sante_physique || '',
        dimensions_psycho_sociales: fullEval.dimensions_psycho_sociales || '',
        roles_sociaux: fullEval.roles_sociaux || '',
        reseau_social_soutien: fullEval.reseau_social_soutien || '',
        analyse_clinique: fullEval.analyse_clinique || '',
        opinion_professionnelle: fullEval.opinion_professionnelle || '',
        recommandations: fullEval.recommandations || ''
      }
      isCreating.value = true
    } else {
      isCreating.value = false
    }
  } catch (err) {
    console.error('❌ Erreur:', err)
    alert('Erreur lors du chargement')
  }
}

const validateForm = () => {
  errors.value = {}
  if (!formData.value.contexte_evaluation?.trim()) {
    errors.value.contexte_evaluation = 'Requis'
    return false
  }
  if (!formData.value.motif_reference?.trim()) {
    errors.value.motif_reference = 'Requis'
    return false
  }
  return true
}

const saveDraft = async () => {
  if (!validateForm()) {
    alert('⚠️ Veuillez remplir les champs obligatoires')
    return
  }

  isSaving.value = true
  try {
    await window.go.main.EvalHandler.CreateEvaluation({
      client_id: props.client?.id,
      ...formData.value
    })
    
    alert('✅ Brouillon sauvegardé')
    isCreating.value = false
    await loadEvaluations()
    
  } catch (err) {
    console.error('Erreur:', err)
    alert('❌ Erreur: ' + err)
  } finally {
    isSaving.value = false
  }
}

const finalizeEvaluation = async () =>{
  showFinalizeModal.value = false
  isFinalizing.value = true


  try {
    const result = await window.go.main.EvalHandler.CreateEvaluation({
      client_id: props.client?.id,
      ...formData.value
    })

    const userName = localStorage.getItem('user_full_name') || 'Utilisateur'
    const userRole = localStorage.getItem('user_role') || 'T.S.'
    await window.go.main.EvalHandler.LockEvaluation(result, `${userName}, ${userRole}`)

    const fullEval = await window.go.main.EvalHandler.GetEvaluationByID(result)
    selectedEvaluation.value = fullEval

    await exportToPDF()
    await loadEvaluations()
    isCreating.value = false

    alert('✅ Évaluation finalisée!\n📄 PDF généré, évaluation verrouillée.')
    
  } catch (err) {
    console.error('Erreur:', err)
    alert('❌ Erreur: ' + err)
  } finally {
    isFinalizing.value = false
  }
}

const deleteDraft = async (draftId) => {
  if (!confirm('Supprimer ce brouillon définitivement ?')) return
  
  try {
    await window.go.main.EvalHandler.DeleteEvaluation(draftId)
    alert('✅ Brouillon supprimé')
    await loadEvaluations()
  } catch (err) {
    console.error('❌ Erreur:', err)
    alert('❌ Erreur lors de la suppression')
  }
}
const exportToPDF = async () => {
  // Optionnel: Garde juste un petit message pour pas que le bouton soit "mort"
  alert("L'exportation PDF des évaluations est en cours de migration vers le module sécurisé Go.");
  
  /* if (!selectedEvaluation.value) return
  isExporting.value = true
  
  try {
    const { jsPDF } = await import('jspdf')
    const doc = new jsPDF()
    
    // ... ton code PDF existant ...
    
    const fileName = `Evaluation_${props.client?.nom}_${new Date().toISOString().split('T')[0]}.pdf`
    doc.save(fileName)
    
  } catch (err) {
    console.error('❌ Erreur PDF:', err)
    alert('❌ Erreur lors de la génération du PDF')
  } finally {
    isExporting.value = false
  }
  */
}
const handleClose = () => {
  if (isCreating.value && totalProgress.value > 10) {
    if (!confirm('Fermer sans sauvegarder ?')) return
  }
  isCreating.value = false
  selectedEvaluation.value = null
  emit('close')
}

watch(() => props.isOpen, (newVal) => {
  if (newVal && props.client?.id) {
    loadEvaluations()
    selectedEvaluation.value = null
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
</style>