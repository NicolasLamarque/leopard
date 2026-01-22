<!-- frontend/src/components/Medecins/MedecinDetailsModal.vue -->
<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/60 backdrop-blur-sm">
    
    <div class="bg-white dark:bg-gray-900 w-full max-w-5xl rounded-2xl shadow-2xl overflow-hidden animate-in zoom-in-95 duration-200">
      
      <!-- HEADER ÉLÉGANT -->
      <div class="relative px-8 py-6 bg-gradient-to-r from-teal-700 to-slate-700 text-white">
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
                  :class="medecin.actif ? 'bg-green-500' : 'bg-gray-500'"
                  class="px-3 py-1 rounded-full font-semibold flex items-center gap-1.5"
                >
                  <span class="w-1.5 h-1.5 bg-white rounded-full"></span>
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
      <div class="p-8 max-h-[70vh] overflow-y-auto space-y-6">
        
        <!-- Informations professionnelles -->
        <div>
          <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
            <Award :size="20" class="text-teal-700" />
            Informations professionnelles
          </h3>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div v-if="medecin.civilite" class="p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
              <p class="text-xs text-gray-500 dark:text-gray-400 mb-1">Civilité</p>
              <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.civilite }}</p>
            </div>

            <div class="p-4 bg-blue-50 dark:bg-blue-900/20 rounded-lg border border-blue-200 dark:border-blue-800">
              <p class="text-xs text-blue-600 dark:text-blue-400 mb-1">Spécialisation</p>
              <p class="font-semibold text-blue-900 dark:text-blue-100">
                {{ medecin.specialisation || 'Médecine générale' }}
              </p>
            </div>

            <div v-if="medecin.statut" class="p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
              <p class="text-xs text-gray-500 dark:text-gray-400 mb-1">Statut</p>
              <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.statut }}</p>
            </div>

            <div v-if="medecin.service" class="p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
              <p class="text-xs text-gray-500 dark:text-gray-400 mb-1">Service</p>
              <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.service }}</p>
            </div>

            <div v-if="medecin.departement" class="p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
              <p class="text-xs text-gray-500 dark:text-gray-400 mb-1">Département</p>
              <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.departement }}</p>
            </div>

            <div v-if="medecin.rls" class="p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
              <p class="text-xs text-gray-500 dark:text-gray-400 mb-1">RLS</p>
              <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.rls }}</p>
            </div>
          </div>
        </div>

        <!-- Coordonnées -->
        <div>
          <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
            <Phone :size="20" class="text-teal-700" />
            Coordonnées
          </h3>

          <div class="space-y-3">
            <div v-if="medecin.telephone" class="flex items-center gap-3 p-4 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
              <div class="p-2 bg-blue-500 rounded-lg">
                <Phone :size="18" class="text-white" />
              </div>
              <div class="flex-1">
                <p class="text-xs text-blue-600 dark:text-blue-400">Téléphone bureau</p>
                <p class="font-semibold text-gray-900 dark:text-white">
                  {{ medecin.telephone }}{{ medecin.extension ? ' poste ' + medecin.extension : '' }}
                </p>
              </div>
              <a 
                :href="`tel:${medecin.telephone}`"
                class="px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm font-medium transition-colors"
              >
                Appeler
              </a>
            </div>

            <div v-if="medecin.cellulaire" class="flex items-center gap-3 p-4 bg-green-50 dark:bg-green-900/20 rounded-lg">
              <div class="p-2 bg-slate-500 rounded-lg">
                <Smartphone :size="18" class="text-white" />
              </div>
              <div class="flex-1">
                <p class="text-xs text-slate-600 dark:text-slate-400">Cellulaire</p>
                <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.cellulaire }}</p>
              </div>
              <a 
                :href="`tel:${medecin.cellulaire}`"
                class="px-3 py-1.5 bg-slate-600 hover:bg-slate-700 text-white rounded-lg text-sm font-medium transition-colors"
              >
                Appeler
              </a>
            </div>

            <div v-if="medecin.email" class="flex items-center gap-3 p-4 bg-slate-50 dark:bg-teal-900/20 rounded-lg">
              <div class="p-2 bg-slate-500 rounded-lg">
                <Mail :size="18" class="text-white" />
              </div>
              <div class="flex-1">
                <p class="text-xs text-slate-600 dark:text-slate-400">Courriel</p>
                <p class="font-semibold text-gray-900 dark:text-white">{{ medecin.email }}</p>
              </div>
              <a 
                :href="`mailto:${medecin.email}`"
                class="px-3 py-1.5 bg-slate-900 hover:bg-slate-500 text-white rounded-lg text-sm font-medium transition-colors"
              >
                Écrire
              </a>
            </div>
          </div>
        </div>

        <!-- Localisation -->
        <div v-if="hasLocationInfo">
          <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
            <MapPin :size="20" class="text-slate-500" />
            Localisation
          </h3>

          <div class="space-y-3">
            <div v-if="medecin.installation_principale" class="p-4 bg-orange-50 dark:bg-orange-900/20 rounded-lg border border-orange-200 dark:border-orange-800">
              <p class="text-xs text-orange-600 dark:text-orange-400 mb-1">Installation principale</p>
              <p class="font-semibold text-orange-900 dark:text-orange-100">{{ medecin.installation_principale }}</p>
            </div>

            <div v-if="fullAddress" class="p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
              <p class="text-xs text-gray-500 dark:text-gray-400 mb-1">Adresse</p>
              <p class="font-semibold text-gray-900 dark:text-white whitespace-pre-line">{{ fullAddress }}</p>
            </div>
          </div>
        </div>

        <!-- Notes permanentes -->
        <div v-if="medecin.noteFixe" class="p-5 bg-amber-50 dark:bg-amber-900/20 rounded-lg border-l-4 border-amber-500">
          <div class="flex items-start gap-3">
            <FileText :size="20" class="text-amber-600 mt-0.5" />
            <div>
              <p class="text-sm font-semibold text-amber-900 dark:text-amber-100 mb-2">Notes permanentes</p>
              <p class="text-gray-700 dark:text-gray-300 whitespace-pre-wrap">{{ medecin.noteFixe }}</p>
            </div>
          </div>
        </div>

        <!-- Lien CMQ -->
        <div class="p-4 bg-blue-50 dark:bg-slate-900/20 rounded-lg border border-blue-200 dark:border-teal-700">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <ExternalLink :size="20" class="text-slate-600" />
              <div>
                <p class="text-sm font-semibold text-slate-800 dark:text-slate-100">Profil au Collège des médecins</p>
                <p class="text-xs text-teal-600 dark:text-slate-400">Consulter le profil officiel sur cmq.org</p>
              </div>
            </div>
            <button 
              @click="openCMQProfile"
              class="px-4 py-2 bg-teal-600 hover:bg-teal-700 text-white rounded-lg text-sm font-medium transition-colors flex items-center gap-2"
            >
              Consulter
              <ExternalLink :size="16" />
            </button>
          </div>
        </div>

      </div>

      <!-- FOOTER -->
      <div class="px-8 py-5 border-t border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-800/50 flex justify-between items-center">
        <div class="text-sm text-gray-500 dark:text-gray-400">
          Ajouté le {{ formatDate(medecin.createdAt) }}
        </div>
        
        <div class="flex gap-3">
          <button 
            @click="$emit('edit', medecin)"
            class="px-5 py-2.5 bg-teal-700 hover:bg-teal-800 text-white rounded-lg font-semibold transition-colors flex items-center gap-2"
          >
            <Edit :size="18" />
            Modifier
          </button>
          
          <button 
            @click="confirmDelete"
            class="px-5 py-2.5 bg-amber-700 hover:bg-amber-800 text-white rounded-lg font-semibold transition-colors flex items-center gap-2"
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

const props = defineProps({
  medecin: { type: Object, required: true }
})

const emit = defineEmits(['close', 'edit', 'delete'])

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