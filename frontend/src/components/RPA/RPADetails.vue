<template>
  <div @click.self="$emit('close')" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4 backdrop-blur-sm">
    <div class="bg-white dark:bg-gray-800 rounded-xl shadow-2xl max-w-4xl w-full max-h-[92vh] overflow-hidden flex flex-col relative">
      
      <div v-if="isUpdating" class="absolute inset-0 bg-white/80 dark:bg-gray-800/80 z-[60] flex flex-col items-center justify-center">
        <RefreshCw :size="48" class="text-blue-600 animate-spin mb-4" />
        <p class="text-lg font-bold text-gray-900 dark:text-white">Synchronisation en direct...</p>
        <p class="text-sm text-gray-500">Récupération des données du MSSS Québec</p>
      </div>

      <div class="px-6 py-4 border-b dark:border-gray-700 bg-gray-50/50 dark:bg-gray-900/20 flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg text-blue-600 dark:text-blue-400">
            <Building2 :size="24" />
          </div>
          <div>
            <div class="flex items-center gap-2">
              <h2 class="text-xl font-bold text-gray-900 dark:text-white">{{ title }}</h2>
              <a v-if="form.source_url" :href="form.source_url" target="_blank" class="p-1 text-gray-400 hover:text-blue-500 transition-colors" title="Fiche officielle du MSSS">
                <ExternalLink :size="18" />
              </a>
            </div>
            <p v-if="mode !== 'create'" class="text-sm font-mono text-gray-500 dark:text-gray-400">#{{ form.registre }}</p>
          </div>
        </div>

        <div class="flex items-center gap-2">
          <button v-if="mode === 'view'" @click="mode = 'edit'" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
            <Edit :size="18" /> Modifier
          </button>
          <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 p-2">
            <X :size="24" />
          </button>
        </div>
      </div>

      <div class="flex-1 overflow-y-auto p-6 space-y-8">
        
        <div v-if="form.statut === 'ferme'" class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 p-4 rounded-lg flex items-center gap-3 text-red-700 dark:text-red-400">
          <Trash2 :size="20" />
          <span>Cette résidence est fermée depuis le {{ formatDate(form.date_fermeture) }}</span>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="md:col-span-2">
            <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Nom du RPA *</label>
            <input v-model="form.titre_rpa" :disabled="mode === 'view'" required class="w-full border dark:border-gray-600 rounded-lg px-4 py-3 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50 dark:disabled:bg-gray-900/50 font-medium" />
          </div>

          <div>
            <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Numéro de registre</label>
            <input v-model="form.registre" disabled class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-gray-50 dark:bg-gray-900 text-gray-500 font-mono" />
          </div>

          <div>
            <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Statut *</label>
            <select v-model="form.statut" :disabled="mode === 'view'" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50">
              <option value="actif">Actif</option>
              <option value="ferme">Fermé</option>
              <option value="suspendu">Suspendu</option>
            </select>
          </div>
        </div>

        <div class="bg-blue-50/50 dark:bg-blue-900/10 rounded-xl border border-blue-100 dark:border-blue-800/30 p-5">
          <h3 class="text-sm font-bold text-blue-800 dark:text-blue-300 uppercase tracking-wider mb-3 flex items-center gap-2">
            <ClipboardList :size="18" /> 7 - Services offerts par la résidence
          </h3>
          <div v-if="mode === 'view'" class="text-gray-700 dark:text-gray-300 text-sm leading-relaxed whitespace-pre-wrap italic">
            {{ form.services || 'Aucune information détaillée n\'a été récupérée.' }}
          </div>
          <textarea v-else v-model="form.services" rows="4" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 resize-none"></textarea>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="space-y-4">
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Adresse</label>
              <textarea v-model="form.adresse" :disabled="mode === 'view'" rows="2" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50 resize-none"></textarea>
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Ville</label>
                <input v-model="form.ville" :disabled="mode === 'view'" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white disabled:bg-gray-50" />
              </div>
              <div>
                <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Code Postal</label>
                <input v-model="form.code_postal" :disabled="mode === 'view'" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white disabled:bg-gray-50 font-mono" />
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Téléphone</label>
              <div class="relative">
                <Phone :size="16" class="absolute left-3 top-3 text-gray-400" />
                <input v-model="form.telephone" :disabled="mode === 'view'" type="tel" class="w-full border dark:border-gray-600 rounded-lg pl-10 pr-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50" />
              </div>
            </div>
            <div>
              <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Capacité (places)</label>
              <input v-model.number="form.capacite" :disabled="mode === 'view'" type="number" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50" />
            </div>
          </div>
        </div>

        <div>
          <label class="block text-xs font-bold text-gray-500 uppercase mb-2">Notes internes</label>
          <textarea v-model="form.notes" :disabled="mode === 'view'" rows="3" placeholder="Notes privées..." class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 disabled:bg-gray-50 resize-none"></textarea>
        </div>

        <div v-if="mode !== 'create'" class="pt-6 border-t dark:border-gray-700 grid grid-cols-2 md:grid-cols-4 gap-4 text-[11px] uppercase tracking-wider font-bold text-gray-400">
          <div>
            <p>Ajouté</p>
            <p class="text-gray-900 dark:text-gray-200 mt-1">{{ formatDate(form.date_ajout) }}</p>
          </div>
          <div>
            <p>Dernière MAJ</p>
            <p class="text-blue-600 dark:text-blue-400 mt-1">{{ formatDate(form.date_maj) }}</p>
          </div>
          <div>
            <p>Vérification</p>
            <p class="text-gray-900 dark:text-gray-200 mt-1">{{ formatDate(form.derniere_verification) }}</p>
          </div>
          <div>
            <p>Certification</p>
            <p class="text-gray-900 dark:text-gray-200 mt-1">{{ form.date_certification || 'N/A' }}</p>
          </div>
        </div>
      </div>

      <div class="px-6 py-4 border-t dark:border-gray-700 flex items-center justify-between bg-gray-50 dark:bg-gray-900/50">
        <button v-if="mode !== 'create'" @click="handleDelete" class="px-4 py-2 bg-red-50 text-red-600 hover:bg-red-600 hover:text-white rounded-lg transition-all flex items-center gap-2 font-bold text-sm">
          <Trash2 :size="18" /> Supprimer
        </button>
        <div v-else></div>

        <div class="flex items-center gap-3">
          <template v-if="mode === 'edit' || mode === 'create'">
            <button @click="handleCancel" class="px-4 py-2 text-gray-700 dark:text-gray-300 font-bold">Annuler</button>
            <button @click="handleSave" :disabled="!isValid" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 font-bold flex items-center gap-2">
              <Save :size="18" /> {{ mode === 'create' ? 'Créer' : 'Sauvegarder' }}
            </button>
          </template>
          <button v-else @click="$emit('close')" class="px-6 py-2 bg-gray-900 dark:bg-gray-700 text-white rounded-lg font-bold">Fermer</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Building2, X, Edit, Save, Trash2, ExternalLink, RefreshCw, MapPin, Phone, ClipboardList } from 'lucide-vue-next'
