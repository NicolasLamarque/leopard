<!-- frontend/src/components/Notaires/NotaireSelector.vue -->
<template>
  <div class="md:col-span-2">
    <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
      Notaire
    </label>
    
    <div class="relative" ref="dropdownContainer">
      <div class="flex gap-2">
        <div class="flex-1 relative">
          <input 
            v-model="searchQuery"
            @focus="handleFocus"
            @input="filterNotaires"
            @keydown.down.prevent="navigateDown"
            @keydown.up.prevent="navigateUp"
            @keydown.enter.prevent="selectHighlighted"
            @keydown.escape="showDropdown = false"
            class="w-full px-4 py-2.5 pr-10 border-2 border-slate-300 dark:border-slate-700 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-white focus:border-slate-500 focus:ring-2 focus:ring-slate-500/20 outline-none transition-all"
            :placeholder="selectedNotaire ? '' : 'Rechercher un notaire...'"
          />
          
          <div 
            v-if="selectedNotaire && !searchQuery"
            class="absolute left-4 top-1/2 -translate-y-1/2 pointer-events-none text-slate-900 dark:text-white"
          >
            {{ selectedNotaire.displayText }}
          </div>
          
          <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-1">
            <button
              v-if="selectedNotaire"
              @click.stop="clearSelection"
              type="button"
              class="p-1 hover:bg-slate-100 dark:hover:bg-slate-700 rounded transition-colors"
              title="Effacer"
            >
              <X :size="16" class="text-slate-500" />
            </button>
            <ChevronDown 
              :size="18" 
              class="text-slate-500 transition-transform"
              :class="{ 'rotate-180': showDropdown }"
            />
          </div>
        </div>

        <button
          v-if="selectedNotaire"
          type="button"
          @click="viewNotaireDetails"
          class="px-4 py-2.5 bg-slate-100 dark:bg-slate-800 hover:bg-slate-200 dark:hover:bg-slate-700 rounded-lg transition-colors"
          title="Voir la fiche du notaire"
        >
          <Eye :size="18" class="text-slate-700 dark:text-slate-300" />
        </button>
      </div>

      <Transition name="dropdown">
        <div 
          v-if="showDropdown && !isLoading"
          class="absolute z-50 w-full mt-2 bg-white dark:bg-slate-800 border-2 border-slate-300 dark:border-slate-700 rounded-lg shadow-xl max-h-80 overflow-y-auto custom-scrollbar"
        >
          <div v-if="filteredNotaires.length > 0" class="p-2 space-y-1">
            <button
              v-for="(notaire, index) in filteredNotaires"
              :key="notaire.id"
              type="button"
              @click="selectNotaire(notaire)"
              @mouseenter="highlightedIndex = index"
              :class="[
                'w-full text-left px-4 py-3 rounded-lg transition-all',
                highlightedIndex === index
                  ? 'bg-slate-100 dark:bg-slate-700 text-slate-900 dark:text-white'
                  : 'hover:bg-slate-50 dark:hover:bg-slate-800/50 text-slate-900 dark:text-white'
              ]"
            >
              <div class="flex items-start justify-between gap-3">
                <div class="flex-1 min-w-0">
                  <div class="font-semibold truncate">
                    {{ notaire.civilite }} {{ notaire.prenom }} {{ notaire.nom }}
                  </div>
                  <div class="text-sm text-slate-600 dark:text-slate-400 mt-0.5 flex items-center gap-1">
                    <Scale :size="14" />
                    <span v-if="notaire.secteur_activite">{{ notaire.secteur_activite }}</span>
                    <span v-else class="text-slate-500 italic">Notaire</span>
                  </div>
                  <div class="text-xs text-slate-500 dark:text-slate-500 mt-1 flex items-center gap-3">
                    <span v-if="notaire.ville" class="flex items-center gap-1">
                      <MapPin :size="12" />
                      {{ notaire.ville }}
                    </span>
                    <span v-if="notaire.telephone" class="flex items-center gap-1">
                      <Phone :size="12" />
                      {{ notaire.telephone }}
                    </span>
                  </div>
                </div>
                
                <div v-if="notaire.secteur_activite" class="flex-shrink-0">
                  <span class="inline-block px-2 py-1 bg-slate-100 dark:bg-slate-700 text-slate-700 dark:text-slate-300 text-xs rounded-full font-medium">
                    {{ getSecteurShort(notaire.secteur_activite) }}
                  </span>
                </div>
              </div>
            </button>
          </div>

          <div 
            v-else
            class="p-6 text-center text-slate-500 dark:text-slate-400"
          >
            <SearchX :size="48" class="mx-auto mb-2 opacity-50" />
            <p class="font-medium">Aucun notaire trouvé</p>
            <p class="text-sm mt-1">Essayez une autre recherche</p>
          </div>
        </div>
      </Transition>

      <div 
        v-if="isLoading"
        class="absolute z-50 w-full mt-2 bg-white dark:bg-slate-800 border-2 border-slate-300 dark:border-slate-700 rounded-lg shadow-xl p-6 text-center"
      >
        <Loader2 :size="32" class="mx-auto animate-spin text-slate-500" />
        <p class="mt-2 text-sm text-slate-600 dark:text-slate-400">Chargement des notaires...</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import { Eye, ChevronDown, X, Scale, MapPin, SearchX, Loader2, Phone } from 'lucide-vue-next'

