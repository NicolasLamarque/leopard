<template>
  <div class="relative" ref="containerRef">

    <label v-if="label" class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">
      {{ label }}
    </label>

    <!-- Champ de recherche -->
    <div class="flex items-center gap-1.5 px-2.5 py-1.5 border rounded-md transition-all bg-white dark:bg-gray-800"
      :class="isFocused
        ? 'border-teal-500 ring-1 ring-teal-500/30'
        : 'border-gray-300 dark:border-gray-700'">

      <Search :size="13" class="text-gray-400 shrink-0" />

      <input
        ref="inputRef"
        type="text"
        v-model="searchQuery"
        :placeholder="placeholder"
        @focus="onFocus"
        @keydown.escape="close"
        @keydown.enter.prevent="selectHighlighted"
        @keydown.arrow-down.prevent="moveDown"
        @keydown.arrow-up.prevent="moveUp"
        class="flex-1 text-sm text-gray-900 dark:text-gray-100 bg-transparent outline-none placeholder-gray-400 dark:placeholder-gray-600 min-w-0"
      />

      <!-- Slot badge sélection (ex: drapeau, icône) -->
      <slot name="badge" :item="selectedItem" />

      <!-- Bouton effacer -->
      <button v-if="modelValue" @click.stop="clear" type="button"
        class="text-gray-400 hover:text-red-500 transition-colors shrink-0">
        <X :size="13" />
      </button>
    </div>

    <!-- Dropdown -->
    <Transition name="dropdown">
      <div v-if="showDropdown && filtered.length > 0"
        class="absolute z-50 w-full mt-1 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg shadow-xl overflow-hidden"
        style="max-height: 220px; overflow-y: auto;">

        <div v-for="(item, index) in filtered" :key="getKey(item)"
          @mousedown.prevent="select(item)"
          :class="[
            'px-3 py-2 cursor-pointer transition-colors border-b border-gray-50 dark:border-gray-700/50 last:border-0',
            index === highlightedIndex
              ? 'bg-teal-50 dark:bg-teal-900/20'
              : 'hover:bg-gray-50 dark:hover:bg-gray-700/50'
          ]">
          <!-- Slot item — personnalisable par parent -->
          <slot name="item" :item="item">
            <p class="text-sm font-medium text-gray-800 dark:text-gray-200">{{ getLabel(item) }}</p>
            <p v-if="getSublabel(item)" class="text-[11px] text-gray-400 mt-0.5">{{ getSublabel(item) }}</p>
          </slot>
        </div>

      </div>

      <!-- Aucun résultat -->
      <div v-else-if="showDropdown && searchQuery && filtered.length === 0"
        class="absolute z-50 w-full mt-1 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg shadow-xl px-3 py-3 text-xs text-gray-400 text-center">
        Aucun résultat pour « {{ searchQuery }} »
      </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { Search, X } from 'lucide-vue-next'

// ── Props ────────────────────────────────────────────────────────
const props = defineProps({
  // Valeur liée (l'ID stocké en DB)
  modelValue: { default: null },

  // La liste complète des items
  items: { type: Array, default: () => [] },

  // Clé unique de chaque item (ex: 'id')
  itemKey: { type: String, default: 'id' },

  // Champ(s) à afficher comme label principal (ex: 'nom' ou ['prenom', 'nom'])
  itemLabel: { type: [String, Array], default: 'nom' },

  // Champ secondaire (ex: 'ville') — optionnel
  itemSublabel: { type: String, default: null },

  // Champ(s) sur lesquels filtrer la recherche
  searchFields: { type: Array, default: null },

  label:       { type: String, default: null },
  placeholder: { type: String, default: 'Rechercher…' },

  // Nombre max de résultats sans recherche
  maxDefault: { type: Number, default: 6 },
})

const emit = defineEmits(['update:modelValue', 'selected'])

