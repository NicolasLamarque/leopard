<template>
  
<div class="bg-white dark:bg-gray-900 p-6 rounded-xl shadow-lg dark:shadow-gray-950/30 border border-gray-100 dark:border-gray-800 transition-colors text-gray-900 dark:text-gray-100">

         

    <FormKit
      type="form"
      v-model="formData"
      submit-label="Sauvegarder les modifications"
      @submit="submitHandler"
      :submit-attrs="{
        inputClass: 'bg-blue-600 hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 text-white font-medium px-6 py-2.5 rounded-lg transition-colors cursor-pointer'
      }"
      :config="{
        classes: {
          label: 'block mb-1 font-semibold text-sm text-gray-700 dark:text-gray-200',
          input: 'w-full px-3 py-2 border rounded-lg dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:placeholder-gray-500 focus:ring-2 focus:ring-blue-500 outline-none',
          help: 'text-xs text-gray-500 dark:text-gray-400 mt-1',
          message: 'text-red-500 text-xs mt-1',
          inner: 'mb-4'
        }
      }"
    >
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
          />
          <FormKit 
            type="text" 
            name="nom" 
            label="Nom" 
            validation="required"
          />
          <FormKit
            type="date"
            name="date_naissance"
            label="Date de naissance"
          />
        </div>
      </div>

      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <Phone :size="20" class="text-green-600 dark:text-green-400" />
          Coordonnées
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <FormKit
            type="tel"
            name="telephone"
            label="Téléphone"
            placeholder="819-555-0123"
          />
          <FormKit
            type="tel"
            name="cellulaire"
            label="Cellulaire"
            placeholder="819-555-0456"
          />
          <FormKit 
            type="email" 
            name="email" 
            label="Courriel"
            placeholder="exemple@email.com"
            outer-class="md:col-span-2"
          />
        </div>
      </div>

      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <MapPin :size="20" class="text-red-600 dark:text-red-400" />
          Adresse
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <FormKit
            type="text"
            name="adresse"
            label="Adresse complète"
            placeholder="123 Rue Principale"
            outer-class="md:col-span-2"
          />
          <FormKit
            type="text"
            name="ville"
            label="Ville"
            placeholder="Gatineau"
          />
          <FormKit
            type="text"
            name="code_postal"
            label="Code postal"
            placeholder="J8V 1A1"
          />
          <FormKit
            type="text"
            name="pays"
            label="Pays"
            placeholder="Canada"
          />
        </div>
      </div>

      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <Heart :size="20" class="text-pink-600 dark:text-pink-400" />
          Informations médicales
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <FormKit
            type="text"
            name="numero_assurance_maladie"
            label="N° Assurance maladie (RAMQ)"
            placeholder="TREJ 4503 1501"
          />
          <FormKit
            type="text"
            name="numero_securite_sociale"
            label="N° Sécurité sociale"
            placeholder="123-456-789"
          />
          <FormKit
            type="text"
            name="medecin_famille_No_Licence"
            label="N° Licence Médecin de famille"
            placeholder="12345"
          />
        </div>
      </div>

      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <FileText :size="20" class="text-purple-600 dark:text-purple-400" />
          Dossier Clinique
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <FormKit 
            type="text" 
            name="no_hcm" 
            label="N° HCM"
            placeholder="HCM001"
          />
          <FormKit 
            type="text" 
            name="no_chaur" 
            label="N° CHAUR"
            placeholder="CH001"
          />

          <div class="space-y-3">
            <FormKit
              type="text"
              name="no_dossier_leopard"
              label="N° Dossier Léopard"
              :value="leopardNumber"
              readonly
              :help="leopardNumberHelp"
            />

            <div class="flex flex-wrap gap-2">
              <button
                v-if="folderExists"
                @click="handleOpenFolder"
                type="button"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium bg-blue-50 dark:bg-blue-900/40 text-blue-700 dark:text-blue-200 hover:bg-blue-100 dark:hover:bg-blue-800 rounded-lg transition-all border border-blue-200 dark:border-blue-700"
              >
                <FolderOpen :size="16" />
                <span>Ouvrir</span>
              </button>

              <button
                v-else
                @click="handleCreateFolder"
                type="button"
                :disabled="isCreatingFolder"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium bg-green-50 dark:bg-green-900/40 text-green-700 dark:text-green-200 hover:bg-green-100 dark:hover:bg-green-800 rounded-lg transition-all disabled:opacity-50 border border-green-200 dark:border-green-700"
              >
                <FolderPlus :size="16" v-if="!isCreatingFolder" />
                <Loader2 :size="16" class="animate-spin" v-else />
                <span>{{ isCreatingFolder ? "Création..." : "Créer dossier" }}</span>
              </button>

              <button
                @click="handleRegenerateLeopardNumber"
                type="button"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium bg-gray-50 dark:bg-gray-800 text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-all border border-gray-200 dark:border-gray-700"
              >
                <RefreshCw :size="16" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <Building :size="20" class="text-orange-600 dark:text-orange-400" />
          Établissements & Contacts
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <FormKit type="number" name="notaire_id" label="ID Notaire" placeholder="Ex: 1" />
          <FormKit type="number" name="pivot_id" label="ID PIVOT" placeholder="Ex: 1" />
          <FormKit type="number" name="rpa_id" label="ID RPA" placeholder="Ex: 1" />
          <FormKit type="number" name="chsld_id" label="ID CHSLD" placeholder="Ex: 1" />
          <FormKit type="number" name="ri_id" label="ID RI" placeholder="Ex: 1" outer-class="md:col-span-2" />
        </div>
      </div>

      <div class="mb-8">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
          <ClipboardList :size="20" class="text-indigo-600 dark:text-indigo-400" />
          Notes & Statut
        </h2>
        <div class="space-y-4">
          <FormKit
            type="textarea"
            name="note_fixe"
            label="Note fixe"
            placeholder="Notes importantes concernant ce client..."
            :rows="4"
          />
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <FormKit type="checkbox" name="Actif" label="Client actif" :value="1" />
            <FormKit type="checkbox" name="dcd" label="Décédé" :value="1" />
          </div>
        </div>
      </div>

      <div v-if="folderInfo" class="mb-6 p-4 bg-blue-50 dark:bg-blue-900/40 rounded-xl border border-blue-200 dark:border-blue-700 transition-colors text-blue-900 dark:text-blue-100">
        <div class="flex items-start gap-3">
          <div class="flex-shrink-0 p-2 bg-blue-100 dark:bg-blue-800 rounded-lg">
            <FolderCheck :size="20" class="text-blue-600 dark:text-blue-300" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-semibold">Dossier client actif</p>
            <p class="text-xs text-blue-700 dark:text-blue-200 mt-1 font-mono break-all opacity-90">{{ folderInfo.path }}</p>
            <div class="flex flex-wrap gap-3 mt-3 text-xs">
              <span class="inline-flex items-center gap-1.5 px-2.5 py-1 bg-blue-100 dark:bg-blue-800 text-blue-700 dark:text-blue-100 rounded-md">
                <Folder :size="14" /> {{ folderInfo.subfoldersCount }} sous-dossiers
              </span>
              <span class="inline-flex items-center gap-1.5 px-2.5 py-1 bg-blue-100 dark:bg-blue-800 text-blue-700 dark:text-blue-100 rounded-md">
                <FileText :size="14" /> {{ folderInfo.filesCount }} fichiers
              </span>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="formData.id && !folderExists" class="mb-6 p-4 bg-yellow-50 dark:bg-yellow-900/40 rounded-xl border border-yellow-200 dark:border-yellow-700 transition-colors text-yellow-900 dark:text-yellow-100">
        <div class="flex items-start gap-3">
          <div class="flex-shrink-0 p-2 bg-yellow-100 dark:bg-yellow-800 rounded-lg">
            <AlertCircle :size="20" class="text-yellow-600 dark:text-yellow-300" />
          </div>
          <div class="flex-1">
            <p class="text-sm font-semibold">Dossier client non créé</p>
            <p class="text-xs text-yellow-700 dark:text-yellow-200 mt-1 opacity-90">Le dossier physique n'existe pas encore. Cliquez sur "Créer dossier".</p>
          </div>
        </div>
      </div>
    </FormKit>
  </div>
  
