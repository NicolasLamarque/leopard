<template>
  <div class="bg-white rounded-lg shadow-lg p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <Folder :size="28" class="text-blue-600" />
        <div>
          <h2 class="text-2xl font-bold text-gray-900">Gestionnaire de Dossiers Clients</h2>
          <p class="text-sm text-gray-500">Gestion des dossiers physiques sur le disque</p>
        </div>
      </div>
      
      <button
        @click="refreshFolderList"
        class="flex items-center gap-2 px-4 py-2 bg-blue-50 text-blue-700 hover:bg-blue-100 rounded-lg transition-colors"
      >
        <RefreshCw :size="18" :class="{ 'animate-spin': isRefreshing }" />
        <span>Actualiser</span>
      </button>
    </div>

    <!-- Statistiques -->
    <div class="grid grid-cols-4 gap-4 mb-6">
      <div class="p-4 bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-blue-700 font-medium">Total Dossiers</p>
            <p class="text-2xl font-bold text-blue-900">{{ stats.totalFolders }}</p>
          </div>
          <FolderOpen :size="32" class="text-blue-400 opacity-50" />
        </div>
      </div>
      
      <div class="p-4 bg-gradient-to-br from-green-50 to-green-100 rounded-lg border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-green-700 font-medium">Créés Aujourd'hui</p>
            <p class="text-2xl font-bold text-green-900">{{ stats.createdToday }}</p>
          </div>
          <FolderPlus :size="32" class="text-green-400 opacity-50" />
        </div>
      </div>
      
      <div class="p-4 bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-purple-700 font-medium">Espace Total</p>
            <p class="text-2xl font-bold text-purple-900">{{ stats.totalSize }}</p>
          </div>
          <HardDrive :size="32" class="text-purple-400 opacity-50" />
        </div>
      </div>
      
      <div class="p-4 bg-gradient-to-br from-orange-50 to-orange-100 rounded-lg border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-orange-700 font-medium">Documents</p>
            <p class="text-2xl font-bold text-orange-900">{{ stats.totalFiles }}</p>
          </div>
          <FileText :size="32" class="text-orange-400 opacity-50" />
        </div>
      </div>
    </div>

    <!-- Barre de recherche -->
    <div class="mb-6">
      <div class="relative">
        <Search :size="20" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher un dossier client (nom, prénom, numéro Leopard)..."
          class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
    </div>

    <!-- Liste des dossiers -->
    <div class="space-y-2 max-h-96 overflow-y-auto">
      <div
        v-for="folder in filteredFolders"
        :key="folder.name"
        class="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3 flex-1">
            <div class="p-2 bg-blue-100 rounded-lg">
              <Folder :size="24" class="text-blue-600" />
            </div>
            
            <div class="flex-1">
              <div class="flex items-center gap-2">
                <h3 class="font-semibold text-gray-900">{{ folder.displayName }}</h3>
                <span class="px-2 py-0.5 text-xs font-mono bg-gray-100 text-gray-700 rounded">
                  {{ folder.leopardNumber }}
                </span>
              </div>
              <div class="flex items-center gap-4 mt-1 text-xs text-gray-500">
                <span class="flex items-center gap-1">
                  <FolderOpen :size="12" />
                  {{ folder.subfoldersCount }} sous-dossiers
                </span>
                <span class="flex items-center gap-1">
                  <FileText :size="12" />
                  {{ folder.filesCount }} fichiers
                </span>
                <span class="flex items-center gap-1">
                  <Calendar :size="12" />
                  {{ formatDate(folder.createdAt) }}
                </span>
              </div>
            </div>
          </div>
          
          <div class="flex items-center gap-2">
            <!-- Bouton ouvrir -->
            <button
              @click="openFolder(folder)"
              class="p-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
              title="Ouvrir le dossier"
            >
              <ExternalLink :size="18" />
            </button>
            
            <!-- Bouton voir détails -->
            <button
              @click="showFolderDetails(folder)"
              class="p-2 text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
              title="Voir les détails"
            >
              <Eye :size="18" />
            </button>
            
            <!-- Bouton options -->
            <button
              @click="showFolderOptions(folder)"
              class="p-2 text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
              title="Options"
            >
              <MoreVertical :size="18" />
            </button>
          </div>
        </div>
      </div>
      
      <!-- Message si aucun dossier -->
      <div v-if="filteredFolders.length === 0" class="text-center py-12">
        <FolderX :size="48" class="mx-auto text-gray-300 mb-4" />
        <p class="text-gray-500">
          {{ searchQuery ? 'Aucun dossier trouvé' : 'Aucun dossier client créé' }}
        </p>
      </div>
    </div>

    <!-- Modal de détails du dossier -->
    <div
      v-if="selectedFolder"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click.self="selectedFolder = null"
    >
      <div class="bg-white rounded-lg shadow-xl p-6 max-w-2xl w-full mx-4 max-h-[80vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-xl font-bold text-gray-900">Détails du Dossier</h3>
          <button
            @click="selectedFolder = null"
            class="p-2 text-gray-500 hover:bg-gray-100 rounded-lg"
          >
            <X :size="20" />
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="text-sm font-medium text-gray-600">Nom complet</label>
            <p class="text-lg font-semibold text-gray-900">{{ selectedFolder.displayName }}</p>
          </div>
          
          <div>
            <label class="text-sm font-medium text-gray-600">Numéro Leopard</label>
            <p class="font-mono text-gray-900">{{ selectedFolder.leopardNumber }}</p>
          </div>
          
          <div>
            <label class="text-sm font-medium text-gray-600">Chemin complet</label>
            <p class="text-sm text-gray-700 bg-gray-50 p-2 rounded font-mono break-all">
              {{ selectedFolder.path }}
            </p>
          </div>
          
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="text-sm font-medium text-gray-600">Sous-dossiers</label>
              <p class="text-2xl font-bold text-blue-600">{{ selectedFolder.subfoldersCount }}</p>
            </div>
            <div>
              <label class="text-sm font-medium text-gray-600">Fichiers</label>
              <p class="text-2xl font-bold text-green-600">{{ selectedFolder.filesCount }}</p>
            </div>
          </div>
          
          <div class="flex gap-2 mt-6">
            <button
              @click="openFolder(selectedFolder); selectedFolder = null"
              class="flex-1 flex items-center justify-center gap-2 px-4 py-3 bg-blue-600 text-white hover:bg-blue-700 rounded-lg transition-colors"
            >
              <ExternalLink :size="18" />
              <span>Ouvrir dans l'explorateur</span>
            </button>
          </div>
        </div>
        <WidgetFoldersShowInfos :leopardPath="selectedFolder.path" class="mt-6" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import {
  Folder,
  FolderOpen,
  FolderPlus,
  FolderX,
  RefreshCw,
  Search,
  Eye,
  MoreVertical,
  ExternalLink,
  X,
  FileText,
  Calendar,
  HardDrive
} from 'lucide-vue-next';
import { listClientFolders, openClientFolder } from '@/utils/clientFolderManager';
import WidgetFoldersShowInfos from '@/widgets/widgetFoldersShowInfos.vue';

