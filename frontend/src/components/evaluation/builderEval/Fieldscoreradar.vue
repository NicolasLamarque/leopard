<template>
  <div class="space-y-4">

    <!-- Radar -->
    <div class="relative flex items-center justify-center">
      <canvas ref="chartCanvas" :width="size" :height="size" />

      <!-- Overlay score central -->
      <div class="absolute inset-0 flex flex-col items-center justify-center pointer-events-none">
        <span class="text-2xl font-bold font-mono" :class="globalScoreColor">{{ globalScore }}</span>
        <span class="text-[10px] text-gray-400 uppercase tracking-wider">score global</span>
      </div>
    </div>

    <!-- Légende / détail des axes -->
    <div :class="['grid gap-2', axes.length > 4 ? 'grid-cols-2' : 'grid-cols-1']">
      <div
        v-for="axis in axesWithScores"
        :key="axis.id"
        class="flex items-center gap-2 px-3 py-2 rounded-lg bg-gray-50 dark:bg-gray-900/50 border border-gray-100 dark:border-gray-800"
      >
        <div class="w-2.5 h-2.5 rounded-full shrink-0" :style="{ background: axis.color }" />
        <span class="text-xs text-gray-600 dark:text-gray-400 flex-1 truncate">{{ axis.label }}</span>
        <span class="text-xs font-bold font-mono" :style="{ color: axis.color }">{{ axis.score }}/{{ axis.max }}</span>
      </div>
    </div>

    <!-- Interprétations seuil -->
    <div v-if="activeAlerts.length" class="space-y-1.5">
      <div
        v-for="alert in activeAlerts"
        :key="alert.id"
        :class="['flex items-start gap-2 px-3 py-2 rounded-lg border text-xs', alertClass(alert.level)]"
      >
        <span class="shrink-0">{{ alertIcon(alert.level) }}</span>
        <span>{{ alert.message }}</span>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  field:   Object,  // { axes: [...], size: 280, sourceFields: {...} }
  payload: Object,
  allPayload: Object  // payload complet du formulaire pour accéder aux autres champs
})

const chartCanvas = ref(null)
let chartInstance = null

const size = computed(() => props.field.size || 280)

// ── Axes avec scores calculés ──────────────────────────────
const axes = computed(() => props.field.axes || [])

const axesWithScores = computed(() => {
  return axes.value.map(axis => {
    const score = computeAxisScore(axis)
    return { ...axis, score, max: axis.max || 10 }
  })
})

const computeAxisScore = (axis) => {
  // L'axe peut calculer son score depuis d'autres champs du payload
  if (!axis.sourceField) return axis.manualScore || 0

  const sourceVal = props.allPayload?.[axis.sourceField]

  // Si c'est une liste de checkboxes (array)
  if (Array.isArray(sourceVal)) {
    const checkedCount = sourceVal.length
    const max = axis.max || 10
    return Math.min(checkedCount, max)
  }

  // Si c'est un nombre direct
  if (typeof sourceVal === 'number') return Math.min(sourceVal, axis.max || 10)

  return 0
}

const globalScore = computed(() => {
  const total = axesWithScores.value.reduce((a, ax) => a + ax.score, 0)
  const max   = axesWithScores.value.reduce((a, ax) => a + ax.max, 0)
  return max > 0 ? Math.round((total / max) * 100) : 0
})

const globalScoreColor = computed(() => {
  if (globalScore.value >= 70) return 'text-red-500 dark:text-red-400'
  if (globalScore.value >= 40) return 'text-orange-500 dark:text-orange-400'
  return 'text-emerald-500 dark:text-emerald-400'
})

// ── Alertes par seuil ─────────────────────────────────────
const activeAlerts = computed(() => {
  const alerts = []
  for (const axis of axesWithScores.value) {
    const thresholds = axis.thresholds || []
    for (const t of [...thresholds].sort((a, b) => b.min - a.min)) {
      if (axis.score >= t.min) {
        alerts.push({ id: `${axis.id}_${t.min}`, level: t.level || 'attention', message: t.message })
        break
      }
    }
  }
  return alerts
})

