<template>
  <Transition name="modal-fade">
    <div v-if="isOpen && client" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/70 backdrop-blur-md" @click="handleClose" />

      <div class="relative bg-white dark:bg-gray-950 w-full max-w-7xl h-[95vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-gray-200 dark:border-gray-800">

        <EvaluationHeader
          :client="client"
          :selected-evaluation="selectedEvaluation"
          :is-creating="isCreating"
          :is-exporting="isExporting"
          :current-model="currentModel"
          @start-new="openTypeSelector"
          @cancel-creation="cancelCreation"
          @export-pdf="exportPdf"
          @close="handleClose"
        />

        <div class="flex-1 flex overflow-hidden">
          <EvaluationSidebar
            :evaluations="evaluationsFinalisees"
            :selected-evaluation="selectedEvaluation"
            :is-creating="isCreating"
            :active-section="activeSection"
            :current-model="currentModel"
            :form-data="formData"
            @section-change="activeSection = $event"
            @view-evaluation="viewEvaluation"
          />

          <EvaluationViewer
            :is-creating="isCreating"
            :selected-evaluation="selectedEvaluation"
            :form-data="formData"
            :active-section="activeSection"
            :current-model="currentModel"
            :is-saving="isSaving"
            :is-finalizing="isFinalizing"
            @update:form-data="formData = $event"
            @save-draft="saveDraft"
            @finalize="showFinalizeModal = true"
            @cancel="cancelCreation"
          />
        </div>

        <EvaluationFooter
          v-if="!isCreating"
          :brouillons="brouillons"
          @view-draft="viewEvaluation"
        />
      </div>
    </div>
  </Transition>

  <!-- ── Sélecteur de modèle (vient de la DB) ── -->
  <Transition name="modal-fade">
    <div v-if="showTypeSelector" class="fixed inset-0 z-[70] flex items-center justify-center p-4 bg-black/70 backdrop-blur-md">
      <div class="bg-white dark:bg-gray-950 rounded-2xl w-full max-w-4xl border border-gray-200 dark:border-gray-800 shadow-2xl overflow-hidden flex flex-col" style="max-height: 88vh;">

        <!-- ── En-tête modal ── -->
        <div class="flex items-center justify-between px-6 py-4 border-b dark:border-gray-800 shrink-0">
          <div>
            <h3 class="text-lg font-bold dark:text-white text-gray-900">Nouvelle évaluation</h3>
            <p class="text-xs text-gray-400 mt-0.5">Sélectionnez un modèle pour commencer</p>
          </div>
          <button @click="showTypeSelector = false; hoveredDef = null"
                  class="p-1.5 hover:bg-gray-100 dark:hover:bg-gray-800 rounded-lg transition-colors">
            <X :size="18" class="text-gray-400" />
          </button>
        </div>

        <!-- ── Corps : 2 panneaux ── -->
        <div class="flex flex-1 overflow-hidden min-h-0">

          <!-- Panneau gauche : liste des modèles -->
          <div class="w-72 shrink-0 border-r dark:border-gray-800 flex flex-col">

            <!-- Chargement -->
            <div v-if="isLoadingDefs" class="flex items-center justify-center flex-1 gap-3 text-gray-400">
              <Loader2 :size="20" class="animate-spin" />
              <span class="text-sm">Chargement...</span>
            </div>

            <!-- Aucun modèle -->
            <div v-else-if="!definitions.length" class="flex-1 flex flex-col items-center justify-center text-gray-400 p-6">
              <FileX :size="36" class="mb-3 opacity-30" />
              <p class="text-sm font-medium text-center">Aucun modèle disponible</p>
              <p class="text-xs mt-1 opacity-70 text-center">Vérifiez la table evaluation_definitions</p>
            </div>

            <!-- Liste -->
            <div v-else class="flex-1 overflow-y-auto p-3 space-y-1.5">
              <button
                v-for="def in definitions"
                :key="def.id"
                @mouseenter="hoveredDef = def"
                @click="startNewEvaluation(def)"
                :class="[
                  'w-full flex items-center gap-3 p-3 rounded-xl border text-left transition-all group',
                  hoveredDef?.id === def.id
                    ? 'border-teal-300 dark:border-teal-700 bg-teal-50 dark:bg-teal-900/20 shadow-sm'
                    : 'border-gray-200 dark:border-gray-800 hover:border-teal-200 dark:hover:border-teal-800'
                ]"
              >
                <!-- Icône colorée -->
                <div :class="['p-2.5 rounded-xl text-white shadow-sm shrink-0 transition-transform group-hover:scale-105', colorClass(def.couleur)]">
                  <component :is="iconComponent(def.icone)" :size="18" />
                </div>

                <div class="flex-1 min-w-0">
                  <div :class="[
                    'font-bold text-sm leading-tight transition-colors truncate',
                    hoveredDef?.id === def.id
                      ? 'text-teal-700 dark:text-teal-300'
                      : 'text-gray-800 dark:text-gray-200'
                  ]">
                    {{ def.nom }}
                  </div>
                  <div class="text-[10px] text-gray-400 mt-0.5 flex items-center gap-1">
                    <Layers :size="10" />
                    {{ def.schema?.sections?.length || 0 }} sections
                    <span v-if="def.schema?.sections?.length" class="mx-0.5">·</span>
                    <span v-if="def.schema?.sections?.length">
                      {{ def.schema.sections.reduce((acc, s) => acc + (s.fields?.length || s.columns?.flatMap(c => c.fields||[]).length || 0), 0) }} champs
                    </span>
                  </div>
                </div>

                <ChevronRight :size="14"
                  :class="hoveredDef?.id === def.id ? 'text-teal-400' : 'text-gray-300'"
                  class="shrink-0 transition-colors" />
              </button>
            </div>

            <!-- Annuler -->
            <div class="p-3 border-t dark:border-gray-800 shrink-0">
              <button @click="showTypeSelector = false; hoveredDef = null"
                      class="w-full py-2 text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 font-medium text-sm transition-colors rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800">
                Annuler
              </button>
            </div>
          </div>

          <!-- Panneau droit : aperçu du modèle -->
          <div class="flex-1 overflow-y-auto bg-slate-50 dark:bg-gray-900/40 min-w-0">

            <!-- Aucun modèle survolé -->
            <div v-if="!hoveredDef" class="flex flex-col items-center justify-center h-full text-gray-300 dark:text-gray-600 p-8">
              <component :is="iconComponent('FileText')" :size="48" class="mb-4 opacity-30" />
              <p class="text-sm font-medium text-center">Survolez un modèle<br>pour en voir l'aperçu</p>
            </div>

            <!-- Aperçu du modèle sélectionné -->
            <div v-else class="p-6 space-y-5 animate-fadeIn">

              <!-- En-tête du modèle -->
              <div class="flex items-start gap-4">
                <div :class="['p-3 rounded-2xl text-white shadow-md shrink-0', colorClass(hoveredDef.couleur)]">
                  <component :is="iconComponent(hoveredDef.icone)" :size="26" />
                </div>
                <div class="flex-1 min-w-0">
                  <h4 class="text-base font-black text-gray-900 dark:text-white leading-tight">
                    {{ hoveredDef.nom }}
                  </h4>
                  <p v-if="hoveredDef.description" class="text-xs text-gray-500 mt-1 italic">
                    {{ hoveredDef.description }}
                  </p>
                  <!-- Stats rapides -->
                  <div class="flex items-center gap-3 mt-2">
                    <span class="inline-flex items-center gap-1 text-[10px] font-bold text-gray-400 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 px-2 py-0.5 rounded-full">
                      <Layers :size="9" />
                      {{ hoveredDef.schema?.sections?.length || 0 }} sections
                    </span>
                    <span class="inline-flex items-center gap-1 text-[10px] font-bold text-gray-400 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 px-2 py-0.5 rounded-full">
                      <ClipboardList :size="9" />
                      {{ hoveredDef.schema?.sections?.reduce((acc, s) => {
                          const flat = s.fields?.length || (s.columns||[]).flatMap(c=>c.fields||[]).length || 0
                          return acc + flat
                        }, 0) || 0 }} champs
                    </span>
                  </div>
                </div>
              </div>

              <!-- Ligne de séparation -->
              <div class="border-t border-gray-200 dark:border-gray-700" />

              <!-- Plan des sections -->
              <div class="space-y-3">
                <p class="text-[9px] font-black tracking-[0.2em] uppercase text-gray-400">
                  Structure du formulaire
                </p>

                <div
                  v-for="(section, sIdx) in (hoveredDef.schema?.sections || [])"
                  :key="section.id"
                  class="bg-white dark:bg-gray-900 rounded-xl border border-gray-200 dark:border-gray-800 overflow-hidden"
                >
                  <!-- En-tête de section -->
                  <div :class="[
                    'flex items-center gap-2.5 px-3 py-2.5 border-b border-gray-100 dark:border-gray-800',
                    sectionBgClass(section.couleur)
                  ]">
                    <div class="flex items-center justify-center w-5 h-5 rounded-full bg-white/60 dark:bg-black/20 shrink-0">
                      <span class="text-[9px] font-black text-gray-600 dark:text-gray-300">{{ sIdx + 1 }}</span>
                    </div>
                    <component :is="iconComponent(section.icone)" :size="12"
                      :class="sectionIconColor(section.couleur)" class="shrink-0" />
                    <span class="text-xs font-bold text-gray-700 dark:text-gray-200 leading-tight">
                      {{ section.label }}
                    </span>
                    <span class="ml-auto text-[9px] text-gray-400 shrink-0">
                      {{ (section.fields || (section.columns||[]).flatMap(c=>c.fields||[])).length }} champs
                    </span>
                  </div>

                  <!-- Liste des champs -->
                  <div class="px-3 py-2 space-y-1">
                    <div
                      v-for="field in (section.fields || (section.columns||[]).flatMap(c=>c.fields||[])).filter(f => f.type !== 'info')"
                      :key="field.id"
                      class="flex items-center gap-2 py-0.5"
                    >
                      <!-- Badge type de champ -->
                      <span :class="['text-[8px] font-bold px-1.5 py-0.5 rounded font-mono shrink-0', fieldTypeBadge(field.type)]">
                        {{ fieldTypeLabel(field.type) }}
                      </span>
                      <span class="text-[11px] text-gray-600 dark:text-gray-400 truncate leading-tight">
                        {{ field.label }}
                      </span>
                      <span v-if="field.required" class="text-red-400 text-[9px] shrink-0">✱</span>
                    </div>
                  </div>
                </div>

                <!-- Aucune section -->
                <div v-if="!hoveredDef.schema?.sections?.length"
                     class="text-center py-8 text-gray-400 text-sm">
                  Modèle sans sections définies
                </div>
              </div>

              <!-- CTA -->
              <button
                @click="startNewEvaluation(hoveredDef)"
                :class="['w-full flex items-center justify-center gap-2 py-3 rounded-xl font-bold text-sm text-white shadow-lg transition-all active:scale-95', colorClass(hoveredDef.couleur)]"
              >
                <component :is="iconComponent(hoveredDef.icone)" :size="16" />
                Commencer avec ce modèle
              </button>

            </div>
          </div>
        </div>

      </div>
    </div>
  </Transition>

  <!-- ── Modal de finalisation / signature ── -->
  <Transition name="modal-fade">
    <div v-if="showFinalizeModal" class="fixed inset-0 z-[80] flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm">
      <div class="bg-white dark:bg-gray-900 rounded-2xl p-6 w-full max-w-md border border-gray-200 dark:border-gray-800 shadow-2xl">
        <div class="flex items-center gap-3 mb-5">
          <div class="p-2 bg-emerald-100 dark:bg-emerald-900/30 rounded-xl">
            <ShieldCheck :size="22" class="text-emerald-600 dark:text-emerald-400" />
          </div>
          <div>
            <h3 class="text-lg font-bold dark:text-white">Sceller l'évaluation</h3>
            <p class="text-xs text-gray-500 mt-0.5">Cette action est irréversible</p>
          </div>
        </div>

        <div class="p-4 bg-amber-50 dark:bg-amber-900/20 border border-amber-200 dark:border-amber-800 rounded-xl mb-5">
          <p class="text-sm text-amber-700 dark:text-amber-300">
            En scellant ce document, vous attestez qu'il reflète fidèlement votre évaluation professionnelle. Il ne pourra plus être modifié.
          </p>
        </div>

        <div class="space-y-3 mb-5">
          <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300">
            Nom du signataire <span class="text-red-500">*</span>
          </label>
          <input
            v-model="signatureNom"
            type="text"
            placeholder="Prénom Nom, T.S."
            class="w-full px-4 py-2.5 rounded-xl border-2 border-gray-200 dark:border-gray-700 dark:bg-gray-800 dark:text-white focus:outline-none focus:border-teal-500 text-sm"
          />
        </div>

        <div class="flex gap-3">
          <button
            @click="showFinalizeModal = false"
            class="flex-1 py-2.5 border border-gray-200 dark:border-gray-700 rounded-xl text-gray-600 dark:text-gray-300 font-medium text-sm hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors"
          >
            Annuler
          </button>
          <button
            @click="finalizeEvaluation"
            :disabled="!signatureNom.trim() || isFinalizing"
            class="flex-1 flex items-center justify-center gap-2 py-2.5 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 disabled:from-gray-300 disabled:to-gray-400 text-white rounded-xl font-bold text-sm shadow transition-all"
          >
            <Loader2 v-if="isFinalizing" :size="16" class="animate-spin" />
            <ShieldCheck v-else :size="16" />
            <span>{{ isFinalizing ? 'Scellement...' : 'Sceller et signer' }}</span>
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import {
  X, Loader2, FileX, ChevronRight, ShieldCheck,
  ClipboardList, Scale, FileText, Brain, Heart, Activity,
  Briefcase, Users, Layers, CheckCircle, Target, Network
} from 'lucide-vue-next'

