<template>
  <div class="bg-white dark:bg-slate-900 rounded-xl shadow-lg border border-slate-200 dark:border-slate-700 overflow-hidden">
    
    <!-- Header Compact -->
    <div class="bg-gradient-to-r from-slate-700 to-slate-600 dark:from-slate-800 dark:to-slate-700 px-6 py-4">
      <div class="flex items-center justify-between">
        <!-- Identité -->
        <div class="flex items-center gap-4">
          <div class="relative">
            <div class="w-16 h-16 rounded-lg bg-gradient-to-br from-slate-500 to-slate-600 flex items-center justify-center text-white font-bold text-xl shadow-md">
              {{ getInitials() }}
            </div>
            <!-- Badge anniversaire -->
            <div v-if="isBirthdayToday" class="absolute -top-1 -right-1">
              <div class="w-7 h-7 bg-teal-500 rounded-full flex items-center justify-center shadow-lg">
                <Cake :size="14" class="text-white" />
              </div>
            </div>
          </div>
          
          <div>
            <h2 class="text-xl font-bold text-white">
              {{ client.prenom }} {{ client.nom }}
            </h2>
            <div class="flex items-center gap-2 mt-1">
              <span class="text-slate-200 text-sm">{{ getSexeBadge() }}{{ age }} ans</span>
              <span class="text-slate-400">•</span>
              <span :class="getStatusClass()" class="text-xs font-semibold px-2 py-0.5 rounded">
                {{ getStatusText() }}
              </span>
            </div>
          </div>
        </div>

        <!-- Actions rapides -->
        <div class="flex items-center gap-2">
          <button
            v-if="isBirthdayToday"
            class="px-3 py-1.5 bg-teal-500/20 text-teal-200 rounded-lg text-xs font-medium flex items-center gap-1"
          >
            🎂 Anniversaire !
          </button>
          <button
            @click="exportToPDF"
            class="p-2 bg-white/10 hover:bg-white/20 rounded-lg transition-colors"
            title="Exporter en PDF"
          >
            <FileDown :size="18" class="text-white" />
          </button>
        </div>
      </div>
    </div>

    <!-- Numéro Leopard -->
    <div class="px-6 py-2 bg-slate-50 dark:bg-slate-800/50 border-b border-slate-200 dark:border-slate-700">
      <div class="flex items-center gap-2 text-sm">
        <Shield :size="14" class="text-slate-400" />
        <span class="text-slate-600 dark:text-slate-400">Dossier:</span>
        <span class="font-mono font-semibold text-slate-900 dark:text-slate-100">
          {{ client.no_dossier_leopard || 'Non généré' }}
        </span>
      </div>
    </div>

    <!-- Contenu principal - Une seule carte fluide -->
    <div class="p-6">
      
      <!-- Contact d'urgence prioritaire -->
      <div v-if="urgenceContact" class="mb-6 p-4 bg-red-50 dark:bg-red-950/20 border-l-4 border-red-500 rounded">
        <div class="flex items-start justify-between">
          <div class="flex items-start gap-3 flex-1">
            <AlertCircle :size="18" class="text-red-600 dark:text-red-400 mt-0.5 flex-shrink-0" />
            <div class="flex-1">
              <p class="text-xs font-bold text-red-900 dark:text-red-100 uppercase tracking-wide mb-1">
                Contact d'urgence
              </p>
              <p class="font-semibold text-red-900 dark:text-red-100">
                {{ urgenceContact.prenom }} {{ urgenceContact.nom }}
              </p>
              <p v-if="urgenceContact.lien_familial" class="text-sm text-red-700 dark:text-red-300">
                {{ urgenceContact.lien_familial }}
              </p>
            </div>
          </div>
          <div class="flex flex-col gap-1 ml-4">
            <a v-if="urgenceContact.telephone" 
               :href="`tel:${urgenceContact.telephone}`"
               class="text-sm text-red-700 dark:text-red-300 hover:underline flex items-center gap-1 whitespace-nowrap">
              <Phone :size="12" />
              {{ formatPhone(urgenceContact.telephone) }}
            </a>
            <a v-if="urgenceContact.cellulaire" 
               :href="`tel:${urgenceContact.cellulaire}`"
               class="text-sm text-red-700 dark:text-red-300 hover:underline flex items-center gap-1 whitespace-nowrap">
              <Smartphone :size="12" />
              {{ formatPhone(urgenceContact.cellulaire) }}
            </a>
          </div>
        </div>
      </div>

      <!-- Grid 3 colonnes - Sans bordures individuelles -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <!-- Colonne 1: Médical -->
        <div class="space-y-4">
          <h3 class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2 mb-3">
            <Heart :size="14" class="text-teal-600 dark:text-teal-400" />
            Médical
          </h3>

          <!-- Médecin de famille -->
          <div v-if="client.medecin_famille_No_Licence || medecinInfo">
            <div class="flex items-start justify-between mb-1">
              <div class="flex items-center gap-2">
                <Stethoscope :size="14" class="text-teal-600 dark:text-teal-400" />
                <span class="text-xs text-slate-500 dark:text-slate-400">Médecin de famille</span>
              </div>
              <button 
                v-if="medecinInfo"
                @click="viewMedecinDetails"
                class="p-1 hover:bg-slate-100 dark:hover:bg-slate-800 rounded transition-colors"
                title="Voir détails"
              >
                <Eye :size="12" class="text-teal-600 dark:text-teal-400" />
              </button>
            </div>
            <p class="font-semibold text-sm text-slate-900 dark:text-slate-100 mb-0.5">
              Dr. {{ getMedecinName() }}
            </p>
            <p v-if="client.medecin_famille_No_Licence" class="text-xs text-slate-500 dark:text-slate-400 font-mono">
              {{ client.medecin_famille_No_Licence }}
            </p>
          </div>

          <div class="h-px bg-slate-200 dark:bg-slate-700"></div>

          <!-- RAMQ -->
          <div v-if="client.numero_assurance_maladie">
            <div class="flex items-center gap-2 mb-1">
              <CreditCard :size="14" class="text-blue-600 dark:text-blue-400" />
              <span class="text-xs text-slate-500 dark:text-slate-400">RAMQ</span>
            </div>
            <p class="font-mono text-sm font-semibold text-slate-900 dark:text-slate-100">
              {{ client.numero_assurance_maladie }}
            </p>
          </div>

          <!-- Dossiers HCM/CHAUR -->
          <div v-if="client.no_hcm || client.no_chaur">
            <div class="h-px bg-slate-200 dark:bg-slate-700 mb-3"></div>
            <div class="flex items-center gap-2 mb-2">
              <FileText :size="14" class="text-purple-600 dark:text-purple-400" />
              <span class="text-xs text-slate-500 dark:text-slate-400">Dossiers cliniques</span>
            </div>
            <div class="space-y-1">
              <p v-if="client.no_hcm" class="text-xs text-slate-700 dark:text-slate-300">
                <span class="font-semibold">HCM:</span> {{ client.no_hcm }}
              </p>
              <p v-if="client.no_chaur" class="text-xs text-slate-700 dark:text-slate-300">
                <span class="font-semibold">CHAUR:</span> {{ client.no_chaur }}
              </p>
            </div>
          </div>
        </div>

        <!-- Colonne 2: Localisation -->
        <div class="space-y-4">
          <h3 class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2 mb-3">
            <MapPin :size="14" class="text-orange-600 dark:text-orange-400" />
            Localisation
          </h3>

          <!-- Adresse -->
          <div v-if="client.adresse">
            <div class="flex items-center justify-between mb-1">
              <span class="text-xs text-slate-500 dark:text-slate-400">Domicile</span>
              <button 
                @click="openGoogleMaps" 
                class="text-xs text-blue-600 dark:text-blue-400 hover:underline flex items-center gap-1"
              >
                <Map :size="12" />
                Voir sur Maps
              </button>
            </div>
            <p class="text-sm font-medium text-slate-900 dark:text-slate-100">
              {{ client.adresse }}
            </p>
            <p class="text-xs text-slate-600 dark:text-slate-400 mt-1">
              {{ client.ville }}, {{ client.code_postal }}
            </p>
          </div>

          <!-- Établissement -->
          <div v-if="getEtablissement()">
            <div class="h-px bg-slate-200 dark:bg-slate-700 mb-3"></div>
            <div class="flex items-center gap-2 mb-1">
              <Building2 :size="14" class="text-orange-600 dark:text-orange-400" />
              <span class="text-xs text-slate-500 dark:text-slate-400">Hébergement</span>
            </div>
            <p class="text-sm font-semibold text-slate-900 dark:text-slate-100">
              {{ getEtablissement() }}
            </p>
          </div>
        </div>

        <!-- Colonne 3: Contact -->
        <div class="space-y-4">
          <h3 class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2 mb-3">
            <Phone :size="14" class="text-slate-600 dark:text-slate-400" />
            Coordonnées
          </h3>

          <!-- Téléphones et email -->
          <div class="space-y-3">
            <a v-if="client.telephone" 
               :href="`tel:${client.telephone}`"
               class="flex items-center gap-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors group">
              <Phone :size="14" class="text-slate-400 group-hover:text-blue-600 dark:group-hover:text-blue-400" />
              <span class="text-sm font-medium text-slate-900 dark:text-slate-100">
                {{ formatPhone(client.telephone) }}
              </span>
            </a>

            <a v-if="client.cellulaire" 
               :href="`tel:${client.cellulaire}`"
               class="flex items-center gap-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors group">
              <Smartphone :size="14" class="text-slate-400 group-hover:text-blue-600 dark:group-hover:text-blue-400" />
              <span class="text-sm font-medium text-slate-900 dark:text-slate-100">
                {{ formatPhone(client.cellulaire) }}
              </span>
            </a>

            <a v-if="client.email" 
               :href="`mailto:${client.email}`"
               class="flex items-center gap-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors group">
              <Mail :size="14" class="text-slate-400 group-hover:text-blue-600 dark:group-hover:text-blue-400" />
              <span class="text-sm font-medium text-slate-900 dark:text-slate-100 truncate">
                {{ client.email }}
              </span>
            </a>
          </div>
        </div>
      </div>

      <!-- Notes importantes (si présentes) -->
      <div v-if="client.note_fixe" class="mt-6 pt-6 border-t border-slate-200 dark:border-slate-700">
        <div class="flex items-start gap-3 p-4 bg-amber-50 dark:bg-amber-950/20 border-l-4 border-amber-500 rounded">
          <AlertTriangle :size="16" class="text-amber-600 dark:text-amber-400 flex-shrink-0 mt-0.5" />
          <div class="flex-1">
            <p class="text-xs font-bold text-amber-900 dark:text-amber-100 uppercase tracking-wide mb-1">
              Note importante
            </p>
            <p class="text-sm text-amber-800 dark:text-amber-200 whitespace-pre-wrap">
              {{ client.note_fixe }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <div class="px-6 py-3 bg-slate-50 dark:bg-slate-800/50 border-t border-slate-200 dark:border-slate-700">
      <div class="flex items-center justify-between text-xs text-slate-500 dark:text-slate-400">
        <span>Dernière MAJ: {{ formatDate(client.created_at) }}</span>
        <span v-if="client.date_maj">Modifié: {{ client.date_maj }}</span>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  Shield, AlertCircle, Heart, MapPin, Phone, Mail, FileDown,
  Cake, Eye, Stethoscope, CreditCard, Building2, FileText,
  Map, Smartphone, AlertTriangle
} from 'lucide-vue-next'

