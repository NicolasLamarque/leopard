<template>
  <Transition name="modal-fade">
    <div
      v-if="isOpen"
      class="fixed inset-0 z-50 flex items-center justify-center p-4"
    >
      <div
        class="absolute inset-0 bg-slate-900/80 backdrop-blur-sm"
        @click="closeModal"
      ></div>

      <div
        class="relative bg-zinc-50 dark:bg-zinc-950 w-full max-w-7xl h-[95vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-zinc-200 dark:border-zinc-800"
      >
        <div
          class="px-6 py-5 border-b dark:border-zinc-800 flex justify-between items-center bg-zinc-900 dark:bg-black"
        >
          <div class="flex items-center gap-4">
            <div
              class="p-2.5 bg-zinc-800 rounded-xl border border-zinc-700 text-teal-400"
            >
              <ClipboardCheck :size="24" />
            </div>
            <div>
              <h2 class="text-xl font-bold text-zinc-100">Tenue de Dossier</h2>
              <p
                class="text-xs text-zinc-400 uppercase font-bold tracking-widest mt-0.5"
              >
                {{ client?.prenom }} {{ client?.nom }}
                <span class="mx-2 text-zinc-600">|</span> {{ leopardNumber }}
              </p>
            </div>
          </div>
          <div class="flex items-center gap-6">
            <div
              class="bg-zinc-800/50 px-5 py-2 rounded-2xl border border-zinc-700/50"
            >
              <div class="flex items-center gap-3">
                <span
                  class="text-xs text-zinc-400 font-bold uppercase tracking-tighter"
                  >Conformité</span
                >
                <div class="flex items-center gap-2">
                  <div
                    class="text-2xl font-black font-mono"
                    :class="scoreColor"
                  >
                    {{ conformiteScore }}%
                  </div>
                  <component :is="scoreIcon" :size="20" :class="scoreColor" />
                </div>
              </div>
            </div>

            <button
              @click="closeModal"
              class="p-2 hover:bg-zinc-800 text-zinc-400 hover:text-white rounded-full transition-colors"
            >
              <X :size="24" />
            </button>
          </div>
        </div>

        <div class="flex-1 overflow-y-auto p-8 bg-zinc-50 dark:bg-zinc-950">
          <div class="max-w-6xl mx-auto grid grid-cols-1 md:grid-cols-2 gap-6">
            <div
              class="bg-white dark:bg-zinc-900/40 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden flex flex-col"
            >
              <div
                class="p-4 border-b border-zinc-100 dark:border-zinc-800 flex items-center justify-between bg-zinc-50/50 dark:bg-zinc-900/50"
              >
                <div class="flex items-center gap-3">
                  <IdCard :size="18" class="text-teal-600 dark:text-teal-500" />
                  <h3
                    class="text-sm font-bold text-zinc-700 dark:text-zinc-300 uppercase tracking-tight"
                  >
                    Fiche d'identification
                  </h3>
                </div>
                <div
                  v-if="sections.ficheIdentification.existe"
                  class="text-teal-600 dark:text-teal-400"
                >
                  <CheckCircle :size="18" />
                </div>
                <div v-else class="text-amber-500 animate-pulse">
                  <AlertTriangle :size="18" />
                </div>
              </div>

              <div class="p-5 flex-1 space-y-4">
                <div
                  v-if="!sections.ficheIdentification.existe"
                  class="space-y-4"
                >
                  <div
                    class="flex items-start gap-3 p-3 bg-amber-500/5 border border-amber-500/20 rounded-lg"
                  >
                    <AlertTriangle
                      :size="16"
                      class="text-amber-600 flex-shrink-0 mt-0.5"
                    />
                    <p
                      class="text-xs text-amber-800 dark:text-amber-400 leading-relaxed"
                    >
                      <strong>Document manquant :</strong> La fiche
                      d'identification doit être générée pour ce dossier.
                    </p>
                  </div>

                  <div
                    v-if="!sections.ficheIdentification.donneesCompletes"
                    class="p-3 bg-zinc-100 dark:bg-zinc-800/50 rounded-lg border border-zinc-200 dark:border-zinc-700"
                  >
                    <p
                      class="text-[10px] font-bold text-zinc-500 uppercase mb-2"
                    >
                      Champs requis manquants :
                    </p>
                    <ul
                      class="text-xs text-zinc-600 dark:text-zinc-400 grid grid-cols-2 gap-1"
                    >
                      <li
                        v-for="champ in sections.ficheIdentification
                          .champsManquants"
                        :key="champ"
                        class="flex items-center gap-1"
                      >
                        <span class="w-1 h-1 bg-zinc-400 rounded-full"></span>
                        {{ champ }}
                      </li>
                    </ul>
                  </div>

                  <button
                    @click="creerFicheIdentification"
                    :disabled="
                      !sections.ficheIdentification.donneesCompletes ||
                      isGenerating.ficheIdentification
                    "
                    class="w-full flex items-center justify-center gap-2 px-4 py-2.5 bg-zinc-800 hover:bg-black dark:bg-teal-600 dark:hover:bg-teal-700 disabled:bg-zinc-200 text-white rounded-lg font-bold transition-all text-sm"
                  >
                    <FilePlus
                      v-if="!isGenerating.ficheIdentification"
                      :size="16"
                    />
                    <Loader2 v-else :size="16" class="animate-spin" />
                    <span>Générer le document</span>
                  </button>
                </div>

                <div
                  v-else
                  class="flex items-center justify-between p-3 bg-teal-500/5 border border-teal-500/20 rounded-lg"
                >
                  <div class="flex items-center gap-3">
                    <FileCheck :size="18" class="text-teal-600" />
                    <div class="flex flex-col">
                      <span
                        class="text-xs font-bold text-zinc-800 dark:text-zinc-200"
                        >{{ sections.ficheIdentification.nomFichier }}</span
                      >
                      <span class="text-[10px] text-zinc-500">{{
                        sections.ficheIdentification.dateCreation
                      }}</span>
                    </div>
                  </div>
                  <button
                    @click="ouvrirDocument('ficheIdentification')"
                    class="p-2 hover:bg-teal-500/10 text-teal-600 rounded-lg transition-all"
                  >
                    <ExternalLink :size="16" />
                  </button>
                </div>
              </div>
            </div>

            <div
              class="bg-white dark:bg-zinc-900/40 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden"
            >
              <div
                class="p-4 border-b border-zinc-100 dark:border-zinc-800 flex items-center justify-between bg-zinc-50/50 dark:bg-zinc-900/50"
              >
                <div class="flex items-center gap-3">
                  <FileSignature
                    :size="18"
                    class="text-teal-600 dark:text-teal-400"
                  />
                  <h3
                    class="text-sm font-bold text-zinc-700 dark:text-zinc-300 uppercase tracking-tight"
                  >
                    Consentement signé
                  </h3>
                </div>
                <div v-if="sections.consentement.existe" class="text-teal-600">
                  <CheckCircle :size="18" />
                </div>
                <div v-else class="text-amber-500">
                  <Clock :size="18" />
                </div>
              </div>

              <div class="p-5 space-y-4">
                <div v-if="!sections.consentement.existe" class="space-y-4">
                  <p class="text-xs text-zinc-500 leading-relaxed italic">
                    Le consentement général est requis pour la conformité légale
                    du dossier.
                  </p>
                  <button
                    @click="ajouterConsentement"
                    class="w-full flex items-center justify-center gap-2 px-4 py-2.5 bg-slate-700 hover:bg-slate-800 text-white rounded-lg font-bold transition-all text-sm"
                  >
                    <Upload :size="16" />
                    <span>Importer le scan</span>
                  </button>
                </div>

                <div
                  v-else
                  class="flex items-center justify-between p-3 bg-zinc-50 dark:bg-zinc-800 rounded-lg border border-zinc-200 dark:border-zinc-700"
                >
                  <div class="flex items-center gap-3">
                    <FileCheck :size="18" class="text-teal-600" />
                    <div>
                      <p
                        class="text-xs font-bold text-zinc-700 dark:text-zinc-300"
                      >
                        Signé le {{ sections.consentement.dateSignature }}
                      </p>
                    </div>
                  </div>
                  <button
                    @click="ouvrirDocument('consentement')"
                    class="p-2 hover:bg-zinc-200 dark:hover:bg-zinc-700 text-zinc-600 rounded-lg"
                  >
                    <ExternalLink :size="16" />
                  </button>
                </div>
              </div>
            </div>

            <div
              class="bg-white dark:bg-zinc-900/40 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden md:col-span-2"
            >
              <div
                class="p-4 border-b border-zinc-100 dark:border-zinc-800 flex items-center justify-between bg-zinc-50/50 dark:bg-zinc-900/50"
              >
                <div class="flex items-center gap-3">
                  <Target :size="18" class="text-teal-600" />
                  <h3
                    class="text-sm font-bold text-zinc-700 dark:text-zinc-300 uppercase tracking-tight"
                  >
                    Plan d'intervention (PI)
                  </h3>
                </div>
                <div
                  v-if="
                    sections.planIntervention.existe &&
                    !sections.planIntervention.echu
                  "
                  class="bg-teal-500/10 text-teal-600 px-3 py-1 rounded-full text-[10px] font-black uppercase border border-teal-500/20"
                >
                  Actif
                </div>
                <div
                  v-else-if="sections.planIntervention.echu"
                  class="bg-red-500/10 text-red-600 px-3 py-1 rounded-full text-[10px] font-black uppercase border border-red-500/20 animate-pulse"
                >
                  À renouveler
                </div>
              </div>

              <div class="p-5">
                <div
                  v-if="sections.planIntervention.existe"
                  class="grid grid-cols-1 md:grid-cols-3 gap-6 items-center"
                >
                  <div class="md:col-span-2 flex items-center gap-6">
                    <div class="flex flex-col">
                      <span
                        class="text-[10px] uppercase font-bold text-zinc-400"
                        >Version actuelle</span
                      >
                      <span
                        class="text-sm font-bold text-zinc-800 dark:text-zinc-200"
                        >PI - v{{ sections.planIntervention.version }}</span
                      >
                    </div>
                    <div
                      class="flex flex-col border-l border-zinc-200 dark:border-zinc-700 pl-6"
                    >
                      <span
                        class="text-[10px] uppercase font-bold text-zinc-400"
                        >Échéance</span
                      >
                      <span
                        class="text-sm font-bold"
                        :class="
                          sections.planIntervention.echu
                            ? 'text-red-600'
                            : 'text-zinc-700 dark:text-zinc-300'
                        "
                      >
                        {{ sections.planIntervention.dateEcheance }}
                      </span>
                    </div>
                  </div>
                  <div class="flex gap-2">
                    <button
                      @click="ouvrirDocument('planIntervention')"
                      class="flex-1 py-2 bg-zinc-100 dark:bg-zinc-800 hover:bg-zinc-200 dark:hover:bg-zinc-700 text-zinc-700 dark:text-zinc-300 rounded-lg text-sm font-bold border border-zinc-200 dark:border-zinc-700 transition-all"
                    >
                      Consulter
                    </button>
                    <button
                      @click="creerNouveauPI"
                      class="flex-1 py-2 bg-zinc-800 dark:bg-teal-600 hover:bg-black dark:hover:bg-teal-700 text-white rounded-lg text-sm font-bold shadow-md transition-all"
                    >
                      Nouvelle version
                    </button>
                  </div>
                </div>

                <button
                  v-else
                  @click="creerNouveauPI"
                  class="w-full flex items-center justify-center gap-2 px-4 py-3 bg-zinc-100 dark:bg-zinc-800 hover:bg-zinc-200 dark:hover:bg-zinc-700 text-zinc-700 dark:text-zinc-300 border-2 border-dashed border-zinc-300 dark:border-zinc-700 rounded-xl font-bold transition-all"
                >
                  <Plus :size="18" />
                  <span>Initialiser le Plan d'Intervention</span>
                </button>
              </div>
            </div>

            <div
              class="bg-white dark:bg-zinc-900/40 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden flex flex-col"
            >
              <div
                class="p-4 border-b border-zinc-100 dark:border-zinc-800 flex items-center justify-between bg-zinc-50/50 dark:bg-zinc-900/50"
              >
                <div class="flex items-center gap-3">
                  <ClipboardList
                    :size="18"
                    class="text-teal-600 dark:text-teal-400"
                  />
                  <h3
                    class="text-sm font-bold text-zinc-700 dark:text-zinc-300 uppercase tracking-tight"
                  >
                    Notes évolutives
                  </h3>
                </div>
                <div
                  class="text-[10px] font-bold text-zinc-500 bg-zinc-100 dark:bg-zinc-800 px-2 py-0.5 rounded"
                >
                  {{ sections.notesEvolutives.notesExportees }} /
                  {{ sections.notesEvolutives.totalNotes }}
                </div>
              </div>

              <div class="p-5 flex-1 space-y-4">
                <div class="space-y-2">
                  <div
                    class="flex justify-between text-[10px] font-bold uppercase text-zinc-500"
                  >
                    <span>Exportation au dossier</span>
                    <span
                      >{{ sections.notesEvolutives.pourcentageExporte }}%</span
                    >
                  </div>
                  <div
                    class="w-full bg-zinc-100 dark:bg-zinc-800 rounded-full h-1.5 overflow-hidden"
                  >
                    <div
                      class="h-full transition-all duration-500"
                      :class="
                        sections.notesEvolutives.pourcentageExporte === 100
                          ? 'bg-teal-500'
                          : 'bg-slate-400'
                      "
                      :style="{
                        width:
                          sections.notesEvolutives.pourcentageExporte + '%',
                      }"
                    ></div>
                  </div>
                </div>

                <button
                  v-if="sections.notesEvolutives.notesNonExportees > 0"
                  @click="exporterNotesManquantes"
                  :disabled="isGenerating.notes"
                  class="w-full py-2.5 bg-zinc-800 dark:bg-zinc-700 hover:bg-black dark:hover:bg-zinc-600 text-white rounded-lg text-sm font-bold transition-all flex items-center justify-center gap-2"
                >
                  <FileDown v-if="!isGenerating.notes" :size="16" />
                  <Loader2 v-else :size="16" class="animate-spin" />
                  <span
                    >Exporter
                    {{ sections.notesEvolutives.notesNonExportees }} notes</span
                  >
                </button>

                <div
                  v-else-if="sections.notesEvolutives.totalNotes === 0"
                  class="text-center py-2 bg-zinc-100 dark:bg-zinc-800/50 border border-dashed border-zinc-300 dark:border-zinc-700 rounded-lg"
                >
                  <span class="text-[10px] font-bold text-zinc-500 uppercase"
                    >Aucune note au dossier</span
                  >
                </div>

                <div
                  v-else
                  class="text-center py-2 bg-teal-500/5 border border-teal-500/10 rounded-lg"
                >
                  <span
                    class="text-xs font-bold text-teal-600 uppercase tracking-tighter"
                    >✓ Toutes les notes sont archivées</span
                  >
                </div>
              </div>
            </div>

            <div
              class="bg-white dark:bg-zinc-900/40 rounded-xl border border-zinc-200 dark:border-zinc-800 overflow-hidden flex flex-col"
            >
              <div
                class="p-4 border-b border-zinc-100 dark:border-zinc-800 flex items-center justify-between bg-zinc-50/50 dark:bg-zinc-900/50"
              >
                <div class="flex items-center gap-3">
                  <ListChecks
                    :size="18"
                    class="text-teal-500 dark:text-teal-400"
                  />
                  <h3
                    class="text-sm font-bold text-zinc-700 dark:text-zinc-300 uppercase tracking-tight"
                  >
                    Sommaire final
                  </h3>
                </div>
                <div v-if="sections.sommaire.aJour" class="text-teal-600">
                  <CheckCircle :size="18" />
                </div>
              </div>

              <div class="p-5 flex-1 flex flex-col justify-between">
                <div>
                  <p class="text-[10px] uppercase font-bold text-zinc-400 mb-1">
                    Dernière génération
                  </p>
                  <p
                    class="text-xs text-zinc-600 dark:text-zinc-400 font-mono italic"
                  >
                    {{
                      sections.sommaire.derniereMiseAJour ||
                      "Aucun document généré"
                    }}
                  </p>
                </div>

                <button
                  @click="genererSommaire"
                  :disabled="isGenerating.sommaire"
                  class="mt-4 w-full py-2.5 bg-white dark:bg-zinc-800 hover:bg-zinc-50 dark:hover:bg-zinc-700 text-zinc-800 dark:text-zinc-200 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm font-bold transition-all flex items-center justify-center gap-2"
                >
                  <FileText v-if="!isGenerating.sommaire" :size="16" />
                  <Loader2 v-else :size="16" class="animate-spin" />
                  <span>Générer le sommaire</span>
                </button>
              </div>
            </div>
          </div>
        </div>

        <div
          class="px-8 py-5 bg-zinc-100 dark:bg-zinc-900/80 border-t dark:border-zinc-800 flex items-center justify-between"
        >
          <div class="flex items-center gap-6">
            <div
              class="flex items-center gap-2 px-3 py-1 bg-white dark:bg-zinc-800 rounded-full border border-zinc-200 dark:border-zinc-700 shadow-sm"
            >
              <div class="w-2 h-2 rounded-full bg-teal-500"></div>
              <span
                class="text-[11px] font-bold text-zinc-600 dark:text-zinc-400 uppercase tracking-tight"
              >
                {{ sectionsConformes }} / {{ totalSections }} Sections validées
              </span>
            </div>
            <span class="text-[10px] text-zinc-400 font-medium italic">
              Vérifié le {{ new Date().toLocaleDateString("fr-CA") }} à
              {{
                new Date().toLocaleTimeString("fr-CA", {
                  hour: "2-digit",
                  minute: "2-digit",
                })
              }}
            </span>
          </div>

          <button
            @click="verifierToutLeDossier"
            :disabled="isVerifying"
            class="flex items-center gap-2 px-6 py-2 bg-zinc-800 hover:bg-black dark:bg-teal-600 dark:hover:bg-teal-700 text-white rounded-xl font-bold transition-all shadow-lg shadow-zinc-500/10 text-sm"
          >
            <RefreshCw
              v-if="!isVerifying"
              :size="16"
              :class="{ 'animate-spin': isVerifying }"
            />
            <span>Actualiser la conformité</span>
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed, watch } from "vue";
import {
  X,
  ClipboardCheck,
  CheckCircle,
  AlertTriangle,
  AlertCircle,
  IdCard,
  FileSignature,
  FileText,
  Target,
  ClipboardList,
  ListChecks,
  FilePlus,
  Upload,
  ExternalLink,
  FileCheck,
  Loader2,
  Plus,
  FileDown,
  RefreshCw,
  Clock,
} from "lucide-vue-next";

