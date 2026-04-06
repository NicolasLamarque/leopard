<template>
  <div class="flex-1 flex flex-col bg-white dark:bg-gray-950 overflow-hidden">

    <!-- ══════════════════════════════════════
         MODE CRÉATION / ÉDITION
    ══════════════════════════════════════ -->
    <div v-if="isCreating && !isPreviewing" class="flex-1 overflow-y-auto">
      <div class="max-w-3xl mx-auto p-8 space-y-6">

        <!-- Bandeau modèle -->
        <div class="flex items-center gap-3 px-4 py-3 rounded-xl bg-teal-50 dark:bg-teal-900/20 border border-teal-200 dark:border-teal-800">
          <component :is="iconComponent(currentModel?.icone)" :size="18" class="text-teal-600 dark:text-teal-400 shrink-0" />
          <span class="text-sm font-bold text-teal-700 dark:text-teal-300">{{ currentModel?.nom || 'Évaluation' }}</span>
          <span class="ml-auto font-mono text-xs text-teal-500 opacity-70">{{ formData.no_leopard }}</span>
        </div>

        <!-- Pas de section active -->
        <div v-if="!activeModelSection" class="text-center py-20 text-gray-400">
          <FileText :size="48" class="mx-auto mb-4 opacity-20" />
          <p class="text-sm">Sélectionnez une section dans la barre latérale</p>
        </div>

        <!-- Section active → SectionLayout -->
        <div v-else class="space-y-6 animate-fadeIn">
          <SectionLayout
            :section="activeModelSection"
            :payload="formData.payload || {}"
            @update="updatePayload"
          />
        </div>

        <!-- Navigation entre sections -->
        <div class="flex items-center justify-between pt-4 border-t dark:border-gray-800">
          <button
            v-if="prevSection"
            @click="$emit('section-change', prevSection.id)"
            class="flex items-center gap-2 text-sm text-gray-500 hover:text-gray-700 dark:hover:text-gray-300 transition-colors"
          >
            <ChevronLeft :size="16" />
            {{ prevSection.label }}
          </button>
          <div v-else />
          <button
            v-if="nextSection"
            @click="$emit('section-change', nextSection.id)"
            class="flex items-center gap-2 text-sm text-teal-600 dark:text-teal-400 hover:text-teal-700 dark:hover:text-teal-300 font-medium transition-colors"
          >
            {{ nextSection.label }}
            <ChevronRight :size="16" />
          </button>
        </div>

      </div>
    </div>

    <!-- ══════════════════════════════════════
         MODE APERÇU / VISUALISATION
         Rendu document clinique complet
    ══════════════════════════════════════ -->
    <div v-else-if="isPreviewing" class="flex-1 overflow-y-auto bg-slate-100 dark:bg-gray-900">
      <div class="max-w-4xl mx-auto py-8 px-4 space-y-0" id="eval-preview-document">

        <!-- ── PAGE DOCUMENT ── -->
        <div class="bg-white dark:bg-gray-950 shadow-xl rounded-none" style="min-height: 29.7cm;">

          <!-- EN-TÊTE DOCUMENT -->
          <div class="border-b-2 border-gray-900 dark:border-gray-100 px-10 pt-8 pb-5">
            <div class="flex items-start justify-between gap-6">

              <!-- Gauche : titre + sous-titre -->
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <component :is="iconComponent(currentModel?.icone)" :size="14" class="text-gray-400" />
                  <p class="text-[10px] font-bold tracking-[0.2em] uppercase text-gray-400">
                    Évaluation psychosociale
                  </p>
                </div>
                <h1 class="text-2xl font-black text-gray-900 dark:text-white leading-tight tracking-tight">
                  {{ currentModel?.nom || 'Évaluation' }}
                </h1>
                <p v-if="currentModel?.description" class="text-xs text-gray-500 mt-1 italic">
                  {{ currentModel.description }}
                </p>
              </div>

              <!-- Droite : méta -->
              <div class="shrink-0 text-right space-y-0.5">
                <p class="font-mono text-xs font-bold text-gray-900 dark:text-gray-100 bg-gray-100 dark:bg-gray-800 px-2 py-1 rounded inline-block">
                  {{ previewSource.no_leopard || '—' }}
                </p>
                <p class="text-[10px] text-gray-400 block mt-1">
                  Créée le {{ formatDate(previewSource.created_at) }}
                </p>
                <p v-if="previewSource.date_signature" class="text-[10px] text-emerald-600 dark:text-emerald-400 font-semibold">
                  Scellée le {{ formatDate(previewSource.date_signature) }}
                </p>
              </div>
            </div>

            <!-- Bloc personne visée -->
            <div class="mt-5 border border-gray-300 dark:border-gray-700 rounded-sm">
              <div class="bg-gray-100 dark:bg-gray-800 px-4 py-1.5 border-b border-gray-300 dark:border-gray-700">
                <p class="text-[9px] font-black tracking-[0.15em] uppercase text-gray-500 dark:text-gray-400">
                  Personne visée par l'évaluation
                </p>
              </div>
              <div class="px-4 py-3 grid grid-cols-3 gap-4">
                <div>
                  <p class="text-[9px] font-bold tracking-widest uppercase text-gray-400 mb-0.5">Dossier</p>
                  <p class="text-sm font-semibold text-gray-800 dark:text-gray-200">
                    {{ previewSource.no_leopard || '—' }}
                  </p>
                </div>
                <div v-if="previewSource.client_nom">
                  <p class="text-[9px] font-bold tracking-widest uppercase text-gray-400 mb-0.5">Client</p>
                  <p class="text-sm font-semibold text-gray-800 dark:text-gray-200">
                    {{ previewSource.client_nom }}
                  </p>
                </div>
                <div v-if="previewSource.signature_nom">
                  <p class="text-[9px] font-bold tracking-widest uppercase text-gray-400 mb-0.5">Intervenant·e</p>
                  <p class="text-sm font-semibold text-gray-800 dark:text-gray-200">
                    {{ previewSource.signature_nom }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Badge scellé -->
            <div v-if="previewSource.signature_nom"
                 class="mt-3 flex items-center gap-2 text-emerald-700 dark:text-emerald-400">
              <ShieldCheck :size="13" />
              <span class="text-[10px] font-bold tracking-wide">DOCUMENT SCELLÉ ET SIGNÉ</span>
            </div>
          </div>

          <!-- ── CORPS : SECTIONS ── -->
          <div class="px-10 py-6 space-y-0">
            <template v-for="(section, sIdx) in sectionsWithContent" :key="section.id">

              <!-- Séparateur de section -->
              <div :class="['preview-section', sIdx > 0 ? 'mt-8' : '']">

                <!-- Titre de section -->
                <div class="flex items-center gap-3 mb-4 pb-2 border-b-2 border-gray-900 dark:border-gray-100">
                  <div class="flex items-center justify-center w-6 h-6 rounded-full bg-gray-900 dark:bg-gray-100 shrink-0">
                    <component :is="iconComponent(section.icone)" :size="12"
                      class="text-white dark:text-gray-900" />
                  </div>
                  <h2 class="text-sm font-black tracking-wide uppercase text-gray-900 dark:text-gray-100">
                    {{ section.label }}
                  </h2>
                  <div class="flex-1 h-px" />
                  <span class="text-[9px] font-mono text-gray-300 dark:text-gray-600">
                    {{ sIdx + 1 }}/{{ sectionsWithContent.length }}
                  </span>
                </div>

                <!-- Champs de la section -->
                <div class="space-y-4">
                  <template v-for="field in allFieldsOf(section)" :key="field.id">
                    <div v-if="field.type !== 'info' && previewHasValue(field)">

                      <!-- Label du champ -->
                      <div class="mb-1">
                        <p class="text-[9px] font-black tracking-[0.15em] uppercase text-gray-400 dark:text-gray-500">
                          {{ field.label }}
                        </p>
                      </div>

                      <!-- TABLE -->
                      <div v-if="field.type === 'table'" class="border border-gray-200 dark:border-gray-700 overflow-hidden rounded-sm">
                        <table class="w-full text-xs">
                          <thead>
                            <tr class="bg-gray-50 dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700">
                              <th v-for="col in field.columns" :key="col.id"
                                  class="text-left px-3 py-2 text-[9px] font-black tracking-widest uppercase text-gray-500">
                                {{ col.label }}
                              </th>
                            </tr>
                          </thead>
                          <tbody>
                            <tr v-for="(row, i) in getPreviewTableRows(field.id)" :key="i"
                                class="border-b last:border-0 border-gray-100 dark:border-gray-800">
                              <td v-for="col in field.columns" :key="col.id"
                                  class="px-3 py-2 text-gray-800 dark:text-gray-200">
                                {{ row[col.id] || '—' }}
                              </td>
                            </tr>
                          </tbody>
                        </table>
                      </div>

                      <!-- CHECKBOXES -->
                      <div v-else-if="field.type === 'checkboxes'" class="flex flex-wrap gap-1.5">
                        <span
                          v-for="item in (previewPayload[field.id] || [])"
                          :key="item"
                          class="inline-flex items-center gap-1 text-[11px] bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 px-2.5 py-1 rounded-sm border border-gray-200 dark:border-gray-700 font-medium"
                        >
                          <span class="w-1.5 h-1.5 rounded-full bg-teal-500 shrink-0" />
                          {{ item }}
                        </span>
                      </div>

                      <!-- SELECT -->
                      <div v-else-if="field.type === 'select'">
                        <p class="text-sm text-gray-800 dark:text-gray-200 font-medium pl-3 border-l-2 border-teal-400">
                          {{ previewPayload[field.id] || '—' }}
                        </p>
                      </div>

                      <!-- DATE -->
                      <div v-else-if="field.type === 'date'">
                        <p class="text-sm text-gray-800 dark:text-gray-200 font-mono">
                          {{ previewPayload[field.id] || '—' }}
                        </p>
                      </div>

                      <!-- GRID CHECKBOX -->
                      <div v-else-if="field.type === 'grid_checkbox'" class="text-xs text-gray-500 italic">
                        [Grille clinique — voir document source]
                      </div>

                      <!-- OBS SELECT / SCORE RADAR -->
                      <div v-else-if="field.type === 'obs_select' || field.type === 'score_radar'" class="text-xs text-gray-500 italic">
                        [Données structurées — voir document source]
                      </div>

                      <!-- TEXTE / TEXTAREA (défaut) -->
                      <div v-else class="pl-0">
                        <p class="text-sm text-gray-800 dark:text-gray-200 whitespace-pre-wrap leading-relaxed">
                          {{ previewPayload[field.id] }}
                        </p>
                      </div>

                      <!-- Ligne de séparation entre champs -->
                      <div class="mt-4 border-b border-dashed border-gray-100 dark:border-gray-800" />

                    </div>
                  </template>
                </div>

              </div>
            </template>

            <!-- Aucun contenu -->
            <div v-if="!sectionsWithContent.length" class="py-16 text-center text-gray-400">
              <FileText :size="36" class="mx-auto mb-3 opacity-20" />
              <p class="text-sm">Aucune donnée saisie dans cette évaluation</p>
            </div>
          </div>

          <!-- PIED DE PAGE DOCUMENT -->
          <div class="border-t border-gray-200 dark:border-gray-800 px-10 py-4 mt-8">
            <div class="flex items-center justify-between">
              <p class="text-[9px] text-gray-300 dark:text-gray-600 font-mono">
                {{ currentModel?.nom }} · {{ previewSource.no_leopard }}
              </p>
              <p class="text-[9px] text-gray-300 dark:text-gray-600">
                Évaluation psychosociale — Confidentiel
              </p>
              <p class="text-[9px] text-gray-300 dark:text-gray-600 font-mono">
                {{ formatDate(new Date().toISOString()) }}
              </p>
            </div>
          </div>

        </div>
        <!-- ── FIN PAGE DOCUMENT ── -->

      </div>
    </div>

    <!-- ══════════════════════════════════════
         MODE VISUALISATION (évaluation scellée)
    ══════════════════════════════════════ -->
    <div v-else-if="selectedEvaluation" class="flex-1 overflow-y-auto">
      <div class="max-w-3xl mx-auto p-8 space-y-6">

        <!-- En-tête document -->
        <div class="rounded-2xl overflow-hidden border dark:border-gray-700 shadow-sm">
          <div class="bg-gradient-to-r from-slate-800 to-slate-700 px-6 py-5">
            <div class="flex items-start justify-between gap-4">
              <div>
                <p class="text-teal-400 font-mono text-xs mb-2">{{ selectedEvaluation.no_leopard }}</p>
                <h3 class="text-xl font-bold text-white">
                  {{ selectedEvaluation.model_id || 'Évaluation du fonctionnement social' }}
                </h3>
              </div>
              <div class="shrink-0 text-right">
                <p class="text-slate-400 text-xs">Créée le</p>
                <p class="text-white text-sm font-semibold">{{ formatDate(selectedEvaluation.created_at) }}</p>
              </div>
            </div>
          </div>
          <div v-if="selectedEvaluation.signature_nom"
               class="bg-emerald-900/20 dark:bg-emerald-900/30 px-6 py-4 border-t border-emerald-700/30">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-2 text-emerald-400">
                <ShieldCheck :size="16" />
                <span class="font-bold text-sm">{{ selectedEvaluation.signature_nom }}</span>
              </div>
              <span class="text-emerald-600 text-xs">
                Scellée le {{ formatDateTime(selectedEvaluation.date_signature) }}
              </span>
            </div>
          </div>
        </div>

        <!-- Sections -->
        <template v-if="currentModel?.schema?.sections">
          <div v-for="section in currentModel.schema.sections" :key="section.id">
            <template v-for="field in (section.fields || [])" :key="field.id">
              <div
                v-if="field.type !== 'info' && hasValue(field)"
                class="bg-gray-50 dark:bg-gray-900/40 rounded-xl border dark:border-gray-800 overflow-hidden mb-4"
              >
                <div class="flex items-center gap-2 px-5 py-3 border-b dark:border-gray-800 bg-white dark:bg-gray-900">
                  <component :is="iconComponent(section.icone)" :size="14" class="text-teal-500 shrink-0" />
                  <h4 class="text-xs font-bold text-gray-500 dark:text-gray-400 uppercase tracking-wider">{{ field.label }}</h4>
                </div>
                <div class="px-5 py-4">
                  <div v-if="field.type === 'table'" class="overflow-x-auto">
                    <table class="w-full text-xs">
                      <thead>
                        <tr class="border-b dark:border-gray-700">
                          <th v-for="col in field.columns" :key="col.id"
                              class="text-left py-2 pr-4 text-gray-500 font-semibold uppercase tracking-wider">
                            {{ col.label }}
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="(row, i) in getTableRows(field.id)" :key="i"
                            class="border-b last:border-0 dark:border-gray-800">
                          <td v-for="col in field.columns" :key="col.id"
                              class="py-2 pr-4 text-gray-700 dark:text-gray-300">
                            {{ row[col.id] || '—' }}
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                  <div v-else-if="field.type === 'checkboxes'" class="flex flex-wrap gap-2">
                    <span
                      v-for="item in (selectedEvaluation.payload?.[field.id] || [])"
                      :key="item"
                      class="inline-flex items-center gap-1.5 text-xs bg-teal-50 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300 px-2.5 py-1 rounded-full border border-teal-200 dark:border-teal-700"
                    >
                      <CheckCircle2 :size="11" />
                      {{ item }}
                    </span>
                  </div>
                  <p v-else class="text-gray-800 dark:text-gray-200 whitespace-pre-wrap leading-relaxed text-sm">
                    {{ selectedEvaluation.payload?.[field.id] }}
                  </p>
                </div>
              </div>
            </template>
          </div>
        </template>

        <div v-else-if="selectedEvaluation.payload" class="bg-gray-50 dark:bg-gray-900/40 rounded-xl p-5 border dark:border-gray-800">
          <pre class="text-xs text-gray-600 dark:text-gray-400 whitespace-pre-wrap">{{ JSON.stringify(selectedEvaluation.payload, null, 2) }}</pre>
        </div>

      </div>
    </div>

    <!-- Placeholder vide -->
    <div v-else class="flex-1 flex items-center justify-center text-gray-400">
      <div class="text-center">
        <ClipboardList :size="56" class="mx-auto mb-4 opacity-15" />
        <p class="text-base font-medium">Sélectionnez une évaluation<br>ou créez-en une nouvelle</p>
        <p class="text-xs mt-1 opacity-60">Évaluations psychosociales OTSTCFQ</p>
      </div>
    </div>

    <!-- ══════════════════════════════════════
         FOOTER BOUTONS (mode création)
    ══════════════════════════════════════ -->
    <div v-if="isCreating"
         class="shrink-0 px-8 py-4 bg-slate-50 dark:bg-gray-900/80 border-t dark:border-gray-800 flex items-center justify-between gap-4">

      <!-- Progression (masquée en aperçu) -->
      <div v-if="!isPreviewing">
        <div class="flex items-center gap-2 mb-1">
          <div class="w-28 h-1.5 bg-gray-200 dark:bg-gray-700 rounded-full overflow-hidden">
            <div
              class="h-full bg-gradient-to-r from-teal-500 to-teal-400 transition-all duration-500"
              :style="{ width: totalProgress + '%' }"
            />
          </div>
          <span class="text-xs font-bold text-teal-600">{{ totalProgress }}%</span>
        </div>
        <div class="flex items-center gap-1.5">
          <span class="w-1.5 h-1.5 rounded-full bg-orange-400 animate-pulse" />
          <span class="text-[10px] font-bold text-orange-500 uppercase tracking-wide">Brouillon</span>
        </div>
      </div>

      <!-- En mode aperçu : label -->
      <div v-else class="flex items-center gap-2 text-violet-600 dark:text-violet-400">
        <Eye :size="14" />
        <span class="text-xs font-bold">Mode aperçu — document tel qu'il sera exporté</span>
      </div>

      <!-- Boutons -->
      <div class="flex items-center gap-3">

        <!-- Bouton Aperçu / Retour édition -->
        <button
          @click="togglePreview"
          :class="[
            'flex items-center gap-2 px-4 py-2.5 rounded-xl font-bold text-sm transition-all active:scale-95 border-2',
            isPreviewing
              ? 'border-violet-500 bg-violet-50 dark:bg-violet-900/20 text-violet-700 dark:text-violet-300 hover:bg-violet-100 dark:hover:bg-violet-900/30'
              : 'border-gray-200 dark:border-gray-700 text-gray-500 dark:text-gray-400 hover:border-violet-400 hover:text-violet-600 dark:hover:text-violet-400'
          ]"
        >
          <Eye v-if="!isPreviewing" :size="15" />
          <EyeOff v-else :size="15" />
          <span>{{ isPreviewing ? 'Retour édition' : 'Aperçu' }}</span>
        </button>

        <template v-if="!isPreviewing">
          <button
            @click="$emit('cancel')"
            class="text-sm font-medium text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 transition-colors px-3 py-2"
          >
            Annuler
          </button>

          <button
            @click="$emit('save-draft')"
            :disabled="isSaving"
            class="flex items-center gap-2 bg-gray-600 hover:bg-gray-700 disabled:bg-gray-300 dark:disabled:bg-gray-700 text-white px-5 py-2.5 rounded-xl font-bold text-sm shadow transition-all active:scale-95"
          >
            <Save :size="16" />
            <span>{{ isSaving ? 'Sauvegarde...' : 'Brouillon' }}</span>
          </button>

          <button
            @click="$emit('finalize')"
            :disabled="isSaving || isFinalizing || totalProgress < 20"
            class="flex items-center gap-2 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 disabled:from-gray-400 disabled:to-gray-500 text-white px-6 py-2.5 rounded-xl font-bold text-sm shadow-lg transition-all active:scale-95"
          >
            <Loader2 v-if="isFinalizing" :size="16" class="animate-spin" />
            <ShieldCheck v-else :size="16" />
            <span>{{ isFinalizing ? 'Finalisation...' : 'Sceller et signer' }}</span>
          </button>
        </template>

        <!-- En aperçu : bouton export PDF -->
        <template v-else>
          <button
            @click="$emit('export-pdf')"
            class="flex items-center gap-2 bg-gradient-to-r from-violet-600 to-indigo-600 hover:from-violet-700 hover:to-indigo-700 text-white px-6 py-2.5 rounded-xl font-bold text-sm shadow-lg transition-all active:scale-95"
          >
            <FileDown :size="15" />
            Exporter en PDF
          </button>
        </template>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import {
  FileText, Brain, Heart, Activity, Briefcase, Users,
  Layers, CheckCircle, Target, Network, ClipboardList,
  Scale, ShieldCheck, Plus, Trash2, Save, Loader2,
  ChevronLeft, ChevronRight, CheckCircle2, Eye, EyeOff,
  FileDown
} from 'lucide-vue-next'
import SectionLayout     from './builderEval/SectionLayout.vue'
import FieldGridCheckbox from './builderEval/FieldGridCheckbox.vue'
import FieldObsSelect    from './builderEval/FieldObsSelect.vue'
import FieldScoreRadar   from './builderEval/FieldScoreRadar.vue'

