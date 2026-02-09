<!-- frontend/src/components/appels/AppelsEditor.vue -->

<template>
  <div class="flex-1 overflow-y-auto p-6 bg-white dark:bg-stone-900">
    <div class="max-w-3xl mx-auto">
      
      <!-- Titre -->
      <div class="mb-6">
        <h3 class="text-2xl font-semibold text-stone-800 dark:text-stone-200 mb-1">
          {{ isEditing ? 'Modifier l\'appel' : 'Nouvel Appel' }}
        </h3>
        <p class="text-sm text-stone-500 dark:text-stone-400">
          {{ isEditing ? 'Mettre à jour les informations' : 'Enregistrer les détails de l\'appel reçu' }}
        </p>
      </div>

      <!-- Formulaire FormKit -->
      <FormKit
        type="form"
        :value="formData"
        @submit="handleSubmit"
        :actions="false"
        :incomplete-message="false"
      >
        <!-- Section Appelant -->
        <div class="mb-8 p-5 bg-stone-50 dark:bg-stone-800/30 rounded-xl border border-stone-200 dark:border-stone-700">
          <h4 class="text-sm font-semibold text-stone-700 dark:text-stone-300 mb-4 flex items-center gap-2">
            <Phone :size="16" class="text-pink-500/40" />
            Informations de l'appelant
          </h4>
          
          <div class="grid grid-cols-2 gap-4">
            <FormKit
              type="text"
              name="appelant_prenom"
              label="Prénom"
              validation="required"
              placeholder="Prénom de l'appelant"
            />
            <FormKit
              type="text"
              name="appelant_nom"
              label="Nom"
              validation="required"
              placeholder="Nom de l'appelant"
            />
          </div>

          <div class="grid grid-cols-2 gap-4 mt-4">
            <FormKit
              type="tel"
              name="appelant_telephone"
              label="Téléphone"
              placeholder="(514) 555-0123"
            />
            <FormKit
              type="select"
              name="appelant_relation"
              label="Relation"
              :options="[
                { label: 'Lui-même / Elle-même', value: 'lui_meme' },
                { label: 'Conjoint(e)', value: 'conjoint' },
                { label: 'Fils / Fille', value: 'enfant' },
                { label: 'Aidant naturel', value: 'aidant' },
                { label: 'Proche', value: 'proche' },
                { label: 'Professionnel', value: 'professionnel' },
                { label: 'Autre', value: 'autre' }
              ]"
            />
          </div>
        </div>

        <!-- Section Client / Prospect -->
        <div class="mb-8 p-5 bg-pink-50/30 dark:bg-pink-900/10 rounded-xl border border-pink-200 dark:border-pink-800/30">
          <h4 class="text-sm font-semibold text-stone-700 dark:text-stone-300 mb-4 flex items-center gap-2">
            <User :size="16" class="text-pink-500/50" />
            Personne concernée
          </h4>

          <!-- Toggle Client existant ou Prospect -->
          <div class="mb-4">
            <FormKit
              type="checkbox"
              name="is_existing_client"
              label="Client existant dans le système"
              :value="!!formData.client_id"
              @input="toggleClientType"
            />
          </div>

          <!-- Si client existant -->
          <div v-if="formData.client_id !== null" class="mb-4">
            <FormKit
              type="select"
              name="client_id"
              label="Sélectionner le client"
              :options="clientOptions"
              placeholder="Choisir un client..."
            />
          </div>

          <!-- Si prospect (nouveau) -->
          <div v-else class="grid grid-cols-2 gap-4">
            <FormKit
              type="text"
              name="prospect_prenom"
              label="Prénom"
              placeholder="Prénom du prospect"
            />
            <FormKit
              type="text"
              name="prospect_nom"
              label="Nom"
              placeholder="Nom du prospect"
            />
            <FormKit
              type="tel"
              name="prospect_telephone"
              label="Téléphone"
              placeholder="(514) 555-0123"
              outer-class="col-span-2"
            />
          </div>
        </div>

        <!-- Section Détails de l'appel -->
        <div class="mb-8 p-5 bg-stone-50 dark:bg-stone-800/30 rounded-xl border border-stone-200 dark:border-stone-700">
          <h4 class="text-sm font-semibold text-stone-700 dark:text-stone-300 mb-4 flex items-center gap-2">
            <FileText :size="16" class="text-pink-500/45" />
            Détails de la demande
          </h4>

          <div class="grid grid-cols-2 gap-4 mb-4">
            <FormKit
              type="select"
              name="type_demande"
              label="Type de demande"
              validation="required"
              :options="[
                { label: 'Évaluation psychosociale', value: 'evaluation' },
                { label: 'Mandat de protection', value: 'mandat_protection' },
                { label: 'Consultation', value: 'consultation' },
                { label: 'Information générale', value: 'information' },
                { label: 'Autre', value: 'autre' }
              ]"
            />
            <FormKit
              type="select"
              name="priorite"
              label="Priorité"
              :options="[
                { label: 'Urgente', value: 'urgente' },
                { label: 'Haute', value: 'haute' },
                { label: 'Normale', value: 'normale' },
                { label: 'Basse', value: 'basse' }
              ]"
            />
          </div>

          <FormKit
            type="textarea"
            name="motif_appel"
            label="Motif de l'appel"
            rows="3"
            placeholder="Décrivez brièvement le motif de l'appel..."
          />

          <FormKit
            type="textarea"
            name="notes_internes"
            label="Notes internes"
            rows="3"
            placeholder="Notes confidentielles pour le suivi..."
            help="Ces notes ne sont visibles que par l'équipe"
          />
        </div>

        <!-- Section Suivi -->
        <div class="mb-8 p-5 bg-stone-50 dark:bg-stone-800/30 rounded-xl border border-stone-200 dark:border-stone-700">
          <h4 class="text-sm font-semibold text-stone-700 dark:text-stone-300 mb-4 flex items-center gap-2">
            <Calendar :size="16" class="text-pink-500/40" />
            Suivi et rendez-vous
          </h4>

          <div class="grid grid-cols-2 gap-4 mb-4">
            <FormKit
              type="select"
              name="statut"
              label="Statut"
              :options="[
                { label: 'Nouveau', value: 'nouveau' },
                { label: 'En attente', value: 'en_attente' },
                { label: 'RDV planifié', value: 'rdv_planifie' },
                { label: 'Complété', value: 'complete' },
                { label: 'Annulé', value: 'annule' }
              ]"
            />
            <FormKit
              type="select"
              name="assigne_a"
              label="Assigné à"
              :options="userOptions"
              placeholder="Non assigné"
            />
          </div>

          <!-- RDV -->
          <div class="grid grid-cols-3 gap-4">
            <FormKit
              type="date"
              name="rdv_date"
              label="Date RDV"
            />
            <FormKit
              type="time"
              name="rdv_heure"
              label="Heure RDV"
            />
            <FormKit
              type="text"
              name="rdv_lieu"
              label="Lieu RDV"
              placeholder="Bureau, Zoom, etc."
            />
          </div>
        </div>

      </FormKit>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Phone, User, FileText, Calendar } from 'lucide-vue-next'

