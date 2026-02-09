<template>
  <header class="relative flex items-center justify-between px-8 py-5 bg-white/90 dark:bg-slate-900/90 backdrop-blur-xl border-b border-slate-200/50 dark:border-slate-800/50 transition-colors shadow-sm">
    <!-- Accent décoratif -->
    <div class="absolute inset-x-0 top-0 h-[2px] bg-gradient-to-r from-transparent via-teal-500/50 to-transparent"></div>
    
    <div class="flex items-center gap-4">
      <div class="relative group">
        <!-- Glow effect -->
        <div class="absolute inset-0 bg-gradient-to-br from-teal-500/20 to-emerald-500/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-500"></div>
        <div class="relative p-3 bg-gradient-to-br from-teal-500/10 to-emerald-500/10 dark:from-teal-500/20 dark:to-emerald-500/20 rounded-2xl backdrop-blur-sm border border-teal-500/20 group-hover:border-teal-500/40 transition-all duration-300">
          <UserCog :size="26" class="text-teal-600 dark:text-teal-400" />
        </div>
      </div>
      <div>
        <h1 class="text-2xl font-black tracking-tight text-slate-900 dark:text-white uppercase leading-none">
          Intervenants & Réseau
        </h1>
        <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5 font-medium">Gestion sécurisée des contacts professionnels</p>
      </div>
    </div>

    <div class="flex items-center gap-3">
      <!-- Bouton Import -->
      <button 
        @click="$emit('import')"
        class="group flex items-center gap-2.5 px-5 py-2.5 bg-white dark:bg-slate-800/80 backdrop-blur-sm border border-slate-200 dark:border-slate-700 hover:border-teal-500/50 dark:hover:border-teal-500/50 text-slate-700 dark:text-slate-200 rounded-xl text-sm font-bold transition-all shadow-sm hover:shadow-md active:scale-95 cursor-pointer"
      >
        <div class="p-1 bg-teal-500/10 rounded-lg group-hover:bg-teal-500/20 transition-colors">
          <FileUp :size="18" class="text-teal-600 dark:text-teal-400" />
        </div>
        <span>Importer Excel</span>
      </button>

      <!-- Bouton Nouveau -->
      <button 
        @click="$emit('create')"
        class="relative group flex items-center gap-2.5 px-6 py-2.5 bg-gradient-to-r from-teal-600 to-emerald-600 hover:from-teal-500 hover:to-emerald-500 text-white rounded-xl text-sm font-bold transition-all shadow-lg shadow-teal-500/25 hover:shadow-xl hover:shadow-teal-500/40 active:scale-95 cursor-pointer overflow-hidden"
      >
        <div class="absolute inset-0 bg-gradient-to-r from-white/0 via-white/20 to-white/0 -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
        <Plus :size="18" class="relative" />
        <span class="relative">Nouveau</span>
      </button>
    </div>
  </header>
</template>

<script setup>
import { UserCog, FileUp, Plus } from 'lucide-vue-next'
defineEmits(['import', 'create'])
</script>