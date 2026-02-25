<template>
  <Transition name="modal-fade">
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/70 backdrop-blur-md" @click="handleClose" />

      <div class="relative bg-white dark:bg-gray-950 w-full max-w-7xl h-[95vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-gray-200 dark:border-gray-800">

        <!-- HEADER -->
        <EvaluationHeader
          :client="client"
          :selected-evaluation="selectedEvaluation"
          :is-creating="isCreating"
          :is-exporting="isExporting"
          :type-evaluation="formData.type_evaluation"
          @start-new="openTypeSelector"
          @cancel-creation="cancelCreation"
          @export-pdf="handleExportPDF"
          @close="handleClose"
        />

        <!-- CORPS -->
        <div class="flex-1 flex overflow-hidden">
          <EvaluationSidebar
            :evaluations="evaluationsFinalisees"
            :selected-evaluation="selectedEvaluation"
            :is-creating="isCreating"
            :active-section="activeSection"
            :sections="currentSections"
            @section-change="activeSection = $event"
            @view-evaluation="viewEvaluation"
          />

          <EvaluationViewer
            :is-creating="isCreating"
            :selected-evaluation="selectedEvaluation"
            :form-data="formData"
            :errors="errors"
            :active-section="activeSection"
            :sections="currentSections"
            :total-progress="totalProgress"
            :is-saving="isSaving"
            :is-finalizing="isFinalizing"
            @update:form-data="formData = $event"
            @save-draft="saveDraft"
            @finalize="showFinalizeModal = true"
            @cancel="cancelCreation"
          />
        </div>

        <!-- FOOTER BROUILLONS -->
        <EvaluationFooter
          v-if="!isCreating"
          :brouillons="brouillons"
          @view-draft="viewEvaluation"
          @delete-draft="deleteDraft"
        />
      </div>
    </div>
  </Transition>

  <!-- ===== MODAL SÉLECTION TYPE ===== -->
  <Transition name="modal-fade">
    <div v-if="showTypeSelector" class="fixed inset-0 z-[60] flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/80 backdrop-blur-sm" @click="showTypeSelector = false" />
      <div class="relative bg-white dark:bg-gray-900 rounded-2xl shadow-2xl w-full max-w-2xl border border-gray-200 dark:border-gray-700 overflow-hidden">

        <!-- En-tête modal -->
        <div class="bg-gradient-to-r from-slate-800 to-slate-700 px-8 py-6">
          <div class="flex items-center gap-3 mb-1">
            <div class="p-2 bg-white/20 rounded-lg"><ClipboardPlus :size="22" class="text-white" /></div>
            <h3 class="text-xl font-bold text-white">Nouvelle évaluation</h3>
          </div>
          <p class="text-slate-300 text-sm ml-13">Choisissez le type de régime à évaluer</p>
        </div>

        <!-- Sélection type -->
        <div class="p-8 space-y-4">
          <div
            v-for="type in typesEvaluation"
            :key="type.id"
            @click="startNewEvaluation(type.id)"
            :class="[
              'group relative p-5 rounded-xl border-2 cursor-pointer transition-all duration-200',
              'hover:border-teal-500 hover:bg-teal-50 dark:hover:bg-teal-900/20 hover:shadow-lg',
              'border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-800'
            ]"
          >
            <div class="flex items-center gap-4">
              <div :class="['p-3 rounded-xl text-white transition-all', type.colorClass]">
                <component :is="type.icon" :size="28" />
              </div>
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <h4 class="text-base font-bold text-gray-900 dark:text-white">{{ type.label }}</h4>
                  <span :class="['text-[10px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wide', type.badgeClass]">
                    {{ type.badge }}
                  </span>
                </div>
                <p class="text-sm text-gray-500 dark:text-gray-400">{{ type.description }}</p>
              </div>
              <ChevronRight :size="20" class="text-gray-300 group-hover:text-teal-500 transition-colors" />
            </div>
          </div>
        </div>

        <div class="px-8 pb-6 text-center">
          <button @click="showTypeSelector = false" class="text-sm text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 transition-colors">
            Annuler
          </button>
        </div>
      </div>
    </div>
  </Transition>

  <!-- ===== MODAL FINALISATION ===== -->
  <Transition name="modal-fade">
    <div v-if="showFinalizeModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/80 backdrop-blur-sm" @click="showFinalizeModal = false" />
      <div class="relative bg-white dark:bg-gray-900 rounded-2xl shadow-2xl max-w-lg w-full p-8 border-2 border-emerald-500">

        <div class="flex flex-col items-center text-center mb-6">
          <div class="p-4 bg-emerald-100 dark:bg-emerald-900/30 rounded-full mb-4">
            <ShieldCheck class="text-emerald-600 dark:text-emerald-400" :size="48" />
          </div>
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">Sceller définitivement ?</h3>
          <p class="text-gray-500 dark:text-gray-400 text-sm">
            Cette action est <strong>irréversible</strong>. L'évaluation sera cryptée, verrouillée et signée.
          </p>
        </div>

        <!-- Nom Léopard généré -->
        <div class="mb-6 bg-slate-900 dark:bg-black rounded-xl p-4 border border-slate-700 font-mono text-center">
          <p class="text-[10px] text-slate-400 uppercase tracking-widest mb-1">Identifiant Léopard</p>
          <p class="text-teal-400 text-lg font-bold tracking-wider">{{ nomLeopardGenere }}</p>
        </div>

        <div class="space-y-2 mb-6 bg-slate-50 dark:bg-gray-800 rounded-xl p-4 text-sm text-gray-700 dark:text-gray-300">
          <div class="flex items-center gap-2"><Check class="text-emerald-500 flex-shrink-0" :size="16" /><span>Verrouillée — aucune modification possible</span></div>
          <div class="flex items-center gap-2"><Check class="text-emerald-500 flex-shrink-0" :size="16" /><span>Signée avec date et heure exactes</span></div>
          <div class="flex items-center gap-2"><Check class="text-emerald-500 flex-shrink-0" :size="16" /><span>Cryptée en base de données (Loi 25)</span></div>
          <div class="flex items-center gap-2"><Check class="text-emerald-500 flex-shrink-0" :size="16" /><span>PDF généré automatiquement</span></div>
        </div>

        <!-- Champ signature -->
        <div class="mb-6">
          <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">Nom de signature (T.S., OTSTCFQ)</label>
          <input
            v-model="signatureNom"
            type="text"
            placeholder="Ex: Marie Tremblay, T.S."
            class="w-full px-4 py-3 rounded-xl border-2 border-gray-200 dark:border-gray-700 dark:bg-gray-800 dark:text-white focus:outline-none focus:border-emerald-500 font-medium"
          />
        </div>

        <div class="flex gap-3">
          <button @click="showFinalizeModal = false" class="flex-1 px-6 py-3 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-800 dark:text-gray-200 rounded-xl font-bold transition-all">
            Annuler
          </button>
          <button
            @click="finalizeEvaluation"
            :disabled="isFinalizing || !signatureNom.trim()"
            class="flex-1 px-6 py-3 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 text-white rounded-xl font-bold shadow-lg disabled:opacity-50 transition-all flex items-center justify-center gap-2"
          >
            <Loader2 v-if="isFinalizing" :size="18" class="animate-spin" />
            <ShieldCheck v-else :size="18" />
            <span>{{ isFinalizing ? 'Finalisation...' : 'Sceller et générer PDF' }}</span>
          </button>
        </div>
      </div>
    </div>
  </Transition>

  <!-- ===== TOAST NOTIFICATIONS ===== -->
  <Transition name="toast-slide">
    <div v-if="toast.visible" :class="['fixed bottom-6 right-6 z-[70] flex items-center gap-3 px-5 py-4 rounded-xl shadow-2xl border text-sm font-bold max-w-sm', toast.type === 'success' ? 'bg-emerald-900 border-emerald-600 text-emerald-100' : toast.type === 'error' ? 'bg-red-900 border-red-600 text-red-100' : 'bg-slate-800 border-slate-600 text-slate-100']">
      <CheckCircle v-if="toast.type === 'success'" :size="20" class="text-emerald-400 flex-shrink-0" />
      <AlertCircle v-else-if="toast.type === 'error'" :size="20" class="text-red-400 flex-shrink-0" />
      <Info v-else :size="20" class="text-slate-400 flex-shrink-0" />
      <span>{{ toast.message }}</span>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ShieldCheck, Check, ChevronRight, ClipboardPlus, Loader2, CheckCircle, AlertCircle, Info, Scale, FileHeart, FilePen, Stethoscope } from 'lucide-vue-next'
