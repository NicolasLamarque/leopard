<template>
  <div class="group flex flex-col gap-1">
    <p class="text-[10px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-wider flex items-center gap-1.5 transition-colors group-hover:text-slate-500">
      <component 
        v-if="icon" 
        :is="icon" 
        :size="12" 
        class="text-slate-300 dark:text-slate-600 group-hover:text-indigo-400 transition-colors"
      />
      {{ label }}
    </p>

    <p 
      :class="[
        'text-sm font-semibold truncate transition-all',
        highlight 
          ? 'text-indigo-600 dark:text-indigo-400' 
          : 'text-slate-700 dark:text-slate-200 group-hover:text-slate-900 dark:group-hover:text-white'
      ]"
      :title="value"
    >
      {{ value || '—' }}
    </p>
  </div>
</template>

<script setup>
/**
 * QuickInfo est un composant atomique utilisé pour afficher 
 * une paire Label/Valeur de manière élégante dans les fiches.
 */
defineProps({
  // Le titre (ex: "Ville", "Licence")
  label: {
    type: String,
    required: true
  },
  // La donnée à afficher
  value: {
    type: [String, Number],
    default: '—'
  },
  // Si vrai, la valeur s'affiche en couleur (ex: pour le statut Actif)
  highlight: {
    type: Boolean,
    default: false
  },
  // Composant Lucide-vue-next (ex: MapPin, Phone)
  icon: {
    type: Object,
    default: null
  }
})
</script>

<style scoped>
/* Un petit effet subtil si la valeur est trop longue */
.truncate {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>