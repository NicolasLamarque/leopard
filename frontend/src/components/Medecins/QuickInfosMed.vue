<!-- frontend/src/components/Medecins/QuickInfosMed.vue -->
<template>
  <div class="group">
    <p class="text-xs font-medium text-slate-500 dark:text-slate-400 mb-1 transition-colors flex items-center gap-1.5">
      <component 
        v-if="icon" 
        :is="icon" 
        :size="14" 
        class="text-slate-400 dark:text-slate-500"
      />
      {{ label }}
    </p>
    <p 
      :class="[
        'text-sm font-semibold transition-colors',
        highlight 
          ? 'text-slate-900 dark:text-white' 
          : 'text-slate-700 dark:text-slate-300'
      ]"
    >
      {{ value }}
    </p>
  </div>
</template>

<script setup>
defineProps({
  label: { type: String, required: true },
  value: { type: String, required: true },
  highlight: { type: Boolean, default: false },
  icon: { type: Object, default: null }
})
</script>