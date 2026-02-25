<!-- FormField.vue — Enveloppe champ de formulaire avec label, erreur, help -->
<template>
  <div class="space-y-1.5">
    <div class="flex items-baseline gap-2">
      <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300">
        {{ label }}
        <span v-if="required" class="text-red-500 ml-0.5">*</span>
      </label>
      <span v-if="error" class="text-red-500 text-xs font-medium">({{ error }})</span>
    </div>

    <p v-if="help" class="text-xs text-gray-400 dark:text-gray-500 -mt-0.5">{{ help }}</p>

    <slot />
  </div>
</template>

<script setup>
defineProps({
  label:    String,
  error:    String,
  help:     String,
  required: Boolean
})
</script>