import EvaluationHeader  from './EvaluationHeader.vue'
import EvaluationSidebar from './EvaluationSidebar.vue'
import EvaluationViewer  from './EvaluationViewer.vue'
import EvaluationFooter  from './EvaluationFooter.vue'

// ── Props & emits ──────────────────────────────────────────
const props = defineProps({
  isOpen: Boolean,
  client: Object
})
const emit = defineEmits(['close'])

// ── États ─────────────────────────────────────────────────
const evaluations     = ref([])
const definitions     = ref([])   // Chargé depuis la DB
const isLoadingDefs   = ref(false)
const isCreating      = ref(false)
const activeSection   = ref('')
const isSaving        = ref(false)
const isFinalizing    = ref(false)
const isExporting     = ref(false)
const showTypeSelector  = ref(false)
const showFinalizeModal = ref(false)
const hoveredDef        = ref(null)   // modèle survolé dans le sélecteur
const selectedEvaluation = ref(null)
const currentModel    = ref(null)  // Le modèle actif (avec schema parsé)
const signatureNom    = ref('')

const formData = ref({
  id:         0,
  client_id:  null,
  model_id:   '',
  no_leopard: '',
  payload:    {}
})

// ── Computed ───────────────────────────────────────────────
const evaluationsFinalisees = computed(() =>
  (evaluations.value || []).filter(e => e.statut === 'verrouille' || e.statut === 'finalisee')
)
const brouillons = computed(() =>
  (evaluations.value || []).filter(e => e.statut === 'brouillon')
)

