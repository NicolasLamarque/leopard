<template>
  <div class="flex items-center space-x-4 p-3 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors">
    <!-- Icône -->
    <div :class="iconBgClass" class="p-2 rounded-lg flex-shrink-0">
      <component :is="iconComponent" :class="iconColorClass" class="w-5 h-5" />
    </div>

    <!-- Contenu -->
    <div class="flex-1 min-w-0">
      <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
        {{ activity.description }}
      </p>
      <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
        {{ activity.time }}
      </p>
    </div>

    <!-- Badge de type -->
    <span :class="badgeClass" class="px-2 py-1 rounded-full text-xs font-medium flex-shrink-0">
      {{ typeLabel }}
    </span>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  activity: { 
    type: Object, 
    required: true 
  }
})

const typeLabel = computed(() => {
  const labels = {
    client: 'Client',
    note: 'Note',
    medecin: 'Médecin',
    appointment: 'RDV'
  }
  return labels[props.activity.type] || 'Activité'
})

const iconComponent = computed(() => {
  const icons = {
    'user-plus': UserPlusIcon,
    'file-text': FileTextIcon,
    stethoscope: StethoscopeIcon,
    calendar: CalendarIcon
  }
  return icons[props.activity.icon] || FileTextIcon
})

const iconBgClass = computed(() => {
  const classes = {
    client: 'bg-blue-100 dark:bg-blue-900/30',
    note: 'bg-green-100 dark:bg-green-900/30',
    medecin: 'bg-purple-100 dark:bg-purple-900/30',
    appointment: 'bg-amber-100 dark:bg-amber-900/30'
  }
  return classes[props.activity.type] || classes.note
})

const iconColorClass = computed(() => {
  const classes = {
    client: 'text-blue-600 dark:text-blue-400',
    note: 'text-green-600 dark:text-green-400',
    medecin: 'text-purple-600 dark:text-purple-400',
    appointment: 'text-amber-600 dark:text-amber-400'
  }
  return classes[props.activity.type] || classes.note
})

const badgeClass = computed(() => {
  const classes = {
    client: 'bg-blue-100 text-blue-700 dark:bg-blue-900/50 dark:text-blue-300',
    note: 'bg-green-100 text-green-700 dark:bg-green-900/50 dark:text-green-300',
    medecin: 'bg-purple-100 text-purple-700 dark:bg-purple-900/50 dark:text-purple-300',
    appointment: 'bg-amber-100 text-amber-700 dark:bg-amber-900/50 dark:text-amber-300'
  }
  return classes[props.activity.type] || classes.note
})
</script>

<script>
const UserPlusIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/></svg>`
}

const FileTextIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>`
}

const StethoscopeIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"/></svg>`
}

const CalendarIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>`
}
</script>