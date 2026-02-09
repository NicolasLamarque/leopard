<!-- frontend/src/components/Notaires/NotaireDetailsModal.vue -->
<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
    
    <div class="bg-white dark:bg-slate-900 w-full max-w-5xl rounded-2xl shadow-2xl overflow-hidden animate-in zoom-in-95 duration-200">
      
      <!-- HEADER sobre et professionnel -->
      <div class="relative px-8 py-6 bg-gradient-to-br from-slate-600 via-slate-700 to-slate-600 text-white">
        <div class="flex justify-between items-start">
          <div class="flex items-center gap-5">
            <div class="p-4 bg-white/10 backdrop-blur-sm rounded-xl">
              <Scale :size="32" />
            </div>
            <div>
              <h2 class="text-3xl font-bold mb-1">
                {{ notaire.civilite }} {{ notaire.prenom }} {{ notaire.nom }}
              </h2>
              <div class="flex items-center gap-3 text-sm">
                <span v-if="notaire.secteur_activite" class="px-3 py-1 bg-white/20 rounded-full">
                  {{ notaire.secteur_activite }}
                </span>
                <span v-if="notaire.ville" class="flex items-center gap-1.5">
                  <MapPin :size="14" />
                  {{ notaire.ville }}
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
        
        <!-- Informations professionnelles -->
        <section v-if="hasProfessionalInfo">
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-4 flex items-center gap-2">
            <Award :size="20" class="text-slate-600 dark:text-slate-400" />
            Informations professionnelles
          </h3>

          <div class="grid grid-cols-2 md:grid-cols-3 gap-x-6 gap-y-3">
            <QuickInfo 
              v-if="notaire.civilite"
              label="Civilité" 
              :value="notaire.civilite"
            />
            
            <QuickInfo 
              v-if="notaire.secteur_activite"
              label="Secteur d'activité" 
              :value="notaire.secteur_activite"
              :icon="Award"
              highlight
            />
          </div>
        </section>

        <!-- Coordonnées -->
        <section>
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-4 flex items-center gap-2">
            <Phone :size="20" class="text-slate-600 dark:text-slate-400" />
            Coordonnées
          </h3>

          <div class="space-y-2">
            <ContactRow 
              v-if="notaire.telephone"
              :icon="Phone"
              label="Bureau"
              :value="notaire.telephone"
              :href="`tel:${notaire.telephone}`"
              action-label="Appeler"
              color="slate"
            />

            <ContactRow 
              v-if="notaire.cellulaire"
              :icon="Smartphone"
              label="Cellulaire"
              :value="notaire.cellulaire"
              :href="`tel:${notaire.cellulaire}`"
              action-label="Appeler"
              color="stone"
            />

            <ContactRow 
              v-if="notaire.telecopieur"
              :icon="Printer"
              label="Télécopieur"
              :value="notaire.telecopieur"
              :href="`tel:${notaire.telecopieur}`"
              action-label="Faxer"
              color="zinc"
            />

            <ContactRow 
              v-if="notaire.email"
              :icon="Mail"
              label="Courriel"
              :value="notaire.email"
              :href="`mailto:${notaire.email}`"
              action-label="Écrire"
              color="slate"
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
        <section v-if="notaire.note_fixe" class="p-5 bg-amber-50/50 dark:bg-amber-900/10 rounded-xl border-l-4 border-amber-400/60">
          <div class="flex items-start gap-3">
            <FileText :size="20" class="text-amber-600 dark:text-amber-500 mt-0.5 flex-shrink-0" />
            <div>
              <p class="text-sm font-semibold text-amber-900 dark:text-amber-100 mb-2">Notes</p>
              <p class="text-sm text-slate-700 dark:text-slate-300 whitespace-pre-wrap leading-relaxed">
                {{ notaire.note_fixe }}
              </p>
            </div>
          </div>
        </section>

        <!-- Lien Chambre des notaires -->
        <section class="p-5 bg-slate-50 dark:bg-slate-800/50 rounded-xl border border-slate-200/60 dark:border-slate-700/60">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="p-2 bg-slate-600 rounded-lg">
                <ExternalLink :size="20" class="text-white" />
              </div>
              <div>
                <p class="text-sm font-semibold text-slate-800 dark:text-slate-100">Chambre des notaires du Québec</p>
                <p class="text-xs text-slate-600 dark:text-slate-400">Consulter le répertoire officiel</p>
              </div>
            </div>
            <button 
              @click="openCNQProfile"
              class="px-4 py-2 bg-slate-700 hover:bg-slate-800 text-white rounded-lg text-sm font-medium transition-colors flex items-center gap-2"
            >
              Consulter
              <ExternalLink :size="16" />
            </button>
          </div>
        </section>

      </div>

      <!-- LISTE DES CLIENTS -->
      <div class="px-8 pb-8">
        <NotaireClientList 
          :notaire-id="notaire.id"
          @view-client="handleViewClient"
        />
      </div>

      <!-- FOOTER -->
      <div class="px-8 py-5 border-t border-slate-200 dark:border-slate-700 bg-slate-50 dark:bg-slate-800/50 flex justify-between items-center">
        <div class="text-sm text-slate-600 dark:text-slate-400">
          Ajouté le {{ formatDate(notaire.created_at) }}
        </div>
        
        <div class="flex gap-3">
          <button 
            @click="$emit('edit', notaire)"
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
  X, Scale, Award, Phone, Mail, MapPin, FileText, 
  Edit, Trash2, Smartphone, ExternalLink, Printer
} from 'lucide-vue-next'
import QuickInfo from '../shared/QuickInfo.vue'
import ContactRow from '../shared/ContactRow.vue'
import NotaireClientList from './NotaireClientList.vue'

const props = defineProps({
  notaire: { type: Object, required: true }
})

const emit = defineEmits(['close', 'edit', 'delete', 'view-client'])

const handleViewClient = (clientId) => {
  emit('view-client', clientId)
}

const hasProfessionalInfo = computed(() => {
  return props.notaire.civilite || props.notaire.secteur_activite
})

const hasLocationInfo = computed(() => {
  return props.notaire.adresse || props.notaire.ville
})

const fullAddress = computed(() => {
  const parts = []
  if (props.notaire.adresse) parts.push(props.notaire.adresse)
  if (props.notaire.ville || props.notaire.code_postal) {
    const cityLine = [props.notaire.ville, props.notaire.code_postal].filter(Boolean).join(', ')
    if (cityLine) parts.push(cityLine)
  }
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



const openCNQProfile = () => {
  const cnqUrl = 'https://www.cnq.org/fr/trouver-un-notaire.html'
  window.open(cnqUrl, '_blank')
}



const confirmDelete = () => {
  if (confirm(`⚠️ Supprimer définitivement ${props.notaire.civilite} ${props.notaire.prenom} ${props.notaire.nom} ?`)) {
    emit('delete', props.notaire.id)
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