const props = defineProps({
  appel: {
    type: Object,
    default: null
  },
  clients: {
    type: Array,
    default: () => []
  },
  users: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['submit'])

const isEditing = computed(() => !!props.appel?.id)

const formData = ref({
  appelant_nom: '',
  appelant_prenom: '',
  appelant_telephone: '',
  appelant_relation: 'lui_meme',
  client_id: null,
  prospect_nom: '',
  prospect_prenom: '',
  prospect_telephone: '',
  type_demande: 'evaluation',
  motif_appel: '',
  priorite: 'normale',
  statut: 'nouveau',
  notes_internes: '',
  rdv_date: '',
  rdv_heure: '',
  rdv_lieu: '',
  assigne_a: null,
  is_existing_client: false
})

// Charger les données si édition
watch(() => props.appel, (newAppel) => {
  if (newAppel) {
    formData.value = {
      ...formData.value,
      ...newAppel,
      is_existing_client: !!newAppel.client_id
    }
  }
}, { immediate: true })

const clientOptions = computed(() => {
  return [
    { label: 'Sélectionner...', value: null },
    ...props.clients.map(c => ({
      label: `${c.prenom} ${c.nom}`,
      value: c.id
    }))
  ]
})

const userOptions = computed(() => {
  return [
    { label: 'Non assigné', value: null },
    ...props.users.map(u => ({
      label: u.full_name,
      value: u.id
    }))
  ]
})

const toggleClientType = (value) => {
  if (value) {
    formData.value.prospect_nom = ''
    formData.value.prospect_prenom = ''
    formData.value.prospect_telephone = ''
  } else {
    formData.value.client_id = null
  }
}

const handleSubmit = (data) => {
  emit('submit', data)
}
</script>

<style>
/* Style FormKit pour le thème zen */
.formkit-outer {
  margin-bottom: 1rem;
}

.formkit-label {
  font-size: 0.875rem;
  font-weight: 500;
  color: rgb(68 64 60);
  margin-bottom: 0.5rem;
}

.dark .formkit-label {
  color: rgb(231 229 228);
}

.formkit-input {
  width: 100%;
  padding: 0.625rem 0.875rem;
  border: 1px solid rgb(231 229 228);
  border-radius: 0.75rem;
  background: white;
  font-size: 0.875rem;
  transition: all 0.2s;
}

.dark .formkit-input {
  background: rgb(28 25 23);
  border-color: rgb(68 64 60);
  color: rgb(231 229 228);
}

.formkit-input:focus {
  outline: none;
  border-color: rgb(236 72 153);
  ring: 2px;
  ring-color: rgb(236 72 153 / 0.2);
}

.formkit-help {
  font-size: 0.75rem;
  color: rgb(120 113 108);
  margin-top: 0.25rem;
}

.dark .formkit-help {
  color: rgb(168 162 158);
}
</style>