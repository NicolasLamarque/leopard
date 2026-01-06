<template>
  <FormKit 
    type="form" 
    v-model="internalData"
    @submit="data => $emit('save', data)"
    :actions="false"
    form-class="flex-1 overflow-y-auto flex flex-col"
  >
    <div class="p-6 space-y-8 flex-1">
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="md:col-span-2">
          <FormKit
            type="text"
            name="titre"
            label="Nom du RPA *"
            validation="required"
            :disabled="mode === 'view'"
            :classes="standardClasses"
            input-class="py-3 font-medium text-lg" 
          />
        </div>

        <FormKit
          type="text"
          name="registre"
          label="Numéro de registre"
          disabled
          :classes="standardClasses"
          input-class="bg-gray-50 dark:bg-gray-900 font-mono opacity-70"
        />
        <FormKit
  type="text"
  name="numero_interne"
  label="Numéro Interne (MSSS)"
  disabled
  :classes="standardClasses"
  input-class="bg-gray-50 dark:bg-gray-900 font-mono opacity-70"
/>

        <FormKit
          type="select"
          name="statut"
          label="Statut *"
          :disabled="mode === 'view'"
          :options="[
            { label: 'Actif', value: 'actif' },
            { label: 'Fermé', value: 'ferme' },
            { label: 'Suspendu', value: 'suspendu' }
          ]"
          :classes="standardClasses"
        />
      </div>

      <div class="bg-blue-50/50 dark:bg-blue-900/10 rounded-xl border border-blue-100 dark:border-blue-800/30 p-5">
        <h3 class="text-sm font-bold text-blue-800 dark:text-blue-300 uppercase tracking-wider mb-3 flex items-center gap-2">
          <ClipboardList :size="18" /> 7 - Services offerts par la résidence
        </h3>
        <div v-if="mode === 'view'" class="text-gray-700 dark:text-gray-300 text-sm leading-relaxed whitespace-pre-wrap italic">
          {{ internalData.services || 'Aucune information détaillée n\'a été récupérée.' }}
        </div>
        <FormKit
          v-else
          type="textarea"
          name="services"
          rows="4"
          :classes="standardClasses"
          input-class="resize-none"
        />
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <FormKit
            type="textarea"
            name="adresse"
            label="Adresse"
            :disabled="mode === 'view'"
            rows="2"
            :classes="standardClasses"
            input-class="resize-none"
          />
          <div class="grid grid-cols-2 gap-4">
            <FormKit type="text" name="ville" label="Ville" :disabled="mode === 'view'" :classes="standardClasses" />
            <FormKit type="text" name="code_postal" label="Code Postal" :disabled="mode === 'view'" :classes="standardClasses" input-class="font-mono" />
          </div>
        </div>

        <div class="space-y-4">
          <FormKit
            type="tel"
            name="telephone"
            label="Téléphone"
            :disabled="mode === 'view'"
            :classes="standardClasses"
          />
          <FormKit
            type="number"
            name="capacite"
            label="Capacité (places)"
            :disabled="mode === 'view'"
            :classes="standardClasses"
          />
        </div>
      </div>

      <FormKit
        type="textarea"
        name="notes"
        label="Notes internes"
        placeholder="Notes privées..."
        :disabled="mode === 'view'"
        :classes="standardClasses"
        input-class="resize-none"
      />

      <slot name="footer-meta" />
    </div>

    <div class="px-6 py-4 border-t dark:border-gray-700 bg-gray-50 dark:bg-gray-900/50">
      <slot name="actions" />
    </div>
  </FormKit>
</template>

<script setup>
import { ref, watch } from 'vue'
import { ClipboardList } from 'lucide-vue-next'

const props = defineProps({
  modelValue: Object,
  mode: String
})

const emit = defineEmits(['update:modelValue', 'save'])

// On garde une copie locale pour éviter les boucles infinies de v-model
const internalData = ref({ ...props.modelValue })

// Crucial : Si le parent reçoit les données de Wails, on met à jour le formulaire
watch(() => props.modelValue, (newVal) => {
  internalData.value = { ...newVal }
}, { deep: true })

// Si on tape dans le formulaire, on informe le parent
watch(internalData, (newVal) => {
  emit('update:modelValue', newVal)
}, { deep: true })

const standardClasses = {
  outer: 'mb-0',
  label: 'block text-xs font-bold text-gray-500 dark:text-gray-400 uppercase mb-2',
  inner: 'w-full border dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 focus-within:ring-2 focus-within:ring-blue-500 overflow-hidden transition-shadow',
  input: 'w-full px-4 py-2 bg-transparent text-gray-900 dark:text-white border-none focus:ring-0 disabled:bg-gray-50/50 dark:disabled:bg-gray-900/50 placeholder:text-gray-400',
  message: 'text-red-500 text-[10px] font-bold mt-1 uppercase',
}
</script>