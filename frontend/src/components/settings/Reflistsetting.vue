<template>
  <div class="space-y-8">

    <!-- En-tête -->
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-3">
        <div class="p-2.5 bg-gradient-to-br from-teal-500/10 to-emerald-500/10 rounded-xl">
          <List :size="22" class="text-teal-600 dark:text-teal-400" />
        </div>
        <div>
          <h2 class="text-base font-black uppercase tracking-[0.15em] text-slate-700 dark:text-slate-300">
            Gestion des listes
          </h2>
          <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">
            Menus déroulants de l'application
          </p>
        </div>
      </div>

      <!-- Bouton Import CSV -->
      <button
        @click="ouvrirImportCSV"
        class="group flex items-center gap-2 px-4 py-2.5 bg-gradient-to-r from-teal-600 to-emerald-600 hover:from-teal-500 hover:to-emerald-500 text-white rounded-xl text-xs font-bold shadow-lg shadow-teal-500/20 transition-all duration-300"
      >
        <Upload :size="16" />
        Importer CSV
      </button>
    </div>

    <!-- Message rapport import -->
    <div v-if="rapportImport" :class="[
      'p-4 rounded-xl text-xs font-medium whitespace-pre-line border',
      rapportImport.includes('⚠️')
        ? 'bg-amber-50 dark:bg-amber-900/10 border-amber-200 dark:border-amber-800 text-amber-800 dark:text-amber-300'
        : 'bg-emerald-50 dark:bg-emerald-900/10 border-emerald-200 dark:border-emerald-800 text-emerald-800 dark:text-emerald-300'
    ]">
      {{ rapportImport }}
      <button @click="rapportImport = ''" class="ml-3 underline opacity-60 hover:opacity-100">fermer</button>
    </div>

    <!-- Chargement -->
    <div v-if="chargement" class="flex items-center justify-center py-16">
      <div class="flex flex-col items-center gap-3">
        <div class="w-8 h-8 border-2 border-teal-500/30 border-t-teal-500 rounded-full animate-spin"></div>
        <p class="text-xs text-slate-400">Chargement des listes...</p>
      </div>
    </div>

    <template v-else>
      <!-- Grille des catégories -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div
          v-for="categorie in categories"
          :key="categorie"
          class="group bg-white/80 dark:bg-slate-900/50 border border-slate-200/50 dark:border-slate-700/50 rounded-2xl overflow-hidden hover:border-teal-500/30 transition-all duration-300"
        >
          <!-- En-tête catégorie -->
          <div class="flex items-center justify-between px-5 py-4 bg-gradient-to-r from-slate-50/80 to-transparent dark:from-slate-800/50 border-b border-slate-200/50 dark:border-slate-700/50">
            <div class="flex items-center gap-2">
              <Tag :size="14" class="text-teal-500" />
              <h3 class="text-xs font-black uppercase tracking-[0.12em] text-slate-600 dark:text-slate-400">
                {{ formaterCategorie(categorie) }}
              </h3>
              <span class="text-[10px] px-2 py-0.5 bg-slate-100 dark:bg-slate-800 text-slate-500 rounded-full font-bold">
                {{ (listesParCategorie[categorie] || []).length }}
              </span>
            </div>
            <!-- Bouton ajouter -->
            <button
              @click="ouvrirAjout(categorie)"
              class="flex items-center gap-1.5 px-3 py-1.5 text-[10px] font-bold text-teal-600 dark:text-teal-400 hover:bg-teal-50 dark:hover:bg-teal-900/20 rounded-lg transition-all duration-200"
            >
              <Plus :size="12" />
              Ajouter
            </button>
          </div>

          <!-- Liste des entrées -->
          <div class="divide-y divide-slate-100/50 dark:divide-slate-800/50 max-h-64 overflow-y-auto scroll-area">
            <div
              v-for="item in listesParCategorie[categorie]"
              :key="item.id"
              class="flex items-center justify-between px-5 py-3 hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-colors group/item"
            >
              <div class="flex items-center gap-3">
                <!-- Badge couleur -->
                <span :class="[
                  'w-2 h-2 rounded-full shrink-0',
                  `bg-${item.couleur}-400`
                ]"></span>
                <span class="text-sm text-slate-700 dark:text-slate-300 font-medium">
                  {{ item.Libelle }}
                </span>
                <!-- Badge système -->
                <span
                  v-if="item.is_systeme"
                  class="text-[9px] px-1.5 py-0.5 bg-slate-100 dark:bg-slate-800 text-slate-400 rounded font-black uppercase tracking-wide"
                >
                  système
                </span>
              </div>

              <!-- Actions (seulement pour entrées utilisateur) -->
              <div v-if="!item.is_systeme" class="flex items-center gap-1 opacity-0 group-hover/item:opacity-100 transition-opacity">
                <button
                  @click="ouvrirModification(item)"
                  class="p-1.5 text-slate-400 hover:text-teal-500 hover:bg-teal-50 dark:hover:bg-teal-900/20 rounded-lg transition-all"
                >
                  <Pencil :size="13" />
                </button>
                <button
                  @click="confirmerSuppression(item)"
                  class="p-1.5 text-slate-400 hover:text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-all"
                >
                  <Trash2 :size="13" />
                </button>
              </div>
            </div>

            <!-- Vide -->
            <div v-if="(listesParCategorie[categorie] || []).length === 0" class="px-5 py-6 text-center">
              <p class="text-xs text-slate-400">Aucune entrée</p>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!-- ── MODAL : Ajouter / Modifier ── -->
    <Teleport to="body">
      <div v-if="modalOuverte" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="fermerModal"></div>
        <div class="relative w-full max-w-md bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border border-slate-200/50 dark:border-slate-700/50 overflow-hidden">

          <!-- En-tête modal -->
          <div class="flex items-center justify-between px-6 py-5 border-b border-slate-100 dark:border-slate-800">
            <h3 class="text-sm font-black uppercase tracking-wide text-slate-700 dark:text-slate-300">
              {{ modeModal === 'ajout' ? 'Ajouter une entrée' : 'Modifier l\'entrée' }}
            </h3>
            <button @click="fermerModal" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
              <X :size="16" class="text-slate-400" />
            </button>
          </div>

          <!-- Formulaire -->
          <div class="p-6 space-y-5">

            <!-- Catégorie (lecture seule en modification) -->
            <div>
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">Catégorie</label>
              <div v-if="modeModal === 'modification'" class="px-4 py-3 bg-slate-50 dark:bg-slate-800 rounded-xl text-sm text-slate-500 font-medium">
                {{ formaterCategorie(formData.categorie) }}
              </div>
              <select v-else v-model="formData.categorie" class="input-base">
                <option value="">-- Choisir --</option>
                <option v-for="cat in categories" :key="cat" :value="cat">
                  {{ formaterCategorie(cat) }}
                </option>
                <option value="__nouvelle__">+ Nouvelle catégorie...</option>
              </select>
            </div>

            <!-- Nouvelle catégorie (si sélectionné) -->
            <div v-if="formData.categorie === '__nouvelle__'">
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">
                Nom de la nouvelle catégorie
              </label>
              <input
                v-model="nouvelleCategorie"
                type="text"
                placeholder="ex: type_document"
                class="input-base"
              />
              <p class="text-[10px] text-slate-400 mt-1.5 ml-1">Utiliser des underscores, pas d'espaces</p>
            </div>

            <!-- Libellé -->
            <div>
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">Libellé *</label>
              <input v-model="formData.libelle" type="text" placeholder="ex: Acupuncteur" class="input-base" />
            </div>

            <!-- Couleur -->
            <div>
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">Couleur du badge</label>
              <div class="flex flex-wrap gap-2">
                <button
                  v-for="couleur in couleursDisponibles"
                  :key="couleur"
                  @click="formData.couleur = couleur"
                  :class="[
                    'w-7 h-7 rounded-lg border-2 transition-all duration-200',
                    `bg-${couleur}-400`,
                    formData.couleur === couleur
                      ? 'border-slate-700 dark:border-white scale-110 shadow-md'
                      : 'border-transparent hover:scale-105'
                  ]"
                  :title="couleur"
                ></button>
              </div>
              <p class="text-[10px] text-slate-400 mt-2 ml-1">Sélectionné : <strong>{{ formData.couleur }}</strong></p>
            </div>

          </div>

          <!-- Actions modal -->
          <div class="flex justify-end gap-3 px-6 py-4 border-t border-slate-100 dark:border-slate-800">
            <button @click="fermerModal" class="px-5 py-2.5 text-sm font-bold text-slate-500 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
              Annuler
            </button>
            <button
              @click="sauvegarder"
              :disabled="!formValide || enSauvegarde"
              class="px-6 py-2.5 bg-gradient-to-r from-teal-600 to-emerald-600 hover:from-teal-500 hover:to-emerald-500 disabled:opacity-50 disabled:cursor-not-allowed text-white text-sm font-bold rounded-xl shadow-lg shadow-teal-500/20 transition-all duration-300"
            >
              <span v-if="enSauvegarde">Sauvegarde...</span>
              <span v-else>{{ modeModal === 'ajout' ? 'Ajouter' : 'Modifier' }}</span>
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- ── MODAL : Confirmation suppression ── -->
    <Teleport to="body">
      <div v-if="itemASupprimer" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="itemASupprimer = null"></div>
        <div class="relative w-full max-w-sm bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border border-red-200/50 dark:border-red-800/30 overflow-hidden">
          <div class="p-6 text-center space-y-4">
            <div class="inline-flex p-3 bg-red-100 dark:bg-red-900/20 rounded-2xl">
              <Trash2 :size="24" class="text-red-500" />
            </div>
            <div>
              <h3 class="text-sm font-black text-slate-700 dark:text-slate-300">Supprimer cette entrée ?</h3>
              <p class="text-xs text-slate-500 mt-1">
                <strong class="text-slate-700 dark:text-slate-300">{{ itemASupprimer?.libelle }}</strong>
                sera supprimé définitivement.
              </p>
            </div>
            <div class="flex gap-3 justify-center pt-2">
              <button @click="itemASupprimer = null" class="px-5 py-2.5 text-sm font-bold text-slate-500 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
                Annuler
              </button>
              <button @click="supprimerConfirme" class="px-5 py-2.5 bg-red-500 hover:bg-red-600 text-white text-sm font-bold rounded-xl transition-colors">
                Supprimer
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { List, Upload, Tag, Plus, Pencil, Trash2, X } from 'lucide-vue-next'

