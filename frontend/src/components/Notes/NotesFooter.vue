<template>
  <div v-if="brouillons.length > 0" class="border-t dark:border-gray-800 bg-gradient-to-r from-slate-50/90 via-amber-50/50 to-slate-50/90 dark:from-slate-900/60 dark:via-amber-900/20 dark:to-slate-900/60 backdrop-blur-md">
    
    <!-- En-tête du footer -->
    <div class="px-4 py-2 border-b dark:border-gray-800 flex items-center justify-between">
      <div class="flex items-center gap-3">
        <div class="p-1.5 bg-amber-100 dark:bg-amber-900/40 rounded-lg">
          <FileEdit :size="16" class="text-amber-600 dark:text-amber-400" />
        </div>
        <div>
          <h3 class="text-xs font-bold text-slate-700 dark:text-slate-200">
            Brouillons en attente
          </h3>
          <p class="text-[10px] text-slate-500 dark:text-slate-400">
            {{ brouillons.length }} note{{ brouillons.length > 1 ? 's' : '' }} non signée{{ brouillons.length > 1 ? 's' : '' }}
          </p>
        </div>
      </div>

      <!-- Actions globales -->
      <div class="flex items-center gap-2">
        <button
          v-if="brouillons.length > 1"
          @click="$emit('delete-all-drafts')"
          class="px-3 py-1.5 text-xs font-medium text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors flex items-center gap-1.5"
        >
          <Trash2 :size="12" />
          Tout supprimer
        </button>
      </div>
    </div>

    <!-- Liste des brouillons -->
    <div class="p-3 flex gap-3 overflow-x-auto">
      <div 
        v-for="draft in brouillons" 
        :key="draft.id"
        class="flex-shrink-0 group"
      >
        <!-- Carte brouillon -->
        <div class="bg-white dark:bg-gray-800 border-2 border-amber-200 dark:border-amber-800 rounded-xl shadow-sm hover:shadow-md transition-all min-w-[280px] max-w-[320px]">
          
          <!-- Header de la carte -->
          <div class="px-4 py-2.5 border-b dark:border-gray-700 bg-gradient-to-r from-amber-50 to-white dark:from-amber-900/20 dark:to-gray-800">
            <div class="flex items-start justify-between gap-2">
              <div class="flex-1 min-w-0">
                <h4 
                  @click="$emit('view-draft', draft)"
                  class="text-sm font-bold text-slate-800 dark:text-slate-100 truncate cursor-pointer hover:text-teal-600 dark:hover:text-teal-400 transition-colors"
                  :title="draft.titre || 'Sans titre'"
                >
                  {{ draft.titre || 'Sans titre' }}
                </h4>
                <div class="flex items-center gap-2 mt-1">
                  <span 
                    :class="getTypeColor(draft.type_note)"
                    class="text-[9px] font-black uppercase px-2 py-0.5 rounded"
                  >
                    {{ draft.type_note }}
                  </span>
                </div>
              </div>

              <!-- Icône brouillon -->
              <div class="flex-shrink-0">
                <div class="p-1 bg-amber-100 dark:bg-amber-900/40 rounded">
                  <AlertTriangle :size="14" class="text-amber-600 dark:text-amber-400" />
                </div>
              </div>
            </div>
          </div>

          <!-- Corps de la carte -->
          <div class="px-4 py-3 space-y-2">
            <!-- Métadonnées -->
            <div class="grid grid-cols-2 gap-2 text-[10px]">
              <div class="flex items-center gap-1.5 text-slate-600 dark:text-slate-400">
                <Calendar :size="11" />
                <span>{{ formatDate(draft.date_note) }}</span>
              </div>
              <div class="flex items-center gap-1.5 text-slate-600 dark:text-slate-400">
                <Type :size="11" />
                <span>{{ getWordCount(draft.contenu) }} mots</span>
              </div>
            </div>

            <!-- Aperçu du contenu -->
            <div class="text-[11px] text-slate-500 dark:text-slate-400 line-clamp-2 italic">
              {{ getPreview(draft.contenu) }}
            </div>
          </div>

          <!-- Footer de la carte avec actions -->
          <div class="px-4 py-2.5 border-t dark:border-gray-700 bg-slate-50 dark:bg-gray-900/30 flex items-center justify-between">
            <button
              @click="$emit('view-draft', draft)"
              class="text-xs font-semibold text-teal-600 dark:text-teal-400 hover:text-teal-700 dark:hover:text-teal-300 flex items-center gap-1.5 transition-colors"
            >
              <Edit3 :size="12" />
              Éditer
            </button>

            <div class="flex items-center gap-1">
              <button 
                @click="$emit('finalize-draft', draft.id)"
                class="p-1.5 text-teal-600 dark:text-teal-400 hover:bg-teal-50 dark:hover:bg-teal-900/30 rounded transition-colors"
                title="Finaliser et signer"
              >
                <ShieldCheck :size="16" />
              </button>
              <button 
                @click="$emit('delete-draft', draft.id)"
                class="p-1.5 text-slate-400 dark:text-slate-500 hover:text-red-500 dark:hover:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded transition-colors"
                title="Supprimer"
              >
                <Trash2 :size="14" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Aide contextuelle -->
    <div class="px-4 py-2 border-t dark:border-gray-800 bg-amber-50/30 dark:bg-amber-900/10">
      <p class="text-[10px] text-slate-500 dark:text-slate-400 text-center">
        💡 <strong>Astuce:</strong> Les brouillons sont automatiquement sauvegardés. Finalisez-les pour les verrouiller définitivement.
      </p>
    </div>
  </div>
</template>

<script setup>
import { FileEdit, Trash2, Edit3, ShieldCheck, AlertTriangle, Type, Calendar } from 'lucide-vue-next'

defineProps({
  brouillons: { type: Array, default: () => [] }
})

defineEmits(['view-draft', 'delete-draft', 'delete-all-drafts', 'finalize-draft'])

const getTypeColor = (type) => {
  const colors = {
    'Suivi': 'bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300',
    'Ouverture': 'bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300',
    'Plan': 'bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300',
    'Fermeture': 'bg-gray-100 dark:bg-gray-900/30 text-gray-700 dark:text-gray-300',
    'Incident': 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300',
    'Correction': 'bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300',
    'Addendum': 'bg-sky-100 dark:bg-sky-900/30 text-sky-700 dark:text-sky-300'
  }
  return colors[type] || 'bg-teal-100 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300'
}

const getWordCount = (content) => {
  if (!content) return 0
  return content.trim().split(/\s+/).filter(w => w.length > 0).length
}

const getPreview = (content) => {
  if (!content) return 'Aucun contenu'
  const cleanContent = content.replace(/\n/g, ' ').trim()
  return cleanContent.length > 80 ? cleanContent.substring(0, 80) + '...' : cleanContent
}

const formatDate = (d) => {
  if (!d) return '-'
  const date = new Date(d)
  const now = new Date()
  const diffMs = now - date
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMs / 3600000)
  const diffDays = Math.floor(diffMs / 86400000)

  if (diffMins < 1) return 'À l\'instant'
  if (diffMins < 60) return `Il y a ${diffMins} min`
  if (diffHours < 24) return `Il y a ${diffHours}h`
  if (diffDays === 0) return 'Aujourd\'hui'
  if (diffDays === 1) return 'Hier'
  if (diffDays < 7) return `Il y a ${diffDays}j`
  
  return date.toLocaleDateString('fr-CA', { 
    year: 'numeric', 
    month: 'short', 
    day: 'numeric'
  })
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>