// ── Props ──────────────────────────────────────────────────
const props = defineProps({
  isCreating:          Boolean,
  selectedEvaluation:  Object,
  formData:            Object,
  activeSection:       String,
  currentModel:        Object,
  isSaving:            Boolean,
  isFinalizing:        Boolean
})

const emit = defineEmits([
  'update:form-data', 'save-draft', 'finalize',
  'cancel', 'section-change', 'export-pdf'
])

// ── État local aperçu ──────────────────────────────────────
const isPreviewing = ref(false)
const togglePreview = () => { isPreviewing.value = !isPreviewing.value }

// ── Source de données pour l'aperçu ───────────────────────
// En mode création : on preview formData (brouillon en cours)
// En mode lecture  : on preview selectedEvaluation (scellée)
const previewSource = computed(() =>
  props.isCreating ? (props.formData || {}) : (props.selectedEvaluation || {})
)
const previewPayload = computed(() =>
  previewSource.value?.payload || {}
)

// ── Icônes ─────────────────────────────────────────────────
const iconMap = {
  FileText, Brain, Heart, Activity, Briefcase, Users,
  Layers, CheckCircle, Target, Network, ClipboardList,
  Scale, ShieldCheck
}
const iconComponent = (name) => iconMap[name] || FileText

// ── Section active (mode édition) ─────────────────────────
const sections = computed(() => props.currentModel?.schema?.sections || [])