import EvaluationHeader from './EvaluationHeader.vue'
import EvaluationSidebar from './EvaluationSidebar.vue'
import EvaluationViewer from './EvaluationViewer.vue'
import EvaluationFooter from './EvaluationFooter.vue'
import { generateEvaluationPDF } from './evaluationPdfGenerator.js'
import { FileText, Brain, Heart, Activity, Briefcase, Users, Layers, CheckCircle as CheckCircleIcon, Target, Gavel, ScrollText, UserCheck } from 'lucide-vue-next'

const props = defineProps({ isOpen: Boolean, client: Object })
const emit = defineEmits(['close'])

// ── États ──────────────────────────────────────────────────
const evaluations     = ref([])
const selectedEvaluation = ref(null)
const isCreating      = ref(false)
const isSaving        = ref(false)
const isExporting     = ref(false)
const isFinalizing    = ref(false)
const showFinalizeModal  = ref(false)
const showTypeSelector   = ref(false)
const activeSection   = ref('contexte')
const signatureNom    = ref('')

const toast = ref({ visible: false, message: '', type: 'success' })

const formData = ref(emptyForm())
const errors   = ref({})

// ── Types d'évaluation ─────────────────────────────────────
const typesEvaluation = [
  {
    id: 'tutelle',
    label: 'Régime de tutelle',
    badge: 'Curateur Public',
    description: 'Évaluation du fonctionnement social pour demande de tutelle au majeur inapte',
    icon: Scale,
    colorClass: 'bg-blue-600',
    badgeClass: 'bg-blue-100 text-blue-700 dark:bg-blue-900/50 dark:text-blue-300'
  },
  {
    id: 'mandat',
    label: 'Homologation de mandat',
    badge: 'Mandat de protection',
    description: 'Évaluation pour homologation d\'un mandat donné en prévision d\'inaptitude',
    icon: FileHeart,
    colorClass: 'bg-purple-600',
    badgeClass: 'bg-purple-100 text-purple-700 dark:bg-purple-900/50 dark:text-purple-300'
  },
  {
    id: 'conseil_tutelle',
    label: 'Conseil de tutelle',
    badge: 'Protection',
    description: 'Évaluation pour la constitution d\'un conseil de tutelle au mineur ou majeur',
    icon: UserCheck,
    colorClass: 'bg-amber-600',
    badgeClass: 'bg-amber-100 text-amber-700 dark:bg-amber-900/50 dark:text-amber-300'
  },
  {
    id: 'suivi_psychosocial',
    label: 'Suivi psychosocial',
    badge: 'Note évolutive',
    description: 'Note de suivi psychosocial, bilan d\'intervention ou rapport de situation',
    icon: Stethoscope,
    colorClass: 'bg-teal-600',
    badgeClass: 'bg-teal-100 text-teal-700 dark:bg-teal-900/50 dark:text-teal-300'
  }
]

