<template>
  <Transition name="modal">
    <div 
      v-if="show"
      class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/60 backdrop-blur-md"
      @click.self="$emit('close')"
    >
      <div class="bg-white dark:bg-slate-900 w-full max-w-5xl rounded-3xl shadow-2xl overflow-hidden border border-slate-200/50 dark:border-slate-800/50 animate-modal-in">
        
        <!-- Header avec gradient -->
        <div class="relative px-10 py-8 bg-gradient-to-br from-teal-600 via-emerald-600 to-teal-700 text-white overflow-hidden">
          <!-- Effet de lumière animé -->
          <div class="absolute inset-0 bg-gradient-to-r from-white/0 via-white/10 to-white/0 -translate-x-full animate-shimmer"></div>
          
          <div class="relative flex justify-between items-start">
            <div class="flex items-center gap-6">
              <div class="p-5 bg-white/10 backdrop-blur-xl rounded-2xl border border-white/20 shadow-2xl">
                <UserCog :size="36" class="drop-shadow-lg" />
              </div>
              <div>
                <h2 class="text-3xl font-black mb-2 tracking-tight">{{ intervenant.nom_complet }}</h2>
                <p class="text-teal-100 font-medium text-lg">{{ intervenant.titre_emploi || 'Poste non défini' }}</p>
              </div>
            </div>
            <button 
              @click="$emit('close')" 
              class="group p-2.5 hover:bg-white/10 rounded-xl transition-all duration-300 backdrop-blur-sm"
            >
              <X :size="24" class="group-hover:rotate-90 transition-transform duration-300" />
            </button>
          </div>
        </div>

        <!-- Contenu scrollable -->
        <div class="p-10 space-y-8 max-h-[70vh] overflow-y-auto custom-scrollbar">
          
          <!-- Grid principal -->
          <section class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <!-- Affectation -->
            <div class="space-y-5 p-6 bg-gradient-to-br from-slate-50/50 to-white dark:from-slate-800/30 dark:to-slate-900/30 rounded-2xl border border-slate-200/50 dark:border-slate-700/50">
              <h3 class="text-sm font-black text-slate-500 dark:text-slate-400 uppercase tracking-[0.15em] mb-5 flex items-center gap-3 pb-3 border-b border-slate-200/50 dark:border-slate-700/50">
                <div class="p-2 bg-blue-500/10 rounded-lg">
                  <Building2 :size="18" class="text-blue-600 dark:text-blue-400" />
                </div>
                Affectation
              </h3>
              <div class="space-y-4">
                <InfoItem label="Direction" :value="intervenant.direction" />
                <InfoItem label="Installation" :value="intervenant.installation" />
                <InfoItem label="Ressource Réseau" :value="intervenant.personne_ressource_reseau_dir" />
              </div>
            </div>

            <!-- Professionnel -->
            <div class="space-y-5 p-6 bg-gradient-to-br from-amber-50/30 to-orange-50/30 dark:from-amber-900/10 dark:to-orange-900/10 rounded-2xl border border-amber-200/50 dark:border-amber-700/50">
              <h3 class="text-sm font-black text-amber-700 dark:text-amber-400 uppercase tracking-[0.15em] mb-5 flex items-center gap-3 pb-3 border-b border-amber-200/50 dark:border-amber-700/50">
                <div class="p-2 bg-amber-500/10 rounded-lg">
                  <Award :size="18" class="text-amber-600 dark:text-amber-400" />
                </div>
                Professionnel
              </h3>
              <div class="space-y-4">
                <div class="flex flex-wrap gap-2">
                  <span 
                    v-if="intervenant.is_intervenant_social" 
                    class="px-3 py-1.5 bg-gradient-to-r from-emerald-100 to-teal-100 dark:from-emerald-900/30 dark:to-teal-900/30 text-emerald-700 dark:text-emerald-400 text-xs rounded-xl font-black uppercase tracking-wide border border-emerald-200 dark:border-emerald-800"
                  >
                    Intervenant Social
                  </span>
                  <span 
                    v-else 
                    class="px-3 py-1.5 bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 text-xs rounded-xl font-black uppercase tracking-wide border border-slate-200 dark:border-slate-700"
                  >
                    Personnel Administratif
                  </span>
                </div>
                <InfoItem v-if="intervenant.numero_permis" label="No. Permis" :value="intervenant.numero_permis" />
                <InfoItem v-if="intervenant.ordre_professionnel" label="Ordre" :value="intervenant.ordre_professionnel" />
              </div>
            </div>
          </section>

          <!-- Coordonnées -->
          <section class="p-8 bg-gradient-to-br from-slate-50 to-slate-100/50 dark:from-slate-800/50 dark:to-slate-900/50 rounded-2xl border border-slate-200/50 dark:border-slate-700/50">
            <h3 class="text-sm font-black text-slate-600 dark:text-slate-400 uppercase tracking-[0.15em] mb-6 flex items-center gap-3 pb-3 border-b border-slate-200/50 dark:border-slate-700/50">
              <div class="p-2 bg-slate-500/10 rounded-lg">
                <Phone :size="18" class="text-slate-600 dark:text-slate-400" />
              </div>
              Coordonnées
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <!-- Professionnel -->
              <div class="space-y-4">
                <div class="flex items-center gap-2 mb-3">
                  <div class="w-2 h-2 rounded-full bg-emerald-500"></div>
                  <p class="text-xs font-black text-emerald-600 dark:text-emerald-400 uppercase tracking-wide">Professionnel</p>
                </div>
                <ContactRow label="Bureau" :value="intervenant.telephone" :ext="intervenant.poste" icon="Phone" />
                <ContactRow label="Cell Pro" :value="intervenant.cellulaire_pro" icon="Smartphone" />
                <ContactRow label="Email Pro" :value="intervenant.courriel_professionnel" icon="Mail" />
              </div>

              <!-- Personnel (Crypté) -->
              <div class="space-y-4 border-l-2 border-amber-200 dark:border-amber-800 pl-8">
                <div class="flex items-center gap-2 mb-3">
                  <Lock :size="14" class="text-amber-600 dark:text-amber-400" />
                  <p class="text-xs font-black text-amber-600 dark:text-amber-400 uppercase tracking-wide">Personnel (Crypté AES-256)</p>
                </div>
                <ContactRow label="Cell Perso" :value="intervenant.cellulaire_perso" icon="Smartphone" is-private />
                <ContactRow label="Email Perso" :value="intervenant.courriel_personnel" icon="Mail" is-private />
              </div>
            </div>
          </section>

          <!-- Notes permanentes -->
          <section 
            v-if="intervenant.note_fixe" 
            class="p-6 bg-gradient-to-br from-amber-50 to-orange-50 dark:from-amber-900/10 dark:to-orange-900/10 rounded-2xl border-l-4 border-amber-500"
          >
            <div class="flex items-center gap-2 mb-3">
              <FileText :size="16" class="text-amber-600 dark:text-amber-400" />
              <p class="text-xs font-black text-amber-800 dark:text-amber-400 uppercase tracking-wide">Notes permanentes</p>
            </div>
            <p class="text-sm text-slate-700 dark:text-slate-300 leading-relaxed">{{ intervenant.note_fixe }}</p>
          </section>
        </div>

        <!-- Footer avec actions -->
        <div class="px-10 py-6 bg-slate-50 dark:bg-slate-900/50 border-t border-slate-200/50 dark:border-slate-700/50 flex justify-between items-center backdrop-blur-sm">
          <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400">
            <ShieldCheck :size="14" class="text-teal-500" />
            <span class="font-bold">Données protégées par chiffrement AES-256</span>
          </div>
          <div class="flex gap-3">
            <button 
              @click="$emit('close')" 
              class="px-5 py-2.5 text-slate-600 dark:text-slate-400 hover:bg-slate-200 dark:hover:bg-slate-800 rounded-xl font-bold text-sm transition-all duration-300"
            >
              Fermer
            </button>
            <button 
              @click="$emit('edit', intervenant)" 
              class="group relative px-6 py-2.5 bg-gradient-to-r from-teal-600 to-emerald-600 hover:from-teal-500 hover:to-emerald-500 text-white rounded-xl font-bold text-sm shadow-lg shadow-teal-500/30 transition-all duration-300 flex items-center gap-2 overflow-hidden"
            >
              <div class="absolute inset-0 bg-gradient-to-r from-white/0 via-white/25 to-white/0 -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
              <Edit2 :size="16" class="relative" />
              <span class="relative">Modifier</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { X, UserCog, Building2, Award, Phone, Smartphone, Mail, Edit2, Lock, FileText, ShieldCheck } from 'lucide-vue-next'

