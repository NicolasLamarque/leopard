// composables/useConfirm.js
import { reactive } from 'vue'

const state = reactive({
  isOpen: false,
  title: '',
  message: '',
  detail: null,
  level: 'warning',   // 'warning' | 'danger'
  confirmLabel: 'Confirmer',
  cancelLabel: 'Annuler',
  resolve: null
})

export function useConfirm() {

  const confirm = ({ title, message, detail = null, level = 'warning', confirmLabel = 'Confirmer', cancelLabel = 'Annuler' }) => {
    return new Promise((resolve) => {
      state.isOpen = true
      state.title = title
      state.message = message
      state.detail = detail
      state.level = level
      state.confirmLabel = confirmLabel
      state.cancelLabel = cancelLabel
      state.resolve = resolve
    })
  }

  const onConfirm = () => {
    state.isOpen = false
    state.resolve?.(true)
  }

  const onCancel = () => {
    state.isOpen = false
    state.resolve?.(false)
  }

  return { confirm, confirmState: state, onConfirm, onCancel }
}