<template>
  <button
    @click="$emit('click')"
    class="flex items-center justify-between p-4 bg-white dark:bg-gray-800 rounded-lg 
           shadow hover:shadow-md transition-all duration-200 border border-gray-200 
           dark:border-gray-700 hover:border-blue-500 dark:hover:border-blue-400 
           group"
  >
    <div class="flex items-center space-x-3">
      <div class="p-2 bg-blue-50 dark:bg-blue-900/30 rounded-lg group-hover:bg-blue-100 
                  dark:group-hover:bg-blue-900/50 transition-colors">
        <component :is="iconComponent" class="w-5 h-5 text-blue-600 dark:text-blue-400" />
      </div>
      <span class="font-medium text-gray-900 dark:text-white">{{ title }}</span>
    </div>
    <svg class="w-5 h-5 text-gray-400 group-hover:text-blue-600 dark:group-hover:text-blue-400 
                group-hover:translate-x-1 transition-all" 
         fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
    </svg>
  </button>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  title: { type: String, required: true },
  icon: { type: String, required: true }
})

defineEmits(['click'])

const iconComponent = computed(() => {
  const icons = {
    'user-plus': UserPlusIcon,
    'file-text': FileTextIcon,
    search: SearchIcon
  }
  return icons[props.icon] || SearchIcon
})
</script>

<script>
const UserPlusIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/></svg>`
}

const FileTextIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>`
}

const SearchIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>`
}
</script>