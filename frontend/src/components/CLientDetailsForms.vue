<template>
  <div class="bg-white p-6 rounded-lg shadow">
    <FormKit
      type="form"
      v-model="formData"
      submit-label="Sauvegarder les modifications"
      @submit="submitHandler"
    >
      <div class="grid grid-cols-2 gap-4">
        <FormKit type="text" name="prenom" label="Prénom" validation="required" />
        <FormKit type="text" name="nom" label="Nom" validation="required" />
        <FormKit  type="tel"  name="telephone"  label="Téléphone"  placeholder="514-555-0123"  validation="required|matches:/^[0-9]{10}$/"  :validation-messages="{    matches: 'Le numéro doit contenir exactement 10 chiffres.',  }"
/>
        <FormKit type="email" name="email" label="Courriel" />
      </div>
      
      <h3 class="text-lg font-semibold mt-6 mb-2 border-b pb-2">Dossier Clinique</h3>
      <div class="grid grid-cols-3 gap-4">
        <FormKit type="text" name="no_hcm" label="No HCM" />
        <FormKit type="text" name="no_chaur" label="No CHAUR" />
        
        <!-- Numéro Leopard avec actions -->
        <div class="relative">
          <FormKit 
            type="text" 
            name="no_dossier_leopard" 
            label="No Dossier Léopard" 
            :value="leopardNumber"
            readonly 
            :help="leopardNumberHelp"
          />
          
          <div class="flex gap-2 mt-2">
            <!-- Bouton ouvrir dossier -->
            <button
              v-if="folderExists"
              @click="handleOpenFolder"
              type="button"
              class="flex items-center gap-2 px-3 py-1.5 text-sm bg-blue-50 text-blue-700 hover:bg-blue-100 rounded-lg transition-colors"
              title="Ouvrir le dossier client"
            >
              <FolderOpen :size="16" />
              <span>Ouvrir</span>
            </button>
            
            <!-- Bouton créer dossier -->
            <button
              v-else
              @click="handleCreateFolder"
              type="button"
              :disabled="isCreatingFolder"
              class="flex items-center gap-2 px-3 py-1.5 text-sm bg-green-50 text-green-700 hover:bg-green-100 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              title="Créer le dossier client sur le disque"
            >
              <FolderPlus :size="16" v-if="!isCreatingFolder" />
              <Loader2 :size="16" class="animate-spin" v-else />
              <span>{{ isCreatingFolder ? 'Création...' : 'Créer dossier' }}</span>
            </button>
            
            <!-- Bouton régénérer -->
            <button
              @click="handleRegenerateLeopardNumber"
              type="button"
              class="flex items-center gap-2 px-3 py-1.5 text-sm bg-gray-50 text-gray-700 hover:bg-gray-100 rounded-lg transition-colors"
              title="Régénérer le numéro Leopard"
            >
              <RefreshCw :size="16" />
            </button>
          </div>
        </div>
      </div>

      <!-- Informations sur le dossier -->
      <div v-if="folderInfo" class="mt-4 p-4 bg-blue-50 rounded-lg border border-blue-200">
        <div class="flex items-start gap-3">
          <FolderCheck :size="20" class="text-blue-600 mt-0.5" />
          <div class="flex-1">
            <p class="text-sm font-semibold text-blue-900">Dossier client actif</p>
            <p class="text-xs text-blue-700 mt-1">{{ folderInfo.path }}</p>
            <div class="flex gap-2 mt-2 text-xs text-blue-600">
              <span>📁 {{ folderInfo.subfoldersCount }} sous-dossiers</span>
              <span>•</span>
              <span>📄 {{ folderInfo.filesCount }} fichiers</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Alerte si dossier n'existe pas -->
      <div v-else-if="formData.id && !folderExists" class="mt-4 p-4 bg-yellow-50 rounded-lg border border-yellow-200">
        <div class="flex items-start gap-3">
          <AlertCircle :size="20" class="text-yellow-600 mt-0.5" />
          <div class="flex-1">
            <p class="text-sm font-semibold text-yellow-900">Dossier client non créé</p>
            <p class="text-xs text-yellow-700 mt-1">
              Le dossier physique pour ce client n'existe pas encore. Cliquez sur "Créer dossier" pour le créer.
            </p>
          </div>
        </div>
      </div>
    </FormKit>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue';
