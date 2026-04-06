<template>
  <!-- Layout de section : 1, 2, ou 3 colonnes configurables -->
  <div :class="sectionWrapClass">

    <!-- Bandeau section -->
    <div v-if="section.label" :class="['flex items-center gap-3 px-5 py-3 rounded-xl border mb-4', bannerClass]">
      <component :is="iconComponent(section.icone)" :size="17" />
      <span class="font-bold text-sm">{{ section.label }}</span>
      <!-- Badge layout -->
      <span class="ml-auto text-[9px] font-mono opacity-50">{{ layoutLabel }}</span>
    </div>

    <!-- Grille de colonnes -->
    <div :class="columnsClass">
      <div
        v-for="(col, cIdx) in columns"
        :key="cIdx"
        :class="['space-y-4', col.span ? `col-span-${col.span}` : '']"
      >
        <!-- Champs de la colonne -->
        <div
          v-for="field in col.fields"
          :key="field.id"
          :class="['space-y-1.5', fieldWrapClass(field)]"
        >
          <!-- Label + aide (sauf info et grille) -->
          <template v-if="field.type !== 'info' && field.type !== 'grid_checkbox' && field.type !== 'obs_select' && field.type !== 'score_radar'">
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300">
              {{ field.label }}
              <span v-if="field.required" class="text-red-500 ml-1">*</span>
            </label>
            <p v-if="field.help" class="text-xs text-gray-400">{{ field.help }}</p>
          </template>

          <!-- ── Rendu selon le type ── -->

          <!-- INFO -->
          <div v-if="field.type === 'info'"
            class="p-4 bg-slate-50 dark:bg-slate-900/50 border border-slate-200 dark:border-slate-700 rounded-xl text-sm text-slate-600 dark:text-slate-400">
            {{ field.contenu }}
          </div>

          <!-- TEXTAREA -->
          <textarea v-else-if="field.type === 'textarea'"
            :value="payload[field.id] || ''"
            @input="$emit('update', field.id, $event.target.value)"
            :rows="field.rows || 5"
            :placeholder="field.placeholder || ''"
            class="w-full px-4 py-3 rounded-xl border-2 border-gray-200 dark:border-gray-700 dark:bg-gray-900 dark:text-white focus:outline-none focus:border-teal-500 text-sm resize-none transition-colors"
          />

          <!-- TEXT -->
          <input v-else-if="field.type === 'text'"
            type="text"
            :value="payload[field.id] || ''"
            @input="$emit('update', field.id, $event.target.value)"
            :placeholder="field.placeholder || ''"
            class="w-full px-4 py-2.5 rounded-xl border-2 border-gray-200 dark:border-gray-700 dark:bg-gray-900 dark:text-white focus:outline-none focus:border-teal-500 text-sm transition-colors"
          />

          <!-- DATE -->
          <input v-else-if="field.type === 'date'"
            type="date"
            :value="payload[field.id] || ''"
            @input="$emit('update', field.id, $event.target.value)"
            class="px-4 py-2.5 rounded-xl border-2 border-gray-200 dark:border-gray-700 dark:bg-gray-900 dark:text-white focus:outline-none focus:border-teal-500 text-sm transition-colors"
          />

          <!-- SELECT -->
          <select v-else-if="field.type === 'select'"
            :value="payload[field.id] || ''"
            @change="$emit('update', field.id, $event.target.value)"
            class="w-full px-4 py-2.5 rounded-xl border-2 border-gray-200 dark:border-gray-700 dark:bg-gray-900 dark:text-white focus:outline-none focus:border-teal-500 text-sm transition-colors"
          >
            <option value="">— Sélectionner —</option>
            <option v-for="opt in (field.options || [])" :key="opt" :value="opt">{{ opt }}</option>
          </select>

          <!-- CHECKBOXES simples -->
          <div v-else-if="field.type === 'checkboxes'" class="grid grid-cols-1 sm:grid-cols-2 gap-2">
            <label v-for="opt in (field.options || [])" :key="opt"
              :class="['flex items-center gap-2.5 p-2.5 rounded-lg border cursor-pointer transition-all text-sm',
                isChecked(field.id, opt)
                  ? 'border-teal-400 bg-teal-50 dark:bg-teal-900/20 dark:border-teal-600 text-teal-700 dark:text-teal-300'
                  : 'border-gray-200 dark:border-gray-700 hover:border-teal-300 dark:text-gray-300']"
            >
              <input type="checkbox" :checked="isChecked(field.id, opt)" @change="toggleCheckbox(field.id, opt)" class="accent-teal-500 w-4 h-4 shrink-0" />
              {{ opt }}
            </label>
          </div>

          <!-- TABLE -->
          <div v-else-if="field.type === 'table'" class="border dark:border-gray-700 rounded-xl overflow-hidden">
            <div class="flex items-center justify-between px-3 py-2 bg-gray-50 dark:bg-gray-800 border-b dark:border-gray-700">
              <span class="text-xs font-bold text-gray-500 uppercase tracking-wider">{{ field.label }}</span>
              <button @click="addTableRow(field)" class="text-xs text-teal-500 hover:text-teal-400 flex items-center gap-1">
                <Plus :size="12" /> Ajouter
              </button>
            </div>
            <div v-if="!getTableRows(field.id).length" class="px-4 py-5 text-center text-xs text-gray-400">Aucune entrée</div>
            <div v-for="(row, rIdx) in getTableRows(field.id)" :key="rIdx"
              class="grid border-b last:border-0 dark:border-gray-700 hover:bg-gray-50/50 dark:hover:bg-gray-900/30"
              :style="{ gridTemplateColumns: `repeat(${(field.columns||[]).length}, 1fr) 32px` }"
            >
              <div v-for="col in (field.columns || [])" :key="col.id" class="px-2 py-1.5">
                <input :type="col.type === 'date' ? 'date' : 'text'" :value="row[col.id] || ''"
                  @input="updateTableCell(field.id, rIdx, col.id, $event.target.value)"
                  class="w-full text-xs px-2 py-1.5 rounded-lg bg-transparent border-transparent hover:border-gray-200 dark:hover:border-gray-700 focus:border-teal-400 dark:text-white focus:outline-none transition-all" />
              </div>
              <div class="flex items-center justify-center">
                <button @click="removeTableRow(field.id, rIdx)" class="p-1 text-gray-300 hover:text-red-400 transition-colors"><Trash2 :size="11" /></button>
              </div>
            </div>
          </div>

          <!-- GRILLE CLINIQUE (FieldGridCheckbox) -->
          <FieldGridCheckbox
            v-else-if="field.type === 'grid_checkbox'"
            :field="field"
            :payload="payload"
            @update="(fId, val) => $emit('update', fId, val)"
          />

          <!-- SELECT OBSERVATIONNEL (FieldObsSelect) -->
          <FieldObsSelect
            v-else-if="field.type === 'obs_select'"
            :field="field"
            :payload="payload"
            @update="(fId, val) => $emit('update', fId, val)"
          />

          <!-- RADAR SCORE (FieldScoreRadar) -->
          <FieldScoreRadar
            v-else-if="field.type === 'score_radar'"
            :field="field"
            :payload="payload"
            :all-payload="payload"
            @update="(fId, val) => $emit('update', fId, val)"
          />

        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import {
  FileText, Brain, Heart, Activity, Briefcase, Users,
  Layers, CheckCircle, Target, Network, ClipboardList,
  Scale, ShieldCheck, Plus, Trash2
} from 'lucide-vue-next'
import FieldGridCheckbox from './FieldGridCheckbox.vue'
import FieldObsSelect    from './FieldObsSelect.vue'
import FieldScoreRadar   from './FieldScoreRadar.vue'