const activeModelSection = computed(() =>
  sections.value.find(s => s.id === props.activeSection) || null
)

const currentSectionIndex = computed(() =>
  sections.value.findIndex(s => s.id === props.activeSection)
)

const prevSection = computed(() =>
  currentSectionIndex.value > 0 ? sections.value[currentSectionIndex.value - 1] : null
)

const nextSection = computed(() =>
  currentSectionIndex.value < sections.value.length - 1
    ? sections.value[currentSectionIndex.value + 1]
    : null
)

// ── Récupère tous les champs d'une section (colonnes ou plat)
const allFieldsOf = (section) => {
  if (section.columns?.length) {
    return section.columns.flatMap(col => col.fields || [])
  }
  return section.fields || []
}

// ── Sections qui ont au moins un champ rempli (aperçu) ────
const sectionsWithContent = computed(() =>
  sections.value.filter(section =>
    allFieldsOf(section).some(f => f.type !== 'info' && previewHasValue(f))
  )
)

// ── Mise à jour du payload ─────────────────────────────────
const updatePayload = (fieldId, value) => {
  emit('update:form-data', {
    ...props.formData,
    payload: { ...(props.formData?.payload || {}), [fieldId]: value }
  })
}

// ── Vérification valeur (mode aperçu) ─────────────────────
const previewHasValue = (field) => {
  const val = previewPayload.value[field.id]
  if (!val) return false
  if (typeof val === 'string') return val.trim().length > 0
  if (Array.isArray(val)) return val.length > 0
  if (typeof val === 'object') return Object.keys(val).length > 0
  return !!val
}