import { FolderOpen, FolderPlus, FolderCheck, RefreshCw, Loader2, AlertCircle } from 'lucide-vue-next';
import { 
  generateLeopardNumber, 
  createClientFolder, 
  openClientFolder, 
  clientFolderExists 
} from '@/utils/clientFolderManager';

const props = defineProps({
  clientData: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['save', 'folderCreated']);

const formData = ref({ ...props.clientData });
const leopardNumber = ref('');
const folderExists = ref(false);
const folderInfo = ref(null);
const isCreatingFolder = ref(false);

/**
 * Génère le numéro Leopard au chargement
 */
const generateNumber = () => {
  if (formData.value.nom && formData.value.prenom) {
    leopardNumber.value = generateLeopardNumber(
      formData.value.nom,
      formData.value.prenom,
      formData.value.created_at
    );
    formData.value.no_dossier_leopard = leopardNumber.value;
  }
};

/**
 * Texte d'aide pour le numéro Leopard
 */
const leopardNumberHelp = computed(() => {
  if (!formData.value.nom || !formData.value.prenom) {
    return 'Remplissez le nom et prénom pour générer automatiquement';
  }
  return `Format: ${formData.value.nom.substring(0, 3).toUpperCase()}${formData.value.prenom.substring(0, 1).toUpperCase()}_DATE`;
});

/**
 * Vérifie si le dossier existe
 */
const checkFolderExists = async () => {
  if (leopardNumber.value) {
    folderExists.value = await clientFolderExists(leopardNumber.value);
    
    if (folderExists.value) {
      // Récupérer les infos du dossier
      await loadFolderInfo();
    }
  }
};

/**
 * Charge les informations du dossier
 */
const loadFolderInfo = async () => {
  try {
    const info = await window.go.main.App.GetClientFolderInfo(leopardNumber.value);
    if (info) {
      folderInfo.value = info;
    }
  } catch (error) {
    console.error('❌ Erreur chargement infos dossier:', error);
  }
};

/**
 * Crée le dossier client
 */
const handleCreateFolder = async () => {
  try {
    isCreatingFolder.value = true;
    
    const result = await createClientFolder(formData.value);
    
    if (result.success) {
      folderExists.value = true;
      leopardNumber.value = result.leopardNumber;
      formData.value.no_dossier_leopard = result.leopardNumber;
      
      await loadFolderInfo();
      
      emit('folderCreated', {
        leopardNumber: result.leopardNumber,
        path: result.path
      });
      
      // Notification succès
      console.log('✅ Dossier créé avec succès');
      alert(`Dossier client créé avec succès!\n\nNuméro: ${result.leopardNumber}\nChemin: ${result.path}`);
    } else {
      throw new Error(result.error);
    }
  } catch (error) {
    console.error('❌ Erreur création dossier:', error);
    alert(`Erreur lors de la création du dossier:\n${error.message}`);
  } finally {
    isCreatingFolder.value = false;
  }
};

/**
 * Ouvre le dossier client
 */
const handleOpenFolder = async () => {
  try {
    const clientName = `${formData.value.nom} ${formData.value.prenom}`;
    const success = await openClientFolder(leopardNumber.value, clientName);
    
    if (!success) {
      alert('Impossible d\'ouvrir le dossier client');
    }
  } catch (error) {
    console.error('❌ Erreur ouverture dossier:', error);
    alert(`Erreur lors de l'ouverture du dossier:\n${error.message}`);
  }
};

/**
 * Régénère le numéro Leopard
 */
const handleRegenerateLeopardNumber = () => {
  if (confirm('Êtes-vous sûr de vouloir régénérer le numéro Leopard?\nCela peut créer des incohérences si un dossier existe déjà.')) {
    generateNumber();
    checkFolderExists();
  }
};

/**
 * Soumission du formulaire
 */
const submitHandler = async (data) => {
  // S'assurer que le numéro Leopard est inclus
  data.no_dossier_leopard = leopardNumber.value;
  
  emit('save', data);
};

// Watchers
watch(() => props.clientData, (newVal) => {
  formData.value = { ...newVal };
  generateNumber();
  checkFolderExists();
}, { deep: true });

watch([
  () => formData.value.nom,
  () => formData.value.prenom
], () => {
  generateNumber();
});

// Au montage
onMounted(() => {
  generateNumber();
  checkFolderExists();
});
</script>

<style scoped>
/* Animations pour le bouton de création */
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>