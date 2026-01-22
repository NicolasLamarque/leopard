<!-- frontend/src/components/RPA/RPASelector.vue -->
<template>
  <div class="md:col-span-3">
    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
      Résidence (RPA)
    </label>
    
    <div class="relative" ref="dropdownContainer">
      <!-- Barre de recherche et filtres -->
      <div class="grid grid-cols-12 gap-2 mb-2">
        <!-- Filtre région - Compact -->
        <select 
          v-model="selectedRegion"
          @change="handleRegionChange"
          class="col-span-3 px-2 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 outline-none text-xs"
        >
          <option value="">📍 Toutes</option>
          <option v-for="region in regions" :key="region" :value="region">
            {{ formatRegionFull(region) }}
          </option>
        </select>

        <!-- Input recherche - Large -->
        <div class="col-span-9 relative">
          <input 
            v-model="searchQuery"
            @focus="showDropdown = true"
            @input="handleSearch"
            @keydown.down.prevent="navigateDown"
            @keydown.up.prevent="navigateUp"
            @keydown.enter.prevent="selectHighlighted"
            @keydown.escape="showDropdown = false"
            class="w-full px-4 py-2.5 pr-10 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
            :placeholder="selectedName ? '' : '🔍 Rechercher par nom, ville, registre...'"
          />
          
          <!-- Affichage sélection -->
          <div 
            v-if="selectedName && !searchQuery"
            class="absolute left-4 top-1/2 -translate-y-1/2 pointer-events-none text-gray-900 dark:text-white font-medium"
          >
            {{ selectedName }}
          </div>

          <!-- Boutons -->
          <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-1">
            <button
              v-if="selectedRPA"
              @click.stop="clearSelection"
              type="button"
              class="p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded transition-colors"
              title="Effacer"
            >
              <X :size="16" class="text-gray-500" />
            </button>
            <ChevronDown 
              :size="18" 
              class="text-gray-500 transition-transform"
              :class="{ 'rotate-180': showDropdown }"
            />
          </div>
        </div>
      </div>

      <!-- Dropdown résultats -->
      <Transition name="dropdown">
        <div 
          v-if="showDropdown && !isLoading"
          class="absolute z-50 w-full mt-1 bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-700 rounded-lg shadow-xl max-h-96 overflow-y-auto custom-scrollbar"
        >
          <!-- Compteur résultats -->
          <div class="sticky top-0 bg-gray-50 dark:bg-gray-900 px-4 py-2 border-b border-gray-200 dark:border-gray-700 text-xs text-gray-600 dark:text-gray-400">
            {{ filteredList.length }} résidence{{ filteredList.length > 1 ? 's' : '' }} trouvée{{ filteredList.length > 1 ? 's' : '' }}
            <span v-if="selectedRegion" class="ml-2 text-blue-600 dark:text-blue-400">
              (Région: {{ selectedRegion }})
            </span>
          </div>

          <div v-if="filteredList.length > 0" class="p-2 space-y-1">
            <button
              v-for="(rpa, index) in filteredList"
              :key="rpa.id"
              type="button"
              @click="select(rpa)"
              @mouseenter="highlightedIndex = index"
              :class="[
                'w-full text-left px-4 py-3 rounded-lg transition-all',
                highlightedIndex === index
                  ? 'bg-blue-50 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300'
                  : 'hover:bg-gray-100 dark:hover:bg-gray-700 text-gray-900 dark:text-white'
              ]"
            >
              <div class="flex items-start justify-between gap-3">
                <div class="flex-1 min-w-0">
                  <div class="font-semibold truncate text-base">
                    {{ rpa.titre }}
                  </div>
                  <div class="flex items-center gap-3 mt-1 text-xs text-gray-600 dark:text-gray-400">
                    <span class="flex items-center gap-1">
                      <MapPin :size="12" />
                      {{ rpa.municipalite || rpa.ville }}
                    </span>
                    <span v-if="rpa.capacite" class="flex items-center gap-1">
                      <Users :size="12" />
                      {{ rpa.capacite }}
                    </span>
                    <span v-if="rpa.registre" class="font-mono text-gray-500">
                      #{{ rpa.registre }}
                    </span>
                  </div>
                </div>
                
                <!-- Badge région compact -->
                <div v-if="rpa.region" class="flex-shrink-0">
                  <span class="inline-block px-2 py-1 bg-blue-100 dark:bg-blue-900/50 text-blue-700 dark:text-blue-300 text-xs rounded-full font-medium">
                    {{ formatRegion(rpa.region) }}
                  </span>
                </div>
              </div>
            </button>
          </div>

          <!-- Aucun résultat -->
          <div v-else class="p-8 text-center text-gray-500 dark:text-gray-400">
            <SearchX :size="48" class="mx-auto mb-2 opacity-50" />
            <p class="font-medium">Aucune résidence trouvée</p>
            <p class="text-sm mt-1">
              {{ searchQuery ? `Aucun résultat pour "${searchQuery}"` : 'Essayez de changer la région ou la recherche' }}
            </p>
          </div>
        </div>
      </Transition>

      <!-- Chargement -->
      <div 
        v-if="isLoading"
        class="absolute z-50 w-full mt-1 bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-700 rounded-lg shadow-xl p-6 text-center"
      >
        <Loader2 :size="32" class="mx-auto animate-spin text-blue-500" />
        <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">Chargement des résidences...</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { X, ChevronDown, MapPin, Users, SearchX, Loader2 } from 'lucide-vue-next'

