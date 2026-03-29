<template>
  <!-- Modal plein écran — même pattern que EvaluationView -->
  <Transition name="modal-fade">
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-gray-900/70 backdrop-blur-md" @click="handleClose" />

      <div class="relative bg-white dark:bg-gray-950 w-full max-w-5xl h-[95vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-gray-200 dark:border-gray-800">

        <!-- ── HEADER ─────────────────────────────────────────────── -->
        <div class="px-6 py-4 border-b dark:border-gray-800 flex justify-between items-center bg-gradient-to-r from-slate-800 via-slate-700 to-slate-600 shrink-0">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-white/15 rounded-lg text-white">
              <Users :size="20" />
            </div>
            <div>
              <h2 class="text-base font-bold text-white leading-tight">Annexe — Personnes consultées</h2>
              <p class="text-xs text-slate-300">
                {{ client?.prenom }} {{ client?.nom }}
                <span class="text-teal-400 font-mono ml-1">{{ client?.no_dossier_leopard }}</span>
              </p>
            </div>
          </div>

          <div class="flex items-center gap-2">
            <!-- Compteur personnes -->
            <span class="text-xs text-slate-400 font-mono">
              {{ personnes.length }} personne{{ personnes.length > 1 ? 's' : '' }}
            </span>

            <!-- Export PDF -->
            <button
              @click="handleExportPDF"
              :disabled="isSaving || personnes.length === 0"
              class="flex items-center gap-2 bg-teal-600 hover:bg-teal-500 disabled:opacity-40 text-white px-4 py-2 rounded-xl text-sm font-bold transition-all shadow-lg"
            >
              <Loader2 v-if="isSaving" :size="15" class="animate-spin" />
              <FileDown v-else :size="15" />
              <span class="hidden sm:inline">{{ isSaving ? 'Sauvegarde...' : 'Exporter PDF' }}</span>
            </button>

            <button @click="handleClose" class="p-2 hover:bg-white/10 rounded-full transition-colors">
              <X :size="18" class="text-white/80" />
            </button>
          </div>
        </div>

        <!-- ── CORPS : zone de prévisualisation "papier" ─────────── -->
        <div class="flex-1 overflow-y-auto bg-gray-300 dark:bg-gray-800 p-6">

          <!-- PAGE PAPIER — fond blanc, look formulaire officiel -->
          <div class="bg-white mx-auto shadow-xl" style="width: 215.9mm; min-height: 279.4mm; padding: 18mm 18mm 16mm 18mm; font-family: Arial, sans-serif; font-size: 11px; color: #000;">

            <!-- EN-TÊTE DOCUMENT -->
            <div style="display:flex; justify-content:space-between; align-items:flex-end; border-bottom: 2px solid #000; padding-bottom: 5px; margin-bottom: 10px;">
              <div>
                <div style="font-size:13px; font-weight:bold;">Annexe à l'évaluation psychosociale</div>
                <div style="font-size:11px; color:#333; margin-top:2px;">dans le cadre de l'homologation d'un mandat de protection</div>
              </div>
              <div style="text-align:right; font-size:9px; color:#555; line-height:1.5;">
                096-DGSP-2025-07<br>
                Annexe — Personnes consultées
              </div>
            </div>

            <!-- ENCADRÉ CLIENT -->
            <div style="border:1.5px solid #000; padding:7px 12px; margin-bottom:12px;">
              <div style="font-size:9px; font-weight:bold; text-transform:uppercase; letter-spacing:0.5px; color:#555; margin-bottom:4px;">Personne visée par l'évaluation</div>
              <div style="font-size:13px; font-weight:bold;">{{ clientName || '—' }}</div>
              <div v-if="client?.date_naissance" style="font-size:11px; color:#333; margin-top:2px;">Date de naissance : {{ client.date_naissance }}</div>
              <div v-if="client?.no_dossier_leopard" style="font-size:11px; color:#777; margin-top:1px;">No de dossier : {{ client.no_dossier_leopard }}</div>
            </div>

            <!-- EN-TÊTE SECTION -->
            <div style="background:#d8d8d8; border:1px solid #888; padding:4px 8px; font-size:11px; font-weight:bold; display:flex; justify-content:space-between; align-items:center;">
              <span>C. Personnes consultées dans le cadre de l'évaluation
                <em style="font-weight:normal; font-size:10px;">(suite du formulaire principal)</em>
              </span>
              <span style="font-weight:normal; font-size:9px; color:#555;">Annexe — page 1</span>
            </div>
            <div style="border:1px solid #888; border-top:none; padding: 10px 10px 8px;">

              <!-- BLOCS PERSONNES -->
              <div
                v-for="(personne, idx) in personnes"
                :key="personne.id"
                style="border:0.5px solid #bbb; padding:10px; margin-bottom:10px; position:relative;"
              >
                <!-- Numéro + actions -->
                <div style="display:flex; justify-content:space-between; align-items:center; margin-bottom:8px;">
                  <span style="font-size:9px; font-weight:bold; color:#555; background:#e8e8e8; border:0.5px solid #bbb; padding:1px 6px;">
                    Personne {{ idx + 5 }}
                  </span>
                  <button
                    @click="retirerPersonne(personne.id)"
                    class="flex items-center gap-1 text-xs text-red-400 hover:text-red-600 transition-colors print:hidden"
                    style="font-size:10px;"
                  >
                    <X :size="11" /> Retirer
                  </button>
                </div>

                <!-- Rangée 1 : Nom / Prénom -->
                <div class="form-row-2" style="display:grid; grid-template-columns:1fr 1fr; gap:8px; margin-bottom:8px;">
                  <FieldLine label="Nom" v-model="personne.nom" />
                  <FieldLine label="Prénom" v-model="personne.prenom" />
                </div>

                <!-- Rangée 2 : Lien / Date / Titre -->
                <div style="display:grid; grid-template-columns:2fr 1.4fr 1fr; gap:8px; margin-bottom:8px;">
                  <FieldLine label="Lien avec la personne visée" v-model="personne.lien" />
                  <FieldLine label="Date de consultation (aaaa-mm-jj)" v-model="personne.dateConsult" placeholder="____-__-__" />
                  <FieldLine label="Titre / fonction" v-model="personne.titre" />
                </div>

                <!-- Rangée 3 : Adresse / CP / Tél / Poste -->
                <div style="display:grid; grid-template-columns:2fr 0.8fr 1.2fr 0.5fr; gap:8px; margin-bottom:8px;">
                  <FieldLine label="Adresse (numéro, rue, ville)" v-model="personne.adresse" />
                  <FieldLine label="Code postal" v-model="personne.codePostal" />
                  <FieldLine label="No tél. (travail)" v-model="personne.tel" />
                  <FieldLine label="Poste" v-model="personne.poste" />
                </div>

                <!-- Rangée 4 : Courriel -->
                <div style="display:grid; grid-template-columns:1fr 1fr; gap:8px; margin-bottom:8px;">
                  <FieldLine label="Courriel" v-model="personne.courriel" />
                </div>

                <!-- Cases à cocher -->
                <div style="margin-bottom:8px;">
                  <div style="font-size:8.5px; font-weight:bold; text-transform:uppercase; letter-spacing:0.4px; color:#555; margin-bottom:4px;">Type de consultation</div>
                  <div style="display:flex; gap:16px; flex-wrap:wrap;">
                    <CheckItem v-model="personne.typeTel"     label="Entretien téléphonique" />
                    <CheckItem v-model="personne.typeRencontre" label="Rencontre" />
                    <CheckItem v-model="personne.typeVideo"   label="Téléconsultation" />
                    <CheckItem v-model="personne.typeCourriel" label="Courriel et / ou courrier" />
                  </div>
                </div>

                <!-- Section organisme (déroulable) -->
                <details style="margin-top:8px; border-top:0.5px dashed #bbb; padding-top:6px;">
                  <summary style="font-size:8.5px; font-weight:bold; text-transform:uppercase; color:#777; cursor:pointer; list-style:none; display:flex; align-items:center; gap:5px;">
                    <ChevronRight :size="10" class="details-arrow" />
                    Contact au travail / Organisme (facultatif)
                  </summary>
                  <div style="margin-top:6px;">
                    <div style="display:grid; grid-template-columns:1fr 1fr; gap:8px; margin-bottom:8px;">
                      <FieldLine label="Organisme / employeur" v-model="personne.organisme" />
                      <FieldLine label="Adresse professionnelle" v-model="personne.adresseOrganisme" />
                    </div>
                    <div style="display:grid; grid-template-columns:1fr 1fr 1fr; gap:8px;">
                      <FieldLine label="No tél. (organisme)" v-model="personne.telOrganisme" />
                      <FieldLine label="Courriel (organisme)" v-model="personne.courrielOrganisme" />
                    </div>
                  </div>
                </details>

                <!-- Notes libres -->
                <div style="margin-top:8px;">
                  <div style="font-size:8.5px; font-weight:bold; text-transform:uppercase; letter-spacing:0.4px; color:#555; margin-bottom:3px;">Notes libres</div>
                  <textarea
                    v-model="personne.notes"
                    style="border:0.5px solid #ccc; width:100%; min-height:40px; font-size:10px; font-family:Arial,sans-serif; padding:4px 6px; resize:vertical; background:transparent; outline:none; color:#000;"
                    placeholder="Observations, précisions, contexte de la consultation…"
                  />
                </div>
              </div>

              <!-- Bouton ajouter -->
              <button
                @click="ajouterPersonne"
                style="display:inline-flex; align-items:center; gap:5px; font-size:10px; color:#555; border:0.5px dashed #999; padding:5px 14px; cursor:pointer; background:none; margin-top:6px; font-family:Arial,sans-serif;"
                class="hover:text-blue-600 hover:border-blue-400 transition-colors"
              >
                <Plus :size="12" /> Ajouter une personne consultée
              </button>

            </div><!-- /section-body -->

            <!-- PIED DE PAGE (formulaire) -->
            <div style="margin-top:16px; display:flex; justify-content:space-between; border-top:0.5px solid #ccc; padding-top:5px; font-size:9px; color:#777;">
              <span>Annexe à l'évaluation psychosociale — Homologation du mandat de protection</span>
              <span v-if="client?.no_dossier_leopard">Dossier : {{ client.no_dossier_leopard }}</span>
            </div>

          </div><!-- /page papier -->
        </div><!-- /overflow zone -->

      </div>
    </div>
  </Transition>

  <!-- TOAST -->
  <Transition name="toast-slide">
    <div
      v-if="toast.visible"
      :class="[
        'fixed bottom-6 right-6 z-[70] flex items-center gap-3 px-5 py-4 rounded-xl shadow-2xl border text-sm font-bold max-w-sm',
        toast.type === 'success' ? 'bg-emerald-900 border-emerald-600 text-emerald-100' : 'bg-red-900 border-red-600 text-red-100'
      ]"
    >
      <CheckCircle v-if="toast.type === 'success'" :size="18" class="text-emerald-400 shrink-0" />
      <AlertCircle v-else :size="18" class="text-red-400 shrink-0" />
      <div>
        <p>{{ toast.message }}</p>
        <p v-if="toast.path" class="text-xs opacity-70 mt-0.5 font-mono truncate max-w-xs">{{ toast.path }}</p>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed } from 'vue'
