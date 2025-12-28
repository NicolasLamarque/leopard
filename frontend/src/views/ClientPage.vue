<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold dark:text-white">Clients</h1>
      <button 
        @click="showForm = true" 
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600"
      >
        + Nouveau client
      </button>
    </div>

    <!-- Table avec dark mode complet -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow overflow-hidden">
      <table class="min-w-full">
        <thead class="bg-gray-50 dark:bg-gray-700">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Nom
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Prénom
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Téléphone
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Email
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
          <tr 
            v-for="client in clients" 
            :key="client.id" 
            class="hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
          >
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ client.nom }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ client.prenom }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ client.telephone }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">
              {{ client.email }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <button 
                @click="viewClient(client.id)" 
                class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 mr-3 font-medium"
              >
                Voir
              </button>
              <button 
                @click="deleteClient(client.id)" 
                class="text-red-600 dark:text-red-400 hover:text-red-800 dark:hover:text-red-300 font-medium"
              >
                Supprimer
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Formulaire client -->
    <ClientForm 
      v-if="showForm" 
      @close="showForm = false" 
      @created="loadClients" 
    />
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { GetClients, DeleteClient } from '../../wailsjs/go/main/App'
import ClientForm from '../components/ClientForm.vue'
import { useClientStore } from '../stores/clientStore'

export default {
  components: { ClientForm },
  setup() {
    const clients = ref([])
    const showForm = ref(false)
    const router = useRouter()
    const store = useClientStore()

    const loadClients = async () => {
      try {
        console.log('🔍 Chargement de la liste des clients...')
        const result = await GetClients()
        clients.value = result || []
      } catch (err) {
        console.error('❌ Erreur chargement clients:', err)
      }
    }

    const viewClient = async (id) => {
      store.clearClient()
      console.log(`📂 Navigation vers le dossier client ID: ${id}`)
      router.push(`/app/clients/${id}`)
    }

    const deleteClient = async (id) => {
      if (confirm("Êtes-vous sûr de vouloir supprimer ce dossier ? Cette action est irréversible.")) {
        try {
          await DeleteClient(id)
          await loadClients()
        } catch (err) {
          console.error('❌ Erreur lors de la suppression:', err)
        }
      }
    }

    onMounted(() => {
      loadClients()
    })

    return { 
      clients, 
      showForm, 
      loadClients, 
      viewClient, 
      deleteClient,
      store 
    }
  }
}
</script>

<style scoped>
/* Styles additionnels si nécessaire */
</style>