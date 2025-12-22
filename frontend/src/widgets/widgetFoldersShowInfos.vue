<template>
  <div class="fixed bottom-4 right-4 bg-white rounded-lg shadow-lg p-4 max-w-md border-2 border-blue-500">
    <div class="flex items-start justify-between gap-3">
      <div class="flex-1">
        <h3 class="text-sm font-bold text-gray-900 mb-2 flex items-center gap-2">
          <Folder :size="18" class="text-blue-600" />
          Dossier Leopard
        </h3>
        
        <div v-if="loading" class="text-sm text-gray-500">
          Chargement...
        </div>
        
        <div v-else-if="error" class="text-sm text-red-600">
          ❌ {{ error }}
        </div>
        
        <div v-else class="space-y-2">
          <!-- Chemin -->
          <div class="bg-gray-50 p-2 rounded border border-gray-200">
            <p class="text-xs font-mono text-gray-700 break-all">
              {{ leopardPath }}
            </p>
          </div>
          
          <!-- Stats -->
          <div class="flex gap-2 text-xs">
            <div class="flex items-center gap-1 text-blue-600">
              <FolderOpen :size="14" />
              <span>{{ clientFolders.length }} dossiers</span>
            </div>
          </div>
          
          <!-- Boutons -->
          <div class="flex gap-2">
            <button
              @click="openFolder"
              class="flex-1 flex items-center justify-center gap-1 px-3 py-1.5 bg-blue-600 text-white text-xs rounded hover:bg-blue-700 transition-colors"
            >
              <ExternalLink :size="14" />
              <span>Ouvrir</span>
            </button>
            
            <button
              @click="refresh"
              class="flex items-center justify-center px-3 py-1.5 bg-gray-100 text-gray-700 text-xs rounded hover:bg-gray-200 transition-colors"
            >
              <RefreshCw :size="14" :class="{ 'animate-spin': loading }" />
            </button>
            
            <button
              @click="copyPath"
              class="flex items-center justify-center px-3 py-1.5 bg-gray-100 text-gray-700 text-xs rounded hover:bg-gray-200 transition-colors"
              title="Copier le chemin"
            >
              <Copy :size="14" />
            </button>
          </div>
          
          <!-- Liste des dossiers (si plusieurs) -->
          <div v-if="clientFolders.length > 0" class="mt-2 pt-2 border-t border-gray-200">
            <p class="text-xs font-semibold text-gray-700 mb-1">Derniers dossiers créés:</p>
            <div class="space-y-1 max-h-32 overflow-y-auto">
              <div
                v-for="folder in clientFolders.slice(0, 5)"
                :key="folder"
                class="text-xs text-gray-600 truncate"
              >
                📁 {{ folder }}
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <button
        @click="$emit('close')"
        class="p-1 text-gray-400 hover:text-gray-600 transition-colors"
      >
        <X :size="16" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { Folder, FolderOpen, ExternalLink, RefreshCw, Copy, X } from 'lucide-vue-next';

defineEmits(['close']);

const leopardPath = ref('');
const clientFolders = ref([]);
const loading = ref(true);
const error = ref('');

/**
 * Charge les informations du dossier
 */
const loadInfo = async () => {
  try {
    loading.value = true;
    error.value = '';
    
    if (!window.go?.main?.App) {
      throw new Error('API Wails non disponible');
    }
    
    // Récupérer le chemin
    const path = await window.go.main.App.GetBasePath();
    leopardPath.value = path + '/Clients';
    
    // Récupérer la liste des dossiers
    const folders = await window.go.main.App.ListClientFolders();
    clientFolders.value = folders || [];
    
  } catch (err) {
    error.value = err.message;
    console.error('❌ Erreur chargement infos:', err);
  } finally {
    loading.value = false;
  }
};

/**
 * Ouvre le dossier
 */
const openFolder = async () => {
  try {
    const result = await window.go.main.App.OpenLeopardBaseFolder();
    if (!result.success) {
      alert('Erreur: ' + result.error);
    }
  } catch (err) {
    alert('Erreur: ' + err.message);
  }
};

/**
 * Rafraîchit les infos
 */
const refresh = () => {
  loadInfo();
};

/**
 * Copie le chemin dans le presse-papier
 */
const copyPath = () => {
  navigator.clipboard.writeText(leopardPath.value);
  alert('✅ Chemin copié dans le presse-papier!');
};

// Charger au montage
onMounted(() => {
  loadInfo();
});
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>