const props = defineProps({
  section: Object,  // section du modèle avec layout
  payload: Object,  // payload courant
})

const emit = defineEmits(['update'])

// ── Layout ─────────────────────────────────────────────────
// section.layout: 'single' | '2col' | '2col-right' | '3col' | 'sidebar-right' | 'sidebar-left'
// section.columns: [{ fields: [...], span?: 2 }, ...]

const columns = computed(() => {
  // Si les colonnes sont déjà définies explicitement
  if (props.section.columns?.length) return props.section.columns

  // Sinon : tous les champs dans une seule colonne
  return [{ fields: props.section.fields || [] }]
})

const layout = computed(() => props.section.layout || 'single')

const columnsClass = computed(() => ({
  'single':        'grid grid-cols-1 gap-4',
  '2col':          'grid grid-cols-2 gap-6',
  '2col-right':    'grid grid-cols-2 gap-6',
  '3col':          'grid grid-cols-3 gap-4',
  'sidebar-right': 'grid grid-cols-3 gap-6',
  'sidebar-left':  'grid grid-cols-3 gap-6',
})[layout.value] || 'grid grid-cols-1 gap-4')

const layoutLabel = computed(() => ({
  'single':        '1 col',
  '2col':          '2 col',
  '2col-right':    '2 col ⬅',
  '3col':          '3 col',
  'sidebar-right': '2/3 + 1/3',
  'sidebar-left':  '1/3 + 2/3',
})[layout.value] || '1 col')

