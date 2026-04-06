<template>
  <div class="sidebar-root h-full flex flex-col">

    <!-- ═══════════════ HEADER ═══════════════ -->
    <div class="sidebar-header px-5 py-4 flex items-center justify-between">
      <div class="flex items-center gap-3">
        <div class="header-icon">
          <Network :size="18" class="text-teal-300" />
        </div>
        <div>
          <h2 class="text-sm font-semibold text-slate-100 tracking-wide uppercase">Réseau social</h2>
          <p class="text-xs text-slate-400 mt-0.5">{{ contacts.length }} contact{{ contacts.length > 1 ? 's' : '' }} · client #{{ clientId }}</p>
        </div>
      </div>
      <button @click="$emit('close')" class="close-btn">
        <X :size="16" class="text-slate-400" />
      </button>
    </div>

    <!-- ═══════════════ SCORE RÉSEAU ═══════════════ -->
    <div class="network-score mx-4 mb-3 p-3 rounded-xl" :class="scoreClass">
      <div class="flex items-center justify-between mb-2">
        <span class="text-xs font-semibold uppercase tracking-wider" :class="scoreLabelClass">
          {{ scoreLabel }}
        </span>
        <span class="score-badge" :class="scoreBadgeClass">{{ networkScore }}/100</span>
      </div>
      <div class="score-bar-bg">
        <div class="score-bar-fill" :class="scoreBarClass" :style="{ width: networkScore + '%' }"></div>
      </div>
      <p class="text-xs mt-2 opacity-75" :class="scoreLabelClass">{{ scoreDescription }}</p>
    </div>

    <!-- ═══════════════ ALERTES CLINIQUES ═══════════════ -->
    <div v-if="alertes.length > 0" class="mx-4 mb-3 space-y-1.5">
      <div
        v-for="alerte in alertes"
        :key="alerte.id"
        class="alerte-item flex items-start gap-2 px-3 py-2 rounded-lg"
        :class="alerte.classe"
      >
        <component :is="alerte.icone" :size="13" class="mt-0.5 shrink-0" />
        <span class="text-xs leading-snug">{{ alerte.message }}</span>
      </div>
    </div>

    <!-- ═══════════════ ACTIONS + RECHERCHE ═══════════════ -->
    <div class="px-4 mb-3 space-y-2">
      <button @click="openNewContact" class="btn-add w-full flex items-center justify-center gap-2 py-2.5 rounded-xl text-sm font-medium">
        <UserPlus :size="15" />
        Ajouter un contact
      </button>
      <div class="search-wrap flex items-center gap-2 px-3 py-2 rounded-lg">
        <Search :size="14" class="text-slate-500 shrink-0" />
        <input
          v-model="recherche"
          placeholder="Rechercher..."
          class="search-input flex-1 text-sm bg-transparent outline-none text-slate-300 placeholder-slate-600"
        />
      </div>
    </div>

    <!-- ═══════════════ FILTRES ═══════════════ -->
    <div class="px-4 mb-3">
      <div class="flex gap-1.5 flex-wrap">
        <button
          v-for="f in filtres"
          :key="f.value"
          @click="filtreActif = f.value"
          class="filtre-btn text-xs px-2.5 py-1 rounded-lg font-medium transition-all"
          :class="filtreActif === f.value ? 'filtre-actif' : 'filtre-inactif'"
        >
          {{ f.label }}
          <span class="ml-1 opacity-60">{{ compteFiltre(f.value) }}</span>
        </button>
      </div>
    </div>

    <!-- ═══════════════ LISTE CONTACTS ═══════════════ -->
    <div class="flex-1 overflow-y-auto px-4 space-y-2 pb-3">

      <div v-if="loading" class="flex items-center justify-center py-12">
        <Loader2 :size="28" class="animate-spin text-teal-500" />
      </div>

      <div v-else-if="contactsFiltres.length === 0" class="text-center py-10">
        <UserX :size="36" class="mx-auto text-slate-600 mb-2" />
        <p class="text-slate-500 text-sm">Aucun contact</p>
      </div>

      <div
        v-else
        v-for="contact in contactsFiltres"
        :key="contact.id"
        class="contact-card rounded-xl p-3.5 cursor-pointer group"
        @click="editContact(contact)"
      >
        <!-- Ligne 1 : Avatar + Nom + Actions -->
        <div class="flex items-start justify-between mb-2.5">
          <div class="flex items-center gap-2.5">
            <div class="avatar" :class="avatarClass(contact)">
              {{ getInitials(contact) }}
            </div>
            <div>
              <p class="text-sm font-semibold text-slate-100 group-hover:text-teal-300 transition-colors">
                {{ contact.prenom }} {{ contact.nom }}
              </p>
              <p class="text-xs text-slate-500">{{ contact.relation_type || '—' }}
                <span v-if="contact.lien_familial" class="text-slate-600"> · {{ contact.lien_familial }}</span>
              </p>
            </div>
          </div>
          <div class="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
            <button @click.stop="editContact(contact)" class="action-btn">
              <Edit :size="13" class="text-teal-400" />
            </button>
            <button @click.stop="confirmerSuppression(contact.id)" class="action-btn">
              <Trash2 :size="13" class="text-red-400" />
            </button>
          </div>
        </div>

        <!-- Ligne 2 : Badges rôles -->
        <div class="flex flex-wrap gap-1.5 mb-2.5">
          <span v-if="contact.contact_urgence === 1" class="badge badge-urgence">
            <AlertTriangle :size="10" /> Urgence
          </span>
          <span v-if="contact.aidant_naturel === 1" class="badge badge-aidant">
            <Heart :size="10" /> Aidant
          </span>
          <span v-if="contact.procuration_bancaire === 1" class="badge badge-legal">
            <Landmark :size="10" /> Banc.
          </span>
          <span v-if="contact.procuration_notariee === 1" class="badge badge-legal">
            <ScrollText :size="10" /> Notarié
          </span>
          <span v-if="contact.type_de_contact" class="badge badge-type">
            {{ contact.type_de_contact }}
          </span>
        </div>

        <!-- Ligne 3 : Indicateurs cliniques (icônes graduelles) -->
        <div class="flex items-center gap-3 mb-2.5">

          <!-- Force du lien -->
          <div class="indicateur-wrap group/ind" :title="`Force du lien : ${contact.force_lien > 0 ? '+' : ''}${contact.force_lien}`">
            <component :is="lienIcon(contact.force_lien)" :size="14" :class="lienClass(contact.force_lien)" />
            <span class="ind-val" :class="lienClass(contact.force_lien)">
              {{ contact.force_lien > 0 ? '+' : '' }}{{ contact.force_lien }}
            </span>
            <span class="ind-tooltip">{{ lienLabel(contact.force_lien) }}</span>
          </div>

          <!-- Maltraitance -->
          <div v-if="contact.force_niv_maltraitance > 0" class="indicateur-wrap" :title="`Maltraitance : ${contact.force_niv_maltraitance}/5`">
            <ShieldAlert :size="14" :class="maltraitanceClass(contact.force_niv_maltraitance)" />
            <span class="ind-val" :class="maltraitanceClass(contact.force_niv_maltraitance)">{{ contact.force_niv_maltraitance }}/5</span>
            <span class="ind-tooltip">{{ maltraitanceLabel(contact.force_niv_maltraitance) }}</span>
          </div>

          <!-- Épuisement -->
          <div v-if="contact.force_niv_epuisement > 0" class="indicateur-wrap" :title="`Épuisement : ${contact.force_niv_epuisement}/5`">
            <BatteryLow :size="14" :class="epuisementClass(contact.force_niv_epuisement)" />
            <span class="ind-val" :class="epuisementClass(contact.force_niv_epuisement)">{{ contact.force_niv_epuisement }}/5</span>
            <span class="ind-tooltip">{{ epuisementLabel(contact.force_niv_epuisement) }}</span>
          </div>

          <!-- Soutien -->
          <div v-if="contact.force_niv_soutien > 0" class="indicateur-wrap" :title="`Soutien : ${contact.force_niv_soutien}/5`">
            <HandHelping :size="14" :class="soutienClass(contact.force_niv_soutien)" />
            <span class="ind-val" :class="soutienClass(contact.force_niv_soutien)">{{ contact.force_niv_soutien }}/5</span>
            <span class="ind-tooltip">{{ soutienLabel(contact.force_niv_soutien) }}</span>
          </div>
        </div>

        <!-- Ligne 4 : Coordonnées -->
        <div class="space-y-1">
          <p v-if="contact.telephone" class="coord-line">
            <Phone :size="11" class="text-slate-600" /> {{ contact.telephone }}
          </p>
          <p v-if="contact.cellulaire" class="coord-line">
            <Smartphone :size="11" class="text-slate-600" /> {{ contact.cellulaire }}
          </p>
          <p v-if="contact.email" class="coord-line">
            <Mail :size="11" class="text-slate-600" /> {{ contact.email }}
          </p>
          <p v-if="contact.ville" class="coord-line">
            <MapPin :size="11" class="text-slate-600" /> {{ contact.ville }}{{ contact.pays ? ', ' + contact.pays : '' }}
          </p>
          <p v-if="contact.profession || contact.employeur" class="coord-line">
            <Briefcase :size="11" class="text-slate-600" />
            {{ contact.profession }}{{ contact.profession && contact.employeur ? ' · ' : '' }}{{ contact.employeur }}
          </p>
        </div>

        <!-- Ligne 5 : Note fixe (tronquée) -->
        <div v-if="contact.note_fixe" class="note-fixe mt-2 pt-2 border-t border-slate-700/50">
          <p class="text-xs text-slate-500 italic line-clamp-2">
            <FileText :size="10" class="inline mr-1 text-slate-600" />{{ contact.note_fixe }}
          </p>
        </div>

        <!-- Ligne 6 : Relation client -->
        <div v-if="contact.relation_client" class="mt-1">
          <p class="text-xs text-slate-600 italic">« {{ contact.relation_client }} »</p>
        </div>

      </div>
    </div>

    <!-- ═══════════════ ANALYSE RÉSEAU ═══════════════ -->
    <div class="analyse-section mx-4 mb-4 rounded-xl p-3.5 space-y-3">
      <p class="text-xs font-semibold uppercase tracking-wider text-slate-400">Analyse du réseau</p>

      <!-- Carte des rôles -->
      <div class="grid grid-cols-2 gap-2">
        <div class="stat-box">
          <AlertTriangle :size="13" class="text-red-400 mb-1" />
          <p class="text-lg font-bold text-slate-100">{{ stats.urgence }}</p>
          <p class="text-xs text-slate-500">Urgence</p>
        </div>
        <div class="stat-box">
          <Heart :size="13" class="text-emerald-400 mb-1" />
          <p class="text-lg font-bold text-slate-100">{{ stats.aidants }}</p>
          <p class="text-xs text-slate-500">Aidants</p>
        </div>
        <div class="stat-box">
          <ScrollText :size="13" class="text-sky-400 mb-1" />
          <p class="text-lg font-bold text-slate-100">{{ stats.procurations }}</p>
          <p class="text-xs text-slate-500">Procurations</p>
        </div>
        <div class="stat-box">
          <Users :size="13" class="text-violet-400 mb-1" />
          <p class="text-lg font-bold text-slate-100">{{ stats.famille }}</p>
          <p class="text-xs text-slate-500">Famille</p>
        </div>
      </div>

      <!-- Indicateurs réseau agrégés -->
      <div class="space-y-1.5">
        <div class="indicateur-ligne">
          <span class="text-xs text-slate-500">Soutien moyen</span>
          <div class="ind-bar-bg flex-1 mx-2">
            <div class="ind-bar-fill bg-sky-500" :style="{ width: (stats.soutienMoyen / 5 * 100) + '%' }"></div>
          </div>
          <span class="text-xs font-medium text-sky-400">{{ stats.soutienMoyen.toFixed(1) }}/5</span>
        </div>
        <div class="indicateur-ligne">
          <span class="text-xs text-slate-500">Épuisement max</span>
          <div class="ind-bar-bg flex-1 mx-2">
            <div class="ind-bar-fill" :class="stats.epuisementMax >= 4 ? 'bg-orange-500' : stats.epuisementMax >= 2 ? 'bg-amber-500' : 'bg-slate-600'" :style="{ width: (stats.epuisementMax / 5 * 100) + '%' }"></div>
          </div>
          <span class="text-xs font-medium" :class="stats.epuisementMax >= 4 ? 'text-orange-400' : 'text-slate-400'">{{ stats.epuisementMax }}/5</span>
        </div>
        <div class="indicateur-ligne">
          <span class="text-xs text-slate-500">Maltraitance max</span>
          <div class="ind-bar-bg flex-1 mx-2">
            <div class="ind-bar-fill" :class="stats.maltraitanceMax >= 3 ? 'bg-red-500' : stats.maltraitanceMax >= 1 ? 'bg-amber-500' : 'bg-slate-600'" :style="{ width: (stats.maltraitanceMax / 5 * 100) + '%' }"></div>
          </div>
          <span class="text-xs font-medium" :class="stats.maltraitanceMax >= 3 ? 'text-red-400' : 'text-slate-400'">{{ stats.maltraitanceMax }}/5</span>
        </div>
      </div>
    </div>

    <!-- ═══════════════ MODAL FORMULAIRE ═══════════════ -->
    <ContactForm
      v-if="showForm"
      :client-id="clientId"
      :contact-data="selectedContact"
      @close="showForm = false; selectedContact = null"
      @saved="onContactSaved"
    />

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import {
  X, UserPlus, Search, Edit, Trash2, Loader2, UserX,
  AlertTriangle, Heart, Phone, Smartphone, Mail, MapPin,
  Briefcase, FileText, Network, Users, Landmark, ScrollText,
  ShieldAlert, BatteryLow, HandHelping, LinkIcon, Link2Off,
  Link2, BatteryFull, BatteryMedium, ShieldCheck, Shield
} from 'lucide-vue-next'
import ContactForm from './ContactForm.vue'
import { GetAllContactsByClientID, DeleteContact } from '../../../wailsjs/go/main/App'

