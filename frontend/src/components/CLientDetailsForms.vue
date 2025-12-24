<template>
  <div class="bg-white dark:bg-gray-900 p-6 rounded-xl shadow-lg dark:shadow-gray-950/30 border border-gray-100 dark:border-gray-800 transition-colors">
    <FormKit
      type="form"
      v-model="formData"
      submit-label="Sauvegarder les modifications"
      @submit="submitHandler"
      :submit-attrs="{
        inputClass: 'bg-blue-600 hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 text-white font-medium px-6 py-2.5 rounded-lg transition-colors cursor-pointer'
      }"
    >
      <!-- Informations personnelles -->
      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <User :size="20" class="text-blue-600 dark:text-blue-400" />
          Informations personnelles
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <FormKit
            type="text"
            name="prenom"
            label="Prénom"
            validation="required"
            outer-class="formkit-outer-dark"
            label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
            input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent transition-colors"
          />
          <FormKit 
            type="text" 
            name="nom" 
            label="Nom" 
            validation="required"
            outer-class="formkit-outer-dark"
            label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
            input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent transition-colors"
          />
          <FormKit
            type="tel"
            name="telephone"
            label="Téléphone"
            placeholder="514-555-0123"
            validation="required|matches:/^[0-9]{10}$/"
            :validation-messages="{
              matches: 'Le numéro doit contenir exactement 10 chiffres.',
            }"
            outer-class="formkit-outer-dark"
            label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
            input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent transition-colors"
            messages-class="text-red-600 dark:text-red-400 text-xs mt-1"
          />
          <FormKit 
            type="email" 
            name="email" 
            label="Courriel"
            outer-class="formkit-outer-dark"
            label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
            input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent transition-colors"
          />
        </div>
      </div>

      <!-- Dossier Clinique -->
      <div class="mb-8">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4 pb-2 border-b border-gray-200 dark:border-gray-700 flex items-center gap-2">
          <FileText :size="20" class="text-purple-600 dark:text-purple-400" />
          Dossier Clinique
        </h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <FormKit 
            type="text" 
            name="no_hcm" 
            label="No HCM"
            outer-class="formkit-outer-dark"
            label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
            input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent transition-colors"
          />
          <FormKit 
            type="text" 
            name="no_chaur" 
            label="No CHAUR"
            outer-class="formkit-outer-dark"
            label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
            input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent transition-colors"
          />

          <!-- Numéro Leopard avec actions -->
          <div class="space-y-3">
            <FormKit
              type="text"
              name="no_dossier_leopard"
              label="No Dossier Léopard"
              :value="leopardNumber"
              readonly
              :help="leopardNumberHelp"
              outer-class="formkit-outer-dark"
              label-class="text-gray-700 dark:text-gray-300 font-medium text-sm"
              input-class="w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-lg bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-white font-mono cursor-not-allowed transition-colors"
              help-class="text-gray-500 dark:text-gray-400 text-xs mt-1"
            />

            <div class="flex flex-wrap gap-2">
              <!-- Bouton ouvrir dossier -->
              <button
                v-if="folderExists"
                @click="handleOpenFolder"
                type="button"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium bg-blue-50 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300 hover:bg-blue-100 dark:hover:bg-blue-900/50 rounded-lg transition-all duration-200 border border-blue-200 dark:border-blue-800 hover:shadow-md"
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
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium bg-green-50 dark:bg-green-900/30 text-green-700 dark:text-green-300 hover:bg-green-100 dark:hover:bg-green-900/50 rounded-lg transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed border border-green-200 dark:border-green-800 hover:shadow-md disabled:hover:shadow-none"
                title="Créer le dossier client sur le disque"
              >
                <FolderPlus :size="16" v-if="!isCreatingFolder" />
                <Loader2 :size="16" class="animate-spin" v-else />
                <span>{{
                  isCreatingFolder ? "Création..." : "Créer dossier"
                }}</span>
              </button>

              <!-- Bouton régénérer -->
              <button
                @click="handleRegenerateLeopardNumber"
                type="button"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium bg-gray-50 dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-all duration-200 border border-gray-200 dark:border-gray-700 hover:shadow-md"
                title="Régénérer le numéro Leopard"
              >
                <RefreshCw :size="16" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Informations sur le dossier (succès) -->
      <div
        v-if="folderInfo"
        class="mb-6 p-4 bg-blue-50 dark:bg-blue-900/20 rounded-xl border border-blue-200 dark:border-blue-800 transition-colors"
      >
        <div class="flex items-start gap-3">
          <div class="flex-shrink-0 p-2 bg-blue-100 dark:bg-blue-900/40 rounded-lg">
            <FolderCheck :size="20" class="text-blue-600 dark:text-blue-400" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-semibold text-blue-900 dark:text-blue-100">
              Dossier client actif
            </p>
            <p class="text-xs text-blue-700 dark:text-blue-300 mt-1 font-mono break-all">
              {{ folderInfo.path }}
            </p>
            <div class="flex flex-wrap gap-3 mt-3 text-xs">
              <span class="inline-flex items-center gap-1.5 px-2.5 py-1 bg-blue-100 dark:bg-blue-900/40 text-blue-700 dark:text-blue-300 rounded-md font-medium">
                <Folder :size="14" />
                {{ folderInfo.subfoldersCount }} sous-dossiers
              </span>
              <span class="inline-flex items-center gap-1.5 px-2.5 py-1 bg-blue-100 dark:bg-blue-900/40 text-blue-700 dark:text-blue-300 rounded-md font-medium">
                <FileText :size="14" />
                {{ folderInfo.filesCount }} fichiers
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Alerte si dossier n'existe pas -->
      <div
        v-else-if="formData.id && !folderExists"
        class="mb-6 p-4 bg-yellow-50 dark:bg-yellow-900/20 rounded-xl border border-yellow-200 dark:border-yellow-800 transition-colors"
      >
        <div class="flex items-start gap-3">
          <div class="flex-shrink-0 p-2 bg-yellow-100 dark:bg-yellow-900/40 rounded-lg">
            <AlertCircle :size="20" class="text-yellow-600 dark:text-yellow-400" />
          </div>
          <div class="flex-1">
            <p class="text-sm font-semibold text-yellow-900 dark:text-yellow-100">
              Dossier client non créé
            </p>
            <p class="text-xs text-yellow-700 dark:text-yellow-300 mt-1">
              Le dossier physique pour ce client n'existe pas encore. Cliquez
              sur "Créer dossier" pour le créer.
            </p>
          </div>
        </div>
      </div>
    </FormKit>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from "vue";
