<!-- frontend/src/components/shared/ContactRow.vue -->
<!-- 
  🚀 CONTACT ROW ULTIMATE EDITION
  ================================
  Composant ultra-flexible pour afficher n'importe quel type de contact
  avec icône, label, valeur et action interactive.
  
  📦 USAGE:
  ---------
  <ContactRow 
    :icon="Phone"              ← Icône Lucide (Phone, Mail, Smartphone, etc.)
    label="Bureau"             ← Petit texte au-dessus (ex: "Bureau", "Email pro")
    value="514-123-4567"       ← La valeur affichée (numéro, email, adresse, etc.)
    href="tel:514-123-4567"    ← Lien cliquable (tel:, mailto:, https://, etc.)
    action-label="Appeler"     ← Texte du bouton d'action
    color="slate"              ← Couleur du thème (slate/stone/zinc)
    :clickable="true"          ← Rendre toute la ligne cliquable (optionnel)
  />
  
  🎨 COULEURS DISPONIBLES:
  ------------------------
  - slate  → Gris-bleu (défaut) - Parfait pour téléphones
  - stone  → Gris chaud - Idéal pour cellulaires
  - zinc   → Gris neutre - Super pour emails
  
  💡 EXEMPLES CONCRETS:
  ---------------------
  Téléphone bureau:
    <ContactRow :icon="Phone" label="Bureau" value="514-555-1234" 
                href="tel:5145551234" action-label="Appeler" color="slate" />
  
  Email:
    <ContactRow :icon="Mail" label="Courriel" value="nom@domaine.ca" 
                href="mailto:nom@domaine.ca" action-label="Écrire" color="zinc" />
  
  Site web:
    <ContactRow :icon="Globe" label="Site web" value="www.exemple.ca" 
                href="https://www.exemple.ca" action-label="Visiter" color="stone" />
-->

<template>
  <div 
    :class="[
      'flex items-center gap-4 p-4 rounded-lg transition-all duration-200 group',
      'hover:bg-slate-50 dark:hover:bg-slate-800/50',
      clickable ? 'cursor-pointer' : ''
    ]"
    @click="handleRowClick"
  >
    
    <!-- 🎯 SECTION ICÔNE 
         Une belle icône colorée dans un cercle
         La couleur change selon la prop "color" -->
    <div 
      :class="[
        'p-2.5 rounded-lg transition-all duration-200 flex-shrink-0',
        'group-hover:scale-110 group-hover:shadow-md',
        colorClasses.iconBg
      ]"
    >
      <component 
        :is="icon" 
        :size="18" 
        class="text-white"
      />
    </div>

    <!-- 📝 SECTION CONTENU 
         Label + Valeur affichés verticalement
         Le texte s'adapte automatiquement à l'espace disponible -->
    <div class="flex-1 min-w-0">
      <!-- Label (petit texte gris) -->
      <p class="text-xs font-medium text-slate-500 dark:text-slate-400 mb-0.5 transition-colors">
        {{ label }}
      </p>
      
      <!-- Valeur (texte principal en gras) -->
      <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
        {{ value }}
      </p>
      
      <!-- 🆕 Badge optionnel (ex: "Préféré", "Principal", etc.) -->
      <span 
        v-if="badge" 
        class="inline-block mt-1 px-2 py-0.5 text-xs font-medium rounded-full bg-slate-100 dark:bg-slate-700 text-slate-700 dark:text-slate-300"
      >
        {{ badge }}
      </span>
    </div>

    <!-- 🔘 BOUTON D'ACTION 
         Apparaît au survol (opacity-0 → opacity-100)
         Change de couleur selon le thème choisi -->
    <a 
      v-if="!hideAction"
      :href="href"
      :class="[
        'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200',
        'flex items-center gap-2 flex-shrink-0',
        'opacity-0 group-hover:opacity-100',
        'transform translate-x-2 group-hover:translate-x-0',
        colorClasses.button
      ]"
      @click.stop="handleActionClick"
    >
      {{ actionLabel }}
      
      <!-- 🆕 Icône optionnelle dans le bouton -->
      <component 
        v-if="actionIcon" 
        :is="actionIcon" 
        :size="14"
      />
    </a>

    <!-- 🆕 Action personnalisée (slot pour flexibilité maximale) -->
    <slot name="action"></slot>

  </div>
