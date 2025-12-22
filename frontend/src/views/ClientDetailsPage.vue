
<template>   
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">
        Dossier de {{ store.currentClient?.prenom }} {{ store.currentClient?.nom }}
      </h1>
      <button 
        @click="router.push('/app/clients')" 
        class="bg-gray-200 hover:bg-gray-300 text-gray-700 px-4 py-2 rounded shadow"
      >
        Fermer le dossier
      </button>
    </div>

    <ClientDetailsForm 
      v-if="store.currentClient" 
      :clientData="store.currentClient" 
      @save="onSave"
    />
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useClientStore } from '../stores/clientStore'
import { UpdateClient } from '../../wailsjs/go/main/App'
import ClientDetailsForm from '../components/ClientDetailsForms.vue'

const route = useRoute()
const store = useClientStore()
const router = useRouter()
const onSave = async (data) => {
  try {
    await UpdateClient(data)
    alert("Modifications enregistrées !")
  } catch (err) {
    console.error(err)
  }
}

onMounted(() => {
  const id = parseInt(route.params.id)
  store.selectClient(id) // Le store Pinia appelle Go
})
</script>