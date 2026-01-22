<template>
  <div class="h-full flex flex-col">
    
    <!-- Header -->
    <div class="bg-gradient-to-r from-slate-600 to-teal-700 dark:from-slate-500 dark:to-teal-600 px-6 py-5 flex items-center justify-between">
      <div>
        <h2 class="text-xl font-bold text-white">Contacts & Relations</h2>
        <p class="text-blue-100 text-sm mt-0.5">{{ contacts.length }} contact(s)</p>
      </div>
      <button 
        @click="$emit('close')"
        class="p-2 hover:bg-white/20 rounded-lg transition-colors"
      >
        <X :size="24" class="text-white" />
      </button>
    </div>

    <!-- Actions rapides -->
    <div class="px-6 py-4 bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-800">
      <button
        @click="showForm = true"
        class="w-full flex items-center justify-center gap-2 bg-slate-600 hover:bg-slate-700 text-white px-4 py-3 rounded-lg font-medium transition-all shadow-sm"
      >
        <UserPlus :size="18" />
        <span>Ajouter un contact</span>
      </button>
    </div>

    <!-- Filtres -->
    <div class="px-6 py-3 bg-gray-50 dark:bg-gray-900/50 border-b border-gray-200 dark:border-gray-800">
      <div class="flex gap-2">
        <button
          v-for="filter in filters"
          :key="filter.value"
          @click="activeFilter = filter.value"
          :class="[
            'px-3 py-1.5 rounded-lg text-sm font-medium transition-all',
            activeFilter === filter.value
              ? 'bg-blue-600 text-white'
              : 'bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'
          ]"
        >
          {{ filter.label }}
        </button>
      </div>
    </div>

    <!-- Liste des contacts -->
    <div class="flex-1 overflow-y-auto px-6 py-4">
      <div v-if="loading" class="flex items-center justify-center py-12">
        <Loader2 :size="32" class="animate-spin text-blue-600 dark:text-blue-400" />
      </div>

      <div v-else-if="filteredContacts.length === 0" class="text-center py-12">
        <UserX :size="48" class="mx-auto text-gray-400 mb-3" />
        <p class="text-gray-600 dark:text-gray-400">Aucun contact trouvé</p>
        <p class="text-sm text-gray-500 dark:text-gray-500 mt-1">
          Ajoutez un premier contact pour ce client
        </p>
      </div>

      <div v-else class="space-y-3">
        <div 
          v-for="contact in filteredContacts"
          :key="contact.id"
          class="bg-white dark:bg-gray-800 rounded-xl p-4 border border-gray-200 dark:border-gray-700 hover:shadow-md transition-all cursor-pointer group"
          @click="viewContact(contact)"
        >
          <!-- Header du contact -->
          <div class="flex items-start justify-between mb-3">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-full bg-gradient-to-br from-slate-500 to-teal-700 flex items-center justify-center text-white font-semibold">
                {{ getInitials(contact) }}
              </div>
              <div>
                <h3 class="font-semibold text-gray-900 dark:text-white group-hover:text-slate-600 dark:group-hover:text-slate-400 transition-colors">
                  {{ contact.prenom }} {{ contact.nom }}
                </h3>
                <p class="text-xs text-gray-500 dark:text-gray-400">
                  {{ contact.relation_type || 'Relation non spécifiée' }}
                </p>
              </div>
            </div>

            <!-- Actions rapides -->
            <div class="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
              <button
                @click.stop="editContact(contact)"
                class="p-1.5 hover:bg-blue-50 dark:hover:bg-blue-900/30 rounded-lg transition-colors"
              >
                <Edit :size="16" class="text-teal-600 dark:text-teal-600" />
              </button>
              <button
                @click.stop="deleteContact(contact.id)"
                class="p-1.5 hover:bg-red-50 dark:hover:bg-red-900/30 rounded-lg transition-colors"
              >
                <Trash2 :size="16" class="text-amber-600 dark:text-amber-800" />
              </button>
            </div>
          </div>

          <!-- Badges -->
          <div class="flex flex-wrap gap-2 mb-3">
            <span v-if="contact.contact_urgence === 1" class="px-2 py-0.5 bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300 text-xs font-medium rounded-full flex items-center gap-1">
              <AlertCircle :size="12" />
              Contact urgence
            </span>
            <span v-if="contact.aidant_naturel === 1" class="px-2 py-0.5 bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300 text-xs font-medium rounded-full flex items-center gap-1">
              <Heart :size="12" />
              Aidant naturel
            </span>
            <span v-if="contact.procuration_bancaire === 1" class="px-2 py-0.5 bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300 text-xs font-medium rounded-full flex items-center gap-1">
              <Briefcase :size="12" />
              Procuration bancaire
            </span>
            <span v-if="contact.procuration_notariee === 1" class="px-2 py-0.5 bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300 text-xs font-medium rounded-full flex items-center gap-1">
              <FileText :size="12" />
              Procuration notariée
            </span>
            <span v-if="contact.lien_familial" class="px-2 py-0.5 bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300 text-xs font-medium rounded-full">
              {{ contact.lien_familial }}
            </span>
          </div>

          <!-- Infos de contact -->
          <div class="space-y-1.5 text-sm">
            <p v-if="contact.telephone" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <Phone :size="14" class="text-gray-400" />
              {{ contact.telephone }}
            </p>
            <p v-if="contact.cellulaire" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <Smartphone :size="14" class="text-gray-400" />
              {{ contact.cellulaire }}
            </p>
            <p v-if="contact.email" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <Mail :size="14" class="text-gray-400" />
              {{ contact.email }}
            </p>
            <p v-if="contact.ville" class="text-gray-600 dark:text-gray-400 flex items-center gap-2">
              <MapPin :size="14" class="text-gray-400" />
              {{ contact.ville }}
            </p>
          </div>

          <!-- Force du lien (si applicable) -->
          <div v-if="contact.force_lien > 0" class="mt-3 pt-3 border-t border-gray-200 dark:border-gray-700">
            <div class="flex items-center gap-2">
              <span class="text-xs text-gray-500 dark:text-gray-400">Force du lien:</span>
              <div class="flex gap-1">
                <div 
                  v-for="i in 5" 
                  :key="i"
                  :class="[
                    'w-5 h-2 rounded-full transition-all',
                    i <= Math.ceil(contact.force_lien / 2)
                      ? 'bg-blue-600 dark:bg-blue-400'
                      : 'bg-gray-200 dark:bg-gray-700'
                  ]"
                ></div>
              </div>
              <span class="text-xs font-medium text-gray-700 dark:text-gray-300">
                {{ contact.force_lien }}/10
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer avec stats -->
    <div class="px-6 py-4 bg-gray-50 dark:bg-gray-900/50 border-t border-gray-200 dark:border-gray-800">
      <div class="grid grid-cols-2 gap-3 text-sm">
        <div class="text-center p-2 bg-white dark:bg-gray-800 rounded-lg">
          <p class="text-xs text-gray-500 dark:text-gray-400">Contacts urgence</p>
          <p class="text-lg font-bold text-gray-900 dark:text-white">{{ stats.contactsUrgence }}</p>
        </div>
        <div class="text-center p-2 bg-white dark:bg-gray-800 rounded-lg">
          <p class="text-xs text-gray-500 dark:text-gray-400">Aidants naturels</p>
          <p class="text-lg font-bold text-gray-900 dark:text-white">{{ stats.aidantsNaturels }}</p>
        </div>
      </div>
    </div>

    <!-- Modal Formulaire Contact -->
    <ContactForm
      v-if="showForm"
      :client-id="clientId"
      :contact-data="selectedContact"
      @close="showForm = false; selectedContact = null"
      @saved="onContactSaved"
    />

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  X, UserPlus, Phone, Mail, MapPin, Edit, Trash2, 
  Loader2, UserX, AlertCircle, Heart, Briefcase, 
  FileText, Smartphone
} from 'lucide-vue-next'
import ContactForm from './ContactForm.vue'
import { GetAllContactsByClientID, DeleteContact } from '../../../wailsjs/go/main/App'

