<template>
  <div class="space-y-2">

    <!-- Label -->
    <label class="fk-label">{{ label || 'Pharmacie rattachée' }}</label>

    <!-- Input + bouton clear -->
    <div class="flex gap-2 max-w-md">
      <div class="relative flex-1">
        <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-400">
          <Search :size="15" />
        </span>
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Nom, ville ou bannière..."
          class="w-full pl-9 pr-3 py-1.5 text-sm bg-white dark:bg-gray-800 dark:text-gray-100
                 border border-gray-300 dark:border-gray-600 rounded-lg
                 focus:ring-2 focus:ring-teal-500 focus:border-transparent
                 placeholder-gray-400 dark:placeholder-gray-500 outline-none transition-all"
          @focus="showDropdown = true"
          @blur="onBlur"
        />

        <!-- Dropdown -->
        <div
          v-if="showDropdown && filteredPharmacies.length > 0"
          class="absolute z-50 w-full mt-1 bg-white dark:bg-gray-800
                 border border-gray-200 dark:border-gray-700 rounded-lg shadow-xl
                 max-h-64 overflow-y-auto"
        >
          <div
            v-for="pharma in filteredPharmacies"
            :key="pharma.ID"
            @mousedown.prevent="selectPharmacie(pharma)"
            class="px-3 py-2.5 hover:bg-teal-50 dark:hover:bg-teal-900/20 cursor-pointer
                   border-b last:border-0 border-gray-100 dark:border-gray-700 transition-colors"
          >
            <!-- Ligne 1 : nom + logo -->
            <div class="flex items-center justify-between gap-2">
              <span class="font-semibold text-sm text-gray-900 dark:text-gray-100 truncate">
                {{ pharma.NomOrganisation }}
              </span>
              <img
                v-if="getBanniereLogo(pharma.Banniere)"
                :src="getBanniereLogo(pharma.Banniere)"
                :alt="pharma.Banniere"
                class="h-5 object-contain flex-shrink-0"
              />
              <span
                v-else-if="pharma.Banniere"
                class="text-[10px] font-medium px-1.5 py-0.5 bg-slate-100 dark:bg-slate-700
                       text-slate-600 dark:text-slate-300 rounded flex-shrink-0"
              >{{ pharma.Banniere }}</span>
            </div>
            <!-- Ligne 2 : adresse + tél -->
            <div class="flex items-center gap-3 mt-0.5 text-xs text-gray-400 dark:text-gray-500">
              <span class="flex items-center gap-1">
                <MapPin :size="10" />
                {{ pharma.Adresse ? `${pharma.Adresse}, ` : '' }}{{ pharma.Ville }}
              </span>
              <span v-if="pharma.Tel" class="flex items-center gap-1">
                <Phone :size="10" />
                {{ formatPhone(pharma.Tel) }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Bouton clear -->
      <button
        v-if="selectedPharmacie"
        type="button"
        @click="clearSelection"
        class="px-2.5 py-1.5 text-rose-500 hover:bg-rose-50 dark:hover:bg-rose-900/20 rounded-lg transition-colors"
        title="Retirer la pharmacie"
      >
        <X :size="15" />
      </button>
    </div>

    <!-- Toast sélection (disparaît après 3s) -->
    <Transition name="toast-slide">
      <div
        v-if="showToast"
        class="flex items-center gap-2 max-w-md px-3 py-2 rounded-lg
               bg-teal-50 dark:bg-teal-900/30 border border-teal-200 dark:border-teal-800
               text-teal-700 dark:text-teal-300 text-xs font-medium"
      >
        <CheckCircle :size="14" class="flex-shrink-0" />
        Pharmacie sélectionnée
      </div>
    </Transition>

    <!-- Carte résumé après sélection -->
    <Transition name="card-fade">
      <div
        v-if="selectedPharmacie"
        class="max-w-md rounded-lg border border-slate-200 dark:border-slate-700
               bg-slate-50 dark:bg-slate-800/60 overflow-hidden"
      >
        <!-- En-tête carte -->
        <div class="flex items-center justify-between px-3 py-2
                    border-b border-slate-200 dark:border-slate-700">
          <div class="flex items-center gap-2">
            <Pill :size="14" class="text-teal-600 dark:text-teal-400 flex-shrink-0" />
            <span class="text-sm font-bold text-slate-900 dark:text-white truncate">
              {{ selectedPharmacie.NomOrganisation }}
            </span>
          </div>
          <img
            v-if="getBanniereLogo(selectedPharmacie.Banniere)"
            :src="getBanniereLogo(selectedPharmacie.Banniere)"
            :alt="selectedPharmacie.Banniere"
            class="h-6 object-contain flex-shrink-0"
          />
          <span
            v-else-if="selectedPharmacie.Banniere"
            class="text-[10px] px-1.5 py-0.5 bg-white dark:bg-slate-700
                   border border-slate-200 dark:border-slate-600
                   text-slate-600 dark:text-slate-300 rounded"
          >{{ selectedPharmacie.Banniere }}</span>
        </div>

        <!-- Détails carte -->
        <div class="px-3 py-2 space-y-1.5 text-xs text-slate-600 dark:text-slate-400">
          <!-- Adresse -->
          <div v-if="selectedPharmacie.Adresse" class="flex items-start gap-2">
            <MapPin :size="12" class="mt-0.5 flex-shrink-0 text-slate-400" />
            <span>
              {{ selectedPharmacie.Adresse }}<br/>
              {{ selectedPharmacie.Ville }}
              <span v-if="selectedPharmacie.Region"> — {{ selectedPharmacie.Region }}</span>
            </span>
          </div>
          <!-- Téléphone cliquable -->
          <div v-if="selectedPharmacie.Tel" class="flex items-center gap-2">
            <Phone :size="12" class="flex-shrink-0 text-slate-400" />
            <a
              :href="`tel:${selectedPharmacie.Tel}`"
              class="hover:text-teal-600 dark:hover:text-teal-400 transition-colors"
            >{{ formatPhone(selectedPharmacie.Tel) }}</a>
          </div>
          <!-- Fax -->
          <div v-if="selectedPharmacie.Fax" class="flex items-center gap-2">
            <Printer :size="12" class="flex-shrink-0 text-slate-400" />
            <span>{{ formatPhone(selectedPharmacie.Fax) }}</span>
          </div>
        </div>

        <!-- Lien fiche -->
        <div class="px-3 py-2 border-t border-slate-200 dark:border-slate-700">
          <button
            type="button"
            @click="$emit('view-pharmacie', selectedPharmacie)"
            class="flex items-center gap-1.5 text-[11px] font-medium
                   text-teal-600 dark:text-teal-400 hover:underline transition-colors"
          >
            <ExternalLink :size="11" />
            Voir la fiche pharmacie
          </button>
        </div>
      </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Search, MapPin, Phone, Printer, X, CheckCircle, ExternalLink, Pill } from 'lucide-vue-next'