// ── Vérification valeur (mode visualisation scellée) ──────
const hasValue = (field) => {
  const val = props.selectedEvaluation?.payload?.[field.id]
  if (!val) return false
  if (typeof val === 'string') return val.trim().length > 0
  if (Array.isArray(val)) return val.length > 0
  return !!val
}

// ── Tables (mode aperçu) ───────────────────────────────────
const getPreviewTableRows = (fieldId) => {
  const val = previewPayload.value[fieldId]
  return Array.isArray(val) ? val : []
}

// ── Tables (mode visualisation scellée) ───────────────────
const getTableRows = (fieldId) => {
  const val = props.selectedEvaluation?.payload?.[fieldId]
  return Array.isArray(val) ? val : []
}

// ── Progress ───────────────────────────────────────────────
const totalProgress = computed(() => {
  if (!sections.value.length || !props.formData?.payload) return 0
  let total = 0, filled = 0
  sections.value.forEach(section => {
    allFieldsOf(section).forEach(field => {
      if (field.type === 'info') return
      total++
      const val = props.formData.payload[field.id]
      if (!val) return
      if (typeof val === 'string' && val.trim()) filled++
      else if (Array.isArray(val) && val.length) filled++
      else if (typeof val === 'object' && Object.keys(val).length) filled++
    })
  })
  return total ? Math.round((filled / total) * 100) : 0
})

// ── Formatage dates ────────────────────────────────────────
const formatDate = (d) => d
  ? new Date(d).toLocaleDateString('fr-CA', { year: 'numeric', month: 'long', day: 'numeric' })
  : '—'

const formatDateTime = (d) => d
  ? new Date(d).toLocaleString('fr-CA', { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' })
  : '—'
</script>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to   { opacity: 1; transform: translateY(0); }
}
.animate-fadeIn { animation: fadeIn 0.25s ease-out; }
.animate-spin   { animation: spin 1s linear infinite; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
.animate-pulse  { animation: pulse 2s cubic-bezier(0.4,0,0.6,1) infinite; }
@keyframes pulse { 0%,100% { opacity:1; } 50% { opacity:.5; } }

/* Styles impression pour le PDF */
@media print {
  .preview-section { page-break-inside: avoid; }
}
</style>