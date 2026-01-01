<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <h2 class="text-2xl font-bold mb-4">Nouveau client</h2>
      
      <form @submit.prevent="save" class="space-y-4">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium mb-1">Nom *</label>
            <input v-model="form.nom" required class="w-full border rounded px-3 py-2" />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Prénom *</label>
            <input v-model="form.prenom" required class="w-full border rounded px-3 py-2" />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Date de naissance</label>
            <input v-model="form.date_naissance" type="date" class="w-full border rounded px-3 py-2" />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">Téléphone</label>
            <input v-model="form.telephone" class="w-full border rounded px-3 py-2" />
          </div>
          <div class="col-span-2">
            <label class="block text-sm font-medium mb-1">Email</label>
            <input v-model="form.email" type="email" class="w-full border rounded px-3 py-2" />
          </div>
          <div class="col-span-2">
            <label class="block text-sm font-medium mb-1">Adresse</label>
            <textarea v-model="form.adresse" rows="2" class="w-full border rounded px-3 py-2"></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">N° Assurance Maladie</label>
            <input v-model="form.numero_assurance_maladie" class="w-full border rounded px-3 py-2" />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">N° Sécurité Sociale</label>
            <input v-model="form.numero_securite_sociale" class="w-full border rounded px-3 py-2" />
          </div>
        </div>

        <div class="flex justify-end gap-2 mt-6">
          <button type="button" @click="$emit('close')" class="px-4 py-2 border rounded hover:bg-gray-50">
            Annuler
          </button>
          <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
            Enregistrer
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { useRouter } from 'vue-router'; // <--- AJOUTE ÇA
import { Login } from '../../wailsjs/go/main/App';
import LeopardLogo from './LeopardLogo.vue';

export default {
  components: { LeopardLogo },
  emits: ['success'],
  setup(props, { emit }) {
    //const router = useRouter(); // <--- INITIALISE LE
    const username = ref('')
    const password = ref('')
    const error = ref('')
    
    const login = async () => {
      try {
        const user = await Login(username.value, password.value)
        
        // 1. On prévient quand même les parents si besoin
        emit('success', user)
        
        // 2. ON DÉCLENCHE LE DASHBOARD ICI !
        //router.push('/clients') 
        
      } catch (e) {
        error.value = 'Identifiants invalides'
      }
    }
    
    return { username, password, error, login }
  }
}
</script>