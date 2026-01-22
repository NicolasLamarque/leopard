<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <!-- Overlay -->
    <div class="fixed inset-0 bg-black bg-opacity-50 transition-opacity" @click="handleClose"></div>

    <!-- Modal -->
    <div class="flex min-h-full items-center justify-center p-4">
      <div class="relative bg-white dark:bg-gray-800 rounded-xl shadow-2xl max-w-5xl w-full max-h-[90vh] overflow-hidden">
        
        <!-- Header -->
        <div class="sticky top-0 z-10 bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 px-6 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <Building2 :size="28" class="text-blue-600 dark:text-blue-400" />
              <div>
                <h2 class="text-xl font-bold text-gray-900 dark:text-white">
                  {{ title }}
                </h2>
                <p class="text-sm text-gray-500 dark:text-gray-400">
                  {{ subtitle }}
                </p>
              </div>
            </div>

            <div class="flex items-center gap-2">
              <button
                v-if="mode === 'view'"
                @click="mode = 'edit'"
                class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
              >
                <Edit :size="16" />
                Modifier
              </button>

              <button
                @click="handleClose"
                class="p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-colors"
              >
                <X :size="20" class="text-gray-500 dark:text-gray-400" />
              </button>
            </div>
          </div>
        </div>

        <!-- Content -->
        <div class="overflow-y-auto max-h-[calc(90vh-140px)] px-6 py-6">
          <FormKit
            type="form"
            v-model="formData"
            @submit="handleSubmit"
            :actions="false"
          >
            <!-- Section Identification -->
            <div class="mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
                <FileText :size="20" />
                Identification
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <FormKit
                  type="text"
                  name="TitreCHSLD"
                  label="Nom du CHSLD"
                  validation="required"
                  :disabled="mode === 'view'"
                  placeholder="Ex: CHSLD de Montréal"
                  outer-class="md:col-span-2"
                />

                <FormKit
                  type="select"
                  name="TypeCHSLD"
                  label="Type de CHSLD"
                  :options="typeCHSLDOptions"
                  :disabled="mode === 'view'"
                />

                <FormKit
                  type="select"
                  name="Statut"
                  label="Statut"
                  :options="statutOptions"
                  :disabled="mode === 'view'"
                />

                <FormKit
                  type="text"
                  name="DateCertification"
                  label="Date de certification"
                  :disabled="mode === 'view'"
                  placeholder="AAAA-MM-JJ"
                />

                <FormKit
                  type="number"
                  name="Capacite"
                  label="Capacité (lits)"
                  validation="number|min:0"
                  :disabled="mode === 'view'"
                />
              </div>
            </div>

            <!-- Section Localisation -->
            <div class="mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
                <MapPin :size="20" />
                Localisation
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <FormKit
                  type="select"
                  name="Region"
                  label="Région administrative"
                  :options="regionOptions"
                  :disabled="mode === 'view'"
                />

                <FormKit
                  type="text"
                  name="Municipalite"
                  label="Municipalité"
                  validation="required"
                  :disabled="mode === 'view'"
                />

                <FormKit
                  type="text"
                  name="Adresse"
                  label="Adresse"
                  :disabled="mode === 'view'"
                  outer-class="md:col-span-2"
                />

                <FormKit
                  type="text"
                  name="CodePostal"
                  label="Code postal"
                  :disabled="mode === 'view'"
                  placeholder="H1A 1A1"
                />
              </div>
            </div>

            <!-- Section Contact -->
            <div class="mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
                <Phone :size="20" />
                Coordonnées
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <FormKit
                  type="text"
                  name="Telephone"
                  label="Téléphone"
                  :disabled="mode === 'view'"
                  placeholder="514-555-1234"
                />

                <FormKit
                  type="text"
                  name="telecopieur"
                  label="Télécopieur"
                  :disabled="mode === 'view'"
                  placeholder="514-555-5678"
                />

                <FormKit
                  type="text"
                  name="Poste_Garde_infirmiere"
                  label="Poste garde infirmière"
                  :disabled="mode === 'view'"
                  placeholder="Poste 123"
                  outer-class="md:col-span-2"
                />
              </div>
            </div>

            <!-- Section Exploitation -->
            <div class="mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
                <Building :size="20" />
                Exploitation
              </h3>
              <div class="grid grid-cols-1 gap-4">
                <FormKit
                  type="text"
                  name="Proprietaire"
                  label="Propriétaire / Exploitant"
                  :disabled="mode === 'view'"
                />
              </div>
            </div>

            <!-- Section Informations supplémentaires -->
            <div class="mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
                <Info :size="20" />
                Informations supplémentaires
              </h3>
              <div class="grid grid-cols-1 gap-4">
                <FormKit
                  type="text"
                  name="SourceURL"
                  label="URL source"
                  :disabled="mode === 'view'"
                  placeholder="https://..."
                />

                <FormKit
                  type="textarea"
                  name="InfosCHSLD"
                  label="Informations CHSLD"
                  :disabled="mode === 'view'"
                  rows="3"
                />

                <FormKit
                  type="textarea"
                  name="Notes"
                  label="Notes"
                  :disabled="mode === 'view'"
                  rows="4"
                  placeholder="Notes additionnelles..."
                />
              </div>
            </div>

            <!-- Section Dates système (lecture seule) -->
            <div v-if="mode !== 'create'" class="mb-6">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
                <Calendar :size="20" />
                Dates système
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <FormKit
                  type="text"
                  name="DateAjout"
                  label="Date d'ajout"
                  :disabled="true"
                />

                <FormKit
                  type="text"
                  name="DateMaj"
                  label="Dernière mise à jour"
                  :disabled="true"
                />
              </div>
            </div>

            <!-- Actions -->
            <div v-if="mode !== 'view'" class="flex items-center justify-end gap-3 pt-4 border-t border-gray-200 dark:border-gray-700">
              <button
                type="button"
                @click="handleClose"
                class="px-6 py-2 border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
              >
                Annuler
              </button>
              <button
                type="submit"
                class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
              >
                <Save :size="16" />
                {{ mode === 'create' ? 'Créer' : 'Enregistrer' }}
              </button>
            </div>
          </FormKit>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Building2, Edit, X, FileText, MapPin, Phone, Building, Info, Calendar, Save } from 'lucide-vue-next'

