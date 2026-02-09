<!-- frontend/src/components/appels/AppelsFooter.vue -->

<template>
  <div class="px-6 py-4 border-t border-stone-200 dark:border-stone-800 bg-gradient-to-r from-stone-50 via-pink-50/30 to-stone-50 dark:from-stone-900 dark:via-stone-800 dark:to-stone-900 flex items-center justify-between">
    
    <!-- Gauche - Infos -->
    <div class="flex items-center gap-3 text-sm text-stone-600 dark:text-stone-400">
      <div v-if="appel?.created_at" class="flex items-center gap-2">
        <Clock :size="14" />
        <span>Créé {{ formatDate(appel.created_at) }}</span>
      </div>
      <div v-if="appel?.updated_at && appel.updated_at !== appel.created_at" class="flex items-center gap-2">
        <span>•</span>
        <span>Modifié {{ formatDate(appel.updated_at) }}</span>
      </div>
    </div>

    <!-- Droite - Actions -->
    <div class="flex items-center gap-3">
      
      <!-- Bouton Supprimer (si édition) -->
      <button
        v-if="isEditing && !isSaving"
        @click="$emit('delete')"
        type="button"
        class="px-4 py-2 text-sm font-medium text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-xl transition-all duration-200 flex items-center gap-2"
      >
        <Trash2 :size="16" />
        Supprimer
      </button>

      <!-- Bouton Annuler -->
      <button
        @click="$emit('cancel')"
        type="button"
        :disabled="isSaving"
        class="px-5 py-2.5 text-sm font-medium text-stone-700 dark:text-stone-300 bg-white dark:bg-stone-800 border border-stone-300 dark:border-stone-700 rounded-xl hover:bg-stone-50 dark:hover:bg-stone-700 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
      >
        <X :size="16" />
        Annuler
      </button>

      <!-- Bouton Sauvegarder -->
      <button
        @click="$emit('save')"
        type="button"
        :disabled="isSaving || !canSave"
        class="px-5 py-2.5 text-sm font-medium text-white bg-gradient-to-r from-pink-500 to-rose-500 hover:from-pink-600 hover:to-rose-600 rounded-xl shadow-md shadow-pink-500/20 hover:shadow-lg hover:shadow-pink-500/30 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
      >
        <Save :size="16" v-if="!isSaving" />
        <Loader2 :size="16" class="animate-spin" v-else />
        {{ isSaving ? 'Enregistrement...' : (isEditing ? 'Mettre à jour' : 'Enregistrer') }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, Save, X, Trash2, Loader2 } from 'lucide-vue-next'

const props = defineProps({
  appel: {
    type: Object,
    default: null
  },
  isSaving: {
    type: Boolean,
    default: false
  },
  canSave: {
    type: Boolean,
    default: true
  }
})

defineEmits(['save', 'cancel', 'delete'])

const isEditing = computed(() => !!props.appel?.id)

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now - date
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMins / 60)
  const diffDays = Math.floor(diffHours / 24)

  if (diffMins < 1) return 'à l\'instant'
  if (diffMins < 60) return `il y a ${diffMins} min`
  if (diffHours < 24) return `il y a ${diffHours}h`
  if (diffDays < 7) return `il y a ${diffDays}j`
  
  return date.toLocaleDateString('fr-CA', { 
    day: 'numeric', 
    month: 'short',
    year: date.getFullYear() !== now.getFullYear() ? 'numeric' : undefined
  })
}
</script>