// ─── État principal ───────────────────────────────────────────
const categories        = ref([])
const listesParCategorie = ref({})
const chargement        = ref(true)
const rapportImport     = ref('')

// ─── Modal Ajout / Modification ──────────────────────────────
const modalOuverte      = ref(false)
const modeModal         = ref('ajout') // 'ajout' | 'modification'
const itemEnModification = ref(null)
const enSauvegarde      = ref(false)
const nouvelleCategorie = ref('')

const formData = ref({
  categorie: '',
  libelle:   '',
  couleur:   'slate',
  icone:     '',
  ordre:     0
})

// ─── Suppression ─────────────────────────────────────────────
const itemASupprimer = ref(null)

// ─── Couleurs disponibles ─────────────────────────────────────
const couleursDisponibles = [
  'slate', 'zinc', 'red', 'orange', 'amber',
  'yellow', 'lime', 'green', 'emerald', 'teal',
  'cyan', 'blue', 'indigo', 'violet', 'purple', 'pink'
]

// ─── Computed ────────────────────────────────────────────────
const formValide = computed(() => {
  const cat = formData.value.categorie === '__nouvelle__'
    ? nouvelleCategorie.value.trim()
    : formData.value.categorie
  return cat !== '' && formData.value.libelle.trim() !== ''
})

