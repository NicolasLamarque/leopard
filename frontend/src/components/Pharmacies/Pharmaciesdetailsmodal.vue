<!-- frontend/src/components/Pharmacies/PharmacieDetailsModal.vue -->
<template>
  <div class="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-slate-900/40 backdrop-blur-sm">
    
    <div class="bg-white dark:bg-slate-900 w-full max-w-5xl rounded-xl shadow-xl overflow-hidden animate-in zoom-in-95 duration-200 border border-slate-200 dark:border-slate-800">
      
      <!-- HEADER - Palette slate/teal subtile -->
      <div class="relative px-8 py-6 bg-gradient-to-br from-slate-50 via-slate-100 to-slate-50 dark:from-slate-800 dark:via-slate-850 dark:to-slate-800 border-b border-slate-200 dark:border-slate-700">
        <div class="flex justify-between items-start">
          <div class="flex items-center gap-5">
            <!-- Icône pilule avec accent teal subtil -->
            <div class="p-3.5 bg-teal-50 dark:bg-teal-950/30 rounded-xl border border-teal-100 dark:border-teal-900/50">
              <Pill :size="28" class="text-teal-600 dark:text-teal-400" />
            </div>
            <div>
              <h2 class="text-2xl font-bold text-slate-950 dark:text-white mb-1.5">
                {{ isEditMode ? 'Modifier la pharmacie' : pharmacie.NomOrganisation }}
              </h2>
              <div v-if="!isEditMode" class="flex items-center gap-2.5 text-sm">
                <!-- Badge bannière épuré -->
                <span 
                  v-if="pharmacie.Banniere"
                  class="px-2.5 py-1 bg-slate-100 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-md font-medium text-slate-700 dark:text-slate-300 text-xs"
                >
                  {{ pharmacie.Banniere }}
                </span>
                <!-- Statut ouvert/fermé - ROUGE si fermé maintenant -->
                <span 
                  :class="isOpenNow 
                    ? 'bg-teal-50 dark:bg-teal-950/30 border-teal-200 dark:border-teal-900/50 text-teal-700 dark:text-teal-400' 
                    : 'bg-rose-50 dark:bg-rose-950/30 border-rose-200 dark:border-rose-900/50 text-rose-700 dark:text-rose-400'"
                  class="px-2.5 py-1 rounded-md font-semibold flex items-center gap-1.5 text-xs border shadow-sm"
                >
                  <span :class="isOpenNow ? 'bg-teal-500 animate-pulse' : 'bg-rose-500'" class="w-1.5 h-1.5 rounded-full"></span>
                  {{ isOpenNow ? 'Ouverte maintenant' : 'Fermée maintenant' }}
                </span>
              </div>
            </div>
          </div>
          
          <button 
            @click="$emit('close')" 
            class="p-2 hover:bg-slate-200 dark:hover:bg-slate-700 rounded-lg transition-colors text-slate-600 dark:text-slate-400"
          >
            <X :size="20" />
          </button>
        </div>
      </div>

      <!-- CONTENU -->
      <div class="p-8 max-h-[70vh] overflow-y-auto space-y-8">
        
        <!-- MODE ÉDITION -->
        <div v-if="isEditMode" class="space-y-6">
          <!-- Informations générales -->
          <section>
            <h3 class="text-base font-semibold text-slate-900 dark:text-slate-100 mb-4 flex items-center gap-2">
              <div class="w-1 h-5 bg-teal-500 rounded-full"></div>
              Informations générales
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  Nom de l'organisation *
                </label>
                <input
                  v-model="editedPharmacie.NomOrganisation"
                  type="text"
                  class="w-full px-3.5 py-2.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 focus:ring-2 focus:ring-teal-500 focus:border-transparent transition-shadow"
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  Bannière
                </label>
                <select
                  v-model="editedPharmacie.Banniere"
                  class="w-full px-3.5 py-2.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                >
                  <option value="">Aucune</option>
                  <option value="Jean Coutu">Jean Coutu</option>
                  <option value="Pharmaprix">Pharmaprix</option>
                  <option value="Familiprix">Familiprix</option>
                  <option value="Uniprix">Uniprix</option>
                  <option value="Proxim">Proxim</option>
                  <option value="Brunet">Brunet</option>
                  <option value="Indépendante">Indépendante</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  Adresse
                </label>
                <input
                  v-model="editedPharmacie.Adresse"
                  type="text"
                  class="w-full px-3.5 py-2.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  Ville
                </label>
                <input
                  v-model="editedPharmacie.Ville"
                  type="text"
                  class="w-full px-3.5 py-2.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  Téléphone
                </label>
                <input
                  v-model="editedPharmacie.Tel"
                  type="tel"
                  class="w-full px-3.5 py-2.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  Télécopieur
                </label>
                <input
                  v-model="editedPharmacie.Fax"
                  type="tel"
                  class="w-full px-3.5 py-2.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 focus:ring-2 focus:ring-teal-500 focus:border-transparent"
                />
              </div>
            </div>
          </section>

          <!-- Horaires -->
          <section>
            <h3 class="text-base font-semibold text-slate-900 dark:text-slate-100 mb-4 flex items-center gap-2">
              <div class="w-1 h-5 bg-teal-500 rounded-full"></div>
              Heures d'ouverture
            </h3>
            <div 
              v-for="jour in jours" 
              :key="jour.key"
              class="flex items-center gap-4 py-3 border-b border-slate-100 dark:border-slate-800 last:border-0"
            >
              <div class="w-28">
                <label class="flex items-center gap-2 cursor-pointer group">
                  <input
                    v-model="editedPharmacie[`${jour.key}Ouvert`]"
                    type="checkbox"
                    class="w-4 h-4 text-teal-600 border-slate-300 rounded focus:ring-teal-500"
                  />
                  <span class="text-sm font-medium text-slate-700 dark:text-slate-300 group-hover:text-slate-900 dark:group-hover:text-slate-100">
                    {{ jour.label }}
                  </span>
                </label>
              </div>
              <div class="flex items-center gap-2 flex-1">
                <input
                  v-model="editedPharmacie[`HeureOuverture${jour.key}`]"
                  type="time"
                  :disabled="!editedPharmacie[`${jour.key}Ouvert`]"
                  class="px-3 py-1.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 disabled:opacity-40 disabled:cursor-not-allowed focus:ring-2 focus:ring-teal-500 focus:border-transparent text-sm"
                />
                <span class="text-slate-500 text-sm">à</span>
                <input
                  v-model="editedPharmacie[`HeureFermeture${jour.key}`]"
                  type="time"
                  :disabled="!editedPharmacie[`${jour.key}Ouvert`]"
                  class="px-3 py-1.5 border border-slate-300 dark:border-slate-600 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-slate-100 disabled:opacity-40 disabled:cursor-not-allowed focus:ring-2 focus:ring-teal-500 focus:border-transparent text-sm"
                />
              </div>
            </div>
          </section>
        </div>

        <!-- MODE LECTURE -->
        <div v-else>
          <!-- Coordonnées avec bouton Google Maps -->
          <section>
            <h3 class="text-base font-semibold text-slate-900 dark:text-slate-100 mb-4 flex items-center gap-2">
              <MapPin :size="18" class="text-teal-600 dark:text-teal-400" />
              Coordonnées
            </h3>

            <div class="space-y-3">
              <!-- Adresse avec bouton Google Maps -->
              <div v-if="pharmacie.Adresse" class="group">
                <div class="flex items-start justify-between p-4 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 hover:border-teal-200 dark:hover:border-teal-800 transition-colors">
                  <div class="flex-1">
                    <p class="text-xs font-medium text-slate-500 dark:text-slate-400 mb-1.5">Adresse</p>
                    <p class="text-sm text-slate-700 dark:text-slate-300 leading-relaxed">
                      {{ pharmacie.Adresse }}<br>
                      {{ pharmacie.Ville }}{{ pharmacie.Region ? `, ${pharmacie.Region}` : '' }}
                    </p>
                  </div>
                  <!-- Bouton Google Maps -->
                  <a 
                    :href="getGoogleMapsUrl()"
                    target="_blank"
                    class="ml-4 p-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 hover:border-teal-500 dark:hover:border-teal-500 rounded-lg transition-all opacity-0 group-hover:opacity-100 transform translate-x-2 group-hover:translate-x-0"
                    title="Ouvrir dans Google Maps"
                  >
                    <MapPin :size="16" class="text-teal-600 dark:text-teal-400" />
                  </a>
                </div>
              </div>

              <ContactRow 
                v-if="pharmacie.Tel"
                :icon="Phone"
                label="Téléphone"
                :value="formatPhone(pharmacie.Tel)"
                :href="`tel:${pharmacie.Tel}`"
                action-label="Appeler"
                color="slate"
              />

              <ContactRow 
                v-if="pharmacie.Fax"
                :icon="Printer"
                label="Télécopieur"
                :value="formatPhone(pharmacie.Fax)"
                :href="`tel:${pharmacie.Fax}`"
                action-label="Faxer"
                color="stone"
                :hide-action="true"
              />
            </div>
          </section>

          <!-- Horaires d'ouverture -->
          <section>
            <h3 class="text-base font-semibold text-slate-900 dark:text-slate-100 mb-4 flex items-center gap-2">
              <Clock :size="18" class="text-teal-600 dark:text-teal-400" />
              Heures d'ouverture
            </h3>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-2.5">
              <div
                v-for="jour in jours"
                :key="jour.key"
                :class="[
                  'p-3.5 rounded-lg border transition-all',
                  // Aujourd'hui ET fermé MAINTENANT → ROUGE
                  isToday(jour.index) && !isOpenNow
                    ? 'bg-rose-50 dark:bg-rose-950/20 border-rose-200 dark:border-rose-900/50 shadow-md ring-2 ring-rose-200 dark:ring-rose-900/50'
                    // Aujourd'hui ET ouvert → teal
                    : isToday(jour.index) && isOpenNow
                      ? 'bg-teal-50 dark:bg-teal-950/20 border-teal-200 dark:border-teal-900/50 shadow-sm ring-2 ring-teal-200 dark:ring-teal-900/50' 
                      // Jour fermé (pas aujourd'hui) → jaune doré subtil
                      : (!pharmacie[`${jour.key}Ouvert`] || pharmacie[`${jour.key}Ouvert`] === 0)
                        ? 'bg-amber-50/50 dark:bg-amber-950/10 border-amber-200 dark:border-amber-900/30'
                        // Sinon → gris normal
                        : 'bg-slate-50 dark:bg-slate-800/50 border-slate-200 dark:border-slate-700 hover:border-slate-300 dark:hover:border-slate-600'
                ]"
              >
                <div class="flex items-center justify-between">
                  <span 
                    :class="[
                      'text-sm font-medium flex items-center gap-1.5',
                      isToday(jour.index) && !isOpenNow
                        ? 'text-rose-700 dark:text-rose-400'
                        : isToday(jour.index)
                          ? 'text-teal-700 dark:text-teal-400'
                          : (!pharmacie[`${jour.key}Ouvert`] || pharmacie[`${jour.key}Ouvert`] === 0)
                            ? 'text-amber-700 dark:text-amber-400'
                            : 'text-slate-900 dark:text-slate-100'
                    ]"
                  >
                    {{ jour.label }}
                    <span v-if="isToday(jour.index)" 
                      :class="isOpenNow ? 'text-teal-500' : 'text-rose-500'"
                      class="text-xs font-bold"
                    >
                      • MAINTENANT
                    </span>
                  </span>
                  <span 
                    :class="[
                      'text-sm font-mono font-medium',
                      isToday(jour.index) && !isOpenNow
                        ? 'text-rose-600 dark:text-rose-400 font-bold'
                        : isToday(jour.index)
                          ? 'text-teal-600 dark:text-teal-400 font-semibold'
                          : (!pharmacie[`${jour.key}Ouvert`] || pharmacie[`${jour.key}Ouvert`] === 0)
                            ? 'text-amber-600 dark:text-amber-400'
                            : 'text-slate-600 dark:text-slate-400'
                    ]"
                  >
                    {{ getHours(jour.key) }}
                  </span>
                </div>
              </div>
            </div>
          </section>
        </div>

      </div>

      <!-- LISTE DES CLIENTS -->
      <div v-if="!isEditMode" class="px-8 pb-8">
        <PharmacieClientsList 
          :pharmacie-id="pharmacie.id"
          @view-client="$emit('view-client', $event)"
        />
      </div>

      <!-- FOOTER -->
      <div class="px-8 py-5 border-t border-slate-200 dark:border-slate-700 bg-slate-50 dark:bg-slate-800/30 flex justify-between items-center">
        <div v-if="!isEditMode && pharmacie.createdAt" class="text-xs text-slate-500 dark:text-slate-400">
          Ajoutée le {{ formatDate(pharmacie.createdAt) }}
        </div>
        <div v-else></div>
        
        <div class="flex gap-2.5">
          <button 
            v-if="isEditMode"
            @click="cancelEdit"
            class="px-4 py-2.5 bg-white hover:bg-slate-50 dark:bg-slate-800 dark:hover:bg-slate-700 text-slate-700 dark:text-slate-200 border border-slate-300 dark:border-slate-600 rounded-lg font-medium transition-colors text-sm"
          >
            Annuler
          </button>

          <button 
            v-if="isEditMode"
            @click="saveChanges"
            class="px-4 py-2.5 bg-teal-600 hover:bg-teal-700 dark:bg-teal-600 dark:hover:bg-teal-700 text-white rounded-lg font-medium transition-colors flex items-center gap-2 shadow-sm text-sm"
          >
            <Save :size="16" />
            Enregistrer
          </button>

          <button 
            v-if="!isEditMode"
            @click="startEdit"
            class="px-4 py-2.5 bg-slate-700 hover:bg-slate-800 dark:bg-slate-700 dark:hover:bg-slate-600 text-white rounded-lg font-medium transition-colors flex items-center gap-2 text-sm"
          >
            <Edit :size="16" />
            Modifier
          </button>
          
          <button 
            v-if="!isEditMode"
            @click="confirmDelete"
            class="px-4 py-2.5 bg-white hover:bg-rose-50 dark:bg-slate-800 dark:hover:bg-rose-950/30 text-rose-600 dark:text-rose-400 border border-rose-200 dark:border-rose-900/50 hover:border-rose-300 dark:hover:border-rose-800 rounded-lg font-medium transition-colors flex items-center gap-2 text-sm"
          >
            <Trash2 :size="16" />
            Supprimer
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Pill, X, MapPin, Phone, Printer, Clock, Edit, Save, Trash2 } from 'lucide-vue-next'
import ContactRow from '../shared/ContactRow.vue'
import PharmacieClientsList from './PharmacieClientsList.vue'