const props = defineProps({
  clientId: { type: Number, required: true }
})
const emit = defineEmits(['close', 'contact-added', 'contact-deleted'])

const contacts = ref([])
const loading = ref(false)
const showForm = ref(false)
const selectedContact = ref(null)
const filtreActif = ref('all')
const recherche = ref('')

const filtres = [
  { value: 'all',     label: 'Tous' },
  { value: 'urgence', label: 'Urgence' },
  { value: 'aidant',  label: 'Aidants' },
  { value: 'famille', label: 'Famille' },
  { value: 'legal',   label: 'Légal' },
]

// ── Filtrage ──────────────────────────────────────────
const contactsFiltres = computed(() => {
  let list = contacts.value
  if (recherche.value.trim()) {
    const q = recherche.value.toLowerCase()
    list = list.filter(c =>
      `${c.nom} ${c.prenom} ${c.telephone ?? ''} ${c.email ?? ''}`.toLowerCase().includes(q)
    )
  }
  if (filtreActif.value === 'urgence') return list.filter(c => c.contact_urgence === 1)
  if (filtreActif.value === 'aidant')  return list.filter(c => c.aidant_naturel === 1)
  if (filtreActif.value === 'famille') return list.filter(c => c.relation_type === 'Famille')
  if (filtreActif.value === 'legal')   return list.filter(c => c.procuration_bancaire === 1 || c.procuration_notariee === 1)
  return list
})