// ── Chart.js ──────────────────────────────────────────────
const COLORS = [
  '#3b82f6', '#ef4444', '#10b981', '#f59e0b',
  '#8b5cf6', '#06b6d4', '#f97316', '#84cc16'
]

const buildChart = () => {
  if (!chartCanvas.value || typeof Chart === 'undefined') return
  if (chartInstance) { chartInstance.destroy(); chartInstance = null }

  const isDark = document.documentElement.classList.contains('dark')
  const gridColor   = isDark ? 'rgba(255,255,255,0.08)' : 'rgba(0,0,0,0.08)'
  const labelColor  = isDark ? '#94a3b8' : '#64748b'

  const scores = axesWithScores.value.map(ax => ax.score)
  const maxes  = axesWithScores.value.map(ax => ax.max)
  const labels = axesWithScores.value.map(ax => ax.label)

  // Normaliser 0-10 pour le radar
  const normalized = scores.map((s, i) => maxes[i] > 0 ? (s / maxes[i]) * 10 : 0)

  chartInstance = new Chart(chartCanvas.value, {
    type: 'radar',
    data: {
      labels,
      datasets: [{
        data: normalized,
        backgroundColor: 'rgba(59,130,246,0.15)',
        borderColor: '#3b82f6',
        borderWidth: 2,
        pointBackgroundColor: axesWithScores.value.map((ax, i) => COLORS[i % COLORS.length]),
        pointBorderColor: '#fff',
        pointBorderWidth: 2,
        pointRadius: 5,
        pointHoverRadius: 7,
      }]
    },
    options: {
      responsive: false,
      animation: { duration: 400 },
      plugins: { legend: { display: false }, tooltip: {
        callbacks: {
          label: (ctx) => {
            const ax = axesWithScores.value[ctx.dataIndex]
            return ` ${ax.score} / ${ax.max}`
          }
        }
      }},
      scales: {
        r: {
          min: 0, max: 10,
          ticks: { display: false, stepSize: 2 },
          grid:        { color: gridColor },
          angleLines:  { color: gridColor },
          pointLabels: { color: labelColor, font: { size: 11, family: 'system-ui' } }
        }
      }
    }
  })
}

// ── Lifecycle ─────────────────────────────────────────────
onMounted(() => {
  // Charger Chart.js si pas encore disponible
  if (typeof Chart === 'undefined') {
    const script = document.createElement('script')
    script.src = 'https://cdn.jsdelivr.net/npm/chart.js@4.4.0/dist/chart.umd.min.js'
    script.onload = () => buildChart()
    document.head.appendChild(script)
  } else {
    buildChart()
  }
})

onUnmounted(() => {
  if (chartInstance) chartInstance.destroy()
})

watch([axesWithScores], () => {
  if (chartInstance) {
    const scores = axesWithScores.value.map(ax => ax.max > 0 ? (ax.score / ax.max) * 10 : 0)
    chartInstance.data.datasets[0].data = scores
    chartInstance.update('active')
  }
}, { deep: true })

// ── Classes alertes ───────────────────────────────────────
const alertClass = (level) => ({
  critique:  'border-red-200 bg-red-50 text-red-700 dark:border-red-800/40 dark:bg-red-950/20 dark:text-red-300',
  attention: 'border-orange-200 bg-orange-50 text-orange-700 dark:border-orange-800/40 dark:bg-orange-950/20 dark:text-orange-300',
  info:      'border-blue-200 bg-blue-50 text-blue-700 dark:border-blue-800/40 dark:bg-blue-950/20 dark:text-blue-300',
})[level] || 'border-gray-200 bg-gray-50 text-gray-600'

const alertIcon = (level) => ({ critique: '🚨', attention: '⚠️', info: 'ℹ️' })[level] || 'ℹ️'
</script>