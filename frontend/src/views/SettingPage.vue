<template>
  <div class="flex h-screen w-full bg-slate-50 dark:bg-slate-950 overflow-hidden text-gray-900 dark:text-slate-100">
    
    <aside class="w-72 flex-shrink-0 border-r border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 p-6 flex flex-col">
      <h1 class="text-2xl font-bold mb-8">Paramètres</h1>
      <nav class="flex-1 space-y-1 overflow-y-auto">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            'w-full flex items-center gap-3 px-4 py-3 rounded-xl text-left transition-all',
            activeTab === tab.id
              ? 'bg-blue-600 text-white shadow-lg shadow-blue-900/20'
              : 'text-gray-600 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800'
          ]"
        >
          <component :is="tab.icon" :size="20" />
          <span class="font-medium">{{ tab.label }}</span>
        </button>
      </nav>
    </aside>

    <main class="flex-1 min-w-0 overflow-y-auto p-8">
      <div :class="activeTab === 'models' ? 'w-full' : 'max-w-5xl mx-auto space-y-6'">

        <RPAManager v-if="activeTab === 'rpa'" />
        <CHSLDSync v-if="activeTab === 'chsld'" />
        <MedecinImporter v-if="activeTab === 'medecins'" />
        <NotairesImporter v-if="activeTab === 'notaires'" />
        <MunImporter v-if="activeTab === 'geo'" />
        <RefListesSettings v-if="activeTab === 'listes'" :key="activeTab" />
        <Cim11Sync v-if="activeTab === 'cim11'" />
        
        <EvalCreator v-if="activeTab === 'models'" />

        <div v-if="activeTab === 'appearance'" class="bg-white dark:bg-gray-800 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-700 p-6 space-y-6">
          <h2 class="text-xl font-semibold mb-4">Apparence</h2>
          <div class="flex items-center justify-between py-4 border-b dark:border-gray-700">
            <div>
              <p class="font-medium">Mode sombre</p>
              <p class="text-sm text-gray-500 dark:text-gray-400">Activer le thème sombre pour l'interface</p>
            </div>
            <button
              @click="handleToggleDarkMode"
              :class="[
                'relative inline-flex h-7 w-12 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500',
                isDark ? 'bg-blue-600' : 'bg-gray-200 dark:bg-gray-600',
              ]"
            >
              <span :class="['inline-block h-5 w-5 transform rounded-full bg-white shadow-lg transition-transform', isDark ? 'translate-x-6' : 'translate-x-1']">
                <Moon v-if="isDark" :size="14" class="m-0.5 text-blue-600" />
                <Sun v-else :size="14" class="m-0.5 text-gray-400" />
              </span>
            </button>
          </div>
          <div class="py-4">
            <label class="block font-medium mb-2">Langue</label>
            <select v-model="settings.language" @change="hasChanges = true" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-blue-500">
              <option value="fr">Français</option>
              <option value="en">English</option>
            </select>
          </div>
        </div>

        <div v-if="activeTab === 'profile'" class="bg-white dark:bg-gray-800 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-700 p-6 space-y-4">
          <h2 class="text-xl font-semibold mb-4">Informations du profil</h2>
          <div>
            <label class="block text-sm font-medium mb-2">Nom complet *</label>
            <input v-model="profile.fullName" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 focus:ring-2 focus:ring-blue-500" />
          </div>
          <div class="flex justify-end pt-4">
            <button @click="saveProfile" :disabled="saving" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50">
              {{ saving ? "Enregistrement..." : "Enregistrer" }}
            </button>
          </div>
        </div>

        <div v-if="activeTab === 'security'" class="bg-white dark:bg-gray-800 rounded-2xl shadow-sm border border-slate-200 dark:border-slate-700 p-6 space-y-4">
          <h2 class="text-xl font-semibold mb-4">Sécurité</h2>
          <input v-model="passwordForm.current" type="password" placeholder="Mot de passe actuel" class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700" />
          <button @click="changePassword" class="px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700">Changer le mot de passe</button>
        </div>

      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import {
  User,
  Palette,
  Lock,
  Bell,
  Moon,
  Sun,
  Building2,
  Heart,
  Shield,
  Globe,
  MapPin,
  List,
  Stethoscope,
  Layout,
} from "lucide-vue-next";
import { useDarkMode } from "../composables/useDarkMode";
import {
  GetSettings,
  UpdateSettings,
  GetCurrentUserProfile,
  UpdateProfile,
  ChangePassword,
} from "../../wailsjs/go/main/App";

import RPAManager from "../components/RPA/RPAManager.vue";
import MedecinImporter from "../components/Medecins/MedecinImporter.vue"; //
import NotairesImporter from "../components/Notaires/NotairesImporter.vue"; //
import CHSLDSync from "../components/CHSLD/CHSLDSync.vue";
import MunImporter from "../components/settings/Munimporter.vue"; //
import RefListesSettings from "../components/settings/Reflistsetting.vue"; //
import EvalCreator from "../components/evaluation/EvaluationModelBuileder.vue"; //

import Cim11Sync from "../components/settings/Cim11Sync.vue";

// ✅ Utiliser le composable GLOBAL pour le dark mode
const { isDark, setTheme } = useDarkMode();

const activeTab = ref("appearance");
const saving = ref(false);
const savingPassword = ref(false);
const hasChanges = ref(false);
const passwordError = ref("");

