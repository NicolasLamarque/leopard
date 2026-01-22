<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-slate-950/40 backdrop-blur-sm animate-in fade-in duration-300">
    
    <div class="bg-white dark:bg-slate-900 w-full max-w-5xl rounded-2xl shadow-2xl border border-slate-200 dark:border-slate-800 overflow-hidden transform transition-all animate-in zoom-in-95 duration-300">
      
      <div class="px-8 py-5 border-b border-slate-100 dark:border-slate-800 flex justify-between items-center bg-slate-50/50 dark:bg-slate-900/50">
        <div class="flex items-center gap-4">
          <div class="p-2.5 bg-teal-500/10 dark:bg-teal-500/20 rounded-lg text-teal-600 dark:text-teal-400">
            <Stethoscope class="w-6 h-6" />
          </div>
          <div>
            <h2 class="text-xl font-bold text-slate-900 dark:text-white tracking-tight">
              {{ medecin?.id ? 'Édition du profil' : 'Nouveau Praticien' }}
            </h2>
            <p class="text-[10px] text-slate-500 uppercase tracking-[0.2em] font-semibold">
              Registre du personnel médical
            </p>
          </div>
        </div>
        <button @click="$emit('close')" class="p-2 text-slate-400 hover:text-slate-600 dark:hover:text-white hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg transition-all">
          <X class="w-5 h-5" />
        </button>
      </div>

      <FormKit
        type="form"
        :value="formInitialValues"
        @submit="handleSubmit"
        :actions="false"
      >
        <div class="p-8 overflow-y-auto max-h-[70vh] custom-scrollbar grid grid-cols-1 md:grid-cols-2 gap-x-12 gap-y-2">
          
          <div class="md:col-span-2 flex items-center gap-3 mb-4 mt-2">
            <div class="h-px flex-1 bg-slate-100 dark:bg-slate-800"></div>
            <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Identité & Licence</span>
            <div class="h-px flex-1 bg-slate-100 dark:bg-slate-800"></div>
          </div>
          
          <FormKit 
            type="text" 
            name="licence" 
            label="Numéro de licence" 
            validation="required" 
            placeholder="Ex: 04602"
          />
          
          <FormKit 
            type="select" 
            name="civilite" 
            label="Civilité" 
            :options="['Docteur', 'Docteure', 'Dr', 'Dre', 'Professeur', 'Professeure']"
          />
          
          <FormKit 
            type="text" 
            name="prenom" 
            label="Prénom" 
            validation="required" 
          />
          
          <FormKit 
            type="text" 
            name="nom" 
            label="Nom" 
            validation="required" 
          />
          
          <div class="md:col-span-2">
            <FormKit 
              type="text" 
              name="nomComplet" 
              label="Nom complet d'affichage" 
              validation="required" 
              help="Format suggéré: Dr Prénom Nom"
            />
          </div>

          <div class="md:col-span-2 flex items-center gap-3 mb-4 mt-6">
            <div class="h-px flex-1 bg-slate-100 dark:bg-slate-800"></div>
            <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Affectation</span>
            <div class="h-px flex-1 bg-slate-100 dark:bg-slate-800"></div>
          </div>
          
          <FormKit 
            type="select" 
            name="specialisation" 
            label="Spécialité" 
            :options="specialites"
          />
          
          <FormKit 
            type="select" 
            name="statut" 
            label="Statut professionnel" 
            :options="['Actif', 'Retraité', 'En congé', 'Autre']"
          />
          
          <FormKit 
            type="text" 
            name="service" 
            label="Service" 
          />
          
          <FormKit 
            type="text" 
            name="departement" 
            label="Département" 
          />

          <div class="md:col-span-2 flex items-center gap-3 mb-4 mt-6">
            <div class="h-px flex-1 bg-slate-100 dark:bg-slate-800"></div>
            <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Coordonnées & Lieu</span>
            <div class="h-px flex-1 bg-slate-100 dark:bg-slate-800"></div>
          </div>

          <div class="md:col-span-2">
             <FormKit type="email" name="email" label="Courriel professionnel" validation="email" />
          </div>
          
          <FormKit type="text" name="telephone" label="Téléphone bureau" />
          <FormKit type="text" name="cellulaire" label="Cellulaire" />
          
          <div class="md:col-span-2">
            <FormKit type="text" name="installation_principale" label="Installation principale" />
          </div>

          <div class="md:col-span-2 grid grid-cols-2 md:grid-cols-4 gap-4">
            <div class="md:col-span-2">
              <FormKit type="text" name="adresse" label="Adresse" />
            </div>
            <FormKit type="text" name="ville" label="Ville" />
            <FormKit type="text" name="code_postal" label="Code postal" />
          </div>

          <div class="md:col-span-2 mt-6">
            <FormKit 
              type="textarea" 
              name="Note_fixe" 
              label="Notes administratives" 
              rows="3"
            />
            
            <div class="mt-4 flex items-center p-4 bg-slate-50 dark:bg-slate-800/50 rounded-xl border border-slate-100 dark:border-slate-800">
              <FormKit 
                type="checkbox" 
                name="actif" 
                label="Le médecin est autorisé à exercer dans le réseau" 
                :value="true"
              />
            </div>
          </div>
        </div>

        <div class="px-8 py-6 border-t border-slate-100 dark:border-slate-800 flex justify-end items-center gap-4 bg-white dark:bg-slate-900">
          <button 
            type="button" 
            @click="$emit('close')" 
            class="px-5 py-2.5 text-slate-500 hover:text-slate-800 dark:hover:text-slate-200 text-sm font-semibold transition-all"
          >
            Annuler
          </button>
          
          <button 
            type="submit" 
            class="bg-slate-900 dark:bg-teal-600 hover:bg-slate-800 dark:hover:bg-teal-500 text-white px-8 py-2.5 rounded-lg font-bold text-sm transition-all shadow-sm flex items-center gap-2"
          >
            <Save class="w-4 h-4" />
            {{ medecin?.id ? 'Sauvegarder les modifications' : 'Enregistrer le praticien' }}
          </button>
        </div>
      </FormKit>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { X, Stethoscope, Save } from 'lucide-vue-next'

