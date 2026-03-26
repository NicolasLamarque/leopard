<template>
  <div class="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center z-[60] p-4">
    <div class="bg-white dark:bg-gray-900 rounded-2xl shadow-2xl w-full max-w-3xl max-h-[90vh] overflow-hidden flex flex-col animate-slideIn">
      
      <!-- Header -->
      <div class="bg-gradient-to-r from-slate-600 to-teal-800 dark:from-slate-500 dark:to-teal-700 px-8 py-6 flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-white/20 rounded-xl backdrop-blur-sm">
            <UserPlus :size="28" class="text-white" />
          </div>
          <div>
            <h2 class="text-2xl font-bold text-white">
              {{ isEdit ? 'Modifier le contact' : 'Nouveau contact' }}
            </h2>
            <p class="text-purple-100 text-sm mt-0.5">
              {{ isEdit ? 'Modifier les informations' : 'Ajouter une personne de confiance' }}
            </p>
          </div>
        </div>
        <button 
          @click="$emit('close')" 
          class="p-2 hover:bg-white/20 rounded-lg transition-colors group"
          :disabled="loading"
        >
          <X :size="24" class="text-white group-hover:rotate-90 transition-transform" />
        </button>
      </div>

      <!-- Form Content -->
      <div class="flex-1 overflow-y-auto px-8 py-6">
        <form @submit.prevent="save" class="space-y-6">
          
          <!-- Identité -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
              <User :size="20" class="text-teal-700 dark:text-teal-700" />
              Identité
            </h3>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Nom <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="form.nom"
                  required
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Nom de famille"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Prénom <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="form.prenom"
                  required
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Prénom"
                />
              </div>
            </div>
          </div>

          <!-- Relation -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
              <Heart :size="20" class="text-pink-600 dark:text-pink-700" />
              Type de relation
            </h3>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
<div>
  <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
    Type de relation <span class="text-red-500">*</span>
  </label>
  <select 
    v-model="form.relation_type"
    required
    class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
  >
    <option value="">Sélectionner...</option>
    <option v-for="type in typesRelation" :key="type" :value="type">
      {{ type }}
    </option>
  </select>
</div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Lien familial
                </label>
                <select 
  v-model="form.lien_familial"
  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
>
  <option value="">Aucun</option>
  <option v-for="lien in liensFamiliaux" :key="lien" :value="lien">
    {{ lien }}
  </option>
</select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Type de contact
                </label>
                <select 
  v-model="form.type_de_contact"
  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
>
  <option value="">Sélectionner...</option>
  <option v-for="type in typesContact" :key="type" :value="type">
    {{ type }}
  </option>