const props = defineProps({
  modelValue: { type: Number, default: null }
})

const emit = defineEmits(['update:modelValue', 'rpa-selected'])

// Mapping des régions du Québec
const REGIONS_QC = {
  '01': 'Bas-Saint-Laurent',
  '02': 'Saguenay–Lac-Saint-Jean',
  '03': 'Capitale-Nationale',
  '04': 'Mauricie',
  '05': 'Estrie',
  '06': 'Montréal',
  '07': 'Outaouais',
  '08': 'Abitibi-Témiscamingue',
  '09': 'Côte-Nord',
  '10': 'Nord-du-Québec',
  '11': 'Gaspésie–Îles-de-la-Madeleine',
  '12': 'Chaudière-Appalaches',
  '13': 'Laval',
  '14': 'Lanaudière',
  '15': 'Laurentides',
  '16': 'Montérégie',
  '17': 'Centre-du-Québec'
}

// Formater la région pour affichage
const formatRegion = (region) => {
  if (!region) return ''
  // Si c'est un code numérique (01, 02, etc.)
  const cleaned = region.trim()
  if (REGIONS_QC[cleaned]) {
    return REGIONS_QC[cleaned]
  }
  // Si déjà au format "04 - Mauricie", extraire le nom
  if (region.includes(' - ')) {
    return region.split(' - ')[1]
  }
  return region
}

// Formater pour le dropdown (avec code)
const formatRegionFull = (region) => {
  const cleaned = region.trim()
  if (REGIONS_QC[cleaned]) {
    return `${cleaned} - ${REGIONS_QC[cleaned]}`
  }
  return region
}

const list = ref([])
const searchQuery = ref('')
const selectedRegion = ref('')
const showDropdown = ref(false)
const highlightedIndex = ref(0)
const isLoading = ref(false)
const dropdownContainer = ref(null)

// Liste unique des régions
const regions = computed(() => {
  const regionSet = new Set()
  list.value.forEach(rpa => {
    if (rpa.region) regionSet.add(rpa.region)
  })
  return Array.from(regionSet).sort()
})

// Liste filtrée par région ET recherche
const filteredList = computed(() => {
  let filtered = [...list.value]
  
  // Filtre par région
  if (selectedRegion.value) {
    filtered = filtered.filter(rpa => rpa.region === selectedRegion.value)
  }
  
  // Filtre par recherche
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(rpa => {
      const titre = (rpa.titre || '').toLowerCase()
      const municipalite = (rpa.municipalite || rpa.ville || '').toLowerCase()
      const registre = (rpa.registre || '').toLowerCase()
      const adresse = (rpa.adresse || '').toLowerCase()
      
      return titre.includes(query) || 
             municipalite.includes(query) || 
             registre.includes(query) ||
             adresse.includes(query)
    })
  }
  
  // Trier par nom
  return filtered.sort((a, b) => 
    (a.titre || '').localeCompare(b.titre || '')
  ).slice(0, 50) // Limiter à 50 résultats
})

// RPA sélectionnée
const selectedRPA = computed(() => {
  if (!props.modelValue) return null
  return list.value.find(rpa => rpa.id === props.modelValue)
})

const selectedName = computed(() => {
  return selectedRPA.value ? selectedRPA.value.titre : ''
})

// Charger les RPA
const loadRPAs = async () => {
  isLoading.value = true
  try {
    const res = await window.go.main.App.GetResidences()
    console.log('📋 RPA chargées:', res.length)
    list.value = res || []
  } catch (error) {
    console.error('❌ Erreur chargement RPA:', error)
    alert('Erreur lors du chargement des résidences')
  } finally {
    isLoading.value = false
  }
}

const handleRegionChange = () => {
  highlightedIndex.value = 0
  showDropdown.value = true
}

const handleSearch = () => {
  highlightedIndex.value = 0
  showDropdown.value = true
}

const navigateDown = () => {
  if (highlightedIndex.value < filteredList.value.length - 1) {
    highlightedIndex.value++
  }
}

const navigateUp = () => {
  if (highlightedIndex.value > 0) {
    highlightedIndex.value--
  }
}

const selectHighlighted = () => {
  if (filteredList.value[highlightedIndex.value]) {
    select(filteredList.value[highlightedIndex.value])
  }
}

const select = (rpa) => {
  emit('update:modelValue', rpa.id)
  emit('rpa-selected', rpa)
  searchQuery.value = ''
  showDropdown.value = false
  highlightedIndex.value = 0
}

const clearSelection = () => {
  emit('update:modelValue', null)
  emit('rpa-selected', null)
  searchQuery.value = ''
  selectedRegion.value = ''
  highlightedIndex.value = 0
}

const handleClickOutside = (event) => {
  if (dropdownContainer.value && !dropdownContainer.value.contains(event.target)) {
    showDropdown.value = false
  }
}

onMounted(() => {
  loadRPAs()
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.2s ease;
}

.dropdown-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.dropdown-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 4px;
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: #475569;
}
</style>