const props = defineProps({
  medecin: { type: Object, default: null },
  loading: { type: Boolean, default: false }
})

const emit = defineEmits(['close', 'saved'])

const specialites = [
  'Médecin de famille', 'Médecine générale', 'Médecine interne', 'Cardiologie', 'Neurologie',
  'Psychiatrie', 'Pédiatrie', 'Gériatrie', 'Obstétrique', 'Gynécologie', 'Chirurgie générale',
  'Orthopédie', 'Dermatologie', 'Ophtalmologie', 'ORL', 'Radiologie', 'Anesthésiologie',
  'Urgence', 'Autre'
]

const formInitialValues = computed(() => {
  if (!props.medecin) {
    return { pays: 'Canada', actif: true, statut: 'Actif' }
  }
  return {
    ...props.medecin,
    actif: props.medecin.actif === 1 || props.medecin.actif === true
  }
})

const handleSubmit = (data) => {
  if (!data.nomComplet && (data.prenom || data.nom)) {
    data.nomComplet = `${data.civilite || ''} ${data.prenom || ''} ${data.nom || ''}`.trim()
  }
  const payload = { ...data, actif: data.actif ? 1 : 0 }
  emit('saved', payload)
}
</script>

<style scoped>
@reference "../../style.css"; 

/* DESIGN PROFESSIONNEL POUR FORMKIT (Tailwind v4 strategy) */
:deep(.formkit-outer) { 
  @apply w-full mb-4; 
}

:deep(.formkit-label) { 
  @apply block text-[11px] font-bold text-slate-500 dark:text-slate-400 uppercase tracking-widest mb-1.5 ml-0.5; 
}

:deep(.formkit-inner) { 
  @apply relative flex items-center bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg transition-all duration-200 shadow-sm ring-0;
}

:deep(.formkit-inner:focus-within) {
  @apply border-teal-500 dark:border-teal-500 ring-2 ring-teal-500/10;
}

:deep(.formkit-input) { 
  @apply w-full bg-transparent px-3.5 py-2.5 text-sm text-slate-700 dark:text-slate-100 placeholder:text-slate-400 outline-none; 
}

:deep(.formkit-help) {
  @apply text-[10px] text-slate-400 mt-1.5 ml-0.5;
}

/* Custom styles for selects and specific inputs */
:deep(select.formkit-input) {
  @apply pr-10 appearance-none bg-no-repeat bg-[right_0.75rem_center] bg-[length:1em_1em];
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='%2394a3b8' stroke-width='2'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' d='M19 9l-7 7-7-7'/%3E%3C/svg%3E");
}

/* Checkbox specific styling */
:deep(.formkit-type-checkbox .formkit-inner) {
  @apply border-none bg-transparent shadow-none ring-0 focus-within:ring-0 w-auto;
}

:deep(.formkit-type-checkbox .formkit-input) {
  @apply w-5 h-5 rounded border-slate-300 text-teal-600 focus:ring-teal-500 mr-3;
}

:deep(.formkit-type-checkbox .formkit-label) {
  @apply normal-case tracking-normal text-sm font-medium text-slate-600 dark:text-slate-300 mb-0;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-700 rounded-full;
}
</style>