import {
  Users, X, FileDown, Plus, Loader2,
  CheckCircle, AlertCircle, ChevronRight
} from 'lucide-vue-next'
import { saveAnnexePDFToFolder } from './annexePdfGenerator.js'

// ── Props / Emits ─────────────────────────────────────────────────────────
const props = defineProps({
  isOpen:      Boolean,
  client:      Object,   // { id, nom, prenom, date_naissance, no_dossier_leopard }
  nomLeopard:  String,   // identifiant Léopard de l'éval liée (pour nommage)
})

const emit = defineEmits(['close'])

// ── État ──────────────────────────────────────────────────────────────────
const personnes = ref([])
const isSaving  = ref(false)
const toast     = ref({ visible: false, message: '', type: 'success', path: '' })

let compteur = 4 // commence à 5 dans le formulaire (suite du formulaire officiel)

// ── Computed ──────────────────────────────────────────────────────────────
const clientName = computed(() => {
  if (!props.client) return ''
  return `${props.client.prenom || ''} ${props.client.nom || ''}`.trim()
})

// ── Helpers ───────────────────────────────────────────────────────────────
function showToast(message, type = 'success', path = '', duration = 5000) {
  toast.value = { visible: true, message, type, path }
  setTimeout(() => { toast.value.visible = false }, duration)
}

