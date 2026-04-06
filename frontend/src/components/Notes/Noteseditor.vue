<template>
  <div class="flex-1 flex flex-col overflow-hidden bg-white dark:bg-gray-950">
    
    <!-- Métadonnées de la note -->
    <div class="p-6 space-y-4 border-b dark:border-gray-800 bg-slate-50/50 dark:bg-gray-900/30">
      <div class="grid grid-cols-3 gap-4">
        <!-- Type de note -->
        <div class="space-y-2">
          <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Type de note *</label>
          <select 
            :value="modelValue.type_note" 
            @input="$emit('update:modelValue', { ...modelValue, type_note: $event.target.value })"
            :class="[
              'w-full bg-white dark:bg-gray-900 border rounded-xl p-3 text-sm dark:text-white outline-none transition-all',
              errors.type_note ? 'border-red-500 ring-2 ring-red-200 dark:ring-red-900' : 'border-gray-300 dark:border-gray-800 focus:ring-2 focus:ring-teal-500'
            ]"
            :disabled="isLocked"
          >
            <option value="">Sélectionner...</option>
            <option value="Suivi">Suivi</option>
            <option value="Ouverture">Ouverture</option>
            <option value="Plan">Plan d'intervention</option>
            <option value="Fermeture">Fermeture</option>
            <option value="Incident">Incident</option>
            <option value="Correction">Correction</option>
            <option value="Addendum">Addendum</option>
          </select>
          <span v-if="errors.type_note" class="text-xs text-red-600 dark:text-red-400">{{ errors.type_note }}</span>
        </div>

        <!-- Date d'intervention -->
        <div class="space-y-2">
          <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Date intervention</label>
          <input 
            type="date" 
            :value="modelValue.date_intervention" 
            @input="$emit('update:modelValue', { ...modelValue, date_intervention: $event.target.value })"
            class="w-full bg-white dark:bg-gray-900 border border-gray-300 dark:border-gray-800 rounded-xl p-3 text-sm dark:text-white outline-none focus:ring-2 focus:ring-teal-500"
            :disabled="isLocked"
          >
        </div>

        <!-- Mode d'intervention -->
        <div class="space-y-2">
          <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Mode</label>
          <select 
            :value="modelValue.mode_intervention" 
            @input="$emit('update:modelValue', { ...modelValue, mode_intervention: $event.target.value })"
            class="w-full bg-white dark:bg-gray-900 border border-gray-300 dark:border-gray-800 rounded-xl p-3 text-sm dark:text-white outline-none focus:ring-2 focus:ring-teal-500"
            :disabled="isLocked"
          >
            <option value="En personne">En personne</option>
            <option value="Téléphone">Téléphone</option>
            <option value="Visioconférence">Visioconférence</option>
            <option value="Courriel">Courriel</option>
            <option value="Domicile">Domicile</option>
          </select>
        </div>
      </div>

      <!-- Sujet -->
      <div class="space-y-2">
        <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1">Sujet de la note *</label>
        <input 
          :value="modelValue.titre" 
          @input="$emit('update:modelValue', { ...modelValue, titre: $event.target.value })"
          :class="[
            'w-full bg-white dark:bg-gray-900 border rounded-xl p-3 text-lg font-semibold dark:text-white placeholder-gray-300 dark:placeholder-gray-700 outline-none transition-all',
            errors.titre ? 'border-red-500 ring-2 ring-red-200 dark:ring-red-900' : 'border-gray-300 dark:border-gray-800 focus:ring-2 focus:ring-teal-500'
          ]"
          placeholder="Ex: Évaluation initiale, Suivi hebdomadaire..."
          :disabled="isLocked"
        />
        <span v-if="errors.titre" class="text-xs text-red-600 dark:text-red-400">{{ errors.titre }}</span>
      </div>

      <!-- Note liée (si correction/addendum) -->
      <div v-if="modelValue.type_note === 'Correction' || modelValue.type_note === 'Addendum'" class="space-y-2">
        <label class="text-[10px] font-bold text-gray-400 uppercase tracking-widest ml-1 flex items-center gap-2">
          <Link2 :size="12" />
          Note associée
        </label>
        <select 
          :value="modelValue.note_liee_id" 
          @input="$emit('update:modelValue', { ...modelValue, note_liee_id: parseInt($event.target.value) || null })"
          class="w-full bg-white dark:bg-gray-900 border border-gray-300 dark:border-gray-800 rounded-xl p-3 text-sm dark:text-white outline-none focus:ring-2 focus:ring-sky-500"
          :disabled="isLocked"
        >
          <option value="">Sélectionner une note...</option>
          <option v-for="note in availableNotes" :key="note.id" :value="note.id">
            {{ formatDate(note.date_note) }} - {{ note.titre }}
          </option>
        </select>
      </div>
    </div>

    <!-- Zone d'édition -->
    <div class="flex-1 flex flex-col overflow-hidden">
      
      <!-- Barre d'outils -->
      <div class="px-6 py-3 border-b dark:border-gray-800 flex items-center justify-between bg-white dark:bg-gray-900">
        <div class="flex items-center gap-4 text-xs text-gray-500 dark:text-gray-400">
          <div class="flex items-center gap-2">
            <Type :size="14" />
            <span>{{ wordCount }} mots</span>
          </div>
          <div class="w-px h-4 bg-gray-300 dark:bg-gray-700"></div>
          <div class="flex items-center gap-2">
            <FileText :size="14" />
            <span>{{ charCount }} caractères</span>
          </div>
        </div>
        
        <div class="flex items-center gap-2">
          <!-- Bouton Templates -->
          <div class="relative">
            <button 
              @click="showTemplates = !showTemplates"
              class="px-3 py-1.5 text-xs bg-teal-100 dark:bg-teal-900/30 text-teal-700 dark:text-teal-300 rounded-lg hover:bg-teal-200 dark:hover:bg-teal-800 transition-colors flex items-center gap-1.5"
              :disabled="isLocked"
            >
              <FileType :size="14" />
              Templates
              <ChevronDown :size="12" :class="{ 'rotate-180': showTemplates }" class="transition-transform" />
            </button>

            <!-- Dropdown templates -->
            <Transition name="fade">
              <div v-if="showTemplates" class="absolute top-full mt-2 right-0 w-72 bg-white dark:bg-gray-800 rounded-xl shadow-2xl border border-gray-200 dark:border-gray-700 z-50 overflow-hidden">
                <div class="p-2 bg-teal-50 dark:bg-teal-900/20 border-b dark:border-gray-700">
                  <p class="text-xs font-bold text-teal-700 dark:text-teal-300 uppercase">Sélectionner un template</p>
                </div>
                <div class="max-h-96 overflow-y-auto">
                  <button
                    v-for="template in templates"
                    :key="template.id"
                    @click="insertTemplate(template)"
                    class="w-full text-left px-4 py-3 hover:bg-teal-50 dark:hover:bg-teal-900/20 transition-colors border-b dark:border-gray-700 last:border-0"
                  >
                    <div class="flex items-start gap-3">
                      <component :is="template.icon" :size="16" class="text-teal-600 dark:text-teal-400 flex-shrink-0 mt-0.5" />
                      <div class="flex-1">
                        <p class="text-sm font-semibold text-gray-800 dark:text-gray-200">{{ template.name }}</p>
                        <p class="text-xs text-gray-500 dark:text-gray-400 mt-0.5">{{ template.description }}</p>
                      </div>
                    </div>
                  </button>
                </div>
              </div>
            </Transition>
          </div>

          <!-- 🎤 Bouton Whisper (préparé pour futur) -->
          <button 
            @click="toggleWhisper"
            :disabled="true"
            class="px-3 py-1.5 text-xs bg-gray-100 dark:bg-gray-800 text-gray-400 dark:text-gray-600 rounded-lg transition-colors flex items-center gap-1.5 cursor-not-allowed opacity-50"
            title="Dictée vocale (bientôt disponible)"
          >
            <Mic :size="14" />
            Dictée
          </button>
          
          <button 
            @click="clearContent"
            v-if="modelValue.contenu && !isLocked"
            class="px-3 py-1.5 text-xs text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
          >
            Effacer
          </button>
        </div>
      </div>

      <!-- Textarea principal -->
      <div class="flex-1 overflow-y-auto p-6">
        <textarea 
          ref="textareaRef"
          :value="modelValue.contenu" 
          @input="$emit('update:modelValue', { ...modelValue, contenu: $event.target.value })"
          :class="[
            'w-full h-full resize-none bg-transparent text-gray-800 dark:text-gray-200 outline-none',
            errors.contenu ? 'text-red-600 dark:text-red-400' : '',
            'font-mono text-sm leading-relaxed'
          ]"
          placeholder="Commencez à rédiger votre note clinique..."
          :disabled="isLocked"
          @keydown.ctrl.s.prevent="$emit('save-draft')"
        ></textarea>
        <span v-if="errors.contenu" class="text-xs text-red-600 dark:text-red-400 mt-2 block">{{ errors.contenu }}</span>
      </div>
    </div>

    <!-- Footer avec actions -->
    <div class="p-4 border-t dark:border-gray-800 bg-slate-50/50 dark:bg-gray-900/30 flex items-center justify-between">
      <div class="flex items-center gap-2 text-xs text-gray-500">
        <Keyboard :size="14" />
        <span>Ctrl+S pour sauvegarder</span>
      </div>

      <div class="flex items-center gap-3">
        <button
          v-if="!isLocked"
          @click="$emit('save-draft')"
          :disabled="isSaving"
          class="px-5 py-2.5 bg-stone-200 dark:bg-stone-700 hover:bg-stone-300 dark:hover:bg-stone-600 text-stone-800 dark:text-stone-200 rounded-xl font-semibold flex items-center gap-2 transition-all disabled:opacity-50"
        >
          <Save :size="16" />
          {{ isSaving ? 'Sauvegarde...' : 'Sauvegarder brouillon' }}
        </button>

        <button
          @click="$emit('finalize')"
          :disabled="isSaving || isLocked"
          class="px-6 py-2.5 bg-gradient-to-r from-teal-600 to-teal-700 hover:from-teal-700 hover:to-teal-800 text-white rounded-xl font-bold flex items-center gap-2 shadow-lg shadow-teal-500/30 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <ShieldCheck :size="18" />
          {{ isLocked ? 'Note verrouillée' : 'Finaliser et signer' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useToast } from '../../composables/useToast.js'

import { useConfirm } from '../../composables/useConfirm.js'
const { confirm } = useConfirm()

import { ref, computed, watch, onMounted } from 'vue'
import { 
  Type, FileText, FileType, ChevronDown, Link2, Mic, Keyboard, 
  Save, ShieldCheck, FileCheck, ClipboardList, FileBarChart, 
  FileSpreadsheet, FileBadge, AlertTriangle 
} from 'lucide-vue-next'

const props = defineProps({
  modelValue: Object,
  errors: { type: Object, default: () => ({}) },
  isSaving: Boolean,
  availableNotes: { type: Array, default: () => [] }
})

defineEmits(['update:modelValue', 'save-draft', 'finalize'])

const textareaRef = ref(null)
const showTemplates = ref(false)

const isLocked = computed(() => 
  props.modelValue.verrouille === true || 
  props.modelValue.verrouille === 1
)

const wordCount = computed(() => {
  if (!props.modelValue.contenu) return 0
  return props.modelValue.contenu.trim().split(/\s+/).filter(w => w.length > 0).length
})

const charCount = computed(() => {
  return props.modelValue.contenu?.length || 0
})

// 📝 TEMPLATES PUISSANTS
const templates = ref([
  {
    id: 'regulier',
    name: 'Note régulière OTSTCFQ',
    description: 'Template conforme aux standards de l\'ordre',
    icon: FileCheck,
    content: `CONTEXTE D'INTERVENTION
[Décrire brièvement le contexte et la raison de l'intervention]

OBSERVATIONS ET DONNÉES RECUEILLIES
[Observations objectives, données factuelles, résultats d'évaluation]

ANALYSE CLINIQUE
[Interprétation professionnelle basée sur le raisonnement clinique]

INTERVENTIONS RÉALISÉES
[Actions posées durant la rencontre]

PLAN DE SUIVI
[Prochaines étapes, recommandations, objectifs]

CONCLUSION
[Synthèse de la rencontre et état de la situation]`
  },
  {
    id: 'ouverture',
    name: 'Note d\'ouverture de dossier',
    description: 'Pour la première rencontre avec le client',
    icon: ClipboardList,
    content: `DATE D'OUVERTURE: [Date]
RÉFÉRENCE REÇUE DE: [Source de référence]
MOTIF DE RÉFÉRENCE: [Raison de la demande]

INFORMATIONS SOCIODÉMOGRAPHIQUES
[Âge, situation familiale, occupation, logement, etc.]

MOTIF DE CONSULTATION
[Demande du client dans ses mots]

HISTORIQUE PERTINENT
[Antécédents médicaux, sociaux, occupationnels pertinents]

ÉVALUATION INITIALE
[Premières observations et impressions cliniques]

ANALYSE PRÉLIMINAIRE
[Interprétation initiale basée sur les données recueillies]

SÉQUENCE DES INTERVENTIONs
[Approches envisagées, objectifs généraux]


OBJECTIFS IDENTIFIÉS
1. [Objectif 1]
2. [Objectif 2]
3. [Objectif 3]

PLAN D'INTERVENTION PRÉLIMINAIRE
[Approches envisagées, fréquence des rencontres]

CONSENTEMENT
☐ Consentement obtenu
☐ Confidentialité expliquée
☐ Droits et limites discutés`
  },
  {
    id: 'plan',
    name: 'Plan d\'intervention',
    description: 'Structuration des objectifs et moyens',
    icon: FileBarChart,
    content: `PLAN D'INTERVENTION - [Date]

PROBLÉMATIQUE CIBLÉE
[Description de la situation nécessitant intervention]

OBJECTIF GÉNÉRAL
[Résultat visé à moyen/long terme]

OBJECTIFS SPÉCIFIQUES
1. [Objectif mesurable 1] - Échéance: [Date]
2. [Objectif mesurable 2] - Échéance: [Date]
3. [Objectif mesurable 3] - Échéance: [Date]

MOYENS ET STRATÉGIES
• [Moyen 1 avec fréquence/durée]
• [Moyen 2 avec fréquence/durée]
• [Moyen 3 avec fréquence/durée]

RÔLES ET RESPONSABILITÉS
Client: [Ce que le client s'engage à faire]
Intervenant: [Ce que l'intervenant s'engage à faire]
Autres: [Partenaires/ressources impliqués]

CRITÈRES D'ÉVALUATION
[Indicateurs pour mesurer l'atteinte des objectifs]

RÉVISION PRÉVUE: [Date]`
  },
  {
    id: 'suivi',
    name: 'Note de suivi rapide',
    description: 'Pour rencontres de suivi courantes',
    icon: FileSpreadsheet,
    content: `SUIVI DU: [Date]

DEPUIS LA DERNIÈRE RENCONTRE
[Évolution, changements, événements significatifs]

OBJECTIFS TRAVAILLÉS AUJOURD'HUI
[Rappel des objectifs ciblés]

INTERVENTIONS
[Ce qui a été fait durant la rencontre]

RÉACTION/PARTICIPATION DU CLIENT
[Engagement, difficultés rencontrées, commentaires]

PROGRÈS OBSERVÉS
[Éléments positifs, acquis, améliorations]

DÉFIS PERSISTANTS
[Obstacles, difficultés qui demeurent]

PROCHAINE RENCONTRE
Date: [Date]
Objectif: [Focus de la prochaine séance]`
  },
  {
    id: 'incident',
    name: 'Note d\'incident',
    description: 'Pour événements critiques ou préoccupants',
    icon: AlertTriangle,
    content: `RAPPORT D'INCIDENT - [Date et heure]

NATURE DE L'INCIDENT
[Description factuelle de l'événement]

PERSONNES IMPLIQUÉES
[Client, témoins, autres intervenants]

CIRCONSTANCES
[Contexte dans lequel l'incident s'est produit]

OBSERVATIONS IMMÉDIATES
[État du client, réactions, comportements observés]

INTERVENTIONS POSÉES
[Actions immédiates entreprises]

MESURES DE SÉCURITÉ
[Évaluation du risque, mesures mises en place]

COMMUNICATIONS EFFECTUÉES
☐ Superviseur informé: [Nom]
☐ Famille contactée: [Nom]
☐ Service d'urgence: [Détails]
☐ Autres: [Préciser]

SUIVI REQUIS
[Actions à entreprendre, surveillance nécessaire]`
  },
  {
    id: 'fermeture',
    name: 'Note de fermeture',
    description: 'Pour la fin du suivi',
    icon: FileBadge,
    content: `NOTE DE FERMETURE - [Date]

PÉRIODE DE SUIVI
Ouverture: [Date]
Fermeture: [Date]
Durée totale: [X mois]
Nombre de rencontres: [X]

RAISON DE LA FERMETURE
☐ Objectifs atteints
☐ Fin du mandat
☐ Transfert vers: [Ressource]
☐ Abandon du client
☐ Autre: [Préciser]

BILAN DES OBJECTIFS
Objectif 1: [Atteint / Partiellement atteint / Non atteint]
Objectif 2: [Atteint / Partiellement atteint / Non atteint]
Objectif 3: [Atteint / Partiellement atteint / Non atteint]

ÉVOLUTION GLOBALE
[Synthèse du parcours, progrès réalisés, défis rencontrés]

ÉTAT ACTUEL DU CLIENT
[Situation occupationnelle, fonctionnement, autonomie]

RECOMMANDATIONS
[Suggestions pour le maintien des acquis, ressources à consulter]

SUIVI FUTUR SI NÉCESSAIRE
[Conditions de reprise de services, ressources disponibles]

TRANSFERT/RÉFÉRENCE
[Si applicable: à qui, pour quoi, coordonnées]`
  }
])


const insertTemplate = (template) => {
  if (props.modelValue.contenu && props.modelValue.contenu.trim().length > 0) {
    if (!confirm('Voulez-vous vraiment remplacer le contenu actuel par ce template ?')) {
      showTemplates.value = false
      return
    }
  }
  
  props.modelValue.contenu = template.content
  showTemplates.value = false
  textareaRef.value?.focus()
}

const clearContent = () => {
  if (confirm('Voulez-vous vraiment effacer tout le contenu ?')) {
    props.modelValue.contenu = ''
    textareaRef.value?.focus()
  }
}

const toggleWhisper = () => {
  // 🎤 Placeholder pour Whisper - sera implémenté avec Wails
  console.log('🎤 Whisper integration coming soon...')
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('fr-CA')
}

// Fermer dropdown si on clique ailleurs
onMounted(() => {
  document.addEventListener('click', (e) => {
    if (!e.target.closest('.relative')) {
      showTemplates.value = false
    }
  })
})
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: all 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>