// ── État ─────────────────────────────────────────────────────────
const searchQuery      = ref('')
const showDropdown     = ref(false)
const isFocused        = ref(false)
const highlightedIndex = ref(-1)
const inputRef         = ref(null)
const containerRef     = ref(null)

// ── Helpers ──────────────────────────────────────────────────────
const getKey = (item) => item[props.itemKey]

const getLabel = (item) => {
  if (Array.isArray(props.itemLabel)) {
    return props.itemLabel.map(k => item[k]).filter(Boolean).join(' ')
  }
  return item[props.itemLabel] || ''
}

const getSublabel = (item) => {
  if (!props.itemSublabel) return null
  return item[props.itemSublabel] || null
}

const getSearchFields = () => {
  if (props.searchFields) return props.searchFields
  if (Array.isArray(props.itemLabel)) return props.itemLabel
  return [props.itemLabel]
}

// ── Item actuellement sélectionné ────────────────────────────────
const selectedItem = computed(() => {
  if (!props.modelValue) return null
  return props.items.find(i => i[props.itemKey] == props.modelValue) || null
})

// ── Filtre ───────────────────────────────────────────────────────
const filtered = computed(() => {
  const q = searchQuery.value.toLowerCase().trim()
  const fields = getSearchFields()

  if (!q) return props.items.slice(0, props.maxDefault)

  return props.items.filter(item =>
    fields.some(field => {
      const val = item[field]
      return val && String(val).toLowerCase().includes(q)
    })
  ).slice(0, 20)
})

// ── Sync label quand sélection externe change ─────────────────────
watch(() => props.modelValue, () => {
  if (selectedItem.value) {
    searchQuery.value = getLabel(selectedItem.value)
  } else {
    searchQuery.value = ''
  }
}, { immediate: true })

// ── Ouverture / fermeture ─────────────────────────────────────────
const onFocus = () => {
  isFocused.value    = true
  showDropdown.value = true
  highlightedIndex.value = -1
  if (selectedItem.value) searchQuery.value = ''
}

const close = () => {
  showDropdown.value = false
  isFocused.value    = false
  highlightedIndex.value = -1
  // Remet le label si une valeur est sélectionnée
  if (selectedItem.value) searchQuery.value = getLabel(selectedItem.value)
  else if (!props.modelValue) searchQuery.value = ''
}

// ── Sélection ─────────────────────────────────────────────────────
const select = (item) => {
  searchQuery.value = getLabel(item)
  showDropdown.value = false
  isFocused.value    = false
  emit('update:modelValue', item[props.itemKey])
  emit('selected', item)
}

const clear = () => {
  searchQuery.value = ''
  showDropdown.value = false
  emit('update:modelValue', null)
  emit('selected', null)
  inputRef.value?.focus()
}

// ── Navigation clavier ────────────────────────────────────────────
const moveDown = () => {
  if (!showDropdown.value) { showDropdown.value = true; return }
  highlightedIndex.value = Math.min(highlightedIndex.value + 1, filtered.value.length - 1)
}
const moveUp = () => {
  highlightedIndex.value = Math.max(highlightedIndex.value - 1, 0)
}
const selectHighlighted = () => {
  if (highlightedIndex.value >= 0 && filtered.value[highlightedIndex.value]) {
    select(filtered.value[highlightedIndex.value])
  }
}

// ── Clic extérieur → ferme ────────────────────────────────────────
const handleClickOutside = (e) => {
  if (containerRef.value && !containerRef.value.contains(e.target)) close()
}
onMounted(() => document.addEventListener('mousedown', handleClickOutside))
onUnmounted(() => document.removeEventListener('mousedown', handleClickOutside))
</script>

<style scoped>
.dropdown-enter-active, .dropdown-leave-active { transition: opacity 0.15s ease, transform 0.15s ease; }
.dropdown-enter-from, .dropdown-leave-to { opacity: 0; transform: translateY(-4px); }
</style>