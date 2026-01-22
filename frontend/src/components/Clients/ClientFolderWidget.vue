<template>
  <div class="bg-gradient-to-br from-purple-50 via-blue-50 to-indigo-50 dark:from-slate-900/20 dark:via-slate-900/20 dark:to-slate-900/20 rounded-xl shadow-lg border-2 border-slate-700 dark:border-teal-800 overflow-hidden transition-all">
    
    <!-- 🎯 EN-TÊTE CLIQUABLE -->
    <div 
      @click="toggleExpand"
      class="bg-gradient-to-r from-slate-600 to-slate-900 dark:from-slate-700 dark:to-slate-500 p-4 cursor-pointer hover:from-slate-600 hover:to-teal-800 transition-all"
    >
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 backdrop-blur-sm rounded-lg">
            <component :is="statusIcon" :size="24" class="text-white" />
          </div>
          <div>
            <h3 class="font-bold text-white text-lg flex items-center gap-2">
              📁 Dossier Client
              <span v-if="!internalExists" class="text-xs bg-teal-700 text-green-900 px-2 py-0.5 rounded-full font-semibold">
                ✓ Actif
              </span>
            </h3>
            <p class="text-blue-100 text-sm font-mono">{{ leopardNumber || '⏳ En attente' }}</p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <div :class="statusBadgeClass" class="px-3 py-1.5 rounded-full text-xs font-bold shadow-lg">
            {{ statusText }}
          </div>

          <div v-if="hasMissingSubfolders && internalExists" 
               class="px-2 py-1 bg-orange-500/90 text-white rounded-full text-xs font-bold flex items-center gap-1 animate-pulse"
               title="Sous-dossiers manquants">
            <AlertTriangle :size="12" />
            <span>{{ missingCount }}</span>
          </div>

          <div class="p-1.5 bg-white/20 rounded-lg transition-transform duration-300"
               :class="{ 'rotate-180': isExpanded }">
            <ChevronDown :size="20" class="text-white" />
          </div>
        </div>
      </div>

      <Transition name="fade">
        <div v-if="!isExpanded && internalExists && folderInfo" 
             class="mt-2 flex items-center gap-4 text-xs text-blue-100">
          <div class="flex items-center gap-1">
            <Folder :size="14" />
            <span>{{ folderInfo.subfoldersCount || 0 }} dossiers</span>
          </div>
          <div class="flex items-center gap-1">
            <FileText :size="14" />
            <span>{{ folderInfo.filesCount || 0 }} fichiers</span>
          </div>
          <div v-if="hasMissingSubfolders" class="flex items-center gap-1 text-orange-200">
            <AlertTriangle :size="14" />
            <span>{{ missingCount }} à réparer</span>
          </div>
        </div>
      </Transition>
    </div>

    <!-- 🎯 CONTENU DÉROULABLE -->
    <Transition name="expand">
      <div v-show="isExpanded" class="p-4 space-y-4">
        
        <!-- Informations dossier EXISTANT -->
        <div v-if="internalExists && folderInfo" class="space-y-3">
          <!-- Chemin du dossier -->
          <div class="p-3 bg-white dark:bg-gray-800 rounded-lg border border-slate-100 dark:border-slate-900 shadow-sm">
            <div class="flex items-start gap-2">
              <FolderOpen :size="16" class="text-slate-600 dark:text-teal-400 flex-shrink-0 mt-0.5" />
              <div class="flex-1 min-w-0">
                <p class="text-xs text-gray-500 dark:text-gray-400 mb-1 font-semibold">Emplacement du dossier</p>
                <p class="text-xs text-gray-700 dark:text-gray-300 font-mono break-all bg-gray-50 dark:bg-gray-900 p-2 rounded border border-gray-200 dark:border-gray-700">
                  📂 {{ folderInfo.path }}
                </p>
              </div>
            </div>
          </div>

          <!-- Statistiques du dossier -->
          <div class="grid grid-cols-2 gap-3">
            <div class="p-3 bg-gradient-to-br from-stone-50 to-blue-100 dark:from-stone-900/30 dark:to-stone-800/30 rounded-lg border border-stone-200 dark:border-stone-700">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-xs text-teal-700 dark:text-teal-700 font-semibold mb-1">Sous-dossiers</p>
                  <p class="text-xl font-bold text-slate-900 dark:text-slate-300">{{ folderInfo.subfoldersCount || 0 }}</p>
                </div>
                <Folder :size="32" class="text-slate-400 dark:text-slate-300 opacity-50" />
              </div>
            </div>

            <div class="p-3 bg-gradient-to-br from-stone-50 to-blue-100 dark:from-stone-900/30 dark:to-stone-800/30 rounded-lg border border-stone-200 dark:border-stone-700">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-xs text-teal-700 dark:text-sky-700 font-semibold mb-1">Fichiers</p>
                  <p class="text-xl font-bold text-slate-900 dark:text-slate-300">{{ folderInfo.filesCount || 0 }}</p>
                </div>
                <FileText :size="32" class="text-green-400 dark:text-green-500 opacity-50" />
              </div>
            </div>
          </div>

          <!-- 🆕 LISTE DES SOUS-DOSSIERS DÉROULABLES -->
          <div class="p-3 bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700">
            <div class="flex items-center justify-between mb-3">
              <p class="text-xs text-gray-500 dark:text-gray-400 font-semibold flex items-center gap-1">
                <Layers :size="14" />
                Structure du dossier
              </p>
              
              <button v-if="hasMissingSubfolders"
                      @click.stop="handleRepairStructure"
                      :disabled="isLoading"
                      class="flex items-center gap-1 px-2 py-1 bg-orange-500 hover:bg-orange-600 text-white rounded text-xs font-semibold transition-all disabled:opacity-50">
                <Wrench v-if="!isLoading" :size="12" />
                <Loader2 v-else :size="12" class="animate-spin" />
                <span>{{ isLoading ? 'Réparation...' : 'Tout réparer' }}</span>
              </button>
            </div>
