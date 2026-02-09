<script setup>
import { computed } from 'vue'
import { 
  FileText, AlertCircle, CheckCircle2, 
  Calendar, User, ChevronRight, Save 
} from 'lucide-vue-next'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  },
  activeSection: {
    type: String,
    required: true
  },
  isSaving: {
    type: Boolean,
    default: false
  },
  isEditing: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'update:activeSection', 'save', 'cancel'])

const formData = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const sections = [
  { id: 'identification', label: 'Identification', icon: FileText },
  { id: 'analyse', label: 'Analyse', icon: AlertCircle },
  { id: 'objectifs', label: 'Objectifs', icon: CheckCircle2 },
  { id: 'interventions', label: 'Interventions', icon: Calendar },
  { id: 'ententes', label: 'Ententes', icon: User }
]

const currentSectionIndex = computed(() => 
  sections.findIndex(s => s.id === props.activeSection)
)

const canGoNext = computed(() => currentSectionIndex.value < sections.length - 1)
const canGoPrev = computed(() => currentSectionIndex.value > 0)

const goNext = () => {
  if (canGoNext.value) {
    emit('update:activeSection', sections[currentSectionIndex.value + 1].id)
  }
}

const goPrev = () => {
  if (canGoPrev.value) {
    emit('update:activeSection', sections[currentSectionIndex.value - 1].id)
  }
}

// Validation basique
const canSubmit = computed(() => {
  return formData.value.titre && 
         formData.value.date_debut &&
         formData.value.contexte
})
</script>