const tabs = [
  { id: "appearance", label: "Apparence", icon: Palette },
  { id: "profile", label: "Profil", icon: User },
  { id: "security", label: "Sécurité", icon: Lock },
  { id: "models", label: "Modèles d'évaluation", icon: Layout },
  { id: "notifications", label: "Notifications", icon: Bell },
  { id: "rpa", label: "Gestion RPA", icon: Building2 },
  { id: "chsld", label: "Sync CHSLD", icon: Globe },
  { id: "medecins", label: "Importation de medecins", icon: Heart },
  { id: "notaires", label: "Importation de notaires", icon: Shield },
  { id: "geo", label: "Données géographiques", icon: MapPin },
  { id: "listes", label: "Listes de référence", icon: List },
  { id: "cim11", label: "Diagnostics CIM-11", icon: Stethoscope },
];

const settings = ref({
  theme: "light",
  language: "fr",
  notificationsEnabled: true,
  emailNotifications: true,
});

const settingsInitial = ref({});

const profile = ref({
  username: "",
  fullName: "",
  role: "",
});

const passwordForm = ref({
  current: "",
  new: "",
  confirm: "",
});

const getRoleLabel = (role) => {
  const labels = {
    admin: "Administrateur",
    psychologue: "Psychologue",
    travailleur_social: "Intervenant",
    lecteur: "Consultation",
  };
  return labels[role] || role;
};

// Charger les données
onMounted(async () => {
  try {
    const [settingsData, profileData] = await Promise.all([
      GetSettings(),
      GetCurrentUserProfile(),
    ]);

    settings.value = {
      theme: settingsData.theme || "light",
      language: settingsData.language || "fr",
      notificationsEnabled: settingsData.notifications_enabled,
      emailNotifications: settingsData.email_notifications,
    };

    settingsInitial.value = { ...settings.value };

    // ✅ APPLIQUER le thème chargé depuis la DB
    setTheme(settings.value.theme);

    profile.value = {
      username: profileData.username,
      fullName: profileData.fullName,
      role: profileData.role,
    };
  } catch (err) {
    console.error("Erreur chargement settings:", err);
  }
});

// Détecter les changements (sauf dark mode qui sauvegarde automatiquement)
watch(
  settings,
  () => {
    hasChanges.value =
      JSON.stringify(settings.value) !== JSON.stringify(settingsInitial.value);
  },
  { deep: true },
);

// ✅ TOGGLE Dark Mode avec sauvegarde immédiate
const handleToggleDarkMode = async () => {
  const newTheme = isDark.value ? "light" : "dark";

  // Mettre à jour l'état local
  settings.value.theme = newTheme;
  setTheme(newTheme);

  // Sauvegarder immédiatement dans la DB
  try {
    await UpdateSettings({
      theme: newTheme,
      language: settings.value.language,
      notifications_enabled: settings.value.notificationsEnabled,
      email_notifications: settings.value.emailNotifications,
    });

    // Mettre à jour la référence initiale
    settingsInitial.value.theme = newTheme;

    console.log("✅ Thème sauvegardé:", newTheme);
  } catch (err) {
    console.error("❌ Erreur sauvegarde thème:", err);
    alert("Erreur lors de la sauvegarde du thème");
  }
};

// Sauvegarder les settings (langue, notifications)
const saveSettings = async () => {
  saving.value = true;
  try {
    await UpdateSettings({
      theme: settings.value.theme,
      language: settings.value.language,
      notifications_enabled: settings.value.notificationsEnabled,
      email_notifications: settings.value.emailNotifications,
    });

    settingsInitial.value = { ...settings.value };
    hasChanges.value = false;
    alert("✅ Paramètres sauvegardés !");
  } catch (err) {
    console.error("Erreur sauvegarde:", err);
    alert("❌ Erreur lors de la sauvegarde");
  } finally {
    saving.value = false;
  }
};

// Sauvegarder le profil
const saveProfile = async () => {
  saving.value = true;
  try {
    await UpdateProfile({
      full_name: profile.value.fullName,
    });
    alert("✅ Profil mis à jour !");
  } catch (err) {
    console.error("Erreur sauvegarde profil:", err);
    alert("❌ Erreur lors de la sauvegarde");
  } finally {
    saving.value = false;
  }
};

// Changer le mot de passe
const changePassword = async () => {
  passwordError.value = "";

  if (passwordForm.value.new !== passwordForm.value.confirm) {
    passwordError.value = "Les mots de passe ne correspondent pas";
    return;
  }

  if (passwordForm.value.new.length < 8) {
    passwordError.value = "Le mot de passe doit contenir au moins 8 caractères";
    return;
  }

  savingPassword.value = true;
  try {
    await ChangePassword({
      current_password: passwordForm.value.current,
      new_password: passwordForm.value.new,
    });

    passwordForm.value = { current: "", new: "", confirm: "" };
    alert("✅ Mot de passe modifié avec succès !");
  } catch (err) {
    console.error("Erreur changement mot de passe:", err);
    passwordError.value =
      err.message || "Erreur lors du changement de mot de passe";
  } finally {
    savingPassword.value = false;
  }
};
</script>

<style scoped>
@keyframes slide-up {
  from {
    transform: translateY(100px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

.animate-slide-up {
  animation: slide-up 0.3s ease-out;
}
</style>
