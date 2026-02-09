<template>
  <div class="flex-1 flex flex-col bg-white dark:bg-gray-950 overflow-hidden">
    
    <!-- MODE CRÉATION/ÉDITION -->
    <div v-if="isCreating" class="flex-1 overflow-y-auto p-8">
      <div class="max-w-4xl mx-auto space-y-8">

        <!-- Section Contexte -->
        <div v-show="activeSection === 'contexte'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <FileText class="text-slate-600" />
            Contexte et Référence
          </h3>

          <div class="space-y-4">
            <div>
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                Contexte de l'évaluation * 
                <span v-if="errors.contexte_evaluation" class="text-red-500 text-xs ml-2">
                  ({{ errors.contexte_evaluation }})
                </span>
              </label>
              <textarea
                :value="formData.contexte_evaluation"
                @input="updateField('contexte_evaluation', $event.target.value)"
                rows="4"
                :class="getInputClass(errors.contexte_evaluation)"
                placeholder="Décrire les circonstances ayant mené à cette évaluation..."
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                Motif de référence *
                <span v-if="errors.motif_reference" class="text-red-500 text-xs ml-2">
                  ({{ errors.motif_reference }})
                </span>
              </label>
              <textarea
                :value="formData.motif_reference"
                @input="updateField('motif_reference', $event.target.value)"
                rows="3"
                :class="getInputClass(errors.motif_reference)"
                placeholder="Raison principale de la demande d'évaluation..."
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                Objet de l'évaluation
              </label>
              <textarea
                :value="formData.objet_evaluation"
                @input="updateField('objet_evaluation', $event.target.value)"
                rows="3"
                :class="getInputClass()"
                placeholder="Qu'est-ce qui doit être évalué spécifiquement..."
              ></textarea>
            </div>
          </div>
        </div>

        <!-- Section Cognitif -->
        <div v-show="activeSection === 'cognitif'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Brain class="text-slate-600" />
            Capacités Cognitives
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Observation générale des capacités cognitives
            </label>
            <textarea
              :value="formData.capacites_cognitives"
              @input="updateField('capacites_cognitives', $event.target.value)"
              rows="6"
              :class="getInputClass()"
              placeholder="Orientation, mémoire, attention, jugement, langage..."
            ></textarea>
          </div>
        </div>

        <!-- Section Santé -->
        <div v-show="activeSection === 'sante'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Heart class="text-slate-600" />
            État de Santé Physique
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Diagnostics médicaux, limitations fonctionnelles, autonomie
            </label>
            <textarea
              :value="formData.etat_sante_physique"
              @input="updateField('etat_sante_physique', $event.target.value)"
              rows="6"
              :class="getInputClass()"
              placeholder="État physique global, maladies chroniques, mobilité..."
            ></textarea>
          </div>
        </div>

        <!-- Section Psycho-social -->
        <div v-show="activeSection === 'psycho'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Activity class="text-slate-600" />
            Dimensions Psycho-sociales
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Humeur, émotions, comportements, adaptation
            </label>
            <textarea
              :value="formData.dimensions_psycho_sociales"
              @input="updateField('dimensions_psycho_sociales', $event.target.value)"
              rows="6"
              :class="getInputClass()"
              placeholder="État émotionnel, relations interpersonnelles..."
            ></textarea>
          </div>
        </div>

        <!-- Section Rôles -->
        <div v-show="activeSection === 'roles'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Briefcase class="text-slate-600" />
            Rôles Sociaux
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Activités quotidiennes, loisirs, implications sociales
            </label>
            <textarea
              :value="formData.roles_sociaux"
              @input="updateField('roles_sociaux', $event.target.value)"
              rows="6"
              :class="getInputClass()"
              placeholder="Implication communautaire, activités, rôles familiaux..."
            ></textarea>
          </div>
        </div>

        <!-- Section Réseau -->
        <div v-show="activeSection === 'reseau'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Users class="text-slate-600" />
            Réseau Social et Soutien
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Famille, proches, soutien disponible
            </label>
            <textarea
              :value="formData.reseau_social_soutien"
              @input="updateField('reseau_social_soutien', $event.target.value)"
              rows="6"
              :class="getInputClass()"
              placeholder="Personnes significatives, qualité des relations..."
            ></textarea>
          </div>
        </div>

        <!-- Section Analyse -->
        <div v-show="activeSection === 'analyse'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Layers class="text-slate-600" />
            Analyse Clinique
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Synthèse et analyse professionnelle
            </label>
            <textarea
              :value="formData.analyse_clinique"
              @input="updateField('analyse_clinique', $event.target.value)"
              rows="8"
              :class="getInputClass()"
              placeholder="Intégration des observations, facteurs de risque..."
            ></textarea>
          </div>
        </div>

        <!-- Section Opinion -->
        <div v-show="activeSection === 'opinion'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <CheckCircle class="text-slate-600" />
            Opinion Professionnelle
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Jugement professionnel et conclusions
            </label>
            <textarea
              :value="formData.opinion_professionnelle"
              @input="updateField('opinion_professionnelle', $event.target.value)"
              rows="8"
              :class="getInputClass()"
              placeholder="Opinion clinique basée sur l'évaluation..."
            ></textarea>
          </div>
        </div>

        <!-- Section Recommandations -->
        <div v-show="activeSection === 'recommandations'" class="space-y-6 animate-fadeIn">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white flex items-center gap-3 pb-4 border-b dark:border-gray-800">
            <Target class="text-slate-600" />
            Recommandations
          </h3>
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Interventions et suivis recommandés
            </label>
            <textarea
              :value="formData.recommandations"
              @input="updateField('recommandations', $event.target.value)"
              rows="8"
              :class="getInputClass()"
              placeholder="Plan d'intervention, services recommandés..."
            ></textarea>
          </div>
        </div>

      </div>
    </div>

    <!-- MODE VISUALISATION (évaluation finalisée) -->
    <div v-else-if="selectedEvaluation" class="flex-1 overflow-y-auto p-8">
      <div class="max-w-4xl mx-auto space-y-6">
        
        <!-- En-tête évaluation -->
        <div class="bg-gradient-to-r from-slate-50 to-slate-100 dark:from-gray-900 dark:to-gray-800 rounded-2xl p-6 border dark:border-gray-700">
          <h3 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">
            {{ selectedEvaluation.objet_evaluation || 'Évaluation du fonctionnement social' }}
          </h3>
          <div class="flex items-center gap-3 text-sm text-gray-600 dark:text-gray-400 mb-4">
            <span>Créée le {{ formatDate(selectedEvaluation.created_at) }}</span>
            <span>•</span>
            <span>Par {{ selectedEvaluation.auteur_nom }}</span>
          </div>

          <!-- Signature si verrouillée -->
          <div v-if="selectedEvaluation.verrouille && selectedEvaluation.signature_nom" class="pt-4 border-t dark:border-gray-700">
            <div class="flex items-center gap-2 text-emerald-600 dark:text-emerald-400">
              <ShieldCheck :size="18" />
              <span class="font-bold text-sm">{{ selectedEvaluation.signature_nom }}</span>
            </div>
            <div class="text-xs text-gray-500 mt-1">
              Signée le {{ formatDateTime(selectedEvaluation.date_signature) }}
            </div>
          </div>
        </div>

        <div v-for="section in displaySections" :key="section.key">
  <div 
    v-if="selectedEvaluation[section.key]?.trim()" 
    class="bg-slate-50 dark:bg-gray-900/30 rounded-xl p-6 border dark:border-gray-800 mb-6"
  >
    <h4 class="text-sm font-bold text-slate-600 dark:text-slate-400 uppercase mb-3 pb-2 border-b dark:border-gray-700">
      {{ section.label }}
    </h4>
    <p class="text-gray-800 dark:text-gray-200 whitespace-pre-wrap leading-relaxed">
      {{ selectedEvaluation[section.key] }}
    </p>
  </div>
