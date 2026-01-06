<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white dark:bg-gray-900 rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <h2 class="text-2xl font-bold mb-4 dark:text-white">Nouveau client</h2>
      
      <form @submit.prevent="save" class="space-y-4">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Nom *</label>
            <input 
              v-model="form.nom" 
              required 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Prénom *</label>
            <input 
              v-model="form.prenom" 
              required 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Date de naissance</label>
            <input 
              v-model="form.date_naissance" 
              type="date" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Téléphone</label>
            <input 
              v-model="form.telephone" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Cellulaire</label>
            <input 
              v-model="form.cellulaire" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Email</label>
            <input 
              v-model="form.email" 
              type="email" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div class="col-span-2">
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Adresse</label>
            <textarea 
              v-model="form.adresse" 
              rows="2" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white"
            ></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Code postal</label>
            <input 
              v-model="form.code_postal" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">Ville</label>
            <input 
              v-model="form.ville" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">N° Assurance Maladie</label>
            <input 
              v-model="form.numero_assurance_maladie" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-1 dark:text-gray-200">N° Sécurité Sociale</label>
            <input 
              v-model="form.numero_securite_sociale" 
              class="w-full border rounded px-3 py-2 dark:bg-gray-800 dark:border-gray-700 dark:text-white" 
            />
          </div>
        </div>

        <div class="flex justify-end gap-2 mt-6">
          <button 
            type="button" 
            @click="$emit('close')" 
            class="px-4 py-2 border rounded hover:bg-gray-50 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-200 dark:hover:bg-gray-700"
          >
            Annuler
          </button>
          <button 
            type="submit" 
            :disabled="loading"
            class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50"
          >
            {{ loading ? 'Enregistrement...' : 'Enregistrer' }}
          </button>
        </div>

        <p v-if="error" class="text-red-500 text-sm mt-2">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { CreateClient } from '../../wailsjs/go/main/App';

export default {
  emits: ['close', 'created'],
  setup(props, { emit }) {
    const loading = ref(false);
    const error = ref('');
    
    const form = ref({
      nom: '',
      prenom: '',
      date_naissance: null,
      telephone: null,
      cellulaire: null,
      email: null,
      adresse: null,
      code_postal: null,
      ville: null,
      pays: 'Canada',
      numero_assurance_maladie: null,
      numero_securite_sociale: null,
      no_hcm: null,
      no_chaur: null,
      no_dossier_leopard: null,
      medecin_famille_No_Licence: null,
      notaire_id: null,
      pivot_id: null,
      rpa_id: null,
      chsld_id: null,
      ri_id: null,
      note_fixe: null,
      Actif: 1,
      dcd: 0
    });
    
    const save = async () => {
      loading.value = true;
      error.value = '';
      
      try {
        console.log('📝 Création client avec données:', form.value);
        
        // Appel à la fonction Go qui va automatiquement chiffrer les données
        const clientId = await CreateClient(form.value);
        
        console.log('✅ Client créé avec ID:', clientId);
        
        // Fermer le formulaire et rafraîchir la liste
        emit('created');
        emit('close');
      } catch (e) {
        console.error('❌ Erreur création client:', e);
        error.value = e.message || 'Erreur lors de la création du client';
      } finally {
        loading.value = false;
      }
    };
    
    return { 
      form, 
      loading,
      error,
      save 
    };
  }
}
</script>

<style scoped>
/* Styles additionnels si nécessaire */
</style>