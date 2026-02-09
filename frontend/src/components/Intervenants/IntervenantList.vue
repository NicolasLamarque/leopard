<template>
  <aside class="w-96 flex flex-col bg-white/80 dark:bg-slate-900/80 backdrop-blur-xl rounded-3xl border border-slate-200/50 dark:border-slate-800/50 shadow-2xl shadow-slate-200/50 dark:shadow-black/20 overflow-hidden transition-colors">
    <!-- Barre de recherche et Filtres -->
    <div class="p-5 border-b border-slate-100/50 dark:border-slate-800/50 space-y-4 bg-gradient-to-b from-slate-50/50 to-transparent dark:from-slate-900/50">
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-teal-500/10 to-emerald-500/10 rounded-2xl opacity-0 group-focus-within:opacity-100 transition-opacity duration-300"></div>
        <Search class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 group-focus-within:text-teal-500 transition-colors duration-300 z-10" :size="18" />
        <input 
          :value="search"
          @input="$emit('update:search', $event.target.value)"
          type="text" 
          placeholder="Rechercher un intervenant..." 
          class="relative w-full pl-12 pr-4 py-3.5 bg-slate-50/50 dark:bg-slate-800/50 backdrop-blur-sm border border-slate-200/50 dark:border-slate-700/50 rounded-2xl text-sm outline-none focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500/50 text-slate-900 dark:text-white placeholder:text-slate-400 transition-all duration-300"
        />
      </div>
      
      <div class="flex gap-2 flex-wrap">
        <button 
          v-for="t in filterOptions" :key="t"
          @click="$emit('update:filter', t)"
          :class="[
            'px-4 py-2 text-xs font-black uppercase tracking-wide rounded-xl border transition-all duration-300',
            filter === t 
              ? 'bg-gradient-to-r from-teal-500 to-emerald-500 border-transparent text-white shadow-lg shadow-teal-500/30' 
              : 'bg-white/50 dark:bg-slate-800/50 backdrop-blur-sm border-slate-200/50 dark:border-slate-700/50 text-slate-500 dark:text-slate-400 hover:border-teal-500/50 hover:text-teal-600 dark:hover:text-teal-400'
          ]"
        >
          {{ t }}
        </button>
      </div>
    </div>
    
    <!-- Liste Scrollable -->
    <div class="flex-1 overflow-y-auto scroll-area">
      <div 
        v-for="(item, index) in items" :key="item.id"
        @click="$emit('select', item.id)"
        :class="[
          'group relative p-5 border-b border-slate-100/30 dark:border-slate-800/30 cursor-pointer transition-all duration-300',
          selectedId === item.id 
            ? 'bg-gradient-to-r from-teal-50/80 to-emerald-50/80 dark:from-teal-900/20 dark:to-emerald-900/20 border-l-4 border-l-teal-500' 
            : 'hover:bg-slate-50/50 dark:hover:bg-slate-800/30'
        ]"
        :style="{ animationDelay: `${index * 50}ms` }"
      >
        <!-- Hover effect -->
        <div v-if="selectedId !== item.id" class="absolute inset-0 bg-gradient-to-r from-teal-500/0 via-teal-500/5 to-teal-500/0 opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
        
        <div class="relative flex justify-between items-start gap-3">
          <div class="flex-1 min-w-0">
            <h4 class="font-bold text-sm text-slate-900 dark:text-white mb-1 truncate group-hover:text-teal-600 dark:group-hover:text-teal-400 transition-colors">
              {{ item.nom_complet }}
            </h4>
            <p class="text-xs text-slate-500 dark:text-slate-400 font-medium truncate">
              {{ item.organisation || 'Indépendant' }}
            </p>
          </div>
          <span :class="[
            'shrink-0 text-[9px] px-2.5 py-1 rounded-lg font-black uppercase tracking-wide',
            getTypeColor(item.type)
          ]">
            {{ item.type }}
          </span>
        </div>

        <!-- Selection indicator -->
        <div v-if="selectedId === item.id" class="absolute left-0 top-1/2 -translate-y-1/2 w-1 h-8 bg-gradient-to-b from-teal-500 to-emerald-500 rounded-r-full"></div>
      </div>

      <!-- Empty state -->
      <div v-if="items.length === 0" class="p-8 text-center space-y-3">
        <div class="inline-flex p-4 bg-slate-100 dark:bg-slate-800 rounded-2xl">
          <Search :size="32" class="text-slate-400" />
        </div>
        <p class="text-sm text-slate-500 dark:text-slate-400 font-medium">Aucun intervenant trouvé</p>
      </div>
    </div>
  </aside>
</template>

<script setup>
import { Search } from 'lucide-vue-next'

defineProps(['items', 'selectedId', 'search', 'filter'])
defineEmits(['update:search', 'update:filter', 'select'])

const filterOptions = ['Tous', 'Médecin', 'Pharmacie', 'Travailleur Social', 'Infirmier(ère)']

const getTypeColor = (type) => {
  const colors = {
    'Médecin': 'bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400 border border-blue-200 dark:border-blue-800',
    'Pharmacie': 'bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-400 border border-purple-200 dark:border-purple-800',
    'Travailleur Social': 'bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-400 border border-emerald-200 dark:border-emerald-800',
    'Infirmier(ère)': 'bg-pink-100 dark:bg-pink-900/30 text-pink-700 dark:text-pink-400 border border-pink-200 dark:border-pink-800'
  }
  return colors[type] || 'bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 border border-slate-200 dark:border-slate-700'
}
</script>

<style scoped>
.scroll-area {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e1 transparent;
}

.scroll-area::-webkit-scrollbar {
  width: 6px;
}

.scroll-area::-webkit-scrollbar-track {
  background: transparent;
}

.scroll-area::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #cbd5e1, #94a3b8);
  border-radius: 10px;
}

.scroll-area::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #94a3b8, #64748b);
}

.dark .scroll-area::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #334155, #1e293b);
}

.dark .scroll-area::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #475569, #334155);
}
</style>