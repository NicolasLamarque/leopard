<template>
  <div class="space-y-4">

    <!-- ══ BARRE D'ACTIONS ══ -->
    <div class="bg-white dark:bg-gray-900 rounded-xl border border-gray-200 dark:border-gray-800 px-4 py-2.5 flex flex-wrap items-center gap-2">
      <button
        @click="showContacts = !showContacts"
        :class="showContacts
          ? 'bg-blue-600 text-white'
          : 'bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700'"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium transition-all"
      >
        <UsersIcon :size="14" />
        Contacts
        <span
          v-if="contactsCount > 0"
          class="px-1.5 py-0.5 rounded-full text-[10px] font-bold"
          :class="showContacts ? 'bg-white/25 text-white' : 'bg-blue-100 dark:bg-blue-900/40 text-blue-700 dark:text-blue-300'"
        >{{ contactsCount }}</span>
      </button>

      <button
        @click="$emit('show-notes')"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all"
      >
        <FileText :size="14" />
        Notes cliniques
      </button>

      <button
        v-if="folderExists"
        @click="handleOpenFolder"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all"
      >
        <FolderOpen :size="14" />
        Dossier fichiers
      </button>

      <button
        @click="showEvaluations = true"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all"
      >
        <ShieldCheck :size="14" />
        Évaluations
      </button>

      <div class="ml-auto">
        <span :class="statusBadgeClass" class="text-xs">{{ statusText }}</span>
      </div>
    </div>

    <!-- Widget dossier -->
    <ClientFolderWidget
      v-if="formData.id"
      :leopard-number="leopardNumber"
      :client-name="`${formData.nom} ${formData.prenom}`"
      :folder-exists="folderExists"
      :folder-info="folderInfo"
      @folder-created="onFolderCreated"
      @folder-opened="onFolderOpened"
      @refresh="checkFolderExists"
    />

    <!-- Fiche rapide -->
    <ClientQuickCard
      v-if="formData.id"
      :client="formData"
      @view-medecin="viewMedecinDetails"
      @open-contact="ouvrirContactDepuisCard"
    />

    <!-- ══ FORMULAIRE PRINCIPAL ══ -->
    <div class="bg-white dark:bg-gray-900 rounded-xl border border-gray-200 dark:border-gray-800 overflow-hidden">
      <div class="flex" style="min-height: 520px">

        <!-- Navigation latérale -->
        <div class="w-44 shrink-0 border-r border-gray-100 dark:border-gray-800 bg-gray-50 dark:bg-gray-900/60 py-3 flex flex-col">
          <nav class="flex flex-col gap-0.5 px-2 flex-1">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="activeTab = tab.id"
              :class="activeTab === tab.id
                ? 'bg-teal-700 text-white shadow-sm'
                : 'text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800 hover:text-gray-700 dark:hover:text-gray-200'"
              class="flex items-center gap-2 px-3 py-2 rounded-lg text-xs font-medium transition-all text-left w-full whitespace-nowrap"
            >
              <component :is="tab.icon" :size="14" class="shrink-0" />
              {{ tab.label }}
            </button>
          </nav>
          <div class="px-2 pt-3 border-t border-gray-100 dark:border-gray-800 mt-2">
            <button
              @click="submitHandler"
              type="button"
              class="flex items-center justify-center gap-1.5 w-full px-3 py-2 rounded-lg text-xs font-semibold bg-teal-700 hover:bg-teal-800 text-white shadow-sm transition-all"
            >
              <Save :size="13" />
              Sauvegarder
            </button>
          </div>
        </div>

        <!-- Contenu onglets -->
        <div class="flex-1 overflow-auto p-5">
          <form @submit.prevent="submitHandler">

            <!-- ══ IDENTITÉ ══ -->
            <div v-show="activeTab === 'identity'">
              <p class="section-title">Identité du client</p>
              <div class="grid grid-cols-3 gap-x-4 gap-y-4">

                <div class="col-span-2">
                  <FormKit type="text" v-model="formData.nom" label="Nom *" validation="required" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.prenom" label="Prénom *" validation="required" :classes="fkClasses" />
                </div>

                <!-- Date de naissance — input natif stylé -->
                <div>
                  <label class="fk-label">Date de naissance</label>
                  <div class="fk-inner">
                    <CalendarDays :size="13" class="text-gray-400 shrink-0 ml-2.5" />
                    <input
                      type="date"
                      v-model="formData.date_naissance"
                      class="fk-input"
                    />
                  </div>
                </div>

                <div>
                  <FormKit
                    type="select"
                    v-model="formData.sexe"
                    label="Sexe"
                    :options="optionsSexe"
                    :classes="fkClasses"
                  />
                </div>
                <div>
                  <FormKit
                    type="select"
                    v-model="formData.statut_marital"
                    label="Statut marital"
                    :options="optionsStatutMarital"
                    :classes="fkClasses"
                  />
                </div>

                <div>
                  <FormKit type="text" v-model="formData.lieu_naissance" label="Lieu de naissance" placeholder="Ville, Pays" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.identite_genre" label="Identité de genre" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.orientation_sexuelle" label="Orientation sexuelle" :classes="fkClasses" />
                </div>

                <!-- N° Léopard -->
                <div class="col-span-3">
                  <label class="fk-label">N° Dossier Léopard</label>
                  <div class="flex items-center gap-2 px-3 py-2 bg-teal-50 dark:bg-teal-900/20 border border-teal-200 dark:border-teal-800 rounded-md">
                    <Shield :size="13" class="text-teal-600 dark:text-teal-400 shrink-0" />
                    <span class="font-mono text-xs font-semibold text-teal-700 dark:text-teal-300">
                      {{ leopardNumber || "En attente de génération" }}
                    </span>
                  </div>
                </div>

              </div>
            </div>

            <!-- ══ COORDONNÉES ══ -->
            <div v-show="activeTab === 'contact'">
              <p class="section-title">Coordonnées</p>
              <div class="grid grid-cols-3 gap-x-4 gap-y-4">

                <!-- Téléphone — computed getter/setter pour éviter le bug curseur -->
                <div>
                  <label class="fk-label">Téléphone</label>
                  <div class="fk-inner">
                    <Phone :size="13" class="text-gray-400 shrink-0 ml-2.5" />
                    <input
                      type="tel"
                      :value="formatPhone(formData.telephone)"
                      @change="formData.telephone = $event.target.value.replace(/\D/g, '')"
                      placeholder="(819) 555-0123"
                      class="fk-input"
                    />
                  </div>
                </div>

                <!-- Cellulaire -->
                <div>
                  <label class="fk-label">Cellulaire</label>
                  <div class="fk-inner">
                    <Smartphone :size="13" class="text-gray-400 shrink-0 ml-2.5" />
                    <input
                      type="tel"
                      :value="formatPhone(formData.cellulaire)"
                      @change="formData.cellulaire = $event.target.value.replace(/\D/g, '')"
                      placeholder="(819) 555-0456"
                      class="fk-input"
                    />
                  </div>
                </div>

                <!-- Courriel -->
                <div>
                  <label class="fk-label">Courriel</label>
                  <div class="fk-inner">
                    <Mail :size="13" class="text-gray-400 shrink-0 ml-2.5" />
                    <input
                      type="email"
                      v-model="formData.email"
                      placeholder="exemple@email.com"
                      class="fk-input"
                    />
                  </div>
                </div>

              </div>

              <!-- Appel rapide -->
              <div
                v-if="formData.telephone || formData.cellulaire || formData.email"
                class="mt-4 flex flex-wrap gap-2 p-3 bg-gray-50 dark:bg-gray-800/60 rounded-lg border border-gray-200 dark:border-gray-700"
              >
                <span class="text-[10px] font-bold uppercase tracking-wider text-gray-400 self-center mr-1">Appel rapide</span>
                <a v-if="formData.telephone" :href="`tel:${formData.telephone}`"
                  class="quick-link hover:border-teal-400 hover:text-teal-700 dark:hover:text-teal-400">
                  <Phone :size="11" /> {{ formatPhone(formData.telephone) }}
                </a>
                <a v-if="formData.cellulaire" :href="`tel:${formData.cellulaire}`"
                  class="quick-link hover:border-blue-400 hover:text-blue-700 dark:hover:text-blue-400">
                  <Smartphone :size="11" /> {{ formatPhone(formData.cellulaire) }}
                </a>
                <a v-if="formData.email" :href="`mailto:${formData.email}`"
                  class="quick-link hover:border-violet-400 hover:text-violet-700 dark:hover:text-violet-400">
                  <Mail :size="11" /> {{ formData.email }}
                </a>
              </div>
            </div>

            <!-- ══ ADRESSE ══ -->
            <div v-show="activeTab === 'address'">
              <p class="section-title">Adresse</p>
              <div class="grid grid-cols-3 gap-x-4 gap-y-4">

                <!-- Adresse + lien Maps -->
                <div class="col-span-2">
                  <div class="flex items-center justify-between mb-1">
                    <label class="fk-label mb-0">Adresse</label>
                    <button
                      v-if="formData.adresse"
                      @click.prevent="openGoogleMaps"
                      type="button"
                      class="flex items-center gap-1 text-[10px] text-blue-500 dark:text-blue-400 hover:underline font-medium"
                    >
                      <MapPin :size="10" /> Google Maps
                    </button>
                  </div>
                  <FormKit type="text" v-model="formData.adresse" placeholder="123 Rue Principale" :classes="fkClasses" />
                </div>

                <div>
                  <FormKit type="text" v-model="formData.appartement" label="Appartement" placeholder="Apt 101" :classes="fkClasses" />
                </div>

                <!-- Municipalité col-span-2, Code postal col-span-1 -->
                <div class="col-span-2">
                  <MunSelector
                    v-model="formData.mcode"
                    v-model:arrValue="formData.arrcod"
                    @change="({ mun, arrondissement }) => {
                      if (mun) formData.ville = mun.munnom
                      if (arrondissement) formData.arrcod = arrondissement.arrcod
                    }"
                  />
                </div>

                <div>
                  <FormKit type="text" v-model="formData.code_postal" label="Code postal" placeholder="J8V 1A1" :classes="fkClasses" />
                </div>

                <!-- Province + Pays sur même ligne -->
                <div>
                  <FormKit
                    type="select"
                    v-model="formData.province"
                    label="Province"
                    :options="optionsProvince"
                    :classes="fkClasses"
                  />
                </div>

                <div class="col-span-2">
                  <PaysSelect v-model="formData.pays" />
                </div>

              </div>
            </div>

            <!-- ══ SOCIO-CULTUREL ══ -->
            <div v-show="activeTab === 'sociocultural'">
              <p class="section-title">Profil socio-culturel</p>
              <div class="grid grid-cols-3 gap-x-4 gap-y-4">

                <div>
                  <FormKit type="text" v-model="formData.occupation" label="Occupation" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.profession" label="Profession" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.employeur" label="Employeur" :classes="fkClasses" />
                </div>

                <div>
                  <FormKit
                    type="select"
                    v-model="formData.niveau_scolaire"
                    label="Niveau scolaire"
                    :options="optionsNiveauScolaire"
                    :classes="fkClasses"
                  />
                </div>
                <div>
                  <FormKit
                    type="select"
                    v-model="formData.langue_preferee"
                    label="Langue préférée"
                    :options="optionsLangue"
                    :classes="fkClasses"
                  />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.premiere_langue" label="Première langue" :classes="fkClasses" />
                </div>

                <div>
                  <FormKit type="text" v-model="formData.origine_ethnique" label="Origine ethnique" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.religion" label="Religion" :classes="fkClasses" />
                </div>

                <div class="col-span-3 mt-1">
                  <label class="flex items-center gap-2.5 cursor-pointer group w-fit">
                    <input
                      type="checkbox"
                      v-model="formData.premiere_nation"
                      :true-value="1"
                      :false-value="0"
                      class="w-3.5 h-3.5 text-teal-600 border-gray-300 rounded focus:ring-teal-500"
                    />
                    <span class="text-xs font-medium text-gray-600 dark:text-gray-400 group-hover:text-gray-800 dark:group-hover:text-gray-200 transition-colors">
                      Personne des Premières Nations
                    </span>
                  </label>
                </div>

              </div>
            </div>

            <!-- ══ MÉDICAL ══ -->
            <div v-show="activeTab === 'medical'">
              <p class="section-title">Dossier médical</p>
              <div class="grid grid-cols-3 gap-x-4 gap-y-4 mb-5">

                <div>
                  <FormKit type="text" v-model="formData.numero_assurance_maladie" label="RAMQ" placeholder="TREJ 4503 1501" :classes="fkMonoClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.numero_securite_sociale" label="NAS" placeholder="123-456-789" :classes="fkMonoClasses" />
                </div>
                <div class="grid grid-cols-2 gap-3">
                  <FormKit type="text" v-model="formData.no_hcm" label="HCM" :classes="fkMonoClasses" />
                  <FormKit type="text" v-model="formData.no_chaur" label="CHAUR" :classes="fkMonoClasses" />
                </div>

              </div>

              <div class="section-divider mb-3">Médecin de famille</div>

              <MedecinSelector
                v-model="formData.medecin_famille_No_Licence"
                @medecin-selected="handleMedecinSelected"
                @view-details="openMedecinModal"
              />
              <MedecinDetailsModal
                v-if="showMedecinModal && selectedMedecinForModal"
                :medecin="selectedMedecinForModal"
                @close="showMedecinModal = false"
                @edit="editMedecin"
                @delete="deleteMedecin"
              />
            </div>

            <!-- ══ ÉTABLISSEMENTS ══ -->
            <div v-show="activeTab === 'establishments'">
              <p class="section-title">Établissements & intervenants</p>

              <div class="section-divider mb-3">Milieu de vie</div>

              <div class="grid grid-cols-3 gap-x-4 gap-y-4 mb-5">
                <div class="col-span-2">
                  <RPAselector v-model="formData.rpa_id" />
                </div>
                <div>
                  <!-- text au lieu de number pour éviter les bugs de valeur -->
                  <FormKit type="text" v-model="formData.chsld_id" label="CHSLD (ID)" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.ri_id" label="RI (ID)" :classes="fkClasses" />
                </div>
                <div>
                  <FormKit type="text" v-model="formData.pivot_id" label="PIVOT (ID)" :classes="fkClasses" />
                </div>
              </div>

              <div class="section-divider mb-3">Notaire & procurations</div>

              <div class="mb-4">
                <NotaireSelector
                  v-model="formData.notaire_id"
                  @notaire-selected="handleNotaireSelected"
                  @view-details="viewNotaireDetails"
                />
              </div>

              <div class="flex flex-wrap gap-4 mt-2">
                <label class="flex items-center gap-2.5 cursor-pointer group">
                  <input
                    type="checkbox"
                    v-model="formData.procuration_bancaire"
                    :true-value="1"
                    :false-value="0"
                    class="w-3.5 h-3.5 text-teal-600 border-gray-300 rounded focus:ring-teal-500"
                  />
                  <span class="text-xs font-medium text-gray-600 dark:text-gray-400 group-hover:text-gray-800 dark:group-hover:text-gray-200 transition-colors">
                    Procuration bancaire
                  </span>
                </label>
                <label class="flex items-center gap-2.5 cursor-pointer group">
                  <input
                    type="checkbox"
                    v-model="formData.procuration_notariee"
                    :true-value="1"
                    :false-value="0"
                    class="w-3.5 h-3.5 text-teal-600 border-gray-300 rounded focus:ring-teal-500"
                  />
                  <span class="text-xs font-medium text-gray-600 dark:text-gray-400 group-hover:text-gray-800 dark:group-hover:text-gray-200 transition-colors">
                    Procuration notariée
                  </span>
                </label>
              </div>
            </div>

            <!-- ══ NOTES & STATUT ══ -->
            <div v-show="activeTab === 'notes'">
              <p class="section-title">Notes & statut du dossier</p>

              <div class="mb-5">
                <FormKit
                  type="textarea"
                  v-model="formData.note_fixe"
                  label="Note fixe (visible sur la fiche)"
                  placeholder="Notes importantes : allergies, directives particulières, comportements…"
                  :rows="4"
                  :classes="fkClasses"
                />
                <div
                  v-if="formData.note_fixe"
                  class="mt-2 flex items-start gap-2 px-3 py-2.5 rounded-lg bg-amber-50 dark:bg-amber-900/20 border-l-2 border-amber-400"
                >
                  <Eye :size="12" class="text-amber-500 shrink-0 mt-0.5" />
                  <p class="text-xs text-amber-700 dark:text-amber-300 leading-relaxed">{{ formData.note_fixe }}</p>
                </div>
              </div>

              <div class="section-divider mb-4">Statut du dossier</div>

              <div class="grid grid-cols-2 gap-3 mb-4">
                <label
                  class="flex items-center gap-3 px-3 py-3 border rounded-lg cursor-pointer transition-all"
                  :class="formData.actif === 1
                    ? 'border-emerald-400 bg-emerald-50 dark:bg-emerald-900/10 dark:border-emerald-700'
                    : 'border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800'"
                >
                  <input
                    v-model="formData.actif"
                    type="checkbox"
                    :true-value="1"
                    :false-value="0"
                    class="w-3.5 h-3.5 text-emerald-600 border-gray-300 rounded shrink-0"
                  />
                  <div>
                    <p class="text-xs font-semibold text-gray-700 dark:text-gray-300">Dossier actif</p>
                    <p class="text-[10px] text-gray-400 mt-0.5">Suivi en cours</p>
                  </div>
                </label>

                <label
                  class="flex items-center gap-3 px-3 py-3 border rounded-lg cursor-pointer transition-all"
                  :class="formData.dcd === 1
                    ? 'border-gray-500 bg-gray-100 dark:bg-gray-800 dark:border-gray-600'
                    : 'border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800'"
                >
                  <input
                    v-model="formData.dcd"
                    type="checkbox"
                    :true-value="1"
                    :false-value="0"
                    class="w-3.5 h-3.5 text-gray-600 border-gray-300 rounded shrink-0"
                  />
                  <div>
                    <p class="text-xs font-semibold text-gray-700 dark:text-gray-300">Décédé (DCD)</p>
                    <p class="text-[10px] text-gray-400 mt-0.5">Ferme le dossier</p>
                  </div>
                </label>
              </div>

              <div v-if="formData.dcd === 1">
                <label class="fk-label">Date de décès</label>
                <div class="fk-inner">
                  <CalendarDays :size="13" class="text-gray-400 shrink-0 ml-2.5" />
                  <input type="date" v-model="formData.date_deces" class="fk-input" />
                </div>
              </div>
            </div>

          </form>
        </div>
      </div>
    </div>

    <!-- Évaluations -->
    <EvaluationView
      :is-open="showEvaluations"
      :client="formData"
      @close="showEvaluations = false"
    />

    <!-- Sidebar Contacts -->
    <Transition name="slide">
      <div
        v-if="showContacts"
        class="fixed inset-y-0 right-0 w-full md:w-[500px] bg-white dark:bg-gray-900 shadow-2xl z-50 overflow-y-auto border-l border-gray-200 dark:border-gray-800"
      >
        <ContactsSidebar
          :client-id="formData.id"
          @close="showContacts = false"
          @contact-added="contactsCount++"
          @contact-deleted="contactsCount > 0 && contactsCount--"
        />
      </div>
    </Transition>
    <Transition name="fade">
      <div v-if="showContacts" @click="showContacts = false" class="fixed inset-0 bg-black/40 backdrop-blur-sm z-40" />
    </Transition>

  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from "vue";
