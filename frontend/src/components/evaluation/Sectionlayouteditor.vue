<template>
  <!-- Éditeur de layout pour une section du builder -->
  <div class="space-y-4">

    <!-- ── Sélecteur de layout ── -->
    <div>
      <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest mb-2">Disposition des colonnes</p>
      <div class="grid grid-cols-3 gap-2">
        <button
          v-for="layout in layouts"
          :key="layout.id"
          @click="$emit('update-layout', layout.id)"
          :class="[
            'flex flex-col items-center gap-1.5 p-3 rounded-xl border transition-all',
            currentLayout === layout.id
              ? 'border-teal-500 bg-teal-500/10 text-teal-400'
              : 'border-gray-700 text-gray-500 hover:border-gray-500 hover:text-gray-300'
          ]"
        >
          <!-- Schéma visuel -->
          <div class="flex gap-0.5 w-full h-6">
            <div
              v-for="col in layout.schema"
              :key="col"
              :style="{ flex: col }"
              class="rounded-sm bg-current opacity-40"
            />
          </div>
          <span class="text-[10px] font-semibold">{{ layout.label }}</span>
        </button>
      </div>
    </div>

    <!-- ── Répartition des champs dans les colonnes ── -->
    <div v-if="section.fields?.length">
      <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest mb-2">Champs → Colonnes</p>

      <div :class="['grid gap-3', previewGridClass]">
        <div
          v-for="(col, cIdx) in previewColumns"
          :key="cIdx"
          class="border border-dashed border-gray-700 rounded-xl p-2 space-y-1.5 min-h-[60px]"
          @dragover.prevent="dragOverCol = cIdx"
          @drop="dropFieldToCol($event, cIdx)"
          :class="dragOverCol === cIdx ? 'border-teal-500 bg-teal-950/20' : ''"
        >
          <p class="text-[9px] font-bold text-gray-600 uppercase tracking-wider px-1">Col. {{ cIdx + 1 }}</p>
          <div
            v-for="field in col.fields"
            :key="field.id"
            draggable="true"
            @dragstart="dragStartField($event, field.id, cIdx)"
            :class="['flex items-center gap-2 px-2 py-1.5 rounded-lg border bg-gray-900 cursor-grab select-none',
              draggingField?.id === field.id ? 'opacity-30 border-teal-500' : 'border-gray-800']"
          >
            <GripVertical :size="11" class="text-gray-600 shrink-0" />
            <span :class="['text-[9px] font-bold px-1.5 py-0.5 rounded-full shrink-0', typeBadge(field.type)]">{{ typeLabel(field.type) }}</span>
            <span class="text-xs text-gray-400 truncate">{{ field.label || field.contenu || 'Sans titre' }}</span>
          </div>
          <div v-if="!col.fields.length" class="text-[10px] text-gray-700 text-center py-3">Déposer ici</div>
        </div>
      </div>
    </div>

    <!-- ── Options avancées de la section ── -->
    <div class="space-y-2">
      <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest">Style de la section</p>
      <div class="grid grid-cols-2 gap-2">
        <label class="flex items-center gap-2 text-xs text-gray-400 cursor-pointer">
          <input type="checkbox" :checked="section.showBanner !== false" @change="$emit('toggle-banner', $event.target.checked)" class="accent-teal-500" />
          Afficher le bandeau
        </label>
        <label class="flex items-center gap-2 text-xs text-gray-400 cursor-pointer">
          <input type="checkbox" :checked="section.divider" @change="$emit('toggle-divider', $event.target.checked)" class="accent-teal-500" />
          Séparateur après
        </label>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { GripVertical } from 'lucide-vue-next'
import SectionLayout from './builderEval/SectionLayout.vue'

const props = defineProps({
  section:       Object,
  currentLayout: String
})

const emit = defineEmits(['update-layout', 'move-field-to-col', 'toggle-banner', 'toggle-divider'])

// ── Layouts disponibles ─────────────────────────────────────
const layouts = [
  { id: 'single',        label: '1 colonne',    schema: [1] },
  { id: '2col',          label: '2 égales',     schema: [1, 1] },
  { id: '2col-right',    label: 'Large + étroit', schema: [2, 1] },
  { id: '3col',          label: '3 colonnes',   schema: [1, 1, 1] },
  { id: 'sidebar-right', label: 'Contenu + sidebar', schema: [2, 1] },
  { id: 'sidebar-left',  label: 'Sidebar + contenu', schema: [1, 2] },
]

// ── Colonnes preview ────────────────────────────────────────
const colCount = computed(() => {
  const map = { single: 1, '2col': 2, '2col-right': 2, '3col': 3, 'sidebar-right': 2, 'sidebar-left': 2 }
  return map[props.currentLayout || 'single'] || 1
})

const previewColumns = computed(() => {
  const cols = []
  for (let i = 0; i < colCount.value; i++) {
    cols.push({ fields: [] })
  }
  // Distribuer les champs selon leur colIndex
  ;(props.section.fields || []).forEach(field => {
    const colIdx = Math.min(field.colIndex || 0, colCount.value - 1)
    cols[colIdx].fields.push(field)
  })
  return cols
})

const previewGridClass = computed(() => ({
  1: 'grid-cols-1',
  2: 'grid-cols-2',
  3: 'grid-cols-3',
})[colCount.value] || 'grid-cols-1')

// ── Drag & Drop entre colonnes ──────────────────────────────
const draggingField = ref(null)
const dragOverCol   = ref(null)

const dragStartField = (e, fieldId, fromCol) => {
  draggingField.value = { id: fieldId, fromCol }
  e.dataTransfer.effectAllowed = 'move'
}

const dropFieldToCol = (e, toCol) => {
  if (!draggingField.value) return
  emit('move-field-to-col', { fieldId: draggingField.value.id, toCol })
  draggingField.value = null
  dragOverCol.value   = null
}

// ── Helpers badges ─────────────────────────────────────────
const typeBadge = (t) => ({
  textarea:      'bg-blue-500/20 text-blue-300',
  text:          'bg-teal-500/20 text-teal-300',
  date:          'bg-purple-500/20 text-purple-300',
  select:        'bg-amber-500/20 text-amber-300',
  checkboxes:    'bg-emerald-500/20 text-emerald-300',
  table:         'bg-pink-500/20 text-pink-300',
  info:          'bg-slate-500/20 text-slate-300',
  grid_checkbox: 'bg-red-500/20 text-red-300',
  obs_select:    'bg-orange-500/20 text-orange-300',
  score_radar:   'bg-violet-500/20 text-violet-300',
})[t] || 'bg-gray-500/20 text-gray-400'

const typeLabel = (t) => ({
  textarea: 'TXT', text: 'TXT', date: 'DATE', select: 'LIST',
  checkboxes: 'CHK', table: 'TAB', info: 'NFO',
  grid_checkbox: 'GRILLE', obs_select: 'OBS', score_radar: 'RADAR',
})[t] || t?.slice(0, 4).toUpperCase()
</script>