const compteFiltre = (val) => {
  if (val === 'all')     return contacts.value.length
  if (val === 'urgence') return contacts.value.filter(c => c.contact_urgence === 1).length
  if (val === 'aidant')  return contacts.value.filter(c => c.aidant_naturel === 1).length
  if (val === 'famille') return contacts.value.filter(c => c.relation_type === 'Famille').length
  if (val === 'legal')   return contacts.value.filter(c => c.procuration_bancaire === 1 || c.procuration_notariee === 1).length
  return 0
}

// ── Stats ─────────────────────────────────────────────
const stats = computed(() => {
  const c = contacts.value
  return {
    urgence:        c.filter(x => x.contact_urgence === 1).length,
    aidants:        c.filter(x => x.aidant_naturel === 1).length,
    procurations:   c.filter(x => x.procuration_bancaire === 1 || x.procuration_notariee === 1).length,
    famille:        c.filter(x => x.relation_type === 'Famille').length,
    soutienMoyen:   c.length ? c.reduce((a, x) => a + (x.force_niv_soutien ?? 0), 0) / c.length : 0,
    epuisementMax:  c.length ? Math.max(...c.map(x => x.force_niv_epuisement ?? 0)) : 0,
    maltraitanceMax:c.length ? Math.max(...c.map(x => x.force_niv_maltraitance ?? 0)) : 0,
    lienMoyen:      c.length ? c.reduce((a, x) => a + (x.force_lien ?? 0), 0) / c.length : 0,
  }
})

