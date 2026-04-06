<template>
  <div class="w-64 border-r dark:border-gray-800 bg-slate-50 dark:bg-gray-900/30 flex flex-col shrink-0">

    <!-- En-tête -->
    <div class="px-4 py-3 border-b dark:border-gray-800 bg-slate-100/70 dark:bg-gray-900/60">
      <span class="text-[10px] font-bold text-gray-400 uppercase tracking-widest">
        {{ isCreating ? (currentModel?.nom || 'Sections') : 'Évaluations finalisées' }}
      </span>
    </div>

    <!-- MODE CRÉATION : navigation par sections du modèle -->
    <div v-if="isCreating" class="flex-1 overflow-y-auto p-2 space-y-0.5">

      <!-- Pas de modèle chargé -->
      <div v-if="!currentModel?.schema?.sections?.length" class="px-3 py-6 text-center text-gray-400">
        <p class="text-xs">Aucune section disponible</p>
      </div>

      <button
        v-else
        v-for="section in currentModel.schema.sections"
        :key="section.id"
        @click="$emit('section-change', section.id)"
        :class="[
          'w-full flex items-center gap-2.5 px-3 py-2.5 rounded-lg text-left transition-all text-sm',
          activeSection === section.id
            ? 'bg-teal-600 text-white shadow-sm'
            : 'text-gray-600 dark:text-gray-400 hover:bg-white dark:hover:bg-gray-800 hover:text-gray-900 dark:hover:text-white'
        ]"
      >
        <!-- Icône de la section -->
        <component
          :is="iconComponent(section.icone)"
          :size="14"
          :class="activeSection === section.id ? 'text-white' : sectionIconColor(section.couleur)"
          class="shrink-0"
        />

        <!-- Label -->
        <span class="flex-1 truncate font-medium text-xs leading-tight">{{ section.label }}</span>

        <!-- Dot vert si section remplie -->
        <span
          v-if="isSectionFilled(section)"
          :class="['w-1.5 h-1.5 rounded-full shrink-0', activeSection === section.id ? 'bg-white/70' : 'bg-emerald-400']"
        />
      </button>
    </div>

    <!-- MODE VISUALISATION : liste des évaluations finalisées -->
    <div v-else class="flex-1 overflow-y-auto p-3 space-y-2">

      <div
        v-for="ev in evaluations"
        :key="ev.id"
        @click="$emit('view-evaluation', ev)"
        :class="[
          'p-3 rounded-xl border cursor-pointer transition-all',
          selectedEvaluation?.id === ev.id
            ? 'ring-2 ring-teal-500 border-teal-200 dark:border-teal-800 bg-teal-50 dark:bg-teal-900/20'
            : 'border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900 hover:border-teal-300 dark:hover:border-teal-700 hover:shadow-sm'
        ]"
      >
        <!-- Badges -->
        <div class="flex items-center justify-between mb-1.5 gap-1">
          <span class="text-[9px] font-bold px-1.5 py-0.5 rounded-full uppercase tracking-wide bg-teal-100 text-teal-700 dark:bg-teal-900/40 dark:text-teal-300 truncate">
            {{ ev.model_id || 'Éval.' }}
          </span>
          <div class="flex items-center gap-1 shrink-0">
            <Lock v-if="ev.statut === 'verrouille'" :size="11" class="text-emerald-500" />
            <span class="text-[10px] text-gray-400">{{ formatDate(ev.created_at) }}</span>
          </div>
        </div>

        <!-- No Léopard -->
        <p v-if="ev.no_leopard" class="text-[9px] font-mono text-teal-600 dark:text-teal-400 mb-1 truncate">
          {{ ev.no_leopard }}
        </p>

        <!-- Signataire -->
        <p v-if="ev.signature_nom" class="text-[10px] text-gray-400 mt-1 truncate italic">
          ✓ {{ ev.signature_nom }}
        </p>
      </div>

      <!-- Vide -->
      <div v-if="!evaluations?.length" class="text-center py-10 text-gray-400">
        <ShieldOff :size="36" class="mx-auto mb-3 opacity-20" />
        <p class="text-xs">Aucune évaluation<br>finalisée</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  Lock, ShieldOff,
  FileText, Brain, Heart, Activity, Briefcase, Users,
  Layers, CheckCircle, Target, Network, ClipboardList,
  Scale, ShieldCheck
} from 'lucide-vue-next'

const props = defineProps({
  evaluations:         Array,
  selectedEvaluation:  Object,
  isCreating:          Boolean,
  activeSection:       String,
  currentModel:        Object,   // ← Le modèle actif avec schema parsé
  formData:            Object
})

defineEmits(['section-change', 'view-evaluation'])

// ── Icônes dynamiques ───────────────────────────────────────
const iconMap = {
  FileText, Brain, Heart, Activity, Briefcase, Users,
  Layers, CheckCircle, Target, Network, ClipboardList,
  Scale, ShieldCheck
}
const iconComponent = (name) => iconMap[name] || FileText

const colorIconMap = {
  blue:    'text-blue-500',
  teal:    'text-teal-500',
  emerald: 'text-emerald-500',
  purple:  'text-purple-500',
  amber:   'text-amber-500',
  red:     'text-red-500',
  slate:   'text-slate-400',
  indigo:  'text-indigo-500',
}
const sectionIconColor = (c) => colorIconMap[c] || 'text-gray-400'

// ── Section remplie ? (vérifie le payload) ─────────────────
const isSectionFilled = (section) => {
  if (!props.formData?.payload || !section?.fields) return false
  return section.fields.some(field => {
    const val = props.formData.payload[field.id]
    if (!val) return false
    if (typeof val === 'string') return val.trim().length > 0
    if (Array.isArray(val)) return val.length > 0
    if (typeof val === 'object') return Object.keys(val).length > 0
    return !!val
  })
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA', { month: 'short', day: 'numeric', year: '2-digit' })
}
</script>