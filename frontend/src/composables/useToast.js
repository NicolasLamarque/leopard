// composables/useToast.js
import { reactive } from 'vue'

// État global partagé entre tous les composants
const state = reactive({
  toasts: []
})

let nextId = 0

const DURATIONS = {
  success: 2000,
  error: 0,      // reste jusqu'à fermeture manuelle
  warning: 4000,
  info: 3000
}

export function useToast() {

  const add = (type, message, detail = null) => {
    const id = ++nextId
    const duration = DURATIONS[type]

    state.toasts.push({ id, type, message, detail })

    if (duration > 0) {
      setTimeout(() => remove(id), duration)
    }
  }

  const remove = (id) => {
    const idx = state.toasts.findIndex(t => t.id === id)
    if (idx !== -1) state.toasts.splice(idx, 1)
  }

  const toast = {
    // ✅ Succès  — message court + détail optionnel
    success: (message, detail = null) => add('success', message, detail),
    // ❌ Erreur  — reste affiché, détail obligatoire recommandé
    error:   (message, detail = null) => add('error',   message, detail),
    // ⚠️ Avertissement
    warning: (message, detail = null) => add('warning', message, detail),
    // ℹ️ Info
    info:    (message, detail = null) => add('info',    message, detail),
  }

  return { toast, toasts: state.toasts, remove }
}