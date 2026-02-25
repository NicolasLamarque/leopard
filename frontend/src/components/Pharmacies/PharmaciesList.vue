<!-- frontend/src/components/Pharmacies/PharmacieList.vue -->
<template>
  <div>
    <!-- Loading -->
    <div v-if="loading" class="flex items-center justify-center py-12">
      <Loader2 :size="32" class="animate-spin text-emerald-600" />
      <span class="ml-3 text-gray-600 dark:text-gray-400">Chargement des pharmacies...</span>
    </div>

    <!-- Empty state -->
    <div v-else-if="!pharmacies || pharmacies.length === 0" class="text-center py-12">
      <div class="inline-flex items-center justify-center w-20 h-20 bg-gray-100 dark:bg-gray-800 rounded-full mb-4">
        <Pill :size="40" class="text-gray-400" />
      </div>
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">
        Aucune pharmacie trouvée
      </h3>
      <p class="text-sm text-gray-500 dark:text-gray-400">
        Essayez de modifier vos filtres de recherche
      </p>
    </div>

    <!-- Liste des pharmacies -->
    <div v-else class="space-y-3">
      <div
        v-for="pharmacie in pharmacies"
        :key="pharmacie.ID"
        @click="$emit('select', pharmacie)"
        class="group bg-white dark:bg-gray-800 rounded-xl border border-gray-200 dark:border-gray-700 p-5 hover:shadow-lg hover:border-teal-700 dark:hover:border-emerald-800 transition-all cursor-pointer"
      >
        <div class="flex items-start justify-between">
          <!-- Info principale -->
          <div class="flex-1">
           <div class="flex items-start gap-3 mb-3">
  <div 
    :class="[
      'p-2 rounded-lg',
      isOpenNow(pharmacie) 
        ? 'bg-teal-100 dark:bg-teal-900/30' 
        : 'bg-rose-100 dark:bg-rose-900/30'
    ]"
  >
    <Pill 
      :size="24" 
      :class="isOpenNow(pharmacie) ? 'text-teal-600 dark:text-teal-400' : 'text-rose-600 dark:text-rose-400'" 
    />
  </div>
  
  <div class="flex-1">
    <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-1 group-hover:text-teal-700 dark:group-hover:text-teal-600 transition-colors">
      {{ pharmacie.NomOrganisation }}
    </h3>
    
    <div class="flex items-center gap-2 flex-wrap">
  <!-- Logo bannière (au lieu du badge texte) -->
  <img 
    v-if="pharmacie.Banniere && getBanniereLogoPath(pharmacie.Banniere)"
    :src="getBanniereLogoPath(pharmacie.Banniere)"
    :alt="pharmacie.Banniere"
    class="h-7 object-contain"
    @error="handleLogoError"
  />
  
  <!-- Fallback si pas de logo -->
  <span 
    v-else-if="pharmacie.Banniere"
    class="px-2.5 py-1 bg-slate-100 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-700 dark:text-slate-300 text-xs font-medium rounded"
  >
    {{ pharmacie.Banniere }}
  </span>
  
  <!-- Statut ouverture -->
  <span 
    :class="isOpenNow(pharmacie) 
      ? 'bg-teal-50 dark:bg-teal-950/30 border-teal-200 dark:border-teal-900/50 text-teal-700 dark:text-teal-400' 
      : 'bg-rose-50 dark:bg-rose-950/30 border-rose-200 dark:border-rose-900/50 text-rose-700 dark:text-rose-400'"
    class="px-2 py-1 text-xs font-semibold rounded flex items-center gap-1 border"
  >
    <span class="w-1.5 h-1.5 rounded-full" :class="isOpenNow(pharmacie) ? 'bg-teal-500' : 'bg-rose-500'"></span>
    {{ isOpenNow(pharmacie) ? 'Ouverte' : 'Fermée' }}
  </span>
</div>
  </div>