const props = defineProps({
  pharmacie: { type: Object, required: true }
})

const emit = defineEmits(['close', 'save', 'delete', 'view-client'])

const isEditMode = ref(false)
const editedPharmacie = ref({})

const jours = [
  { key: 'Dimanche', label: 'Dimanche', index: 0 },
  { key: 'Lundi', label: 'Lundi', index: 1 },
  { key: 'Mardi', label: 'Mardi', index: 2 },
  { key: 'Mercredi', label: 'Mercredi', index: 3 },
  { key: 'Jeudi', label: 'Jeudi', index: 4 },
  { key: 'Vendredi', label: 'Vendredi', index: 5 },
  { key: 'Samedi', label: 'Samedi', index: 6 }
]

// Vérifier si c'est aujourd'hui
const isToday = (index) => {
  return new Date().getDay() === index
}

// Vérifier si ouvert maintenant
const isOpenNow = computed(() => {
  const now = new Date()
  const today = now.getDay()
  const jour = jours.find(j => j.index === today)
  
  if (!jour) return false
  
  const ouvertKey = `${jour.key}Ouvert`
  if (!props.pharmacie[ouvertKey] || props.pharmacie[ouvertKey] === 0) return false
  
  const ouvertureKey = `HeureOuverture${jour.key}`
  const fermetureKey = `HeureFermeture${jour.key}`
  
  const ouverture = props.pharmacie[ouvertureKey]
  const fermeture = props.pharmacie[fermetureKey]
  
  if (!ouverture || !fermeture) return false
  
  try {
    const currentTime = now.getHours() * 60 + now.getMinutes()
    const [openH, openM] = ouverture.split(':').map(Number)
    const [closeH, closeM] = fermeture.split(':').map(Number)
    
    const openTime = openH * 60 + openM
    const closeTime = closeH * 60 + closeM
    
    return currentTime >= openTime && currentTime <= closeTime
  } catch {
    return false
  }
})

