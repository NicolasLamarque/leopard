<template>
  <div class="flex h-[calc(100vh-120px)] gap-5 p-6 animate-in fade-in duration-300">
    <!-- SIDEBAR - Liste des notaires -->
    <div class="w-1/3 flex flex-col gap-3 bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden">
      
      <!-- Header avec recherche -->
      <div class="p-5 border-b border-slate-200 dark:border-slate-800 bg-slate-50/80 dark:bg-slate-900/50">
        <div class="relative">
          <Search :size="16" class="absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Rechercher un notaire..."
            class="w-full pl-10 pr-4 py-2.5 rounded-xl border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 placeholder-slate-400 focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all"
          />
        </div>
        
        <!-- Bouton nouveau notaire -->
        <button 
          @click="toggleQuickForm"
          class="mt-3 w-full flex items-center justify-center gap-2 px-4 py-2.5 bg-teal-600 hover:bg-teal-700 text-white rounded-xl transition-all shadow-sm hover:shadow-md active:scale-[0.98]"
        >
          <Plus :size="18" />
          <span class="text-sm font-semibold">{{ showQuickForm ? 'Annuler' : 'Nouveau Notaire' }}</span>
        </button>
      </div>

      <!-- Liste des notaires -->
      <div class="flex-1 overflow-y-auto p-3 space-y-2">
        <div
          v-for="n in filteredNotaires"
          :key="n.id"
          @click="selectNotaire(n)"
          :class="[
            'flex items-center justify-between p-3.5 rounded-xl cursor-pointer transition-all',
            selectedNotaire?.id === n.id
              ? 'bg-teal-50 dark:bg-teal-900/20 border-2 border-teal-500 shadow-sm'
              : 'hover:bg-slate-50 dark:hover:bg-slate-800/50 border-2 border-transparent',
          ]"
        >
          <div class="flex items-center gap-3">
            <div :class="[
              'w-10 h-10 rounded-lg flex items-center justify-center font-semibold text-sm transition-colors',
              selectedNotaire?.id === n.id
                ? 'bg-teal-600 text-white'
                : 'bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400',
            ]">
              {{ n.nom.charAt(0) }}{{ n.prenom.charAt(0) }}
            </div>
            
            <div>
              <p :class="['font-semibold text-sm leading-tight', selectedNotaire?.id === n.id ? 'text-teal-900 dark:text-teal-100' : 'text-slate-700 dark:text-slate-300']">
                {{ n.prenom }} {{ n.nom }}
              </p>
              <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">
                {{ n.ville || "Mauricie" }}
              </p>
            </div>
          </div>
          
          <div class="flex items-center gap-1">
            <button
              v-if="selectedNotaire?.id === n.id"
              @click.stop="editNotaire(n)"
              class="p-1.5 hover:bg-teal-100 dark:hover:bg-teal-800/30 rounded-lg transition-colors"
              title="Modifier rapidement"
            >
              <Edit2 :size="14" class="text-teal-600 dark:text-teal-400" />
            </button>
            <Eye :size="16" :class="selectedNotaire?.id === n.id ? 'text-teal-600 dark:text-teal-400' : 'text-slate-300 dark:text-slate-700'" />
          </div>
        </div>

        <!-- Empty state -->
        <div v-if="filteredNotaires.length === 0" class="flex flex-col items-center justify-center py-12 text-slate-400">
          <Scale :size="40" class="mb-3 opacity-30" />
          <p class="text-sm">Aucun notaire trouvé</p>
        </div>

        <!-- MINI-FORM RAPIDE (sous le notaire sélectionné) -->
        <div v-if="showQuickForm && selectedNotaire" class="p-4 bg-teal-50/50 dark:bg-teal-900/10 rounded-xl border border-teal-200 dark:border-teal-800 animate-in slide-in-from-top-2 duration-200">
          <h3 class="text-sm font-bold text-teal-900 dark:text-teal-100 mb-3 flex items-center gap-2">
            <Sparkles :size="16" />
            {{ editingNotaire ? 'Modification rapide' : 'Création rapide' }}
          </h3>
          
          <FormKit
            type="form"
            v-model="quickFormData"
            @submit="handleQuickSubmit"
            :submit-attrs="{
              inputClass: 'w-full bg-teal-600 hover:bg-teal-700 text-white font-semibold py-2 px-4 rounded-lg transition-all text-sm active:scale-[0.98]'
            }"
            :submit-label="editingNotaire ? 'Sauvegarder' : 'Créer'"
          >
            <div class="space-y-3">
              <div class="grid grid-cols-2 gap-2">
                <FormKit 
                  type="select" 
                  name="civilite" 
                  label="Civilité" 
                  :options="['Me', 'Mtre']"
                  :classes="{ label: 'text-xs font-medium text-slate-700 dark:text-slate-300 mb-1', input: 'text-sm px-3 py-2 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20' }"
                />
                <FormKit 
                  type="text" 
                  name="ville" 
                  label="Ville"
                  placeholder="Mauricie"
                  :classes="{ label: 'text-xs font-medium text-slate-700 dark:text-slate-300 mb-1', input: 'text-sm px-3 py-2 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 placeholder-slate-400 focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20' }"
                />
              </div>
              
              <FormKit 
                type="text" 
                name="prenom" 
                label="Prénom" 
                validation="required"
                :classes="{ label: 'text-xs font-medium text-slate-700 dark:text-slate-300 mb-1', input: 'text-sm px-3 py-2 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20' }"
              />
              
              <FormKit 
                type="text" 
                name="nom" 
                label="Nom" 
                validation="required"
                :classes="{ label: 'text-xs font-medium text-slate-700 dark:text-slate-300 mb-1', input: 'text-sm px-3 py-2 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20' }"
              />
              
              <FormKit 
                type="text" 
                name="telephone" 
                label="Téléphone"
                :classes="{ label: 'text-xs font-medium text-slate-700 dark:text-slate-300 mb-1', input: 'text-sm px-3 py-2 rounded-lg border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-900 dark:text-slate-100 focus:border-teal-500 focus:ring-1 focus:ring-teal-500/20' }"
              />
            </div>

            <button 
              type="button"
              @click="openFullModal"
              class="w-full mt-2 text-xs text-teal-600 dark:text-teal-400 hover:text-teal-700 dark:hover:text-teal-300 font-medium flex items-center justify-center gap-1"
            >
              <MoreVertical :size="14" />
              Formulaire complet
            </button>
          </FormKit>
        </div>
      </div>
    </div>

    <!-- MAIN CONTENT -->
    <div class="flex-1 flex flex-col gap-5 overflow-y-auto">
      <template v-if="selectedNotaire">
        <NotaireDetails :notaire="selectedNotaire" />
        
        <!-- Form complet (accessible via scroll) -->
        <NotaireForm :notaire="selectedNotaire" @updated="loadNotaires" />
      </template>

      <!-- Empty state -->
      <div v-else class="flex-1 flex items-center justify-center border-2 border-dashed border-slate-200 dark:border-slate-700 rounded-2xl bg-slate-50/50 dark:bg-slate-900/50">
        <div class="text-center text-slate-400 space-y-3">
          <Scale :size="48" class="mx-auto opacity-20" />
          <div>
            <p class="text-base font-medium text-slate-500 dark:text-slate-400">Sélectionnez un notaire</p>
            <p class="text-sm text-slate-400 dark:text-slate-500">Consultez les détails et les clients liés</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal complète -->
    <NotaireAddModal 
      :is-open="showAddModal"
      :notaire="modalNotaire"
      @close="showAddModal = false; modalNotaire = null" 
      @saved="handleModalSaved" 
    />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { Search, Eye, Scale, Plus, Edit2, Sparkles, MoreVertical } from "lucide-vue-next";
