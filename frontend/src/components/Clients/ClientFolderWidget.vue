<template>
  <div class="rounded-xl border border-slate-700/60 bg-slate-900 shadow-md overflow-hidden transition-all">

    <!-- EN-TÊTE CLIQUABLE -->
    <div
      @click="toggleExpand"
      class="flex items-center justify-between px-4 py-3 cursor-pointer hover:bg-slate-800/60 transition-colors"
    >
      <div class="flex items-center gap-3">
        <div class="p-1.5 rounded-lg bg-slate-800 border border-slate-700">
          <component :is="statusIcon" :size="18" class="text-teal-400" />
        </div>
        <div>
          <p class="text-sm font-semibold text-slate-100 leading-tight">Dossier Client</p>
          <p class="text-xs text-slate-500 font-mono mt-0.5">{{ leopardNumber || '—' }}</p>
        </div>
      </div>

      <div class="flex items-center gap-2">
        <!-- Statistiques inline quand replié -->
        <Transition name="fade">
          <div v-if="!isExpanded && internalExists && folderInfo"
               class="hidden sm:flex items-center gap-3 text-xs text-slate-500 mr-2">
            <span class="flex items-center gap-1">
              <Folder :size="13" class="text-slate-600" />
              {{ folderInfo.subfoldersCount || 0 }}
            </span>
            <span class="flex items-center gap-1">
              <FileText :size="13" class="text-slate-600" />
              {{ folderInfo.filesCount || 0 }}
            </span>
          </div>
        </Transition>

        <!-- Badge manquants -->
        <span v-if="hasMissingSubfolders && internalExists"
              class="flex items-center gap-1 px-2 py-0.5 rounded-full bg-amber-500/15 text-amber-400 text-xs font-medium border border-amber-500/30">
          <AlertTriangle :size="11" />
          {{ missingCount }}
        </span>

        <!-- Badge statut -->
        <span :class="[
          'px-2.5 py-0.5 rounded-full text-xs font-semibold border',
          internalExists
            ? 'bg-teal-500/10 text-teal-400 border-teal-500/30'
            : 'bg-amber-500/10 text-amber-400 border-amber-500/30'
        ]">
          {{ statusText }}
        </span>

        <ChevronDown :size="16" class="text-slate-500 transition-transform duration-200"
                     :class="{ 'rotate-180': isExpanded }" />
      </div>
    </div>

    <!-- CONTENU DÉROULABLE -->
    <Transition name="expand">
      <div v-show="isExpanded" class="border-t border-slate-700/50">

        <!-- Dossier EXISTANT -->
        <div v-if="internalExists && folderInfo" class="p-4 space-y-3">

          <!-- Chemin -->
          <div class="flex items-start gap-2 p-2.5 rounded-lg bg-slate-800/50 border border-slate-700/40">
            <FolderOpen :size="14" class="text-slate-500 flex-shrink-0 mt-0.5" />
            <p class="text-xs font-mono text-slate-400 break-all leading-relaxed">{{ folderInfo.path }}</p>
          </div>

          <!-- Statistiques -->
          <div class="grid grid-cols-2 gap-2">
            <div class="p-3 rounded-lg bg-slate-800/50 border border-slate-700/40 flex items-center justify-between">
              <div>
                <p class="text-xs text-slate-500 mb-0.5">Sous-dossiers</p>
                <p class="text-lg font-bold text-slate-200">{{ folderInfo.subfoldersCount || 0 }}</p>
              </div>
              <Folder :size="24" class="text-slate-700" />
            </div>
            <div class="p-3 rounded-lg bg-slate-800/50 border border-slate-700/40 flex items-center justify-between">
              <div>
                <p class="text-xs text-slate-500 mb-0.5">Fichiers</p>
                <p class="text-lg font-bold text-slate-200">{{ folderInfo.filesCount || 0 }}</p>
              </div>
              <FileText :size="24" class="text-slate-700" />
            </div>
          </div>

          <!-- Structure des sous-dossiers -->
          <div class="rounded-lg border border-slate-700/40 bg-slate-800/30 overflow-hidden">
            <div class="flex items-center justify-between px-3 py-2 border-b border-slate-700/40">
              <span class="text-xs text-slate-500 font-medium flex items-center gap-1.5">
                <Layers :size="13" />
                Structure
              </span>
              <button v-if="hasMissingSubfolders"
                      @click.stop="handleRepairStructure"
                      :disabled="isLoading"
                      class="flex items-center gap-1.5 px-2.5 py-1 rounded-md bg-amber-500/15 hover:bg-amber-500/25 text-amber-400 border border-amber-500/30 text-xs font-medium transition-colors disabled:opacity-50">
                <Wrench v-if="!isLoading" :size="11" />
                <Loader2 v-else :size="11" class="animate-spin" />
                {{ isLoading ? 'Réparation...' : 'Réparer' }}
              </button>
            </div>
            <div class="p-2 space-y-0.5">
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
                @open-file="handleOpenFile"
              />
            </div>
          </div>
        </div>

        <!-- PAS de dossier -->
        <div v-else-if="!internalExists" class="m-4 p-3.5 rounded-lg bg-amber-500/8 border border-amber-500/25 flex items-start gap-3">
          <AlertTriangle :size="16" class="text-amber-400 flex-shrink-0 mt-0.5" />
          <div>
            <p class="text-sm font-semibold text-amber-300 leading-tight">Dossier non créé</p>
            <p class="text-xs text-slate-500 mt-1 leading-relaxed">
              La structure physique n'existe pas encore. Créez-la pour initialiser tous les sous-dossiers.
            </p>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex flex-wrap gap-2 px-4 pb-4" :class="{ 'pt-0': internalExists, 'pt-0': !internalExists }">

          <button
            v-if="internalExists"
            @click.stop="handleOpenFolder"
            :disabled="isLoading"
            class="flex-1 flex items-center justify-center gap-2 px-3 py-2 rounded-lg bg-slate-700 hover:bg-slate-600 text-slate-200 text-xs font-medium border border-slate-600/60 transition-colors disabled:opacity-50"
          >
            <FolderOpen :size="15" />
            Ouvrir
            <ExternalLink :size="12" class="text-slate-500" />
          </button>

          <button
            v-if="internalExists"
            @click.stop="showTenueDossier = true"
            :disabled="isLoading"
            class="flex-1 flex items-center justify-center gap-2 px-3 py-2 rounded-lg bg-teal-500/10 hover:bg-teal-500/20 text-teal-400 text-xs font-medium border border-teal-500/30 transition-colors disabled:opacity-50"
          >
            <ClipboardCheck :size="15" />
            Tenue de dossier
          </button>

          <button
            v-else
            @click.stop="handleCreateFolder"
            :disabled="isLoading || !leopardNumber"
            class="flex-1 flex items-center justify-center gap-2 px-3 py-2 rounded-lg bg-teal-600 hover:bg-teal-500 text-white text-xs font-semibold transition-colors disabled:opacity-50"
          >
            <FolderPlus v-if="!isLoading" :size="15" />
            <Loader2 v-else :size="15" class="animate-spin" />
            {{ isLoading ? 'Création...' : 'Créer le dossier' }}
          </button>

          <button
            @click.stop="handleRefresh"
            :disabled="isLoading"
            class="flex items-center justify-center gap-1.5 px-3 py-2 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-400 text-xs border border-slate-700/60 transition-colors disabled:opacity-50"
            title="Rafraîchir"
          >
            <RefreshCw :size="15" :class="{ 'animate-spin': isLoading }" />
          </button>

          <button
            v-if="internalExists"
            @click.stop="showDetails = !showDetails"
            class="flex items-center justify-center px-3 py-2 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-400 border border-slate-700/60 transition-colors"
            title="Détails"
          >
            <Info :size="15" />
          </button>
        </div>

        <!-- Panneau détails -->
        <Transition name="slide-down">
          <div v-if="showDetails && internalExists"
               class="mx-4 mb-4 p-3 rounded-lg bg-slate-800/50 border border-slate-700/40 space-y-2">
            <div class="flex items-center justify-between text-xs">
              <span class="text-slate-500">Créé le</span>
              <span class="font-mono text-slate-300">{{ creationDate }}</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-slate-500">Numéro Leopard</span>
              <span class="font-mono text-teal-400">{{ leopardNumber }}</span>
            </div>
          </div>
        </Transition>

        <!-- Modal Tenue de dossier -->
        <GestionTenueDossier
          :is-open="showTenueDossier"
          :client="client"
          :leopard-number="leopardNumber"
          @close="showTenueDossier = false"
          @refresh="handleRefresh"
        />

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

const folderStructure = ref([]);

const creationDate = computed(() => {
  const today = new Date();
  return today.toLocaleDateString('fr-CA');
});

const statusIcon = computed(() => {
  return internalExists.value ? FolderCheck : FolderPlus;
});

const statusText = computed(() => {
  return internalExists.value ? 'Actif' : 'Non créé';
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

const handleOpenFile = async (filePath) => {
  if (!filePath) return;
  
  isLoading.value = true;
  try {
    const result = await window.go.main.App.OpenFile(filePath);
    if (!result.success) {
      alert(`Erreur : ${result.error}`);
    }
  } catch (e) {
    console.error("Erreur lors de l'ouverture du fichier:", e);
  } finally {
    isLoading.value = false;
  }
};

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
.animate-spin {
  animation: spin 1s linear infinite;
}
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.expand-enter-active,
.expand-leave-active {
  transition: all 0.25s ease;
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
  transition: all 0.2s ease;
}
.slide-down-enter-from,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.15s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>