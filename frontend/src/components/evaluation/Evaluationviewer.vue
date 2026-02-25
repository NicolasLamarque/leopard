<template>
  <div class="flex-1 flex flex-col bg-white dark:bg-gray-950 overflow-hidden">

    <!-- ===== MODE CRÉATION/ÉDITION ===== -->
    <div v-if="isCreating" class="flex-1 overflow-y-auto">
      <div class="max-w-3xl mx-auto p-8 space-y-6">

        <!-- Type banner -->
        <div :class="['flex items-center gap-3 px-4 py-3 rounded-xl border text-sm font-semibold', typeBannerClass]">
          <component :is="typeIcon" :size="18" />
          <span>{{ typeLabel }}</span>
          <span class="ml-auto font-mono text-xs opacity-70">{{ formData.nom_leopard }}</span>
        </div>

        <!-- ── SECTIONS COMMUNES ── -->

        <!-- CONTEXTE -->
        <div v-show="activeSection === 'contexte'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="FileText" label="Contexte et Référence" />

          <FormField
            label="Contexte de l'évaluation"
            :required="true"
            :error="errors.contexte_evaluation"
            help="Circonstances ayant mené à la demande d'évaluation"
          >
            <textarea
              :value="formData.contexte_evaluation"
              @input="update('contexte_evaluation', $event.target.value)"
              :class="fieldClass(errors.contexte_evaluation)"
              rows="5"
              placeholder="Décrire les circonstances: qui référait, pourquoi, dans quel contexte juridique..."
            />
          </FormField>

          <FormField
            label="Motif de référence"
            :required="true"
            :error="errors.motif_reference"
            help="Raison principale de la demande"
          >
            <textarea
              :value="formData.motif_reference"
              @input="update('motif_reference', $event.target.value)"
              :class="fieldClass(errors.motif_reference)"
              rows="3"
              placeholder="Motif principal de la référence..."
            />
          </FormField>

          <FormField label="Objet de l'évaluation" help="Ce qui doit spécifiquement être évalué ou déterminé">
            <textarea
              :value="formData.objet_evaluation"
              @input="update('objet_evaluation', $event.target.value)"
              :class="fieldClass()"
              rows="3"
              placeholder="Préciser l'objet précis de l'évaluation..."
            />
          </FormField>
        </div>

        <!-- COGNITIF (tutelle / mandat) -->
        <div v-show="activeSection === 'cognitif'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Brain" label="Capacités Cognitives" />
          <FormField label="Observation des capacités cognitives" help="Orientation, mémoire, attention, jugement, langage, praxies">
            <textarea
              :value="formData.capacites_cognitives"
              @input="update('capacites_cognitives', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Décrire les observations: orientation temporo-spatiale, mémoire à court et long terme, capacité de jugement et de compréhension, langage, reconnaissance..."
            />
          </FormField>
        </div>

        <!-- INAPTITUDE (mandat) -->
        <div v-show="activeSection === 'inaptitude'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Brain" label="Constat d'inaptitude" />
          <div class="p-4 bg-purple-50 dark:bg-purple-900/20 border border-purple-200 dark:border-purple-800 rounded-xl text-sm text-purple-700 dark:text-purple-300">
            <strong>Note :</strong> Pour l'homologation de mandat, l'inaptitude doit être constatée par deux professionnels (dont un médecin). Documenter votre constat clinique ici.
          </div>
          <FormField label="Inaptitude constatée — observations cliniques" help="Votre constat clinique d'inaptitude (partielle ou totale)">
            <textarea
              :value="formData.inaptitude_constatee"
              @input="update('inaptitude_constatee', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Décrire le niveau d'inaptitude constatée, sa nature (totale/partielle), son évolution attendue..."
            />
          </FormField>
        </div>

        <!-- SANTÉ -->
        <div v-show="activeSection === 'sante'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Heart" label="État de Santé Physique" />
          <FormField label="Diagnostics médicaux, limitations fonctionnelles, autonomie" help="État physique global, maladies chroniques, mobilité, AVQ/AVD">
            <textarea
              :value="formData.etat_sante_physique"
              @input="update('etat_sante_physique', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Diagnostics connus, médication, limitations dans les activités de la vie quotidienne et domestique, mobilité, autonomie résiduelle..."
            />
          </FormField>
        </div>

        <!-- PSYCHO-SOCIAL -->
        <div v-show="activeSection === 'psycho'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Activity" label="Dimensions Psycho-sociales" />
          <FormField label="Humeur, émotions, comportements, adaptation" help="État émotionnel, traits de personnalité, mécanismes d'adaptation">
            <textarea
              :value="formData.dimensions_psycho_sociales"
              @input="update('dimensions_psycho_sociales', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Humeur générale, anxiété, dépression, comportements, anosognosie, conscience des limitations, capacité d'adaptation..."
            />
          </FormField>
        </div>

        <!-- RÔLES SOCIAUX -->
        <div v-show="activeSection === 'roles'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Briefcase" label="Rôles Sociaux" />
          <FormField label="Activités quotidiennes, loisirs, implications sociales" help="Maintien des rôles familiaux, professionnels, communautaires">
            <textarea
              :value="formData.roles_sociaux"
              @input="update('roles_sociaux', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Implication familiale, occupations quotidiennes, loisirs, bénévolat, participation sociale, gestion administrative et financière..."
            />
          </FormField>
        </div>

        <!-- MANDATAIRE (mandat) -->
        <div v-show="activeSection === 'mandataire'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="UserCheck" label="Évaluation du Mandataire" />
          <div class="p-4 bg-purple-50 dark:bg-purple-900/20 border border-purple-200 dark:border-purple-800 rounded-xl text-sm text-purple-700 dark:text-purple-300">
            <strong>Important :</strong> L'évaluation de la capacité du mandataire à assumer les responsabilités est un élément essentiel du rapport pour homologation.
          </div>
          <FormField label="Identification et évaluation du mandataire" help="Relation avec le mandant, capacité à assumer le mandat, acceptation">
            <textarea
              :value="formData.evaluation_mandataire"
              @input="update('evaluation_mandataire', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Identité du mandataire désigné, relation avec le mandant, évaluation de sa capacité à assumer les responsabilités, acceptation du rôle, ressources disponibles..."
            />
          </FormField>
        </div>

        <!-- RÉSEAU SOUTIEN -->
        <div v-show="activeSection === 'reseau'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Users" label="Réseau Social et Soutien" />
          <FormField label="Famille, proches, soutien disponible" help="Qualité et disponibilité du réseau de soutien naturel et professionnel">
            <textarea
              :value="formData.reseau_social_soutien"
              @input="update('reseau_social_soutien', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Personnes significatives, qualité des relations, implication du réseau, services professionnels déjà en place, besoins identifiés..."
            />
          </FormField>
        </div>

        <!-- ÉVOLUTION (suivi psychosocial) -->
        <div v-show="activeSection === 'evolution'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="TrendingUp" label="Évolution de la Situation" />
          <FormField label="Évolution depuis la dernière note" help="Changements observés, progrès, nouvelles problématiques">
            <textarea
              :value="formData.evolution_situation"
              @input="update('evolution_situation', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Décrire l'évolution de la situation depuis le dernier contact: progrès, régressions, nouvelles problématiques, événements significatifs..."
            />
          </FormField>
        </div>

        <!-- INTERVENTION (suivi) -->
        <div v-show="activeSection === 'intervention'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Briefcase" label="Interventions Réalisées" />
          <FormField label="Nature des interventions et objectifs travaillés" help="Ce qui a été fait, avec qui, pour quoi">
            <textarea
              :value="formData.objectifs_intervention"
              @input="update('objectifs_intervention', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Décrire les interventions réalisées, les objectifs travaillés, les approches utilisées, les ressources mobilisées..."
            />
          </FormField>
        </div>

        <!-- SITUATION GLOBALE (conseil tutelle) -->
        <div v-show="activeSection === 'situation'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Activity" label="Situation Globale" />
          <FormField label="Présentation de la situation globale" help="Portrait complet de la personne et de sa situation">
            <textarea
              :value="formData.situation_globale"
              @input="update('situation_globale', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Présentation complète de la situation: contexte familial, histoire, difficultés actuelles, ressources..."
            />
          </FormField>
        </div>

        <!-- ANALYSE CLINIQUE -->
        <div v-show="activeSection === 'analyse'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Layers" label="Analyse Clinique" />
          <FormField label="Synthèse et analyse professionnelle" help="Intégration des observations, facteurs de protection et de risque">
            <textarea
              :value="formData.analyse_clinique"
              @input="update('analyse_clinique', $event.target.value)"
              :class="fieldClass()"
              rows="9"
              placeholder="Analyse intégrée des dimensions évaluées, facteurs de protection et de risque, dynamiques familiales et sociales, diagnostic fonctionnel..."
            />
          </FormField>
        </div>

        <!-- OPINION PROFESSIONNELLE -->
        <div v-show="activeSection === 'opinion'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="CheckCircle" label="Opinion Professionnelle" />
          <div class="p-4 bg-slate-50 dark:bg-slate-900/50 border border-slate-200 dark:border-slate-700 rounded-xl text-sm text-slate-600 dark:text-slate-400">
            <strong>Note OTSTCFQ :</strong> L'opinion professionnelle doit être clairement formulée, fondée sur les observations et l'analyse, et distinguée des faits rapportés.
          </div>
          <FormField label="Opinion professionnelle du T.S." help="Votre opinion clinique fondée sur l'évaluation">
            <textarea
              :value="formData.opinion_professionnelle"
              @input="update('opinion_professionnelle', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="En ma qualité de travailleur social, mon opinion professionnelle est que... (formuler clairement votre opinion sur la situation, l'inaptitude, le besoin de protection)..."
            />
          </FormField>
        </div>

        <!-- RECOMMANDATIONS -->
        <div v-show="activeSection === 'recommandations'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Target" label="Recommandations" />
          <FormField label="Recommandations et plan d'intervention" help="Actions recommandées, ressources à mobiliser, suivi préconisé">
            <textarea
              :value="formData.recommandations"
              @input="update('recommandations', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Sur la base de mon évaluation, je recommande: (régime de protection approprié, mesures à mettre en place, ressources recommandées, suivi préconisé)..."
            />
          </FormField>
        </div>

        <!-- PLAN D'ACTION (suivi) -->
        <div v-show="activeSection === 'plan'" class="space-y-5 animate-fadeIn">
          <SectionTitle icon="Target" label="Plan d'Action" />
          <FormField label="Objectifs et plan d'intervention" help="Prochaines étapes, objectifs mesurables, échéancier">
            <textarea
              :value="formData.plan_action"
              @input="update('plan_action', $event.target.value)"
              :class="fieldClass()"
              rows="7"
              placeholder="Prochains objectifs d'intervention, stratégies planifiées, ressources à mobiliser, échéancier, prochain rendez-vous..."
            />
          </FormField>
        </div>

      </div>
    </div>

    <!-- ===== MODE VISUALISATION (évaluation finalisée) ===== -->
    <div v-else-if="selectedEvaluation" class="flex-1 overflow-y-auto">
      <div class="max-w-3xl mx-auto p-8 space-y-6">

        <!-- En-tête doc -->
        <div class="rounded-2xl overflow-hidden border dark:border-gray-700 shadow-sm">
          <div class="bg-gradient-to-r from-slate-800 to-slate-700 px-6 py-5">
            <div class="flex items-start justify-between gap-4">
              <div>
                <div class="flex items-center gap-2 mb-2">
                  <span :class="['text-[10px] font-bold px-2 py-1 rounded-full uppercase tracking-wide', typeClass(selectedEvaluation.type_evaluation)]">
                    {{ typeLabel2(selectedEvaluation.type_evaluation) }}
                  </span>
                </div>
                <h3 class="text-xl font-bold text-white mb-1">
                  {{ selectedEvaluation.objet_evaluation || selectedEvaluation.motif_reference || 'Évaluation du fonctionnement social' }}
                </h3>
                <p v-if="selectedEvaluation.nom_leopard" class="text-teal-400 font-mono text-xs">
                  {{ selectedEvaluation.nom_leopard }}
                </p>
              </div>
              <div class="shrink-0 text-right">
                <p class="text-slate-400 text-xs">Créée le</p>
                <p class="text-white text-sm font-semibold">{{ formatDate(selectedEvaluation.created_at) }}</p>
              </div>
            </div>
          </div>

          <!-- Signature -->
          <div v-if="selectedEvaluation.verrouille && selectedEvaluation.signature_nom"
               class="bg-emerald-900/20 dark:bg-emerald-900/30 px-6 py-4 border-t border-emerald-700/30">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-2 text-emerald-400">
                <ShieldCheck :size="16" />
                <span class="font-bold text-sm">{{ selectedEvaluation.signature_nom }}</span>
              </div>
              <span class="text-emerald-600 text-xs">Scellée le {{ formatDateTime(selectedEvaluation.date_signature) }}</span>
            </div>
          </div>
        </div>

        <!-- Sections du document -->
        <div
          v-for="section in docSections"
          :key="section.key"
        >
          <div
            v-if="selectedEvaluation[section.key]?.trim()"
            class="bg-gray-50 dark:bg-gray-900/40 rounded-xl border dark:border-gray-800 overflow-hidden"
          >
            <div class="flex items-center gap-2 px-5 py-3 border-b dark:border-gray-800 bg-white dark:bg-gray-900">
              <component :is="section.icon" :size="14" class="text-teal-500 shrink-0" />
              <h4 class="text-xs font-bold text-gray-500 dark:text-gray-400 uppercase tracking-wider">{{ section.label }}</h4>
            </div>
            <p class="px-5 py-4 text-gray-800 dark:text-gray-200 whitespace-pre-wrap leading-relaxed text-sm">
              {{ selectedEvaluation[section.key] }}
            </p>
          </div>
        </div>

      </div>
    </div>

    <!-- Placeholder vide -->
    <div v-else class="flex-1 flex items-center justify-center text-gray-400">
      <div class="text-center">
        <Scale :size="56" class="mx-auto mb-4 opacity-15" />
        <p class="text-base font-medium">Sélectionnez une évaluation ou créez-en une nouvelle</p>
        <p class="text-xs mt-1 opacity-60">Tutelle · Mandat · Suivi psychosocial</p>
      </div>
    </div>

    <!-- ===== FOOTER BOUTONS (mode création) ===== -->
    <div v-if="isCreating" class="shrink-0 px-8 py-4 bg-slate-50 dark:bg-gray-900/80 border-t dark:border-gray-800 flex items-center justify-between gap-4">

      <!-- Progress -->
      <div class="flex items-center gap-4">
        <div>
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
            <span class="text-[10px] font-bold text-orange-500 uppercase tracking-wide">Brouillon non scellé</span>
          </div>
        </div>
      </div>

      <!-- Boutons -->
      <div class="flex items-center gap-3">
        <button
          @click="$emit('cancel')"
          class="text-sm font-medium text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 transition-colors px-3 py-2"
        >
          Annuler
        </button>

        <button
          @click="$emit('save-draft')"
          :disabled="isSaving || totalProgress < 15"
          class="flex items-center gap-2 bg-gray-600 hover:bg-gray-700 disabled:bg-gray-300 dark:disabled:bg-gray-700 text-white px-5 py-2.5 rounded-xl font-bold text-sm shadow transition-all active:scale-95"
        >
          <Save :size="16" />
          <span>{{ isSaving ? 'Sauvegarde...' : 'Brouillon' }}</span>
        </button>

        <button
          @click="$emit('finalize')"
          :disabled="isSaving || isFinalizing || totalProgress < 40"
          class="flex items-center gap-2 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 disabled:from-gray-400 disabled:to-gray-500 text-white px-6 py-2.5 rounded-xl font-bold text-sm shadow-lg transition-all active:scale-95"
        >
          <Loader2 v-if="isFinalizing" :size="16" class="animate-spin" />
          <ShieldCheck v-else :size="16" />
          <span>{{ isFinalizing ? 'Finalisation...' : 'Sceller et signer' }}</span>
        </button>
      </div>
    </div>

  </div>
</template>

<script setup>
import { computed } from 'vue'
import { FileText, Brain, Heart, Activity, Briefcase, Users, Layers, CheckCircle, Target, ShieldCheck, Save, Scale, Loader2, UserCheck, TrendingUp } from 'lucide-vue-next'

const props = defineProps({
  isCreating:         Boolean,
  selectedEvaluation: Object,
  formData:           Object,
  errors:             Object,
  activeSection:      String,
  sections:           Array,
  totalProgress:      Number,
  isSaving:           Boolean,
  isFinalizing:       Boolean
})

const emit = defineEmits(['update:form-data', 'save-draft', 'finalize', 'cancel'])

const update = (field, value) => {
  emit('update:form-data', { ...props.formData, [field]: value })
}

const fieldClass = (error) => {
  const base = 'w-full px-4 py-3 rounded-xl border-2 dark:bg-gray-900 dark:text-white transition-all focus:outline-none focus:ring-2 text-sm leading-relaxed resize-none'
  return error
    ? `${base} border-red-300 dark:border-red-700 focus:border-red-500 focus:ring-red-100 dark:focus:ring-red-900/30`
    : `${base} border-gray-200 dark:border-gray-700 focus:border-teal-500 focus:ring-teal-100 dark:focus:ring-teal-900/30`
}

// Type meta pour le viewer
const typeMeta = {
  tutelle:            { label: 'Régime de tutelle',       color: 'bg-blue-600',   banner: 'bg-blue-50 border-blue-200 text-blue-700 dark:bg-blue-900/20 dark:border-blue-800 dark:text-blue-300', icon: Scale },
  mandat:             { label: 'Homologation de mandat',   color: 'bg-purple-600', banner: 'bg-purple-50 border-purple-200 text-purple-700 dark:bg-purple-900/20 dark:border-purple-800 dark:text-purple-300', icon: FileText },
  conseil_tutelle:    { label: 'Conseil de tutelle',       color: 'bg-amber-600',  banner: 'bg-amber-50 border-amber-200 text-amber-700 dark:bg-amber-900/20 dark:border-amber-800 dark:text-amber-300', icon: UserCheck },
  suivi_psychosocial: { label: 'Suivi psychosocial',       color: 'bg-teal-600',   banner: 'bg-teal-50 border-teal-200 text-teal-700 dark:bg-teal-900/20 dark:border-teal-800 dark:text-teal-300', icon: Activity }
}

const typeBannerClass = computed(() => {
  const m = typeMeta[props.formData?.type_evaluation]
  return m ? `${m.banner} border` : 'bg-slate-50 border-slate-200 text-slate-600'
})

const typeLabel = computed(() => typeMeta[props.formData?.type_evaluation]?.label || 'Évaluation')
const typeIcon  = computed(() => typeMeta[props.formData?.type_evaluation]?.icon || Scale)

const typeLabel2 = (t) => typeMeta[t]?.label || 'Évaluation'
const typeClass  = (t) => {
  const m = {
    tutelle:            'bg-blue-500/20 text-blue-300',
    mandat:             'bg-purple-500/20 text-purple-300',
    conseil_tutelle:    'bg-amber-500/20 text-amber-300',
    suivi_psychosocial: 'bg-teal-500/20 text-teal-300'
  }
  return m[t] || 'bg-gray-500/20 text-gray-300'
}

// Toutes les sections possibles pour le viewer read-only
const iconMap = { FileText, Brain, Heart, Activity, Briefcase, Users, Layers, CheckCircle, Target, Scale, UserCheck, TrendingUp }
const docSections = [
  { label: 'Contexte de l\'évaluation',      key: 'contexte_evaluation',       icon: FileText },
  { label: 'Motif de référence',             key: 'motif_reference',           icon: FileText },
  { label: 'Objet de l\'évaluation',          key: 'objet_evaluation',          icon: FileText },
  { label: 'Capacités cognitives',           key: 'capacites_cognitives',      icon: Brain },
  { label: 'Inaptitude constatée',           key: 'inaptitude_constatee',      icon: Brain },
  { label: 'État de santé physique',         key: 'etat_sante_physique',       icon: Heart },
  { label: 'Dimensions psycho-sociales',     key: 'dimensions_psycho_sociales',icon: Activity },
  { label: 'Rôles sociaux',                  key: 'roles_sociaux',             icon: Briefcase },
  { label: 'Évaluation du mandataire',       key: 'evaluation_mandataire',     icon: UserCheck },
  { label: 'Réseau social et soutien',       key: 'reseau_social_soutien',     icon: Users },
  { label: 'Évolution de la situation',      key: 'evolution_situation',       icon: TrendingUp },
  { label: 'Interventions réalisées',        key: 'objectifs_intervention',    icon: Briefcase },
  { label: 'Situation globale',              key: 'situation_globale',         icon: Activity },
  { label: 'Analyse clinique',               key: 'analyse_clinique',          icon: Layers },
  { label: 'Opinion professionnelle',        key: 'opinion_professionnelle',   icon: CheckCircle },
  { label: 'Recommandations',                key: 'recommandations',           icon: Target },
  { label: 'Plan d\'action',                  key: 'plan_action',               icon: Target }
]

const formatDate = (d) => d ? new Date(d).toLocaleDateString('fr-CA', { year: 'numeric', month: 'long', day: 'numeric' }) : '-'
const formatDateTime = (d) => d ? new Date(d).toLocaleString('fr-CA', { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' }) : '-'
</script>

<!-- SectionTitle (inline component via defineComponent not needed — just use locally) -->
<script>
// Sub-composant inline SectionTitle + FormField définis comme composants globaux dans main.js
// ou on les met ici comme petit helper via defineComponent
</script>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to   { opacity: 1; transform: translateY(0); }
}
.animate-fadeIn { animation: fadeIn 0.25s ease-out; }

.animate-spin { animation: spin 1s linear infinite; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4,0,0.6,1) infinite;
}
@keyframes pulse { 0%,100% { opacity:1; } 50% { opacity:.5; } }
</style>