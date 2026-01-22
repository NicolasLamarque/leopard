<template>
  <div class="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
    <div class="bg-white dark:bg-gray-900 rounded-2xl shadow-2xl w-full max-w-4xl max-h-[90vh] overflow-hidden flex flex-col animate-slideIn">
      
      <!-- Header -->
      <div class="bg-gradient-to-r from-slate-600 to-teal-700 dark:from-slate-600 dark:to-teal-700 px-8 py-6 flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-white/20 rounded-xl backdrop-blur-sm">
            <UserPlus :size="28" class="text-white" />
          </div>
          <div>
            <h2 class="text-2xl font-bold text-slate-100">Nouveau dossier client</h2>
            <p class="text-blue-100 text-sm mt-0.5">Création d'un nouveau dossier</p>
          </div>
        </div>
        <button 
          @click="$emit('close')" 
          class="p-2 hover:bg-white/20 rounded-lg transition-colors group"
          :disabled="loading"
        >
          <X :size="24" class="text-white group-hover:rotate-90 transition-transform" />
        </button>
      </div>

      <!-- Progress Bar -->
      <div class="h-1 bg-gray-200 dark:bg-gray-800">
        <div 
          class="h-full bg-gradient-to-r from-teal-600 to-teal-700 transition-all duration-500"
          :style="{ width: progressWidth }"
        ></div>
      </div>

      <!-- Form Content -->
      <div class="flex-1 overflow-y-auto px-8 py-6">
        <form @submit.prevent="save" class="space-y-8">
          
          <!-- Section 1: Identité -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 pb-3 border-b-2 border-blue-100 dark:border-teal-900">
              <div class="p-2 bg-blue-50 dark:bg-blue-900/30 rounded-lg">
                <User :size="20" class="text-slate-600 dark:text-slate-400" />
              </div>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Identité du client</h3>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-2">
                  Nom <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="form.nom" 
                  @input="updateProgress"
                  required 
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-slate-400 focus:ring-4 focus:ring-blue-500/20 transition-all outline-none"
                  placeholder="Nom de famille"
                />
              </div>

              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Prénom <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="form.prenom" 
                  @input="updateProgress"
                  required 
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-slate-400 focus:ring-4 focus:ring-blue-500/20 transition-all outline-none"
                  placeholder="Prénom"
                />
              </div>

              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-2">
                  <Calendar :size="16" />
                  Date de naissance
                </label>
                <input 
                  v-model="form.date_naissance" 
                  @input="updateProgress"
                  type="date" 
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-slate-400 focus:ring-4 focus:ring-blue-500/20 transition-all outline-none"
                />
              </div>

              <!-- Preview du numéro Leopard -->
              <div v-if="form.nom && form.prenom" class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-2">
                  <FileText :size="16" />
                  N° Dossier (auto-généré)
                </label>
                <div class="w-full px-4 py-3 bg-gradient-to-r from-green-50 to-emerald-50 dark:from-green-900/20 dark:to-emerald-900/20 border-2 border-green-200 dark:border-green-800 rounded-xl flex items-center gap-2">
                  <CheckCircle :size="18" class="text-green-600 dark:text-green-400" />
                  <span class="font-mono text-sm text-green-700 dark:text-green-300 font-semibold">
                    {{ previewLeopardNumber() }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Section 2: Coordonnées -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 pb-3 border-b-2 border-purple-100 dark:border-teal-900">
              <div class="p-2 bg-purple-50 dark:bg-purple-900/30 rounded-lg">
                <Phone :size="20" class="text-teal-600 dark:text-teal-400" />
              </div>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Coordonnées</h3>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-2">
                  <Phone :size="16" />
                  Téléphone
                </label>
                <input 
                  v-model="form.telephone" 
                  @input="updateProgress"
                  type="tel"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-slate-500 focus:ring-4 focus:ring-purple-500/20 transition-all outline-none"
                  placeholder="819-555-0123"
                />
              </div>

              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-2">
                  <Smartphone :size="16" />
                  Cellulaire
                </label>
                <input 
                  v-model="form.cellulaire" 
                  @input="updateProgress"
                  type="tel"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-slate-500 focus:ring-4 focus:ring-purple-500/20 transition-all outline-none"
                  placeholder="819-555-0456"
                />
              </div>

              <div class="group md:col-span-2">
  <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-2">
    <Mail :size="16" />
    Courriel
  </label>
  <input 
    v-model="form.email" 
    @input="updateProgress"
    type="email"
    class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-slate-500 focus:ring-4 focus:ring-purple-500/20 transition-all outline-none"
    placeholder="exemple@email.com"
  />
</div>
            </div>
          </div>

          <!-- Section 3: Adresse -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 pb-3 border-b-2 border-orange-100 dark:border-teal-900">
              <div class="p-2 bg-orange-50 dark:bg-orange-900/30 rounded-lg">
                <MapPin :size="20" class="text-orange-600 dark:text-orange-400" />
              </div>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Adresse</h3>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="group md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Adresse complète
                </label>
                <textarea 
                  v-model="form.adresse" 
                  @input="updateProgress"
                  rows="2"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-orange-500 focus:ring-4 focus:ring-orange-500/20 transition-all outline-none resize-none"
                  placeholder="123 Rue Principale"
                ></textarea>
              </div>

              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Ville
                </label>
                <input 
                  v-model="form.ville" 
                  @input="updateProgress"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-orange-500 focus:ring-4 focus:ring-orange-500/20 transition-all outline-none"
                  placeholder="Gatineau"
                />
              </div>

              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Code postal
                </label>
                <input 
                  v-model="form.code_postal" 
                  @input="updateProgress"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-orange-500 focus:ring-4 focus:ring-orange-500/20 transition-all outline-none"
                  placeholder="J8V 1A1"
                />
              </div>
            </div>
          </div>

          <!-- Section 4: Informations médicales -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 pb-3 border-b-2 border-red-100 dark:border-teal-900">
              <div class="p-2 bg-red-50 dark:bg-red-900/30 rounded-lg">
                <Heart :size="20" class="text-red-600 dark:text-red-400" />
              </div>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Informations médicales</h3>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° Assurance maladie (RAMQ)
                </label>
                <input 
                  v-model="form.numero_assurance_maladie" 
                  @input="updateProgress"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-red-500 focus:ring-4 focus:ring-red-500/20 transition-all outline-none font-mono"
                  placeholder="TREJ 4503 1501"
                />
              </div>

              <div class="group">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° Sécurité sociale
                </label>
                <input 
                  v-model="form.numero_securite_sociale" 
                  @input="updateProgress"
                  class="w-full px-4 py-3 border-2 border-gray-300 dark:border-gray-700 rounded-xl bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-red-500 focus:ring-4 focus:ring-red-500/20 transition-all outline-none font-mono"
                  placeholder="123-456-789"
                />
              </div>
            </div>
          </div>

        </form>
      </div>

      <!-- Footer avec actions -->
      <div class="border-t-2 border-gray-200 dark:border-gray-800 px-8 py-5 bg-gray-50 dark:bg-gray-900/50">
        <div class="flex items-center justify-between">
          <!-- Indicateur de progression -->
          <div class="flex items-center gap-3">
            <div class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
              <div class="w-2 h-2 rounded-full animate-pulse" :class="progressColor"></div>
              <span>{{ progressText }}</span>
            </div>
          </div>

          <!-- Boutons -->
          <div class="flex gap-3">
            <button 
              type="button" 
              @click="$emit('close')" 
              :disabled="loading"
              class="px-6 py-3 border-2 border-gray-300 dark:border-gray-700 text-gray-700 dark:text-gray-300 rounded-xl hover:bg-gray-100 dark:hover:bg-gray-800 transition-all font-medium disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Annuler
            </button>
            <button 
              @click="save"
              :disabled="loading || !canSubmit"
              class="px-8 py-3 bg-gradient-to-r from-teal-300 to-slate-600 hover:from-teal-700 hover:to-slate-800 text-white rounded-xl font-semibold shadow-lg hover:shadow-xl transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            >
              <Loader2 v-if="loading" :size="20" class="animate-spin" />
              <Save v-else :size="20" />
              <span>{{ loading ? 'Création en cours...' : 'Créer le dossier' }}</span>
            </button>
          </div>
        </div>

        <!-- Message d'erreur -->
        <div v-if="error" class="mt-4 p-4 bg-red-50 dark:bg-red-900/20 border-2 border-red-200 dark:border-red-800 rounded-xl flex items-start gap-3">
          <AlertCircle :size="20" class="text-red-600 dark:text-red-400 flex-shrink-0 mt-0.5" />
          <div>
            <p class="text-sm font-medium text-red-900 dark:text-red-100">Erreur lors de la création</p>
            <p class="text-sm text-red-700 dark:text-red-300 mt-1">{{ error }}</p>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { 
  UserPlus, X, User, Phone, Mail, MapPin, Heart, FileText, Calendar, 
  CheckCircle, Loader2, Save, AlertCircle, Smartphone 
} from 'lucide-vue-next';
import { CreateClient } from '../../../wailsjs/go/main/App';

const emit = defineEmits(['close', 'created']);

const loading = ref(false);
const error = ref('');

const form = ref({
  // --- Identité (Champs existants) ---
  nom: '',
  prenom: '',
  date_naissance: null,
  sexe: null,
  lieu_naissance: null,
  statut_marital: null,
  occupation: null,
  employeur: null,
  profession: null,
  niveau_scolaire: null,
  langue_preferee: 'Français',
  origine_ethnique: null,
  identite_genre: null,
  orientation_sexuelle: null,
  religion: null,
  premiere_langue: null,
  premiere_nationale: 0, 
  telephone: null,
  cellulaire: null,
  email: null,
  adresse: null,
  appartement: null, // Ajouté
  code_postal: null,
  ville: null,
  pays: 'Canada',
  province: 'Québec',
  numero_assurance_maladie: null,
  numero_securite_sociale: null,
  no_hcm: null,
  no_chaur: null,
  no_dossier_leopard: null, // Sera généré par le backend
  medecin_famille_No_Licence: null,

  // --- Références (Relations) ---
  notaire_id: null,
  pivot_id: null,
  rpa_id: null,
  chsld_id: null,
  ri_id: null,

  // --- États et Notes ---
  note_fixe: null,
  actif: 1, // Note le 'a' minuscule pour matcher le JSON Go
  dcd: 0,
  date_deces: null
});

// Calcul de la progression
const filledFields = computed(() => {
  let count = 0;
  if (form.value.nom) count++;
  if (form.value.prenom) count++;
  if (form.value.date_naissance) count++;
  if (form.value.telephone) count++;
  if (form.value.cellulaire) count++;
  if (form.value.email) count++;
  if (form.value.adresse) count++;
  if (form.value.ville) count++;
  if (form.value.code_postal) count++;
  if (form.value.numero_assurance_maladie) count++;
  if (form.value.numero_securite_sociale) count++;
  return count;
});

const totalFields = 11;

const progressWidth = computed(() => {
  const percentage = (filledFields.value / totalFields) * 100;
  return `${Math.min(percentage, 100)}%`;
});

const progressText = computed(() => {
  const percentage = Math.round((filledFields.value / totalFields) * 100);
  return `${percentage}% complété (${filledFields.value}/${totalFields} champs)`;
});

const progressColor = computed(() => {
  const percentage = (filledFields.value / totalFields) * 100;
  if (percentage < 30) return 'bg-red-500';
  if (percentage < 70) return 'bg-yellow-500';
  return 'bg-green-500';
});

const canSubmit = computed(() => {
  return form.value.nom && form.value.prenom;
});

const updateProgress = () => {
  // Force le recalcul de la progression
};

// Preview du numéro Leopard
const previewLeopardNumber = () => {
  if (!form.value.nom || !form.value.prenom) return '';
  
  // Nettoyage basique pour preview (le vrai sera généré côté serveur)
  const cleanNom = form.value.nom.normalize('NFD').replace(/[\u0300-\u036f]/g, '').toUpperCase();
  const cleanPrenom = form.value.prenom.normalize('NFD').replace(/[\u0300-\u036f]/g, '').toUpperCase();
  
  const nomPart = cleanNom.substring(0, 3).padEnd(3, 'X');
  const prenomPart = cleanPrenom.substring(0, 1) || 'X';
  
  const today = new Date();
  const year = today.getFullYear();
  const month = String(today.getMonth() + 1).padStart(2, '0');
  const day = String(today.getDate()).padStart(2, '0');
  
  // Preview approximatif - le compteur réel sera généré par le serveur
  return `${nomPart}${prenomPart}_${year}${month}${day}_XXX`;
};

const save = async () => {
  loading.value = true;
  error.value = '';
  
  try {
    console.log('📝 Création client avec données:', form.value);
    
    const clientId = await CreateClient(form.value);
    
    console.log('✅ Client créé avec ID:', clientId);
    
    emit('created');
    emit('close');
  } catch (e) {
    console.error('❌ Erreur création client:', e);
    error.value = e.message || 'Erreur lors de la création du client' ;
    // amelioration du debug puissant
    console.error('Données du formulaire au moment de l\'erreur:', form.value);
    

  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
@keyframes slideIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

.animate-slideIn {
  animation: slideIn 0.3s ease-out;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>