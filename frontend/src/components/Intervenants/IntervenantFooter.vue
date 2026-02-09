<template>
  <footer class="relative h-12 bg-white/90 dark:bg-slate-900/90 backdrop-blur-xl border-t border-slate-200/50 dark:border-slate-800/50 flex items-center justify-between px-8 transition-colors duration-300 shadow-sm">
    <!-- Accent décoratif -->
    <div class="absolute inset-x-0 top-0 h-[1px] bg-gradient-to-r from-transparent via-teal-500/30 to-transparent"></div>
    
    <div class="flex items-center gap-3">
      <!-- Indicateur visuel de conformité avec animation -->
      <div class="relative flex items-center gap-2">
        <div class="relative">
          <div class="absolute inset-0 bg-teal-500/30 rounded-full blur-md animate-pulse"></div>
          <div class="relative w-2 h-2 rounded-full bg-gradient-to-br from-teal-400 to-emerald-500 shadow-lg shadow-teal-500/50"></div>
        </div>
        <span class="text-[10px] font-black text-slate-500 dark:text-slate-400 uppercase tracking-[0.15em]">
          Sécurité Loi 25 Active
        </span>
      </div>

      <!-- Séparateur -->
      <div class="w-px h-4 bg-slate-200 dark:bg-slate-700"></div>

      <!-- Badge de cryptage -->
      <div class="flex items-center gap-1.5 px-2.5 py-1 bg-gradient-to-r from-amber-500/10 to-orange-500/10 rounded-lg border border-amber-200/50 dark:border-amber-800/50">
        <ShieldCheck :size="12" class="text-amber-600 dark:text-amber-400" />
        <span class="text-[9px] font-black text-amber-700 dark:text-amber-400 uppercase tracking-wide">
          AES-256
        </span>
      </div>
    </div>

    <div class="flex items-center gap-6">
      <!-- Stats optionnelles (peuvent être alimentées par props) -->
      <div class="flex items-center gap-4">
        <div class="flex items-center gap-2">
          <div class="w-1.5 h-1.5 rounded-full bg-teal-500"></div>
          <span class="text-[10px] text-slate-500 dark:text-slate-400 font-bold uppercase tracking-wide">
            En ligne
          </span>
        </div>
      </div>

      <!-- Séparateur -->
      <div class="w-px h-4 bg-slate-200 dark:bg-slate-700"></div>

      <!-- Version -->
      <div class="flex items-center gap-2">
        <span class="text-[10px] text-slate-400 dark:text-slate-500 font-black uppercase tracking-[0.15em]">
          Leopard v1.0
        </span>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { ShieldCheck } from 'lucide-vue-next'
// Composant pur : statique pour le moment, mais prêt pour recevoir des props de statut
</script>

<style scoped>
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>