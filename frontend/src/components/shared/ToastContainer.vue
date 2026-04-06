<template>
  <Teleport to="body">
    <div class="fixed bottom-5 right-5 z-[9999] flex flex-col gap-2 items-end">
      <TransitionGroup name="toast">
        <div
          v-for="t in toasts"
          :key="t.id"
          :class="[
            'flex items-start gap-3 px-4 py-3 rounded-xl shadow-xl border',
            'min-w-[280px] max-w-[380px] cursor-pointer select-none',
            styles[t.type].wrapper
          ]"
          @click="remove(t.id)"
        >
          <!-- Icône -->
          <div :class="['mt-0.5 shrink-0', styles[t.type].icon]">
            <component :is="icons[t.type]" :size="18" />
          </div>

          <!-- Texte -->
          <div class="flex-1 min-w-0">
            <p :class="['text-sm font-semibold leading-tight', styles[t.type].title]">
              {{ t.message }}
            </p>
            <p v-if="t.detail" :class="['text-xs mt-0.5 leading-snug', styles[t.type].detail]">
              {{ t.detail }}
            </p>
          </div>

          <!-- Bouton fermer (toujours visible sur error) -->
          <button
            :class="['shrink-0 opacity-50 hover:opacity-100 transition-opacity', styles[t.type].icon]"
            @click.stop="remove(t.id)"
          >
            <X :size="14" />
          </button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup>
import { CheckCircle2, AlertTriangle, XCircle, Info, X } from 'lucide-vue-next'
import { useToast } from '@/composables/useToast.js'

const { toasts, remove } = useToast()

const icons = {
  success: CheckCircle2,
  error:   XCircle,
  warning: AlertTriangle,
  info:    Info
}

const styles = {
  success: {
    wrapper: 'bg-emerald-950/95 border-emerald-700/60 backdrop-blur-sm',
    icon:    'text-emerald-400',
    title:   'text-emerald-100',
    detail:  'text-emerald-300/80'
  },
  error: {
    wrapper: 'bg-red-950/95 border-red-700/60 backdrop-blur-sm',
    icon:    'text-red-400',
    title:   'text-red-100',
    detail:  'text-red-300/80'
  },
  warning: {
    wrapper: 'bg-amber-950/95 border-amber-700/60 backdrop-blur-sm',
    icon:    'text-amber-400',
    title:   'text-amber-100',
    detail:  'text-amber-300/80'
  },
  info: {
    wrapper: 'bg-sky-950/95 border-sky-700/60 backdrop-blur-sm',
    icon:    'text-sky-400',
    title:   'text-sky-100',
    detail:  'text-sky-300/80'
  }
}
</script>

<style scoped>
.toast-enter-active {
  transition: all 0.25s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.toast-leave-active {
  transition: all 0.2s ease-in;
}
.toast-enter-from {
  opacity: 0;
  transform: translateX(40px) scale(0.95);
}
.toast-leave-to {
  opacity: 0;
  transform: translateX(40px) scale(0.95);
}
</style>