// ─── Chargement des données ───────────────────────────────────
async function chargerTout() {
  chargement.value = true
  try {
    const cats = await window.go.main.App.GetAllRefCategories()
    console.log('CATS:', cats)
    categories.value = cats || []

    // Charger les entrées de chaque catégorie en parallèle
    const resultats = await Promise.all(
      categories.value.map(cat => window.go.main.App.GetRefListeByCategorie(cat))
    )
    const map = {}
    categories.value.forEach((cat, i) => {
      map[cat] = resultats[i] || []
    })
    listesParCategorie.value = map
  } catch (err) {
    console.error('Erreur chargement ref_listes:', err)
  } finally {
    chargement.value = false
  }
}

// ─── Formatage ────────────────────────────────────────────────
function formaterCategorie(cat) {
  return cat
    .replace(/_/g, ' ')
    .replace(/\b\w/g, l => l.toUpperCase())
}

// ─── Modal Ajout ─────────────────────────────────────────────
function ouvrirAjout(categorie = '') {
  modeModal.value = 'ajout'
  itemEnModification.value = null
  nouvelleCategorie.value = ''
  formData.value = { categorie, libelle: '', couleur: 'slate', icone: '', ordre: 0 }
  modalOuverte.value = true
}

// ─── Modal Modification ───────────────────────────────────────
function ouvrirModification(item) {
  modeModal.value = 'modification'
  itemEnModification.value = item
  formData.value = {
    categorie: item.categorie,
    libelle:   item.libelle,
    couleur:   item.couleur || 'slate',
    icone:     item.icone || '',
    ordre:     item.ordre || 0
  }
  modalOuverte.value = true
}

