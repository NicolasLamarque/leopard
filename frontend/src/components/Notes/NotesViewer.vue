<template>
  <div class="flex-1 flex flex-col overflow-hidden bg-white dark:bg-gray-950">
    
    <!-- Header de la note -->
    <div class="p-6 border-b dark:border-gray-800 bg-gradient-to-r from-slate-50 to-teal-50 dark:from-slate-900/50 dark:to-teal-900/20">
      <div class="flex items-start justify-between mb-4">
        <div class="flex-1">
          <div class="flex items-center gap-3 mb-2">
            <span :class="[
              'text-xs font-bold uppercase px-3 py-1 rounded-lg',
              getTypeColor(note.type_note)
            ]">
              {{ note.type_note }}
            </span>
            <span class="text-xs text-gray-500 dark:text-gray-400">
              {{ formatDate(note.date_intervention) }}
            </span>
            <span class="text-xs text-gray-500 dark:text-gray-400">
              {{ note.mode_intervention }}
            </span>
          </div>
          <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-1">
            {{ note.titre || 'Sans titre' }}
          </h2>
          
          <!-- Note liée -->
          <div v-if="note.note_liee_id" class="flex items-center gap-2 mt-2">
            <Link2 :size="14" class="text-sky-500" />
            <span class="text-sm text-sky-600 dark:text-sky-400">
              {{ note.type_lien || 'Lié' }} à la note: 
              <button 
                @click="$emit('view-linked-note', note.note_liee_id)"
                class="font-semibold hover:underline"
              >
                {{ note.note_liee_titre || '#' + note.note_liee_id }}
              </button>
            </span>
          </div>
        </div>

        <div class="flex items-center gap-2">
          <button
            @click="$emit('create-correction')"
            class="px-4 py-2 bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300 rounded-lg text-sm font-semibold hover:bg-orange-200 dark:hover:bg-orange-900/50 transition-all flex items-center gap-2"
          >
            <Edit3 :size="16" />
            Créer correction
          </button>
          
          <button
            @click="$emit('create-addendum')"
            class="px-4 py-2 bg-sky-100 dark:bg-sky-900/30 text-sky-700 dark:text-sky-300 rounded-lg text-sm font-semibold hover:bg-sky-200 dark:hover:bg-sky-900/50 transition-all flex items-center gap-2"
          >
            <Plus :size="16" />
            Créer addendum
          </button>

          <button
            @click="$emit('export-single')"
            class="px-4 py-2 bg-teal-100 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300 rounded-lg text-sm font-semibold hover:bg-teal-200 dark:hover:bg-teal-900/50 transition-all flex items-center gap-2"
          >
            <FileDown :size="16" />
            Export PDF
          </button>
        </div>
      </div>

      <!-- Stats de la note -->
      <div class="grid grid-cols-4 gap-3">
        <div class="p-3 bg-white dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-gray-800">
          <div class="text-xs text-gray-500 dark:text-gray-400 mb-1">Créée le</div>
          <div class="text-sm font-semibold text-gray-900 dark:text-white">
            {{ formatDateTime(note.date_note) }}
          </div>
        </div>
        
        <div class="p-3 bg-white dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-gray-800">
          <div class="text-xs text-gray-500 dark:text-gray-400 mb-1">Mots</div>
          <div class="text-sm font-semibold text-gray-900 dark:text-white">
            {{ wordCount }}
          </div>
        </div>

        <div class="p-3 bg-white dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-gray-800">
          <div class="text-xs text-gray-500 dark:text-gray-400 mb-1">Caractères</div>
          <div class="text-sm font-semibold text-gray-900 dark:text-white">
            {{ charCount }}
          </div>
        </div>

        <div class="p-3 bg-emerald-50 dark:bg-emerald-900/20 rounded-lg border border-emerald-200 dark:border-emerald-800">
          <div class="text-xs text-emerald-600 dark:text-emerald-400 mb-1 flex items-center gap-1">
            <ShieldCheck :size="12" />
            Statut
          </div>
          <div class="text-sm font-semibold text-emerald-700 dark:text-emerald-300">
            Verrouillée
          </div>
        </div>
      </div>
    </div>

    <!-- Contenu de la note -->
    <div class="flex-1 overflow-y-auto p-8">
      <div class="max-w-4xl mx-auto">
        <div class="prose prose-slate dark:prose-invert max-w-none">
          <div 
            class="whitespace-pre-wrap text-gray-800 dark:text-gray-200 text-base leading-relaxed font-mono bg-white dark:bg-gray-900 p-6 rounded-xl border border-gray-200 dark:border-gray-800 shadow-sm"
            v-html="formattedContent"
          ></div>
        </div>
      </div>
    </div>

    <!-- Footer avec signature -->
    <div class="p-6 border-t dark:border-gray-800 bg-gradient-to-r from-emerald-50 to-teal-50 dark:from-emerald-900/10 dark:to-teal-900/10">
      <div class="max-w-4xl mx-auto">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <div class="p-3 bg-emerald-100 dark:bg-emerald-900/30 rounded-lg">
              <ShieldCheck :size="24" class="text-emerald-600 dark:text-emerald-400" />
            </div>
            <div>
              <div class="text-sm font-semibold text-gray-700 dark:text-gray-300">Note signée par</div>
              <div class="text-lg font-bold text-gray-900 dark:text-white">{{ note.signature_nom }}</div>
              <div class="text-xs text-gray-500 dark:text-gray-400 mt-1">
                Le {{ formatDateTime(note.signature_date) }}
              </div>
            </div>
          </div>

          <div class="text-right">
            <div class="text-xs text-gray-500 dark:text-gray-400 mb-1">ID de la note</div>
            <div class="text-sm font-mono font-semibold text-gray-700 dark:text-gray-300">
              #{{ note.id }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ShieldCheck, FileDown, Edit3, Plus, Link2 } from 'lucide-vue-next'

const props = defineProps({
  note: { type: Object, required: true }
})

defineEmits(['create-correction', 'create-addendum', 'export-single', 'view-linked-note'])

const wordCount = computed(() => {
  if (!props.note.contenu) return 0
  return props.note.contenu.trim().split(/\s+/).filter(w => w.length > 0).length
})

const charCount = computed(() => {
  return props.note.contenu?.length || 0
})

const formattedContent = computed(() => {
  if (!props.note.contenu) return ''
  
  // Conversion simple en HTML avec préservation des sauts de ligne
  return props.note.contenu
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/\n/g, '<br>')
    // Mettre en gras les sections (mots en MAJUSCULES suivis de :)
    .replace(/^([A-ZÀÂÄÉÈÊËÏÎÔÖÙÛÜŸÇÆŒ\s]+):$/gm, '<strong class="text-teal-700 dark:text-teal-400">$1:</strong>')
    // Mettre en évidence les éléments de liste
    .replace(/^[•\-*]\s(.+)$/gm, '<span class="text-gray-700 dark:text-gray-300">• $1</span>')
})

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
</script>