<template>
  <div 
    v-if="isOpen" 
    class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-slate-900/50 backdrop-blur-sm animate-in fade-in duration-200"
    @click.self="$emit('close')"
  >
    <div class="bg-white dark:bg-slate-900 w-full max-w-3xl rounded-2xl shadow-xl border border-slate-200 dark:border-slate-800 overflow-hidden animate-in zoom-in-95 duration-200">
      
      <!-- Header -->
      <div class="px-7 py-5 border-b border-slate-200 dark:border-slate-800 bg-slate-50/80 dark:bg-slate-900/50">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <div class="p-2.5 bg-teal-100 dark:bg-teal-900/30 rounded-xl text-teal-600 dark:text-teal-400">
              <Scale :size="24" />
            </div>
            <div>
              <h2 class="text-xl font-bold text-slate-900 dark:text-white">
                {{ notaire ? 'Modifier le notaire' : 'Formulaire Complet' }}
              </h2>
              <p class="text-sm text-teal-600 dark:text-teal-400 font-medium mt-0.5">
                {{ notaire ? 'Toutes les informations du notaire' : 'Créer un nouveau notaire' }}
              </p>
            </div>
          </div>
          
          <button 
            @click="$emit('close')" 
            class="p-2 hover:bg-slate-200 dark:hover:bg-slate-800 rounded-xl transition-colors text-slate-400 hover:text-slate-600 dark:hover:text-slate-300"
          >
            <X :size="20" />
          </button>
        </div>
      </div>

      <!-- Formulaire COMPLET -->
      <div class="p-7 max-h-[70vh] overflow-y-auto">
        <FormKit
          type="form"
          id="addNotaireForm"
          v-model="formData"
          @submit="handleCreate"
          :submit-attrs="{
            inputClass: 'w-full bg-teal-600 hover:bg-teal-700 text-white font-semibold py-3 px-6 rounded-xl transition-all shadow-sm hover:shadow-md active:scale-[0.98] mt-6'
          }"
          :submit-label="notaire ? 'Sauvegarder les modifications' : 'Créer le notaire'"
        >
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-5">
            
            <!-- Civilité -->
            <FormKit 
              type="select" 
              name="civilite" 
              label="Civilité" 
              :options="['Me', 'Mtre', 'M.', 'Mme']" 
              value="Me"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all cursor-pointer'
              }"
            />
            
            <!-- Secteur d'activité -->
            <FormKit 
              type="text" 
              name="secteur_activite" 
              label="Secteur d'activité" 
              placeholder="ex: Immobilier, Succession..."
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 placeholder-slate-400 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Prénom -->
            <FormKit 
              type="text" 
              name="prenom" 
              label="Prénom" 
              validation="required"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Nom -->
            <FormKit 
              type="text" 
              name="nom" 
              label="Nom" 
              validation="required"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Téléphone -->
            <FormKit 
              type="text" 
              name="telephone" 
              label="Téléphone"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Cellulaire -->
            <FormKit 
              type="text" 
              name="cellulaire" 
              label="Cellulaire"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Télécopieur -->
            <FormKit 
              type="text" 
              name="telecopieur" 
              label="Télécopieur / Fax"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Email -->
            <FormKit 
              type="text" 
              name="email" 
              label="Courriel professionnel" 
              validation="email"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Adresse complète - Full width -->
            <div class="md:col-span-2">
              <FormKit 
                type="text" 
                name="adresse" 
                label="Adresse complète"
                :classes="{
                  label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                  input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
                }"
              />
            </div>
            
            <!-- Ville -->
            <FormKit 
              type="text" 
              name="ville" 
              label="Ville"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />
            
            <!-- Code Postal -->
            <FormKit 
              type="text" 
              name="code_postal" 
              label="Code Postal"
              :classes="{
                label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
              }"
            />

            <!-- Note fixe - Full width -->
            <div class="md:col-span-2">
              <FormKit 
                type="textarea" 
                name="note_fixe" 
                label="Notes / Informations complémentaires"
                :classes="{
                  label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
                  input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all min-h-[100px]'
                }"
              />
            </div>
          </div>
        </FormKit>
      </div>
    </div>
  </div>
</template>

<script setup>
import { X, Scale } from 'lucide-vue-next'
import { ref, watch } from 'vue'

const props = defineProps({ 
  isOpen: Boolean,
  notaire: { type: Object, default: null }
})

const emit = defineEmits(['close', 'saved'])

const formData = ref({})

// Pré-remplir le formulaire si on édite
watch(() => props.notaire, (newVal) => {
  if (newVal) {
    formData.value = { ...newVal }
  } else {
    formData.value = {}
  }
}, { immediate: true })

const handleCreate = async (data) => {
  try {
    if (props.notaire?.id) {
      // Modification
      await window.go.main.App.UpdateNotaire({ ...data, id: props.notaire.id })
    } else {
      // Création
      await window.go.main.App.CreateNotaire({ ...data, created_by: 1 })
    }
    
    emit('saved')
    emit('close')
  } catch (err) {
    console.error("Erreur:", err)
  }
}
</script>