const props = defineProps({
  isOpen: Boolean,
  client: Object,
  leopardNumber: String,
});

const emit = defineEmits(["close", "refresh"]);

// États de génération
const isGenerating = ref({
  ficheIdentification: false,
  notes: false,
  sommaire: false,
});

const isVerifying = ref(false);

// État des sections obligatoires
const sections = ref({
  ficheIdentification: {
    existe: false,
    donneesCompletes: false,
    champsManquants: [],
    nomFichier: "",
    dateCreation: "",
    chemin: "",
  },
  consentement: {
    existe: false,
    nomFichier: "",
    dateSignature: "",
    chemin: "",
  },
  ficheOuverture: {
    existe: false,
    nomFichier: "",
    chemin: "",
  },
  planIntervention: {
    existe: false,
    version: 0,
    dateCreation: "",
    dateEcheance: "",
    echu: false,
    chemin: "",
  },
  notesEvolutives: {
    totalNotes: 0,
    notesExportees: 0,
    notesNonExportees: 0,
    pourcentageExporte: 0,
  },
  sommaire: {
    aJour: false,
    derniereMiseAJour: "",
  },
});

// Calculs de conformité
const totalSections = computed(() => 6);

const sectionsConformes = computed(() => {
  let count = 0;
  if (sections.value.ficheIdentification.existe) count++;
  if (sections.value.consentement.existe) count++;
  if (sections.value.ficheOuverture.existe) count++;
  if (
    sections.value.planIntervention.existe &&
    !sections.value.planIntervention.echu
  )
    count++;
  if (sections.value.notesEvolutives.pourcentageExporte === 100) count++;
  if (sections.value.sommaire.aJour) count++;
  return count;
});

