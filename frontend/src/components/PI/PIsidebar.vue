<script setup>
import { ref, computed } from 'vue'
import { 
  Lock, FileText, Edit3, Copy, Archive, 
  ChevronDown, ChevronRight, MoreVertical,
  Calendar, User, AlertCircle, CheckCircle2
} from 'lucide-vue-next'

const props = defineProps({
  plans: {
    type: Array,
    default: () => []
  },
  plansEnCours: {
    type: Array,
    default: () => []
  },
  plansActifs: {
    type: Array,
    default: () => []
  },
  plansScelles: {
    type: Array,
    default: () => []
  },
  plansArchives: {
    type: Array,
    default: () => []
  },
  selectedPi: Object,
  currentView: {
    type: String,
    default: 'list'
  },
  activeSection: String
})

const emit = defineEmits([
  'view-pi', 'edit-pi', 'duplicate-pi', 
  'archive-pi', 'section-change'
])

// Sections du formulaire wizard
const formSections = [
  {
    id: 'identification',
    label: 'Identification',
    icon: 'FileText',
    description: 'Titre et dates'
  },
  {
    id: 'analyse',
    label: 'Analyse',
    icon: 'AlertCircle',
    description: 'Contexte, problématique, forces'
  },
  {
    id: 'objectifs',
    label: 'Objectifs',
    icon: 'CheckCircle2',
    description: 'Objectifs SMART'
  },
  {
    id: 'interventions',
    label: 'Interventions',
    icon: 'Calendar',
    description: 'Plan d\'action'
  },
  {
    id: 'ententes',
    label: 'Ententes',
    icon: 'User',
    description: 'Engagements & suivi'
  }
]

// État d'expansion des catégories
const expandedCategories = ref({
  enCours: true,
  actifs: true,
  scelles: false,
  archives: false
})

const toggleCategory = (category) => {
  expandedCategories.value[category] = !expandedCategories.value[category]
}

// Menu contextuel
const showContextMenu = ref(null)

const openContextMenu = (planId, event) => {
  event.stopPropagation()
  showContextMenu.value = showContextMenu.value === planId ? null : planId
}

const closeContextMenu = () => {
  showContextMenu.value = null
}

