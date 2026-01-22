<template>
  <div class="bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-900 dark:to-slate-950 rounded-xl shadow-xl border border-slate-200 dark:border-slate-700 overflow-hidden">
    
    <!-- Header avec avatar et infos clés -->
    <div class="bg-gradient-to-r from-slate-700 via-slate-600 to-slate-700 dark:from-slate-800 dark:via-slate-700 dark:to-slate-800 p-6 border-b border-slate-600">
      <div class="flex items-start gap-4">
        <!-- Avatar avec initiales -->
        <div class="relative">
          <div class="w-20 h-20 rounded-lg bg-gradient-to-br from-slate-500 to-slate-600 dark:from-slate-600 dark:to-slate-700 flex items-center justify-center text-white font-semibold text-2xl border border-slate-400 dark:border-slate-500 shadow-lg">
            {{ getInitials() }}
          </div>
          <!-- Badge anniversaire si c'est aujourd'hui ou cette semaine -->
          <div v-if="isBirthdayToday" class="absolute -top-2 -right-2 animate-bounce">
            <div class="w-10 h-10 bg-teal-500 rounded-full flex items-center justify-center shadow-lg border-2 border-white">
              <Cake :size="20" class="text-white" />
            </div>
          </div>
          <div v-else-if="isBirthdayThisWeek" class="absolute -top-2 -right-2">
            <div class="w-8 h-8 bg-teal-400 rounded-full flex items-center justify-center shadow border-2 border-white">
              <Gift :size="14" class="text-white" />
            </div>
          </div>
        </div>

        <!-- Infos principales -->
        <div class="flex-1">
          <div class="flex items-start justify-between">
            <div>
              <h2 class="text-2xl font-semibold text-white mb-1">
                {{ client.prenom }} {{ client.nom }}
              </h2>
              <div class="flex items-center gap-2 flex-wrap">
                <!-- Badge Sexe + Âge -->
                <span class="px-3 py-1 bg-slate-600 rounded text-slate-100 font-medium text-sm border border-slate-500">
                  {{ getSexeBadge() }}{{ age }} ans
                </span>
                
                <!-- Badge Statut -->
                <span 
        :class="[
          'px-2.5 py-1 rounded text-[10px] font-black uppercase tracking-widest flex items-center gap-1.5 border',
          getStatusClass()
        ]"
      >
        <span v-if="client.actif == 1 && client.dcd != 1" class="h-1.5 w-1.5 rounded-full bg-green-500 animate-pulse"></span>
        {{ getStatusText() }}
      </span>

                <!-- Badge Anniversaire -->
                <span v-if="isBirthdayToday" class="px-3 py-1 bg-teal-500 text-white rounded text-xs font-semibold animate-pulse">
                  🎂 Anniversaire aujourd'hui !
                </span>
                <span v-else-if="isBirthdayThisWeek" class="px-3 py-1 bg-teal-600 text-white rounded text-xs font-medium">
                  🎁 Anniversaire dans {{ daysUntilBirthday }} jour(s)
                </span>
              </div>
            </div>

            <!-- Bouton Export PDF -->
            <button
              @click="exportToPDF"
              class="px-4 py-2 bg-slate-600 hover:bg-slate-500 rounded transition-colors flex items-center gap-2 text-white font-medium border border-slate-500"
              title="Exporter en PDF"
            >
              <FileDown :size="18" />
              <span class="hidden md:inline">PDF</span>
            </button>
          </div>

          <!-- N° Dossier -->
          <div class="mt-3 flex items-center gap-2">
            <Shield :size="16" class="text-slate-300" />
            <span class="text-slate-300 text-sm">N° Dossier:</span>
            <span class="font-mono text-white font-medium">{{ client.no_dossier_leopard || 'N/A' }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Contenu des sections -->
    <div class="p-6 space-y-4">
      
      <!-- Section Contact d'urgence -->
      <div v-if="urgenceContact" class="bg-red-50 dark:bg-red-950/30 border border-red-200 dark:border-red-900 rounded-lg p-4">
        <div class="flex items-center gap-2 mb-3">
          <div class="p-2 bg-red-100 dark:bg-red-900/50 rounded">
            <AlertCircle :size="20" class="text-red-700 dark:text-red-400" />
          </div>
          <h3 class="font-semibold text-red-900 dark:text-red-100">Contact d'urgence</h3>
        </div>
        <div class="space-y-2">
          <p class="text-sm">
            <span class="font-medium text-red-900 dark:text-red-100">
              {{ urgenceContact.prenom }} {{ urgenceContact.nom }}
            </span>
            <span v-if="urgenceContact.lien_familial" class="text-xs text-red-700 dark:text-red-300 ml-2">
              ({{ urgenceContact.lien_familial }})
            </span>
          </p>
          <div class="flex flex-wrap gap-3 text-sm">
            <a v-if="urgenceContact.telephone" 
               :href="`tel:${urgenceContact.telephone}`"
               class="flex items-center gap-1 text-red-700 dark:text-red-300 hover:underline">
              <Phone :size="14" />
              {{ formatPhone(urgenceContact.telephone) }}
            </a>
            <a v-if="urgenceContact.cellulaire" 
               :href="`tel:${urgenceContact.cellulaire}`"
               class="flex items-center gap-1 text-red-700 dark:text-red-300 hover:underline">
              <Smartphone :size="14" />
              {{ formatPhone(urgenceContact.cellulaire) }}
            </a>
          </div>
        </div>
      </div>

      <!-- Section Informations médicales -->
      <div class="bg-gradient-to-br from-white to-slate-50 dark:from-slate-800/50 dark:to-slate-900/30 border border-slate-200 dark:border-slate-700 rounded-lg p-4 shadow-sm">
        <div class="flex items-center gap-2 mb-3">
          <div class="p-2 bg-gradient-to-br from-teal-100 to-teal-50 dark:from-teal-900/40 dark:to-teal-800/20 rounded">
            <Heart :size="20" class="text-teal-700 dark:text-teal-400" />
          </div>
          <h3 class="font-semibold text-slate-900 dark:text-slate-100">Informations médicales</h3>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-3 text-sm">
          <!-- Médecin de famille -->
          <div v-if="client.medecin_famille_No_Licence" class="flex items-start gap-2">
            <Stethoscope :size="16" class="text-slate-400 mt-0.5 flex-shrink-0" />
            <div>
              <p class="text-xs text-slate-500 dark:text-slate-400">Médecin de famille</p>
              <div class="flex items-center gap-2">
                <p class="font-medium text-slate-900 dark:text-slate-100">Dr. {{ getMedecinName() }}</p>
                <button 
                  @click="viewMedecinDetails"
                  class="p-1 hover:bg-slate-100 dark:hover:bg-slate-700 rounded transition-colors"
                  title="Voir la fiche"
                >
                  <Eye :size="14" class="text-teal-600 dark:text-teal-400" />
                </button>
              </div>
              <p class="text-xs text-slate-600 dark:text-slate-400">Licence: {{ client.medecin_famille_No_Licence }}</p>
            </div>
          </div>

          <!-- N° RAMQ -->
          <div v-if="client.numero_assurance_maladie" class="flex items-start gap-2">
            <CreditCard :size="16" class="text-slate-400 mt-0.5 flex-shrink-0" />
            <div>
              <p class="text-xs text-slate-500 dark:text-slate-400">N° RAMQ</p>
              <p class="font-mono font-medium text-slate-900 dark:text-slate-100">{{ client.numero_assurance_maladie }}</p>
            </div>
          </div>

         
          <!-- Dossiers cliniques -->
          <div v-if="client.no_hcm || client.no_chaur" class="flex items-start gap-2">
            <FileText :size="16" class="text-slate-400 mt-0.5 flex-shrink-0" />
            <div>
              <p class="text-xs text-slate-500 dark:text-slate-400">Dossiers</p>
              <div class="space-y-0.5">
                <p v-if="client.no_hcm" class="text-xs text-slate-700 dark:text-slate-300">HCM: {{ client.no_hcm }}</p>
                <p v-if="client.no_chaur" class="text-xs text-slate-700 dark:text-slate-300">CHAUR: {{ client.no_chaur }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Section Localisation en deux colonnes-->
       
<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
  
  <div v-if="client.adresse || client.ville" 
       class="bg-gradient-to-br from-white to-slate-50 dark:from-slate-800/50 dark:to-slate-900/30 border border-slate-200 dark:border-slate-700 rounded-lg p-4 shadow-sm flex flex-col justify-between">
    <div>
      <div class="flex items-center gap-2 mb-3">
        <div class="p-2 bg-stone-100 dark:bg-stone-900/40 rounded">
          <MapPin :size="20" class="text-stone-700 dark:text-stone-400" />
        </div>
        <h3 class="font-semibold text-slate-900 dark:text-slate-100 italic">Localisation</h3>
      </div>
      <div class="text-sm text-slate-700 dark:text-slate-300">
        <p class="font-medium">{{ client.adresse }}</p>
        <p class="text-slate-500">{{ client.ville }}, {{ client.code_postal }}</p>
      </div>
    </div>
    <button @click="openGoogleMaps" class="mt-4 w-full flex items-center justify-center gap-2 px-3 py-1.5 bg-white dark:bg-slate-800 hover:bg-slate-50 rounded border border-slate-200 text-[10px] font-black tracking-widest transition-colors">
      <Map :size="14" /> GOOGLE MAPS
    </button>
  </div>

  <div v-if="getEtablissement()" 
       class="bg-gradient-to-br from-white to-slate-50 dark:from-slate-800/50 dark:to-slate-900/30 border border-slate-200 dark:border-slate-700 rounded-lg p-4 shadow-sm flex flex-col justify-between">
    <div>
      <div class="flex items-center gap-2 mb-3">
        <div class="p-2 bg-blue-100 dark:bg-blue-900/40 rounded">
          <Building2 :size="20" class="text-blue-700 dark:text-blue-400" />
        </div>
        <h3 class="font-semibold text-slate-900 dark:text-slate-100 italic">Hébergement</h3>
      </div>
      <div class="space-y-1">
        <p class="text-[10px] text-slate-500 uppercase font-bold tracking-tighter">Établissement</p>
        <p class="text-sm font-medium text-slate-900 dark:text-slate-100 leading-snug">
          {{ getEtablissement() }}
        </p>
      </div>
    </div>
    
    <div class="mt-4 pt-3 border-t border-slate-200 dark:border-slate-700 flex items-center justify-between">
      

      <span v-if="client.date_maj" class="text-[10px] text-slate-400 italic">
        MàJ: {{ client.date_maj }}
      </span>
    </div>
  </div>

</div>

      <!-- Section Coordonnées -->
      <div class="bg-gradient-to-br from-white to-slate-50 dark:from-slate-800/50 dark:to-slate-900/30 border border-slate-200 dark:border-slate-700 rounded-lg p-4 shadow-sm">
        <div class="flex items-center gap-2 mb-3">
          <div class="p-2 bg-gradient-to-br from-teal-100 to-teal-50 dark:from-teal-900/40 dark:to-teal-800/20 rounded">
            <Phone :size="20" class="text-teal-700 dark:text-teal-400" />
          </div>
          <h3 class="font-semibold text-slate-900 dark:text-slate-100">Coordonnées</h3>
        </div>

        <div class="flex flex-wrap gap-3 text-sm">
          <a v-if="client.telephone" 
             :href="`tel:${client.telephone}`"
             class="flex items-center gap-2 px-3 py-2 bg-white dark:bg-slate-700 hover:bg-slate-100 dark:hover:bg-slate-600 rounded transition-colors border border-slate-200 dark:border-slate-600">
            <Phone :size="14" class="text-slate-500 dark:text-slate-400" />
            <span class="text-slate-900 dark:text-slate-100">{{ formatPhone(client.telephone) }}</span>
          </a>
          <a v-if="client.cellulaire" 
             :href="`tel:${client.cellulaire}`"
             class="flex items-center gap-2 px-3 py-2 bg-white dark:bg-slate-700 hover:bg-slate-100 dark:hover:bg-slate-600 rounded transition-colors border border-slate-200 dark:border-slate-600">
            <Smartphone :size="14" class="text-slate-500 dark:text-slate-400" />
            <span class="text-slate-900 dark:text-slate-100">{{ formatPhone(client.cellulaire) }}</span>
          </a>
          <a v-if="client.email" 
             :href="`mailto:${client.email}`"
             class="flex items-center gap-2 px-3 py-2 bg-white dark:bg-slate-700 hover:bg-slate-100 dark:hover:bg-slate-600 rounded transition-colors border border-slate-200 dark:border-slate-600">
            <Mail :size="14" class="text-slate-500 dark:text-slate-400" />
            <span class="text-slate-900 dark:text-slate-100">{{ client.email }}</span>
          </a>
        </div>
      </div>

      <!-- Notes importantes -->
      <div v-if="client.note_fixe" class="bg-amber-50 dark:bg-amber-950/30 border border-amber-200 dark:border-amber-900 rounded-lg p-4">
        <div class="flex items-center gap-2 mb-2">
          <AlertTriangle :size="18" class="text-amber-700 dark:text-amber-400" />
          <h3 class="font-semibold text-amber-900 dark:text-amber-100">Notes importantes</h3>
        </div>
        <p class="text-sm text-amber-800 dark:text-amber-200 whitespace-pre-wrap">{{ client.note_fixe }}</p>
      </div>

    </div>

    <!-- Footer avec timestamp -->
    <div class="px-6 py-3 bg-slate-50 dark:bg-slate-800/50 border-t border-slate-200 dark:border-slate-700">
      <p class="text-xs text-slate-500 dark:text-slate-400 text-center">
        Dernière mise à jour: {{ new Date(client.created_at).toLocaleDateString('fr-CA') }}
      </p>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  Shield, AlertCircle, Heart, MapPin, Phone, Mail, FileDown,
  Cake, Gift, Eye, Stethoscope, CreditCard, Building2, FileText,
  Map, Smartphone, AlertTriangle
} from 'lucide-vue-next'
import { GetAllContactsByClientID } from '../../../wailsjs/go/main/App'

const props = defineProps({
  client: { type: Object, required: true }
})

const emit = defineEmits(['view-medecin'])

const contacts = ref([])

// Calcul de l'âge
const age = computed(() => {
  if (!props.client.date_naissance) return '?'
  
  const birthDate = new Date(props.client.date_naissance)
  const today = new Date()
  let age = today.getFullYear() - birthDate.getFullYear()
  const monthDiff = today.getMonth() - birthDate.getMonth()
  
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
    age--
  }
  
  return age
})