const conformiteScore = computed(() => {
  return Math.round((sectionsConformes.value / totalSections.value) * 100);
});

const scoreColor = computed(() => {
  if (conformiteScore.value >= 100) return "text-green-400";
  if (conformiteScore.value >= 75) return "text-yellow-400";
  return "text-red-400";
});

const scoreIcon = computed(() => {
  if (conformiteScore.value >= 100) return CheckCircle;
  if (conformiteScore.value >= 75) return Clock;
  return AlertCircle;
});

// C'est ici qu'on stocke l'état réel du dossier après vérification
const tenueStatus = ref({
  notesEvolutives: {
    totalNotes: 0,
    notesExportees: 0,
    notesNonExportees: 0,
    pourcentageExporte: 0,
  },
  consentement: { existe: false },
  ficheOuverture: { existe: false },
  planIntervention: { existe: false },
});

// Actions
const creerFicheIdentification = async () => {
  isGenerating.value.ficheIdentification = true;
  try {
    // TODO: Appeler la génération de PDF
    await new Promise((resolve) => setTimeout(resolve, 2000)); // Simulation
    alert("✅ Fiche d'identification créée!");
    await verifierToutLeDossier();
  } catch (err) {
    alert("❌ Erreur: " + err.message);
  } finally {
    isGenerating.value.ficheIdentification = false;
  }
};