const searchQuery = ref('');
const isRefreshing = ref(false);
const folders = ref([]);
const selectedFolder = ref(null);

const stats = ref({
  totalFolders: 0,
  createdToday: 0,
  totalSize: '0 MB',
  totalFiles: 0
});

/**
 * Dossiers filtrés par recherche
 */
const filteredFolders = computed(() => {
  if (!searchQuery.value) return folders.value;
  
  const query = searchQuery.value.toLowerCase();
  return folders.value.filter(folder => 
    folder.displayName.toLowerCase().includes(query) ||
    folder.leopardNumber.toLowerCase().includes(query)
  );
});

/**
 * Charge la liste des dossiers
 */
const loadFolders = async () => {
  try {
    const folderNames = await listClientFolders();
    
    // Transformer les noms en objets avec détails
    folders.value = await Promise.all(
      folderNames.map(async (name) => {
        const leopardNumber = name.split(' ')[0];
        const displayName = name.replace(leopardNumber + ' - ', '');
        
        const info = await window.go.main.App.GetClientFolderInfo(leopardNumber);
        
        return {
          name,
          leopardNumber,
          displayName,
          path: info?.path || '',
          subfoldersCount: info?.subfoldersCount || 0,
          filesCount: info?.filesCount || 0,
          createdAt: new Date() // À récupérer du système de fichiers
        };
      })
    );
    
    // Calculer les stats
    updateStats();
  } catch (error) {
    console.error('❌ Erreur chargement dossiers:', error);
  }
};

/**
 * Met à jour les statistiques
 */
const updateStats = () => {
  stats.value.totalFolders = folders.value.length;
  
  const today = new Date().toDateString();
  stats.value.createdToday = folders.value.filter(
    f => f.createdAt.toDateString() === today
  ).length;
  
  stats.value.totalFiles = folders.value.reduce((sum, f) => sum + f.filesCount, 0);
  
  // Taille totale (à implémenter côté backend)
  stats.value.totalSize = '125 MB';
};

/**
 * Actualise la liste
 */
const refreshFolderList = async () => {
  isRefreshing.value = true;
  await loadFolders();
  setTimeout(() => {
    isRefreshing.value = false;
  }, 500);
};

/**
 * Ouvre un dossier
 */
const openFolder = async (folder) => {
  await openClientFolder(folder.leopardNumber, folder.displayName);
};

/**
 * Affiche les détails d'un dossier
 */
const showFolderDetails = (folder) => {
  selectedFolder.value = folder;
};

/**
 * Options du dossier
 */
const showFolderOptions = (folder) => {
  // À implémenter: menu contextuel avec options
  console.log('Options pour:', folder);
};

/**
 * Formate une date
 */
const formatDate = (date) => {
  return new Intl.DateTimeFormat('fr-CA', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date);
};

// Charger au montage
onMounted(() => {
  loadFolders();
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