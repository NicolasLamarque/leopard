<template>
  <div class="border rounded-lg overflow-hidden transition-all"
       :class="subfolder.exists 
         ? 'bg-green-50 dark:bg-green-900/20 border-green-200 dark:border-green-800' 
         : 'bg-yellow-50 dark:bg-yellow-900/20 border-yellow-200 dark:border-yellow-800'">
    
    <!-- 🎯 EN-TÊTE DU SOUS-DOSSIER (cliquable pour déplier) -->
    <div @click="toggleExpand"
         class="flex items-center gap-2 p-2 cursor-pointer hover:bg-black/5 dark:hover:bg-white/5 transition-all">
      
      <!-- Chevron de dépliage -->
      <div class="flex-shrink-0 transition-transform duration-200"
           :class="{ 'rotate-90': isExpanded }">
        <ChevronRight :size="14" class="text-gray-500 dark:text-gray-400" />
      </div>
      
      <!-- Icône de statut -->
      <div class="flex-shrink-0">
        <CheckCircle v-if="subfolder.exists" 
                     :size="16" 
                     class="text-green-500" />
        <AlertCircle v-else 
                     :size="16" 
                     class="text-yellow-500" />
      </div>
      
      <!-- Icône de dossier avec couleur selon type -->
      <component :is="folderIcon" 
                 :size="16" 
                 :class="iconColorClass"
                 class="flex-shrink-0" />
      
      <!-- Nom du sous-dossier -->
      <span class="flex-1 font-medium text-sm text-gray-700 dark:text-gray-300">
        {{ subfolder.name }}
      </span>

      <!-- Badge de comptage (fichiers + sous-dossiers) -->
      <div v-if="subfolder.exists" class="flex items-center gap-2 text-xs">
        <span v-if="subfolder.fileCount > 0" 
              class="flex items-center gap-1 text-gray-600 dark:text-gray-400">
          <FileText :size="12" />
          {{ subfolder.fileCount }}
        </span>
        <span v-if="subfolder.subfolderCount > 0" 
              class="flex items-center gap-1 text-gray-600 dark:text-gray-400">
          <FolderOpen :size="12" />
          {{ subfolder.subfolderCount }}
        </span>
      </div>
      
      <!-- Badge OK ou Bouton Créer -->
      <div class="flex-shrink-0" @click.stop>
        <span v-if="subfolder.exists" 
              class="text-[10px] bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300 px-2 py-0.5 rounded-full font-semibold">
          ✓ OK
        </span>
        <button v-else
                @click="$emit('create-subfolder', subfolder.name)"
                class="text-[10px] bg-orange-500 hover:bg-orange-600 text-white px-2 py-1 rounded-md font-semibold transition-all flex items-center gap-1">
          <FolderPlus :size="10" />
          <span>Créer</span>
        </button>
      </div>
    </div>

    <!-- 🎯 CONTENU DÉROULÉ (sous-dossiers imbriqués + fichiers + actions) -->
    <div v-show="isExpanded && subfolder.exists" class="border-t border-current/10">
        
        <!-- Sous-dossiers imbriqués -->
        <div v-if="subfolder.children && subfolder.children.length > 0" 
             class="p-2 space-y-1 bg-white/30 dark:bg-black/10">
          <div class="text-[10px] uppercase tracking-wide text-gray-500 dark:text-gray-400 font-semibold mb-1 px-2">
            Sous-dossiers
          </div>
          <div v-for="child in subfolder.children" 
               :key="child.name"
               class="ml-4 flex items-center gap-2 p-1.5 rounded bg-white/50 dark:bg-black/20 text-xs">
            <FolderOpen :size="12" class="text-blue-500 flex-shrink-0" />
            <span class="flex-1 text-gray-700 dark:text-gray-300">{{ child.name }}</span>
            <span v-if="child.fileCount > 0" class="text-[10px] text-gray-500 flex items-center gap-1">
              <FileText :size="10" />
              {{ child.fileCount }}
            </span>
          </div>
        </div>

        <!-- Liste des fichiers récents (max 5) -->
        <div v-if="subfolder.files && subfolder.files.length > 0" 
             class="p-2 space-y-1 bg-white/30 dark:bg-black/10">
          <div class="flex items-center justify-between mb-1">
            <div class="text-[10px] uppercase tracking-wide text-gray-500 dark:text-gray-400 font-semibold px-2">
              Fichiers récents
            </div>
            <span class="text-[10px] text-gray-400 px-2">
              {{ subfolder.fileCount }} total
            </span>
          </div>
          <div v-for="(file, idx) in recentFiles" 
               :key="idx"
               class="ml-2 flex items-center gap-2 p-1 rounded hover:bg-white/50 dark:hover:bg-white/5 transition-all cursor-pointer group text-xs">
            <component :is="getFileIcon(file.name)" 
                       :size="12" 
                       class="text-gray-400 group-hover:text-blue-500 flex-shrink-0 transition-colors" />
            <span class="flex-1 text-gray-600 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-gray-100 truncate transition-colors">
              {{ file.name }}
            </span>
            <span class="text-[9px] text-gray-400 group-hover:text-gray-500">
              {{ formatFileSize(file.size) }}
            </span>
          </div>
          <div v-if="subfolder.files.length > 5" 
               class="ml-2 text-[10px] text-gray-400 italic px-2">
            + {{ subfolder.files.length - 5 }} autres fichiers...
          </div>
        </div>

        <!-- Actions spécifiques au dossier -->
        <div class="p-2 flex flex-wrap gap-1 bg-gray-50 dark:bg-gray-900/30 border-t border-current/10">
          
          <!-- Ouvrir ce sous-dossier -->
          <button @click.stop="openSubfolder"
                  class="flex items-center gap-1 px-2 py-1 bg-blue-500 hover:bg-blue-600 text-white rounded text-[10px] font-semibold transition-all">
            <ExternalLink :size="10" />
            <span>Ouvrir</span>
          </button>

          <!-- Actions contextuelles selon le type de dossier -->
          <template v-if="subfolder.name === 'Evaluations'">
            <button @click.stop="showEvaluationCreator = true"
                    class="flex items-center gap-1 px-2 py-1 bg-purple-500 hover:bg-purple-600 text-white rounded text-[10px] font-semibold transition-all">
              <Plus :size="10" />
              <span>Nouvelle évaluation</span>
            </button>
          </template>

          <template v-if="subfolder.name === 'Documents'">
            <button @click.stop="showDocumentUploader = true"
                    class="flex items-center gap-1 px-2 py-1 bg-green-500 hover:bg-green-600 text-white rounded text-[10px] font-semibold transition-all">
              <Upload :size="10" />
              <span>Importer</span>
            </button>
          </template>
          <!-- Dans SubfolderItem.vue, après le bloc "Actions contextuelles" -->



          <!-- Rafraîchir -->
          <button @click.stop="$emit('refresh')"
                  class="flex items-center gap-1 px-2 py-1 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 rounded text-[10px] font-semibold transition-all">
            <RefreshCw :size="10" />
          </button>
        </div>

    </div>

    <!-- 🆕 MODAL: Créateur d'évaluation par discipline -->
    <Teleport to="body">
      <Transition name="fade">
        <div v-if="showEvaluationCreator" 
             @click="showEvaluationCreator = false"
             class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4">
          <div @click.stop 
               class="bg-white dark:bg-gray-800 rounded-xl shadow-2xl max-w-md w-full p-6 space-y-4">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-bold text-gray-900 dark:text-gray-100 flex items-center gap-2">
                <ClipboardList :size="20" class="text-purple-500" />
                Nouvelle évaluation
              </h3>
              <button @click="showEvaluationCreator = false" 
                      class="text-gray-400 hover:text-gray-600">
                <X :size="20" />
              </button>
            </div>

            <div class="space-y-3">
              <div>
                <label class="block text-xs font-semibold text-gray-700 dark:text-gray-300 mb-1">
                  Discipline
                </label>
                <select v-model="selectedDiscipline"
                        class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 text-sm">
                  <option value="">-- Sélectionner --</option>
                  <option value="Ergotherapie">Ergothérapie</option>
                  <option value="Psychologie">Psychologie</option>
                  <option value="Travail_Social">Travail social</option>
                  <option value="Medicale">Médicale</option>
                  <option value="Physiotherapie">Physiothérapie</option>
                  <option value="Autre">Autre...</option>
                </select>
              </div>

              <div v-if="selectedDiscipline === 'Autre'">
                <label class="block text-xs font-semibold text-gray-700 dark:text-gray-300 mb-1">
                  Nom personnalisé
                </label>
                <input v-model="customDisciplineName"
                       type="text"
                       placeholder="Ex: Orthophonie"
                       class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 text-sm">
              </div>
            </div>

            <div class="flex gap-2 pt-2">
              <button @click="createEvaluationFolder"
                      :disabled="!canCreateEvaluation"
                      class="flex-1 px-4 py-2 bg-purple-500 hover:bg-purple-600 disabled:bg-gray-300 disabled:cursor-not-allowed text-white rounded-lg font-semibold text-sm transition-all">
                Créer le sous-dossier
              </button>
              <button @click="showEvaluationCreator = false"
                      class="px-4 py-2 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 rounded-lg font-medium text-sm transition-all">
                Annuler
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { 
  ChevronRight,
  CheckCircle, 
  AlertCircle,
  FolderOpen,
  FolderPlus,
  FileText,
  ExternalLink,
  RefreshCw,
  Plus,
  Upload,
  File,
  FileImage,
  FileSpreadsheet,
  ClipboardList,
  X
} from 'lucide-vue-next';


