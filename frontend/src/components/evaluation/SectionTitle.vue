<!-- SectionTitle.vue — Titre de section dans le formulaire -->
<template>
  <div class="flex items-center gap-3 pb-4 border-b dark:border-gray-800">
    <div class="p-2 bg-slate-100 dark:bg-slate-800 rounded-lg">
      <component :is="iconComponent" :size="18" class="text-slate-600 dark:text-slate-400" />
    </div>
    <h3 class="text-xl font-bold text-gray-900 dark:text-white">{{ label }}</h3>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import * as LucideIcons from 'lucide-vue-next'

const props = defineProps({
  icon:  String,
  label: String
})

const iconComponent = computed(() => LucideIcons[props.icon] || LucideIcons.FileText)
</script>