const props = defineProps({
  client: { type: Object, required: true }
})

const emit = defineEmits(['view-medecin'])

const contacts = ref([])
const medecinInfo = ref(null)

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

// Contact d'urgence
const urgenceContact = computed(() => {
  return contacts.value.find(c => c.contact_urgence === 1) || null
})

const getInitials = () => {
  return `${props.client.prenom[0]}${props.client.nom[0]}`.toUpperCase()
}

const getSexeBadge = () => {
  const sexe = props.client.sexe || 'F'
  return sexe === 'M' ? '♂ ' : '♀ '
}

const getStatusText = () => {
  if (props.client.dcd == 1) return 'Décédé'
  if (props.client.actif == 1) return 'Actif'
  return 'Archivé'
}

const getStatusClass = () => {
  if (props.client.dcd == 1) return 'bg-gray-200 text-gray-700'
  if (props.client.actif == 1) return 'bg-green-100 text-green-700'
  return 'bg-red-100 text-red-700'
}

const getEtablissement = () => {
  if (props.client.rpa_id) return `RPA #${props.client.rpa_id}`
  if (props.client.chsld_id) return `CHSLD #${props.client.chsld_id}`
  if (props.client.ri_id) return `RI #${props.client.ri_id}`
  return null
}

const getMedecinName = () => {
  if (medecinInfo.value) {
    return `${medecinInfo.value.prenom} ${medecinInfo.value.nom}`
  }
  return props.client.medecin_famille_No_Licence || 'Non spécifié'
}