<div class="space-y-1">
  <!-- Composant sous-dossier pour chaque entrée -->
  <SubfolderItem
    v-for="subfolder in folderStructure"
    :key="subfolder.name"
    :subfolder="subfolder"
    :leopard-number="leopardNumber"
    :client-path="folderInfo.path"
    :client-data="client"
    @create-subfolder="handleCreateSubfolder"
    @create-nested="handleCreateNestedFolder"
    @refresh="loadFolderStructure"
  />
</div>
          </div>
        </div>

        <!-- Avertissement si PAS de dossier -->
        <div v-else-if="!internalExists" class="p-4 bg-gradient-to-r from-yellow-50 to-orange-50 dark:from-yellow-900/20 dark:to-orange-900/20 rounded-lg border-2 border-yellow-300 dark:border-yellow-700">
          <div class="flex items-start gap-3">
            <div class="p-2 bg-yellow-100 dark:bg-yellow-900/50 rounded-lg flex-shrink-0">
              <AlertTriangle :size="20" class="text-yellow-600 dark:text-yellow-400" />
            </div>
            <div>
              <p class="text-sm font-bold text-yellow-900 dark:text-yellow-100 mb-1">
                Dossier non créé
              </p>
              <p class="text-xs text-yellow-700 dark:text-yellow-300">
                Le dossier physique n'existe pas encore. Cliquez sur "Créer le dossier" pour initialiser la structure complète avec tous les sous-dossiers nécessaires.
              </p>
            </div>
          </div>
        </div>

        <!-- Boutons d'action -->
        <div class="flex flex-wrap gap-2">
          <button
            v-if="internalExists"
            @click.stop="handleOpenFolder"
            :disabled="isLoading"
            class="flex-1 flex items-center justify-center gap-2 px-4 py-2.5 bg-gradient-to-r from-slate-600 to-teal-700 hover:from-slate-600 hover:to-teal-800 text-white rounded-lg font-semibold shadow-md hover:shadow-lg transition-all disabled:opacity-50 disabled:cursor-not-allowed text-sm"
          >
            <FolderOpen :size="18" />
            <span>Ouvrir le dossier</span>
            <ExternalLink :size="14" class="opacity-70" />
          </button>
          <!-- ✅ NOUVEAU BOUTON TENUE DE DOSSIER -->
  <button
    v-if="internalExists"
    @click.stop="showTenueDossier = true"
    :disabled="isLoading"
    class="flex-1 flex items-center justify-center gap-2 px-4 py-2.5 bg-gradient-to-r from-neutral-400 to-neutral-700 dark: hover:from-neutral-400 to-neutral-700 dark:hover:from-neutral-400 dark:hover:to-neutral-800 text-white rounded-lg font-semibold shadow-md hover:shadow-lg transition-all disabled:opacity-50 disabled:cursor-not-allowed text-sm"
  >
    <ClipboardCheck :size="18" />
    <span>Tenue de dossier</span>
  </button>




          <button
            v-else
            @click.stop="handleCreateFolder"
            :disabled="isLoading || !leopardNumber"
            class="flex-1 flex items-center justify-center gap-2 px-4 py-2.5 bg-gradient-to-r from-green-600 to-green-700 hover:from-green-700 hover:to-green-800 text-white rounded-lg font-semibold shadow-md hover:shadow-lg transition-all disabled:opacity-50 disabled:cursor-not-allowed text-sm"
          >
            <FolderPlus v-if="!isLoading" :size="18" />
            <Loader2 v-else :size="18" class="animate-spin" />
            <span>{{ isLoading ? 'Création...' : 'Créer le dossier' }}</span>
          </button>

          <button
            @click.stop="handleRefresh"
            :disabled="isLoading"
            class="flex items-center justify-center gap-2 px-4 py-2.5 bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 text-gray-700 dark:text-gray-300 rounded-lg font-medium transition-all disabled:opacity-50 shadow-sm text-sm"
            title="Rafraîchir les informations"
          >
            <RefreshCw :size="18" :class="{ 'animate-spin': isLoading }" />
          </button>

          <button
            v-if="internalExists"
            @click.stop="showDetails = !showDetails"
            class="flex items-center justify-center gap-2 px-4 py-2.5 bg-stone-100 dark:bg-purple-900/40 hover:bg-stone-200 dark:hover:bg-stone-700 text-stone-600 dark:text-stone-300 rounded-lg font-medium transition-all shadow-sm text-sm"
            title="Détails du dossier"
          >
            <Info :size="18" />
          </button>
        </div>