// ── Sections par type ──────────────────────────────────────
const sectionsByType = {
  tutelle: [
    { id: 'contexte',      label: 'Contexte',           icon: FileText },
    { id: 'cognitif',      label: 'Cognitif',            icon: Brain },
    { id: 'sante',         label: 'Santé physique',      icon: Heart },
    { id: 'psycho',        label: 'Psycho-Social',       icon: Activity },
    { id: 'roles',         label: 'Rôles sociaux',       icon: Briefcase },
    { id: 'reseau',        label: 'Réseau soutien',      icon: Users },
    { id: 'analyse',       label: 'Analyse clinique',    icon: Layers },
    { id: 'opinion',       label: 'Opinion prof.',       icon: CheckCircleIcon },
    { id: 'recommandations', label: 'Recommandations',   icon: Target }
  ],
  mandat: [
    { id: 'contexte',      label: 'Contexte',            icon: FileText },
    { id: 'inaptitude',    label: 'Inaptitude',          icon: Brain },
    { id: 'sante',         label: 'Santé',               icon: Heart },
    { id: 'psycho',        label: 'Psycho-Social',       icon: Activity },
    { id: 'mandataire',    label: 'Mandataire',          icon: UserCheck },
    { id: 'reseau',        label: 'Réseau soutien',      icon: Users },
    { id: 'analyse',       label: 'Analyse',             icon: Layers },
    { id: 'opinion',       label: 'Opinion',             icon: CheckCircleIcon },
    { id: 'recommandations', label: 'Recommandations',   icon: Target }
  ],
  conseil_tutelle: [
    { id: 'contexte',       label: 'Contexte',           icon: FileText },
    { id: 'situation',      label: 'Situation globale',  icon: Activity },
    { id: 'sante',          label: 'Santé',              icon: Heart },
    { id: 'reseau',         label: 'Réseau',             icon: Users },
    { id: 'analyse',        label: 'Analyse',            icon: Layers },
    { id: 'recommandations', label: 'Recommandations',   icon: Target }
  ],
  suivi_psychosocial: [
    { id: 'contexte',       label: 'Contexte',           icon: FileText },
    { id: 'evolution',      label: 'Évolution',          icon: Activity },
    { id: 'intervention',   label: 'Intervention',       icon: Briefcase },
    { id: 'reseau',         label: 'Réseau',             icon: Users },
    { id: 'analyse',        label: 'Analyse',            icon: Layers },
    { id: 'plan',           label: 'Plan d\'action',     icon: Target }
  ]
}