// ── Icônes dynamiques ──────────────────────────────────────
const iconMap = {
  ClipboardList, Scale, FileText, Brain, Heart, Activity,
  Briefcase, Users, Layers, CheckCircle, Target, Network,
  ShieldCheck
}
const iconComponent = (name) => iconMap[name] || FileText

const colorMap = {
  teal:    'bg-teal-600',
  blue:    'bg-blue-600',
  purple:  'bg-purple-600',
  amber:   'bg-amber-600',
  emerald: 'bg-emerald-600',
  slate:   'bg-slate-600',
  red:     'bg-red-600',
}
const colorClass = (c) => colorMap[c] || 'bg-teal-600'

// ── Helpers aperçu modèle ─────────────────────────────────
const sectionBgMap = {
  teal:    'bg-teal-50 dark:bg-teal-900/20',
  blue:    'bg-blue-50 dark:bg-blue-900/20',
  purple:  'bg-purple-50 dark:bg-purple-900/20',
  amber:   'bg-amber-50 dark:bg-amber-900/20',
  emerald: 'bg-emerald-50 dark:bg-emerald-900/20',
  red:     'bg-red-50 dark:bg-red-900/20',
  slate:   'bg-slate-50 dark:bg-slate-800/40',
}
const sectionBgClass = (c) => sectionBgMap[c] || sectionBgMap.teal

