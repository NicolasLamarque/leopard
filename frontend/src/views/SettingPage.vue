<template>
  <div class="p-6 max-w-5xl mx-auto">
    <h1 class="text-3xl font-bold mb-8 text-gray-900 dark:text-white">Paramètres</h1>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Menu latéral -->
      <div class="lg:col-span-1">
        <nav class="bg-white dark:bg-gray-800 rounded-lg shadow p-4 space-y-2 sticky top-6">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'w-full flex items-center gap-3 px-4 py-3 rounded-lg text-left transition-all',
              activeTab === tab.id
                ? 'bg-blue-50 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400 font-medium'
                : 'text-gray-600 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700'
            ]"
          >
            <component :is="tab.icon" :size="20" />
            <span>{{ tab.label }}</span>
          </button>
        </nav>
      </div>

      <!-- Contenu -->
      <div class="lg:col-span-2 space-y-6">

        <!-- RPA Manager -->
  <div v-if="activeTab === 'rpa'">
    <RPAManager />
  </div>
  <!-- 🆕 Import Médecins -->
        <div v-if="activeTab === 'medecins'">
          <MedecinImporter />
        </div>

        <!-- 🆕 Import Notaires -->
        <div v-if="activeTab === 'notaires'">
          <NotairesImporter />
        </div>
        
        <!-- Onglet Apparence -->
        <div v-if="activeTab === 'appearance'" class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6">
          <div>
            <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Apparence</h2>
            
            <!-- Dark Mode -->
            <div class="flex items-center justify-between py-4 border-b dark:border-gray-700">
              <div>
                <p class="font-medium text-gray-900 dark:text-white">Mode sombre</p>
                <p class="text-sm text-gray-500 dark:text-gray-400">Activer le thème sombre pour l'interface</p>
              </div>


              <button 
                @click="handleToggleDarkMode"
                :class="[
                  'relative inline-flex h-7 w-12 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2',
                  isDark ? 'bg-blue-600' : 'bg-gray-200 dark:bg-gray-600'
                ]"
              >
                <span 
                  :class="[
                    'inline-block h-5 w-5 transform rounded-full bg-white shadow-lg transition-transform',
                    isDark ? 'translate-x-6' : 'translate-x-1'
                  ]"
                >
                  <Moon v-if="isDark" :size="14" class="m-0.5 text-blue-600" />
                  <Sun v-else :size="14" class="m-0.5 text-gray-400" />
                </span>
              </button>
            </div>

            <!-- Langue -->
            <div class="py-4">
              <label class="block font-medium mb-2 text-gray-900 dark:text-white">Langue</label>
              <select 
                v-model="settings.language"
                @change="hasChanges = true"
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              >
                <option value="fr">Français</option>
                <option value="en">English</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Onglet Profil -->
        <div v-if="activeTab === 'profile'" class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6">
          <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Informations du profil</h2>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium mb-2 text-gray-900 dark:text-white">Nom complet *</label>
              <input 
                v-model="profile.fullName"
                required
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium mb-2 text-gray-900 dark:text-white">Nom d'utilisateur</label>
              <input 
                :value="profile.username"
                disabled
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-gray-100 dark:bg-gray-900 text-gray-500 dark:text-gray-400"
              />
              <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">Le nom d'utilisateur ne peut pas être modifié</p>
            </div>

            <div>
              <label class="block text-sm font-medium mb-2 text-gray-900 dark:text-white">Rôle</label>
              <input 
                :value="getRoleLabel(profile.role)"
                disabled
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-gray-100 dark:bg-gray-900 text-gray-500 dark:text-gray-400"
              />
            </div>
          </div>

          <div class="flex justify-end pt-4">
            <button 
              @click="saveProfile"
              :disabled="saving"
              class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {{ saving ? 'Enregistrement...' : 'Enregistrer' }}
            </button>
          </div>
        </div>

        <!-- Onglet Sécurité -->
        <div v-if="activeTab === 'security'" class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6">
          <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Sécurité</h2>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium mb-2 text-gray-900 dark:text-white">Mot de passe actuel *</label>
              <input 
                v-model="passwordForm.current"
                type="password"
                required
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium mb-2 text-gray-900 dark:text-white">Nouveau mot de passe *</label>
              <input 
                v-model="passwordForm.new"
                type="password"
                required
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium mb-2 text-gray-900 dark:text-white">Confirmer le nouveau mot de passe *</label>
              <input 
                v-model="passwordForm.confirm"
                type="password"
                required
                class="w-full border dark:border-gray-600 rounded-lg px-4 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <p v-if="passwordError" class="text-sm text-red-600 dark:text-red-400">{{ passwordError }}</p>
          </div>

          <div class="flex justify-end pt-4">
            <button 
              @click="changePassword"
              :disabled="savingPassword"
              class="px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {{ savingPassword ? 'Modification...' : 'Changer le mot de passe' }}
            </button>
          </div>
        </div>

        <!-- Onglet Notifications -->
        <div v-if="activeTab === 'notifications'" class="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-6">
          <h2 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white">Notifications</h2>
          
          <div class="space-y-4">
            <label class="flex items-center justify-between py-3 border-b dark:border-gray-700">
              <div>
                <p class="font-medium text-gray-900 dark:text-white">Notifications dans l'application</p>
                <p class="text-sm text-gray-500 dark:text-gray-400">Recevoir des alertes dans l'interface</p>
              </div>
              <input 
                v-model="settings.notificationsEnabled"
                @change="hasChanges = true"
                type="checkbox" 
                class="w-5 h-5 text-blue-600 rounded focus:ring-2 focus:ring-blue-500"
              />
            </label>

            <label class="flex items-center justify-between py-3">
              <div>
                <p class="font-medium text-gray-900 dark:text-white">Notifications par email</p>
                <p class="text-sm text-gray-500 dark:text-gray-400">Recevoir des emails pour les événements importants</p>
              </div>
              <input 
                v-model="settings.emailNotifications"
                @change="hasChanges = true"
                type="checkbox" 
                class="w-5 h-5 text-blue-600 rounded focus:ring-2 focus:ring-blue-500"
              />
            </label>
          </div>
        </div>

      </div>
    </div>

    <!-- Bouton de sauvegarde flottant -->
    <div 
      v-if="hasChanges" 
      class="fixed bottom-6 right-6 flex items-center gap-3 bg-white dark:bg-gray-800 shadow-2xl rounded-lg p-4 border border-gray-200 dark:border-gray-700 animate-slide-up"
    >
      <p class="text-sm font-medium text-gray-700 dark:text-gray-300">Modifications non sauvegardées</p>
      <button 
        @click="saveSettings"
        :disabled="saving"
        class="px-5 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 transition-all shadow-lg"
      >
        {{ saving ? 'Enregistrement...' : 'Enregistrer' }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { User, Palette, Lock, Bell, Moon, Sun, Building2, Heart, Shield } from 'lucide-vue-next'
import { useDarkMode } from '../composables/useDarkMode'
import { 
  GetSettings, 
  UpdateSettings, 
  GetCurrentUserProfile, 
  UpdateProfile,
  ChangePassword 
} from '../../wailsjs/go/main/App'

import RPAManager from '../components/RPA/RPAManager.vue'
import MedecinImporter from '../components/Medecins/MedecinImporter.vue' //
import NotairesImporter from '../components/Notaires/NotairesImporter.vue' //



// ✅ Utiliser le composable GLOBAL pour le dark mode
const { isDark, setTheme } = useDarkMode()

const activeTab = ref('appearance')
const saving = ref(false)
const savingPassword = ref(false)
const hasChanges = ref(false)
const passwordError = ref('')

const tabs = [
  { id: 'appearance', label: 'Apparence', icon: Palette },
  { id: 'profile', label: 'Profil', icon: User },
  { id: 'security', label: 'Sécurité', icon: Lock },
  { id: 'notifications', label: 'Notifications', icon: Bell },
  { id: 'rpa', label: 'Gestion RPA', icon: Building2 }, // ← NOUVEAU
  { id: 'medecins', label: 'Importation de medecins', icon: Heart },
  { id: 'notaires', label: 'Importation de notaires', icon: Shield }
]

const settings = ref({
  theme: 'light',
  language: 'fr',
  notificationsEnabled: true,
  emailNotifications: true
})

const settingsInitial = ref({})

const profile = ref({
  username: '',
  fullName: '',
  role: ''
})

const passwordForm = ref({
  current: '',
  new: '',
  confirm: ''
})

const getRoleLabel = (role) => {
  const labels = {
    admin: 'Administrateur',
    psychologue: 'Psychologue',
    travailleur_social: 'Intervenant',
    lecteur: 'Consultation'
  }
  return labels[role] || role
}

// Charger les données
onMounted(async () => {
  try {
    const [settingsData, profileData] = await Promise.all([
      GetSettings(),
      GetCurrentUserProfile()
    ])
    
    settings.value = {
      theme: settingsData.theme || 'light',
      language: settingsData.language || 'fr',
      notificationsEnabled: settingsData.notifications_enabled,
      emailNotifications: settingsData.email_notifications
    }
    
    settingsInitial.value = { ...settings.value }
    
    // ✅ APPLIQUER le thème chargé depuis la DB
    setTheme(settings.value.theme)
    
    profile.value = {
      username: profileData.username,
      fullName: profileData.fullName,
      role: profileData.role
    }
  } catch (err) {
    console.error('Erreur chargement settings:', err)
  }
})

// Détecter les changements (sauf dark mode qui sauvegarde automatiquement)
watch(settings, () => {
  hasChanges.value = JSON.stringify(settings.value) !== JSON.stringify(settingsInitial.value)
}, { deep: true })

// ✅ TOGGLE Dark Mode avec sauvegarde immédiate
const handleToggleDarkMode = async () => {
  const newTheme = isDark.value ? 'light' : 'dark'
  
  // Mettre à jour l'état local
  settings.value.theme = newTheme
  setTheme(newTheme)
  
  // Sauvegarder immédiatement dans la DB
  try {
    await UpdateSettings({
      theme: newTheme,
      language: settings.value.language,
      notifications_enabled: settings.value.notificationsEnabled,
      email_notifications: settings.value.emailNotifications
    })
    
    // Mettre à jour la référence initiale
    settingsInitial.value.theme = newTheme
    
    console.log('✅ Thème sauvegardé:', newTheme)
  } catch (err) {
    console.error('❌ Erreur sauvegarde thème:', err)
    alert('Erreur lors de la sauvegarde du thème')
  }
}

// Sauvegarder les settings (langue, notifications)
const saveSettings = async () => {
  saving.value = true
  try {
    await UpdateSettings({
      theme: settings.value.theme,
      language: settings.value.language,
      notifications_enabled: settings.value.notificationsEnabled,
      email_notifications: settings.value.emailNotifications
    })
    
    settingsInitial.value = { ...settings.value }
    hasChanges.value = false
    alert('✅ Paramètres sauvegardés !')
  } catch (err) {
    console.error('Erreur sauvegarde:', err)
    alert('❌ Erreur lors de la sauvegarde')
  } finally {
    saving.value = false
  }
}

// Sauvegarder le profil
const saveProfile = async () => {
  saving.value = true
  try {
    await UpdateProfile({
      full_name: profile.value.fullName
    })
    alert('✅ Profil mis à jour !')
  } catch (err) {
    console.error('Erreur sauvegarde profil:', err)
    alert('❌ Erreur lors de la sauvegarde')
  } finally {
    saving.value = false
  }
}

// Changer le mot de passe
const changePassword = async () => {
  passwordError.value = ''
  
  if (passwordForm.value.new !== passwordForm.value.confirm) {
    passwordError.value = 'Les mots de passe ne correspondent pas'
    return
  }
  
  if (passwordForm.value.new.length < 8) {
    passwordError.value = 'Le mot de passe doit contenir au moins 8 caractères'
    return
  }
  
  savingPassword.value = true
  try {
    await ChangePassword({
      current_password: passwordForm.value.current,
      new_password: passwordForm.value.new
    })
    
    passwordForm.value = { current: '', new: '', confirm: '' }
    alert('✅ Mot de passe modifié avec succès !')
  } catch (err) {
    console.error('Erreur changement mot de passe:', err)
    passwordError.value = err.message || 'Erreur lors du changement de mot de passe'
  } finally {
    savingPassword.value = false
  }
}
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