const currentSections = computed(() =>
  sectionsByType[formData.value.type_evaluation] || sectionsByType.tutelle
)

// ── Computed ───────────────────────────────────────────────
const brouillons = computed(() => evaluations.value.filter(e => e.verrouille === 0))
const evaluationsFinalisees = computed(() => evaluations.value.filter(e => e.verrouille === 1))

const totalProgress = computed(() => {
  const fields = Object.keys(formData.value).filter(k =>
    !['client_id', 'type_evaluation', 'nom_leopard'].includes(k) && typeof formData.value[k] === 'string'
  )
  if (!fields.length) return 0
  const filled = fields.filter(f => formData.value[f]?.trim()).length
  return Math.round((filled / fields.length) * 100)
})

const nomLeopardGenere = computed(() => {
  if (!props.client?.no_dossier_leopard) return 'EVA-XXXXXXXX'
  const date = new Date().toISOString().slice(0, 10).replace(/-/g, '')
  const type = (formData.value.type_evaluation || 'tutelle').slice(0, 3).toUpperCase()
  return `EVA-${props.client.no_dossier_leopard}-${type}-${date}`
})

// ── Helpers ────────────────────────────────────────────────
function emptyForm(type = 'tutelle') {
  return {
    client_id:               null,
    type_evaluation:         type,
    nom_leopard:             '',
    // sections communes
    contexte_evaluation:     '',
    motif_reference:         '',
    objet_evaluation:        '',
    // cognitif / inaptitude
    capacites_cognitives:    '',
    inaptitude_constatee:    '',
    // santé
    etat_sante_physique:     '',
    // psycho-social
    dimensions_psycho_sociales: '',
    // rôles
    roles_sociaux:           '',
    // réseau
    reseau_social_soutien:   '',
    // mandat spécifique
    evaluation_mandataire:   '',
    // suivi
    evolution_situation:     '',
    objectifs_intervention:  '',
    plan_action:             '',
    // analyse / opinion / reco
    analyse_clinique:        '',
    opinion_professionnelle: '',
    recommandations:         ''
  }
}

function showToast(message, type = 'success', duration = 4000) {
  toast.value = { visible: true, message, type }
  setTimeout(() => { toast.value.visible = false }, duration)
}

// ── Actions ────────────────────────────────────────────────
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

const openTypeSelector = () => {
  showTypeSelector.value = true
}

const startNewEvaluation = (type) => {
  showTypeSelector.value = false
  formData.value = emptyForm(type)
  formData.value.client_id = props.client?.id
  formData.value.nom_leopard = nomLeopardGenere.value
  errors.value = {}
  isCreating.value = true
  selectedEvaluation.value = null
  activeSection.value = 'contexte'
}

