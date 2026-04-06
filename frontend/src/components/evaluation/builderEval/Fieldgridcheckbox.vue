<template>
  <div class="space-y-3">

    <!-- Sous-groupes -->
    <div v-for="group in field.groups" :key="group.id" class="space-y-2">

      <!-- Label du groupe -->
      <div v-if="group.label" class="flex items-center gap-2">
        <span :class="['text-[10px] font-bold uppercase tracking-widest', groupLabelColor(group.type)]">
          {{ group.label }}
        </span>
        <div class="flex-1 h-px" :class="groupDividerColor(group.type)" />
        <!-- Score du groupe -->
        <span v-if="groupScore(group) > 0" :class="['text-[10px] font-mono font-bold px-2 py-0.5 rounded-full', groupBadgeClass(group.type)]">
          {{ groupScore(group) }} / {{ group.items.length }}
        </span>
      </div>

      <!-- Grille d'items -->
      <div :class="['grid gap-1.5', gridColsClass(field.columns || 2)]">
        <label
          v-for="item in group.items"
          :key="item.id"
          :class="[
            'flex items-start gap-2 px-3 py-2 rounded-lg border cursor-pointer transition-all text-xs select-none',
            itemClass(item, group.type, isChecked(item.id))
          ]"
        >
          <input
            type="checkbox"
            :checked="isChecked(item.id)"
            @change="toggle(item)"
            :class="['mt-0.5 shrink-0 w-3.5 h-3.5', checkAccent(group.type)]"
          />
          <span class="leading-snug">{{ item.label }}</span>
          <!-- Badge poids -->
          <span v-if="item.poids && item.poids !== 1" class="ml-auto shrink-0 text-[9px] font-mono opacity-50">×{{ item.poids }}</span>
        </label>
      </div>
    </div>

    <!-- Barre de score globale si scoring activé -->
    <div v-if="field.showScore && totalMax > 0" class="mt-3 p-3 rounded-xl border" :class="scoreContainerClass">
      <div class="flex items-center justify-between mb-2">
        <span class="text-xs font-semibold" :class="scoreLabelClass">Score global</span>
        <span class="text-sm font-bold font-mono" :class="scoreLabelClass">{{ totalScore }} / {{ totalMax }}</span>
      </div>
      <div class="h-2 rounded-full bg-gray-200 dark:bg-gray-700 overflow-hidden">
        <div
          class="h-full rounded-full transition-all duration-500"
          :class="scoreBarClass"
          :style="{ width: scorePercent + '%' }"
        />
      </div>
      <p v-if="scoreInterpretation" class="text-[10px] mt-1.5 font-medium" :class="scoreLabelClass">
        {{ scoreInterpretation }}
      </p>
    </div>

  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  field:   Object,  // { groups: [...], columns: 2, showScore: true }
  payload: Object   // { fieldId: ['item1', 'item2', ...] }
})

const emit = defineEmits(['update'])

// ── Valeur courante ────────────────────────────────────────
const checkedIds = computed(() => {
  const val = props.payload?.[props.field.id]
  return Array.isArray(val) ? val : []
})

const isChecked = (itemId) => checkedIds.value.includes(itemId)

const toggle = (item) => {
  const current = [...checkedIds.value]
  const idx = current.indexOf(item.id)
  if (idx === -1) current.push(item.id)
  else current.splice(idx, 1)
  emit('update', props.field.id, current)
}

// ── Scoring ────────────────────────────────────────────────
const groupScore = (group) => {
  return group.items.reduce((acc, item) => {
    if (!isChecked(item.id)) return acc
    const poids = item.poids || 1
    return group.type === 'prot' ? acc : acc + poids
  }, 0)
}

const totalScore = computed(() => {
  return (props.field.groups || []).reduce((acc, g) => {
    if (g.type === 'prot') return acc
    return acc + groupScore(g)
  }, 0)
})

const totalMax = computed(() => {
  return (props.field.groups || []).reduce((acc, g) => {
    if (g.type === 'prot') return acc
    return acc + g.items.reduce((a, i) => a + (i.poids || 1), 0)
  }, 0)
})

const scorePercent = computed(() =>
  totalMax.value > 0 ? Math.round((totalScore.value / totalMax.value) * 100) : 0
)

const scoreInterpretation = computed(() => {
  const thresholds = props.field.thresholds || []
  for (const t of [...thresholds].sort((a, b) => b.min - a.min)) {
    if (totalScore.value >= t.min) return t.label
  }
  return null
})