const props = defineProps({
  chsld: {
    type: Object,
    default: null
  },
  mode: {
    type: String,
    default: 'view',
    validator: (value) => ['view', 'edit', 'create'].includes(value)
  }
})

const emit = defineEmits(['close', 'save', 'delete'])

const isOpen = ref(true)
const mode = ref(props.mode)
const formData = ref({})

// Options pour les selects
const statutOptions = [
  { label: 'Actif', value: 'Actif' },
  { label: 'Fermé', value: 'Fermé' },
  { label: 'Suspendu', value: 'Suspendu' },
  { label: 'En rénovation', value: 'En rénovation' }
]

const typeCHSLDOptions = [
  { label: 'Public', value: 'Public' },
  { label: 'Privé conventionné', value: 'Privé conventionné' },
  { label: 'Privé non conventionné', value: 'Privé non conventionné' }
]

const regionOptions = [
  { label: '01 - Bas-Saint-Laurent', value: '01' },
  { label: '02 - Saguenay–Lac-Saint-Jean', value: '02' },
  { label: '03 - Capitale-Nationale', value: '03' },
  { label: '04 - Mauricie', value: '04' },
  { label: '05 - Estrie', value: '05' },
  { label: '06 - Montréal', value: '06' },
  { label: '07 - Outaouais', value: '07' },
  { label: '08 - Abitibi-Témiscamingue', value: '08' },
  { label: '09 - Côte-Nord', value: '09' },
  { label: '10 - Nord-du-Québec', value: '10' },
  { label: '11 - Gaspésie–Îles-de-la-Madeleine', value: '11' },
  { label: '12 - Chaudière-Appalaches', value: '12' },
  { label: '13 - Laval', value: '13' },
  { label: '14 - Lanaudière', value: '14' },
  { label: '15 - Laurentides', value: '15' },
  { label: '16 - Montérégie', value: '16' },
  { label: '17 - Centre-du-Québec', value: '17' }
]

const title = computed(() => {
  if (mode.value === 'create') return 'Nouveau CHSLD'
  if (mode.value === 'edit') return 'Modifier le CHSLD'
  return formData.value.TitreCHSLD || 'Détails du CHSLD'
})

const subtitle = computed(() => {
  if (mode.value === 'create') return 'Ajouter un nouveau centre'
  const parts = []
  if (formData.value.Municipalite) parts.push(formData.value.Municipalite)
  if (formData.value.Region) parts.push(`Région ${formData.value.Region}`)
  return parts.join(' • ')
})

// Initialiser les données du formulaire
watch(() => props.chsld, (newVal) => {
  if (newVal) {
    formData.value = { ...newVal }
  } else {
    // Valeurs par défaut pour nouveau CHSLD
    const now = new Date().toISOString()
    formData.value = {
      Region: '',
      TitreCHSLD: '',
      Adresse: '',
      Municipalite: '',
      CodePostal: '',
      Telephone: '',
      telecopieur: '',
      Poste_Garde_infirmiere: '',
      Capacite: 0,
      TypeCHSLD: 'Public',
      Proprietaire: '',
      DateCertification: '',
      Statut: 'Actif',
      SourceURL: '',
      InfosCHSLD: '',
      Notes: '',
      DateAjout: now,
      DateMaj: now
    }
  }
}, { immediate: true })

const handleClose = () => {
  isOpen.value = false
  setTimeout(() => emit('close'), 200)
}

const handleSubmit = () => {
  // Mettre à jour DateMaj
  formData.value.DateMaj = new Date().toISOString()
  emit('save', formData.value)
  handleClose()
}
</script>

<style scoped>

</style>