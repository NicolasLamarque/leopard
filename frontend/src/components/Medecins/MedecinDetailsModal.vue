<!-- frontend/src/components/Medecins/MedecinDetailsModal.vue -->
<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm">
    
    <div class="bg-white dark:bg-slate-900 w-full max-w-5xl rounded-2xl shadow-2xl overflow-hidden animate-in zoom-in-95 duration-200">
      
      <!-- HEADER ÉLÉGANT -->
      <div class="relative px-8 py-6 bg-gradient-to-br from-slate-700 via-slate-600 to-stone-700 text-white">
        <div class="flex justify-between items-start">
          <div class="flex items-center gap-5">
            <div class="p-4 bg-white/10 backdrop-blur-sm rounded-xl">
              <Stethoscope :size="32" />
            </div>
            <div>
              <h2 class="text-3xl font-bold mb-1">
                {{ medecin.nomComplet }}
              </h2>
              <div class="flex items-center gap-3 text-sm">
                <span class="px-3 py-1 bg-white/20 rounded-full font-mono">
                  {{ medecin.licence }}
                </span>
                <span 
                  :class="medecin.actif ? 'bg-emerald-500/90' : 'bg-slate-500/90'"
                  class="px-3 py-1 rounded-full font-semibold flex items-center gap-1.5"
                >
                  <span class="w-1.5 h-1.5 bg-white rounded-full animate-pulse"></span>
                  {{ medecin.actif ? 'Actif' : 'Inactif' }}
                </span>
              </div>
            </div>
          </div>
          
          <button 
            @click="$emit('close')" 
            class="p-2 hover:bg-white/10 rounded-lg transition-colors"
          >
            <X :size="24" />
          </button>
        </div>
      </div>

      <!-- CONTENU -->
      <div class="p-8 max-h-[70vh] overflow-y-auto space-y-8">
        
        <!-- Informations professionnelles - VERSION COMPACTE -->
        <section>
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-4 flex items-center gap-2">
            <Award :size="20" class="text-slate-600 dark:text-slate-400" />
            Informations professionnelles
          </h3>

          <div class="space-y-3">
            <!-- Quick Info Grid -->
            <div class="grid grid-cols-2 md:grid-cols-4 gap-x-6 gap-y-3">
              <QuickInfo 
                v-if="medecin.civilite"
                label="Civilité" 
                :value="medecin.civilite"
              />
              
              <QuickInfo 
                label="Spécialisation" 
                :value="medecin.specialisation || 'Médecine générale'"
                :icon="Award"
                highlight
              />
              
              <QuickInfo 
                v-if="medecin.statut"
                label="Statut" 
                :value="medecin.statut"
              />
              
              <QuickInfo 
                v-if="medecin.service"
                label="Service" 
                :value="medecin.service"
              />
              
              <QuickInfo 
                v-if="medecin.departement"
                label="Département" 
                :value="medecin.departement"
              />
              
              <QuickInfo 
                v-if="medecin.rls"
                label="RLS" 
                :value="medecin.rls"
              />
            </div>
          </div>
        </section>

        <!-- Coordonnées - VERSION ÉPURÉE -->
        <section>
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-4 flex items-center gap-2">
            <Phone :size="20" class="text-slate-600 dark:text-slate-400" />
            Coordonnées
          </h3>

          <div class="space-y-2">
            <!-- Téléphone bureau -->
            <ContactRow 
              v-if="medecin.telephone"
              :icon="Phone"
              label="Bureau"
              :value="medecin.telephone + (medecin.extension ? ' poste ' + medecin.extension : '')"
              :href="`tel:${medecin.telephone}`"
              action-label="Appeler"
              color="slate"
            />

            <!-- Cellulaire -->
            <ContactRow 
              v-if="medecin.cellulaire"
              :icon="Smartphone"
              label="Cellulaire"
              :value="medecin.cellulaire"
              :href="`tel:${medecin.cellulaire}`"
              action-label="Appeler"
              color="stone"
            />

            <!-- Email -->
            <ContactRow 
              v-if="medecin.email"
              :icon="Mail"
              label="Courriel"
              :value="medecin.email"
              :href="`mailto:${medecin.email}`"
              action-label="Écrire"
              color="zinc"
            />
          </div>
        </section>

        <!-- Localisation -->
        <section v-if="hasLocationInfo">
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-4 flex items-center gap-2">
            <MapPin :size="20" class="text-slate-600 dark:text-slate-400" />
            Localisation
          </h3>

          <div class="space-y-3">
            <QuickInfo 
              v-if="medecin.installation_principale"
              label="Installation principale" 
              :value="medecin.installation_principale"
              :icon="Building2"
              highlight
            />

            <div v-if="fullAddress" class="pl-0">
              <p class="text-xs font-medium text-slate-500 dark:text-slate-400 mb-1 flex items-center gap-1.5">
                <MapPin :size="14" class="text-slate-400 dark:text-slate-500" />
                Adresse
              </p>
              <p class="text-sm text-slate-700 dark:text-slate-300 whitespace-pre-line leading-relaxed">
                {{ fullAddress }}
              </p>
            </div>
          </div>
        </section>

        <!-- Notes permanentes -->
        <section v-if="medecin.noteFixe" class="p-5 bg-amber-50/50 dark:bg-amber-900/10 rounded-xl border-l-4 border-amber-400/60">
          <div class="flex items-start gap-3">
            <FileText :size="20" class="text-amber-600 dark:text-amber-500 mt-0.5 flex-shrink-0" />
            <div>
              <p class="text-sm font-semibold text-amber-900 dark:text-amber-100 mb-2">Notes permanentes</p>
              <p class="text-sm text-slate-700 dark:text-slate-300 whitespace-pre-wrap leading-relaxed">
                {{ medecin.noteFixe }}
              </p>
            </div>
          </div>
        </section>

        <!-- Lien CMQ -->
        <section class="p-5 bg-slate-50 dark:bg-slate-800/50 rounded-xl border border-slate-200/60 dark:border-slate-700/60">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="p-2 bg-slate-600 rounded-lg">
                <ExternalLink :size="20" class="text-white" />
              </div>
              <div>
                <p class="text-sm font-semibold text-slate-800 dark:text-slate-100">Profil au Collège des médecins</p>
                <p class="text-xs text-slate-600 dark:text-slate-400">Consulter le profil officiel sur cmq.org</p>
              </div>
            </div>
            <button 
              @click="openCMQProfile"
              class="px-4 py-2 bg-slate-700 hover:bg-slate-800 text-white rounded-lg text-sm font-medium transition-colors flex items-center gap-2"
            >
              Consulter
              <ExternalLink :size="16" />
            </button>
          </div>
        </section>

      </div>

      <!-- 🆕 LISTE DES PATIENTS EN COMMUN -->
      <div class="px-8 pb-8">
        <MedecinClientsList 
          :medecin-licence="medecin.licence"
          @view-client="handleViewClient"
        />
      </div>

      <!-- FOOTER -->
      <div class="px-8 py-5 border-t border-slate-200 dark:border-slate-700 bg-slate-50 dark:bg-slate-800/50 flex justify-between items-center">
        <div class="text-sm text-slate-600 dark:text-slate-400">
          Ajouté le {{ formatDate(medecin.createdAt) }}
        </div>
        
        <div class="flex gap-3">
          <button 
            @click="$emit('edit', medecin)"
            class="px-5 py-2.5 bg-slate-700 hover:bg-slate-800 text-white rounded-lg font-semibold transition-colors flex items-center gap-2"
          >
            <Edit :size="18" />
            Modifier
          </button>
          
          <button 
            @click="confirmDelete"
            class="px-5 py-2.5 bg-rose-600 hover:bg-rose-700 text-white rounded-lg font-semibold transition-colors flex items-center gap-2"
          >
            <Trash2 :size="18" />
            Supprimer
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { 
  X, Stethoscope, Award, Phone, Mail, MapPin, FileText, 
  Edit, Trash2, Smartphone, ExternalLink
} from 'lucide-vue-next'
import QuickInfo from './QuickInfosMed.vue'
import ContactRow from './../shared/ContactRow.vue'
import MedecinClientsList from './MedecinClientList.vue'