const formatPhone = (phone) => {
  if (!phone) return ''
  const cleaned = phone.replace(/\D/g, '')
  
  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }
  
  if (cleaned.length === 11 && cleaned[0] === '1') {
    return `+1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`
  }
  
  return phone
}

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString('fr-CA')
}

const openGoogleMaps = () => {
  const address = `${props.client.adresse}, ${props.client.ville}, ${props.client.code_postal}`
  window.open(`https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(address)}`, '_blank')
}

const viewMedecinDetails = () => {
  emit('view-medecin', props.client.medecin_famille_No_Licence)
}

const exportToPDF = () => {
  alert('Export PDF à implémenter')
}

const loadContacts = async () => {
  try {
    const result = await window.go.main.App.GetAllContactsByClientID(props.client.id)
    contacts.value = result || []
  } catch (error) {
    console.error('Erreur chargement contacts:', error)
  }
}

const loadMedecinInfo = async () => {
  if (!props.client.medecin_famille_No_Licence) return
  
  try {
    const result = await window.go.main.App.GetMedecinByLicence(props.client.medecin_famille_No_Licence)
    medecinInfo.value = result
  } catch (error) {
    console.error('Erreur chargement médecin:', error)
  }
}

onMounted(() => {
  loadContacts()
  loadMedecinInfo()
})
</script>

<style scoped>
/* Transitions fluides */
a {
  transition: all 0.2s ease;
}
</style>