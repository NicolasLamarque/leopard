<template>
  <div class="space-y-1">
    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
      Pays / Origine
    </label>
    
    <div class="relative flex items-center gap-2">
      <!-- Toggle strict/libre -->
      <button
        type="button"
        @click="strictMode = !strictMode"
        :class="[
          'p-2 rounded-lg border-2 transition-all flex-shrink-0',
          strictMode 
            ? 'bg-blue-50 dark:bg-blue-900/30 border-blue-500 text-blue-600' 
            : 'bg-gray-50 dark:bg-gray-800 border-gray-300 dark:border-gray-600 text-gray-400'
        ]"
        :title="strictMode ? 'Mode liste stricte' : 'Mode saisie libre'"
      >
        <component :is="strictMode ? Lock : Unlock" :size="18" />
      </button>

      <!-- Container stable pour éviter l'erreur Vue -->
      <div class="relative flex-1">
        <!-- Mode libre : input avec suggestions -->
        <template v-if="!strictMode">
          <input 
            type="text"
            :value="modelValue"
            @input="$emit('update:modelValue', $event.target.value)"
            list="liste-pays-db"
            placeholder="Taper ou choisir..."
            class="w-full pl-4 pr-10 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 outline-none transition-all"
          />
        </template>
        
        <!-- Mode strict : select -->
        <template v-else>
          <select
            :value="getAlpha2FromName(modelValue)"
            @change="updateFromAlpha2($event.target.value)"
            class="w-full pl-4 pr-10 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 outline-none transition-all appearance-none"
          >
            <option value="">Sélectionner...</option>
            <option v-for="p in paysListe" :key="p.id" :value="p.alpha2">
              {{ getFlagEmoji(p.alpha2) }} {{ p.pays }}
            </option>
          </select>
        </template>
        
        <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none flex items-center">
  <div 
  v-if="paysActuel && paysActuel.alpha2" 
  class="w-7 h-5 overflow-hidden rounded-sm border border-gray-200 dark:border-gray-700 shadow-sm flex"
>
  <img 
    :src="getFlagUrl(paysActuel.alpha2)" 
    class="w-full h-full object-cover"
    :alt="paysActuel.alpha2"
  />
</div>
  <Globe v-else :size="20" class="text-gray-400 opacity-50" />
</div>
      </div>
    </div>

    <!-- Datalist pour mode libre -->
    <datalist id="liste-pays-db">
      <option v-for="p in paysListe" :key="p.id" :value="p.pays" />
    </datalist>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { Globe, Lock, Unlock } from 'lucide-vue-next'
import { GetAllPays } from '../../../wailsjs/go/main/App'

const props = defineProps({
  modelValue: String
})

const emit = defineEmits(['update:modelValue'])

const paysListe = ref([])
const strictMode = ref(true)

onMounted(async () => {
  try {
    const data = await GetAllPays()
    console.log("DATA BRUTE:", data)        // ← ajoute ça
    console.log("TYPE:", typeof data)       // ← et ça
    paysListe.value = data || []
    console.log("Pays chargés:", paysListe.value.length)
  } catch (err) {
    console.error('Erreur GetAllPays:', err)  // ← et change ça
  }
})

const paysActuel = computed(() => {
  if (!props.modelValue) return null
  return paysListe.value.find(p => 
    p.pays.toLowerCase() === props.modelValue.trim().toLowerCase()
  )
})

function getFlagEmoji(alpha2) {
  if (!alpha2) return ''
  return alpha2.toUpperCase().replace(/./g, char => 
    String.fromCodePoint(char.charCodeAt(0) + 127397)
  )
}

function getAlpha2FromName(name) {
  const pays = paysListe.value.find(p => 
    p.pays.toLowerCase() === name?.toLowerCase()
  )
  return pays?.alpha2 || ''
}

function updateFromAlpha2(alpha2) {
  const pays = paysListe.value.find(p => p.alpha2 === alpha2)
  emit('update:modelValue', pays?.pays || '')
}

// Fonction pour récupérer l'image localement
const getFlagUrl = (alpha2) => {
  if (!alpha2) return null
  // On force en minuscule pour matcher les fichiers téléchargés
  const code = alpha2.toLowerCase().trim()
  // Utilisation de l'API native de Vite pour les assets dynamiques
  return new URL(`../../assets/flags/${code}.svg`, import.meta.url).href
  
     
}
</script>