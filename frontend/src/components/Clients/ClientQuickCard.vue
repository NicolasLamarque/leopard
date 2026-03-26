<template>
  <div
    class="bg-white dark:bg-slate-900 rounded-xl shadow-lg border border-slate-200 dark:border-slate-700 overflow-hidden"
  >
    <!-- Header Compact -->
    <div
      class="bg-gradient-to-r from-slate-700 to-slate-600 dark:from-slate-800 dark:to-slate-700 px-6 py-4"
    >
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div class="relative">
            <div
              class="w-16 h-16 rounded-lg bg-gradient-to-br from-slate-500 to-slate-600 flex items-center justify-center text-white font-bold text-xl shadow-md"
            >
              {{ getInitials() }}
            </div>
            <div v-if="isBirthdayToday" class="absolute -top-1 -right-1">
              <div
                class="w-7 h-7 bg-teal-500 rounded-full flex items-center justify-center shadow-lg"
              >
                <Cake :size="14" class="text-white" />
              </div>
            </div>
          </div>
          <div>
            <h2 class="text-xl font-bold text-white">
              {{ client.prenom }} {{ client.nom }}
            </h2>
            <div class="flex items-center gap-2 mt-1">
              <span class="text-slate-200 text-sm"
                >{{ getSexeBadge() }}{{ age }} ans</span
              >
              <span class="text-slate-400">•</span>
              <span
                :class="getStatusClass()"
                class="text-xs font-semibold px-2 py-0.5 rounded"
              >
                {{ getStatusText() }}
              </span>
            </div>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <button
            v-if="isBirthdayToday"
            class="px-3 py-1.5 bg-teal-500/20 text-teal-200 rounded-lg text-xs font-medium flex items-center gap-1"
          >
            🎂 Anniversaire !
          </button>
          <button
            @click="exportToPDF"
            class="p-2 bg-white/10 hover:bg-white/20 rounded-lg transition-colors"
            title="Exporter en PDF"
          >
            <FileDown :size="18" class="text-white" />
          </button>
        </div>
      </div>
    </div>

    <!-- Numéro Leopard -->
    <div
      class="px-6 py-2 bg-slate-50 dark:bg-slate-800/50 border-b border-slate-200 dark:border-slate-700"
    >
      <div class="flex items-center gap-2 text-sm">
        <Shield :size="14" class="text-slate-400" />
        <span class="text-slate-600 dark:text-slate-400">Dossier:</span>
        <span
          class="font-mono font-semibold text-slate-900 dark:text-slate-100"
        >
          {{ client.no_dossier_leopard || "Non généré" }}
        </span>
      </div>
    </div>

    <!-- Contenu principal -->
    <div class="p-6 space-y-6">
      <!-- ══ PERSONNES CLÉS ══ -->
      <div v-if="personnesCles.length > 0">
        <!-- Titre + toggle vue -->
        <div class="flex items-center justify-between mb-3">
          <h3
            class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2"
          >
            <Users :size="13" />
            Personnes clés
          </h3>
          <div class="flex gap-1">
            <button
              v-for="v in vues"
              :key="v.id"
              @click="vueActive = v.id"
              :title="v.label"
              :class="[
                'p-1.5 rounded-md transition-colors',
                vueActive === v.id
                  ? 'bg-slate-200 dark:bg-slate-700 text-slate-800 dark:text-white'
                  : 'text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800',
              ]"
            >
              <component :is="v.icon" :size="13" />
            </button>
          </div>
        </div>

        <!-- Vue liste -->
        <div v-if="vueActive === 'liste'" class="space-y-2">
          <div
            v-for="p in personnesCles"
            :key="p.contact.id"
            class="cursor-pointer hover:brightness-110 transition-all"
            @click.stop="contactOuvert = p.contact"
            :style="`border-left: 3px solid ${p.couleur}; border-color: var(--tw-border-opacity)`"
          >
            <div
              class="px-4 py-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl"
              :style="`border-left: 3px solid ${p.couleur};`"
            >
              <!-- Rôle -->
              <div
                class="text-xs font-semibold uppercase tracking-wide mb-1.5 flex items-center gap-1.5"
                :style="`color: ${p.couleur};`"
              >
                {{ p.label }}
                <span
                  v-if="p.contact.lien_familial"
                  class="font-normal normal-case text-slate-400"
                >
                  · {{ p.contact.lien_familial }}
                </span>
              </div>

              <!-- Nom -->
              <div
                :class="[
                  'text-sm font-semibold text-slate-900 dark:text-white',
                  p.interdit ? 'line-through opacity-50' : '',
                ]"
              >
                {{ p.contact.prenom }} {{ p.contact.nom }}
              </div>

              <!-- Sous-titre contextuel -->
              <div
                v-if="p.sous"
                class="text-xs text-slate-500 dark:text-slate-400 mt-0.5"
              >
                {{ p.sous }}
              </div>

              <!-- Téléphones cliquables (bloqués si interdit) -->
              <div v-if="!p.interdit" class="flex flex-wrap gap-3 mt-2">
                <a
                  v-if="p.contact.telephone"
                  :href="`tel:${p.contact.telephone}`"
                  class="text-xs text-blue-600 dark:text-blue-400 hover:underline flex items-center gap-1"
                >
                  <Phone :size="11" /> {{ formatPhone(p.contact.telephone) }}
                </a>
                <a
                  v-if="p.contact.cellulaire"
                  :href="`tel:${p.contact.cellulaire}`"
                  class="text-xs text-blue-600 dark:text-blue-400 hover:underline flex items-center gap-1"
                >
                  <Smartphone :size="11" />
                  {{ formatPhone(p.contact.cellulaire) }}
                </a>
              </div>
              <div v-else class="flex gap-3 mt-2 opacity-40">
                <span
                  v-if="p.contact.telephone"
                  class="text-xs text-slate-500 line-through flex items-center gap-1"
                >
                  <Phone :size="11" /> {{ formatPhone(p.contact.telephone) }}
                </span>
              </div>

              <!-- Barre force du lien -->
              <div class="flex items-center gap-2 mt-2">
                <span class="text-xs text-slate-400 whitespace-nowrap"
                  >Lien</span
                >
                <div
                  class="flex-1 h-1.5 bg-slate-200 dark:bg-slate-700 rounded-full overflow-hidden"
                >
                  <div
                    class="h-full rounded-full transition-all"
                    :style="`width:${((p.contact.force_lien + 10) / 20) * 100}%; background:${couleurForceLien(p.contact.force_lien)};`"
                  ></div>
                </div>
                <span class="text-xs text-slate-400"
                  >{{ p.contact.force_lien }}/10</span
                >
              </div>

              <!-- Alerte maltraitance -->
              <div
                v-if="p.contact.signalement_maltraitance === 1"
                class="mt-2 flex items-center gap-1.5 text-xs font-medium px-2 py-1.5 rounded-lg"
                style="background: #fcebeb; color: #a32d2d"
              >
                <AlertTriangle :size="11" /> Signalement maltraitance actif
              </div>

              <!-- Note refus -->
              <div
                v-if="p.interdit && p.contact.note_refus"
                class="mt-2 text-xs px-2 py-1.5 rounded-lg"
                style="background: #faeeda; color: #633806"
              >
                {{ p.contact.note_refus }}
              </div>
            </div>
          </div>
        </div>

        <!-- Vue grille -->
        <div
          v-else
          :class="[
            'grid gap-2',
            vueActive === 'grid3'
              ? 'grid-cols-3'
              : vueActive === 'grid2'
                ? 'grid-cols-2'
                : 'grid-cols-1',
          ]"
        >
          <div
            v-for="p in personnesCles"
            :key="p.contact.id"
            class="rounded-xl border border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 p-3 flex flex-col gap-2 cursor-pointer hover:brightness-110 transition-all"
            :style="`border-top: 3px solid ${p.couleur};`"
            @click.stop="contactOuvert = p.contact"
          >
            <div
              class="text-xs font-semibold uppercase tracking-wide"
              :style="`color:${p.couleur};`"
            >
              {{ p.label }}
            </div>

            <div class="flex items-center gap-2">
              <!-- Avatar -->
              <div
  class="w-8 h-8 rounded-full flex items-center justify-center text-xs font-bold flex-shrink-0"
  :style="`background:${p.couleurLight}; color:${p.couleur};`"
