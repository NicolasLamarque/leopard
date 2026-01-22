<!-- frontend/src/components/Medecins/MedecinSelector.vue -->
<template>
  <div class="md:col-span-2">
    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
      Médecin de famille
    </label>
    
    <div class="relative" ref="dropdownContainer">
      <div class="flex gap-2">
        <div class="flex-1 relative">
          <input 
            v-model="searchQuery"
            @focus="handleFocus"
            @input="filterMedecins"
            @keydown.down.prevent="navigateDown"
            @keydown.up.prevent="navigateUp"
            @keydown.enter.prevent="selectHighlighted"
            @keydown.escape="showDropdown = false"
            class="w-full px-4 py-2.5 pr-10 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
            :placeholder="selectedMedecin ? '' : 'Rechercher un médecin de famille...'"
          />
          
          <div 
            v-if="selectedMedecin && !searchQuery"
            class="absolute left-4 top-1/2 -translate-y-1/2 pointer-events-none text-gray-900 dark:text-white"
          >
            {{ selectedMedecin.displayText }}
          </div>
          
          <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-1">
            <button
              v-if="selectedMedecin"
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

        <button
          v-if="selectedMedecin"
          type="button"
          @click="viewMedecinDetails"
          class="px-4 py-2.5 bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 rounded-lg transition-colors"
          title="Voir la fiche du médecin"
        >
          <Eye :size="18" class="text-gray-700 dark:text-gray-300" />
        </button>
      </div>

      <Transition name="dropdown">
        <div 
          v-if="showDropdown && !isLoading"
          class="absolute z-50 w-full mt-2 bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-700 rounded-lg shadow-xl max-h-80 overflow-y-auto custom-scrollbar"
        >
          <div v-if="filteredMedecins.length > 0" class="p-2 space-y-1">
            <button
              v-for="(medecin, index) in filteredMedecins"
              :key="medecin.licence"
              type="button"
              @click="selectMedecin(medecin)"
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
                  <div class="font-semibold truncate">
                    {{ medecin.nomComplet || `${medecin.prenom} ${medecin.nom}` }}
                  </div>
                  <div class="text-sm text-gray-600 dark:text-gray-400 mt-0.5 flex items-center gap-1">
                    <Stethoscope :size="14" />
                    <span v-if="medecin.specialisation">{{ medecin.specialisation }}</span>
                    <span v-else class="text-gray-500 italic">Médecine générale</span>
                  </div>
                  <div class="text-xs text-gray-500 dark:text-gray-500 mt-1 flex items-center gap-3">
                    <span class="font-mono">{{ medecin.licence }}</span>
                    <span v-if="medecin.installation_principale" class="flex items-center gap-1">
                      <Building2 :size="12" />
                      {{ medecin.installation_principale }}
                    </span>
                  </div>
                </div>
                
                <div v-if="medecin.specialisation" class="flex-shrink-0">
                  <span class="inline-block px-2 py-1 bg-blue-100 dark:bg-blue-900/50 text-blue-700 dark:text-blue-300 text-xs rounded-full font-medium">
                    {{ getSpecialtyShort(medecin.specialisation) }}
                  </span>
                </div>
              </div>
            </button>
          </div>

          <div 
            v-else
            class="p-6 text-center text-gray-500 dark:text-gray-400"
          >
            <SearchX :size="48" class="mx-auto mb-2 opacity-50" />
            <p class="font-medium">Aucun médecin trouvé</p>
            <p class="text-sm mt-1">Essayez une autre recherche</p>
          </div>
        </div>
      </Transition>

      <div 
        v-if="isLoading"
        class="absolute z-50 w-full mt-2 bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-700 rounded-lg shadow-xl p-6 text-center"
      >
        <Loader2 :size="32" class="mx-auto animate-spin text-blue-500" />
        <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">Chargement des médecins...</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import { Eye, ChevronDown, X, Stethoscope, Building2, SearchX, Loader2, ExternalLink } from 'lucide-vue-next'

const props = defineProps({
  modelValue: { type: String, default: '' }
})

const emit = defineEmits(['update:modelValue', 'medecin-selected', 'view-details'])

const searchQuery = ref('')
const showDropdown = ref(false)
const highlightedIndex = ref(0)
const medecinsList = ref([])
const isLoading = ref(false)
const dropdownContainer = ref(null)

const selectedMedecin = computed(() => {
  if (!props.modelValue) return null
  return medecinsList.value.find(m => m.licence === props.modelValue)
})

const filteredMedecins = computed(() => {
  let medecins = [...medecinsList.value]
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    medecins = medecins.filter(m => {
      const nomComplet = (m.nomComplet || `${m.prenom} ${m.nom}`).toLowerCase()
      const licence = m.licence.toLowerCase()
      const specialisation = (m.specialisation || '').toLowerCase()
      
      return nomComplet.includes(query) || 
             licence.includes(query) || 
             specialisation.includes(query)
    })
  }
  
  // Trier: médecins de famille en premier
  return medecins.sort((a, b) => {
    const aIsFamille = isMedecinFamille(a)
    const bIsFamille = isMedecinFamille(b)
    if (aIsFamille && !bIsFamille) return -1
    if (!aIsFamille && bIsFamille) return 1
    return (a.nomComplet || `${a.prenom} ${a.nom}`).localeCompare(
      b.nomComplet || `${b.prenom} ${b.nom}`
    )
  })
})