function nouvellePersonne() {
  compteur++
  return {
    id:               compteur,
    nom:              '',
    prenom:           '',
    lien:             '',
    dateConsult:      '',
    titre:            '',
    adresse:          '',
    codePostal:       '',
    tel:              '',
    poste:            '',
    courriel:         '',
    typeTel:          false,
    typeRencontre:    false,
    typeVideo:        false,
    typeCourriel:     false,
    organisme:        '',
    adresseOrganisme: '',
    telOrganisme:     '',
    courrielOrganisme:'',
    notes:            '',
  }
}

// ── Actions ───────────────────────────────────────────────────────────────
function ajouterPersonne() {
  personnes.value.push(nouvellePersonne())
}

function retirerPersonne(id) {
  if (!confirm('Retirer cette personne de la liste ?')) return
  personnes.value = personnes.value.filter(p => p.id !== id)
}

async function handleExportPDF() {
  if (personnes.value.length === 0) {
    showToast('Ajoutez au moins une personne avant d\'exporter', 'error')
    return
  }

  if (!props.client?.no_dossier_leopard) {
    showToast('Numéro Léopard du client introuvable', 'error')
    return
  }

  isSaving.value = true
  try {
    // Préparer les données pour le générateur PDF
    const pdfData = {
      client:      props.client,
      nomLeopard:  props.nomLeopard || '',
      personnes:   personnes.value.map(p => ({
        ...p,
        typesConsultation: [
          p.typeTel       ? 'tel'      : null,
          p.typeRencontre ? 'rencontre': null,
          p.typeVideo     ? 'video'    : null,
          p.typeCourriel  ? 'courriel' : null,
        ].filter(Boolean)
      }))
    }

    const savedPath = await saveAnnexePDFToFolder(
      pdfData,
      props.client.no_dossier_leopard
    )

    showToast('✅ Annexe sauvegardée dans le dossier Evaluations', 'success', savedPath)

  } catch (err) {
    console.error('❌ Erreur export annexe:', err)
    showToast('Erreur lors de l\'export : ' + err.message, 'error')
  } finally {
    isSaving.value = false
  }
}

