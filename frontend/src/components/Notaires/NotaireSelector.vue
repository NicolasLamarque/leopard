<template>
  <BaseSelector
    v-model="props.modelValue"
    :items="notaires"
    item-key="id"
    :item-label="['prenom', 'nom']"
    item-sublabel="ville"
    :search-fields="['nom', 'prenom', 'ville']"
    label="Notaire"
    placeholder="Rechercher un notaire…"
    @update:modelValue="$emit('update:modelValue', $event)"
    @selected="$emit('notaire-selected', $event)"
  >
    <!-- Badge : icône dans le champ une fois sélectionné -->
    <template #badge="{ item }">
      <span v-if="item" class="text-[10px] font-bold text-purple-600 dark:text-purple-400 shrink-0">Me</span>
    </template>

    <!-- Ligne dans le dropdown -->
    <template #item="{ item }">
      <p class="text-sm font-semibold text-gray-800 dark:text-gray-200">
        Me {{ item.prenom }} {{ item.nom }}
      </p>
      <p class="text-[11px] text-gray-400 mt-0.5 flex items-center gap-1">
        <MapPin :size="10" /> {{ item.ville || 'Ville inconnue' }}
      </p>
    </template>
  </BaseSelector>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { MapPin } from 'lucide-vue-next'
import { GetAllNotaires } from '../../../wailsjs/go/main/App'
import BaseSelector from '@/components/shared/BaseSelector.vue'

const props = defineProps(['modelValue'])
const emit  = defineEmits(['update:modelValue', 'notaire-selected'])

const notaires = ref([])

onMounted(async () => {
  try { notaires.value = (await GetAllNotaires()) || [] } catch {}
})
</script>