const isMedecinFamille = (medecin) => {
  const spec = (medecin.specialisation || '').toLowerCase()
  return spec.includes('famille') || 
         spec.includes('général') || 
         spec.includes('omnipraticien') ||
         spec === ''
}

const getSpecialtyShort = (specialisation) => {
  if (!specialisation) return 'MG'
  if (specialisation.toLowerCase().includes('famille')) return 'MF'
  if (specialisation.toLowerCase().includes('général')) return 'MG'
  return specialisation.substring(0, 3).toUpperCase()
}

const loadMedecins = async () => {
  isLoading.value = true
  try {
    // Vérifier que window.go existe
    if (!window.go || !window.go.main || !window.go.main.App) {
      throw new Error('Wails runtime non disponible')
    }

    // Essayer différents noms de méthode possibles
    let response
    if (window.go.main.App.GetAllMedecins) {
      response = await window.go.main.App.GetAllMedecins()
    } else if (window.go.main.App.GetMedecins) {
      response = await window.go.main.App.GetMedecins()
    } else {
      throw new Error('Méthode GetAllMedecins non trouvée dans le backend')
    }

    console.log('📋 Médecins chargés:', response)
    
    // Vérifier que la réponse est un tableau
    if (!Array.isArray(response)) {
      console.error('⚠️ Réponse invalide:', response)
      throw new Error('La réponse n\'est pas un tableau')
    }
    
    // Filtrer seulement les médecins de famille (spécialité famille ou vide)
    medecinsList.value = response
      .filter(m => isMedecinFamille(m))
      .map(m => ({
        ...m,
        displayText: `${m.nomComplet || `${m.prenom} ${m.nom}`} (${m.licence})`
      }))
    
    console.log('✅ Médecins de famille filtrés:', medecinsList.value.length)
  } catch (error) {
    console.error('❌ Erreur chargement médecins:', error)
    console.error('Stack:', error.stack)
    medecinsList.value = []
    // Afficher l'erreur détaillée
    alert(`Erreur lors du chargement des médecins:\n${error.message}\n\nVérifiez la console pour plus de détails.`)
  } finally {
    isLoading.value = false
  }
}

const handleFocus = () => {
  showDropdown.value = true
  highlightedIndex.value = 0
  
  if (selectedMedecin.value) {
    const index = filteredMedecins.value.findIndex(
      m => m.licence === selectedMedecin.value.licence
    )
    if (index !== -1) highlightedIndex.value = index
  }
}

const filterMedecins = () => {
  highlightedIndex.value = 0
  showDropdown.value = true
}

const navigateDown = () => {
  if (highlightedIndex.value < filteredMedecins.value.length - 1) {
    highlightedIndex.value++
    scrollToHighlighted()
  }
}

const navigateUp = () => {
  if (highlightedIndex.value > 0) {
    highlightedIndex.value--
    scrollToHighlighted()
  }
}

const scrollToHighlighted = () => {
  const container = dropdownContainer.value?.querySelector('.overflow-y-auto')
  const highlighted = container?.children[0]?.children[highlightedIndex.value]
  if (highlighted && container) {
    highlighted.scrollIntoView({ block: 'nearest', behavior: 'smooth' })
  }
}

const selectHighlighted = () => {
  if (filteredMedecins.value[highlightedIndex.value]) {
    selectMedecin(filteredMedecins.value[highlightedIndex.value])
  }
}

const selectMedecin = (medecin) => {
  emit('update:modelValue', medecin.licence)
  emit('medecin-selected', medecin)
  searchQuery.value = ''
  showDropdown.value = false
  highlightedIndex.value = 0
}

const clearSelection = () => {
  emit('update:modelValue', '')
  emit('medecin-selected', null)
  searchQuery.value = ''
  highlightedIndex.value = 0
}

const viewMedecinDetails = () => {
  if (selectedMedecin.value) {
    emit('view-details', selectedMedecin.value)
  }
}
const openCMQProfile = () => {
  if (selectedMedecin.value) {
    const cmqUrl = `https://www.cmq.org/fr/bottin/medecins?number=${selectedMedecin.value.licence}&lastname=&firstname=&specialty=0&city=&unlisted=false`
    window.open(cmqUrl, '_blank')
  }
}

const handleClickOutside = (event) => {
  if (dropdownContainer.value && !dropdownContainer.value.contains(event.target)) {
    showDropdown.value = false
    searchQuery.value = ''
  }
}

onMounted(() => {
  loadMedecins()
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})

watch(() => props.modelValue, (newVal) => {
  if (!newVal) searchQuery.value = ''
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

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

.dark .custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #64748b;
}
</style>