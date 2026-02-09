<template>
  <div class="space-y-8">
    
    <!-- Alerte Loi 25 : Rappel visuel de la confidentialité -->
    <div class="group relative overflow-hidden bg-gradient-to-br from-amber-50 to-orange-50 dark:from-amber-900/10 dark:to-orange-900/10 border border-amber-200/50 dark:border-amber-800/50 p-6 rounded-2xl shadow-sm hover:shadow-md transition-all duration-300">
      <div class="absolute inset-0 bg-gradient-to-r from-amber-500/0 via-amber-500/5 to-amber-500/0 opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
      <div class="relative flex gap-4 items-start">
        <div class="shrink-0 p-3 bg-gradient-to-br from-amber-500/10 to-orange-500/10 backdrop-blur-sm rounded-2xl border border-amber-200/50 dark:border-amber-800/50">
          <ShieldCheck :size="24" class="text-amber-600 dark:text-amber-400" />
        </div>
        <div class="flex-1">
          <h4 class="text-sm font-black text-amber-800 dark:text-amber-300 uppercase tracking-wide mb-2">Conformité Loi 25</h4>
          <p class="text-xs text-amber-700/90 dark:text-amber-400/90 leading-relaxed">
            Les champs marqués d'un cadenas <Lock :size="12" class="inline-block mx-1 text-amber-600 dark:text-amber-400" /> sont automatiquement chiffrés (AES-256) avant la sauvegarde en base de données pour garantir la protection des données personnelles.
          </p>
        </div>
      </div>
    </div>

    <!-- Formulaire Principal via FormKit -->
    <FormKit
      type="form"
      :actions="false"
      @submit="submitForm"
      #default="{ value }"
    >
      <div class="space-y-10">
        
        <!-- SECTION : Identification -->
        <div class="group space-y-6 p-8 bg-gradient-to-br from-slate-50/50 to-white dark:from-slate-800/30 dark:to-slate-900/30 rounded-2xl border border-slate-200/50 dark:border-slate-700/50 hover:border-teal-500/30 transition-all duration-300">
          <div class="flex items-center gap-3 pb-3 border-b border-slate-200/50 dark:border-slate-700/50">
            <div class="p-2 bg-gradient-to-br from-teal-500/10 to-emerald-500/10 rounded-xl">
              <User :size="20" class="text-teal-600 dark:text-teal-400" />
            </div>
            <h3 class="text-sm font-black uppercase tracking-[0.15em] text-slate-600 dark:text-slate-400">Identification</h3>
          </div>

          <FormKit
            type="text"
            name="nom_complet"
            label="Nom complet de l'intervenant"
            placeholder="ex: Dr. Julie Tremblay"
            validation="required"
          />

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <FormKit
              type="select"
              name="type_intervenant"
              label="Type d'intervenant"
              :options="['Médecin', 'Travailleur Social', 'Notaire', 'Infirmier(ère)', 'Pharmacien', 'Autre']"
              validation="required"
            />
            <FormKit
              type="text"
              name="organisation"
              label="Organisation / Clinique"
              placeholder="ex: CIUSSS Capitale-Nationale"
            />
          </div>
        </div>

        <!-- SECTION : Coordonnées (Zone Sensible) -->
        <div class="group space-y-6 p-8 bg-gradient-to-br from-amber-50/30 to-orange-50/30 dark:from-amber-900/10 dark:to-orange-900/10 rounded-2xl border border-amber-200/50 dark:border-amber-700/50 hover:border-amber-500/50 transition-all duration-300">
          <div class="flex items-center gap-3 pb-3 border-b border-amber-200/50 dark:border-amber-700/50">
            <div class="p-2 bg-gradient-to-br from-amber-500/10 to-orange-500/10 rounded-xl">
              <Lock :size="20" class="text-amber-600 dark:text-amber-400" />
            </div>
            <h3 class="text-sm font-black uppercase tracking-[0.15em] text-amber-700 dark:text-amber-400">Coordonnées Sécurisées</h3>
          </div>

          <FormKit
            type="email"
            name="email_prive"
            label="Courriel personnel (Crypté)"
            validation="required|email"
          
          >
            <template #label="{ label }">
              <span class="flex items-center gap-2">
                {{ label }} 
                <Lock :size="14" class="text-amber-500 dark:text-amber-400" />
              </span>
            </template>
          </FormKit>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <FormKit
              type="tel"
              name="telephone_bureau"
              label="Tél. Bureau"
              placeholder="514-555-0000"
            />
            <FormKit
              type="tel"
              name="cellulaire_prive"
              label="Cellulaire"
            >
              <template #label="{ label }">
                <span class="flex items-center gap-2">
                  {{ label }} 
                  <Lock :size="14" class="text-amber-500 dark:text-amber-400" />
                </span>
              </template>
            </FormKit>
          </div>
        </div>

        <!-- SECTION : Notes et Expertise -->
        <div class="group space-y-6 p-8 bg-gradient-to-br from-slate-50/50 to-white dark:from-slate-800/30 dark:to-slate-900/30 rounded-2xl border border-slate-200/50 dark:border-slate-700/50 hover:border-slate-500/30 transition-all duration-300">
          <div class="flex items-center gap-3 pb-3 border-b border-slate-200/50 dark:border-slate-700/50">
            <div class="p-2 bg-gradient-to-br from-slate-500/10 to-slate-600/10 rounded-xl">
              <FileText :size="20" class="text-slate-600 dark:text-slate-400" />
            </div>
            <h3 class="text-sm font-black uppercase tracking-[0.15em] text-slate-600 dark:text-slate-400">Notes de collaboration</h3>
          </div>

          <FormKit
            type="textarea"
            name="notes"
            label="Observations ou champs d'expertise"
            rows="5"
            placeholder="Détails sur la spécialité, heures de contact préférées, informations complémentaires..."
          />
        </div>
      </div>

      <!-- Actions du Formulaire -->
      <div class="mt-10 pt-8 border-t border-slate-200/50 dark:border-slate-700/50 flex justify-end gap-4">
        <button 
          type="button"
          @click="$emit('cancel')"
          class="group px-6 py-3.5 rounded-2xl text-sm font-bold text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-all duration-300 flex items-center gap-2"
        >
          <X :size="18" class="group-hover:rotate-90 transition-transform duration-300" />
          Annuler
        </button>
        <button 
          type="submit"
          class="relative group px-8 py-3.5 bg-gradient-to-r from-teal-600 to-emerald-600 hover:from-teal-500 hover:to-emerald-500 text-white rounded-2xl text-sm font-bold shadow-xl shadow-teal-500/30 hover:shadow-2xl hover:shadow-teal-500/50 transition-all duration-300 active:scale-95 flex items-center gap-3 overflow-hidden"
        >
          <div class="absolute inset-0 bg-gradient-to-r from-white/0 via-white/25 to-white/0 -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
          <Save :size="18" class="relative" />
          <span class="relative">Enregistrer l'intervenant</span>
        </button>
      </div>
    </FormKit>
  </div>