</select>
              </div>

              <!-- Force du lien -10 à +10 — positionné par le client -->
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                  Force du lien affectif
                  <span class="ml-1 text-xs font-normal text-gray-400">— évalué par le client</span>
                </label>
                <input
                  v-model.number="form.force_lien"
                  type="range" min="-10" max="10" step="1"
                  class="w-full h-2 rounded-lg appearance-none cursor-pointer transition-all outline-none slider-lien"
                  :style="styleLien"
                />
                <div class="flex justify-between text-xs mt-1">
                  <span class="text-red-500">-10 Lien rompu</span>
                  <span :class="['font-bold', form.force_lien < -5 ? 'text-red-600' : form.force_lien < 0 ? 'text-orange-500' : form.force_lien === 0 ? 'text-slate-400' : 'text-teal-600']">
                    {{ form.force_lien > 0 ? '+' : '' }}{{ form.force_lien }}
                    <span class="font-normal ml-1 text-gray-500 dark:text-gray-400">{{ labelForceLien(form.force_lien) }}</span>
                  </span>
                  <span class="text-teal-600">+10 Lien pivot</span>
                </div>
              </div>

              <!-- Indicateur maltraitance 0-5 -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                  Indicateur maltraitance
                  <span class="ml-1 text-xs font-normal text-gray-400">(0–5)</span>
                </label>
                <input
                  v-model.number="form.force_niv_maltraitance"
                  type="range" min="0" max="5" step="1"
                  class="w-full h-2 rounded-lg appearance-none cursor-pointer transition-all outline-none slider-gradient"
                  :style="styleMalt"
                />
                <div class="flex justify-between text-xs mt-1">
                  <span class="text-slate-400">0 Aucun</span>
                  <span :class="['font-semibold', form.force_niv_maltraitance === 0 ? 'text-slate-400' : form.force_niv_maltraitance <= 2 ? 'text-amber-600' : 'text-red-600']">
                    {{ form.force_niv_maltraitance }} — {{ labelMaltraitance(form.force_niv_maltraitance) }}
                  </span>
                  <span class="text-red-500">5 Grave</span>
                </div>
              </div>

              <!-- Niveau soutien 0-5 -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                  Soutien apporté
                  <span class="ml-1 text-xs font-normal text-gray-400">(0–5)</span>
                </label>
                <input
                  v-model.number="form.force_niv_soutien"
                  type="range" min="0" max="5" step="1"
                  class="w-full h-2 rounded-lg appearance-none cursor-pointer transition-all outline-none slider-gradient"
                  :style="styleSoutien"
                />
                <div class="flex justify-between text-xs mt-1">
                  <span class="text-slate-400">0 Aucun</span>
                  <span class="font-semibold text-blue-600">{{ form.force_niv_soutien }} — {{ labelSoutien(form.force_niv_soutien) }}</span>
                  <span class="text-blue-600">5 Total</span>
                </div>
              </div>

              <!-- Niveau épuisement 0-5 -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                  Épuisement observé
                  <span class="ml-1 text-xs font-normal text-gray-400">(0–5)</span>
                </label>
                <input
                  v-model.number="form.force_niv_epuisement"
                  type="range" min="0" max="5" step="1"
                  class="w-full h-2 rounded-lg appearance-none cursor-pointer transition-all outline-none slider-gradient"
                  :style="styleEpuisement"
                />
                <div class="flex justify-between text-xs mt-1">
                  <span class="text-slate-400">0 Aucun</span>
                  <span :class="['font-semibold', form.force_niv_epuisement === 0 ? 'text-slate-400' : form.force_niv_epuisement <= 2 ? 'text-amber-600' : 'text-orange-600']">
                    {{ form.force_niv_epuisement }} — {{ labelEpuisement(form.force_niv_epuisement) }}
                  </span>
                  <span class="text-orange-600">5 Effondrement</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Options spéciales -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
              <Shield :size="20" class="text-teal-600 dark:text-teal-700" />
              Rôles spéciaux
            </h3>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
              <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-xl cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                <input 
                  v-model="form.contact_urgence"
                  type="checkbox"
                  :true-value="1"
                  :false-value="0"
                  class="w-5 h-5 text-red-600 border-gray-300 rounded accent-teal-600"
                />
                <div>
                  <span class="text-sm font-medium text-gray-700 dark:text-gray-300 block">Contact d'urgence</span>
                  <span class="text-xs text-gray-500 dark:text-gray-400">À contacter en priorité</span>
                </div>
              </label>

              <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-xl cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                <input 
                  v-model="form.aidant_naturel"
                  type="checkbox"
                  :true-value="1"
                  :false-value="0"
                  class="w-5 h-5 text-green-600 border-gray-300 rounded accent-teal-600"
                />
                <div>
                  <span class="text-sm font-medium text-gray-700 dark:text-gray-300 block">Aidant naturel</span>
                  <span class="text-xs text-gray-500 dark:text-gray-400">Aide régulière</span>
                </div>
              </label>

              <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-xl cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                <input 
                  v-model="form.procuration_bancaire"
                  type="checkbox"
                  :true-value="1"
                  :false-value="0"
                  class="w-5 h-5 text-blue-600 border-gray-300 rounded accent-teal-600"
                />
                <div>
                  <span class="text-sm font-medium text-gray-700 dark:text-gray-300 block">Procuration bancaire</span>
                  <span class="text-xs text-gray-500 dark:text-gray-400">Gestion finances</span>
                </div>
              </label>

              <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-xl cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                <input 
                  v-model="form.procuration_notariee"
                  type="checkbox"
                  :true-value="1"
                  :false-value="0"
                  class="w-5 h-5 text-purple-600 border-gray-300 rounded accent-teal-600"
                />
                <div>
                  <span class="text-sm font-medium text-gray-700 dark:text-gray-300 block">Procuration notariée</span>
                  <span class="text-xs text-gray-500 dark:text-gray-400">Mandat légal</span>
                </div>
              </label>
            </div>
          </div>

          <!-- Coordonnées -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
              <Phone :size="20" class="text-teal-600 dark:text-teal-700" />
              Coordonnées
            </h3>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Téléphone
                </label>
                <input 
                  v-model="form.telephone"
                  type="tel"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="819-555-0123"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Cellulaire
                </label>
                <input 
                  v-model="form.cellulaire"
                  type="tel"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="819-555-0456"
                />
              </div>

              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Courriel
                </label>
                <input 
                  v-model="form.email"
                  type="email"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="exemple@email.com"
                />
              </div>
            </div>
          </div>

          <!-- Adresse -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
              <MapPin :size="20" class="text-orange-600 dark:text-sky-700" />
              Adresse
            </h3>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Adresse
                </label>
                <input 
                  v-model="form.adresse"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="123 Rue Principale"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Ville
                </label>
                <input 
                  v-model="form.ville"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Gatineau"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Code postal
                </label>
                <input 
                  v-model="form.code_postal"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="J8V 1A1"
                />
              </div>

              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Pays
                </label>
                <input 
                  v-model="form.pays"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Canada"
                />
              </div>
            </div>
          </div>

          <!-- Profession -->
          <div class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
              <Briefcase :size="20" class="text-indigo-600 dark:text-indigo-400" />
              Informations professionnelles
            </h3>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Profession
                </label>
                <input 
                  v-model="form.profession"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Ex: Infirmière"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Employeur
                </label>
                <input 
                  v-model="form.employeur"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Nom de l'employeur"
                />
              </div>
            </div>
          </div>

          <!-- Notes -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              Note fixe
            </label>
            <textarea 
              v-model="form.note_fixe"
              rows="3"
              class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none resize-none"
              placeholder="Notes importantes..."
            ></textarea>
          </div>

        </form>
      </div>

      <!-- Footer -->
      <div class="border-t-2 border-gray-200 dark:border-gray-800 px-8 py-5 bg-gray-50 dark:bg-gray-900/50">
        <div class="flex justify-end gap-3">
          <button 
            type="button"
            @click="$emit('close')"
            :disabled="loading"
            class="px-6 py-3 border-2 border-gray-300 dark:border-gray-700 text-gray-700 dark:text-gray-300 rounded-xl hover:bg-gray-100 dark:hover:bg-gray-800 transition-all font-medium disabled:opacity-50"
          >
            Annuler
          </button>
          <button 
            @click="save"
            :disabled="loading || !canSubmit"
            class="px-8 py-3 bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 text-white rounded-xl font-semibold shadow-lg hover:shadow-xl transition-all disabled:opacity-50 flex items-center gap-2"
          >
            <Loader2 v-if="loading" :size="20" class="animate-spin" />
            <Save v-else :size="20" />
            <span>{{ loading ? 'Enregistrement...' : (isEdit ? 'Modifier' : 'Ajouter') }}</span>
          </button>
        </div>

        <div v-if="error" class="mt-4 p-4 bg-red-50 dark:bg-red-900/20 border-2 border-red-200 dark:border-red-800 rounded-xl flex items-start gap-3">
          <AlertCircle :size="20" class="text-red-600 dark:text-red-400 flex-shrink-0 mt-0.5" />
          <div>
            <p class="text-sm font-medium text-red-900 dark:text-red-100">Erreur</p>
            <p class="text-sm text-red-700 dark:text-red-300 mt-1">{{ error }}</p>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { 
  X, UserPlus, User, Heart, Shield, Phone, MapPin, Briefcase,
  Save, Loader2, AlertCircle
} from 'lucide-vue-next'
import { CreateContact, UpdateContact } from '../../../wailsjs/go/main/App'
import { useRefListes } from '@/composables/useRefListes' 


