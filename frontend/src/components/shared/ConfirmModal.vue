<template>
  <Teleport to="body">
    <Transition name="confirm-fade">
      <div
        v-if="confirmState.isOpen"
        class="fixed inset-0 z-[10000] flex items-center justify-center p-4"
      >
        <!-- Backdrop -->
        <div
          class="absolute inset-0 bg-gray-950/70 backdrop-blur-sm"
          @click="onCancel"
        />

        <!-- Modal -->
        <div class="relative bg-gray-900 border rounded-2xl shadow-2xl w-full max-w-md p-6 flex flex-col gap-4"
          :class="confirmState.level === 'danger' ? 'border-red-800/60' : 'border-amber-800/60'"
        >
          <!-- Icône + Titre -->
          <div class="flex items-center gap-3">
            <div class="p-2 rounded-xl shrink-0"
              :class="confirmState.level === 'danger' ? 'bg-red-950 text-red-400' : 'bg-amber-950 text-amber-400'"
            >
              <AlertTriangle v-if="confirmState.level === 'warning'" :size="22" />
              <ShieldAlert v-else :size="22" />
            </div>
            <h3 class="text-base font-bold text-white leading-tight">
              {{ confirmState.title }}
            </h3>
          </div>

          <!-- Message -->
          <p class="text-sm text-gray-300 leading-relaxed">
            {{ confirmState.message }}
          </p>

          <!-- Détail optionnel -->
          <p v-if="confirmState.detail"
            class="text-xs text-gray-500 border-t border-gray-800 pt-3 leading-relaxed"
          >
            {{ confirmState.detail }}
          </p>

          <!-- Actions -->
          <div class="flex items-center justify-end gap-3 pt-1">
            <button
              @click="onCancel"
              class="px-4 py-2 rounded-lg text-sm font-medium text-gray-400 hover:text-white hover:bg-gray-800 transition-all"
            >
              {{ confirmState.cancelLabel }}
            </button>
            <button
              @click="onConfirm"
              class="px-5 py-2 rounded-lg text-sm font-bold transition-all"
              :class="confirmState.level === 'danger'
                ? 'bg-red-600 hover:bg-red-500 text-white shadow-lg shadow-red-900/40'
                : 'bg-amber-600 hover:bg-amber-500 text-white shadow-lg shadow-amber-900/40'"
            >
              {{ confirmState.confirmLabel }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { AlertTriangle, ShieldAlert } from 'lucide-vue-next'
import { useConfirm } from '../../composables/useConfirm.js'

const { confirmState, onConfirm, onCancel } = useConfirm()
</script>

<style scoped>
.confirm-fade-enter-active,
.confirm-fade-leave-active {
  transition: opacity 0.2s ease;
}
.confirm-fade-enter-active .relative,
.confirm-fade-leave-active .relative {
  transition: transform 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.confirm-fade-enter-from,
.confirm-fade-leave-to {
  opacity: 0;
}
.confirm-fade-enter-from .relative {
  transform: scale(0.95) translateY(8px);
}
</style>