</div>

      </div>
    </div>

    <!-- Message si rien sélectionné -->
    <div v-else class="flex-1 flex items-center justify-center text-gray-400">
      <div class="text-center">
        <FileText :size="64" class="mx-auto mb-4 opacity-20" />
        <p class="text-lg">Sélectionnez une évaluation ou créez-en une nouvelle</p>
      </div>
    </div>

    <!-- FOOTER AVEC BOUTONS (seulement en mode création) -->
    <div v-if="isCreating" class="px-8 py-4 bg-slate-50 dark:bg-gray-900 border-t dark:border-gray-800 flex items-center justify-between">
      <div class="flex gap-8 items-center text-[11px]">
        <div class="flex flex-col">
          <span class="font-bold text-gray-400 uppercase tracking-tighter leading-none mb-1">Progression</span>
          <div class="flex items-center gap-2">
            <div class="w-32 h-2 bg-gray-200 dark:bg-gray-700 rounded-full overflow-hidden">
              <div 
                class="h-full bg-gradient-to-r from-teal-500 to-teal-600 transition-all duration-500"
                :style="{ width: `${totalProgress}%` }"
              ></div>
            </div>
            <span class="font-bold text-teal-600">{{ totalProgress }}%</span>
          </div>
        </div>
        <div class="flex items-center gap-2 px-3 py-1 bg-orange-100 dark:bg-orange-900/30 text-orange-600 rounded-full">
          <span class="w-1.5 h-1.5 rounded-full bg-orange-500 animate-pulse"></span>
          <span class="font-bold uppercase tracking-tighter text-[9px]">Brouillon non scellé</span>
        </div>
      </div>

      <div class="flex gap-3">
        <button 
          @click="$emit('cancel')" 
          class="px-4 py-2 text-sm font-bold text-gray-500 hover:text-gray-700 dark:hover:text-gray-300 transition-colors"
        >
          Annuler
        </button>
        
        <!-- Bouton Brouillon -->
        <button 
          @click="$emit('save-draft')" 
          :disabled="isSaving || totalProgress < 20"
          class="bg-gray-600 hover:bg-gray-700 disabled:bg-gray-400 text-white px-6 py-2.5 rounded-xl font-bold shadow-lg transition-all active:scale-95 flex items-center gap-2"
        >
          <Save :size="18" /> 
          <span v-if="!isSaving">Sauvegarder brouillon</span>
          <span v-else>Sauvegarde...</span>
        </button>

        <!-- Bouton Définitif -->
        <button 
          @click="$emit('finalize')" 
          :disabled="isSaving || isFinalizing || totalProgress < 50"
          class="bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 disabled:from-gray-400 disabled:to-gray-500 text-white px-8 py-2.5 rounded-xl font-bold shadow-xl transition-all active:scale-95 flex items-center gap-2"
        >
          <ShieldCheck :size="18" /> 
          <span v-if="!isFinalizing">Sauvegarder définitivement</span>
          <span v-else>Finalisation...</span>
        </button>
      </div>
    </div>

  </div>
