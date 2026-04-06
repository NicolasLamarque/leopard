<template>
  <div class="space-y-2">

    <!-- Grille de selects -->
    <div :class="['grid gap-3', gridColsClass(field.columns || 2)]">
      <div v-for="sel in field.selects" :key="sel.id" class="space-y-1">
        <label class="block text-[10px] font-bold uppercase tracking-wider text-gray-500 dark:text-gray-400">
          {{ sel.label }}
        </label>
        <select
          :value="getValue(sel.id)"
          @change="onChange(sel.id, $event.target.value, sel)"
          class="w-full px-3 py-2 rounded-lg border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-900 text-sm text-gray-800 dark:text-gray-200 focus:outline-none focus:border-blue-400 dark:focus:border-blue-500 transition-colors"
        >
          <option value="">— Observer —</option>
          <option
            v-for="opt in sel.options"
            :key="opt.value"
            :value="opt.value"
            :data-type="opt.type"
          >
            {{ opt.label }}
          </option>
        </select>
      </div>
    </div>

    <!-- Messages cliniques actifs -->
    <TransitionGroup name="msg-list" tag="div" class="space-y-2 mt-2">
      <div
        v-for="msg in activeMessages"
        :key="msg.key"
        :class="['flex items-start gap-2.5 px-3 py-2.5 rounded-lg border text-xs', msgClass(msg.type)]"
      >
        <span class="text-base shrink-0 mt-0.5">{{ msgIcon(msg.type) }}</span>
        <div class="flex-1 min-w-0">
          <p class="font-semibold mb-0.5 leading-snug">{{ msg.title }}</p>
          <p class="opacity-80 leading-snug">{{ msg.text }}</p>
        </div>
        <!-- Badge type -->
        <span :class="['text-[9px] font-bold px-1.5 py-0.5 rounded-full uppercase tracking-wide shrink-0', msgBadgeClass(msg.type)]">
          {{ msg.type }}
        </span>
      </div>
    </TransitionGroup>

    <!-- Combinaisons détectées -->
    <div v-if="activeCombos.length" class="mt-2 space-y-1.5">
      <div
        v-for="combo in activeCombos"
        :key="combo.id"
        class="flex items-center gap-2 px-3 py-2 rounded-lg border border-purple-200 bg-purple-50 dark:border-purple-800/40 dark:bg-purple-950/20 text-xs"
      >
        <span class="text-purple-500 shrink-0">⚡</span>
        <span class="text-purple-700 dark:text-purple-300 font-medium">{{ combo.label }}</span>
      </div>
    </div>

  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  field:   Object,  // { selects: [...], columns: 2, combos: [...] }
  payload: Object
})

const emit = defineEmits(['update'])

// ── Valeurs ────────────────────────────────────────────────
const getValue = (selId) => props.payload?.[`${props.field.id}_${selId}`] || ''

const onChange = (selId, value, sel) => {
  emit('update', `${props.field.id}_${selId}`, value)
}

// ── Messages cliniques actifs ─────────────────────────────
const activeMessages = computed(() => {
  const msgs = []
  for (const sel of (props.field.selects || [])) {
    const val = getValue(sel.id)
    if (!val) continue
    const opt = sel.options?.find(o => o.value === val)
    if (opt?.message) {
      msgs.push({
        key:   `${sel.id}_${val}`,
        type:  opt.type || 'info',
        title: opt.label,
        text:  opt.message
      })
    }
  }
  return msgs
})

// ── Combinaisons ──────────────────────────────────────────
const activeCombos = computed(() => {
  const combos = props.field.combos || []
  return combos.filter(combo => {
    return combo.requires.every(req => {
      const [selId, optVal] = req.split(':')
      return getValue(selId) === optVal
    })
  })
})

// ── Classes CSS ──────────────────────────────────────────
const gridColsClass = (n) => ({
  1: 'grid-cols-1',
  2: 'grid-cols-1 sm:grid-cols-2',
  3: 'grid-cols-1 sm:grid-cols-2 lg:grid-cols-3',
  4: 'grid-cols-2 lg:grid-cols-4',
})[n] || 'grid-cols-2'

const msgClass = (type) => ({
  alerte:    'border-red-200 bg-red-50 text-red-800 dark:border-red-800/40 dark:bg-red-950/20 dark:text-red-300',
  attention: 'border-orange-200 bg-orange-50 text-orange-800 dark:border-orange-800/40 dark:bg-orange-950/20 dark:text-orange-300',
  info:      'border-blue-200 bg-blue-50 text-blue-800 dark:border-blue-800/40 dark:bg-blue-950/20 dark:text-blue-300',
  positif:   'border-emerald-200 bg-emerald-50 text-emerald-800 dark:border-emerald-800/40 dark:bg-emerald-950/20 dark:text-emerald-300',
})[type] || 'border-gray-200 bg-gray-50 text-gray-700'

const msgBadgeClass = (type) => ({
  alerte:    'bg-red-100 text-red-600 dark:bg-red-900/40 dark:text-red-400',
  attention: 'bg-orange-100 text-orange-600 dark:bg-orange-900/40 dark:text-orange-400',
  info:      'bg-blue-100 text-blue-600 dark:bg-blue-900/40 dark:text-blue-400',
  positif:   'bg-emerald-100 text-emerald-600 dark:bg-emerald-900/40 dark:text-emerald-400',
})[type] || 'bg-gray-100 text-gray-600'

const msgIcon = (type) => ({
  alerte:    '🚨',
  attention: '⚠️',
  info:      'ℹ️',
  positif:   '✅',
})[type] || 'ℹ️'
</script>

<style scoped>
.msg-list-enter-active, .msg-list-leave-active { transition: all 0.25s ease; }
.msg-list-enter-from, .msg-list-leave-to { opacity: 0; transform: translateY(-6px); }
</style>