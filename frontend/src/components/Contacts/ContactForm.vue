<template>
  <div class="cf-overlay">
    <div class="cf-modal animate-slideIn">

      <!-- ═══ HEADER ═══ -->
      <div class="cf-header">
        <div class="flex items-center gap-3">
          <div class="cf-header-icon">
            <UserPlus :size="16" />
          </div>
          <div>
            <h2 class="cf-title">{{ isEdit ? 'Modifier le contact' : 'Nouveau contact' }}</h2>
            <p class="cf-subtitle">{{ isEdit ? 'Mise à jour du dossier' : 'Ajout au réseau du client' }}</p>
          </div>
        </div>
        <button @click="$emit('close')" class="cf-close" :disabled="loading">
          <X :size="16" />
        </button>
      </div>

      <!-- ═══ BODY ═══ -->
      <div class="cf-body">
        <form @submit.prevent="save" class="cf-form">

          <!-- ── IDENTITÉ ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <User :size="13" />
              Identité
            </div>
            <div class="cf-grid-2">
              <div class="cf-field">
                <label class="cf-label">Nom <span class="cf-req">*</span></label>
                <input v-model="form.nom" required class="cf-input" placeholder="Nom de famille" />
              </div>
              <div class="cf-field">
                <label class="cf-label">Prénom <span class="cf-req">*</span></label>
                <input v-model="form.prenom" required class="cf-input" placeholder="Prénom" />
              </div>
            </div>
          </section>

          <!-- ── RELATION ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <Network :size="13" />
              Relation
            </div>
            <div class="cf-grid-2">
              <div class="cf-field">
                <label class="cf-label">Type de relation <span class="cf-req">*</span></label>
                <select v-model="form.relation_type" required class="cf-select">
                  <option value="">Sélectionner…</option>
                  <option v-for="type in typesRelation" :key="type" :value="type">{{ type }}</option>
                </select>
              </div>
              <div class="cf-field">
                <label class="cf-label">Lien familial</label>
                <select v-model="form.lien_familial" class="cf-select">
                  <option value="">Aucun</option>
                  <option v-for="lien in liensFamiliaux" :key="lien" :value="lien">{{ lien }}</option>
                </select>
              </div>
              <div class="cf-field">
                <label class="cf-label">Type de contact</label>
                <select v-model="form.type_de_contact" class="cf-select">
                  <option value="">Sélectionner…</option>
                  <option v-for="type in typesContact" :key="type" :value="type">{{ type }}</option>
                </select>
              </div>
              <div class="cf-field">
                <label class="cf-label">Description de la relation</label>
                <input v-model="form.relation_client" class="cf-input" placeholder="Ex: Voisine de longue date…" />
              </div>
            </div>
          </section>

          <!-- ── INDICATEURS CLINIQUES ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <Activity :size="13" />
              Indicateurs cliniques
            </div>

            <!-- Force du lien -->
            <div class="cf-field cf-field-full">
              <div class="cf-slider-header">
                <label class="cf-label">Force du lien affectif</label>
                <span class="cf-slider-val" :class="lienValClass">
                  {{ form.force_lien > 0 ? '+' : '' }}{{ form.force_lien }}
                  <span class="cf-slider-label">{{ labelForceLien(form.force_lien) }}</span>
                </span>
              </div>
              <input v-model.number="form.force_lien" type="range" min="-10" max="10" step="1"
                class="cf-slider" :style="styleLien" />
              <div class="cf-slider-bounds">
                <span>−10 Lien rompu</span>
                <span>+10 Lien pivot</span>
              </div>
            </div>

            <div class="cf-grid-3">
              <!-- Maltraitance -->
              <div class="cf-field">
                <div class="cf-slider-header">
                  <label class="cf-label flex items-center gap-1">
                    <ShieldAlert :size="12" :class="maltraitanceIconClass" />
                    Maltraitance
                  </label>
                  <span class="cf-slider-val" :class="maltraitanceValClass">{{ form.force_niv_maltraitance }}/5</span>
                </div>
                <input v-model.number="form.force_niv_maltraitance" type="range" min="0" max="5" step="1"
                  class="cf-slider" :style="styleMalt" />
                <p class="cf-slider-hint">{{ labelMaltraitance(form.force_niv_maltraitance) }}</p>
              </div>

              <!-- Soutien -->
              <div class="cf-field">
                <div class="cf-slider-header">
                  <label class="cf-label flex items-center gap-1">
                    <HandHelping :size="12" :class="soutienIconClass" />
                    Soutien apporté
                  </label>
                  <span class="cf-slider-val text-sky-400 dark:text-sky-400">{{ form.force_niv_soutien }}/5</span>
                </div>
                <input v-model.number="form.force_niv_soutien" type="range" min="0" max="5" step="1"
                  class="cf-slider" :style="styleSoutien" />
                <p class="cf-slider-hint">{{ labelSoutien(form.force_niv_soutien) }}</p>
              </div>

              <!-- Épuisement -->
              <div class="cf-field">
                <div class="cf-slider-header">
                  <label class="cf-label flex items-center gap-1">
                    <BatteryLow :size="12" :class="epuisementIconClass" />
                    Épuisement
                  </label>
                  <span class="cf-slider-val" :class="epuisementValClass">{{ form.force_niv_epuisement }}/5</span>
                </div>
                <input v-model.number="form.force_niv_epuisement" type="range" min="0" max="5" step="1"
                  class="cf-slider" :style="styleEpuisement" />
                <p class="cf-slider-hint">{{ labelEpuisement(form.force_niv_epuisement) }}</p>
              </div>
            </div>
          </section>

          <!-- ── RÔLES ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <Shield :size="13" />
              Rôles &amp; mandats
            </div>
            <div class="cf-grid-2">
              <label class="cf-checkbox">
                <input v-model="form.contact_urgence" type="checkbox" :true-value="1" :false-value="0" class="cf-check-input" />
                <div class="cf-check-box">
                  <AlertTriangle :size="14" class="text-red-400" />
                  <div>
                    <span class="cf-check-title">Contact d'urgence</span>
                    <span class="cf-check-sub">À contacter en priorité</span>
                  </div>
                </div>
              </label>
              <label class="cf-checkbox">
                <input v-model="form.aidant_naturel" type="checkbox" :true-value="1" :false-value="0" class="cf-check-input" />
                <div class="cf-check-box">
                  <Heart :size="14" class="text-emerald-400" />
                  <div>
                    <span class="cf-check-title">Aidant naturel</span>
                    <span class="cf-check-sub">Aide régulière au quotidien</span>
                  </div>
                </div>
              </label>
              <label class="cf-checkbox">
                <input v-model="form.procuration_bancaire" type="checkbox" :true-value="1" :false-value="0" class="cf-check-input" />
                <div class="cf-check-box">
                  <Landmark :size="14" class="text-sky-400" />
                  <div>
                    <span class="cf-check-title">Procuration bancaire</span>
                    <span class="cf-check-sub">Gestion des finances</span>
                  </div>
                </div>
              </label>
              <label class="cf-checkbox">
                <input v-model="form.procuration_notariee" type="checkbox" :true-value="1" :false-value="0" class="cf-check-input" />
                <div class="cf-check-box">
                  <ScrollText :size="14" class="text-violet-400" />
                  <div>
                    <span class="cf-check-title">Procuration notariée</span>
                    <span class="cf-check-sub">Mandat légal</span>
                  </div>
                </div>
              </label>
            </div>
          </section>

          <!-- ── COORDONNÉES ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <Phone :size="13" />
              Coordonnées
            </div>
            <div class="cf-grid-2">
              <div class="cf-field">
                <label class="cf-label">Téléphone</label>
                <input v-model="form.telephone" type="tel" class="cf-input" placeholder="819-555-0123" />
              </div>
              <div class="cf-field">
                <label class="cf-label">Cellulaire</label>
                <input v-model="form.cellulaire" type="tel" class="cf-input" placeholder="819-555-0456" />
              </div>
              <div class="cf-field cf-field-full">
                <label class="cf-label">Courriel</label>
                <input v-model="form.email" type="email" class="cf-input" placeholder="exemple@courriel.com" />
              </div>
            </div>
          </section>

          <!-- ── ADRESSE ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <MapPin :size="13" />
              Adresse
            </div>
            <div class="cf-grid-2">
              <div class="cf-field cf-field-full">
                <label class="cf-label">Adresse</label>
                <input v-model="form.adresse" class="cf-input" placeholder="123 Rue Principale" />
              </div>
              <div class="cf-field">
                <label class="cf-label">Ville</label>
                <input v-model="form.ville" class="cf-input" placeholder="Gatineau" />
              </div>
              <div class="cf-field">
                <label class="cf-label">Code postal</label>
                <input v-model="form.code_postal" class="cf-input" placeholder="J8V 1A1" />
              </div>
              <div class="cf-field cf-field-full">
                <label class="cf-label">Pays</label>
                <input v-model="form.pays" class="cf-input" placeholder="Canada" />
              </div>
            </div>
          </section>

          <!-- ── PROFESSION ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <Briefcase :size="13" />
              Informations professionnelles
            </div>
            <div class="cf-grid-2">
              <div class="cf-field">
                <label class="cf-label">Profession</label>
                <input v-model="form.profession" class="cf-input" placeholder="Ex: Infirmière" />
              </div>
              <div class="cf-field">
                <label class="cf-label">Employeur</label>
                <input v-model="form.employeur" class="cf-input" placeholder="Nom de l'employeur" />
              </div>
            </div>
          </section>

          <!-- ── NOTE FIXE ── -->
          <section class="cf-section">
            <div class="cf-section-title">
              <FileText :size="13" />
              Note fixe
            </div>
            <div class="cf-field cf-field-full">
              <textarea v-model="form.note_fixe" rows="3" class="cf-textarea"
                placeholder="Observations importantes, contexte clinique…"></textarea>
            </div>
          </section>

        </form>
      </div>

      <!-- ═══ FOOTER ═══ -->
      <div class="cf-footer">
        <div v-if="error" class="cf-error">
          <AlertCircle :size="14" class="shrink-0" />
          <span>{{ error }}</span>
        </div>
        <div class="cf-footer-actions">
          <button type="button" @click="$emit('close')" :disabled="loading" class="cf-btn-cancel">
            Annuler
          </button>
          <button @click="save" :disabled="loading || !canSubmit" class="cf-btn-save">
            <Loader2 v-if="loading" :size="14" class="animate-spin" />
            <Save v-else :size="14" />
            {{ loading ? 'Enregistrement…' : isEdit ? 'Modifier' : 'Ajouter' }}
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import {
  X, UserPlus, User, Network, Heart, Shield, Phone, MapPin,
  Briefcase, FileText, Save, Loader2, AlertCircle, Activity,
  ShieldAlert, BatteryLow, HandHelping, AlertTriangle,
  Landmark, ScrollText
} from 'lucide-vue-next'
import { CreateContact, UpdateContact } from '../../../wailsjs/go/main/App'
import { useRefListes } from '@/composables/useRefListes'


