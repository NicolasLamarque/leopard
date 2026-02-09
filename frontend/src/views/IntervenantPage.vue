<template>
  <div class="flex flex-col h-full bg-gradient-to-br from-slate-50 via-white to-slate-50 dark:from-gray-950 dark:via-gray-900 dark:to-gray-950 transition-colors duration-500">
    
    <!-- HEADER : Actions globales -->
    <PageHeader 
      @import="handleImportExcel" 
      @create="initNewIntervenant" 
    />

    <main class="flex-1 overflow-hidden p-6">
      <div class="h-full flex gap-6">
        
        <!-- LISTE : Navigation Master -->
        <IntervenantList 
          :items="filteredIntervenants" 
          :selected-id="selectedId"
          :is-loading="isLoading"
          v-model:search="searchQuery"
          v-model:filter="filterType"
          @select="selectedId = $event"
        />

        <!-- VIEWER CENTRAL : Conteneur dynamique -->
        <section class="flex-1 flex flex-col bg-white/80 dark:bg-slate-900/80 backdrop-blur-xl rounded-3xl border border-slate-200/50 dark:border-slate-800/50 shadow-2xl shadow-slate-200/50 dark:shadow-black/20 overflow-hidden text-slate-900 dark:text-white transition-all">
          
          <!-- En-tête du Viewer -->
          <div class="flex border-b border-slate-100/50 dark:border-slate-800/50 px-8 bg-gradient-to-r from-slate-50/50 to-transparent dark:from-slate-900/50">
            <div class="px-6 py-5 text-xs font-black uppercase tracking-[0.2em] border-b-2 border-teal-500 text-teal-600 dark:text-teal-400 flex items-center gap-3">
              <div class="w-1.5 h-1.5 rounded-full bg-teal-500 animate-pulse"></div>
              Fiche d'identification
            </div>
          </div>

          <!-- Zone de contenu -->
          <div class="flex-1 p-10 overflow-y-auto custom-scrollbar">
             <!-- Affichage conditionnel du formulaire -->
             <div v-if="selectedId" class="max-w-5xl mx-auto animate-fade-in">
                <IntervenantForm 
                  :intervenant-id="selectedId" 
                  @saved="handleSaveSuccess"
                  @cancel="selectedId = null"
                />
             </div>

             <!-- État vide si aucune sélection -->
             <div v-else class="h-full flex flex-col items-center justify-center text-center space-y-6 animate-fade-in">
                <div class="relative">
                  <div class="absolute inset-0 bg-gradient-to-r from-teal-500/20 to-emerald-500/20 blur-3xl rounded-full"></div>
                  <div class="relative p-8 bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-800/50 dark:to-slate-900/50 rounded-3xl border border-slate-200/50 dark:border-slate-700/50 shadow-xl">
                    <UserCog :size="72" class="text-slate-300 dark:text-slate-600" />
                  </div>
                </div>
                <div class="space-y-2">
                  <h3 class="text-xl font-black text-slate-400 dark:text-slate-500 uppercase tracking-wide">Aucune sélection</h3>
                  <p class="text-sm text-slate-500 dark:text-slate-600 max-w-md">Sélectionnez un intervenant dans la liste ou créez-en un nouveau pour commencer.</p>
                </div>
             </div>
          </div>
        </section>

      </div>
    </main>

    <!-- FOOTER : Statut et Conformité -->
    <PageFooter />
    
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { UserCog } from 'lucide-vue-next'
import { GetAllIntervenants } from '../../wailsjs/go/main/App'

import PageHeader from '../components/Intervenants/IntervenantHeader.vue'
import IntervenantList from '../components/Intervenants/IntervenantList.vue'
import PageFooter from '../components/Intervenants/IntervenantFooter.vue'
import IntervenantForm from '../components/Intervenants/IntervenantForm.vue'

const searchQuery = ref('')
const filterType = ref('Tous')
const selectedId = ref(null) 
const intervenants = ref([])
const isLoading = ref(false)

onMounted(async () => {
  await loadIntervenants()
})

const loadIntervenants = async () => {
  try {
    isLoading.value = true
    const data = await GetAllIntervenants()
    intervenants.value = data || []
  } catch (error) {
    console.error("Erreur chargement intervenants:", error)
    intervenants.value = []
  } finally {
    isLoading.value = false
  }
}

const filteredIntervenants = computed(() => {
  const query = searchQuery.value.toLowerCase()
  return intervenants.value.filter(i => {
    const matchesSearch = i.nom_complet?.toLowerCase().includes(query)
    const matchesFilter = filterType.value === 'Tous' || i.titre_emploi === filterType.value
    return matchesSearch && matchesFilter
  })
})

const handleImportExcel = async () => { 
  console.log("Import Excel - À implémenter") 
}

const initNewIntervenant = () => { 
  selectedId.value = 'new' 
}

const handleSaveSuccess = async () => {
  await loadIntervenants()
  selectedId.value = null
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #cbd5e1, #94a3b8);
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #94a3b8, #64748b);
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #334155, #1e293b);
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #475569, #334155);
}

@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.5s ease-out;
}
</style>