import { useDarkMode } from '@/composables/useDarkMode'
// ✅ Importe juste le composable, pas besoin de faire quoi que ce soit d'autre
const { isDark } = useDarkMode()

import {
  FolderOpen,
  FolderPlus,
  FolderCheck,
  RefreshCw,
  Loader2,
  AlertCircle,
  User,
  FileText,
  Folder,
} from "lucide-vue-next";
import {
  generateLeopardNumber,
  createClientFolder,
  openClientFolder,
  clientFolderExists,
} from "@/utils/clientFolderManager";

const props = defineProps({
  clientData: {
    type: Object,
    required: true,
  },
});

const emit = defineEmits(["save", "folderCreated"]);

const formData = ref({ ...props.clientData });
const leopardNumber = ref("");
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
    return "Remplissez le nom et prénom pour générer automatiquement";
  }
  return `Format: ${formData.value.nom
    .substring(0, 3)
    .toUpperCase()}${formData.value.prenom.substring(0, 1).toUpperCase()}_DATE`;
});

/**
 * Vérifie si le dossier existe
 */
const checkFolderExists = async () => {
  if (leopardNumber.value) {
    folderExists.value = await clientFolderExists(leopardNumber.value);

    if (folderExists.value) {
      await loadFolderInfo();
    }
  }
};

/**
 * Charge les informations du dossier
 */
const loadFolderInfo = async () => {
  try {
    const info = await window.go.main.App.GetClientFolderInfo(
      leopardNumber.value
    );
    if (info) {
      folderInfo.value = info;
    }
  } catch (error) {
    console.error("❌ Erreur chargement infos dossier:", error);
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

      emit("folderCreated", {
        leopardNumber: result.leopardNumber,
        path: result.path,
      });

      console.log("✅ Dossier créé avec succès");
      alert(
        `Dossier client créé avec succès!\n\nNuméro: ${result.leopardNumber}\nChemin: ${result.path}`
      );
    } else {
      throw new Error(result.error);
    }
  } catch (error) {
    console.error("❌ Erreur création dossier:", error);
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
      alert("Impossible d'ouvrir le dossier client");
    }
  } catch (error) {
    console.error("❌ Erreur ouverture dossier:", error);
    alert(`Erreur lors de l'ouverture du dossier:\n${error.message}`);
  }
};

/**
 * Régénère le numéro Leopard
 */
const handleRegenerateLeopardNumber = () => {
  if (
    confirm(
      "Êtes-vous sûr de vouloir régénérer le numéro Leopard?\nCela peut créer des incohérences si un dossier existe déjà."
    )
  ) {
    generateNumber();
    checkFolderExists();
  }
};

/**
 * Soumission du formulaire
 */
const submitHandler = async (data) => {
  data.no_dossier_leopard = leopardNumber.value;
  emit("save", data);
};

// Watchers
watch(
  () => props.clientData,
  (newVal) => {
    formData.value = { ...newVal };
    generateNumber();
    checkFolderExists();
  },
  { deep: true }
);

watch([() => formData.value.nom, () => formData.value.prenom], () => {
  generateNumber();
});

// Au montage
onMounted(() => {
  generateNumber();
  checkFolderExists();
});
</script>

<style scoped>
/* Animations */
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Override FormKit styles for dark mode */
:deep(.formkit-outer-dark) {
  margin-bottom: 1rem;
}

:deep(.formkit-wrapper) {
  max-width: 100%;
}
</style>