</template>

<script setup>
import { computed } from 'vue'

// 📌 PROPS DU COMPOSANT
// =====================
const props = defineProps({
  // Icône principale (composant Lucide)
  icon: { 
    type: Object, 
    required: true 
  },
  
  // Label au-dessus de la valeur
  label: { 
    type: String, 
    required: true 
  },
  
  // Valeur affichée (numéro, email, etc.)
  value: { 
    type: String, 
    required: true 
  },
  
  // Lien cliquable (tel:, mailto:, https://)
  href: { 
    type: String, 
    required: true 
  },
  
  // Texte du bouton d'action
  actionLabel: { 
    type: String, 
    required: true 
  },
  
  // Thème de couleur (slate, stone, zinc)
  color: { 
    type: String, 
    default: 'slate',
    validator: (val) => ['slate', 'stone', 'zinc'].includes(val)
  },
  
  // 🆕 Badge optionnel (ex: "Préféré", "Principal")
  badge: {
    type: String,
    default: null
  },
  
  // 🆕 Icône optionnelle dans le bouton
  actionIcon: {
    type: Object,
    default: null
  },
  
  // 🆕 Cacher le bouton d'action
  hideAction: {
    type: Boolean,
    default: false
  },
  
  // 🆕 Rendre toute la ligne cliquable
  clickable: {
    type: Boolean,
    default: false
  }
})

// 📌 ÉVÉNEMENTS ÉMIS
// ==================
const emit = defineEmits(['click', 'action-click'])

// 🎨 CLASSES DE COULEUR DYNAMIQUES
// =================================
const colorClasses = computed(() => {
  // console.log('🎨 [ContactRow] Thème sélectionné:', props.color)
  
  const colors = {
    slate: {
      iconBg: 'bg-slate-600 group-hover:bg-slate-700',
      button: 'bg-slate-700 hover:bg-slate-800 text-white shadow-sm hover:shadow-md'
    },
    stone: {
      iconBg: 'bg-stone-600 group-hover:bg-stone-700',
      button: 'bg-stone-700 hover:bg-stone-800 text-white shadow-sm hover:shadow-md'
    },
    zinc: {
      iconBg: 'bg-zinc-600 group-hover:bg-zinc-700',
      button: 'bg-zinc-700 hover:bg-zinc-800 text-white shadow-sm hover:shadow-md'
    }
  }
  
  return colors[props.color]
})

// 🎯 GESTION DU CLIC SUR LA LIGNE
// ================================
const handleRowClick = () => {
  if (props.clickable) {
    // console.log('👆 [ContactRow] Ligne cliquée:', props.label, props.value)
    emit('click', { label: props.label, value: props.value, href: props.href })
    
    // Si clickable, ouvrir le lien directement
    window.open(props.href, '_self')
  }
}

// 🎯 GESTION DU CLIC SUR LE BOUTON
// =================================
const handleActionClick = (event) => {
  // console.log('🔘 [ContactRow] Bouton action cliqué:', props.actionLabel)
  emit('action-click', { 
    label: props.label, 
    value: props.value, 
    href: props.href,
    event 
  })
}

// 🔍 DEBUG AU MONTAGE
// ===================
// import { onMounted } from 'vue'
// onMounted(() => {
//   console.log('✅ [ContactRow] Composant monté avec:', {
//     label: props.label,
//     value: props.value,
//     color: props.color,
//     href: props.href
//   })
// })
</script>

<style scoped>
/* 
  🎨 STYLES PERSONNALISÉS
  =======================
  Animation douce pour le bouton d'action qui glisse depuis la droite
*/

/* Animation d'apparition du bouton */
a {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

/* Effet de survol sur toute la ligne */
.group:hover {
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.05);
}

/* Mode sombre - bordure subtile au survol */
.dark .group:hover {
  border: 1px solid rgba(71, 85, 105, 0.3);
}
</style>