// ── Score réseau ──────────────────────────────────────
const networkScore = computed(() => {
  if (!contacts.value.length) return 0
  const s = stats.value
  let score = 50
  score += Math.min(s.urgence, 2) * 8
  score += Math.min(s.aidants, 3) * 10
  score += s.soutienMoyen * 4
  score -= s.epuisementMax * 6
  score -= s.maltraitanceMax * 8
  score += Math.max(0, s.lienMoyen) * 2
  return Math.max(0, Math.min(100, Math.round(score)))
})

const scoreLabel = computed(() => {
  const s = networkScore.value
  if (s >= 75) return 'Réseau solide'
  if (s >= 50) return 'Réseau modéré'
  if (s >= 25) return 'Réseau fragile'
  return 'Réseau critique'
})

const scoreDescription = computed(() => {
  const s = networkScore.value
  if (s >= 75) return 'Le client bénéficie d\'un réseau de soutien robuste.'
  if (s >= 50) return 'Réseau fonctionnel mais des vulnérabilités sont présentes.'
  if (s >= 25) return 'Attention — le réseau présente des lacunes significatives.'
  return 'Intervention prioritaire — réseau insuffisant ou à risque.'
})

const scoreClass = computed(() => {
  const s = networkScore.value
  if (s >= 75) return 'score-solide'
  if (s >= 50) return 'score-modere'
  if (s >= 25) return 'score-fragile'
  return 'score-critique'
})
const scoreLabelClass = computed(() => {
  const s = networkScore.value
  if (s >= 75) return 'text-emerald-400'
  if (s >= 50) return 'text-sky-400'
  if (s >= 25) return 'text-amber-400'
  return 'text-red-400'
})
const scoreBadgeClass = computed(() => {
  const s = networkScore.value
  if (s >= 75) return 'badge-score-solide'
  if (s >= 50) return 'badge-score-modere'
  if (s >= 25) return 'badge-score-fragile'
  return 'badge-score-critique'
})
const scoreBarClass = computed(() => {
  const s = networkScore.value
  if (s >= 75) return 'bg-emerald-500'
  if (s >= 50) return 'bg-sky-500'
  if (s >= 25) return 'bg-amber-500'
  return 'bg-red-500'
})