<template>
  <div class="flex-1 flex flex-col overflow-hidden">
    
    <!-- Progress Bar -->
    <div class="px-6 py-3 bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800">
      <div class="flex items-center gap-2">
        <div 
          v-for="(section, idx) in sections" 
          :key="section.id"
          class="flex items-center flex-1"
        >
          <div class="flex-1">
            <div 
              :class="[
                'h-1.5 rounded-full transition-all',
                idx < currentSectionIndex ? 'bg-rose-500' :
                idx === currentSectionIndex ? 'bg-rose-500' :
                'bg-slate-200 dark:bg-slate-700'
              ]"
            ></div>
          </div>
          <ChevronRight 
            v-if="idx < sections.length - 1" 
            :size="16" 
            class="mx-1 text-slate-300 dark:text-slate-600 flex-shrink-0" 
          />
        </div>
      </div>
      
      <div class="flex items-center justify-between mt-2">
        <div class="flex items-center gap-2">
          <component 
            :is="sections[currentSectionIndex].icon" 
            :size="18" 
            class="text-rose-500"
          />
          <span class="text-sm font-semibold text-slate-900 dark:text-slate-100">
            {{ sections[currentSectionIndex].label }}
          </span>
        </div>
        <span class="text-xs text-slate-500 dark:text-slate-400">
          Étape {{ currentSectionIndex + 1 }} / {{ sections.length }}
        </span>
      </div>
    </div>

    <!-- Form Content -->
    <div class="flex-1 overflow-y-auto p-6">
      <div class="max-w-3xl mx-auto space-y-6">
        
        <!-- Section: Identification -->
        <div v-if="activeSection === 'identification'" class="space-y-6 animate-fade-in">
          <div class="bg-rose-50 dark:bg-rose-900/10 border border-rose-200 dark:border-rose-800 rounded-xl p-4 mb-6">
            <h3 class="font-semibold text-sm text-rose-900 dark:text-rose-100 mb-1">
              Informations générales
            </h3>
            <p class="text-xs text-rose-700 dark:text-rose-300">
              Définissez les informations de base du plan d'intervention
            </p>
          </div>

          <FormKit
            type="text"
            label="Titre du plan *"
            v-model="formData.titre"
            validation="required"
            placeholder="Ex: Plan d'intervention psychosocial - Automne 2024"
            help="Donnez un titre descriptif à votre plan"
          />

          <FormKit
            type="select"
            label="Type de plan"
            v-model="formData.type_plan"
            :options="[
              { label: 'Psychosocial', value: 'psychosocial' },
              { label: 'Médical', value: 'medical' },
              { label: 'Éducatif', value: 'educatif' },
              { label: 'Comportemental', value: 'comportemental' }
            ]"
          />

          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <FormKit
              type="date"
              label="Date de début *"
              v-model="formData.date_debut"
              validation="required"
            />

            <FormKit
              type="date"
              label="Date de fin prévue"
              v-model="formData.date_fin_prevue"
              help="Échéance du plan"
            />

            <FormKit
              type="date"
              label="Date de révision prévue"
              v-model="formData.date_revision_prevue"
              help="Quand réviser ce plan?"
            />
          </div>
        </div>

        <!-- Section: Analyse -->
        <div v-else-if="activeSection === 'analyse'" class="space-y-6 animate-fade-in">
          <div class="bg-amber-50 dark:bg-amber-900/10 border border-amber-200 dark:border-amber-800 rounded-xl p-4 mb-6">
            <h3 class="font-semibold text-sm text-amber-900 dark:text-amber-100 mb-1">
              Analyse de la situation
            </h3>
            <p class="text-xs text-amber-700 dark:text-amber-300">
              Analysez le contexte, la problématique et les forces du client
            </p>
          </div>

          <FormKit
            type="textarea"
            label="Contexte de vie *"
            v-model="formData.contexte"
            validation="required"
            :rows="6"
            placeholder="Décrivez la situation actuelle du client, son environnement familial, social, professionnel..."
            help="Situation familiale, sociale, professionnelle, etc."
          />

          <FormKit
            type="textarea"
            label="Problématique identifiée"
            v-model="formData.problematique"
            :rows="6"
            placeholder="Quels sont les défis, difficultés ou besoins identifiés?"
            help="Difficultés, besoins, défis identifiés"
          />

          <FormKit
            type="textarea"
            label="Forces et ressources"
            v-model="formData.forces"
            :rows="6"
            placeholder="Quelles sont les forces, capacités et ressources du client?"
            help="Forces personnelles, soutien familial, ressources disponibles"
          />
        </div>

        <!-- Section: Objectifs -->
        <div v-else-if="activeSection === 'objectifs'" class="space-y-6 animate-fade-in">
          <div class="bg-teal-50 dark:bg-teal-900/10 border border-teal-200 dark:border-teal-800 rounded-xl p-4 mb-6">
            <h3 class="font-semibold text-sm text-teal-900 dark:text-teal-100 mb-1">
              Objectifs SMART
            </h3>
            <p class="text-xs text-teal-700 dark:text-teal-300">
              Définissez des objectifs Spécifiques, Mesurables, Atteignables, Réalistes et Temporels
            </p>
          </div>

          <div class="bg-teal-50/50 dark:bg-teal-900/5 rounded-lg p-4 space-y-2 mb-4">
            <p class="text-xs font-semibold text-teal-900 dark:text-teal-100">Guide SMART:</p>
            <ul class="text-xs text-teal-800 dark:text-teal-200 space-y-1 ml-4">
              <li><strong>S</strong>pécifique - Clair et précis</li>
              <li><strong>M</strong>esurable - Quantifiable</li>
              <li><strong>A</strong>tteignable - Réalisable</li>
              <li><strong>R</strong>éaliste - Pertinent au contexte</li>
              <li><strong>T</strong>emporel - Délai défini</li>
            </ul>
          </div>

          <FormKit
            type="textarea"
            label="Objectifs du plan"
            v-model="formData.objectifs"
            :rows="12"
            placeholder="Objectif 1: [Description]&#10;- Indicateur de mesure: ...&#10;- Échéance: ...&#10;&#10;Objectif 2: ..."
            help="Listez les objectifs en utilisant la méthode SMART"
          />
        </div>

        <!-- Section: Interventions -->
        <div v-else-if="activeSection === 'interventions'" class="space-y-6 animate-fade-in">
          <div class="bg-purple-50 dark:bg-purple-900/10 border border-purple-200 dark:border-purple-800 rounded-xl p-4 mb-6">
            <h3 class="font-semibold text-sm text-purple-900 dark:text-purple-100 mb-1">
              Plan d'action
            </h3>
            <p class="text-xs text-purple-700 dark:text-purple-300">
              Détaillez les interventions prévues pour atteindre les objectifs
            </p>
          </div>

          <FormKit
            type="textarea"
            label="Interventions planifiées"
            v-model="formData.interventions"
            :rows="12"
            placeholder="Intervention 1:&#10;- Actions concrètes: ...&#10;- Fréquence: ...&#10;- Responsable: ...&#10;&#10;Intervention 2: ..."
            help="Actions, moyens, stratégies à mettre en place"
          />

          <FormKit
            type="textarea"
            label="Résultats attendus"
            v-model="formData.resultats"
            :rows="6"
            placeholder="Quels résultats espérez-vous observer?"
            help="Impacts attendus des interventions"
          />
        </div>

        <!-- Section: Ententes -->
        <div v-else-if="activeSection === 'ententes'" class="space-y-6 animate-fade-in">
          <div class="bg-indigo-50 dark:bg-indigo-900/10 border border-indigo-200 dark:border-indigo-800 rounded-xl p-4 mb-6">
            <h3 class="font-semibold text-sm text-indigo-900 dark:text-indigo-100 mb-1">
              Engagements et modalités
            </h3>
            <p class="text-xs text-indigo-700 dark:text-indigo-300">
              Définissez les engagements du client et les modalités de suivi
            </p>
          </div>

          <FormKit
            type="textarea"
            label="Ententes avec le client"
            v-model="formData.ententes"
            :rows="8"
            placeholder="- Engagement du client: ...&#10;- Fréquence des rencontres: ...&#10;- Modalités de suivi: ...&#10;- Révision du plan: ..."
            help="Engagements mutuels, fréquence des suivis, modalités de révision"
          />

          <div class="bg-indigo-50 dark:bg-indigo-900/10 rounded-lg p-4">
            <p class="text-xs font-semibold text-indigo-900 dark:text-indigo-100 mb-2">
              Avant de finaliser:
            </p>
            <ul class="text-xs text-indigo-800 dark:text-indigo-200 space-y-1 ml-4">
              <li>✓ Vérifiez que toutes les sections sont complétées</li>
              <li>✓ Assurez-vous que les objectifs sont SMART</li>
              <li>✓ Confirmez les dates et échéances</li>
              <li>✓ Validez les ententes avec le client</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Navigation Footer -->
    <div class="px-6 py-4 bg-white dark:bg-slate-900 border-t border-slate-200 dark:border-slate-800">
      <div class="max-w-3xl mx-auto flex items-center justify-between">
        <button
          v-if="canGoPrev"
          @click="goPrev"
          class="px-4 py-2 text-sm font-medium text-slate-600 dark:text-slate-400 hover:text-slate-900 dark:hover:text-slate-100 transition-colors"
        >
          ← Précédent
        </button>
        <div v-else></div>

        <div class="flex items-center gap-3">
          <button
            @click="$emit('cancel')"
            class="px-5 py-2 text-sm font-medium text-slate-600 dark:text-slate-400 hover:text-slate-900 dark:hover:text-slate-100 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
          >
            Annuler
          </button>

          <button
            v-if="canGoNext"
            @click="goNext"
            class="px-5 py-2 text-sm font-semibold bg-rose-500 hover:bg-rose-600 text-white rounded-lg transition-colors flex items-center gap-2"
          >
            <span>Suivant</span>
            <ChevronRight :size="16" />
          </button>

          <button
            v-else
            @click="$emit('save')"
            :disabled="!canSubmit || isSaving"
            class="px-6 py-2 text-sm font-bold bg-emerald-500 hover:bg-emerald-600 text-white rounded-lg transition-colors flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg hover:shadow-xl"
          >
            <Save :size="18" />
            <span>{{ isSaving ? 'Enregistrement...' : (isEditing ? 'Mettre à jour' : 'Enregistrer le plan') }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.3s ease-out;
}
</style>