const colorIconMap = {
  blue:    'text-blue-500', teal: 'text-teal-500', emerald: 'text-emerald-500',
  purple:  'text-purple-500', amber: 'text-amber-500', red: 'text-red-500',
  slate:   'text-slate-400',
}
const sectionIconColor = (c) => colorIconMap[c] || 'text-gray-400'

const fieldTypeLabelMap = {
  text: 'TXT', textarea: 'AREA', date: 'DATE', select: 'SEL',
  checkboxes: 'CHK', table: 'TAB', grid_checkbox: 'GRID',
  obs_select: 'OBS', score_radar: 'RADAR', info: 'INFO',
}
const fieldTypeLabel = (t) => fieldTypeLabelMap[t] || t?.toUpperCase()?.slice(0,4) || '?'

const fieldTypeBadgeMap = {
  text:          'bg-gray-100 dark:bg-gray-800 text-gray-500 dark:text-gray-400',
  textarea:      'bg-blue-50 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400',
  date:          'bg-purple-50 dark:bg-purple-900/30 text-purple-600 dark:text-purple-400',
  select:        'bg-amber-50 dark:bg-amber-900/30 text-amber-600 dark:text-amber-400',
  checkboxes:    'bg-teal-50 dark:bg-teal-900/30 text-teal-600 dark:text-teal-400',
  table:         'bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400',
  grid_checkbox: 'bg-emerald-50 dark:bg-emerald-900/30 text-emerald-600 dark:text-emerald-400',
  obs_select:    'bg-orange-50 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400',
  score_radar:   'bg-rose-50 dark:bg-rose-900/30 text-rose-600 dark:text-rose-400',
  info:          'bg-slate-100 dark:bg-slate-800 text-slate-400',
}
const fieldTypeBadge = (t) => fieldTypeBadgeMap[t] || 'bg-gray-100 text-gray-500'