const ajouterConsentement = () => {
  alert("📄 Fonctionnalité scanner à venir");
};

const creerFicheOuverture = () => {
  alert("📝 Formulaire de fiche d'ouverture à venir");
};

const creerNouveauPI = () => {
  alert("🎯 Formulaire de Plan d'Intervention à venir");
};
// Dans GestionTenueDossier.vue
const exporterNotesManquantes = async () => {
  isGenerating.value.notes = true;
  try {
    // 1. Aller chercher les notes réelles via ton App Go
    const notes = await window.go.main.App.GetClientNotes(props.client.id);

    if (!notes || notes.length === 0) {
      alert("Aucune note à exporter.");
      return;
    }

    // 2. Utiliser jsPDF pour créer le document de tenue de dossier
    // C'est ici qu'on "boucle" sur le tableau 'notes' pour les écrire dans le PDF
    const doc = new jsPDF();
    let y = 20;

    notes.forEach((note, index) => {
      doc.setFontSize(12);
      doc.text(`Date: ${note.date_intervention}`, 10, y);
      doc.text(`Sujet: ${note.sujet}`, 10, y + 7);
      // ... ajouter le contenu et la signature ...
      y += 50; // On descend pour la note suivante

      // Gérer le saut de page si on arrive en bas
      if (y > 270) {
        doc.addPage();
        y = 20;
      }
    });

    // 3. Convertir en Base64 pour l'envoyer à ton fichier files.go
    const pdfBase64 = doc.output("datauristring").split(",")[1];

    // 4. Sauvegarder dans le dossier Leopard du client
    const folderInfo = await window.go.main.App.GetClientFolderInfo(
      leopardNumber.value,
    );
    const fileName = `Export_Notes_Manquantes_${new Date().getTime()}.pdf`;

    await window.go.main.App.SavePDFToFolder(
      `${folderInfo.path}/notes`,
      fileName,
      pdfBase64,
    );

    alert("✅ Les notes ont été assemblées et exportées avec succès !");
    await verifierToutLeDossier();
  } catch (err) {
    alert("❌ Erreur: " + err.message);
  } finally {
    isGenerating.value.notes = false;
  }
};