// Anniversaire aujourd'hui
const isBirthdayToday = computed(() => {
  if (!props.client.date_naissance) return false
  
  const birthDate = new Date(props.client.date_naissance)
  const today = new Date()
  
  return birthDate.getMonth() === today.getMonth() && 
         birthDate.getDate() === today.getDate()
})

// Anniversaire cette semaine
const isBirthdayThisWeek = computed(() => {
  if (!props.client.date_naissance || isBirthdayToday.value) return false
  
  const birthDate = new Date(props.client.date_naissance)
  const today = new Date()
  
  // Anniversaire cette année
  const thisYearBirthday = new Date(today.getFullYear(), birthDate.getMonth(), birthDate.getDate())
  
  const diffTime = thisYearBirthday - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  return diffDays > 0 && diffDays <= 7
})

const daysUntilBirthday = computed(() => {
  if (!props.client.date_naissance) return 0
  
  const birthDate = new Date(props.client.date_naissance)
  const today = new Date()
  const thisYearBirthday = new Date(today.getFullYear(), birthDate.getMonth(), birthDate.getDate())
  
  const diffTime = thisYearBirthday - today
  return Math.ceil(diffTime / (1000 * 60 * 60 * 24))
})

// Formatte un numéro de téléphone au format lisible
const formatPhone = (phone) => {
  if (!phone) return '';
  
  // Enlever tous les caractères non-numériques
  const cleaned = phone.replace(/\D/g, '');
  
  // Format nord-américain (10 chiffres)
  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`;
  }
  
  // Si 11 chiffres (avec le 1 au début)
  if (cleaned.length === 11 && cleaned[0] === '1') {
    return `+1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`;
  }
  
  // Sinon retourner le numéro tel quel
  return phone;
};
// Contact d'urgence
const urgenceContact = computed(() => {
  return contacts.value.find(c => c.contact_urgence === 1) || null
})

const getInitials = () => {
  return `${props.client.prenom[0]}${props.client.nom[0]}`.toUpperCase()
}

const getSexeBadge = () => {
  // Si tu ajoutes un champ sexe dans ta table
  const sexe = props.client.sexe || 'F' // Par défaut F
  return sexe === 'M' ? 'H' : 'F'
}


const getStatusText = () => {
  // 1. Priorité au décès
  if (props.client.dcd == 1) return '✝ Ce client est Décédé le : ' + (props.client.date_deces || 'N/A')
  
  // 2. Vérification du statut actif (attention à la minuscule 'actif')
  if (props.client.actif == 1) return '✓ Dossier Actif'
  
  // 3. Sinon archivé
  return '⊗ Archivé'
}

// Pour la couleur dynamique du badge
const getStatusClass = () => {
  if (props.client.dcd == 1) return 'bg-gray-100 text-gray-700 border-gray-200'
  if (props.client.actif == 1) return 'bg-teal-100/60 text-teal-700 border-green-200 shadow-sm'
  return 'bg-red-50 text-red-600 border-red-100'
}
const getEtablissement = () => {
  if (props.client.rpa_id) return `RPA #${props.client.rpa_id}`
  if (props.client.chsld_id) return `CHSLD #${props.client.chsld_id}`
  if (props.client.ri_id) return `RI #${props.client.ri_id}`
  return null
}

const getMedecinName = () => {
  // À remplacer par une vraie requête vers la table médecins
  return 'Non spécifié'
}

const openGoogleMaps = () => {
  const address = `${props.client.adresse}, ${props.client.ville}, ${props.client.code_postal}`
  window.open(`https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(address)}`, '_blank')
}

const viewMedecinDetails = () => {
  emit('view-medecin', props.client.medecin_famille_No_Licence)
}

const exportToPDF = () => {
  // À implémenter avec une lib comme jsPDF ou html2pdf
  alert('Export PDF à implémenter')
}

const loadContacts = async () => {
  try {
    const result = await GetAllContactsByClientID(props.client.id)
    contacts.value = result || []
  } catch (error) {
    console.error('Erreur chargement contacts:', error)
  }
}

onMounted(() => {
  loadContacts()
})
</script>

<style scoped>
@keyframes bounce {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
}

.animate-bounce {
  animation: bounce 1s infinite;
}
</style>