import NotaireDetails from "../components/Notaires/Notairedetails.vue";
import NotaireForm from "../components/Notaires/Notaireform.vue";
import NotaireAddModal from '../components/Notaires/NotaireAddModal.vue';

const notaires = ref([]);
const selectedNotaire = ref(null);
const searchQuery = ref("");
const showAddModal = ref(false);
const showQuickForm = ref(false);
const editingNotaire = ref(null);
const modalNotaire = ref(null); // Données pour la modale
const quickFormData = ref({
  civilite: 'Me',
  prenom: '',
  nom: '',
  ville: '',
  telephone: ''
});

const loadNotaires = async () => {
  try {
    notaires.value = await window.go.main.App.GetAllNotaires();
  } catch (err) {
    console.error(err);
  }
};

const filteredNotaires = computed(() => {
  const q = searchQuery.value.toLowerCase();
  return notaires.value.filter(
    (n) =>
      `${n.prenom} ${n.nom}`.toLowerCase().includes(q) ||
      (n.ville && n.ville.toLowerCase().includes(q)),
  );
});

const toggleQuickForm = () => {
  showQuickForm.value = !showQuickForm.value;
  if (!showQuickForm.value) {
    editingNotaire.value = null;
    resetQuickForm();
  }
};

const resetQuickForm = () => {
  quickFormData.value = {
    civilite: 'Me',
    prenom: '',
    nom: '',
    ville: '',
    telephone: ''
  };
};

const selectNotaire = (notaire) => {
  selectedNotaire.value = notaire;
  showQuickForm.value = false;
  editingNotaire.value = null;
  resetQuickForm();
};

const editNotaire = (notaire) => {
  editingNotaire.value = notaire;
  showQuickForm.value = true;
  quickFormData.value = {
    id: notaire.id,
    civilite: notaire.civilite,
    prenom: notaire.prenom,
    nom: notaire.nom,
    ville: notaire.ville || '',
    telephone: notaire.telephone || ''
  };
};

const handleQuickSubmit = async (data) => {
  try {
    if (editingNotaire.value) {
      // Modification
      await window.go.main.App.UpdateNotaire({
        ...editingNotaire.value,
        ...data
      });
    } else {
      // Création
      await window.go.main.App.CreateNotaire({ ...data, created_by: 1 });
    }
    
    await loadNotaires();
    showQuickForm.value = false;
    editingNotaire.value = null;
    resetQuickForm();
  } catch (err) {
    console.error("Erreur:", err);
  }
};

const handleModalSaved = async () => {
  await loadNotaires();
  showAddModal.value = false;
  modalNotaire.value = null;
};

const openFullModal = () => {
  // Si on édite un notaire, on pré-remplit la modale
  if (editingNotaire.value) {
    modalNotaire.value = { ...editingNotaire.value };
  } else {
    modalNotaire.value = null;
  }
  showAddModal.value = true;
  showQuickForm.value = false;
};

onMounted(loadNotaires);
</script>