// Obtenir les heures d'un jour
const getHours = (jour) => {
  const ouvertKey = `${jour}Ouvert`
  const ouvertureKey = `HeureOuverture${jour}`
  const fermetureKey = `HeureFermeture${jour}`
  
  if (!props.pharmacie[ouvertKey] || props.pharmacie[ouvertKey] === 0) {
    return 'Fermée'
  }
  
  const ouverture = props.pharmacie[ouvertureKey] || 'N/D'
  const fermeture = props.pharmacie[fermetureKey] || 'N/D'
  
  return `${ouverture} - ${fermeture}`
}

// Formater téléphone
const formatPhone = (phone) => {
  if (!phone) return ''
  
  const cleaned = phone.replace(/\D/g, '')
  
  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }
  
  if (cleaned.length === 11 && cleaned.startsWith('1')) {
    return `1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`
  }
  
  return phone
}

// Générer URL Google Maps
const getGoogleMapsUrl = () => {
  const address = `${props.pharmacie.Adresse}, ${props.pharmacie.Ville}, ${props.pharmacie.Region || 'Québec'}, Canada`
  return `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(address)}`
}

// Formater date
const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('fr-CA')
}

// Mode édition
const startEdit = () => {
  isEditMode.value = true
  editedPharmacie.value = { ...props.pharmacie }
}

const cancelEdit = () => {
  isEditMode.value = false
  editedPharmacie.value = {}
}

const saveChanges = () => {
  emit('save', editedPharmacie.value)
  isEditMode.value = false
}

const confirmDelete = () => {
  if (confirm(`Supprimer "${props.pharmacie.NomOrganisation}" ?`)) {
    emit('delete', props.pharmacie)
  }
}
</script>