import {
  User, Phone, Mail, MapPin, Heart, Building, ClipboardList,
  FileText, UsersIcon, FolderOpen, Shield, Eye, Save,
  Smartphone, Globe, ShieldCheck, CalendarDays,
} from "lucide-vue-next";
import ContactsSidebar from "../../components/Contacts/ContactSidebar.vue";
import ClientQuickCard from "../Clients/ClientQuickCard.vue";
import MedecinSelector from "@/components/Medecins/MedecinSelector.vue";
import MedecinDetailsModal from "@/components/Medecins/MedecinDetailsModal.vue";
import {
  generateLeopardNumber,
  createClientFolder,
  openClientFolder,
  clientFolderExists,
} from "@/utils/clientFolderManager";
import ClientFolderWidget from "./ClientFolderWidget.vue";
import RPAselector from "../RPA/RPAselector.vue";
import NotaireSelector from "@/components/Notaires/NotaireSelector.vue";
import EvaluationView from "../evaluation/EvaluationView.vue";
import PaysSelect from "../Selectors/PaysSelector.vue";
import MunSelector from "../Selectors/MunSelector.vue";

// ── formatPhone ──────────────────────────────────────────────────
const formatPhone = (phone) => {
  if (!phone) return "";
  const cleaned = String(phone).replace(/\D/g, "");
  if (cleaned.length === 10)
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`;
  if (cleaned.length === 11 && cleaned[0] === "1")
    return `+1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`;
  return phone;
};

// ── Options listes déroulantes ───────────────────────────────────
const optionsSexe = [
  { label: "—", value: "" },
  { label: "Masculin", value: "M" },
  { label: "Féminin", value: "F" },
  { label: "Non-binaire", value: "X" },
];

const optionsStatutMarital = [
  { label: "—", value: "" },
  { label: "Célibataire", value: "Célibataire" },
  { label: "Marié(e)", value: "Marié(e)" },
  { label: "Conjoint(e) de fait", value: "Conjoint(e) de fait" },
  { label: "Divorcé(e)", value: "Divorcé(e)" },
  { label: "Veuf(ve)", value: "Veuf(ve)" },
  { label: "Séparé(e)", value: "Séparé(e)" },
];

const optionsNiveauScolaire = [
  { label: "—", value: "" },
  { label: "Primaire", value: "Primaire" },
  { label: "Secondaire", value: "Secondaire" },
  { label: "Collégial", value: "Collégial" },
  { label: "Universitaire", value: "Universitaire" },
  { label: "Autre", value: "Autre" },
];

const optionsLangue = [
  { label: "—", value: "" },
  { label: "Français", value: "Français" },
  { label: "Anglais", value: "Anglais" },
  { label: "Autre", value: "Autre" },
];

const optionsProvince = [
  { label: "—", value: "" },
  { label: "Québec", value: "QC" },
  { label: "Ontario", value: "ON" },
  { label: "Colombie-Britannique", value: "BC" },
  { label: "Alberta", value: "AB" },
  { label: "Manitoba", value: "MB" },
  { label: "Saskatchewan", value: "SK" },
  { label: "Nouvelle-Écosse", value: "NS" },
  { label: "Nouveau-Brunswick", value: "NB" },
  { label: "Île-du-Prince-Édouard", value: "PE" },
  { label: "Terre-Neuve-et-Labrador", value: "NL" },
];

// ── Classes FormKit uniformes ────────────────────────────────────
const fkClasses = {
  outer: "mb-0",
  label: "fk-label",
  inner: "fk-inner",
  input: "fk-input",
  select: "fk-input",
  message: "text-[10px] text-red-500 mt-0.5",
};

const fkMonoClasses = {
  ...fkClasses,
  input: "fk-input font-mono tracking-wider",
};

// ── Props / Emits ────────────────────────────────────────────────
const props = defineProps({
  clientData: { type: Object, required: true },
});
const emit = defineEmits(["save", "folderCreated", "show-notes", "show-evaluations"]);

// ── État ─────────────────────────────────────────────────────────
const formData = ref({ ...props.clientData });
const leopardNumber = ref("");
const folderExists = ref(false);
const folderInfo = ref(null);
const showContacts = ref(false);
const contactsCount = ref(0);
const activeTab = ref("identity");
const showEvaluations = ref(false);
const showMedecinModal = ref(false);
const selectedMedecinForModal = ref(null);

// ── Onglets ──────────────────────────────────────────────────────
const tabs = [
  { id: "identity", label: "Identité", icon: User },
  { id: "contact", label: "Coordonnées", icon: Phone },
  { id: "address", label: "Adresse", icon: MapPin },
  { id: "sociocultural", label: "Socio-culturel", icon: Globe },
  { id: "medical", label: "Médical", icon: Heart },
  { id: "establishments", label: "Établissements", icon: Building },
  { id: "notes", label: "Notes & Statut", icon: ClipboardList },
];

// ── Statut calculé depuis formData ───────────────────────────────
const statusBadgeClass = computed(() => {
  const base = "px-3 py-1 rounded-full text-xs font-medium";
  if (formData.value.dcd === 1)
    return `${base} bg-gray-200 dark:bg-gray-700 text-gray-600 dark:text-gray-400`;
  if (formData.value.actif === 1)
    return `${base} bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300`;
  return `${base} bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300`;
});

const statusText = computed(() => {
  if (formData.value.dcd === 1) return "Décédé";
  if (formData.value.actif === 1) return "Actif";
  return "Inactif";
});

// ── Actions ──────────────────────────────────────────────────────
const openGoogleMaps = () => {
  const address = `${formData.value.adresse}, ${formData.value.ville}, ${formData.value.code_postal}`;
  window.open(`https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(address)}`, "_blank");
};

const viewMedecinDetails = () => { showMedecinModal.value = true; };
const openMedecinModal = (medecin) => {
  selectedMedecinForModal.value = medecin;
  showMedecinModal.value = true;
};
const handleMedecinSelected = (medecin) => { console.log("Médecin:", medecin); };
const editMedecin = (medecin) => {
  console.log("Éditer:", medecin);
  showMedecinModal.value = false;
};
const deleteMedecin = async (medecinId) => {
  try {
    await window.go.main.App.DeleteMedecin(medecinId);
    showMedecinModal.value = false;
    if (formData.value.medecin_famille_No_Licence === selectedMedecinForModal.value?.licence) {
      formData.value.medecin_famille_No_Licence = "";
    }
  } catch (e) { console.error(e); }
};

const handleNotaireSelected = (n) => { console.log("Notaire:", n); };
const viewNotaireDetails = (n) => { console.log("Voir notaire:", n); };

const onFolderCreated = (data) => {
  folderExists.value = true;
  leopardNumber.value = data.leopardNumber;
  formData.value.no_dossier_leopard = data.leopardNumber;
  loadFolderInfo();
  emit("folderCreated", data);
};
const onFolderOpened = () => { console.log("Dossier ouvert"); };
const handleOpenFolder = () => { console.log("Ouvrir:", leopardNumber.value); };

const checkFolderExists = async () => {
  try {
    folderExists.value = await window.go.main.App.ClientFolderExists(leopardNumber.value);
  } catch {}
};
const contactDepuisCard = ref(null)

const ouvrirContactDepuisCard = (contact) => {
  contactDepuisCard.value = contact
  showContacts.value = true  // ouvre la sidebar
}

const loadFolderInfo = async () => {
  try {
    const info = await window.go.main.App.GetClientFolderInfo(leopardNumber.value);
    if (info) folderInfo.value = info;
  } catch {}
};

const submitHandler = () => emit("save", formData.value);

// ── Watchers ─────────────────────────────────────────────────────
watch(
  () => props.clientData,
  (newVal) => {
    formData.value = { ...newVal };
    leopardNumber.value = newVal.no_dossier_leopard || "";
  },
  { deep: true },
);

// ── Montage ──────────────────────────────────────────────────────
onMounted(async () => {
  leopardNumber.value = formData.value.no_dossier_leopard || "";
  await checkFolderExists();
  await loadFolderInfo();
});
</script>

<style scoped>
@reference "tailwindcss";

/* ── Classes utilitaires uniformes pour les champs ── */
.fk-label {
  @apply block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide;
}

.fk-inner {
  @apply flex items-center gap-0 border border-gray-300 dark:border-gray-700 rounded-md overflow-hidden
         focus-within:border-teal-500 focus-within:ring-1 focus-within:ring-teal-500/30
         transition-all bg-white dark:bg-gray-800;
}

.fk-input {
  @apply w-full px-2.5 py-1.5 text-sm text-gray-900 dark:text-gray-100
         bg-transparent outline-none placeholder-gray-400 dark:placeholder-gray-600;
}

/* Section title */
.section-title {
  @apply text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500
         mb-4 pb-2 border-b border-gray-100 dark:border-gray-800;
}

/* Section divider (ligne avec texte centré) */
.section-divider {
  @apply flex items-center gap-3 text-[10px] font-bold uppercase tracking-wider text-gray-400;
}
.section-divider::before,
.section-divider::after {
  content: '';
  @apply flex-1 h-px bg-gray-100 dark:bg-gray-800;
}

/* Liens appel rapide */
.quick-link {
  @apply flex items-center gap-1.5 px-2.5 py-1 rounded-lg text-xs font-medium
         bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700
         text-gray-700 dark:text-gray-300 transition-all;
}

/* Transitions sidebar */
.slide-enter-active,
.slide-leave-active { transition: transform 0.28s ease; }
.slide-enter-from,
.slide-leave-to { transform: translateX(100%); }

.fade-enter-active,
.fade-leave-active { transition: opacity 0.25s ease; }
.fade-enter-from,
.fade-leave-to { opacity: 0; }
</style>