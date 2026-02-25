<template>
  <div class="space-y-2">
    <FormKit
      type="list"
      :label="label || 'Pharmacie rattachée'"
      v-model="selectedId"
    >
      <div class="relative">
        <div class="flex gap-2">
          <div class="relative flex-1">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-400">
              <Search :size="18" />
            </span>
            <input
              type="text"
              v-model="searchQuery"
              placeholder="Rechercher par nom ou ville..."
              class="w-full pl-10 pr-4 py-2 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none transition-all"
              @focus="showDropdown = true"
            />
          </div>
          
          <button 
            v-if="selectedId"
            type="button"
            @click="clearSelection"
            class="px-3 py-2 text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
          >
            <X :size="18" />
          </button>
        </div>

        <div 
          v-if="showDropdown && filteredPharmacies.length > 0"
          class="absolute z-50 w-full mt-1 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg shadow-xl max-h-60 overflow-y-auto"
        >
          <div
            v-for="pharma in filteredPharmacies"
            :key="pharma.id"
            @click="selectPharmacie(pharma)"
            class="p-3 hover:bg-emerald-50 dark:hover:bg-emerald-900/20 cursor-pointer border-b last:border-0 border-gray-100 dark:border-gray-700 transition-colors"
          >
            <div class="font-bold text-gray-900 dark:text-white">{{ pharma.NomOrganisation }}</div>
            <div class="text-xs text-gray-500 flex items-center gap-1">
              <MapPin :size="12" /> {{ pharma.Ville }} ({{ pharma.Banniere || 'Indépendante' }})
            </div>
          </div>
        </div>
      </div>
    </FormKit>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Search, MapPin, X } from 'lucide-vue-next'
import { GetAllPharmacies } from '../../wailsjs/go/main/App'

const props = defineProps(['modelValue', 'label'])
const emit = defineEmits(['update:modelValue', 'change'])

const pharmacies = ref([])
const searchQuery = ref('')
const showDropdown = ref(false)
const selectedId = ref(props.modelValue)

const filteredPharmacies = computed(() => {
  if (!searchQuery.value) return pharmacies.value.slice(0, 5) // Top 5 par défaut
  const q = searchQuery.value.toLowerCase()
  return pharmacies.value.filter(p => 
    p.NomOrganisation.toLowerCase().includes(q) || 
    p.Ville?.toLowerCase().includes(q)
  )
})

const selectPharmacie = (pharma) => {
  selectedId.value = pharma.id
  searchQuery.value = pharma.NomOrganisation
  showDropdown.value = false
  emit('update:modelValue', pharma.id)
  emit('change', pharma)
}

const clearSelection = () => {
  selectedId.value = null
  searchQuery.value = ''
  emit('update:modelValue', null)
}

onMounted(async () => {
  try {
    const data = await GetAllPharmacies()
    pharmacies.value = data || []
    
    // Si on a déjà un ID, on remplit le libellé
    if (selectedId.value) {
      const p = pharmacies.value.find(x => x.id === selectedId.value)
      if (p) searchQuery.value = p.NomOrganisation
    }
  } catch (err) {
    console.error("Erreur chargement sélecteur:", err)
  }
})
</script>