const props = defineProps({
  clientId: { type: Number, required: true }
})

const emit = defineEmits(['close', 'contact-added', 'contact-deleted'])

const contacts = ref([])
const loading = ref(false)
const showForm = ref(false)
const selectedContact = ref(null)
const activeFilter = ref('all')

const filters = [
  { value: 'all', label: 'Tous' },
  { value: 'urgence', label: 'Urgence' },
  { value: 'aidant', label: 'Aidants' },
  { value: 'famille', label: 'Famille' }
]

const filteredContacts = computed(() => {
  if (activeFilter.value === 'all') return contacts.value
  if (activeFilter.value === 'urgence') return contacts.value.filter(c => c.contact_urgence === 1)
  if (activeFilter.value === 'aidant') return contacts.value.filter(c => c.aidant_naturel === 1)
  if (activeFilter.value === 'famille') return contacts.value.filter(c => c.relation_type === 'Famille')
  return contacts.value
})

const stats = computed(() => ({
  contactsUrgence: contacts.value.filter(c => c.contact_urgence === 1).length,
  aidantsNaturels: contacts.value.filter(c => c.aidant_naturel === 1).length
}))

const getInitials = (contact) => {
  return `${contact.prenom[0]}${contact.nom[0]}`.toUpperCase()
}

const loadContacts = async () => {
  if (!props.clientId) return;
  
  loading.value = true;
  try {
    const result = await GetAllContactsByClientID(props.clientId);
    contacts.value = result || [];
  } catch (error) {
    // Si Go envoie une erreur, on la voit ici dans la console F12
    console.error('Erreur reçue de Go:', error);
    contacts.value = []; 
  } finally {
    // QUOI QU'IL ARRIVE, on arrête le cercle qui tourne
    loading.value = false;
  }
}

const viewContact = (contact) => {
  selectedContact.value = contact
  showForm.value = true
}

const editContact = (contact) => {
  selectedContact.value = contact
  showForm.value = true
}

const deleteContact = async (id) => {
  if (confirm('Êtes-vous sûr de vouloir supprimer ce contact ?')) {
    try {
      await DeleteContact(id)
      await loadContacts()
      emit('contact-deleted')
    } catch (error) {
      console.error('Erreur suppression contact:', error)
      alert('Erreur lors de la suppression')
    }
  }
}

const onContactSaved = () => {
  loadContacts()
  emit('contact-added')
  showForm.value = false
  selectedContact.value = null
}

onMounted(() => {
  loadContacts()
})
</script>