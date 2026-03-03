<template>
  <div class="flex flex-col gap-2">

    <!-- Sélecteur Municipalité -->
    <BaseSelector
      v-model="selectedMcode"
      :items="municipalities"
      item-key="mcode"
      item-label="munnom"
      item-sublabel="mrc"
      label="Municipalité"
      placeholder="Commencez à taper (ex: Shawinigan...)"
      :search-fields="['munnom', 'mrc', 'mcode']"
      @selected="onMunSelect"
    >
      <template #item="{ item }">
        <div class="flex items-center justify-between group">
          <div class="flex flex-col">
            <span class="text-sm font-semibold text-gray-900 dark:text-gray-100 group-hover:text-teal-600 transition-colors">
              {{ item.munnom }}
            </span>
            <span class="text-[10px] text-gray-400 uppercase tracking-tighter">
              MRC : {{ item.mrc || 'N/A' }}
            </span>
          </div>
          <div class="text-[9px] font-mono text-gray-300 dark:text-gray-600 border border-gray-100 dark:border-gray-700 px-1 rounded">
            {{ item.mcode }}
          </div>
        </div>
      </template>
      <template #badge="{ item }">
        <div v-if="item" class="flex items-center px-1.5 py-0.5 bg-teal-100 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300 text-[10px] rounded font-bold">
          QC
        </div>
      </template>
    </BaseSelector>

    <!-- Arrondissement — apparaît seulement si la mun en a -->
    <Transition name="fade">
      <BaseSelector
        v-if="arrondissements.length > 0"
        v-model="selectedArrcod"
        :items="arrondissements"
        item-key="arrcod"
        item-label="arrnom"
        label="Arrondissement / Secteur"
        placeholder="Sélectionner un secteur..."
        :search-fields="['arrnom', 'arrcod']"
        @selected="onArrSelect"
      >
        <template #item="{ item }">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-800 dark:text-gray-200">{{ item.arrnom }}</span>
            <span class="text-[9px] font-mono text-gray-400 px-1">{{ item.arrcod }}</span>
          </div>
        </template>
      </BaseSelector>
    </Transition>

    <div v-if="loadingArr" class="flex items-center gap-1.5 text-[11px] text-gray-400">
      <Loader2 class="w-3 h-3 animate-spin" />
      Chargement des secteurs...
    </div>

  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { Loader2 } from 'lucide-vue-next'
import BaseSelector from '@/components/shared/BaseSelector.vue'
import { GetMunicipalities, GetArrondissements } from '../../../wailsjs/go/main/App'

const props = defineProps({
  modelValue: { type: String, default: null }, // mcode
  arrValue:   { type: String, default: null }, // arrcod (optionnel)
})

const emit = defineEmits(['update:modelValue', 'update:arrValue', 'change'])

const municipalities  = ref([])
const arrondissements = ref([])
const selectedMcode   = ref(props.modelValue)
const selectedArrcod  = ref(props.arrValue)
const loadingArr      = ref(false)

const onMunSelect = async (mun) => {
  selectedArrcod.value  = null
  arrondissements.value = []
  emit('update:modelValue', mun.mcode)
  emit('update:arrValue', null)

  loadingArr.value = true
  try {
    arrondissements.value = (await GetArrondissements(mun.mcode)) ?? []
  } catch (e) {
    console.error('Erreur arrondissements:', e)
  } finally {
    loadingArr.value = false
  }

  emit('change', { mun, arrondissement: null })
}

const onArrSelect = (arr) => {
  emit('update:arrValue', arr.arrcod)
  emit('change', {
    mun: municipalities.value.find(m => m.mcode === selectedMcode.value) ?? null,
    arrondissement: arr,
  })
}

watch(() => props.modelValue, async (mcode) => {
  selectedMcode.value   = mcode
  arrondissements.value = []
  if (mcode) {
    arrondissements.value = (await GetArrondissements(mcode)) ?? []
  }
})

watch(() => props.arrValue, (v) => { selectedArrcod.value = v })

onMounted(async () => {
  try {
    municipalities.value = await GetMunicipalities()
    if (props.modelValue) {
      arrondissements.value = (await GetArrondissements(props.modelValue)) ?? []
    }
  } catch (e) {
    console.error('Erreur chargement municipalités:', e)
  }
})
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s, transform 0.2s; }
.fade-enter-from, .fade-leave-to       { opacity: 0; transform: translateY(-4px); }
</style>