const props = defineProps({
  medecin: { type: Object, required: true }
})

const emit = defineEmits(['close', 'edit', 'delete', 'view-client'])

const handleViewClient = (clientId) => {
  emit('view-client', clientId)
}

const hasLocationInfo = computed(() => {
  return props.medecin.installation_principale || 
         props.medecin.adresse || 
         props.medecin.ville
})

const fullAddress = computed(() => {
  const parts = []
  if (props.medecin.adresse) parts.push(props.medecin.adresse)
  if (props.medecin.ville || props.medecin.codePostal) {
    const cityLine = [props.medecin.ville, props.medecin.codePostal].filter(Boolean).join(', ')
    if (cityLine) parts.push(cityLine)
  }
  if (props.medecin.pays && props.medecin.pays !== 'Canada') parts.push(props.medecin.pays)
  return parts.length > 0 ? parts.join('\n') : null
})

const formatDate = (dateString) => {
  if (!dateString) return 'Date inconnue'
  const date = new Date(dateString)
  return date.toLocaleDateString('fr-CA', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
}

const openCMQProfile = () => {
  const cmqUrl = `https://www.cmq.org/fr/bottin/medecins?number=${props.medecin.licence}&lastname=&firstname=&specialty=0&city=&unlisted=false`
  window.open(cmqUrl, '_blank')
}

const confirmDelete = () => {
  if (confirm(`⚠️ Supprimer définitivement ${props.medecin.nomComplet} ?`)) {
    emit('delete', props.medecin.id)
  }
}
</script>

<style scoped>
.animate-in {
  animation: fadeIn 0.2s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
</style>