// Format date
const formatDate = (dateStr) => {
  if (!dateStr) return 'Non défini'
  return new Date(dateStr).toLocaleDateString('fr-CA', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

// Calcul du temps écoulé
const getTimeAgo = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  const now = new Date()
  const days = Math.floor((now - date) / (1000 * 60 * 60 * 24))
  
  if (days === 0) return 'Aujourd\'hui'
  if (days === 1) return 'Hier'
  if (days < 7) return `Il y a ${days}j`
  if (days < 30) return `Il y a ${Math.floor(days / 7)}sem`
  return `Il y a ${Math.floor(days / 30)}mois`
}
</script>

<template>
  <aside class="w-80 border-r border-slate-200 dark:border-slate-800 
                bg-white dark:bg-slate-900/50 flex flex-col h-full overflow-hidden">
    
    <!-- Wizard Steps (Mode création/édition) -->
    <div v-if="currentView === 'create' || currentView === 'edit'" 
         class="flex-1 flex flex-col overflow-hidden">
      
      <div class="px-4 py-3 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900/50">
        <span class="text-[10px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest">
          Étapes du formulaire
        </span>
      </div>

      <div class="flex-1 overflow-y-auto p-4 space-y-2">
        <button
          v-for="(section, index) in formSections" 
          :key="section.id"
          @click="$emit('section-change', section.id)"
          :class="[
            'w-full flex items-start gap-3 px-4 py-3 rounded-xl transition-all text-left group',
            activeSection === section.id
              ? 'bg-rose-50 dark:bg-rose-900/20 text-rose-700 dark:text-rose-400 border-2 border-rose-200 dark:border-rose-800 shadow-sm'
              : 'hover:bg-slate-50 dark:hover:bg-slate-800/50 text-slate-700 dark:text-slate-300 border border-transparent'
          ]"
        >
          <div :class="[
            'flex-shrink-0 w-8 h-8 rounded-lg flex items-center justify-center text-xs font-bold',
            activeSection === section.id 
              ? 'bg-rose-500 text-white' 
              : 'bg-slate-100 dark:bg-slate-800 text-slate-500 dark:text-slate-400 group-hover:bg-slate-200 dark:group-hover:bg-slate-700'
          ]">
            {{ index + 1 }}
          </div>
          
          <div class="flex-1 min-w-0">
            <div class="font-semibold text-sm">{{ section.label }}</div>
            <div class="text-xs opacity-60 mt-0.5">{{ section.description }}</div>
          </div>
        </button>
      </div>
    </div>

    <!-- Liste des Plans (Mode consultation) -->
    <div v-else class="flex-1 flex flex-col overflow-hidden">
      
      <div class="px-4 py-3 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900/50">
        <span class="text-[10px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest">
          Plans d'intervention
        </span>
      </div>

      <div class="flex-1 overflow-y-auto">
        
        <!-- Plans en cours (Brouillons) -->
        <div v-if="plansEnCours.length > 0" class="border-b border-slate-200 dark:border-slate-800">
          <button 
            @click="toggleCategory('enCours')"
            class="w-full px-4 py-3 flex items-center justify-between hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors"
          >
            <div class="flex items-center gap-2">
              <component 
                :is="expandedCategories.enCours ? ChevronDown : ChevronRight" 
                :size="16" 
                class="text-slate-400"
              />
              <span class="text-xs font-bold text-amber-600 dark:text-amber-400 uppercase tracking-wide">
                En cours
              </span>
              <span class="text-xs font-bold text-amber-600/60 dark:text-amber-400/60">
                ({{ plansEnCours.length }})
              </span>
            </div>
          </button>
          
          <div v-if="expandedCategories.enCours" class="px-3 pb-3 space-y-2">
            <div
              v-for="pi in plansEnCours"
              :key="pi.id"
              @click="$emit('view-pi', pi)"
              :class="[
                'relative p-3 rounded-lg border cursor-pointer transition-all group',
                selectedPi?.id === pi.id
                  ? 'bg-amber-50 dark:bg-amber-900/10 border-amber-300 dark:border-amber-700 ring-2 ring-amber-400/50'
                  : 'bg-white dark:bg-slate-800/50 border-slate-200 dark:border-slate-700 hover:border-amber-300 dark:hover:border-amber-700 hover:shadow-sm'
              ]"
            >
              <div class="flex items-start justify-between gap-2 mb-2">
                <div class="flex items-center gap-2 min-w-0">
                  <div class="w-1.5 h-1.5 rounded-full bg-amber-500 flex-shrink-0"></div>
                  <span class="text-[10px] font-bold text-amber-600 dark:text-amber-400 uppercase">Brouillon</span>
                </div>
                <button
                  @click.stop="openContextMenu(pi.id, $event)"
                  class="opacity-0 group-hover:opacity-100 p-1 hover:bg-slate-100 dark:hover:bg-slate-700 rounded transition-opacity"
                >
                  <MoreVertical :size="14" class="text-slate-400" />
                </button>
              </div>
              
              <h3 class="font-semibold text-sm text-slate-900 dark:text-slate-100 truncate mb-1">
                {{ pi.titre || 'Plan sans titre' }}
              </h3>
              
              <div class="flex items-center justify-between text-[10px] text-slate-500 dark:text-slate-400">
                <span>{{ getTimeAgo(pi.created_at) }}</span>
                <span>{{ formatDate(pi.date_debut) }}</span>
              </div>

              <!-- Context Menu -->
              <div
                v-if="showContextMenu === pi.id"
                @click.stop
                class="absolute right-2 top-12 bg-white dark:bg-slate-800 rounded-lg shadow-lg border border-slate-200 dark:border-slate-700 py-1 z-10 min-w-[160px]"
              >
                <button
                  @click="$emit('edit-pi', pi); closeContextMenu()"
                  class="w-full px-3 py-2 text-left text-sm hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2"
                >
                  <Edit3 :size="14" />
                  <span>Modifier</span>
                </button>
                <button
                  @click="$emit('duplicate-pi', pi); closeContextMenu()"
                  class="w-full px-3 py-2 text-left text-sm hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2"
                >
                  <Copy :size="14" />
                  <span>Dupliquer</span>
                </button>
                <button
                  @click="$emit('archive-pi', pi.id); closeContextMenu()"
                  class="w-full px-3 py-2 text-left text-sm hover:bg-slate-50 dark:hover:bg-slate-700 flex items-center gap-2 text-red-600"
                >
                  <Archive :size="14" />
                  <span>Archiver</span>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Plans actifs -->
        <div v-if="plansActifs.length > 0" class="border-b border-slate-200 dark:border-slate-800">
          <button 
            @click="toggleCategory('actifs')"
            class="w-full px-4 py-3 flex items-center justify-between hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors"
          >
            <div class="flex items-center gap-2">
              <component 
                :is="expandedCategories.actifs ? ChevronDown : ChevronRight" 
                :size="16" 
                class="text-slate-400"
              />
              <span class="text-xs font-bold text-teal-600 dark:text-teal-400 uppercase tracking-wide">
                Actifs
              </span>
              <span class="text-xs font-bold text-teal-600/60 dark:text-teal-400/60">
                ({{ plansActifs.length }})
              </span>
            </div>
          </button>
          
          <div v-if="expandedCategories.actifs" class="px-3 pb-3 space-y-2">
            <div
              v-for="pi in plansActifs"
              :key="pi.id"
              @click="$emit('view-pi', pi)"
              :class="[
                'p-3 rounded-lg border cursor-pointer transition-all',
                selectedPi?.id === pi.id
                  ? 'bg-teal-50 dark:bg-teal-900/10 border-teal-300 dark:border-teal-700 ring-2 ring-teal-400/50'
                  : 'bg-white dark:bg-slate-800/50 border-slate-200 dark:border-slate-700 hover:border-teal-300 dark:hover:border-teal-700 hover:shadow-sm'
              ]"
            >
              <div class="flex items-start justify-between gap-2 mb-2">
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 rounded-full bg-teal-500"></div>
                  <span class="text-[10px] font-bold text-teal-600 dark:text-teal-400 uppercase">Actif</span>
                </div>
              </div>
              
              <h3 class="font-semibold text-sm text-slate-900 dark:text-slate-100 truncate mb-1">
                {{ pi.titre }}
              </h3>
              
              <div class="flex items-center justify-between text-[10px] text-slate-500 dark:text-slate-400">
                <span>{{ getTimeAgo(pi.date_debut) }}</span>
                <span v-if="pi.date_fin_prevue">
                  Fin: {{ formatDate(pi.date_fin_prevue) }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Plans scellés -->
        <div v-if="plansScelles.length > 0" class="border-b border-slate-200 dark:border-slate-800">
          <button 
            @click="toggleCategory('scelles')"
            class="w-full px-4 py-3 flex items-center justify-between hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors"
          >
            <div class="flex items-center gap-2">
              <component 
                :is="expandedCategories.scelles ? ChevronDown : ChevronRight" 
                :size="16" 
                class="text-slate-400"
              />
              <span class="text-xs font-bold text-emerald-600 dark:text-emerald-400 uppercase tracking-wide">
                Scellés
              </span>
              <span class="text-xs font-bold text-emerald-600/60 dark:text-emerald-400/60">
                ({{ plansScelles.length }})
              </span>
            </div>
          </button>
          
          <div v-if="expandedCategories.scelles" class="px-3 pb-3 space-y-2">
            <div
              v-for="pi in plansScelles"
              :key="pi.id"
              @click="$emit('view-pi', pi)"
              :class="[
                'p-3 rounded-lg border cursor-pointer transition-all',
                selectedPi?.id === pi.id
                  ? 'bg-emerald-50 dark:bg-emerald-900/10 border-emerald-300 dark:border-emerald-700 ring-2 ring-emerald-400/50'
                  : 'bg-white dark:bg-slate-800/50 border-slate-200 dark:border-slate-700 hover:border-emerald-300 dark:hover:border-emerald-700 hover:shadow-sm'
              ]"
            >
              <div class="flex items-start justify-between gap-2 mb-2">
                <div class="flex items-center gap-2">
                  <Lock :size="12" class="text-emerald-500" />
                  <span class="text-[10px] font-bold text-emerald-600 dark:text-emerald-400 uppercase">Scellé</span>
                </div>
              </div>
              
              <h3 class="font-semibold text-sm text-slate-900 dark:text-slate-100 truncate mb-1">
                {{ pi.titre }}
              </h3>
              
              <div class="text-[10px] text-slate-500 dark:text-slate-400">
                Signé: {{ formatDate(pi.date_signature) }}
              </div>
            </div>
          </div>
        </div>

        <!-- Plans archivés -->
        <div v-if="plansArchives.length > 0">
          <button 
            @click="toggleCategory('archives')"
            class="w-full px-4 py-3 flex items-center justify-between hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors"
          >
            <div class="flex items-center gap-2">
              <component 
                :is="expandedCategories.archives ? ChevronDown : ChevronRight" 
                :size="16" 
                class="text-slate-400"
              />
              <span class="text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wide">
                Archivés
              </span>
              <span class="text-xs font-bold text-slate-500/60 dark:text-slate-400/60">
                ({{ plansArchives.length }})
              </span>
            </div>
          </button>
          
          <div v-if="expandedCategories.archives" class="px-3 pb-3 space-y-2">
            <div
              v-for="pi in plansArchives"
              :key="pi.id"
              @click="$emit('view-pi', pi)"
              class="p-3 rounded-lg border bg-white dark:bg-slate-800/50 border-slate-200 dark:border-slate-700 cursor-pointer transition-all hover:shadow-sm opacity-75 hover:opacity-100"
            >
              <h3 class="font-semibold text-sm text-slate-700 dark:text-slate-300 truncate mb-1">
                {{ pi.titre }}
              </h3>
              <div class="text-[10px] text-slate-400">
                Archivé: {{ formatDate(pi.updated_at) }}
              </div>
            </div>
          </div>
        </div>

        <!-- Empty state -->
        <div v-if="plans.length === 0" class="p-6 text-center">
          <div class="w-16 h-16 mx-auto mb-3 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center">
            <FileText :size="28" class="text-slate-400" />
          </div>
          <p class="text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">
            Aucun plan
          </p>
          <p class="text-xs text-slate-500 dark:text-slate-500">
            Créez votre premier plan d'intervention
          </p>
        </div>
      </div>
    </div>

    <!-- Click outside to close context menu -->
    <div 
      v-if="showContextMenu" 
      @click="closeContextMenu"
      class="fixed inset-0 z-0"
    ></div>
  </aside>
</template>