function fermerModal() {
  modalOuverte.value = false
  itemEnModification.value = null
}

// ─── Sauvegarde (Ajout ou Modification) ───────────────────────
async function sauvegarder() {
  if (!formValide.value) return
  enSauvegarde.value = true

  // Résoudre la catégorie finale
  const categorieFinale = formData.value.categorie === '__nouvelle__'
    ? nouvelleCategorie.value.trim().toLowerCase().replace(/\s+/g, '_')
    : formData.value.categorie

  try {
    const payload = { ...formData.value, categorie: categorieFinale }

    if (modeModal.value === 'ajout') {
      await window.go.main.App.CreateRefListe(payload)
    } else {
      await window.go.main.App.UpdateRefListe(itemEnModification.value.id, payload)
    }

    fermerModal()
    await chargerTout()
  } catch (err) {
    console.error('Erreur sauvegarde:', err)
  } finally {
    enSauvegarde.value = false
  }
}

// ─── Suppression ─────────────────────────────────────────────
function confirmerSuppression(item) {
  itemASupprimer.value = item
}

async function supprimerConfirme() {
  if (!itemASupprimer.value) return
  try {
    await window.go.main.App.DeleteRefListe(itemASupprimer.value.id)
    itemASupprimer.value = null
    await chargerTout()
  } catch (err) {
    console.error('Erreur suppression:', err)
  }
}

// ─── Import CSV ───────────────────────────────────────────────
async function ouvrirImportCSV() {
  try {
    rapportImport.value = '⏳ Import en cours...'
    const rapport = await window.go.main.App.OuvrirDialogCSVRefListes()
    if (rapport) {
      rapportImport.value = rapport
      await chargerTout()
    } else {
      rapportImport.value = ''
    }
  } catch (err) {
    rapportImport.value = `❌ Erreur: ${err}`
  }
}

// ─── Init ─────────────────────────────────────────────────────
onMounted(chargerTout)
</script>

<style scoped>
@reference "tailwindcss";

.input-base {
  @apply w-full bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70
         rounded-xl px-4 py-3 text-sm outline-none transition-all duration-300
         focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500
         dark:text-white placeholder:text-slate-400;
}

.scroll-area {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e1 transparent;
}
.scroll-area::-webkit-scrollbar { width: 4px; }
.scroll-area::-webkit-scrollbar-track { background: transparent; }
.scroll-area::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #cbd5e1, #94a3b8);
  border-radius: 10px;
}
</style>