</template>

<script setup>
import { ref, computed, watch, onMounted } from "vue";
import { useDarkMode } from '@/composables/useDarkMode'
const { isDark } = useDarkMode()

import {
  FolderOpen, FolderPlus, FolderCheck, RefreshCw, Loader2,
  AlertCircle, User, FileText, Folder, Phone, MapPin, Heart, Building, ClipboardList,
} from "lucide-vue-next";
import {
  generateLeopardNumber, createClientFolder, openClientFolder, clientFolderExists,
} from "@/utils/clientFolderManager";

const props = defineProps({
  clientData: { type: Object, required: true },
});

const emit = defineEmits(["save", "folderCreated"]);

const formData = ref({ ...props.clientData });
const leopardNumber = ref("");
const folderExists = ref(false);
const folderInfo = ref(null);
const isCreatingFolder = ref(false);

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

const leopardNumberHelp = computed(() => {
  if (!formData.value.nom || !formData.value.prenom) {
    return "Remplissez le nom et prénom pour générer automatiquement";
  }
  return `Format: ${formData.value.nom.substring(0, 3).toUpperCase()}${formData.value.prenom.substring(0, 1).toUpperCase()}_DATE`;
});

const checkFolderExists = async () => {
  if (leopardNumber.value) {
    folderExists.value = await clientFolderExists(leopardNumber.value);
    if (folderExists.value) {
      await loadFolderInfo();
    }
  }
};