// ── Classes CSS ─────────────────────────────────────────────
const gridColsClass = (n) => ({
  1: 'grid-cols-1',
  2: 'grid-cols-1 sm:grid-cols-2',
  3: 'grid-cols-1 sm:grid-cols-2 lg:grid-cols-3',
  4: 'grid-cols-2 lg:grid-cols-4',
})[n] || 'grid-cols-2'

const groupLabelColor = (type) => ({
  vuln:    'text-red-500 dark:text-red-400',
  prot:    'text-emerald-600 dark:text-emerald-400',
  risk:    'text-orange-500 dark:text-orange-400',
  neutral: 'text-gray-500 dark:text-gray-400',
})[type] || 'text-gray-500'

const groupDividerColor = (type) => ({
  vuln:    'bg-red-200 dark:bg-red-900/40',
  prot:    'bg-emerald-200 dark:bg-emerald-900/40',
  risk:    'bg-orange-200 dark:bg-orange-900/40',
  neutral: 'bg-gray-200 dark:bg-gray-800',
})[type] || 'bg-gray-200'

const groupBadgeClass = (type) => ({
  vuln:    'bg-red-100 text-red-700 dark:bg-red-900/40 dark:text-red-300',
  prot:    'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/40 dark:text-emerald-300',
  risk:    'bg-orange-100 text-orange-700 dark:bg-orange-900/40 dark:text-orange-300',
  neutral: 'bg-gray-100 text-gray-700 dark:bg-gray-800 dark:text-gray-300',
})[type] || 'bg-gray-100 text-gray-700'

const checkAccent = (type) => ({
  vuln:    'accent-red-500',
  prot:    'accent-emerald-500',
  risk:    'accent-orange-500',
  neutral: 'accent-blue-500',
})[type] || 'accent-blue-500'

const itemClass = (item, type, checked) => {
  const base = {
    vuln: checked
      ? 'border-red-400 bg-red-50 text-red-700 dark:bg-red-950/40 dark:border-red-600 dark:text-red-300'
      : 'border-red-200 bg-red-50/50 text-gray-700 dark:border-red-900/50 dark:bg-red-950/20 dark:text-gray-300 hover:border-red-400 hover:bg-red-50',
    prot: checked
      ? 'border-emerald-400 bg-emerald-50 text-emerald-700 dark:bg-emerald-950/40 dark:border-emerald-600 dark:text-emerald-300'
      : 'border-emerald-200 bg-emerald-50/50 text-gray-700 dark:border-emerald-900/50 dark:bg-emerald-950/20 dark:text-gray-300 hover:border-emerald-400 hover:bg-emerald-50',
    risk: checked
      ? 'border-orange-400 bg-orange-50 text-orange-700 font-semibold dark:bg-orange-950/40 dark:border-orange-600 dark:text-orange-300'
      : 'border-orange-200 bg-orange-50/50 text-gray-700 dark:border-orange-900/50 dark:bg-orange-950/20 dark:text-gray-300 hover:border-orange-400',
    neutral: checked
      ? 'border-blue-400 bg-blue-50 text-blue-700 dark:bg-blue-950/40 dark:border-blue-600 dark:text-blue-300'
      : 'border-gray-200 bg-gray-50 text-gray-700 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-300 hover:border-blue-300',
  }
  return base[type] || base.neutral
}

// Score container
const scoreContainerClass = computed(() => {
  if (scorePercent.value >= 70) return 'border-red-200 bg-red-50 dark:border-red-900/40 dark:bg-red-950/20'
  if (scorePercent.value >= 40) return 'border-orange-200 bg-orange-50 dark:border-orange-900/40 dark:bg-orange-950/20'
  return 'border-emerald-200 bg-emerald-50 dark:border-emerald-900/40 dark:bg-emerald-950/20'
})
const scoreBarClass = computed(() => {
  if (scorePercent.value >= 70) return 'bg-red-500'
  if (scorePercent.value >= 40) return 'bg-orange-400'
  return 'bg-emerald-500'
})
const scoreLabelClass = computed(() => {
  if (scorePercent.value >= 70) return 'text-red-700 dark:text-red-300'
  if (scorePercent.value >= 40) return 'text-orange-700 dark:text-orange-300'
  return 'text-emerald-700 dark:text-emerald-300'
})
</script>