>
  {{ initiales(p.contact) }}
</div>
              <div>
                <div
                  :class="[
                    'text-lg font-semibold text-slate-900 dark:text-white leading-tight',
                    p.interdit ? 'line-through opacity-50' : '',
                  ]"
                >
                  {{ p.contact.prenom }} {{ p.contact.nom }}
                </div>
                <div
                  v-if="p.contact.lien_familial"
                  class="text-xs text-slate-400 leading-tight"
                >
                  {{ p.contact.lien_familial }}
                </div>
              </div>
            </div>

            <a
              v-if="p.contact.telephone && !p.interdit"
              :href="`tel:${p.contact.telephone}`"
              class="text-xs text-blue-600 dark:text-blue-400 hover:underline truncate flex items-center gap-1"
            >
              <Phone :size="10" /> {{ formatPhone(p.contact.telephone) }}
            </a>

            <div class="flex items-center gap-1 mt-auto">
              <div
                class="flex-1 h-1 bg-slate-200 dark:bg-slate-700 rounded-full overflow-hidden"
              >
                <div
                  class="h-full rounded-full"
                  :style="`width:${(p.contact.force_lien / 10) * 100}%; background:${p.couleur};`"
                ></div>
              </div>
              <span class="text-xs text-slate-400"
                >{{ p.contact.force_lien }}/10</span
              >
            </div>

            <div
              v-if="p.contact.signalement_maltraitance === 1"
              class="text-xs font-medium px-1.5 py-1 rounded text-center"
              style="background: #fcebeb; color: #a32d2d"
            >
              ⚠ Maltraitance
            </div>
            <div
              v-if="p.interdit"
              class="text-xs font-medium px-1.5 py-1 rounded text-center"
              style="background: #faeeda; color: #854f0b"
            >
              Contact interdit
            </div>
          </div>
        </div>
      </div>

      <!-- Grid 3 colonnes -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Colonne 1: Médical -->
        <div class="space-y-4">
          <h3
            class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2 mb-3"
          >
            <Heart :size="14" class="text-teal-600 dark:text-teal-400" />
            Médical
          </h3>
          <div v-if="client.medecin_famille_No_Licence || medecinInfo">
            <div class="flex items-start justify-between mb-1">
              <div class="flex items-center gap-2">
                <Stethoscope
                  :size="14"
                  class="text-teal-600 dark:text-teal-400"
                />
                <span class="text-xs text-slate-500 dark:text-slate-400"
                  >Médecin de famille</span
                >
              </div>
              <button
                v-if="medecinInfo"
                @click="viewMedecinDetails"
                class="p-1 hover:bg-slate-100 dark:hover:bg-slate-800 rounded transition-colors"
                title="Voir détails"
              >
                <Eye :size="12" class="text-teal-600 dark:text-teal-400" />
              </button>
            </div>
            <p
              class="font-semibold text-sm text-slate-900 dark:text-slate-100 mb-0.5"
            >
              Dr. {{ getMedecinName() }}
            </p>
            <p
              v-if="client.medecin_famille_No_Licence"
              class="text-xs text-slate-500 dark:text-slate-400 font-mono"
            >
              {{ client.medecin_famille_No_Licence }}
            </p>
          </div>
          <div class="h-px bg-slate-200 dark:bg-slate-700"></div>
          <div v-if="client.numero_assurance_maladie">
            <div class="flex items-center gap-2 mb-1">
              <CreditCard :size="14" class="text-blue-600 dark:text-blue-400" />
              <span class="text-xs text-slate-500 dark:text-slate-400"
                >RAMQ</span
              >
            </div>
            <p
              class="font-mono text-sm font-semibold text-slate-900 dark:text-slate-100"
            >
              {{ client.numero_assurance_maladie }}
            </p>
          </div>
          <div v-if="client.no_hcm || client.no_chaur">
            <div class="h-px bg-slate-200 dark:bg-slate-700 mb-3"></div>
            <div class="flex items-center gap-2 mb-2">
              <FileText
                :size="14"
                class="text-purple-600 dark:text-purple-400"
              />
              <span class="text-xs text-slate-500 dark:text-slate-400"
                >Dossiers cliniques</span
              >
            </div>
            <div class="space-y-1">
              <p
                v-if="client.no_hcm"
                class="text-xs text-slate-700 dark:text-slate-300"
              >
                <span class="font-semibold">HCM:</span> {{ client.no_hcm }}
              </p>
              <p
                v-if="client.no_chaur"
                class="text-xs text-slate-700 dark:text-slate-300"
              >
                <span class="font-semibold">CHAUR:</span> {{ client.no_chaur }}
              </p>
            </div>
          </div>
        </div>

        <!-- Colonne 2: Localisation -->
        <div class="space-y-4">
          <h3
            class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2 mb-3"
          >
            <MapPin :size="14" class="text-orange-600 dark:text-orange-400" />
            Localisation
          </h3>
          <div v-if="client.adresse">
            <div class="flex items-center justify-between mb-1">
              <span class="text-xs text-slate-500 dark:text-slate-400"
                >Domicile</span
              >
              <button
                @click="openGoogleMaps"
                class="text-xs text-blue-600 dark:text-blue-400 hover:underline flex items-center gap-1"
              >
                <Map :size="12" /> Voir sur Maps
              </button>
            </div>
            <p class="text-sm font-medium text-slate-900 dark:text-slate-100">
              {{ client.adresse }}
            </p>
            <div class="flex items-center gap-2 mt-1">
              <p
                class="text-xs text-slate-600 dark:text-slate-400 mt-1 flex items-center gap-2"
              >
                <span>{{ client.ville }}, {{ client.province }}</span>
                <span
                  v-if="paysAlpha2"
                  class="inline-flex w-5 h-3.5 overflow-hidden rounded-sm border border-slate-200 dark:border-slate-700 shadow-sm"
                >
                  <img
                    :src="getFlagUrl(paysAlpha2)"
                    class="w-full h-full object-cover"
                    @error="(e) => (e.target.style.display = 'none')"
                  />
                </span>
                <span>{{ client.pays }} - {{ client.code_postal }}</span>
              </p>
            </div>
          </div>
          <div v-if="getEtablissement()">
            <div class="h-px bg-slate-200 dark:bg-slate-700 mb-3"></div>
            <div class="flex items-center gap-2 mb-1">
              <Building2
                :size="14"
                class="text-orange-600 dark:text-orange-400"
              />
              <span class="text-xs text-slate-500 dark:text-slate-400"
                >Hébergement</span
              >
            </div>
            <p class="text-sm font-semibold text-slate-900 dark:text-slate-100">
              {{ getEtablissement() }}
            </p>
          </div>
        </div>

        <!-- Colonne 3: Coordonnées -->
        <div class="space-y-4">
          <h3
            class="text-xs font-bold text-slate-400 uppercase tracking-wider flex items-center gap-2 mb-3"
          >
            <Phone :size="14" class="text-slate-600 dark:text-slate-400" />
            Coordonnées
          </h3>
          <div class="space-y-3">
            <a
              v-if="client.telephone"
              :href="`tel:${client.telephone}`"
              class="flex items-center gap-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors group"
            >
              <Phone
                :size="14"
                class="text-slate-400 group-hover:text-blue-600 dark:group-hover:text-blue-400"
              />
              <span
                class="text-sm font-medium text-slate-900 dark:text-slate-100"
                >{{ formatPhone(client.telephone) }}</span
              >
            </a>
            <a
              v-if="client.cellulaire"
              :href="`tel:${client.cellulaire}`"
              class="flex items-center gap-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors group"
            >
              <Smartphone
                :size="14"
                class="text-slate-400 group-hover:text-blue-600 dark:group-hover:text-blue-400"
              />
              <span
                class="text-sm font-medium text-slate-900 dark:text-slate-100"
                >{{ formatPhone(client.cellulaire) }}</span
              >
            </a>
            <a
              v-if="client.email"
              :href="`mailto:${client.email}`"
              class="flex items-center gap-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors group"
            >
              <Mail
                :size="14"
                class="text-slate-400 group-hover:text-blue-600 dark:group-hover:text-blue-400"
              />
              <span
                class="text-sm font-medium text-slate-900 dark:text-slate-100 truncate"
                >{{ client.email }}</span
              >
            </a>
          </div>
        </div>
      </div>

      <!-- Note fixe client -->
      <div
        v-if="client.note_fixe"
        class="pt-2 border-t border-slate-200 dark:border-slate-700"
      >
        <div
          class="flex items-start gap-3 p-4 bg-amber-50 dark:bg-amber-950/20 border-l-4 border-amber-500 rounded"
        >
          <AlertTriangle
            :size="16"
            class="text-amber-600 dark:text-amber-400 flex-shrink-0 mt-0.5"
          />
          <div>
            <p
              class="text-xs font-bold text-amber-900 dark:text-amber-100 uppercase tracking-wide mb-1"
            >
              Note importante
            </p>
            <p
              class="text-sm text-amber-800 dark:text-amber-200 whitespace-pre-wrap"
            >
              {{ client.note_fixe }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <div
      class="px-6 py-3 bg-slate-50 dark:bg-slate-800/50 border-t border-slate-200 dark:border-slate-700"
    >
      <div
        class="flex items-center justify-between text-xs text-slate-500 dark:text-slate-400"
      >
        <span>Dernière MAJ: {{ formatDate(client.created_at) }}</span>
        <span v-if="client.date_maj">Modifié: {{ client.date_maj }}</span>
      </div>
    </div>
  </div>
  <ContactForm
    v-if="contactOuvert"
    :client-id="client.id"
    :contact-data="contactOuvert"
    @close="contactOuvert = null"
    @saved="
      contactOuvert = null;
      loadContacts();
    "
  />
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import {
  Shield,
  AlertTriangle,
  Heart,
  MapPin,
  Phone,
  Mail,
  FileDown,
  Cake,
  Eye,
  Stethoscope,
  CreditCard,
  Building2,
  FileText,
  Map,
  Smartphone,
  Users,
  List,
  LayoutGrid,
  Square,
} from "lucide-vue-next";
import ContactForm from "../Contacts/ContactForm.vue";
const props = defineProps({
  client: { type: Object, required: true },
  paysAlpha2: { type: String, default: "" },
});
const emit = defineEmits(["view-medecin", "open-contact"]);
const contactOuvert = ref(null);
const contacts = ref([]);
const medecinInfo = ref(null);
const vueActive = ref("liste");

const vues = [
  { id: "liste", label: "Liste", icon: List },
  { id: "grid3", label: "3 colonnes", icon: LayoutGrid },
  { id: "grid2", label: "2 colonnes", icon: Square },
  { id: "grid1", label: "1 colonne", icon: Square },
];

// ─── Personnes clés — logique sur les FLAGS existants ────────────────────────
// On ne touche pas aux libellés de ref_listes.
// On utilise uniquement les booléens déjà dans ta table.

// ─── Détection par libellé ref_listes (mots-clés probables) ─────────────────
// L'utilisateur définit ses libellés librement — on cherche des mots-clés.
// Si son libellé contient "mandataire" → on enrichit l'affichage.
const motsCles = {
  mandataire: ["mandataire", "mandat"],
  tuteur: ["tuteur", "curatelle", "curateur", "tutelle"],
  conjoint: [
    "conjoint",
    "époux",
    "épouse",
    "conjointe",
    "mari",
    "femme",
    "partenaire",
  ],
  aidant: ["aidant", "proche aidant"],
  urgence: ["urgence"],
  interdit: ["interdit", "refus", "banni"],
};

const contientMotCle = (texte, groupe) =>
  motsCles[groupe]?.some((m) => (texte || "").toLowerCase().includes(m)) ??
  false;

// Label enrichi : si le type_de_contact donne plus d'info que le rôle générique, on l'affiche
const labelEnrichi = (contact, labelDefaut) => {
  const t = contact.type_de_contact || "";
  return t.trim() !== "" ? t : labelDefaut;
};

const ROLES_CONFIG = [
  {
    id: "mandataire",
    label: "Mandataire légal",
    couleur: "#185FA5",
    couleurLight: "#E6F1FB",
    match: (c) =>
      !!c.procuration_notariee ||
      contientMotCle(c.type_de_contact, "mandataire") ||
      contientMotCle(c.relation_type, "mandataire"),
    sous: (c) =>
      [
        c.lien_familial,
        c.relation_type,
        c.procuration_bancaire ? "Proc. bancaire" : null,
      ]
        .filter(Boolean)
        .join(" · "),
    labelFn: (c) => labelEnrichi(c, "Mandataire légal"),
  },
  {
    id: "tuteur",
    label: "Tuteur / Curateur",
    couleur: "#534AB7",
    couleurLight: "#EEEDFE",
    match: (c) =>
      contientMotCle(c.type_de_contact, "tuteur") ||
      contientMotCle(c.relation_type, "tuteur"),
    sous: (c) => [c.lien_familial, c.relation_type].filter(Boolean).join(" · "),
    labelFn: (c) => labelEnrichi(c, "Tuteur / Curateur"),
  },
  {
    id: "conjoint",
    label: "Conjoint(e)",
    couleur: "#D4537E",
    couleurLight: "#FBEAF0",
    match: (c) =>
      contientMotCle(c.lien_familial, "conjoint") ||
      contientMotCle(c.type_de_contact, "conjoint") ||
      contientMotCle(c.relation_type, "conjoint"),
    sous: (c) =>
      [
        c.lien_familial,
        c.aidant_naturel ? "Aidant(e) naturel(le)" : null,
        c.procuration_bancaire ? "Proc. bancaire" : null,
      ]
        .filter(Boolean)
        .join(" · "),
    labelFn: (c) => labelEnrichi(c, "Conjoint(e)"),
  },
  {
    id: "aidant_principal",
    label: "Aidant principal",
    couleur: "#3B6D11",
    couleurLight: "#EAF3DE",
    // Aidant naturel — seuil force_lien abaissé à 3 pour être réaliste
    match: (c) =>
      (!!c.aidant_naturel && c.force_lien >= 3) ||
      (!!c.aidant_naturel && !c.procuration_notariee) ||
      contientMotCle(c.type_de_contact, "aidant"),
    sous: (c) =>
      [
        c.lien_familial,
        c.relation_type && c.relation_type !== c.lien_familial
          ? c.relation_type
          : null,
        `Lien ${c.force_lien > 0 ? "+" : ""}${c.force_lien}`,
      ]
        .filter(Boolean)
        .join(" · "),
    labelFn: (c) => labelEnrichi(c, "Aidant principal"),
  },
  {
    id: "urgence",
    label: "Contact d'urgence",
    couleur: "#E24B4A",
    couleurLight: "#FCEBEB",
    match: (c) =>
      !!c.contact_urgence || contientMotCle(c.type_de_contact, "urgence"),
    sous: (c) => [c.lien_familial, c.relation_type].filter(Boolean).join(" · "),
    labelFn: (c) => "Contact d'urgence",
  },
  {
    id: "bancaire",
    label: "Procuration bancaire",
    couleur: "#0F6E56",
    couleurLight: "#E1F5EE",
    match: (c) => !!c.procuration_bancaire && !c.procuration_notariee,
    sous: (c) => [c.lien_familial, c.relation_type].filter(Boolean).join(" · "),
    labelFn: (c) => "Procuration bancaire",
  },
];

const personnesCles = computed(() => {
  const result = [];
  const dejaVus = new Set();

  // 1. Rôles prioritaires dans l'ordre
  for (const role of ROLES_CONFIG) {
    const contact = contacts.value.find(
      (c) => !dejaVus.has(c.id) && role.match(c),
    );
    if (contact) {
      dejaVus.add(contact.id);
      result.push({
        contact,
        label: role.labelFn ? role.labelFn(contact) : role.label,
        couleur: role.couleur,
        couleurLight: role.couleurLight,
        sous: role.sous(contact),
        interdit: contact.contact_interdit === 1,
      });
    }
  }

  // 2. Contacts interdits non encore affichés
  for (const c of contacts.value) {
    if (!dejaVus.has(c.id) && c.contact_interdit === 1) {
      dejaVus.add(c.id);
      result.push({
        contact: c,
        label: "Contact interdit",
        couleur: "#BA7517",
        couleurLight: "#FAEEDA",
        sous: c.lien_familial || "",
        interdit: true,
      });
    }
  }

  // 3. Signalements maltraitance non encore affichés
  for (const c of contacts.value) {
    if (!dejaVus.has(c.id) && c.signalement_maltraitance === 1) {
      dejaVus.add(c.id);
      result.push({
        contact: c,
        label: "Sous surveillance",
        couleur: "#E24B4A",
        couleurLight: "#FCEBEB",
        sous: c.lien_familial || "",
        interdit: false,
      });
    }
  }

  return result;
});

// ─── Helpers ─────────────────────────────────────────────────────────────────

const initiales = (c) =>
  ((c.prenom?.[0] || "") + (c.nom?.[0] || "")).toUpperCase();

const age = computed(() => {
  if (!props.client.date_naissance) return "?";
  const b = new Date(props.client.date_naissance);
  const t = new Date();
  let a = t.getFullYear() - b.getFullYear();
  if (
    t.getMonth() < b.getMonth() ||
    (t.getMonth() === b.getMonth() && t.getDate() < b.getDate())
  )
    a--;
  return a;
});
// Couleur dynamique force_lien -10 à +10
const couleurForceLien = (v) => {
  if (v < -5) return "#dc2626"; // rouge
  if (v < 0) return "#f97316"; // orange
  if (v === 0) return "#94a3b8"; // gris neutre
  if (v <= 5) return "#0d9488"; // teal
  return "#0f6e56"; // vert foncé
};

// Style du track force_lien
const styleForceLien = (v) => {
  const pct = ((v + 10) / 20) * 100;
  if (v < 0)
    return {
      background: `linear-gradient(to right, #dc2626 0%, #f97316 ${pct * 0.6}%, #94a3b8 ${pct}%, #e2e8f0 ${pct}%, #e2e8f0 100%)`,
    };
  if (v === 0) return { background: "#94a3b8" };
  return {
    background: `linear-gradient(to right, #e2e8f0 0%, #e2e8f0 50%, #0d9488 ${pct}%, #0f6e56 100%)`,
  };
};

const isBirthdayToday = computed(() => {
  if (!props.client.date_naissance) return false;
  const b = new Date(props.client.date_naissance);
  const t = new Date();
  return b.getMonth() === t.getMonth() && b.getDate() === t.getDate();
});

const getInitials = () =>
  `${props.client.prenom[0]}${props.client.nom[0]}`.toUpperCase();
const getSexeBadge = () => ((props.client.sexe || "F") === "M" ? "♂ " : "♀ ");
const getStatusText = () => {
  if (props.client.dcd == 1) return "Décédé";
  if (props.client.actif == 1) return "Actif";
  return "Archivé";
};
const getStatusClass = () => {
  if (props.client.dcd == 1) return "bg-gray-200 text-gray-700";
  if (props.client.actif == 1) return "bg-green-100 text-green-700";
  return "bg-red-100 text-red-700";
};
const getEtablissement = () => {
  if (props.client.rpa_id) return `RPA #${props.client.rpa_id}`;
  if (props.client.chsld_id) return `CHSLD #${props.client.chsld_id}`;
  if (props.client.ri_id) return `RI #${props.client.ri_id}`;
  return null;
};
const getMedecinName = () => {
  if (medecinInfo.value)
    return `${medecinInfo.value.prenom} ${medecinInfo.value.nom}`;
  return props.client.medecin_famille_No_Licence || "Non spécifié";
};
const formatPhone = (phone) => {
  if (!phone) return "";
  const c = phone.replace(/\D/g, "");
  if (c.length === 10)
    return `(${c.slice(0, 3)}) ${c.slice(3, 6)}-${c.slice(6)}`;
  if (c.length === 11 && c[0] === "1")
    return `+1 (${c.slice(1, 4)}) ${c.slice(4, 7)}-${c.slice(7)}`;
  return phone;
};
const formatDate = (d) => (d ? new Date(d).toLocaleDateString("fr-CA") : "N/A");
const getFlagUrl = (alpha2) => {
  if (!alpha2) return null;
  return new URL(
    `../../assets/flags/${alpha2.toLowerCase().trim()}.svg`,
    import.meta.url,
  ).href;
};
const openGoogleMaps = () => {
  const addr = `${props.client.adresse}, ${props.client.ville}, ${props.client.code_postal}`;
  window.open(
    `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(addr)}`,
    "_blank",
  );
};
const viewMedecinDetails = () =>
  emit("view-medecin", props.client.medecin_famille_No_Licence);
const exportToPDF = () => alert("Export PDF à implémenter");

const loadContacts = async () => {
  try {
    contacts.value =
      (await window.go.main.App.GetAllContactsByClientID(props.client.id)) ||
      [];
  } catch (e) {
    console.error("Erreur chargement contacts:", e);
  }
};
const loadMedecinInfo = async () => {
  if (!props.client.medecin_famille_No_Licence) return;
  try {
    medecinInfo.value = await window.go.main.App.GetMedecinByLicence(
      props.client.medecin_famille_No_Licence,
    );
  } catch (e) {
    console.error("Erreur chargement médecin:", e);
  }
};

onMounted(() => {
  loadContacts();
  loadMedecinInfo();
});
</script>

<style scoped>
a {
  transition: all 0.2s ease;
}
</style>