<!-- ✅ COMPOSANT MODAL TENUE DE DOSSIER -->
<GestionTenueDossier
  :is-open="showTenueDossier"
  :client="client"
  :leopard-number="leopardNumber"
  @close="showTenueDossier = false"
  @refresh="handleRefresh"
/>
        <!-- Panneau de détails -->
        <Transition name="slide-down">
          <div v-if="showDetails && internalExists" class="p-4 bg-purple-50 dark:bg-purple-900/20 rounded-lg border border-purple-200 dark:border-purple-800 space-y-2">
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-600 dark:text-gray-400">Créé le</span>
              <span class="font-mono font-semibold text-gray-900 dark:text-gray-100">{{ creationDate }}</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-600 dark:text-gray-400">Dernier accès</span>
              <span class="font-mono font-semibold text-gray-900 dark:text-gray-100">Maintenant</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-600 dark:text-gray-400">Numéro Leopard</span>
              <span class="font-mono font-semibold text-gray-900 dark:text-gray-100">{{ leopardNumber }}</span>
            </div>
          </div>
        </Transition>

      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue';
import { 
  FolderCheck, 
  FolderPlus, 
  FolderOpen, 
  RefreshCw, 
  AlertTriangle,
  Loader2,
  ExternalLink,
  Info,
  Folder,
  FileText,
  Layers,
  CheckCircle,    
  AlertCircle,    
  Wrench,
  ChevronDown,
  ClipboardCheck  
} from 'lucide-vue-next';
import { createClientFolder, openClientFolder } from '@/utils/clientFolderManager';
import SubfolderItem from './SubfolderItem.vue';
import GestionTenueDossier from '../Files/GestionTenueDossier.vue';