const { getCategorie } = useRefListes()
const typesRelation  = ref([])
const typesContact   = ref([])
const liensFamiliaux = ref([])

onMounted(async () => {
  const tr = await getCategorie('type_relation')
  typesRelation.value = tr.map(t => t.Libelle)
  const tc = await getCategorie('type_contact')
  typesContact.value = tc.map(t => t.Libelle)
  const lf = await getCategorie('lien_familial')
  liensFamiliaux.value = lf.map(t => t.Libelle)
})

const props = defineProps({
  clientId:    { type: Number, required: true },
  contactData: { type: Object, default: null }
})
const emit = defineEmits(['close', 'saved'])

const loading = ref(false)
const error   = ref('')
const isEdit  = computed(() => !!props.contactData)

const form = ref({
  nom: '', prenom: '',
  telephone: null, cellulaire: null,
  adresse: null, code_postal: null, ville: null, pays: 'Canada',
  email: null, employeur: null, profession: null,
  relation_client: null, lien_familial: null,
  force_lien: 0, force_niv_maltraitance: 0,
  force_niv_soutien: 0, force_niv_epuisement: 0,
  contact_urgence: 0, aidant_naturel: 0,
  type_de_contact: null,
  procuration_bancaire: 0, procuration_notariee: 0,
  note_fixe: null, relation_type: '',
  client_id: props.clientId
})