// ── Chargement depuis la DB ────────────────────────────────
const loadDefinitions = async () => {
  isLoadingDefs.value = true
  try {
    const res = await window.go.main.EvalHandler.GetDefinitions()
    definitions.value = (res || []).map(d => ({
      ...d,
      schema: d.schema_json ? JSON.parse(d.schema_json) : { sections: [] }
    }))
  } catch (e) {
    console.error('Erreur chargement definitions:', e)
    definitions.value = []
  } finally {
    isLoadingDefs.value = false
  }
}

const loadEvaluations = async () => {
  if (!props.client?.id) return
  try {
    const res = await window.go.main.EvalHandler.GetClientEvaluationsV2(props.client.id)
    evaluations.value = (res || []).map(e => ({
      ...e,
      payload: e.payload
        ? (typeof e.payload === 'string' ? JSON.parse(e.payload) : e.payload)
        : {}
    }))
  } catch (e) {
    console.error('Erreur chargement evaluations:', e)
    evaluations.value = []
  }
}

// ── Actions ────────────────────────────────────────────────
const handleClose = () => emit('close')

const openTypeSelector = () => {
  hoveredDef.value = null
  showTypeSelector.value = true
}

const cancelCreation = () => {
  isCreating.value      = false
  showTypeSelector.value = false
  selectedEvaluation.value = null
  currentModel.value    = null
  formData.value = { id: 0, client_id: null, model_id: '', no_leopard: '', payload: {} }
}

