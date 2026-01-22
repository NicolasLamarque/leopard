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
                  <option value="Famille">Famille</option>
                  <option value="Ami">Ami</option>
                  <option value="Professionnel">Professionnel</option>
                  <option value="Voisin">Voisin</option>
                  <option value="Autre">Autre</option>
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
                  <option value="Conjoint(e)">Conjoint(e)</option>
                  <option value="Fils">Fils</option>
                  <option value="Fille">Fille</option>
                  <option value="Père">Père</option>
                  <option value="Mère">Mère</option>
                  <option value="Frère">Frère</option>
                  <option value="Sœur">Sœur</option>
                  <option value="Petit-fils">Petit-fils</option>
                  <option value="Petite-fille">Petite-fille</option>
                  <option value="Autre">Autre</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Type de contact
                </label>
                <input 
                  v-model="form.type_de_contact"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-teal-700 focus:ring-4 focus:ring-teal-500/20 transition-all outline-none"
                  placeholder="Ex: Personne de confiance"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Force du lien (0-10)
                </label>
                <input 
                  v-model.number="form.force_lien"
                  type="range"
                  min="0"
                  max="10"
                  class="w-full h-2 bg-slate-400 rounded-lg appearance-none cursor-pointer accent-teal-600 dark:accent-teal-600 focus:ring-2 focus:ring-teal-500/20 transition-all outline-none"
                />
                <div class="flex justify-between text-xs text-gray-500 dark:text-gray-400 mt-1">
                  <span>Faible</span>
                  <span class="font-semibold text-stone-600 dark:text-stone-400">{{ form.force_lien }}/10</span>
                  <span>Fort</span>
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
import { ref, computed, watch } from 'vue'
import { 
  X, UserPlus, User, Heart, Shield, Phone, MapPin, Briefcase,
  Save, Loader2, AlertCircle
} from 'lucide-vue-next'
import { CreateContact, UpdateContact } from '../../../wailsjs/go/main/App'

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
  force_lien: 5,
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
</script>

<style scoped>
@keyframes slideIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

.animate-slideIn {
  animation: slideIn 0.3s ease-out;
}
</style>