const canSubmit = computed(() => form.value.nom && form.value.prenom && form.value.relation_type)

const save = async () => {
  loading.value = true
  error.value = ''
  try {
    if (isEdit.value) {
      await UpdateContact({ ...form.value, id: props.contactData.id })
    } else {
      await CreateContact(form.value)
    }
    emit('saved')
  } catch (e) {
    console.error('Erreur sauvegarde contact:', e)
    error.value = e.message || 'Erreur lors de la sauvegarde'
  } finally {
    loading.value = false
  }
}

watch(() => props.contactData, (newVal) => {
  if (newVal) {
    form.value = { ...newVal, client_id: props.clientId }
  }
}, { immediate: true })

// ── Styles sliders ────────────────────────────────────
const styleLien = computed(() => {
  const v   = form.value.force_lien
  const pct = ((v + 10) / 20) * 100
  if (v < 0) return {
    background: `linear-gradient(to right, #dc2626 0%, #f97316 ${pct * 0.6}%, #475569 ${pct}%, #1e293b ${pct}%, #1e293b 100%)`,
    accentColor: '#dc2626'
  }
  if (v === 0) return { background: '#334155', accentColor: '#64748b' }
  return {
    background: `linear-gradient(to right, #1e293b 0%, #1e293b 50%, #0d9488 ${50 + (pct - 50)}%, #0f6e56 100%)`,
    accentColor: '#0d9488'
  }
})
const styleMalt = computed(() => ({
  background: `linear-gradient(to right, #22c55e 0%, #eab308 50%, #ef4444 100%)`,
  accentColor: form.value.force_niv_maltraitance >= 3 ? '#ef4444' : '#f59e0b'
}))
const styleSoutien = computed(() => {
  const pct = (form.value.force_niv_soutien / 5) * 100
  return {
    background: `linear-gradient(to right, #0ea5e9 0% ${pct}%, #1e293b ${pct}% 100%)`,
    accentColor: '#0ea5e9'
  }
})
const styleEpuisement = computed(() => ({
  background: `linear-gradient(to right, #22c55e 0%, #f59e0b 60%, #ea580c 100%)`,
  accentColor: form.value.force_niv_epuisement >= 3 ? '#ea580c' : '#f59e0b'
}))