function handleClose() {
  if (personnes.value.length > 0 && !confirm('Fermer sans sauvegarder ?')) return
  personnes.value = []
  compteur = 4
  emit('close')
}

// Initialiser avec 2 blocs par défaut à l'ouverture
import { watch } from 'vue'
watch(() => props.isOpen, (val) => {
  if (val && personnes.value.length === 0) {
    ajouterPersonne()
    ajouterPersonne()
  }
})
</script>

<!-- ── Sous-composants inline ──────────────────────────────────────────── -->
<script>
// FieldLine — champ de saisie style "formulaire papier"
export const FieldLine = {
  props: { label: String, modelValue: String, placeholder: String },
  emits: ['update:modelValue'],
  template: `
    <div>
      <div style="font-size:8.5px; font-weight:bold; text-transform:uppercase; letter-spacing:0.4px; color:#555; margin-bottom:2px;">{{ label }}</div>
      <input
        type="text"
        :value="modelValue"
        :placeholder="placeholder || ''"
        @input="$emit('update:modelValue', $event.target.value)"
        style="border:none; border-bottom:1px solid #555; width:100%; font-size:11px; font-family:Arial,sans-serif; padding:1px 2px; background:transparent; color:#000; outline:none;"
      >
    </div>
  `
}

// CheckItem — case à cocher style papier
export const CheckCheckItem = {
  props: { label: String, modelValue: Boolean },
  emits: ['update:modelValue'],
  template: `
    <label style="display:flex; align-items:center; gap:5px; font-size:10px; cursor:pointer; user-select:none;">
      <input
        type="checkbox"
        :checked="modelValue"
        @change="$emit('update:modelValue', $event.target.checked)"
        style="width:11px; height:11px; accent-color:#1a6bb5; cursor:pointer;"
      >
      {{ label }}
    </label>
  `
}
</script>

<style scoped>
.modal-fade-enter-active, .modal-fade-leave-active { transition: opacity 0.25s ease; }
.modal-fade-enter-from, .modal-fade-leave-to { opacity: 0; }

.toast-slide-enter-active, .toast-slide-leave-active { transition: all 0.3s ease; }
.toast-slide-enter-from { opacity: 0; transform: translateX(2rem); }
.toast-slide-leave-to   { opacity: 0; transform: translateX(2rem); }

.animate-spin { animation: spin 1s linear infinite; }
@keyframes spin { from { transform: rotate(0deg); } to { transform: rotate(360deg); } }

details summary::-webkit-details-marker { display: none; }
details[open] .details-arrow { transform: rotate(90deg); }
.details-arrow { transition: transform 0.15s; }
</style>