const sectionWrapClass = computed(() => 'space-y-0')

// ── Bannière ───────────────────────────────────────────────
const colorBannerMap = {
  teal:    'bg-teal-50 dark:bg-teal-900/20 border-teal-200 dark:border-teal-800 text-teal-700 dark:text-teal-300',
  blue:    'bg-blue-50 dark:bg-blue-900/20 border-blue-200 dark:border-blue-800 text-blue-700 dark:text-blue-300',
  purple:  'bg-purple-50 dark:bg-purple-900/20 border-purple-200 dark:border-purple-800 text-purple-700 dark:text-purple-300',
  amber:   'bg-amber-50 dark:bg-amber-900/20 border-amber-200 dark:border-amber-800 text-amber-700 dark:text-amber-300',
  emerald: 'bg-emerald-50 dark:bg-emerald-900/20 border-emerald-200 dark:border-emerald-800 text-emerald-700 dark:text-emerald-300',
  red:     'bg-red-50 dark:bg-red-900/20 border-red-200 dark:border-red-800 text-red-700 dark:text-red-300',
  slate:   'bg-slate-50 dark:bg-slate-900/50 border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-400',
}
const bannerClass = computed(() => colorBannerMap[props.section.couleur] || colorBannerMap.teal)

// ── Icônes ─────────────────────────────────────────────────
const iconMap = { FileText, Brain, Heart, Activity, Briefcase, Users, Layers, CheckCircle, Target, Network, ClipboardList, Scale, ShieldCheck }
const iconComponent = (n) => iconMap[n] || FileText

// ── Champs wrapper ─────────────────────────────────────────
const fieldWrapClass = (field) => {
  // Certains types prennent toute la largeur naturellement
  const fullWidth = ['score_radar', 'grid_checkbox', 'obs_select', 'table']
  return fullWidth.includes(field.type) ? '' : ''
}

// ── Checkboxes simples ─────────────────────────────────────
const isChecked = (fId, opt) => {
  const val = props.payload?.[fId]
  return Array.isArray(val) && val.includes(opt)
}
const toggleCheckbox = (fId, opt) => {
  const current = Array.isArray(props.payload?.[fId]) ? [...props.payload[fId]] : []
  const idx = current.indexOf(opt)
  if (idx === -1) current.push(opt)
  else current.splice(idx, 1)
  emit('update', fId, current)
}

// ── Tables ─────────────────────────────────────────────────
const getTableRows = (fId) => {
  const val = props.payload?.[fId]
  return Array.isArray(val) ? val : []
}
const addTableRow = (field) => {
  const rows = [...getTableRows(field.id)]
  const row = {}
  ;(field.columns || []).forEach(col => { row[col.id] = '' })
  rows.push(row)
  emit('update', field.id, rows)
}
const removeTableRow = (fId, idx) => {
  const rows = [...getTableRows(fId)]
  rows.splice(idx, 1)
  emit('update', fId, rows)
}
const updateTableCell = (fId, rIdx, colId, val) => {
  const rows = [...getTableRows(fId)]
  rows[rIdx] = { ...rows[rIdx], [colId]: val }
  emit('update', fId, rows)
}
</script>