const genererSommaire = async () => {
  is;
  Generating.value.sommaire = true;
  try {
    await new Promise((resolve) => setTimeout(resolve, 2000));
    alert("✅ Sommaire généré!");
    await verifierToutLeDossier();
  } finally {
    isGenerating.value.sommaire = false;
  }
};
const ouvrirDocument = async (type) => {
  const chemin = sections.value[type]?.chemin;
  if (chemin) {
    await window.go.main.App.OpenFile(chemin);
  }
};
const verifierToutLeDossier = async () => {
  if (!props.client?.id) return;
  isVerifying.value = true;

  try {
    // 1. Récupération des vraies données
    const notesEnDB = await window.go.main.App.GetClientNotes(props.client.id);
    console.log("Notes trouvées en DB pour ce client:", notesEnDB)
    const folderInfo = await window.go.main.App.GetClientFolderInfo(
      leopardNumber.value,
    );

    const dossierNotes = folderInfo.children?.find((c) => c.name === "notes");
    const nbFichiers = dossierNotes ? dossierNotes.filesCount : 0;
    const totalNotesDB = notesEnDB ? notesEnDB.length : 0;

    // 2. Mise à jour de l'UI avec les VARIABLES et non des chiffres fixes
    sections.value = {
      ...sections.value, // On garde les autres sections
      notesEvolutives: {
        totalNotes: totalNotesDB,
        notesExportees: nbFichiers,
        notesNonExportees: Math.max(0, totalNotesDB - nbFichiers),
        pourcentageExporte:
          totalNotesDB > 0 ? Math.round((nbFichiers / totalNotesDB) * 100) : 0,
      },
      // Tu peux aussi brancher le consentement ici si tu as le fichier
      consentement: {
        existe: folderInfo.children?.some(
          (c) => c.name === "consentement" && c.filesCount > 0,
        ),
        nomFichier: "",
        dateSignature: "",
        chemin: "",
      },
    };
  } catch (err) {
    console.error("Erreur synchro réelle:", err);
  } finally {
    isVerifying.value = false;
  }
};
const closeModal = () => {
  emit("close");
};
watch(
  () => props.isOpen,
  (newVal) => {
    if (newVal && props.client?.id) {
      verifierToutLeDossier();
    }
  },
);
</script>
<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

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

@keyframes pulse {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0.7;
  }
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

::-webkit-scrollbar {
  width: 6px;
}
::-webkit-scrollbar-thumb {
  background: #9333ea;
  border-radius: 10px;
}
</style>
