<!-- frontend/src/components/Medecins/ContactRow.vue -->
<template>
  <div 
    class="flex items-center gap-4 p-4 rounded-lg transition-all duration-200 hover:bg-slate-50 dark:hover:bg-slate-800/50 group"
  >
    <!-- Icône -->
    <div 
      :class="[
        'p-2.5 rounded-lg transition-colors flex-shrink-0',
        colorClasses.iconBg
      ]"
    >
      <component :is="icon" :size="18" class="text-white" />
    </div>

    <!-- Contenu -->
    <div class="flex-1 min-w-0">
      <p class="text-xs font-medium text-slate-500 dark:text-slate-400 mb-0.5">
        {{ label }}
      </p>
      <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
        {{ value }}
      </p>
    </div>

    <!-- Action button -->
    <a 
      :href="href"
      :class="[
        'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200 flex items-center gap-2 flex-shrink-0',
        'opacity-70 group-hover:opacity-100',
        colorClasses.button
      ]"
    >
      {{ actionLabel }}
    </a>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  icon: { type: Object, required: true },
  label: { type: String, required: true },
  value: { type: String, required: true },
  href: { type: String, required: true },
  actionLabel: { type: String, required: true },
  color: { 
    type: String, 
    default: 'slate',
    validator: (val) => ['slate', 'stone', 'zinc'].includes(val)
  }
})

const colorClasses = computed(() => {
  const colors = {
    slate: {
      iconBg: 'bg-slate-600',
      button: 'bg-slate-700 hover:bg-slate-800 text-white'
    },
    stone: {
      iconBg: 'bg-stone-600',
      button: 'bg-stone-700 hover:bg-stone-800 text-white'
    },
    zinc: {
      iconBg: 'bg-zinc-600',
      button: 'bg-zinc-700 hover:bg-zinc-800 text-white'
    }
  }
  return colors[props.color]
})
</script>