const props = defineProps({
  subfolder: {
    type: Object,
    required: true
  },
  leopardNumber: {
    type: String,
    required: true
  },
  clientPath: {
    type: String,
    required: true
  },
  clientData: { 
    type: Object, required: true }, 
});

const emit = defineEmits(['create-subfolder', 'create-nested', 'refresh']);

const isExpanded = ref(false);
const showEvaluationCreator = ref(false);
const showDocumentUploader = ref(false);
const selectedDiscipline = ref('');
const customDisciplineName = ref('');

const toggleExpand = () => {
  if (props.subfolder.exists) {
    isExpanded.value = !isExpanded.value;
  }
};

// Icône selon le type de dossier
const folderIcon = computed(() => {
  const icons = {
    'Evaluations': ClipboardList,
    'Documents': FileText,
    'Rapports': File,
    'Correspondance': FileText,
    'Notes': FileText,
    'Contrats': FileSpreadsheet
  };
  return icons[props.subfolder.name] || FolderOpen;
});

// Couleur selon le type
const iconColorClass = computed(() => {
  const colors = {
    'Evaluations': 'text-purple-500',
    'Documents': 'text-blue-500',
    'Rapports': 'text-green-500',
    'Correspondance': 'text-orange-500',
    'Notes': 'text-yellow-600',
    'Contrats': 'text-red-500'
  };
  return colors[props.subfolder.name] || 'text-gray-500';
});

