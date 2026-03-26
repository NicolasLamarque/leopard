<template>
  <div class="space-y-6">

    <!-- ══ EN-TÊTE ══ -->
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-3">
        <div class="p-2.5 bg-gradient-to-br from-teal-500/10 to-sky-500/10 rounded-xl">
          <Stethoscope :size="22" class="text-teal-600 dark:text-teal-400" />
        </div>
        <div>
          <h2 class="text-base font-black uppercase tracking-[0.15em] text-slate-700 dark:text-slate-300">
            Diagnostics CIM-11
          </h2>
          <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">
            {{ totalCodes.toLocaleString('fr-CA') }} codes · v{{ meta?.version_oms || '—' }} · OMS
          </p>
        </div>
      </div>

      <div class="flex items-center gap-2">
        <!-- Statut sync -->
        <div class="flex items-center gap-1.5 px-3 py-1.5 rounded-xl text-xs font-bold"
          :class="meta?.statut === 'ok'
            ? 'bg-teal-50 dark:bg-teal-900/20 text-teal-600 dark:text-teal-400'
            : 'bg-amber-50 dark:bg-amber-900/20 text-amber-600 dark:text-amber-400'"
        >
          <component :is="meta?.statut === 'ok' ? CheckCircle2 : AlertCircle" :size="13" />
          {{ meta?.statut === 'ok' ? 'Synchronisé' : 'Non synchronisé' }}
        </div>
        <!-- Bouton sync -->
        <button
          @click="lancerSync"
          :disabled="syncing"
          class="group flex items-center gap-2 px-4 py-2.5
                 bg-gradient-to-r from-teal-600 to-sky-600
                 hover:from-teal-500 hover:to-sky-500
                 disabled:opacity-50 disabled:cursor-not-allowed
                 text-white rounded-xl text-xs font-bold
                 shadow-lg shadow-teal-500/20 transition-all duration-300"
        >
          <div v-if="syncing" class="w-3.5 h-3.5 border-2 border-white border-t-transparent rounded-full animate-spin" />
          <RefreshCw v-else :size="14" />
          {{ syncing ? syncLabel : 'Synchroniser OMS' }}
        </button>
      </div>
    </div>

    <!-- ══ BARRE DE RECHERCHE + FILTRES ══ -->
    <div class="flex items-center gap-3">
      <div class="relative flex-1">
        <Search :size="15" class="absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
        <input
          v-model="recherche"
          @input="onRecherche"
          type="text"
          placeholder="Rechercher un code ou un diagnostic…"
          class="w-full bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70 rounded-xl px-4 py-2.5 pl-10 pr-10 text-sm outline-none transition-all duration-300 focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500 dark:text-white placeholder:text-slate-400"
        />
        <button
          v-if="recherche"
          @click="recherche = ''; chargerCodes()"
          class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600"
        >
          <X :size="14" />
        </button>
      </div>

      <!-- Filtre chapitre -->
      <select v-model="filtreChapitreCode" @change="chargerCodes()" class="w-64 bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70 rounded-xl px-4 py-2.5 text-sm outline-none transition-all duration-300 focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500 dark:text-white placeholder:text-slate-400">
        <option value="">Tous les chapitres</option>
        <option v-for="ch in chapitres" :key="ch.code" :value="ch.code">
          {{ ch.code }} — {{ ch.titre_fr || ch.titre_en }}
        </option>
      </select>

      <!-- Toggle langue -->
      <div class="flex rounded-xl border border-slate-200 dark:border-slate-700 overflow-hidden shrink-0">
        <button
          @click="langue = 'fr'"
          :class="langue === 'fr'
            ? 'bg-teal-600 text-white'
            : 'bg-white dark:bg-slate-900 text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800'"
          class="px-3 py-2 text-xs font-bold transition-all"
        >FR</button>
        <button
          @click="langue = 'en'"
          :class="langue === 'en'
            ? 'bg-sky-600 text-white'
            : 'bg-white dark:bg-slate-900 text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800'"
          class="px-3 py-2 text-xs font-bold transition-all"
        >EN</button>
      </div>
    </div>

    <!-- ══ MESSAGE SYNC ══ -->
    <div v-if="syncResult" class="flex items-center gap-3 p-4 rounded-xl border"
      :class="syncResult.erreur
        ? 'bg-red-50 dark:bg-red-900/10 border-red-200 dark:border-red-800'
        : 'bg-teal-50 dark:bg-teal-900/10 border-teal-200 dark:border-teal-800'"
    >
      <component :is="syncResult.erreur ? AlertCircle : CheckCircle2" :size="16"
        :class="syncResult.erreur ? 'text-red-500' : 'text-teal-500'" />
      <div class="flex-1">
        <p class="text-xs font-bold" :class="syncResult.erreur ? 'text-red-700 dark:text-red-300' : 'text-teal-700 dark:text-teal-300'">
          {{ syncResult.erreur ? 'Erreur de synchronisation' : 'Synchronisation terminée' }}
        </p>
        <p class="text-xs mt-0.5" :class="syncResult.erreur ? 'text-red-500' : 'text-teal-500'">
          {{ syncResult.erreur || `${syncResult.nb_inseres?.toLocaleString('fr-CA')} insérés · ${syncResult.nb_mis_a_jour?.toLocaleString('fr-CA')} mis à jour · ${syncResult.duree}` }}
        </p>
      </div>
      <button @click="syncResult = null" class="text-slate-400 hover:text-slate-600">
        <X :size="14" />
      </button>
    </div>

    <!-- ══ CHARGEMENT ══ -->
    <div v-if="chargement" class="flex justify-center py-16">
      <div class="flex flex-col items-center gap-3">
        <div class="w-8 h-8 border-2 border-teal-500/30 border-t-teal-500 rounded-full animate-spin" />
        <p class="text-xs text-slate-400">Chargement des diagnostics...</p>
      </div>
    </div>

    <!-- ══ TABLE ══ -->
    <div v-else class="bg-white/80 dark:bg-slate-900/50 border border-slate-200/50 dark:border-slate-700/50 rounded-2xl overflow-hidden">

      <!-- Header table -->
      <div class="grid grid-cols-12 gap-4 px-5 py-3
                  bg-gradient-to-r from-slate-50/80 to-transparent dark:from-slate-800/50
                  border-b border-slate-200/50 dark:border-slate-700/50">
        <div class="col-span-2 text-[10px] font-black uppercase tracking-[0.12em] text-slate-500">Code</div>
        <div class="col-span-5 text-[10px] font-black uppercase tracking-[0.12em] text-slate-500">Diagnostic</div>
        <div class="col-span-3 text-[10px] font-black uppercase tracking-[0.12em] text-slate-500">Chapitre</div>
        <div class="col-span-2 text-[10px] font-black uppercase tracking-[0.12em] text-slate-500 text-right">Actions</div>
      </div>

      <!-- Lignes -->
      <div class="divide-y divide-slate-100/50 dark:divide-slate-800/50 max-h-[calc(100vh-400px)] overflow-y-auto scroll-area">
        <div
          v-for="code in codes"
          :key="code.id"
          class="grid grid-cols-12 gap-4 px-5 py-3
                 hover:bg-slate-50/50 dark:hover:bg-slate-800/30
                 transition-colors group/row"
        >
          <!-- Code -->
          <div class="col-span-2 flex items-center">
            <span class="font-mono text-xs font-bold px-2 py-1 rounded-lg
                         bg-teal-50 dark:bg-teal-900/30
                         text-teal-700 dark:text-teal-300
                         border border-teal-200/50 dark:border-teal-800/50">
              {{ code.code }}
            </span>
          </div>

          <!-- Titre -->
          <div class="col-span-5 flex flex-col justify-center min-w-0">
            <p class="text-sm text-slate-700 dark:text-slate-300 font-medium truncate">
              {{ langue === 'fr' ? (code.titre_fr || code.titre_en) : (code.titre_en || code.titre_fr) }}
            </p>
            <p v-if="code.titre_fr && code.titre_en && langue === 'fr'" class="text-[10px] text-slate-400 italic truncate mt-0.5">
              {{ code.titre_en }}
            </p>
          </div>

          <!-- Chapitre -->
          <div class="col-span-3 flex items-center min-w-0">
            <span class="text-xs text-slate-400 truncate">
              {{ code.chapitre_titre || code.chapitre_code || '—' }}
            </span>
          </div>

          <!-- Actions -->
          <div class="col-span-2 flex items-center justify-end gap-1
                      opacity-0 group-hover/row:opacity-100 transition-opacity">
            <button
              @click="ouvrirDetail(code)"
              class="p-1.5 text-slate-400 hover:text-sky-500 hover:bg-sky-50
                     dark:hover:bg-sky-900/20 rounded-lg transition-all"
              title="Voir détails"
            >
              <Eye :size="13" />
            </button>
            <button
              @click="ouvrirModification(code)"
              class="p-1.5 text-slate-400 hover:text-teal-500 hover:bg-teal-50
                     dark:hover:bg-teal-900/20 rounded-lg transition-all"
              title="Modifier"
            >
              <Pencil :size="13" />
            </button>
            <button
              @click="confirmerSuppression(code)"
              class="p-1.5 text-slate-400 hover:text-red-500 hover:bg-red-50
                     dark:hover:bg-red-900/20 rounded-lg transition-all"
              title="Désactiver"
            >
              <Trash2 :size="13" />
            </button>
          </div>
        </div>

        <!-- Vide -->
        <div v-if="codes.length === 0" class="flex flex-col items-center justify-center py-16 text-center">
          <div class="p-4 bg-slate-100 dark:bg-slate-800 rounded-full mb-3">
            <Search :size="24" class="text-slate-400" />
          </div>
          <p class="text-sm font-medium text-slate-500">Aucun résultat</p>
          <p class="text-xs text-slate-400 mt-1">Modifiez votre recherche ou synchronisez les données OMS</p>
        </div>
      </div>

      <!-- Footer pagination -->
      <div class="flex items-center justify-between px-5 py-3
                  border-t border-slate-200/50 dark:border-slate-700/50
                  bg-slate-50/50 dark:bg-slate-800/30">
        <p class="text-xs text-slate-400">
          {{ codes.length }} résultat{{ codes.length !== 1 ? 's' : '' }} affichés sur {{ totalCodes.toLocaleString('fr-CA') }}
        </p>
        <div class="flex items-center gap-2">
          <button
            @click="page--; chargerCodes()"
            :disabled="page <= 1"
            class="p-1.5 rounded-lg text-slate-400 hover:text-slate-600 hover:bg-slate-100
                   dark:hover:bg-slate-700 disabled:opacity-30 disabled:cursor-not-allowed transition-all"
          >
            <ChevronLeft :size="15" />
          </button>
          <span class="text-xs font-bold text-slate-500 px-2">{{ page }} / {{ totalPages }}</span>
          <button
            @click="page++; chargerCodes()"
            :disabled="page >= totalPages"
            class="p-1.5 rounded-lg text-slate-400 hover:text-slate-600 hover:bg-slate-100
                   dark:hover:bg-slate-700 disabled:opacity-30 disabled:cursor-not-allowed transition-all"
          >
            <ChevronRight :size="15" />
          </button>
        </div>
      </div>
    </div>

    <!-- ══ MODAL DÉTAIL / ENRICHISSEMENT ══ -->
    <Teleport to="body">
      <div v-if="codeDetail" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="codeDetail = null" />
        <div class="relative w-full max-w-lg bg-white dark:bg-slate-900 rounded-2xl shadow-2xl
                    border border-slate-200/50 dark:border-slate-700/50 overflow-hidden">

          <!-- Header -->
          <div class="flex items-start justify-between px-6 py-5
                      border-b border-slate-100 dark:border-slate-800
                      bg-gradient-to-r from-teal-50/50 dark:from-teal-900/10 to-transparent">
            <div>
              <div class="flex items-center gap-2 mb-1">
                <span class="font-mono text-sm font-black text-teal-600 dark:text-teal-400">
                  {{ codeDetail.code }}
                </span>
                <span class="text-[10px] px-2 py-0.5 bg-teal-100 dark:bg-teal-900/40
                             text-teal-600 dark:text-teal-400 rounded-full font-bold">
                  CIM-11
                </span>
              </div>
              <h3 class="text-sm font-bold text-slate-800 dark:text-slate-200 leading-snug">
                {{ codeDetail.titre_fr || codeDetail.titre_en }}
              </h3>
              <p v-if="codeDetail.titre_en && codeDetail.titre_fr" class="text-xs text-slate-400 italic mt-0.5">
                {{ codeDetail.titre_en }}
              </p>
            </div>
            <button @click="codeDetail = null"
              class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors shrink-0 ml-3">
              <X :size="16" class="text-slate-400" />
            </button>
          </div>

          <!-- Corps -->
          <div class="p-6 space-y-4">

            <!-- Infos -->
            <div class="grid grid-cols-2 gap-3">
              <div class="p-3 bg-slate-50 dark:bg-slate-800/60 rounded-xl">
                <p class="text-[10px] font-black uppercase tracking-wider text-slate-400 mb-1">Chapitre</p>
                <p class="text-xs font-medium text-slate-700 dark:text-slate-300">
                  {{ codeDetail.chapitre_titre || codeDetail.chapitre_code || '—' }}
                </p>
              </div>
              <div class="p-3 bg-slate-50 dark:bg-slate-800/60 rounded-xl">
                <p class="text-[10px] font-black uppercase tracking-wider text-slate-400 mb-1">Niveau</p>
                <p class="text-xs font-medium text-slate-700 dark:text-slate-300">
                  {{ niveauLabel(codeDetail.niveau) }}
                </p>
              </div>
            </div>

            <!-- Recherche API OMS en temps réel -->
            <div class="border border-sky-200/50 dark:border-sky-800/50 rounded-xl overflow-hidden">
              <div class="flex items-center justify-between px-4 py-3
                          bg-sky-50/50 dark:bg-sky-900/10 border-b border-sky-200/50 dark:border-sky-800/50">
                <div class="flex items-center gap-2">
                  <Globe :size="13" class="text-sky-500" />
                  <p class="text-[10px] font-black uppercase tracking-wider text-sky-600 dark:text-sky-400">
                    Informations OMS en ligne
                  </p>
                </div>
                <button
                  @click="enrichirDepuisOMS(codeDetail.code)"
                  :disabled="enrichLoading"
                  class="flex items-center gap-1.5 px-3 py-1.5 text-[10px] font-bold
                         bg-sky-500 hover:bg-sky-600 text-white rounded-lg transition-all
                         disabled:opacity-50"
                >
                  <div v-if="enrichLoading" class="w-3 h-3 border border-white border-t-transparent rounded-full animate-spin" />
                  <Sparkles v-else :size="11" />
                  {{ enrichLoading ? 'Chargement...' : 'Enrichir depuis OMS' }}
                </button>
              </div>
              <div class="px-4 py-3">
                <p v-if="!enrichData" class="text-xs text-slate-400 italic">
                  Cliquez pour récupérer la description complète, les exclusions et inclusions depuis l'API OMS.
                </p>
                <div v-else class="space-y-2">
                  <p v-if="enrichData.definition" class="text-xs text-slate-700 dark:text-slate-300 leading-relaxed">
                    {{ enrichData.definition }}
                  </p>
                  <div v-if="enrichData.inclusions?.length" class="space-y-1">
                    <p class="text-[10px] font-bold text-teal-600 uppercase">Inclusions</p>
                    <p v-for="i in enrichData.inclusions" :key="i" class="text-xs text-slate-500">• {{ i }}</p>
                  </div>
                  <div v-if="enrichData.exclusions?.length" class="space-y-1">
                    <p class="text-[10px] font-bold text-red-500 uppercase">Exclusions</p>
                    <p v-for="e in enrichData.exclusions" :key="e" class="text-xs text-slate-500">• {{ e }}</p>
                  </div>
                </div>
              </div>
            </div>

          </div>

          <!-- Footer -->
          <div class="flex justify-end gap-3 px-6 py-4 border-t border-slate-100 dark:border-slate-800">
            <button @click="codeDetail = null"
              class="px-5 py-2.5 text-sm font-bold text-slate-500 hover:bg-slate-100
                     dark:hover:bg-slate-800 rounded-xl transition-colors">
              Fermer
            </button>
            <button @click="ouvrirModification(codeDetail); codeDetail = null"
              class="px-5 py-2.5 bg-gradient-to-r from-teal-600 to-sky-600
                     hover:from-teal-500 hover:to-sky-500 text-white text-sm font-bold
                     rounded-xl shadow-lg shadow-teal-500/20 transition-all">
              Modifier
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- ══ MODAL MODIFICATION ══ -->
    <Teleport to="body">
      <div v-if="modalModif" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="modalModif = null" />
        <div class="relative w-full max-w-md bg-white dark:bg-slate-900 rounded-2xl shadow-2xl
                    border border-slate-200/50 dark:border-slate-700/50 overflow-hidden">

          <div class="flex items-center justify-between px-6 py-5 border-b border-slate-100 dark:border-slate-800">
            <h3 class="text-sm font-black uppercase tracking-wide text-slate-700 dark:text-slate-300">
              Modifier le code {{ modalModif.code }}
            </h3>
            <button @click="modalModif = null" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl">
              <X :size="16" class="text-slate-400" />
            </button>
          </div>

          <div class="p-6 space-y-4">
            <div>
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">Titre FR</label>
              <input v-model="formModif.titre_fr" type="text" class="w-full bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70 rounded-xl px-4 py-2.5 text-sm outline-none transition-all duration-300 focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500 dark:text-white placeholder:text-slate-400" />
            </div>
            <div>
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">Titre EN</label>
              <input v-model="formModif.titre_en" type="text" class="w-full bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70 rounded-xl px-4 py-2.5 text-sm outline-none transition-all duration-300 focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500 dark:text-white placeholder:text-slate-400" />
            </div>
            <div>
              <label class="block text-xs font-black uppercase tracking-[0.1em] text-slate-500 mb-2">Chapitre</label>
              <input v-model="formModif.chapitre_titre" type="text" class="w-full bg-white dark:bg-slate-900 border border-slate-200/70 dark:border-slate-700/70 rounded-xl px-4 py-2.5 text-sm outline-none transition-all duration-300 focus:ring-2 focus:ring-teal-500/20 focus:border-teal-500 dark:text-white placeholder:text-slate-400" />
            </div>
          </div>

          <div class="flex justify-end gap-3 px-6 py-4 border-t border-slate-100 dark:border-slate-800">
            <button @click="modalModif = null"
              class="px-5 py-2.5 text-sm font-bold text-slate-500 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl">
              Annuler
            </button>
            <button @click="sauvegarderModif"
              :disabled="enSauvegarde"
              class="px-6 py-2.5 bg-gradient-to-r from-teal-600 to-sky-600
                     hover:from-teal-500 hover:to-sky-500 disabled:opacity-50
                     text-white text-sm font-bold rounded-xl shadow-lg transition-all">
              {{ enSauvegarde ? 'Sauvegarde...' : 'Sauvegarder' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- ══ MODAL SUPPRESSION ══ -->
    <Teleport to="body">
      <div v-if="codeASupprimer" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="codeASupprimer = null" />
        <div class="relative w-full max-w-sm bg-white dark:bg-slate-900 rounded-2xl shadow-2xl
                    border border-red-200/50 dark:border-red-800/30 overflow-hidden">
          <div class="p-6 text-center space-y-4">
            <div class="inline-flex p-3 bg-red-100 dark:bg-red-900/20 rounded-2xl">
              <Trash2 :size="24" class="text-red-500" />
            </div>
            <div>
              <h3 class="text-sm font-black text-slate-700 dark:text-slate-300">Désactiver ce code ?</h3>
              <p class="text-xs text-slate-500 mt-1">
                <strong class="font-mono text-teal-600">{{ codeASupprimer.code }}</strong>
                sera masqué des menus diagnostics.
              </p>
            </div>
            <div class="flex gap-3 justify-center pt-2">
              <button @click="codeASupprimer = null"
                class="px-5 py-2.5 text-sm font-bold text-slate-500 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl">
                Annuler
              </button>
              <button @click="supprimerConfirme"
                class="px-5 py-2.5 bg-red-500 hover:bg-red-600 text-white text-sm font-bold rounded-xl">
                Désactiver
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
import {
  Stethoscope, RefreshCw, Search, X, Eye, Pencil, Trash2,
  ChevronLeft, ChevronRight, CheckCircle2, AlertCircle,
  Globe, Sparkles
} from 'lucide-vue-next'

// ── État ─────────────────────────────────────────────────────────
const codes              = ref([])
const chapitres          = ref([])
const meta               = ref(null)
const chargement         = ref(true)
const recherche          = ref('')
const filtreChapitreCode = ref('')
const langue             = ref('fr')
const page               = ref(1)
const parPage            = 50
const totalCodes         = ref(0)

const syncing    = ref(false)
const syncLabel  = ref('Connexion OMS...')
const syncResult = ref(null)

const codeDetail    = ref(null)
const enrichLoading = ref(false)
const enrichData    = ref(null)

const modalModif  = ref(null)
const formModif   = ref({})
const enSauvegarde = ref(false)

const codeASupprimer = ref(null)

let rechercheDebounce = null

// ── Computed ──────────────────────────────────────────────────────
const totalPages = computed(() => Math.max(1, Math.ceil(totalCodes.value / parPage)))

// ── Chargement ────────────────────────────────────────────────────
const chargerMeta = async () => {
  try {
    meta.value = await window.go.main.App.GetCim11SyncMeta()
    totalCodes.value = meta.value?.nb_codes_fr || 0
  } catch (e) { console.error(e) }
}

const chargerChapitres = async () => {
  try {
    chapitres.value = await window.go.main.App.GetCim11Chapitres() || []
  } catch (e) { console.error(e) }
}

const chargerCodes = async () => {
  chargement.value = true
  try {
    const query = recherche.value.trim() || filtreChapitreCode.value || '%'
    const limit = parPage
    const offset = (page.value - 1) * parPage

    const results = await window.go.main.App.GetCim11Page(
      query,
      filtreChapitreCode.value,
      langue.value,
      limit,
      offset
    )
    codes.value = results?.codes || []
    totalCodes.value = results?.total || 0
  } catch (e) {
    console.error(e)
    codes.value = []
  } finally {
    chargement.value = false
  }
}

const onRecherche = () => {
  clearTimeout(rechercheDebounce)
  rechercheDebounce = setTimeout(() => {
    page.value = 1
    chargerCodes()
  }, 300)
}

// ── Sync OMS ──────────────────────────────────────────────────────
const lancerSync = async () => {
  syncing.value = true
  syncResult.value = null
  const msgs = ['Connexion OMS...', 'Récupération EN...', 'Récupération FR...', 'Insertion SQLite...']
  let i = 0
  syncLabel.value = msgs[0]
  const interval = setInterval(() => {
    syncLabel.value = msgs[++i % msgs.length]
  }, 4000)

  try {
    const result = await window.go.main.App.ImportCim11()
    syncResult.value = result
    await chargerMeta()
    await chargerCodes()
  } catch (e) {
    syncResult.value = { erreur: e?.message || 'Erreur inconnue' }
  } finally {
    clearInterval(interval)
    syncing.value = false
  }
}

// ── Enrichissement OMS ────────────────────────────────────────────
const enrichirDepuisOMS = async (code) => {
  enrichLoading.value = true
  enrichData.value = null
  try {
    const result = await window.go.main.App.GetCim11DetailOMS(code)
    enrichData.value = result
  } catch (e) {
    enrichData.value = { definition: 'Impossible de récupérer les données OMS.' }
  } finally {
    enrichLoading.value = false
  }
}

// ── CRUD ──────────────────────────────────────────────────────────
const ouvrirDetail = (code) => {
  codeDetail.value = code
  enrichData.value = null
}

const ouvrirModification = (code) => {
  modalModif.value = code
  formModif.value = {
    id:            code.id,
    code:          code.code,
    titre_fr:      code.titre_fr,
    titre_en:      code.titre_en,
    chapitre_code: code.chapitre_code,
    chapitre_titre:code.chapitre_titre,
    bloc_code:     code.bloc_code,
    bloc_titre:    code.bloc_titre,
    parent_code:   code.parent_code,
    niveau:        code.niveau,
    is_billable:   code.is_billable,
    actif:         code.actif,
  }
}

const sauvegarderModif = async () => {
  enSauvegarde.value = true
  try {
    await window.go.main.App.UpdateCim11Code(formModif.value)
    modalModif.value = null
    await chargerCodes()
  } catch (e) {
    console.error(e)
  } finally {
    enSauvegarde.value = false
  }
}

const confirmerSuppression = (code) => {
  codeASupprimer.value = code
}

const supprimerConfirme = async () => {
  if (!codeASupprimer.value) return
  try {
    await window.go.main.App.DeleteCim11Code(codeASupprimer.value.id)
    codeASupprimer.value = null
    await chargerCodes()
    await chargerMeta()
  } catch (e) { console.error(e) }
}

// ── Helpers ───────────────────────────────────────────────────────
const niveauLabel = (n) => {
  switch (n) {
    case 0: return 'Chapitre'
    case 1: return 'Bloc'
    case 2: return 'Catégorie'
    default: return 'Sous-catégorie'
  }
}

// ── Init ─────────────────────────────────────────────────────────
onMounted(async () => {
  await chargerMeta()
  await chargerChapitres()
  await chargerCodes()
})
</script>

<style scoped>
.scroll-area {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e1 transparent;
}
.scroll-area::-webkit-scrollbar { width: 4px; }
.scroll-area::-webkit-scrollbar-track { background: transparent; }
.scroll-area::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #0d9488, #0ea5e9);
  border-radius: 10px;
}
</style>