const getInitialExpandedState = () => {
  const saved = localStorage.getItem('leopard_folder_widget_expanded');
  return saved === null ? true : saved === 'true';
};

const isExpanded = ref(getInitialExpandedState());
const showTenueDossier = ref(false);
const toggleExpand = () => {
  isExpanded.value = !isExpanded.value;
};

watch(isExpanded, (newValue) => {
  localStorage.setItem('leopard_folder_widget_expanded', newValue.toString());
});

const isCreatingSubfolder = ref(false);

const props = defineProps({
  leopardNumber: {
    type: String,
    required: true
  },
  folderExists: {
    type: Boolean,
    default: false
  },
  client: { type: Object, required: true }
});

const emit = defineEmits(['folder-created', 'refresh']);

const internalExists = ref(false);
const isLoading = ref(false);
const showDetails = ref(false);
const folderInfo = ref(null);

// 🆕 Structure enrichie des sous-dossiers avec métadonnées
const folderStructure = ref([]);

const creationDate = computed(() => {
  const today = new Date();
  return today.toLocaleDateString('fr-CA');
});

const statusIcon = computed(() => {
  return internalExists.value ? FolderCheck : FolderPlus;
});

const statusText = computed(() => {
  return internalExists.value ? 'Dossier créé' : 'Non créé';
});

const statusBadgeClass = computed(() => {
  return internalExists.value 
    ? 'bg-green-500 text-white'
    : 'bg-yellow-500 text-yellow-900';
});

const missingCount = computed(() => {
  return folderStructure.value.filter(f => !f.exists).length;
});

const hasMissingSubfolders = computed(() => {
  return folderStructure.value.some(f => !f.exists);
});

const syncWithDisk = async () => {
  if (!props.leopardNumber || props.leopardNumber.includes('XXX')) {
    console.log('⏭️ Numéro Leopard invalide, sync ignoré');
    return;
  }
  
  isLoading.value = true;
  try {
    const exists = await window.go.main.App.ClientFolderExists(props.leopardNumber);
    internalExists.value = exists;
    
    if (exists) {
      const info = await window.go.main.App.GetClientFolderInfo(props.leopardNumber);
      folderInfo.value = info;
      await loadFolderStructure();
    }
    
  } catch (e) {
    console.error("❌ Erreur sync disque:", e);
  } finally {
    isLoading.value = false;
  }
};

const handleCreateFolder = async () => {
  if (!props.leopardNumber) {
    alert('Aucun numéro Leopard disponible');
    return;
  }
  
  isLoading.value = true;
  try {
    const result = await createClientFolder({
      no_dossier_leopard: props.leopardNumber
    });

    if (result.success) {
      internalExists.value = true;
      emit('folder-created', result);
      await syncWithDisk();
      alert(`✅ Dossier créé avec succès!\n\nEmplacement:\n${result.path}`);
    } else {
      alert(`❌ Erreur: ${result.error}`);
    }
  } catch (e) {
    console.error('❌ Erreur création:', e);
    alert(`❌ Erreur lors de la création:\n${e.message}`);
  } finally {
    isLoading.value = false;
  }
};

const handleOpenFolder = async () => {
  if (!internalExists.value && !props.folderExists) {
    alert('Le dossier n\'existe pas encore. Créez-le d\'abord.');
    return;
  }
  
  isLoading.value = true;
  try {
    const result = await openClientFolder(props.leopardNumber);
    
    if (!result.success) {
      alert(`❌ Impossible d'ouvrir le dossier:\n${result.error || 'Erreur inconnue'}`);
    }
  } catch (e) {
    console.error('❌ Erreur ouverture:', e);
    alert(`❌ Erreur:\n${e.message}`);
  } finally {
    isLoading.value = false;
  }
};

const handleRefresh = async () => {
  await syncWithDisk();
  emit('refresh');
};