defineProps({ 
  intervenant: Object,
  show: { type: Boolean, default: true }
})

// Composants simplifiés inline (ou importer les vrais)
const InfoItem = {
  props: ['label', 'value'],
  template: `
    <div v-if="value" class="group">
      <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-wide mb-1">{{ label }}</p>
      <p class="text-sm font-bold text-slate-700 dark:text-slate-300">{{ value }}</p>
    </div>
  `
}

const ContactRow = {
  props: ['label', 'value', 'ext', 'icon', 'isPrivate'],
  template: `
    <div v-if="value" class="flex items-center gap-3 group">
      <div class="p-1.5 bg-slate-100 dark:bg-slate-800 rounded-lg">
        <component :is="iconComponent" :size="14" class="text-slate-600 dark:text-slate-400" />
      </div>
      <div class="flex-1">
        <p class="text-xs text-slate-500 dark:text-slate-400 font-bold">{{ label }}</p>
        <p class="text-sm font-bold text-slate-700 dark:text-slate-300">
          {{ value }}
          <span v-if="ext" class="text-xs text-slate-400"> (poste {{ ext }})</span>
        </p>
      </div>
      <Lock v-if="isPrivate" :size="12" class="text-amber-500" />
    </div>
  `,
  computed: {
    iconComponent() {
      return { Phone, Smartphone, Mail }[this.icon] || Phone
    }
  }
}
</script>

<style scoped>
/* Modal Transitions */
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-active .animate-modal-in {
  animation: modal-in 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}

@keyframes modal-in {
  from {
    opacity: 0;
    transform: scale(0.9) translateY(20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

@keyframes shimmer {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

.animate-shimmer {
  animation: shimmer 3s infinite;
}

/* Custom Scrollbar */
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
</style>