</template>

<script setup>
import { ShieldCheck, Lock, User, Save, FileText, X } from 'lucide-vue-next'

const props = defineProps({
  intervenantId: [String, Number]
})

const emit = defineEmits(['saved', 'cancel'])

const submitForm = (formData) => {
  console.log("Données prêtes pour le backend Go (Wails) :", formData)
  // Simulation de sauvegarde
  emit('saved')
}
</script>

<style>
@reference "tailwindcss";

/* FormKit Custom Styling pour Design Zen */
.formkit-outer {
  @apply mb-6;
}

.formkit-label {
  @apply block text-xs font-black uppercase tracking-[0.1em] text-slate-600 dark:text-slate-400 mb-3 ml-1;
}

.formkit-inner {
  @apply relative shadow-none;
}

.formkit-input,
.formkit-input[type="text"],
.formkit-input[type="email"],
.formkit-input[type="tel"],
.formkit-input[type="textarea"],
select.formkit-input {
  @apply w-full bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70 rounded-2xl px-4 py-3.5 text-sm outline-none transition-all duration-300 
         focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500 dark:text-white placeholder:text-slate-400
         hover:border-slate-300 dark:hover:border-slate-600;
}

textarea.formkit-input {
  @apply resize-none;
}

.formkit-messages {
  @apply list-none p-0 mt-2 ml-1 space-y-1;
}

.formkit-message {
  @apply text-red-500 dark:text-red-400 text-xs font-bold uppercase tracking-tight flex items-center gap-1.5;
}

.formkit-message::before {
  content: "⚠";
  @apply text-red-500;
}

/* Select dropdown styling */
.formkit-input[type="select"] {
  @apply appearance-none bg-no-repeat cursor-pointer;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3E%3Cpath stroke='%236B7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3E%3C/svg%3E");
  background-position: right 1rem center;
  background-size: 1.25rem;
}

/* Dark mode select arrow */
.dark .formkit-input[type="select"] {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3E%3Cpath stroke='%9CA3AF' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3E%3C/svg%3E");
}
</style>