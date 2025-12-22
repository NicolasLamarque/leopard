<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold">Clients</h1>
      <button @click="showForm = true" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
        + Nouveau client
      </button>
    </div>


    <!-- Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nom</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Prénom</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Téléphone</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Email</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="client in clients" :key="client.id" class="hover:bg-gray-50">
            <td class="px-6 py-4">{{ client.nom }}</td>
            <td class="px-6 py-4">{{ client.prenom }}</td>
            <td class="px-6 py-4">{{ client.telephone }}</td>
            <td class="px-6 py-4">{{ client.email }}</td>
            <td class="px-6 py-4">
              <!-- Bouton pour afficher les détails -->
              <button @click="viewClient(client.id)" class="text-blue-600 hover:text-blue-800 mr-3">Voir</button>

<Dialog v-model:visible="showDetails" modal header="Détails du client" :style="{ width: '50vw' }">
  <ClientDetails @close="showDetails = false" @updated="loadClients" />
</Dialog>
              <button @click="deleteClient(client.id)" class="text-red-600 hover:text-red-800">
                Supprimer
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { GetClients, DeleteClient } from '../../wailsjs/go/main/App'
import ClientForm from '../components/ClientForm.vue'

// Import du store Pinia
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
  
  // Correction ici : on ajoute /app/
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