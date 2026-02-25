<!-- frontend/src/components/Pharmacies/PharmacieClientsList.vue -->
<template>
  <section class="border-t border-slate-200 dark:border-slate-700 pt-6">
    <div class="flex items-center justify-between mb-4">
      <h3
        class="text-lg font-bold text-slate-900 dark:text-white flex items-center gap-2"
      >
        <Users :size="20" class="text-slate-600 dark:text-slate-400" />
        Clients utilisant cette pharmacie
        <span
          v-if="!loading"
          class="text-sm font-normal text-slate-500 dark:text-slate-400"
        >
          ({{ clients.length }})
        </span>
      </h3>

      <!-- Recherche -->
      <div v-if="clients.length > 0" class="relative">
        <Search
          :size="16"
          class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400"
        />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher un client..."
          class="pl-9 pr-4 py-2 text-sm border border-slate-200 dark:border-slate-700 rounded-lg bg-white dark:bg-slate-800 text-slate-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-emerald-500"
        />
      </div>
    </div>

    <!-- Chargement -->
    <div v-if="loading" class="flex items-center justify-center py-8">
      <Loader2 :size="24" class="animate-spin text-slate-400" />
      <span class="ml-3 text-sm text-slate-600 dark:text-slate-400"
        >Chargement des clients...</span
      >
    </div>

    <!-- Erreur -->
    <div
      v-else-if="error"
      class="p-4 bg-rose-50 dark:bg-rose-900/20 rounded-lg border border-rose-200 dark:border-rose-800"
    >
      <p class="text-sm text-rose-700 dark:text-rose-300">{{ error }}</p>
    </div>

    <!-- État vide -->
    <div v-else-if="clients.length === 0" class="text-center py-8">
      <div
        class="inline-flex items-center justify-center w-16 h-16 bg-slate-100 dark:bg-slate-800 rounded-full mb-4"
      >
        <Users :size="32" class="text-slate-400" />
      </div>
      <p class="text-sm font-medium text-slate-700 dark:text-slate-300 mb-1">
        Aucun client assigné
      </p>
      <p class="text-xs text-slate-500 dark:text-slate-400">
        Cette pharmacie n'a pas encore de clients assignés
      </p>
    </div>

    <!-- Liste des clients -->
    <div v-else class="space-y-2 max-h-[400px] overflow-y-auto pr-2 custom-scrollbar">
  <div
    v-for="client in filteredClients"
    :key="client.ID"
    :class="[
      'group flex items-center gap-4 p-3 rounded-lg transition-all duration-200 border border-transparent',
      client.dcd === 1
        ? 'bg-slate-50/50 dark:bg-slate-900/30 opacity-75 border-slate-100 dark:border-slate-800'
        : 'hover:bg-emerald-50/50 dark:hover:bg-emerald-900/10 hover:border-emerald-100 dark:hover:border-emerald-900/30',
    ]"
  >
    <div
      :class="[
        'flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center shadow-sm transition-transform group-hover:scale-105',
        client.dcd === 1 
          ? 'bg-slate-400 dark:bg-slate-600' 
          : 'bg-emerald-600 dark:bg-emerald-500',
      ]"
    >
      <span
        v-if="client.dcd !== 1"
        class="text-sm font-bold text-white uppercase"
      >
        {{ getInitials(client.nom, client.prenom) }}
      </span>
      <Flower v-else :size="20" class="text-white/90" />
    </div>

    <div class="flex-1 min-w-0">
      <div class="flex items-center gap-2">
        <p
          :class="[
            'text-sm font-bold truncate',
            client.dcd === 1
              ? 'text-slate-500 italic'
              : 'text-slate-900 dark:text-white',
          ]"
        >
          {{ client.prenom }} {{ client.nom }}
        </p>
        
        <span
          v-if="client.dcd === 1"
          class="inline-flex items-center px-2 py-0.5 rounded text-[10px] font-bold bg-slate-200 dark:bg-slate-700 text-slate-600 dark:text-slate-400 uppercase"
        >
          DCD
        </span>
      </div>
      
      <p class="text-[11px] text-slate-400 dark:text-slate-500 truncate">
        #{{ client.no_dossier_leopard || 'Sans dossier' }}
      </p>
    </div>

    <div class="flex items-center gap-2">
      <button
        @click="$emit('view-client', client.id)"
        class="opacity-0 group-hover:opacity-100 px-3 py-1.5 rounded-lg text-xs font-semibold transition-all duration-200 flex items-center gap-1.5 shadow-sm"
        :class="
          client.dcd === 1
            ? 'bg-slate-600 hover:bg-slate-700 text-white'
            : 'bg-emerald-600 hover:bg-emerald-700 text-white'
        "
      >
        <FileText :size="14" />
        {{ client.dcd === 1 ? "Consulter archive" : "Voir dossier" }}
      </button>
    </div>
  </div>
</div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { Users, Search, Loader2, FileText, Calendar, Flower } from "lucide-vue-next";
import { GetClientsForPharmacie } from "../../../wailsjs/go/main/App";

const props = defineProps({
  pharmacieId: { type: Number, required: true },
});

const emit = defineEmits(["view-client"]);

const clients = ref([]);
const loading = ref(false);
const error = ref(null);
const searchQuery = ref("");

const filteredClients = computed(() => {
  if (!searchQuery.value) return clients.value;

  const query = searchQuery.value.toLowerCase();
  return clients.value.filter((client) => {
    const fullName = `${client.prenom} ${client.nom}`.toLowerCase();
    const dossier = client.no_dossier_leopard?.toLowerCase() || "";
    return fullName.includes(query) || dossier.includes(query);
  });
});

const getInitials = (nom, prenom) => {
  const n = nom?.charAt(0)?.toUpperCase() || "";
  const p = prenom?.charAt(0)?.toUpperCase() || "";
  return `${p}${n}`;
};

const formatDateNaissance = (date) => {
  if (!date) return "";
  try {
    const d = new Date(date);
    const age = new Date().getFullYear() - d.getFullYear();
    return `${age} ans`;
  } catch {
    return "";
  }
};

const loadClients = async () => {
  loading.value = true;
  error.value = null;

  try {
    clients.value = await GetClientsForPharmacie(props.pharmacieId);
  } catch (err) {
    error.value = "Impossible de charger les clients";
    console.error("Erreur chargement clients:", err);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  loadClients();
});
</script>

<style scoped>
/* Scrollbar personnalisée */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.dark .overflow-y-auto::-webkit-scrollbar-thumb {
  background: #475569;
}
</style>