// ── Alertes cliniques ─────────────────────────────────
const alertes = computed(() => {
  const liste = []
  const c = contacts.value

  if (stats.value.urgence === 0)
    liste.push({ id: 'no-urgence', icone: AlertTriangle, message: 'Aucun contact d\'urgence défini.', classe: 'alerte-warning' })

  const aidantEpuise = c.find(x => x.aidant_naturel === 1 && x.force_niv_epuisement >= 4)
  if (aidantEpuise)
    liste.push({ id: 'aidant-epuise', icone: BatteryLow, message: `Aidant épuisé : ${aidantEpuise.prenom} ${aidantEpuise.nom} (${aidantEpuise.force_niv_epuisement}/5)`, classe: 'alerte-danger' })

  if (stats.value.maltraitanceMax >= 3)
    liste.push({ id: 'maltraitance', icone: ShieldAlert, message: `Indicateur maltraitance élevé dans le réseau (${stats.value.maltraitanceMax}/5).`, classe: 'alerte-danger' })

  const lienRompu = c.find(x => x.force_lien <= -7)
  if (lienRompu)
    liste.push({ id: 'lien-rompu', icone: Link2Off, message: `Lien rompu : ${lienRompu.prenom} ${lienRompu.nom}`, classe: 'alerte-warning' })

  if (stats.value.aidants > 0 && stats.value.soutienMoyen < 1.5)
    liste.push({ id: 'soutien-faible', icone: HandHelping, message: 'Soutien global très faible malgré la présence d\'aidants.', classe: 'alerte-info' })

  return liste
})

