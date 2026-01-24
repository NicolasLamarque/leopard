<template>
  <div class="bg-white dark:bg-slate-900 rounded-2xl p-8 border border-slate-200 dark:border-slate-800 shadow-sm">
    <!-- Header -->
    <div class="mb-6 pb-5 border-b border-slate-200 dark:border-slate-700">
      <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-1">
        Modifier les informations
      </h3>
      <p class="text-sm text-slate-600 dark:text-slate-400">
        Mettez à jour les coordonnées professionnelles du notaire
      </p>
    </div>

    <FormKit
      type="form"
      v-model="formData"
      @submit="handleSubmit"
      :submit-attrs="{
        inputClass: 'w-full bg-teal-600 hover:bg-teal-700 text-white font-semibold py-3 px-6 rounded-xl transition-all shadow-sm hover:shadow-md active:scale-[0.98] mt-6'
      }"
      submit-label="Sauvegarder les modifications"
    >
      <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-5">
        
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
        
        <FormKit 
          type="select" 
          name="civilite" 
          label="Civilité" 
          :options="['Me', 'Mtre', 'M.', 'Mme']"
          :classes="{
            label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
            input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all cursor-pointer'
          }"
        />
        
        <FormKit 
          type="text" 
          name="email" 
          label="Courriel" 
          validation="email"
          :classes="{
            label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
            input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
          }"
        />
        
        <FormKit 
          type="text" 
          name="telephone" 
          label="Téléphone"
          :classes="{
            label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
            input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
          }"
        />
        
        <FormKit 
          type="text" 
          name="cellulaire" 
          label="Cellulaire"
          :classes="{
            label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
            input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
          }"
        />
        
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
        
        <FormKit 
          type="text" 
          name="ville" 
          label="Ville"
          :classes="{
            label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
            input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
          }"
        />
        
        <FormKit 
          type="text" 
          name="code_postal" 
          label="Code Postal"
          :classes="{
            label: 'block text-sm font-medium text-slate-700 dark:text-slate-300 mb-1.5',
            input: 'w-full px-4 py-2.5 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all'
          }"
        />
      </div>
    </FormKit>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  notaire: { type: Object, default: () => ({}) }
})

const formData = ref({ ...props.notaire })

watch(() => props.notaire, (newVal) => {
  formData.value = { ...newVal }
}, { deep: true })

const handleSubmit = async (data) => {
  try {
    await window.go.main.App.UpdateNotaire(data)
  } catch (err) {
    console.error(err)
  }
}
</script>