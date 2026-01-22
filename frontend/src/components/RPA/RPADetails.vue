<template>
  <div @click.self="$emit('close')" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4 backdrop-blur-sm">
    <div class="bg-white dark:bg-gray-800 rounded-xl shadow-2xl max-w-4xl w-full max-h-[92vh] overflow-hidden flex flex-col relative">
      
      <div v-if="isUpdating" class="absolute inset-0 bg-white/80 dark:bg-gray-800/80 z-[60] flex flex-col items-center justify-center">
        <RefreshCw :size="48" class="text-blue-600 animate-spin mb-4" />
        <p class="text-lg font-bold text-gray-900 dark:text-white">Synchronisation en direct...</p>
      </div>

      <div class="px-6 py-4 border-b dark:border-gray-700 bg-gray-50/50 dark:bg-gray-900/20 flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg text-blue-600">
            <Building2 :size="24" />
          </div>
          <div>
            <h2 class="text-xl font-bold text-gray-900 dark:text-white">{{ title }}</h2>
            <p v-if="mode !== 'create'" class="text-sm font-mono text-gray-500">#{{ form.registre }}</p>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <button v-if="mode === 'view'" @click="mode = 'edit'" class="px-4 py-2 bg-blue-600 text-white rounded-lg flex items-center gap-2 text-sm font-bold">
            <Edit :size="18" /> Modifier
          </button>
          <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600 p-2"><X :size="24" /></button>
        </div>
        <div class="flex items-center gap-2">
  
  <button 
    v-if="mode === 'view'"
    @click="handleManualSync"
    :disabled="isUpdating"
    class="flex items-center gap-2 px-3 py-1.5 bg-amber-100 hover:bg-amber-200 text-amber-700 rounded-lg transition-colors disabled:opacity-50"
    title="Mettre à jour via le MSSS"
  >
    <RefreshCw :size="16" :class="{ 'animate-spin': isUpdating }" />
    <span class="text-sm font-medium">{{ isUpdating ? 'Synchro...' : 'Sync MSSS' }}</span>
  </button>

  <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 transition-colors">
    <X :size="24" />
  </button>
</div>
      </div>

      <RPAForm v-model="form" :mode="mode" @save="handleSave">
        <template #footer-meta>
          <div v-if="mode !== 'create'" class="pt-6 border-t dark:border-gray-700 grid grid-cols-2 md:grid-cols-4 gap-4 text-[11px] uppercase tracking-wider font-bold text-gray-400">
            <div><p>Ajouté</p><p class="text-gray-900 dark:text-gray-200 mt-1">{{ formatDate(form.date_ajout) }}</p></div>
            <div><p>Dernière MAJ</p><p class="text-blue-600 dark:text-blue-400 mt-1">{{ formatDate(form.date_maj) }}</p></div>
            <div><p>Vérification</p><p class="text-gray-900 dark:text-gray-200 mt-1">{{ formatDate(form.derniere_verification) }}</p></div>
            <div><p>Certification</p><p class="text-gray-900 dark:text-gray-200 mt-1">{{ form.date_certification || 'N/A' }}</p></div>
          </div>
        </template>

        <template #actions>
          <div class="flex items-center justify-between">
            <button v-if="mode !== 'create'" type="button" @click="handleDelete" class="px-4 py-2 bg-red-50 dark:bg-red-900/20 text-red-600 rounded-lg flex items-center gap-2 font-bold text-sm">
              <Trash2 :size="18" /> Supprimer
            </button>
            <div v-else></div>

            <div class="flex items-center gap-3">
              <template v-if="mode === 'edit' || mode === 'create'">
                <button type="button" @click="handleCancel" class="px-4 py-2 text-gray-700 dark:text-gray-300 font-bold">Annuler</button>
                <FormKit type="submit" input-class="px-6 py-2 bg-blue-600 text-white rounded-lg font-bold flex items-center gap-2">
                  <Save :size="18" /> {{ mode === 'create' ? 'Créer' : 'Sauvegarder' }}
                </FormKit>
              </template>
              <button v-else type="button" @click="$emit('close')" class="px-6 py-2 bg-gray-900 dark:bg-gray-700 text-white rounded-lg font-bold">Fermer</button>
            </div>
          </div>
        </template>
      </RPAForm>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Building2, X, Edit, Save, Trash2, RefreshCw } from 'lucide-vue-next'
import { GetResidenceForDetails } from '../../../wailsjs/go/main/App'
import RPAForm from './RPAForm.vue'

const props = defineProps({
  rpa: { type: Object, default: () => ({}) },
  mode: { type: String, default: 'view' }
})

const emit = defineEmits(['close', 'save', 'delete'])
const mode = ref(props.mode)
const isUpdating = ref(false)
const form = ref({ ...props.rpa })

onMounted(() => {
  // On charge juste les données passées en props
  form.value = { ...props.rpa }
})
// NOUVELLE FONCTION : Synchro manuelle
const handleManualSync = async () => {
  if (!form.value.registre) return
  
  console.log("📊 AVANT synchro:", JSON.stringify(form.value))
  
  isUpdating.value = true
  try {
    const freshData = await GetResidenceForDetails(form.value.registre, true)
    
    console.log("📥 Données reçues de Go:", JSON.stringify(freshData))
    
    if (freshData) {
      Object.assign(form.value, freshData)
      
      console.log("📊 APRÈS synchro:", JSON.stringify(form.value))
    }
  } catch (err) {
    console.error("Erreur synchro:", err)
    alert("Impossible de synchroniser : " + err)
  } finally {
    isUpdating.value = false
  }
}

const title = computed(() => mode.value === 'create' ? 'Nouveau RPA' : mode.value === 'edit' ? 'Modifier le RPA' : 'Détails du RPA')
const handleSave = (data) => emit('save', data)
const handleCancel = () => {
  if (mode.value === 'create') emit('close')
  else {
    form.value = { ...props.rpa }
    mode.value = 'view'
  }
}
const handleDelete = () => confirm(`Supprimer "${form.value.titre_rpa}" ?`) && emit('delete', form.value)
const formatDate = (d) => d ? new Date(d).toLocaleDateString('fr-CA') : 'N/A'
</script>