</div>

            <!-- Coordonnées -->
            <div class="space-y-2 text-sm">
              <!-- Adresse -->
              <div v-if="pharmacie.Adresse" class="flex items-start gap-2 text-gray-600 dark:text-gray-400">
                <MapPin :size="16" class="mt-0.5 flex-shrink-0" />
                <div>
                  <p>{{ pharmacie.Adresse }}</p>
                  <p v-if="pharmacie.Ville || pharmacie.Region">
                    {{ pharmacie.Ville }}{{ pharmacie.Region ? `, ${pharmacie.Region}` : '' }}
                  </p>
                </div>
              </div>

              <!-- Téléphone -->
              <div v-if="pharmacie.Tel" class="flex items-center gap-2 text-gray-600 dark:text-gray-400">
                <Phone :size="16" />
                <a 
                  :href="`tel:${pharmacie.Tel}`"
                  @click.stop
                  class="hover:text-emerald-600 dark:hover:text-emerald-400 transition-colors"
                >
                  {{ formatPhone(pharmacie.Tel) }}
                </a>
              </div>

              <!-- Horaires aujourd'hui -->
              <div class="flex items-center gap-2 text-gray-600 dark:text-gray-400">
                <Clock :size="16" />
                <span>{{ getTodayHours(pharmacie) }}</span>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex items-center gap-2">
            <button
              @click.stop="$emit('select', pharmacie)"
              class="p-2 hover:bg-emerald-100 dark:hover:bg-emerald-900/30 rounded-lg transition-colors group/btn"
              title="Voir les détails"
            >
              <Eye :size="20" class="text-gray-400 group-hover/btn:text-emerald-600 dark:group-hover/btn:text-emerald-400" />
            </button>
            
            <button
              @click.stop="$emit('delete', pharmacie)"
              class="p-2 hover:bg-red-100 dark:hover:bg-red-900/30 rounded-lg transition-colors group/btn"
              title="Supprimer"
            >
              <Trash2 :size="20" class="text-gray-400 group-hover/btn:text-red-600 dark:group-hover/btn:text-red-400" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Pill, MapPin, Phone, Clock, Eye, Trash2, Loader2 } from 'lucide-vue-next'
import JeanCoutu from '../../assets/logoPharma/jean-coutu.png'
import Pharmaprix from '../../assets/logoPharma/pharmaprix.png'
import Familiprix from '../../assets/logoPharma/familiprix.png' 
import Uniprix from '../../assets/logoPharma/uniprix.png'
import Proxim from '../../assets/logoPharma/proxim.png'
import Brunet from '../../assets/logoPharma/brunet.png'
import Independante from '../../assets/logoPharma/independante.png'


const props = defineProps({
  pharmacies: { type: Array, default: () => [] },
  loading: { type: Boolean, default: false }
})

defineEmits(['select', 'delete'])

const formatPhone = (phone) => {
  if (!phone) return ''
  
  // Enlever tout sauf les chiffres
  const cleaned = phone.replace(/\D/g, '')
  
  // Si 10 chiffres : format (819) 999-0011
  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }
  
  // Si 11 chiffres commençant par 1 : format 1 (819) 999-0011
  if (cleaned.length === 11 && cleaned.startsWith('1')) {
    return `1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`
  }
  
  // Sinon retourner tel quel
  return phone
}
const getBanniereLogoPath = (banniere) => {
  if (!banniere) return null
  
  // Mapping des noms de bannières vers les fichiers
  const logoMap = {
    'Jean Coutu': JeanCoutu,
    'Pharmaprix': Pharmaprix,
    'Familiprix': Familiprix,
    'Uniprix': Uniprix,
    'Proxim': Proxim,
    'Brunet': Brunet,
    'Indépendante': Independante
  }
  
  return logoMap[banniere] || null
}

// Gestion d'erreur si le logo n'existe pas
const handleLogoError = (event) => {
  event.target.src = '/logos/independante.png'
}
// Obtenir les heures d'ouverture d'aujourd'hui
const getTodayHours = (pharmacie) => {
  const jours = ['Dimanche', 'Lundi', 'Mardi', 'Mercredi', 'Jeudi', 'Vendredi', 'Samedi']
  const today = new Date().getDay()
  const jour = jours[today]
  
  const ouvertKey = `${jour}Ouvert`
  const ouvertureKey = `HeureOuverture${jour}`
  const fermetureKey = `HeureFermeture${jour}`
  
  if (!pharmacie[ouvertKey]) {
    return 'Fermée aujourd\'hui'
  }
  
  const ouverture = pharmacie[ouvertureKey] || 'N/D'
  const fermeture = pharmacie[fermetureKey] || 'N/D'
  
  return `${ouverture} - ${fermeture}`
}

// Vérifier si la pharmacie est ouverte maintenant
const isOpenNow = (pharmacie) => {
  const jours = ['Dimanche', 'Lundi', 'Mardi', 'Mercredi', 'Jeudi', 'Vendredi', 'Samedi']
  const now = new Date()
  const today = now.getDay()
  const jour = jours[today]
  
  const ouvertKey = `${jour}Ouvert`
  
  if (!pharmacie[ouvertKey]) {
    return false
  }
  
  const ouvertureKey = `HeureOuverture${jour}`
  const fermetureKey = `HeureFermeture${jour}`
  
  const ouverture = pharmacie[ouvertureKey]
  const fermeture = pharmacie[fermetureKey]
  
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
}
</script>