import { GetAllPharmacies } from '../../../wailsjs/go/main/App'
import JeanCoutu    from '../../assets/logoPharma/jean-coutu.png'
import Pharmaprix   from '../../assets/logoPharma/pharmaprix.png'
import Familiprix   from '../../assets/logoPharma/familiprix.png'
import Uniprix      from '../../assets/logoPharma/uniprix.png'
import Proxim       from '../../assets/logoPharma/proxim.png'
import Brunet       from '../../assets/logoPharma/brunet.png'
import Independante from '../../assets/logoPharma/independante.png'

const banniereLogos = {
  'Jean Coutu':   JeanCoutu,
  'Pharmaprix':   Pharmaprix,
  'Familiprix':   Familiprix,
  'Uniprix':      Uniprix,
  'Proxim':       Proxim,
  'Brunet':       Brunet,
  'Indépendante': Independante,
}

const getBanniereLogo = (b) => banniereLogos[b] || null

const formatPhone = (phone) => {
  if (!phone) return ''
  const cleaned = String(phone).replace(/\D/g, '')
  if (cleaned.length === 10) return `(${cleaned.slice(0,3)}) ${cleaned.slice(3,6)}-${cleaned.slice(6)}`
  if (cleaned.length === 11 && cleaned[0] === '1') return `1 (${cleaned.slice(1,4)}) ${cleaned.slice(4,7)}-${cleaned.slice(7)}`
  return phone
}

const props = defineProps(['modelValue', 'label'])
const emit  = defineEmits(['update:modelValue', 'change', 'view-pharmacie'])

const pharmacies      = ref([])
const searchQuery     = ref('')
const showDropdown    = ref(false)
const showToast       = ref(false)
const selectedId      = ref(props.modelValue)
const selectedPharmacie = ref(null)
let toastTimer = null

const filteredPharmacies = computed(() => {
  const q = searchQuery.value.toLowerCase()
  const list = q
    ? pharmacies.value.filter(p =>
        p.NomOrganisation.toLowerCase().includes(q) ||
        p.Ville?.toLowerCase().includes(q) ||
        p.Banniere?.toLowerCase().includes(q)
      )
    : pharmacies.value.slice(0, 6)
  return list
})

const selectPharmacie = (pharma) => {
  selectedId.value        = pharma.ID
  selectedPharmacie.value = pharma
  searchQuery.value       = pharma.NomOrganisation
  showDropdown.value      = false
  emit('update:modelValue', pharma.ID)
  emit('change', pharma)
  // Toast
  showToast.value = true
  clearTimeout(toastTimer)
  toastTimer = setTimeout(() => { showToast.value = false }, 3000)
}

const clearSelection = () => {
  selectedId.value        = null
  selectedPharmacie.value = null
  searchQuery.value       = ''
  showToast.value         = false
  emit('update:modelValue', null)
}

const onBlur = () => {
  setTimeout(() => { showDropdown.value = false }, 150)
}

onMounted(async () => {
  try {
    const data = await GetAllPharmacies()
    pharmacies.value = data || []
    if (props.modelValue) {
      const p = pharmacies.value.find(x => x.ID === props.modelValue)
      if (p) {
        selectedPharmacie.value = p
        searchQuery.value       = p.NomOrganisation
        selectedId.value        = p.ID
      }
    }
  } catch (err) {
    console.error('Erreur chargement sélecteur pharmacie:', err)
  }
})
</script>

<style scoped>
.toast-slide-enter-active,
.toast-slide-leave-active { transition: all 0.25s ease; }
.toast-slide-enter-from   { opacity: 0; transform: translateY(-6px); }
.toast-slide-leave-to     { opacity: 0; transform: translateY(-6px); }

.card-fade-enter-active,
.card-fade-leave-active { transition: all 0.2s ease; }
.card-fade-enter-from,
.card-fade-leave-to     { opacity: 0; transform: translateY(4px); }

.fk-label {
  display: block;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: #6b7280;
  margin-bottom: 4px;
}
.dark .fk-label { color: #9ca3af; }
</style>