import { GetResidenceForDetails } from '../../../wailsjs/go/main/App'

const props = defineProps({
  rpa: { type: Object, default: null },
  mode: { type: String, default: 'view' }
})

const emit = defineEmits(['close', 'save', 'delete'])
const mode = ref(props.mode)
const isUpdating = ref(false)

// Ton formulaire original étendu
const form = ref({
  id: null,
  region: '',
  registre: '',
  titre_rpa: '',
  municipalite: '',
  ville: '',
  code_postal: '',
  adresse: '',
  telephone: '',
  capacite: null,
  type_resid: '',
  proprietaires: '',
  services: '',
  statut: 'actif',
  notes: '',
  date_certification: '',
  date_ajout: null,
  date_maj: null,
  date_fermeture: null,
  derniere_verification: null,
  source_url: ''
})

const initialForm = ref({})

if (props.rpa) {
  form.value = { ...props.rpa }
  initialForm.value = { ...props.rpa }
}

onMounted(async () => {
  if (mode.value === 'view' && form.value.registre) {
    isUpdating.value = true
    try {
      const freshData = await GetResidenceForDetails(form.value.registre, true)
      if (freshData) {
        form.value = { ...freshData }
        initialForm.value = { ...freshData }
      }
    } catch (err) {
      console.error("Synchro error:", err)
    } finally {
      isUpdating.value = false
    }
  }
})

const title = computed(() => mode.value === 'create' ? 'Nouveau RPA' : mode.value === 'edit' ? 'Modifier le RPA' : 'Détails du RPA')
const isValid = computed(() => form.value.titre_rpa && form.value.registre && form.value.statut)

const handleSave = () => emit('save', { ...form.value })
const handleCancel = () => {
  if (mode.value === 'create') emit('close')
  else {
    form.value = { ...initialForm.value }
    mode.value = 'view'
  }
}
const handleDelete = () => {
  if (confirm(`Supprimer "${form.value.titre_rpa}" ?`)) emit('delete', form.value)
}

const formatDate = (d) => d ? new Date(d).toLocaleDateString('fr-CA') : 'N/A'
const getStatutClass = (s) => {
  const c = { 'actif': 'bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400', 'ferme': 'bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400' }
  return c[s] || 'bg-gray-100 text-gray-800'
}
</script>