const { getCategorie } = useRefListes()
const typesRelation = ref([])   // type_relation  → grande catégorie (Famille, Légal, Social...)
const typesContact  = ref([])   // type_contact   → rôle précis (Mandataire, Conjoint, Aidant...)
const liensFamiliaux = ref([])

onMounted(async () => {
  const tr = await getCategorie('type_relation')
  typesRelation.value = tr.map(t => t.Libelle)

  const tc = await getCategorie('type_contact')
  typesContact.value = tc.map(t => t.Libelle)

  const lf = await getCategorie('lien_familial')
  liensFamiliaux.value = lf.map(t => t.Libelle)
})

const props = defineProps({
  clientId: { type: Number, required: true },
  contactData: { type: Object, default: null }
})


const emit = defineEmits(['close', 'saved'])

const loading = ref(false)
const error = ref('')

const isEdit = computed(() => !!props.contactData)

const form = ref({
  nom: '',
  prenom: '',
  telephone: null,
  cellulaire: null,
  adresse: null,
  code_postal: null,
  ville: null,
  pays: 'Canada',
  email: null,
  employeur: null,
  profession: null,
  relation_client: null,
  lien_familial: null,
  force_lien: 0,
  force_niv_maltraitance: 0,
  force_niv_soutien: 0,
  force_niv_epuisement: 0,
  contact_urgence: 0,
  aidant_naturel: 0,
  type_de_contact: null,
  procuration_bancaire: 0,
  procuration_notariee: 0,
  note_fixe: null,
  relation_type: '',
  client_id: props.clientId
})

const canSubmit = computed(() => {
  return form.value.nom && form.value.prenom && form.value.relation_type
})

const save = async () => {
  loading.value = true
  error.value = ''

  try {
    if (isEdit.value) {
      await UpdateContact({ ...form.value, id: props.contactData.id })
    } else {
      await CreateContact(form.value)
    }
    
    emit('saved')
  } catch (e) {
    console.error('Erreur sauvegarde contact:', e)
    error.value = e.message || 'Erreur lors de la sauvegarde'
  } finally {
    loading.value = false
  }
}

