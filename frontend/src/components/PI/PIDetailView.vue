<script setup>
import { ref } from 'vue'
import { 
  FileText, AlertCircle, Target, Calendar, 
  User, Lock, Edit3, Download, CheckCircle2,
  ShieldCheck, Clock
} from 'lucide-vue-next'

const props = defineProps({
  plan: {
    type: Object,
    required: true
  },
  client: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['edit', 'lock', 'export'])

const showLockDialog = ref(false)
const signatureName = ref('')

const formatDate = (dateStr) => {
  if (!dateStr) return 'Non défini'
  return new Date(dateStr).toLocaleDateString('fr-CA', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getStatusColor = (statut) => {
  const colors = {
    brouillon: 'amber',
    actif: 'teal',
    termine: 'slate',
    archive: 'gray'
  }
  return colors[statut] || 'slate'
}

const confirmLock = () => {
  if (signatureName.value.trim()) {
    emit('lock', props.plan.id, signatureName.value)
    showLockDialog.value = false
    signatureName.value = ''
  }
}
</script>

<template>
  <div class="flex-1 overflow-y-auto">
    <div class="max-w-5xl mx-auto p-8 space-y-8">
      
      <!-- En-tête du plan -->
      <div class="bg-gradient-to-br from-rose-50 to-pink-50 dark:from-rose-900/10 dark:to-pink-900/10 rounded-2xl p-6 border border-rose-200 dark:border-rose-800 shadow-sm">
        <div class="flex items-start justify-between mb-4">
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-2">
              <h1 class="text-2xl font-bold text-slate-900 dark:text-slate-100">
                {{ plan.titre }}
              </h1>
              <span 
                :class="[
                  'px-3 py-1 rounded-full text-xs font-bold uppercase',
                  `bg-${getStatusColor(plan.statut)}-100 dark:bg-${getStatusColor(plan.statut)}-900/30`,
                  `text-${getStatusColor(plan.statut)}-700 dark:text-${getStatusColor(plan.statut)}-300`
                ]"
              >
                {{ plan.statut }}
              </span>
            </div>
            
            <div class="flex items-center gap-4 text-sm text-slate-600 dark:text-slate-400">
              <div class="flex items-center gap-1.5">
                <FileText :size="16" />
                <span>{{ plan.type_plan }}</span>
              </div>
              <span class="text-slate-300 dark:text-slate-700">•</span>
              <div class="flex items-center gap-1.5">
                <Calendar :size="16" />
                <span>Créé le {{ formatDate(plan.created_at) }}</span>
              </div>
            </div>
          </div>

          <div v-if="plan.verrouille" class="flex flex-col items-end gap-2">
            <div class="flex items-center gap-2 bg-emerald-500 text-white px-4 py-2 rounded-lg shadow-md">
              <ShieldCheck :size="20" />
              <span class="font-bold text-sm">Plan Scellé</span>
            </div>
            <div class="text-xs text-slate-600 dark:text-slate-400">
              Signé par {{ plan.signature_nom }}
              <br>
              le {{ formatDate(plan.date_signature) }}
            </div>
          </div>
        </div>

        <!-- Dates importantes -->
        <div class="grid grid-cols-3 gap-4 mt-6 pt-4 border-t border-rose-200 dark:border-rose-800">
          <div>
            <p class="text-xs font-semibold text-slate-500 dark:text-slate-400 mb-1">Date de début</p>
            <p class="text-sm font-medium text-slate-900 dark:text-slate-100">
              {{ formatDate(plan.date_debut) }}
            </p>
          </div>
          <div>
            <p class="text-xs font-semibold text-slate-500 dark:text-slate-400 mb-1">Date de fin prévue</p>
            <p class="text-sm font-medium text-slate-900 dark:text-slate-100">
              {{ formatDate(plan.date_fin_prevue) }}
            </p>
          </div>
          <div>
            <p class="text-xs font-semibold text-slate-500 dark:text-slate-400 mb-1">Révision prévue</p>
            <p class="text-sm font-medium text-slate-900 dark:text-slate-100">
              {{ formatDate(plan.date_revision_prevue) }}
            </p>
          </div>
        </div>
      </div>

      <!-- Métadonnées -->
      <div class="bg-white dark:bg-slate-800 rounded-xl p-5 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-3">
          <User :size="18" class="text-slate-400" />
          <h3 class="font-semibold text-sm text-slate-900 dark:text-slate-100">Informations</h3>
        </div>
        <div class="grid grid-cols-2 gap-4 text-sm">
          <div>
            <span class="text-slate-500 dark:text-slate-400">Client:</span>
            <span class="ml-2 font-medium text-slate-900 dark:text-slate-100">
              {{ client.prenom }} {{ client.nom }}
            </span>
          </div>
          <div>
            <span class="text-slate-500 dark:text-slate-400">Dossier:</span>
            <span class="ml-2 font-mono font-medium text-slate-900 dark:text-slate-100">
              {{ client.no_dossier_leopard }}
            </span>
          </div>
          <div>
            <span class="text-slate-500 dark:text-slate-400">Auteur:</span>
            <span class="ml-2 font-medium text-slate-900 dark:text-slate-100">
              {{ plan.auteur_nom || 'N/A' }}
            </span>
          </div>
          <div>
            <span class="text-slate-500 dark:text-slate-400">Dernière modification:</span>
            <span class="ml-2 font-medium text-slate-900 dark:text-slate-100">
              {{ formatDate(plan.updated_at) }}
            </span>
          </div>
        </div>
      </div>

      <!-- Analyse: Contexte -->
      <section v-if="plan.contexte" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-amber-100 dark:bg-amber-900/30 rounded-lg">
            <AlertCircle :size="20" class="text-amber-600 dark:text-amber-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Contexte de vie</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.contexte }}</p>
        </div>
      </section>

      <!-- Analyse: Problématique -->
      <section v-if="plan.problematique" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-red-100 dark:bg-red-900/30 rounded-lg">
            <AlertCircle :size="20" class="text-red-600 dark:text-red-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Problématique identifiée</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.problematique }}</p>
        </div>
      </section>

      <!-- Analyse: Forces -->
      <section v-if="plan.forces" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-emerald-100 dark:bg-emerald-900/30 rounded-lg">
            <CheckCircle2 :size="20" class="text-emerald-600 dark:text-emerald-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Forces et ressources</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.forces }}</p>
        </div>
      </section>

      <!-- Objectifs -->
      <section v-if="plan.objectifs" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-teal-100 dark:bg-teal-900/30 rounded-lg">
            <Target :size="20" class="text-teal-600 dark:text-teal-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Objectifs SMART</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.objectifs }}</p>
        </div>
      </section>

      <!-- Interventions -->
      <section v-if="plan.interventions" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-purple-100 dark:bg-purple-900/30 rounded-lg">
            <Calendar :size="20" class="text-purple-600 dark:text-purple-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Interventions planifiées</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.interventions }}</p>
        </div>
      </section>

      <!-- Résultats attendus -->
      <section v-if="plan.resultats" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg">
            <CheckCircle2 :size="20" class="text-blue-600 dark:text-blue-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Résultats attendus</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.resultats }}</p>
        </div>
      </section>

      <!-- Ententes -->
      <section v-if="plan.ententes" class="bg-white dark:bg-slate-800 rounded-xl p-6 border border-slate-200 dark:border-slate-700 shadow-sm">
        <div class="flex items-center gap-2 mb-4">
          <div class="p-2 bg-indigo-100 dark:bg-indigo-900/30 rounded-lg">
            <User :size="20" class="text-indigo-600 dark:text-indigo-400" />
          </div>
          <h2 class="text-lg font-bold text-slate-900 dark:text-slate-100">Ententes et engagements</h2>
        </div>
        <div class="prose prose-sm dark:prose-invert max-w-none">
          <p class="text-slate-700 dark:text-slate-300 whitespace-pre-wrap">{{ plan.ententes }}</p>
        </div>
      </section>

      <!-- Actions -->
      <div v-if="!plan.verrouille" class="flex items-center justify-between gap-4 pt-6 border-t border-slate-200 dark:border-slate-700">
        <button
          @click="$emit('edit')"
          class="flex items-center gap-2 px-6 py-3 bg-slate-100 dark:bg-slate-800 hover:bg-slate-200 dark:hover:bg-slate-700 text-slate-900 dark:text-slate-100 rounded-xl font-semibold transition-colors"
        >
          <Edit3 :size="18" />
          <span>Modifier le plan</span>
        </button>

        <button
          @click="showLockDialog = true"
          class="flex items-center gap-2 px-6 py-3 bg-emerald-500 hover:bg-emerald-600 text-white rounded-xl font-bold transition-colors shadow-lg"
        >
          <Lock :size="18" />
          <span>Sceller et activer</span>
        </button>
      </div>

      <!-- Dialog de verrouillage -->
      <div 
        v-if="showLockDialog" 
        class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4"
        @click.self="showLockDialog = false"
      >
        <div class="bg-white dark:bg-slate-800 rounded-2xl p-6 max-w-md w-full shadow-2xl">
          <div class="flex items-center gap-3 mb-4">
            <div class="p-3 bg-emerald-100 dark:bg-emerald-900/30 rounded-xl">
              <ShieldCheck :size="24" class="text-emerald-600 dark:text-emerald-400" />
            </div>
            <div>
              <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100">Sceller le plan</h3>
              <p class="text-sm text-slate-600 dark:text-slate-400">Cette action est irréversible</p>
            </div>
          </div>

          <div class="bg-amber-50 dark:bg-amber-900/20 border border-amber-200 dark:border-amber-800 rounded-lg p-3 mb-4">
            <p class="text-xs text-amber-800 dark:text-amber-200">
              ⚠️ Une fois scellé, le plan ne pourra plus être modifié. 
              Il passera automatiquement au statut "Actif".
            </p>
          </div>

          <FormKit
            type="text"
            label="Nom du signataire *"
            v-model="signatureName"
            placeholder="Votre nom complet"
            validation="required"
            help="Cette signature officialise le plan"
          />

          <div class="flex gap-3 mt-6">
            <button
              @click="showLockDialog = false"
              class="flex-1 px-4 py-2 bg-slate-100 dark:bg-slate-700 hover:bg-slate-200 dark:hover:bg-slate-600 text-slate-900 dark:text-slate-100 rounded-lg font-medium transition-colors"
            >
              Annuler
            </button>
            <button
              @click="confirmLock"
              :disabled="!signatureName.trim()"
              class="flex-1 px-4 py-2 bg-emerald-500 hover:bg-emerald-600 text-white rounded-lg font-bold transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Confirmer
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>