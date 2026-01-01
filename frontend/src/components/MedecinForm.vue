<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-gray-950/60 backdrop-blur-md animate-in fade-in duration-300">
    
    <div class="bg-white dark:bg-gray-900 w-full max-w-4xl rounded-[2.5rem] shadow-2xl border border-gray-200 dark:border-gray-800 overflow-hidden transform transition-all animate-in zoom-in-95 duration-300">
      
      <div class="px-10 py-7 border-b border-gray-100 dark:border-gray-800 flex justify-between items-center bg-gray-50/50 dark:bg-gray-800/30">
        <div class="flex items-center gap-5">
          <div class="p-4 bg-emerald-500/10 rounded-2xl text-emerald-500 shadow-inner">
            <Stethoscope class="w-7 h-7" />
          </div>
          <div>
            <h2 class="text-2xl font-black text-gray-900 dark:text-white tracking-tighter uppercase">
              {{ medecin?.id ? 'Édition Profil' : 'Nouveau Praticien' }}
            </h2>
            <p class="text-[10px] text-emerald-600 dark:text-emerald-500 uppercase tracking-[0.3em] font-black opacity-70">Certification Médicale</p>
          </div>
        </div>
        <button @click="$emit('close')" class="p-3 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-full transition-all text-gray-400 hover:text-red-500">
          <X class="w-6 h-6" />
        </button>
      </div>

      <FormKit
        type="form"
        :value="formInitialValues"
        @submit="handleSubmit"
        :actions="false"
      >
        <div class="p-10 overflow-y-auto max-h-[70vh] custom-scrollbar">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-x-12 gap-y-8">
            
            <div class="md:col-span-2 space-y-6">
              <div class="flex items-center gap-3">
                <span class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></span>
                <span class="text-[10px] font-black uppercase tracking-[0.2em] text-gray-400">01. Informations Légales</span>
                <span class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></span>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <FormKit type="text" name="licence" label="Numéro de licence *" validation="required" placeholder="RPPS / ADELI" />
                <FormKit type="text" name="nomComplet" label="Nom complet du médecin *" validation="required" placeholder="Dr. Jean Tremblay" />
                <div class="md:col-span-2">
                   <FormKit type="text" name="specialisation" label="Spécialisation" placeholder="ex: Médecine générale, Cardiologie..." />
                </div>
              </div>
            </div>

            <div class="space-y-6">
              <div class="flex items-center gap-3">
                <span class="text-[10px] font-black uppercase tracking-[0.2em] text-gray-400 px-2">02. Contact direct</span>
                <span class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></span>
              </div>
              <div class="grid grid-cols-3 gap-4">
                <div class="col-span-2"><FormKit type="text" name="telephone" label="Bureau" placeholder="514-..." /></div>
                <div class="col-span-1"><FormKit type="text" name="extension" label="Ext." placeholder="123" /></div>
              </div>
              <FormKit type="text" name="cellulaire" label="Cellulaire personnel" placeholder="438-..." />
              <FormKit type="email" name="email" label="Courriel" placeholder="docteur@exemple.com" validation="email" />
            </div>

            <div class="space-y-6">
              <div class="flex items-center gap-3">
                <span class="text-[10px] font-black uppercase tracking-[0.2em] text-gray-400 px-2">03. Localisation</span>
                <span class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></span>
              </div>
              <FormKit type="text" name="adresse" label="Adresse du cabinet" />
              <div class="grid grid-cols-2 gap-4">
                <FormKit type="text" name="ville" label="Ville" />
                <FormKit type="text" name="code_postal" label="Code postal" />
              </div>
              <FormKit type="text" name="pays" label="Pays" value="Canada" />
            </div>

            <div class="md:col-span-2 space-y-6 pt-4">
              <div class="flex items-center gap-3">
                <span class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></span>
                <span class="text-[10px] font-black uppercase tracking-[0.2em] text-gray-400">04. Administration</span>
                <span class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></span>
              </div>
              <FormKit type="textarea" name="note_fixe" label="Notes permanentes" rows="3" />
              <div class="flex items-center p-4 bg-emerald-500/5 rounded-2xl border border-emerald-500/10">
                <FormKit type="checkbox" name="actif" label="Le médecin est actuellement actif dans le réseau" :value="true" />
              </div>
            </div>
          </div>
        </div>

        <div class="px-10 py-8 border-t border-gray-100 dark:border-gray-800 flex justify-end items-center gap-6 bg-gray-50/50 dark:bg-gray-800/30">
          <button type="button" @click="$emit('close')" class="text-xs font-black text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 uppercase tracking-[0.2em] transition-all">
            Annuler
          </button>
          <button 
            type="submit" 
            class="bg-emerald-600 hover:bg-emerald-500 text-white px-12 py-4 rounded-2xl font-black text-xs uppercase tracking-[0.2em] shadow-2xl shadow-emerald-500/30 transition-all active:scale-95 flex items-center gap-3"
          >
            <Save class="w-4 h-4" />
            {{ medecin?.id ? 'Mettre à jour' : 'Confirmer la création' }}
          </button>
        </div>
      </FormKit>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { X, Stethoscope, Save, User, Phone, MapPin } from 'lucide-vue-next'

const props = defineProps({
  medecin: { type: Object, default: null },
  loading: { type: Boolean, default: false }
})

const emit = defineEmits(['close', 'saved'])

// Mapping technique identique à ton code initial
const formInitialValues = computed(() => {
  if (!props.medecin) return { pays: 'Canada', actif: true }
  return {
    ...props.medecin,
    actif: props.medecin.actif === 1
  }
})

const handleSubmit = (data) => {
  const payload = {
    ...data,
    actif: data.actif ? 1 : 0
  }
  emit('saved', payload)
}
</script>

<style scoped>
@reference "tailwindcss";

/* DESIGN PREMIUM POUR FORMKIT */
:deep(.formkit-outer) { @apply mb-0; }
:deep(.formkit-label) { 
  @apply block text-[10px] font-black text-emerald-600 dark:text-emerald-500 uppercase tracking-[0.2em] mb-2 ml-1; 
}
:deep(.formkit-input) {
  @apply w-full px-6 py-4 bg-gray-100 dark:bg-gray-800/80 border-2 border-transparent dark:border-gray-800/50 rounded-2xl outline-none focus:ring-4 focus:ring-emerald-500/10 focus:border-emerald-500/50 dark:text-white text-sm font-bold transition-all placeholder:text-gray-400 shadow-inner;
}
:deep(.formkit-input[type="checkbox"]) {
  @apply w-6 h-6 rounded-lg bg-emerald-500;
}
:deep(.formkit-message) { @apply text-[10px] font-bold text-red-500 uppercase tracking-widest mt-2 ml-2; }

/* Custom Scrollbar */
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { @apply bg-gray-200 dark:bg-gray-800 rounded-full; }
</style>