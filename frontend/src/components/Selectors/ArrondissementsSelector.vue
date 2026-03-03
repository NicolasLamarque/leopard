<template>
  <BaseSelector
    v-model="localModelValue"
    :items="municipalities"
    item-key="mcode"
    item-label="munnom"
    item-sublabel="mrc"
    label="Municipalité"
    placeholder="Rechercher une ville ou MRC..."
    :search-fields="['munnom', 'mrc', 'mcode']"
    @selected="onSelect"
  >
    <template #item="{ item }">
      <div class="flex justify-between items-center w-full">
        <div class="flex flex-col">
          <span class="text-sm font-bold text-gray-900 dark:text-gray-100">{{ item.munnom }}</span>
          <span class="text-[10px] text-gray-500 uppercase tracking-tighter">{{ item.mrc }}</span>
        </div>
        <span class="text-[10px] font-mono text-gray-400 bg-gray-100 dark:bg-gray-800 px-1 rounded">
          {{ item.mcode }}
        </span>
      </div>
    </template>
  </BaseSelector>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import BaseSelector from '@/components/shared/BaseSelector.vue'
import { GetAllMuns } from '../../wailsjs/go/main/App'

const props = defineProps(['modelValue'])
const emit = defineEmits(['update:modelValue', 'change'])

const municipalities = ref([])
const localModelValue = ref(props.modelValue)

const onSelect = (mun) => {
  emit('update:modelValue', mun.mcode)
  emit('change', mun)
}

watch(() => props.modelValue, (newVal) => {
  localModelValue.value = newVal
})

onMounted(async () => {
  try {
    municipalities.value = await GetAllMuns()
  } catch (err) {
    console.error("Erreur chargement municipalités:", err)
  }
})
</script>