</template>

<script setup>
import { FileText, Brain, Heart, Activity, Briefcase, Users, Layers, CheckCircle, Target, ShieldCheck, Save } from 'lucide-vue-next'
import { computed } from 'vue'
const props = defineProps({
  isCreating: Boolean,
  selectedEvaluation: Object,
  formData: Object,
  errors: Object,
  activeSection: String,
  sections: Array,
  totalProgress: Number,
  isSaving: Boolean,
  isFinalizing: Boolean
})

const emit = defineEmits(['update:form-data', 'save-draft', 'finalize', 'cancel'])

const updateField = (field, value) => {
  emit('update:form-data', { ...props.formData, [field]: value })
}

const getInputClass = (error) => {
  const base = 'w-full px-4 py-3 rounded-xl border-2 dark:bg-gray-900 dark:text-white transition-all focus:outline-none focus:ring-2'
  return error 
    ? `${base} border-red-300 focus:border-red-500 focus:ring-red-200`
    : `${base} border-gray-200 dark:border-gray-700 focus:border-teal-500 focus:ring-teal-200`
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


const displaySections = [
  { label: "1. Contexte", key: "contexte_evaluation" },
  { label: "2. Motif de référence", key: "motif_reference" },
  { label: "3. Objet", key: "objet_evaluation" },
  { label: "4. Capacités cognitives", key: "capacites_cognitives" },
  { label: "5. Santé physique", key: "etat_sante_physique" },
  { label: "6. Dimensions psycho-sociales", key: "dimensions_psycho_sociales" },
  { label: "7. Rôles sociaux", key: "roles_sociaux" },
  { label: "8. Réseau social", key: "reseau_social_soutien" },
  { label: "9. Analyse clinique", key: "analyse_clinique" },
  { label: "10. Opinion professionnelle", key: "opinion_professionnelle" },
  { label: "11. Recommandations", key: "recommandations" }
]


</script>

<style scoped>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fadeIn {
  animation: fadeIn 0.3s ease-out;
}
</style>