const handleRepairStructure = async () => {
  if (!props.leopardNumber) return;
  
  if (!confirm('🔧 Réparer la structure?\n\nCette opération créera UNIQUEMENT les sous-dossiers manquants.\nLes dossiers existants ne seront PAS touchés.')) {
    return;
  }
  
  isLoading.value = true;
  
  try {
    const result = await window.go.main.App.RepairClientFolderStructure(props.leopardNumber);
    
    if (result.success) {
      await loadFolderStructure();
      
      let message = '✅ Réparation terminée!\n\n';
      
      if (result.created && result.created.length > 0) {
        message += `📁 Créés:\n${result.created.map(f => '  • ' + f).join('\n')}\n\n`;
      }
      
      if (result.alreadyExists && result.alreadyExists.length > 0) {
        message += `ℹ️ Déjà existants (ignorés):\n${result.alreadyExists.map(f => '  • ' + f).join('\n')}\n\n`;
      }
      
      if (result.errors && result.errors.length > 0) {
        message += `⚠️ Erreurs:\n${result.errors.map(e => '  • ' + e).join('\n')}`;
      }
      
      alert(message);
    } else {
      alert(`❌ Erreur: ${result.error}`);
    }
  } catch (e) {
    console.error('❌ Erreur réparation:', e);
    alert(`❌ Erreur: ${e.message}`);
  } finally {
    isLoading.value = false;
  }
};

const handleCreateSubfolder = async (subfolderName) => {
  if (!props.leopardNumber) return;
  
  isCreatingSubfolder.value = true;
  
  try {
    const result = await window.go.main.App.CreateSubfolder(
      props.leopardNumber,
      subfolderName
    );
    
    if (result.success) {
      await loadFolderStructure();
      
      if (result.error) {
        alert(`ℹ️ ${result.error}`);
      }
    } else {
      alert(`⚠️ ${result.error}`);
    }
  } catch (e) {
    console.error('❌ Erreur création sous-dossier:', e);
    alert(`❌ Erreur: ${e.message}`);
  } finally {
    isCreatingSubfolder.value = false;
  }
};

// 🆕 Création de sous-dossier imbriqué
const handleCreateNestedFolder = async ({ parentFolder, subfolderName }) => {
  if (!props.leopardNumber) return;
  
  isCreatingSubfolder.value = true;
  
  try {
    const fullPath = `${parentFolder}/${subfolderName}`;
    const result = await window.go.main.App.CreateSubfolder(
      props.leopardNumber,
      fullPath
    );
    
    if (result.success) {
      await loadFolderStructure();
    } else {
      alert(`⚠️ ${result.error}`);
    }
  } catch (e) {
    console.error('❌ Erreur création sous-dossier imbriqué:', e);
    alert(`❌ Erreur: ${e.message}`);
  } finally {
    isCreatingSubfolder.value = false;
  }
};



// 🆕 Charge la structure enrichie avec détails
const loadFolderStructure = async () => {
  if (!props.leopardNumber || props.leopardNumber.includes('XXX')) {
    return;
  }
  
  try {
    const structure = await window.go.main.App.GetDetailedFolderStructure(props.leopardNumber);
    
    if (structure.success) {
      folderStructure.value = structure.folders || [];
      console.log('📋 Structure enrichie chargée:', structure);
    }
  } catch (e) {
    console.error('❌ Erreur chargement structure:', e);
  }
};

onMounted(() => {
  syncWithDisk();
});

watch(() => props.leopardNumber, (newVal) => {
  if (newVal && !newVal.includes('XXX')) {
    syncWithDisk();
  }
}, { immediate: true });
</script>

<style scoped>
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.7; }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.expand-enter-active,
.expand-leave-active {
  transition: all 0.3s ease;
  overflow: hidden;
}

.expand-enter-from,
.expand-leave-to {
  max-height: 0;
  opacity: 0;
}

.expand-enter-to,
.expand-leave-from {
  max-height: 2000px;
  opacity: 1;
}

.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease;
}

.slide-down-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>