const cancelCreation = () => {
  if (totalProgress.value > 10 && !confirm('Quitter sans sauvegarder ?')) return
  isCreating.value = false
  formData.value = emptyForm()
}

const viewEvaluation = async (evalItem) => {
  try {
    const fullEval = await window.go.main.EvalHandler.GetEvaluationByID(evalItem.id)
    selectedEvaluation.value = fullEval
    if (fullEval.verrouille === 0) {
      // Reprise brouillon
      formData.value = { ...emptyForm(fullEval.type_evaluation || 'tutelle'), ...fullEval, client_id: props.client?.id }
      isCreating.value = true
    } else {
      isCreating.value = false
    }
  } catch (err) {
    console.error('❌ Erreur:', err)
    showToast('Erreur lors du chargement', 'error')
  }
}

const validateForm = () => {
  errors.value = {}
  if (!formData.value.contexte_evaluation?.trim()) {
    errors.value.contexte_evaluation = 'Le contexte est requis'
    return false
  }
  if (!formData.value.motif_reference?.trim()) {
    errors.value.motif_reference = 'Le motif est requis'
    return false
  }
  return true
}

const saveDraft = async () => {
  if (!validateForm()) {
    showToast('⚠️ Veuillez remplir les champs obligatoires', 'error')
    return
  }
  isSaving.value = true
  try {
    await window.go.main.EvalHandler.CreateEvaluation({ ...formData.value, client_id: props.client?.id })
    showToast('Brouillon sauvegardé', 'success')
    isCreating.value = false
    await loadEvaluations()
  } catch (err) {
    showToast('Erreur lors de la sauvegarde: ' + err, 'error')
  } finally {
    isSaving.value = false
  }
}

const finalizeEvaluation = async () => {
  if (!signatureNom.value.trim()) return
  showFinalizeModal.value = false
  isFinalizing.value = true

  try {
    const nomLeopard = nomLeopardGenere.value
    const result = await window.go.main.EvalHandler.CreateEvaluation({
      ...formData.value,
      client_id: props.client?.id,
      nom_leopard: nomLeopard
    })

    await window.go.main.EvalHandler.LockEvaluation(result, signatureNom.value)

    const fullEval = await window.go.main.EvalHandler.GetEvaluationByID(result)
    selectedEvaluation.value = fullEval

    // Génération PDF automatique
    await handleExportPDF(fullEval)

    await loadEvaluations()
    isCreating.value = false
    signatureNom.value = ''
    showToast('✅ Évaluation scellée — PDF généré avec succès', 'success')

  } catch (err) {
    console.error('Erreur finalisation:', err)
    showToast('Erreur lors de la finalisation: ' + err, 'error')
  } finally {
    isFinalizing.value = false
  }
}

const handleExportPDF = async (evalOverride = null) => {
  const evalData = evalOverride || selectedEvaluation.value
  if (!evalData) {
    showToast('Aucune évaluation sélectionnée', 'error')
    return
  }
  isExporting.value = true
  try {
    await generateEvaluationPDF(evalData, props.client)
    showToast('PDF généré avec succès', 'success')
  } catch (err) {
    console.error('❌ Erreur PDF:', err)
    showToast('Erreur génération PDF: ' + err.message, 'error')
  } finally {
    isExporting.value = false
  }
}

const deleteDraft = async (draftId) => {
  if (!confirm('Supprimer ce brouillon définitivement ?')) return
  try {
    await window.go.main.EvalHandler.DeleteEvaluation(draftId)
    showToast('Brouillon supprimé', 'success')
    await loadEvaluations()
  } catch (err) {
    showToast('Erreur lors de la suppression', 'error')
  }
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
    signatureNom.value = localStorage.getItem('user_full_name') || ''
  }
})
</script>

<style scoped>
.modal-fade-enter-active, .modal-fade-leave-active { transition: opacity 0.25s ease; }
.modal-fade-enter-from, .modal-fade-leave-to { opacity: 0; }

.toast-slide-enter-active, .toast-slide-leave-active { transition: all 0.3s ease; }
.toast-slide-enter-from { opacity: 0; transform: translateX(2rem); }
.toast-slide-leave-to { opacity: 0; transform: translateX(2rem); }

.animate-spin { animation: spin 1s linear infinite; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
</style>