const loadFolderInfo = async () => {
  try {
    const info = await window.go.main.App.GetClientFolderInfo(leopardNumber.value);
    if (info) folderInfo.value = info;
  } catch (error) {
    console.error("❌ Erreur chargement infos dossier:", error);
  }
};

const handleCreateFolder = async () => {
  try {
    isCreatingFolder.value = true;
    const result = await createClientFolder(formData.value);
    if (result.success) {
      folderExists.value = true;
      leopardNumber.value = result.leopardNumber;
      formData.value.no_dossier_leopard = result.leopardNumber;
      await loadFolderInfo();
      emit("folderCreated", { leopardNumber: result.leopardNumber, path: result.path });
    } else {
      throw new Error(result.error);
    }
  } catch (error) {
    alert(`Erreur: ${error.message}`);
  } finally {
    isCreatingFolder.value = false;
  }
};

const handleOpenFolder = async () => {
  try {
    const clientName = `${formData.value.nom} ${formData.value.prenom}`;
    await openClientFolder(leopardNumber.value, clientName);
  } catch (error) {
    alert(`Erreur: ${error.message}`);
  }
};

const handleRegenerateLeopardNumber = () => {
  if (confirm("Régénérer le numéro Leopard ?")) {
    generateNumber();
    checkFolderExists();
  }
};

const submitHandler = async (data) => {
  data.no_dossier_leopard = leopardNumber.value;
  emit("save", data);
};

watch(() => props.clientData, (newVal) => {
  formData.value = { ...newVal };
  generateNumber();
  checkFolderExists();
}, { deep: true });

watch([() => formData.value.nom, () => formData.value.prenom], () => generateNumber());

onMounted(() => {
  generateNumber();
  checkFolderExists();
});
</script>

<style scoped>
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }
.animate-spin { animation: spin 1s linear infinite; }
</style>