// ── Indicateurs visuels ───────────────────────────────
const lienIcon = (v) => {
  if (v <= -5) return Link2Off
  if (v < 0)   return Link2Off
  if (v === 0) return LinkIcon
  return Link2
}
const lienClass = (v) => {
  if (v <= -7) return 'text-red-500'
  if (v <= -3) return 'text-orange-400'
  if (v < 0)   return 'text-amber-400'
  if (v === 0) return 'text-slate-500'
  if (v <= 4)  return 'text-sky-400'
  if (v <= 7)  return 'text-teal-400'
  return 'text-emerald-400'
}
const lienLabel = (v) => {
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

const maltraitanceClass = (v) => {
  if (v >= 4) return 'text-red-500'
  if (v >= 2) return 'text-orange-400'
  return 'text-amber-400'
}
const maltraitanceLabel = (v) => ['', 'Suspectés', 'Détectés', 'Admis', 'Involontaire avérée', 'Grave avérée'][v] || ''

const epuisementClass = (v) => {
  if (v >= 4) return 'text-orange-500'
  if (v >= 2) return 'text-amber-400'
  return 'text-yellow-400'
}
const epuisementLabel = (v) => ['', 'Léger', 'Modéré', 'Élevé', 'Risque effondrement', 'Effondrement'][v] || ''

const soutienClass = (v) => {
  if (v >= 4) return 'text-emerald-400'
  if (v >= 2) return 'text-sky-400'
  return 'text-slate-400'
}
const soutienLabel = (v) => ['', 'Occasionnel', 'Régulier', 'Significatif', 'Aidant principal', 'Total'][v] || ''

// ── Avatar ────────────────────────────────────────────
const avatarClass = (contact) => {
  if (contact.contact_urgence === 1) return 'avatar-urgence'
  if (contact.aidant_naturel === 1)  return 'avatar-aidant'
  if (contact.relation_type === 'Famille') return 'avatar-famille'
  if (contact.procuration_bancaire === 1 || contact.procuration_notariee === 1) return 'avatar-legal'
  return 'avatar-default'
}
const getInitials = (c) => `${(c.prenom?.[0] ?? '')}${(c.nom?.[0] ?? '')}`.toUpperCase()

// ── Actions ───────────────────────────────────────────
const openNewContact = () => {
  selectedContact.value = null
  showForm.value = true
}
const editContact = (contact) => {
  selectedContact.value = contact
  showForm.value = true
}
const confirmerSuppression = async (id) => {
  if (confirm('Supprimer ce contact ?')) {
    try {
      await DeleteContact(id)
      await loadContacts()
      emit('contact-deleted')
    } catch (e) {
      console.error('Erreur suppression:', e)
    }
  }
}
const onContactSaved = () => {
  loadContacts()
  emit('contact-added')
  showForm.value = false
  selectedContact.value = null
}
const loadContacts = async () => {
  if (!props.clientId) return
  loading.value = true
  try {
    const result = await GetAllContactsByClientID(props.clientId)
    contacts.value = result || []
  } catch (e) {
    console.error('Erreur contacts:', e)
    contacts.value = []
  } finally {
    loading.value = false
  }
}

onMounted(loadContacts)
</script>

<style scoped>
/* ── Base ──────────────────────────────────────────── */
.sidebar-root {
  background: #0f1623;
  color: #e2e8f0;
  font-family: 'DM Sans', system-ui, sans-serif;
}

/* ── Header ─────────────────────────────────────────── */
.sidebar-header {
  background: linear-gradient(135deg, #1e293b 0%, #0f2336 100%);
  border-bottom: 1px solid #1e3a4a;
}
.header-icon {
  width: 32px; height: 32px;
  background: rgba(20, 184, 166, 0.15);
  border: 1px solid rgba(20, 184, 166, 0.3);
  border-radius: 8px;
  display: flex; align-items: center; justify-content: center;
}
.close-btn {
  width: 28px; height: 28px;
  display: flex; align-items: center; justify-content: center;
  border-radius: 6px;
  transition: background 0.15s;
}
.close-btn:hover { background: rgba(255,255,255,0.08); }

/* ── Score réseau ────────────────────────────────────── */
.network-score { border: 1px solid; }
.score-solide  { background: rgba(16, 185, 129, 0.06); border-color: rgba(16,185,129,0.2); }
.score-modere  { background: rgba(14, 165, 233, 0.06); border-color: rgba(14,165,233,0.2); }
.score-fragile { background: rgba(245, 158, 11, 0.06); border-color: rgba(245,158,11,0.2); }
.score-critique{ background: rgba(239, 68, 68, 0.06);  border-color: rgba(239,68,68,0.2); }
.score-badge {
  font-size: 11px; font-weight: 700;
  padding: 1px 7px; border-radius: 999px;
}
.badge-score-solide  { background: rgba(16,185,129,0.15); color: #34d399; }
.badge-score-modere  { background: rgba(14,165,233,0.15); color: #38bdf8; }
.badge-score-fragile { background: rgba(245,158,11,0.15); color: #fbbf24; }
.badge-score-critique{ background: rgba(239,68,68,0.15);  color: #f87171; }
.score-bar-bg  { height: 4px; background: rgba(255,255,255,0.06); border-radius: 99px; overflow: hidden; }
.score-bar-fill{ height: 100%; border-radius: 99px; transition: width 0.6s ease; }

/* ── Alertes ─────────────────────────────────────────── */
.alerte-danger  { background: rgba(239,68,68,0.08);  border: 1px solid rgba(239,68,68,0.2);  color: #fca5a5; }
.alerte-warning { background: rgba(245,158,11,0.08); border: 1px solid rgba(245,158,11,0.2); color: #fcd34d; }
.alerte-info    { background: rgba(14,165,233,0.08); border: 1px solid rgba(14,165,233,0.2); color: #7dd3fc; }

/* ── Bouton ajouter ──────────────────────────────────── */
.btn-add {
  background: rgba(20, 184, 166, 0.12);
  border: 1px solid rgba(20, 184, 166, 0.3);
  color: #5eead4;
  transition: all 0.15s;
}
.btn-add:hover {
  background: rgba(20, 184, 166, 0.2);
  border-color: rgba(20, 184, 166, 0.5);
}

/* ── Recherche ───────────────────────────────────────── */
.search-wrap {
  background: rgba(255,255,255,0.03);
  border: 1px solid rgba(255,255,255,0.07);
}
.search-input::placeholder { color: #475569; }

/* ── Filtres ─────────────────────────────────────────── */
.filtre-actif   { background: rgba(20,184,166,0.15); color: #5eead4; border: 1px solid rgba(20,184,166,0.3); }
.filtre-inactif { background: rgba(255,255,255,0.03); color: #64748b; border: 1px solid rgba(255,255,255,0.06); }
.filtre-inactif:hover { background: rgba(255,255,255,0.06); color: #94a3b8; }

/* ── Carte contact ───────────────────────────────────── */
.contact-card {
  background: rgba(255,255,255,0.03);
  border: 1px solid rgba(255,255,255,0.07);
  transition: all 0.15s;
}
.contact-card:hover {
  background: rgba(255,255,255,0.055);
  border-color: rgba(20,184,166,0.25);
}

/* ── Avatar ──────────────────────────────────────────── */
.avatar {
  width: 34px; height: 34px; border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
  font-size: 12px; font-weight: 700; flex-shrink: 0;
}
.avatar-urgence { background: rgba(239,68,68,0.2);    color: #fca5a5; border: 1px solid rgba(239,68,68,0.3); }
.avatar-aidant  { background: rgba(16,185,129,0.2);   color: #6ee7b7; border: 1px solid rgba(16,185,129,0.3); }
.avatar-famille { background: rgba(167,139,250,0.2);  color: #c4b5fd; border: 1px solid rgba(167,139,250,0.3); }
.avatar-legal   { background: rgba(14,165,233,0.2);   color: #7dd3fc; border: 1px solid rgba(14,165,233,0.3); }
.avatar-default { background: rgba(100,116,139,0.2);  color: #94a3b8; border: 1px solid rgba(100,116,139,0.2); }

/* ── Badges ──────────────────────────────────────────── */
.badge {
  display: inline-flex; align-items: center; gap: 3px;
  font-size: 10px; font-weight: 600;
  padding: 2px 6px; border-radius: 999px;
}
.badge-urgence { background: rgba(239,68,68,0.15);   color: #fca5a5; }
.badge-aidant  { background: rgba(16,185,129,0.15);  color: #6ee7b7; }
.badge-legal   { background: rgba(14,165,233,0.15);  color: #7dd3fc; }
.badge-type    { background: rgba(100,116,139,0.15); color: #94a3b8; }

/* ── Indicateurs cliniques ───────────────────────────── */
.indicateur-wrap {
  position: relative;
  display: flex; align-items: center; gap: 3px;
  cursor: default;
}
.ind-val { font-size: 10px; font-weight: 600; }
.ind-tooltip {
  display: none;
  position: absolute;
  bottom: calc(100% + 4px);
  left: 50%; transform: translateX(-50%);
  background: #1e293b;
  border: 1px solid #334155;
  color: #cbd5e1;
  font-size: 10px;
  padding: 3px 7px;
  border-radius: 5px;
  white-space: nowrap;
  z-index: 10;
  pointer-events: none;
}
.indicateur-wrap:hover .ind-tooltip { display: block; }

/* ── Coordonnées ─────────────────────────────────────── */
.coord-line {
  display: flex; align-items: center; gap: 5px;
  font-size: 11px; color: #64748b;
}

/* ── Note fixe ───────────────────────────────────────── */
.note-fixe { }

/* ── Action buttons ──────────────────────────────────── */
.action-btn {
  width: 24px; height: 24px;
  display: flex; align-items: center; justify-content: center;
  border-radius: 5px;
  transition: background 0.15s;
}
.action-btn:hover { background: rgba(255,255,255,0.08); }

/* ── Analyse réseau ──────────────────────────────────── */
.analyse-section {
  background: rgba(255,255,255,0.02);
  border: 1px solid rgba(255,255,255,0.06);
}
.stat-box {
  background: rgba(255,255,255,0.03);
  border: 1px solid rgba(255,255,255,0.06);
  border-radius: 10px;
  padding: 10px;
  display: flex; flex-direction: column; align-items: center; text-align: center;
}
.indicateur-ligne {
  display: flex; align-items: center;
}
.ind-bar-bg {
  height: 3px;
  background: rgba(255,255,255,0.06);
  border-radius: 99px; overflow: hidden;
}
.ind-bar-fill {
  height: 100%; border-radius: 99px;
  transition: width 0.5s ease;
}
</style>