// Fichiers récents (max 5)
const recentFiles = computed(() => {
  if (!props.subfolder.files) return [];
  return props.subfolder.files.slice(0, 5);
});

// Validation pour créer évaluation
const canCreateEvaluation = computed(() => {
  if (selectedDiscipline.value === 'Autre') {
    return customDisciplineName.value.trim() !== '';
  }
  return selectedDiscipline.value !== '';
});

// Ouvrir le sous-dossier dans l'explorateur
const openSubfolder = async () => {
  try {
    console.log('🔍 Tentative ouverture:', props.subfolder.path);
    const result = await window.go.main.App.OpenFolder(props.subfolder.path);
    
    if (!result.success) {
      console.error('❌ Erreur:', result.error);
      alert(`Impossible d'ouvrir le dossier: ${result.error}`);
    } else {
      console.log('✅ Dossier ouvert:', result.path);
    }
  } catch (e) {
    console.error('❌ Erreur ouverture sous-dossier:', e);
    alert(`Erreur: ${e.message}`);
  }
};

// Créer un dossier d'évaluation
const createEvaluationFolder = () => {
  const disciplineName = selectedDiscipline.value === 'Autre' 
    ? customDisciplineName.value.trim()
    : selectedDiscipline.value;
  
  const folderName = `${disciplineName}_Evaluation`;
  
  emit('create-nested', {
    parentFolder: props.subfolder.name,
    subfolderName: folderName
  });
  
  // Reset et fermer
  selectedDiscipline.value = '';
  customDisciplineName.value = '';
  showEvaluationCreator.value = false;
};

// Icône selon extension de fichier
const getFileIcon = (filename) => {
  const ext = filename.split('.').pop().toLowerCase();
  const iconMap = {
    'pdf': FileText,
    'doc': FileText,
    'docx': FileText,
    'txt': FileText,
    'jpg': FileImage,
    'jpeg': FileImage,
    'png': FileImage,
    'gif': FileImage,
    'xls': FileSpreadsheet,
    'xlsx': FileSpreadsheet,
    'csv': FileSpreadsheet
  };
  return iconMap[ext] || File;
};

// Formater taille de fichier
const formatFileSize = (bytes) => {
  if (!bytes) return '';
  if (bytes < 1024) return bytes + ' B';
  if (bytes < 1048576) return (bytes / 1024).toFixed(1) + ' KB';
  return (bytes / 1048576).toFixed(1) + ' MB';
};
</script>

<style scoped>
.expand-enter-active,
.expand-leave-active {
  transition: all 0.2s ease;
  overflow: hidden;
}

.expand-enter-from,
.expand-leave-to {
  max-height: 0;
  opacity: 0;
}

.expand-enter-to,
.expand-leave-from {
  max-height: 500px;
  opacity: 1;
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