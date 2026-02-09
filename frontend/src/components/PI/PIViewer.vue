<script setup>
import { ref, computed, onMounted } from 'vue'
// On importe les bindings générés par Wails (ajuste le chemin si nécessaire)
import { GetPlansByClient, CreatePlan, UpdatePlan } from '../../../wailsjs/go/main/App'

import PIHeader from './PiHeader.vue' // Attention à la casse 'PiHeader'
import PISidebar from './PIsidebar.vue' // Attention à la casse 'PIsidebar'
import PIFormWizard from './PIFormWizard.vue'
import PIDetailView from './PIDetailView.vue'
import PIFooter from './PIFooter.vue'
import PIQuickActions from './Piquickactions.vue'

const props = defineProps({
  client: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['close'])

// --- ÉTAT ---
const currentView = ref('list') 
const selectedPI = ref(null)
const plans = ref([])
const activeSection = ref('identification')
const isLoading = ref(false)
const isSaving = ref(false)

// --- LOGIQUE BACKEND ---

const loadPlans = async () => {
  if (!props.client?.id) return
  isLoading.value = true
  try {
    // APPEL DIRECT AU GO
    const result = await GetPlansByClientId(props.client.id)
    plans.value = result || []
  } catch (err) {
    console.error("Erreur lors du chargement des PI:", err)
  } finally {
    isLoading.value = false
  }
}

const savePI = async (data) => {
  isSaving.value = true
  try {
    if (data.id) {
      await UpdatePlan(data)
    } else {
      await CreatePlan(data)
    }
    await loadPlans() // Recharger la liste
    currentView.value = 'list'
  } catch (err) {
    console.error("Erreur sauvegarde PI:", err)
  } finally {
    isSaving.value = false
  }
}

// --- ACTIONS UI ---
const startNewPI = () => {
  selectedPI.value = {
    client_id: props.client.id,
    titre: '',
    statut: 'brouillon',
    objectifs: [],
    interventions: []
  }
  currentView.value = 'create'
}

// Dans PIViewer.vue (ou ton bouton de sauvegarde)
const handleSave = async () => {
  if (selectedPI.value.id) {
    // On passe l'ID séparément, puis le reste des données
    await UpdatePlan(selectedPI.value.id, formData.value)
  } else {
    await CreatePlan(formData.value)
  }
}

const viewPI = (pi) => {
  selectedPI.value = pi
  currentView.value = 'view'
}

onMounted(() => {
  loadPlans()
})

// Filtrage pour la sidebar
const plansActifs = computed(() => plans.value.filter(p => p.statut === 'actif'))
const plansEnCours = computed(() => plans.value.filter(p => p.statut === 'brouillon'))
const plansArchives = computed(() => plans.value.filter(p => p.statut === 'archive'))

</script>

<template>
  <div class="flex h-screen bg-slate-50 dark:bg-slate-900 overflow-hidden">
    <PISidebar 
      :plans="plans"
      :plans-actifs="plansActifs"
      :plans-en-cours="plansEnCours"
      :plans-archives="plansArchives"
      :selected-pi="selectedPI"
      :current-view="currentView"
      @view-pi="viewPI"
      @section-change="(s) => activeSection = s"
    />

    <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <PIHeader 
        :client="props.client"
        :selected-pi="selectedPI"
        :current-view="currentView"
        @start-new="startNewPI"
        @close="emit('close')"
      />

      <main class="flex-1 overflow-y-auto p-6">
        <div v-if="isLoading" class="flex items-center justify-center h-full">
           <p>Chargement des données sécurisées...</p>
        </div>

        <template v-else>
          <PIFormWizard 
            v-if="currentView === 'create' || currentView === 'edit'"
            v-model="selectedPI"
            :active-section="activeSection"
            @save="savePI(selectedPI)"
            @cancel="currentView = 'list'"
          />

          <PIDetailView 
            v-else-if="currentView === 'view' && selectedPI"
            :plan="selectedPI"
            :client="props.client"
            @edit="currentView = 'edit'"
          />

          <div v-else class="h-full flex flex-col items-center justify-center text-slate-400">
             <p>Sélectionnez un Plan d'Intervention dans la barre latérale.</p>
          </div>
        </template>
      </main>

      <PIFooter 
        v-if="currentView !== 'list' && selectedPI"
        :current-pi="selectedPI"
        :is-saving="isSaving"
        :current-view="currentView"
      />
    </div>
  </div>
</template>

<style scoped>
/* Animations douces */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}

.slide-enter-from {
  transform: translateX(-100%);
}

.slide-leave-to {
  transform: translateX(100%);
}
</style>