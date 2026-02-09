<script setup>
import { ref } from 'vue'
import { 
  Edit3, Copy, Archive, Trash2, 
  MoreVertical, X 
} from 'lucide-vue-next'

const props = defineProps({
  plan: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['edit', 'duplicate', 'archive', 'delete'])

const showMenu = ref(false)

const toggleMenu = () => {
  showMenu.value = !showMenu.value
}

const handleAction = (action) => {
  emit(action)
  showMenu.value = false
}
</script>

<template>
  <div class="fixed bottom-8 right-8 z-40">
    <!-- Menu déroulant -->
    <Transition name="slide-up">
      <div 
        v-if="showMenu"
        class="mb-4 bg-white dark:bg-slate-800 rounded-2xl shadow-2xl border border-slate-200 dark:border-slate-700 overflow-hidden"
      >
        <button
          @click="handleAction('edit')"
          class="w-full flex items-center gap-3 px-5 py-3 hover:bg-slate-50 dark:hover:bg-slate-700 transition-colors text-left"
        >
          <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg">
            <Edit3 :size="16" class="text-blue-600 dark:text-blue-400" />
          </div>
          <div>
            <p class="text-sm font-semibold text-slate-900 dark:text-slate-100">Modifier</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">Éditer le plan</p>
          </div>
        </button>

        <div class="h-px bg-slate-200 dark:bg-slate-700"></div>

        <button
          @click="handleAction('duplicate')"
          class="w-full flex items-center gap-3 px-5 py-3 hover:bg-slate-50 dark:hover:bg-slate-700 transition-colors text-left"
        >
          <div class="p-2 bg-purple-100 dark:bg-purple-900/30 rounded-lg">
            <Copy :size="16" class="text-purple-600 dark:text-purple-400" />
          </div>
          <div>
            <p class="text-sm font-semibold text-slate-900 dark:text-slate-100">Dupliquer</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">Créer une copie</p>
          </div>
        </button>

        <div class="h-px bg-slate-200 dark:bg-slate-700"></div>

        <button
          @click="handleAction('archive')"
          class="w-full flex items-center gap-3 px-5 py-3 hover:bg-slate-50 dark:hover:bg-slate-700 transition-colors text-left"
        >
          <div class="p-2 bg-amber-100 dark:bg-amber-900/30 rounded-lg">
            <Archive :size="16" class="text-amber-600 dark:text-amber-400" />
          </div>
          <div>
            <p class="text-sm font-semibold text-slate-900 dark:text-slate-100">Archiver</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">Déplacer vers archives</p>
          </div>
        </button>

        <div class="h-px bg-slate-200 dark:bg-slate-700"></div>

        <button
          @click="handleAction('delete')"
          class="w-full flex items-center gap-3 px-5 py-3 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors text-left"
        >
          <div class="p-2 bg-red-100 dark:bg-red-900/30 rounded-lg">
            <Trash2 :size="16" class="text-red-600 dark:text-red-400" />
          </div>
          <div>
            <p class="text-sm font-semibold text-red-600 dark:text-red-400">Supprimer</p>
            <p class="text-xs text-red-500/70 dark:text-red-400/70">Action irréversible</p>
          </div>
        </button>
      </div>
    </Transition>

    <!-- Bouton principal -->
    <button
      @click="toggleMenu"
      :class="[
        'p-4 rounded-full shadow-2xl transition-all transform hover:scale-110',
        showMenu 
          ? 'bg-slate-600 hover:bg-slate-700' 
          : 'bg-gradient-to-br from-rose-500 to-pink-500 hover:from-rose-600 hover:to-pink-600'
      ]"
    >
      <Transition name="rotate" mode="out-in">
        <X v-if="showMenu" :size="24" class="text-white" :stroke-width="2.5" />
        <MoreVertical v-else :size="24" class="text-white" :stroke-width="2.5" />
      </Transition>
    </button>

    <!-- Backdrop -->
    <Transition name="fade">
      <div 
        v-if="showMenu"
        @click="showMenu = false"
        class="fixed inset-0 -z-10"
      ></div>
    </Transition>
  </div>
</template>

<style scoped>
.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.slide-up-enter-from {
  opacity: 0;
  transform: translateY(20px) scale(0.9);
}

.slide-up-leave-to {
  opacity: 0;
  transform: translateY(10px) scale(0.95);
}

.rotate-enter-active,
.rotate-leave-active {
  transition: all 0.2s ease;
}

.rotate-enter-from {
  transform: rotate(-90deg);
  opacity: 0;
}

.rotate-leave-to {
  transform: rotate(90deg);
  opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>