const props = defineProps({
  modelValue: { type: Number, default: null }
})

const emit = defineEmits(['update:modelValue', 'notaire-selected', 'view-details'])

const searchQuery = ref('')
const showDropdown = ref(false)
const highlightedIndex = ref(0)
const notairesList = ref([])
const isLoading = ref(false)
const dropdownContainer = ref(null)

const selectedNotaire = computed(() => {
  if (!props.modelValue) return null
  return notairesList.value.find(n => n.id === props.modelValue)
})

const filteredNotaires = computed(() => {
  let notaires = [...notairesList.value]
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    notaires = notaires.filter(n => {
      const nomComplet = `${n.prenom} ${n.nom}`.toLowerCase()
      const ville = (n.ville || '').toLowerCase()
      const secteur = (n.secteur_activite || '').toLowerCase()
      
      return nomComplet.includes(query) || 
             ville.includes(query) || 
             secteur.includes(query)
    })
  }
  
  // Trier par nom
  return notaires.sort((a, b) => 
    `${a.prenom} ${a.nom}`.localeCompare(`${b.prenom} ${b.nom}`)
  )
})

const getSecteurShort = (secteur) => {
  if (!secteur) return 'N'
  const words = secteur.split(' ')
  if (words.length > 1) {
    return words.map(w => w.charAt(0)).join('').toUpperCase().substring(0, 3)
  }
  return secteur.substring(0, 3).toUpperCase()
}

const loadNotaires = async () => {
  isLoading.value = true
  try {
    if (!window.go || !window.go.main || !window.go.main.App) {
      throw new Error('Wails runtime non disponible')
    }

    const response = await window.go.main.App.GetAllNotaires()
    
    console.log('📋 Notaires chargés:', response)
    
    if (!Array.isArray(response)) {
      console.error('⚠️ Réponse invalide:', response)
      throw new Error('La réponse n\'est pas un tableau')
    }
    
    notairesList.value = response.map(n => ({
      ...n,
      displayText: `${n.civilite || 'Me'} ${n.prenom} ${n.nom}${n.ville ? ' - ' + n.ville : ''}`
    }))
    
    console.log('✅ Notaires formatés:', notairesList.value.length)
  } catch (error) {
    console.error('❌ Erreur chargement notaires:', error)
    console.error('Stack:', error.stack)
    notairesList.value = []
    alert(`Erreur lors du chargement des notaires:\n${error.message}\n\nVérifiez la console pour plus de détails.`)
  } finally {
    isLoading.value = false
  }
}

const handleFocus = () => {
  showDropdown.value = true
  highlightedIndex.value = 0
  
  if (selectedNotaire.value) {
    const index = filteredNotaires.value.findIndex(
      n => n.id === selectedNotaire.value.id
    )
    if (index !== -1) highlightedIndex.value = index
  }
}

const filterNotaires = () => {
  highlightedIndex.value = 0
  showDropdown.value = true
}

const navigateDown = () => {
  if (highlightedIndex.value < filteredNotaires.value.length - 1) {
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
  if (filteredNotaires.value[highlightedIndex.value]) {
    selectNotaire(filteredNotaires.value[highlightedIndex.value])
  }
}

const selectNotaire = (notaire) => {
  emit('update:modelValue', notaire.id)
  emit('notaire-selected', notaire)
  searchQuery.value = ''
  showDropdown.value = false
  highlightedIndex.value = 0
}

const clearSelection = () => {
  emit('update:modelValue', null)
  emit('notaire-selected', null)
  searchQuery.value = ''
  highlightedIndex.value = 0
}

const viewNotaireDetails = () => {
  if (selectedNotaire.value) {
    emit('view-details', selectedNotaire.value)
  }
}

const handleClickOutside = (event) => {
  if (dropdownContainer.value && !dropdownContainer.value.contains(event.target)) {
    showDropdown.value = false
    searchQuery.value = ''
  }
}

onMounted(() => {
  loadNotaires()
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