const startNewEvaluation = (def) => {
  showTypeSelector.value = false
  hoveredDef.value = null
  currentModel.value = def

  const date = new Date()
  const stamp = `${date.getFullYear()}${String(date.getMonth()+1).padStart(2,'0')}${String(date.getDate()).padStart(2,'0')}`

  formData.value = {
    id:         0,
    client_id:  props.client.id,
    model_id:   def.id,
    no_leopard: `EVA-${props.client.no_dossier_leopard || '000'}-${stamp}`,
    payload:    {}
  }

  // Activer la première section
  if (def.schema?.sections?.length) {
    activeSection.value = def.schema.sections[0].id
  }

  selectedEvaluation.value = null
  isCreating.value = true
}

const viewEvaluation = (evalItem) => {
  // Retrouver le modèle correspondant
  const def = definitions.value.find(d => d.id === evalItem.model_id)
  currentModel.value = def || null

  formData.value = {
    ...evalItem,
    payload: typeof evalItem.payload === 'string'
      ? JSON.parse(evalItem.payload)
      : (evalItem.payload || {})
  }
  selectedEvaluation.value = evalItem

  // Si brouillon → mode édition, sinon → mode lecture
  if (evalItem.statut === 'brouillon') {
    isCreating.value = true
    if (def?.schema?.sections?.length) {
      activeSection.value = def.schema.sections[0].id
    }
  } else {
    isCreating.value = false
  }
}

const saveDraft = async () => {
  isSaving.value = true
  try {
    const payloadString = JSON.stringify(formData.value.payload)
    if (formData.value.id === 0) {
      const newId = await window.go.main.EvalHandler.CreateEvaluationV2({
        client_id:  formData.value.client_id,
        model_id:   formData.value.model_id,
        no_leopard: formData.value.no_leopard,
        payload:    payloadString
      })
      formData.value.id = Number(newId)
    } else {
      await window.go.main.EvalHandler.UpdateEvaluationV2(formData.value.id, payloadString)
    }
    await loadEvaluations()
  } catch (err) {
    console.error('Erreur sauvegarde:', err)
  } finally {
    isSaving.value = false
  }
}

const finalizeEvaluation = async () => {
  if (!signatureNom.value.trim() || formData.value.id === 0) return
  isFinalizing.value = true
  try {
    // Sauvegarder d'abord si nécessaire
    await saveDraft()
    // Puis sceller
    await window.go.main.EvalHandler.SignerEvaluation(formData.value.id, signatureNom.value.trim())
    showFinalizeModal.value = false
    signatureNom.value = ''
    await loadEvaluations()
    cancelCreation()
  } catch (err) {
    console.error('Erreur finalisation:', err)
  } finally {
    isFinalizing.value = false
  }
}

const exportPdf = () => {
  // TODO: brancher le générateur PDF existant
  console.log('Export PDF à brancher')
}

// ── Lifecycle ──────────────────────────────────────────────
onMounted(() => {
  loadDefinitions()
  loadEvaluations()
})

watch(() => props.client?.id, (newId) => {
  if (props.isOpen && newId) loadEvaluations()
})
</script>

<style scoped>
.modal-fade-enter-active, .modal-fade-leave-active { transition: opacity 0.25s ease; }
.modal-fade-enter-from, .modal-fade-leave-to { opacity: 0; }
.animate-spin { animation: spin 1s linear infinite; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
</style>