watch(() => props.contactData, (newVal) => {
  if (newVal) {
    form.value = { ...newVal, client_id: props.clientId }
  }
}, { immediate: true })


// ─── Styles dynamiques des sliders (track coloré) ────────────────────────────
// Tailwind V4 ne colore pas le track — on utilise background inline

const styleLien = computed(() => {
  const v = form.value.force_lien  // -10 à +10
  const pct = ((v + 10) / 20) * 100  // 0% à 100%

  if (v < 0) {
    // Rouge → Orange → Gris au centre
    return {
      background: `linear-gradient(to right, 
        #dc2626 0%, 
        #f97316 ${pct * 0.6}%, 
        #94a3b8 ${pct}%, 
        #e2e8f0 ${pct}%, 
        #e2e8f0 100%)`,
      accentColor: '#dc2626'
    }
  } else if (v === 0) {
    return {
      background: '#94a3b8',
      accentColor: '#94a3b8'
    }
  } else {
    // Gris → Teal → Vert foncé
    return {
      background: `linear-gradient(to right,
        #e2e8f0 0%,
        #e2e8f0 50%,
        #0d9488 ${50 + (pct - 50)}%,
        #0f6e56 100%)`,
      accentColor: '#0d9488'
    }
  }
})

const styleMalt = computed(() => {
  // Vert (0) → jaune → rouge (5)
  const v = form.value.force_niv_maltraitance
  const pct = (v / 5) * 100
  const couleur = v === 0 ? '#94a3b8' : v <= 2 ? '#f59e0b' : '#ef4444'
  return {
    background: `linear-gradient(to right, #22c55e 0%, #eab308 50%, #ef4444 100%)`,
    accentColor: couleur
  }
})

const styleSoutien = computed(() => {
  // Gris (0) → bleu (5)
  const pct = (form.value.force_niv_soutien / 5) * 100
  return {
    background: `linear-gradient(to right, #3b82f6 0% ${pct}%, #e2e8f0 ${pct}% 100%)`,
    accentColor: '#3b82f6'
  }
})

const styleEpuisement = computed(() => {
  // Vert (0) → amber → orange (5)
  const v = form.value.force_niv_epuisement
  const couleur = v === 0 ? '#94a3b8' : v <= 2 ? '#f59e0b' : '#ea580c'
  return {
    background: `linear-gradient(to right, #22c55e 0%, #f59e0b 60%, #ea580c 100%)`,
    accentColor: couleur
  }
})

// ─── Labels des échelles ─────────────────────────────────────────────────────

const labelForceLien = (v) => {
  if (v <= -8) return 'Lien rompu / rupture grave'
  if (v <= -5) return 'Conflit sévère actif'
  if (v <= -3) return 'Relation toxique / ambivalence'
  if (v <= -1) return 'Lien fragilisé / distant'
  if (v === 0) return 'Neutre / inconnu'
  if (v <= 3)  return 'Lien présent'
  if (v <= 5)  return 'Lien significatif'
  if (v <= 7)  return 'Lien fort'
  return 'Lien pivot / attachement profond'
}

const labelMaltraitance = (v) => {
  const labels = [
    'Aucun indicateur',
    'Indicateurs suspectés (tiers)',
    'Indicateurs détectés',
    'Indicateurs admis',
    'Maltraitance involontaire avérée',
    'Maltraitance grave avérée'
  ]
  return labels[v] || ''
}

const labelSoutien = (v) => {
  const labels = [
    'Aucun soutien observable',
    'Présence occasionnelle',
    'Soutien régulier (AVQ, transport)',
    'Soutien significatif / pivot informel',
    'Aidant principal engagé',
    'Présence et coordination totales'
  ]
  return labels[v] || ''
}

const labelEpuisement = (v) => {
  const labels = [
    'Aucun signe',
    'Signes légers (fatigue, commentaires)',
    'Épuisement modéré (isolement)',
    'Épuisement élevé (santé affectée)',
    'Risque d\'effondrement',
    'Effondrement / retrait de l\'aide'
  ]
  return labels[v] || ''
}
</script>

<style scoped>
@keyframes slideIn {
  from { opacity: 0; transform: scale(0.95) translateY(20px); }
  to   { opacity: 1; transform: scale(1) translateY(0); }
}
.animate-slideIn { animation: slideIn 0.3s ease-out; }

/* Track + thumb des sliders */
.slider-lien, .slider-gradient {
  -webkit-appearance: none;
  appearance: none;
  height: 6px;
  border-radius: 9999px;
}
.slider-lien::-webkit-slider-thumb,
.slider-gradient::-webkit-slider-thumb {
  -webkit-appearance: none;
  width: 18px; height: 18px;
  border-radius: 50%;
  background: white;
  border: 2px solid #0d9488;
  cursor: pointer;
  box-shadow: 0 1px 3px rgba(0,0,0,0.3);
}
</style>