// ── Classes dynamiques indicateurs ───────────────────
const lienValClass = computed(() => {
  const v = form.value.force_lien
  if (v <= -5) return 'text-red-400'
  if (v < 0)   return 'text-amber-400'
  if (v === 0) return 'text-slate-500'
  if (v <= 5)  return 'text-sky-400'
  return 'text-teal-400'
})
const maltraitanceValClass  = computed(() => form.value.force_niv_maltraitance >= 3 ? 'text-red-400' : form.value.force_niv_maltraitance >= 1 ? 'text-amber-400' : 'text-slate-500')
const maltraitanceIconClass = computed(() => form.value.force_niv_maltraitance >= 3 ? 'text-red-400' : form.value.force_niv_maltraitance >= 1 ? 'text-amber-400' : 'text-slate-600')
const epuisementValClass    = computed(() => form.value.force_niv_epuisement >= 4 ? 'text-orange-400' : form.value.force_niv_epuisement >= 2 ? 'text-amber-400' : 'text-slate-500')
const epuisementIconClass   = computed(() => form.value.force_niv_epuisement >= 3 ? 'text-orange-400' : 'text-slate-600')
const soutienIconClass      = computed(() => form.value.force_niv_soutien >= 3 ? 'text-sky-400' : 'text-slate-600')

// ── Labels ────────────────────────────────────────────
const labelForceLien = (v) => {
  if (v <= -8) return 'Lien rompu'
  if (v <= -5) return 'Conflit sévère'
  if (v <= -3) return 'Relation toxique'
  if (v < 0)   return 'Lien fragilisé'
  if (v === 0) return 'Neutre'
  if (v <= 3)  return 'Lien présent'
  if (v <= 5)  return 'Lien significatif'
  if (v <= 7)  return 'Lien fort'
  return 'Lien pivot'
}
const labelMaltraitance = (v) => ['Aucun indicateur', 'Suspectés (tiers)', 'Indicateurs détectés', 'Indicateurs admis', 'Maltraitance avérée', 'Maltraitance grave'][v] || ''
const labelSoutien      = (v) => ['Aucun soutien', 'Présence occasionnelle', 'Soutien régulier', 'Soutien significatif', 'Aidant principal', 'Présence totale'][v] || ''
const labelEpuisement   = (v) => ['Aucun signe', 'Signes légers', 'Épuisement modéré', 'Épuisement élevé', "Risque d'effondrement", 'Effondrement'][v] || ''
</script>

