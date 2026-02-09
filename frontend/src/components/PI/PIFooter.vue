<script setup>
import { computed } from 'vue'
import { 
  Save, Calendar, FileText, AlertCircle, 
  Target, Users, Clock 
} from 'lucide-vue-next'

const props = defineProps({
  currentPi: {
    type: Object,
    required: true
  },
  isSaving: {
    type: Boolean,
    default: false
  },
  currentView: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['save', 'cancel'])

const completionPercentage = computed(() => {
  const fields = [
    props.currentPi.titre,
    props.currentPi.date_debut,
    props.currentPi.contexte,
    props.currentPi.problematique,
    props.currentPi.forces,
    props.currentPi.objectifs,
    props.currentPi.interventions,
    props.currentPi.ententes
  ]
  
  const filled = fields.filter(f => f && f.toString().trim()).length
  return Math.round((filled / fields.length) * 100)
})

const getCompletionColor = (percentage) => {
  if (percentage >= 80) return 'emerald'
  if (percentage >= 50) return 'amber'
  return 'rose'
}

const completionColor = computed(() => getCompletionColor(completionPercentage.value))

const formatDate = (dateStr) => {
  if (!dateStr) return 'Non définie'
  return new Date(dateStr).toLocaleDateString('fr-CA', {
    month: 'short',
    day: 'numeric'
  })
}

const getSectionStatus = (field) => {
  return field && field.trim() ? 'complete' : 'incomplete'
}
</script>

<template>
  <footer class="border-t border-slate-200 dark:border-slate-800 
                 bg-gradient-to-r from-rose-500/95 via-pink-500/95 to-purple-500/95
                 backdrop-blur-sm shadow-lg">
    <div class="px-6 py-4">
      
      <!-- Résumé compact -->
      <div class="flex items-center justify-between mb-3">
        <div class="flex items-center gap-4">
          <!-- Titre et type -->
          <div class="flex items-center gap-3">
            <div class="p-2 bg-white/20 rounded-lg backdrop-blur-sm">
              <FileText :size="18" class="text-white" />
            </div>
            <div>
              <p class="text-sm font-bold text-white">
                {{ currentPi.titre || 'Plan sans titre' }}
              </p>
              <p class="text-xs text-white/75">
                {{ currentPi.type_plan || 'psychosocial' }}
              </p>
            </div>
          </div>

          <!-- Séparateur -->
          <div class="h-8 w-px bg-white/20"></div>

          <!-- Dates -->
          <div class="flex items-center gap-2 text-xs text-white/90">
            <Calendar :size="14" />
            <span>Début: {{ formatDate(currentPi.date_debut) }}</span>
            <span v-if="currentPi.date_fin_prevue" class="text-white/60">→</span>
            <span v-if="currentPi.date_fin_prevue">Fin: {{ formatDate(currentPi.date_fin_prevue) }}</span>
          </div>

          <!-- Séparateur -->
          <div class="h-8 w-px bg-white/20"></div>

          <!-- Progression -->
          <div class="flex items-center gap-3">
            <div class="flex items-center gap-2">
              <div class="w-32 h-2 bg-white/20 rounded-full overflow-hidden">
                <div 
                  :class="`h-full bg-${completionColor}-400 transition-all duration-500`"
                  :style="{ width: `${completionPercentage}%` }"
                ></div>
              </div>
              <span class="text-xs font-bold text-white">
                {{ completionPercentage }}%
              </span>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center gap-3">
          <button
            @click="$emit('cancel')"
            class="px-4 py-2 text-sm font-medium text-white/90 hover:text-white hover:bg-white/10 rounded-lg transition-colors backdrop-blur-sm"
          >
            Annuler
          </button>
          
          <button
            @click="$emit('save')"
            :disabled="isSaving || !currentPi.titre || !currentPi.date_debut"
            class="px-6 py-2 bg-white hover:bg-white/95 text-rose-600 rounded-lg font-bold text-sm flex items-center gap-2 transition-all shadow-lg hover:shadow-xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <Save :size="18" />
            <span>{{ isSaving ? 'Enregistrement...' : 'Enregistrer' }}</span>
          </button>
        </div>
      </div>

      <!-- Indicateurs de sections -->
      <div class="flex items-center gap-2 overflow-x-auto pb-1">
        <div 
          :class="[
            'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium backdrop-blur-sm',
            getSectionStatus(currentPi.contexte) === 'complete' 
              ? 'bg-emerald-400/20 text-white border border-emerald-400/30' 
              : 'bg-white/10 text-white/60 border border-white/20'
          ]"
        >
          <AlertCircle :size="12" />
          <span>Contexte</span>
        </div>

        <div 
          :class="[
            'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium backdrop-blur-sm',
            getSectionStatus(currentPi.problematique) === 'complete' 
              ? 'bg-emerald-400/20 text-white border border-emerald-400/30' 
              : 'bg-white/10 text-white/60 border border-white/20'
          ]"
        >
          <AlertCircle :size="12" />
          <span>Problématique</span>
        </div>

        <div 
          :class="[
            'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium backdrop-blur-sm',
            getSectionStatus(currentPi.forces) === 'complete' 
              ? 'bg-emerald-400/20 text-white border border-emerald-400/30' 
              : 'bg-white/10 text-white/60 border border-white/20'
          ]"
        >
          <Target :size="12" />
          <span>Forces</span>
        </div>

        <div 
          :class="[
            'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium backdrop-blur-sm',
            getSectionStatus(currentPi.objectifs) === 'complete' 
              ? 'bg-emerald-400/20 text-white border border-emerald-400/30' 
              : 'bg-white/10 text-white/60 border border-white/20'
          ]"
        >
          <Target :size="12" />
          <span>Objectifs</span>
        </div>

        <div 
          :class="[
            'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium backdrop-blur-sm',
            getSectionStatus(currentPi.interventions) === 'complete' 
              ? 'bg-emerald-400/20 text-white border border-emerald-400/30' 
              : 'bg-white/10 text-white/60 border border-white/20'
          ]"
        >
          <Calendar :size="12" />
          <span>Interventions</span>
        </div>

        <div 
          :class="[
            'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium backdrop-blur-sm',
            getSectionStatus(currentPi.ententes) === 'complete' 
              ? 'bg-emerald-400/20 text-white border border-emerald-400/30' 
              : 'bg-white/10 text-white/60 border border-white/20'
          ]"
        >
          <Users :size="12" />
          <span>Ententes</span>
        </div>
      </div>
    </div>
  </footer>
</template>