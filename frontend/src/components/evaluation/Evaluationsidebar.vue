<template>
  <div class="w-64 border-r dark:border-gray-800 bg-slate-50 dark:bg-gray-900/30 flex flex-col shrink-0">

    <!-- En-tête sidebar -->
    <div class="px-4 py-3 border-b dark:border-gray-800 bg-slate-100/70 dark:bg-gray-900/60">
      <span class="text-[10px] font-bold text-gray-400 uppercase tracking-widest">
        {{ isCreating ? 'Sections' : 'Évaluations finalisées' }}
      </span>
    </div>

    <!-- MODE CRÉATION : Navigation sections -->
    <div v-if="isCreating" class="flex-1 overflow-y-auto p-3 space-y-1">
      <button
        v-for="section in sections"
        :key="section.id"
        @click="$emit('section-change', section.id)"
        :class="[
          'w-full flex items-center gap-3 px-3 py-2.5 rounded-xl transition-all text-left group',
          activeSection === section.id
            ? 'bg-teal-600 text-white shadow-md shadow-teal-200 dark:shadow-teal-900/50'
            : 'hover:bg-white dark:hover:bg-gray-800 text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'
        ]"
      >
        <component :is="section.icon" :size="16" class="shrink-0" />
        <span class="flex-1 font-medium text-xs">{{ section.label }}</span>
        <!-- Dot si remplie -->
        <div
          :class="[
            'w-1.5 h-1.5 rounded-full shrink-0 transition-all',
            activeSection === section.id ? 'bg-white/60' : (filledSections.has(section.id) ? 'bg-teal-500' : 'bg-gray-200 dark:bg-gray-700')
          ]"
        />
      </button>
    </div>

    <!-- MODE VISUALISATION : Liste évaluations -->
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
        <!-- Type badge -->
        <div class="flex items-center justify-between mb-1.5">
          <span :class="['text-[9px] font-bold px-1.5 py-0.5 rounded-full uppercase tracking-wide', typeClass(ev.type_evaluation)]">
            {{ typeLabel(ev.type_evaluation) }}
          </span>
          <div class="flex items-center gap-1">
            <Lock v-if="ev.verrouille" :size="11" class="text-emerald-500" />
            <span class="text-[10px] text-gray-400">{{ formatDate(ev.created_at) }}</span>
          </div>
        </div>

        <!-- Nom Léopard -->
        <p v-if="ev.nom_leopard" class="text-[9px] font-mono text-teal-600 dark:text-teal-400 mb-1 truncate">
          {{ ev.nom_leopard }}
        </p>

        <!-- Objet/Titre -->
        <p class="text-xs font-semibold text-gray-800 dark:text-gray-200 truncate leading-snug">
          {{ ev.objet_evaluation || ev.motif_reference || 'Évaluation sans titre' }}
        </p>

        <!-- Signataire -->
        <p v-if="ev.signature_nom" class="text-[10px] text-gray-400 mt-1 truncate italic">
          {{ ev.signature_nom }}
        </p>
      </div>

      <!-- Vide -->
      <div v-if="!evaluations.length" class="text-center py-10 text-gray-400">
        <ShieldOff :size="36" class="mx-auto mb-3 opacity-20" />
        <p class="text-xs">Aucune évaluation<br>finalisée</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Lock, ShieldOff } from 'lucide-vue-next'

const props = defineProps({
  evaluations:         Array,
  selectedEvaluation:  Object,
  isCreating:          Boolean,
  activeSection:       String,
  sections:            Array,
  formData:            Object
})

defineEmits(['section-change', 'view-evaluation'])

// Détecte les sections déjà remplies (dot vert)
const filledSections = computed(() => {
  if (!props.formData) return new Set()
  const fieldToSection = {
    contexte_evaluation: 'contexte', motif_reference: 'contexte', objet_evaluation: 'contexte',
    capacites_cognitives: 'cognitif', inaptitude_constatee: 'inaptitude',
    etat_sante_physique: 'sante',
    dimensions_psycho_sociales: 'psycho',
    roles_sociaux: 'roles',
    reseau_social_soutien: 'reseau',
    evaluation_mandataire: 'mandataire',
    evolution_situation: 'evolution',
    objectifs_intervention: 'intervention',
    plan_action: 'plan',
    situation_globale: 'situation',
    analyse_clinique: 'analyse',
    opinion_professionnelle: 'opinion',
    recommandations: 'recommandations'
  }
  const filled = new Set()
  Object.entries(props.formData).forEach(([key, val]) => {
    if (val?.trim?.() && fieldToSection[key]) filled.add(fieldToSection[key])
  })
  return filled
})

const typeLabel = (t) => {
  const m = { tutelle: 'Tutelle', mandat: 'Mandat', conseil_tutelle: 'Conseil', suivi_psychosocial: 'Suivi' }
  return m[t] || 'Éval.'
}

const typeClass = (t) => {
  const m = {
    tutelle:            'bg-blue-100 text-blue-700 dark:bg-blue-900/40 dark:text-blue-300',
    mandat:             'bg-purple-100 text-purple-700 dark:bg-purple-900/40 dark:text-purple-300',
    conseil_tutelle:    'bg-amber-100 text-amber-700 dark:bg-amber-900/40 dark:text-amber-300',
    suivi_psychosocial: 'bg-teal-100 text-teal-700 dark:bg-teal-900/40 dark:text-teal-300'
  }
  return m[t] || 'bg-gray-100 text-gray-600'
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA', { month: 'short', day: 'numeric', year: '2-digit' })
}
</script>