<style scoped>
/* ── Overlay & Modal ─────────────────────────────────── */
.cf-overlay {
  position: fixed; inset: 0; z-index: 60;
  background: rgba(0,0,0,0.65);
  backdrop-filter: blur(4px);
  display: flex; align-items: center; justify-content: center;
  padding: 1rem;
}
.cf-modal {
  background: #f8fafc;
  color: #1e293b;
  width: 100%; max-width: 760px;
  max-height: 90vh;
  border-radius: 14px;
  overflow: hidden;
  display: flex; flex-direction: column;
  box-shadow: 0 25px 60px rgba(0,0,0,0.4);
}
:root.dark .cf-modal, .dark .cf-modal {
  background: #111827;
  color: #e2e8f0;
}

/* ── Header ──────────────────────────────────────────── */
.cf-header {
  padding: 18px 24px;
  background: #f1f5f9;
  border-bottom: 1px solid #e2e8f0;
  display: flex; align-items: center; justify-content: space-between;
  flex-shrink: 0;
}
.dark .cf-header {
  background: #0f172a;
  border-bottom-color: #1e293b;
}
.cf-header-icon {
  width: 34px; height: 34px;
  background: rgba(13,148,136,0.12);
  border: 1px solid rgba(13,148,136,0.25);
  border-radius: 8px;
  display: flex; align-items: center; justify-content: center;
  color: #0d9488;
}
.cf-title {
  font-size: 15px; font-weight: 600;
  color: #0f172a;
}
.dark .cf-title { color: #f1f5f9; }
.cf-subtitle {
  font-size: 12px; color: #64748b; margin-top: 1px;
}
.cf-close {
  width: 30px; height: 30px;
  display: flex; align-items: center; justify-content: center;
  border-radius: 7px; color: #64748b;
  transition: background 0.15s, color 0.15s;
}
.cf-close:hover { background: rgba(0,0,0,0.06); color: #0f172a; }
.dark .cf-close:hover { background: rgba(255,255,255,0.07); color: #f1f5f9; }

/* ── Body ────────────────────────────────────────────── */
.cf-body { flex: 1; overflow-y: auto; padding: 20px 24px; }
.cf-form { display: flex; flex-direction: column; gap: 20px; }

/* ── Sections ────────────────────────────────────────── */
.cf-section {
  display: flex; flex-direction: column; gap: 12px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e2e8f0;
}
.dark .cf-section { border-bottom-color: #1e293b; }
.cf-section:last-child { border-bottom: none; padding-bottom: 0; }

.cf-section-title {
  display: flex; align-items: center; gap: 6px;
  font-size: 11px; font-weight: 600;
  text-transform: uppercase; letter-spacing: 0.07em;
  color: #64748b;
}
.dark .cf-section-title { color: #475569; }

/* ── Grilles ─────────────────────────────────────────── */
.cf-grid-2 { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; }
.cf-grid-3 { display: grid; grid-template-columns: 1fr 1fr 1fr; gap: 12px; }
@media (max-width: 600px) {
  .cf-grid-2, .cf-grid-3 { grid-template-columns: 1fr; }
}

/* ── Champs ──────────────────────────────────────────── */
.cf-field { display: flex; flex-direction: column; gap: 5px; }
.cf-field-full { grid-column: 1 / -1; }
.cf-label {
  font-size: 12px; font-weight: 500; color: #475569;
  display: flex; align-items: center; gap: 4px;
}
.dark .cf-label { color: #64748b; }
.cf-req { color: #ef4444; }

.cf-input, .cf-select, .cf-textarea {
  width: 100%;
  padding: 9px 12px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  background: #ffffff;
  color: #0f172a;
  font-size: 13px;
  outline: none;
  transition: border-color 0.15s, box-shadow 0.15s;
}
.dark .cf-input, .dark .cf-select, .dark .cf-textarea {
  background: #1e293b;
  border-color: #334155;
  color: #e2e8f0;
}
.cf-input:focus, .cf-select:focus, .cf-textarea:focus {
  border-color: #0d9488;
  box-shadow: 0 0 0 3px rgba(13,148,136,0.12);
}
.cf-input::placeholder, .cf-textarea::placeholder { color: #94a3b8; }
.dark .cf-input::placeholder, .dark .cf-textarea::placeholder { color: #475569; }
.cf-textarea { resize: none; }

/* ── Sliders ─────────────────────────────────────────── */
.cf-slider-header {
  display: flex; align-items: center; justify-content: space-between;
}
.cf-slider-val {
  font-size: 12px; font-weight: 600;
  display: flex; align-items: center; gap: 5px;
}
.cf-slider-label { font-weight: 400; color: #64748b; font-size: 11px; }
.cf-slider {
  width: 100%; height: 5px;
  border-radius: 99px;
  -webkit-appearance: none; appearance: none;
  outline: none; cursor: pointer;
}
.cf-slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  width: 16px; height: 16px;
  border-radius: 50%;
  background: #ffffff;
  border: 2px solid #0d9488;
  box-shadow: 0 1px 4px rgba(0,0,0,0.25);
  cursor: pointer;
}
.dark .cf-slider::-webkit-slider-thumb { background: #1e293b; }
.cf-slider-bounds {
  display: flex; justify-content: space-between;
  font-size: 10px; color: #94a3b8; margin-top: 2px;
}
.cf-slider-hint {
  font-size: 11px; color: #64748b; margin-top: 2px; min-height: 14px;
}

/* ── Checkboxes ──────────────────────────────────────── */
.cf-checkbox { display: flex; cursor: pointer; }
.cf-check-input { display: none; }
.cf-check-box {
  width: 100%;
  display: flex; align-items: center; gap: 10px;
  padding: 10px 12px;
  border: 1px solid #e2e8f0;
  border-radius: 9px;
  background: #f8fafc;
  transition: all 0.15s;
}
.dark .cf-check-box {
  border-color: #1e293b;
  background: rgba(255,255,255,0.02);
}
.cf-check-input:checked + .cf-check-box {
  border-color: #0d9488;
  background: rgba(13,148,136,0.06);
}
.dark .cf-check-input:checked + .cf-check-box {
  border-color: rgba(13,148,136,0.4);
  background: rgba(13,148,136,0.08);
}
.cf-check-title {
  display: block; font-size: 13px; font-weight: 500; color: #1e293b;
}
.dark .cf-check-title { color: #e2e8f0; }
.cf-check-sub {
  display: block; font-size: 11px; color: #94a3b8;
}

/* ── Footer ──────────────────────────────────────────── */
.cf-footer {
  padding: 14px 24px;
  background: #f1f5f9;
  border-top: 1px solid #e2e8f0;
  display: flex; align-items: center; justify-content: space-between;
  flex-shrink: 0; gap: 12px;
}
.dark .cf-footer {
  background: #0f172a;
  border-top-color: #1e293b;
}
.cf-footer-actions { display: flex; gap: 8px; margin-left: auto; }

.cf-error {
  display: flex; align-items: center; gap: 6px;
  font-size: 12px; color: #ef4444;
  background: rgba(239,68,68,0.08);
  border: 1px solid rgba(239,68,68,0.2);
  padding: 6px 10px; border-radius: 7px;
  flex: 1;
}

.cf-btn-cancel {
  padding: 8px 16px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 13px; font-weight: 500;
  color: #475569;
  background: transparent;
  transition: all 0.15s;
}
.cf-btn-cancel:hover { background: #f1f5f9; color: #1e293b; }
.dark .cf-btn-cancel { border-color: #334155; color: #64748b; }
.dark .cf-btn-cancel:hover { background: rgba(255,255,255,0.05); color: #e2e8f0; }

.cf-error {
  display: flex; align-items: center; gap: 6px;
  font-size: 12px; color: #ef4444;
  background: rgba(239,68,68,0.08);
  border: 1px solid rgba(239,68,68,0.2);
  padding: 6px 10px; border-radius: 7px;
  flex: 1;
}

.cf-btn-save {
  padding: 8px 16px;
  border: none;
  border-radius: 8px;
  font-size: 13px; font-weight: 500;
  color: #ffffff;
  background: #0d9488;
  display: flex; align-items: center; gap: 6px;
  transition: background 0.15s; 
}
.cf-btn-save:hover:not(:disabled) { background: #0f766e; }
.cf-btn-save:disabled { opacity: 0.45; cursor: not-allowed; }

/* ── Animation ───────────────────────────────────────── */
@keyframes slideIn {
  from { opacity: 0; transform: scale(0.97) translateY(12px); }
  to   { opacity: 1; transform: scale(1) translateY(0); }
}
.animate-slideIn { animation: slideIn 0.2s ease-out; }
</style>