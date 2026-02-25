<template>
  <div class="min-h-screen bg-slate-100 dark:bg-slate-950 flex flex-col">

    <!-- ══════════════════════════════════════════════════════
         BARRE D'EN-TÊTE FIXE
    ══════════════════════════════════════════════════════ -->
    <header class="sticky top-0 z-30 bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 shadow-sm">
      <div class="px-6 py-3 flex items-center justify-between gap-4">

        <!-- Titre -->
        <div class="flex items-center gap-3">
          <div class="w-8 h-8 rounded-lg bg-teal-700 flex items-center justify-center shadow">
            <FolderOpen :size="16" class="text-white" />
          </div>
          <div class="leading-tight">
            <p class="text-[10px] font-bold tracking-[0.15em] text-slate-400 dark:text-slate-500 uppercase">Léopard · Clinique</p>
            <h1 class="text-base font-bold text-slate-800 dark:text-slate-100">Dossiers clients</h1>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center gap-2">
          <span v-if="isLoading" class="flex items-center gap-1.5 text-xs text-slate-400 mr-1">
            <Loader2 :size="13" class="animate-spin" /> Chargement…
          </span>

          <!-- Recherche rapide (raccourci clavier) -->
          <div
            class="hidden md:flex items-center gap-2 px-3 py-1.5 rounded-lg border border-slate-200 dark:border-slate-700 bg-slate-50 dark:bg-slate-800 text-slate-400 text-xs cursor-pointer hover:border-teal-400 transition-colors"
            @click="focusSearch"
          >
            <Search :size="13" />
            <span>Rechercher…</span>
            <kbd class="ml-2 px-1.5 py-0.5 text-[10px] bg-slate-200 dark:bg-slate-700 rounded font-mono">Ctrl K</kbd>
          </div>

          <!-- Filtres rapides -->
          <div class="flex items-center gap-0 border border-slate-200 dark:border-slate-700 rounded-lg overflow-hidden">
            <button
              v-for="f in filtresRapides"
              :key="f.value"
              @click="filterStatut = f.value"
              :class="[
                'px-3 py-1.5 text-xs font-semibold transition-all',
                filterStatut === f.value
                  ? 'bg-teal-700 text-white'
                  : 'bg-white dark:bg-slate-900 text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800'
              ]"
            >
              {{ f.label }}
            </button>
          </div>

          <!-- Rafraîchir -->
          <button
            @click="loadClients"
            :disabled="isLoading"
            title="Actualiser"
            class="p-2 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-teal-50 dark:hover:bg-teal-900/20 transition-all disabled:opacity-40"
          >
            <RefreshCw :size="15" :class="{ 'animate-spin': isLoading }" />
          </button>

          <!-- Basculer vue -->
          <button
            @click="toggleView"
            title="Changer la vue"
            class="p-2 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-teal-50 dark:hover:bg-teal-900/20 transition-all"
          >
            <component :is="viewMode === 'table' ? LayoutGrid : List" :size="15" />
          </button>

          <!-- Nouveau dossier -->
          <button
            @click="showForm = true"
            class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-semibold bg-teal-700 hover:bg-teal-800 text-white shadow-sm hover:shadow transition-all"
          >
            <UserPlus :size="15" />
            <span class="hidden sm:inline">Nouveau dossier</span>
          </button>
        </div>
      </div>
    </header>

    <!-- ══════════════════════════════════════════════════════
         CORPS
    ══════════════════════════════════════════════════════ -->
    <main class="flex-1 px-6 py-5 max-w-screen-2xl mx-auto w-full space-y-4">

      <!-- ─── STATISTIQUES ─── -->
      <div class="grid grid-cols-2 xl:grid-cols-4 gap-3">

        <!-- Total -->
        <div
          class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4 flex items-center gap-3 cursor-pointer hover:border-slate-400 dark:hover:border-slate-600 transition-all"
          @click="filterStatut = 'tous'"
        >
          <div class="w-10 h-10 rounded-lg bg-slate-100 dark:bg-slate-800 flex items-center justify-center shrink-0">
            <Users :size="19" class="text-slate-600 dark:text-slate-400" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-2xl font-bold text-slate-800 dark:text-slate-100 leading-none tabular-nums">{{ stats.total }}</p>
            <p class="text-xs text-slate-500 mt-0.5">Total dossiers</p>
          </div>
          <span class="text-[10px] font-bold tracking-widest text-slate-300 dark:text-slate-600 uppercase shrink-0">ALL</span>
        </div>

        <!-- Actifs -->
        <div
          class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4 flex items-center gap-3 cursor-pointer hover:border-emerald-400 transition-all"
          @click="filterStatut = 'actif'"
        >
          <div class="w-10 h-10 rounded-lg bg-emerald-50 dark:bg-emerald-900/20 flex items-center justify-center shrink-0">
            <UserCheck :size="19" class="text-emerald-600 dark:text-emerald-400" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-2xl font-bold text-slate-800 dark:text-slate-100 leading-none tabular-nums">{{ stats.actifs }}</p>
            <p class="text-xs text-slate-500 mt-0.5">Actifs</p>
          </div>
          <span class="text-xs font-semibold px-2 py-0.5 rounded-full bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-400 shrink-0">
            {{ stats.total > 0 ? Math.round((stats.actifs / stats.total) * 100) : 0 }}%
          </span>
        </div>

        <!-- Nouveaux ce mois -->
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4 flex items-center gap-3">
          <div class="w-10 h-10 rounded-lg bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center shrink-0">
            <CalendarPlus :size="19" class="text-blue-600 dark:text-blue-400" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-2xl font-bold text-slate-800 dark:text-slate-100 leading-none tabular-nums">{{ stats.nouveaux }}</p>
            <p class="text-xs text-slate-500 mt-0.5">Nouveaux ce mois</p>
          </div>
          <TrendingUp :size="16" class="text-blue-400 shrink-0" />
        </div>

        <!-- Archivés / DCD -->
        <div
          class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4 flex items-center gap-3 cursor-pointer hover:border-orange-400 transition-all"
          @click="filterStatut = 'inactif'"
        >
          <div class="w-10 h-10 rounded-lg bg-orange-50 dark:bg-orange-900/20 flex items-center justify-center shrink-0">
            <Archive :size="19" class="text-orange-600 dark:text-orange-400" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-2xl font-bold text-slate-800 dark:text-slate-100 leading-none tabular-nums">{{ stats.archives }}</p>
            <p class="text-xs text-slate-500 mt-0.5">Archivés</p>
          </div>
          <span v-if="stats.dcd > 0" class="text-[10px] font-semibold px-1.5 py-0.5 rounded bg-slate-100 dark:bg-slate-800 text-slate-500 shrink-0">
            {{ stats.dcd }} DCD
          </span>
        </div>

      </div>

      <!-- ─── BARRE RECHERCHE + FILTRES ─── -->
      <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 px-4 py-3">
        <div class="flex flex-col sm:flex-row gap-3 items-center">

          <!-- Champ recherche -->
          <div class="flex-1 relative w-full">
            <Search :size="15" class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none" />
            <input
              ref="searchInput"
              v-model="searchQuery"
              type="text"
              placeholder="Nom, prénom, N° dossier, ville, téléphone…"
              class="w-full pl-9 pr-8 py-2 text-sm border border-slate-200 dark:border-slate-700 rounded-lg bg-slate-50 dark:bg-slate-800 text-slate-800 dark:text-slate-100 placeholder-slate-400 focus:outline-none focus:border-teal-500 focus:ring-2 focus:ring-teal-500/20 transition-all"
            />
            <button
              v-if="searchQuery"
              @click="searchQuery = ''"
              class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600"
            >
              <X :size="14" />
            </button>
          </div>

          <!-- Tri -->
          <select
            v-model="sortBy"
            class="px-3 py-2 text-sm border border-slate-200 dark:border-slate-700 rounded-lg bg-slate-50 dark:bg-slate-800 text-slate-700 dark:text-slate-300 focus:outline-none focus:border-teal-500 transition-all shrink-0"
          >
            <option value="nom">Trier : Nom A→Z</option>
            <option value="prenom">Trier : Prénom</option>
            <option value="recent">Trier : Plus récent</option>
            <option value="age">Trier : Âge ↓</option>
          </select>

          <!-- Compteur résultats -->
          <div class="flex items-center gap-2 shrink-0 text-xs text-slate-400">
            <span>
              <span class="font-semibold text-slate-600 dark:text-slate-300">{{ filteredClients.length }}</span>
              résultat{{ filteredClients.length !== 1 ? 's' : '' }}
            </span>
            <button
              v-if="searchQuery || filterStatut !== 'tous'"
              @click="resetFilters"
              class="text-teal-700 dark:text-teal-400 font-semibold hover:underline"
            >
              Réinitialiser
            </button>
          </div>

        </div>
      </div>

      <!-- ══════════════════════════════════════════════════════
           VUE TABLEAU
      ══════════════════════════════════════════════════════ -->
      <div v-if="viewMode === 'table'" class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 overflow-hidden shadow-sm">
        <div class="overflow-x-auto">
          <table class="min-w-full">

            <thead>
              <tr class="border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60">
                <th class="px-5 py-3 text-left">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">Client</span>
                </th>
                <th class="px-4 py-3 text-left hidden sm:table-cell">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">N° Dossier</span>
                </th>
                <th class="px-4 py-3 text-left hidden md:table-cell">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">Âge / Sexe</span>
                </th>
                <th class="px-4 py-3 text-left hidden lg:table-cell">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">Contact</span>
                </th>
                <th class="px-4 py-3 text-left hidden xl:table-cell">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">Localisation</span>
                </th>
                <th class="px-4 py-3 text-left">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">Statut</span>
                </th>
                <th class="px-4 py-3 text-right">
                  <span class="text-[10px] font-bold uppercase tracking-wider text-slate-400">Actions</span>
                </th>
              </tr>
            </thead>

            <tbody class="divide-y divide-slate-100 dark:divide-slate-800">

              <!-- Lignes clients -->
              <tr
                v-for="client in filteredClients"
                :key="client.id"
                @click="viewClient(client.id)"
                class="hover:bg-teal-50/60 dark:hover:bg-teal-900/10 transition-colors cursor-pointer group"
                :class="{ 'opacity-55': client.dcd === 1 }"
              >
                <!-- Identité -->
                <td class="px-5 py-3.5 whitespace-nowrap">
                  <div class="flex items-center gap-3">
                    <div class="relative shrink-0">
                      <div
                        :class="getAvatarClass(client)"
                        class="w-9 h-9 rounded-lg flex items-center justify-center text-white text-sm font-bold shadow-sm"
                      >
                        {{ getInitials(client) }}
                      </div>
                      <span
                        v-if="isBirthday(client)"
                        class="absolute -top-1.5 -right-1.5 text-base leading-none"
                        title="Anniversaire aujourd'hui !"
                      >🎂</span>
                    </div>
                    <div class="min-w-0">
                      <p class="text-sm font-semibold text-slate-800 dark:text-slate-100 group-hover:text-teal-700 dark:group-hover:text-teal-400 transition-colors leading-tight">
                        {{ client.nom }},&nbsp;{{ client.prenom }}
                      </p>
                      <p v-if="client.date_naissance" class="text-xs text-slate-400 font-mono mt-0.5">
                        {{ formatDateCourt(client.date_naissance) }}
                      </p>
                    </div>
                  </div>
                </td>

                <!-- N° Dossier -->
                <td class="px-4 py-3.5 whitespace-nowrap hidden sm:table-cell">
                  <div class="flex items-center gap-1.5">
                    <ShieldCheck :size="12" class="text-teal-600 dark:text-teal-500 shrink-0" />
                    <span class="font-mono text-xs font-semibold text-slate-600 dark:text-slate-300">
                      {{ client.no_dossier_leopard || '—' }}
                    </span>
                  </div>
                </td>

                <!-- Âge / Sexe -->
                <td class="px-4 py-3.5 whitespace-nowrap hidden md:table-cell">
                  <p class="text-sm font-semibold text-slate-700 dark:text-slate-300 tabular-nums">
                    {{ calcAge(client.date_naissance) }}<span class="text-xs font-normal text-slate-400"> ans</span>
                  </p>
                  <p class="text-[10px] text-slate-400">{{ getSexeLabel(client.sexe) }}</p>
                </td>

                <!-- Contact -->
                <td class="px-4 py-3.5 hidden lg:table-cell">
                  <div class="space-y-0.5">
                    <p v-if="client.telephone" class="flex items-center gap-1.5 text-xs text-slate-600 dark:text-slate-400">
                      <Phone :size="11" class="text-slate-400 shrink-0" />
                      {{ formatPhone(client.telephone) }}
                    </p>
                    <p v-if="client.cellulaire" class="flex items-center gap-1.5 text-xs text-slate-500">
                      <Smartphone :size="11" class="text-slate-400 shrink-0" />
                      {{ formatPhone(client.cellulaire) }}
                    </p>
                    <p v-if="!client.telephone && !client.cellulaire" class="text-xs text-slate-300 dark:text-slate-600 italic">Aucun contact</p>
                  </div>
                </td>

                <!-- Localisation -->
                <td class="px-4 py-3.5 hidden xl:table-cell">
                  <p v-if="client.ville" class="text-xs text-slate-600 dark:text-slate-400 flex items-center gap-1.5">
                    <MapPin :size="11" class="text-slate-400 shrink-0" />
                    {{ client.ville }}<span v-if="client.province" class="text-slate-400">, {{ client.province }}</span>
                  </p>
                  <p v-if="getEtablissement(client)" class="text-[10px] text-slate-400 mt-0.5 flex items-center gap-1">
                    <Building2 :size="10" class="shrink-0" />
                    {{ getEtablissement(client) }}
                  </p>
                </td>

                <!-- Statut -->
                <td class="px-4 py-3.5 whitespace-nowrap">
                  <span
                    :class="getStatusBadgeClass(client)"
                    class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[11px] font-semibold"
                  >
                    <span class="w-1.5 h-1.5 rounded-full" :class="getStatusDotClass(client)"></span>
                    {{ getStatusText(client) }}
                  </span>
                  <p v-if="client.note_fixe" class="mt-0.5 flex items-center gap-1 text-[10px] text-amber-600 dark:text-amber-400">
                    <AlertTriangle :size="10" />
                    Note
                  </p>
                </td>

                <!-- Actions -->
                <td class="px-4 py-3.5 whitespace-nowrap text-right" @click.stop>
                  <div class="flex items-center justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                    <button
                      @click.stop="viewClient(client.id)"
                      class="p-1.5 rounded-lg hover:bg-teal-100 dark:hover:bg-teal-900/30 transition-colors"
                      title="Ouvrir le dossier"
                    >
                      <FolderOpen :size="15" class="text-teal-700 dark:text-teal-400" />
                    </button>
                    <button
                      v-if="client.telephone || client.cellulaire"
                      @click.stop="quickCall(client)"
                      class="p-1.5 rounded-lg hover:bg-blue-100 dark:hover:bg-blue-900/30 transition-colors"
                      title="Appeler"
                    >
                      <Phone :size="15" class="text-blue-600 dark:text-blue-400" />
                    </button>
                    <button
                      @click.stop="confirmDelete(client)"
                      class="p-1.5 rounded-lg hover:bg-red-100 dark:hover:bg-red-900/30 transition-colors"
                      title="Supprimer"
                    >
                      <Trash2 :size="15" class="text-red-500 dark:text-red-400" />
                    </button>
                  </div>
                  <ChevronRight :size="14" class="text-slate-300 dark:text-slate-600 inline group-hover:hidden" />
                </td>
              </tr>

              <!-- État vide -->
              <tr v-if="filteredClients.length === 0 && !isLoading">
                <td colspan="7" class="px-6 py-16 text-center">
                  <UserX :size="40" class="mx-auto text-slate-300 dark:text-slate-700 mb-3" />
                  <p class="text-sm font-semibold text-slate-500 dark:text-slate-400 mb-1">Aucun dossier trouvé</p>
                  <p class="text-xs text-slate-400 dark:text-slate-600 mb-4">Modifiez vos critères de recherche ou de filtre.</p>
                  <button @click="resetFilters" class="text-xs text-teal-700 dark:text-teal-400 font-semibold hover:underline">
                    Réinitialiser les filtres
                  </button>
                </td>
              </tr>

              <!-- Skeleton chargement -->
              <template v-if="isLoading">
                <tr v-for="n in 6" :key="`skel-${n}`" class="animate-pulse">
                  <td class="px-5 py-3.5">
                    <div class="flex items-center gap-3">
                      <div class="w-9 h-9 rounded-lg bg-slate-200 dark:bg-slate-700 shrink-0"></div>
                      <div class="space-y-1.5">
                        <div class="h-3 w-32 bg-slate-200 dark:bg-slate-700 rounded"></div>
                        <div class="h-2 w-20 bg-slate-100 dark:bg-slate-800 rounded"></div>
                      </div>
                    </div>
                  </td>
                  <td class="px-4 py-3.5 hidden sm:table-cell"><div class="h-3 w-28 bg-slate-200 dark:bg-slate-700 rounded"></div></td>
                  <td class="px-4 py-3.5 hidden md:table-cell"><div class="h-3 w-10 bg-slate-200 dark:bg-slate-700 rounded"></div></td>
                  <td class="px-4 py-3.5 hidden lg:table-cell"><div class="h-3 w-24 bg-slate-200 dark:bg-slate-700 rounded"></div></td>
                  <td class="px-4 py-3.5 hidden xl:table-cell"><div class="h-3 w-20 bg-slate-200 dark:bg-slate-700 rounded"></div></td>
                  <td class="px-4 py-3.5"><div class="h-5 w-16 bg-slate-200 dark:bg-slate-700 rounded-full"></div></td>
                  <td class="px-4 py-3.5"><div class="h-3 w-8 bg-slate-200 dark:bg-slate-700 rounded ml-auto"></div></td>
                </tr>
              </template>

            </tbody>
          </table>
        </div>

        <!-- Pied de tableau -->
        <div
          v-if="!isLoading && filteredClients.length > 0"
          class="px-5 py-2.5 bg-slate-50 dark:bg-slate-800/50 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between"
        >
          <p class="text-xs text-slate-400">
            <span class="font-semibold text-slate-600 dark:text-slate-300">{{ filteredClients.length }}</span>
            dossier{{ filteredClients.length > 1 ? 's' : '' }}
            <span v-if="filteredClients.length < clients.length"> sur {{ clients.length }}</span>
          </p>
          <p class="text-[10px] text-slate-300 dark:text-slate-600">Léopard · {{ dateAujourdhui }}</p>
        </div>
      </div>

      <!-- ══════════════════════════════════════════════════════
           VUE CARTES
      ══════════════════════════════════════════════════════ -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3">

        <div
          v-for="client in filteredClients"
          :key="client.id"
          @click="viewClient(client.id)"
          :class="[{ 'opacity-55': client.dcd === 1 }, 'bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4 cursor-pointer hover:border-teal-400 dark:hover:border-teal-600 hover:shadow-md transition-all group flex flex-col gap-3']"
        >
          <!-- En-tête -->
          <div class="flex items-start justify-between gap-2">
            <div class="flex items-center gap-2.5">
              <div
                :class="getAvatarClass(client)"
                class="w-10 h-10 rounded-lg flex items-center justify-center text-white font-bold text-sm shadow-sm shrink-0 relative"
              >
                {{ getInitials(client) }}
                <span v-if="isBirthday(client)" class="absolute -top-1.5 -right-1.5 text-sm leading-none">🎂</span>
              </div>
              <div class="min-w-0">
                <p class="text-sm font-bold text-slate-800 dark:text-slate-100 group-hover:text-teal-700 dark:group-hover:text-teal-400 transition-colors leading-tight truncate">
                  {{ client.prenom }} {{ client.nom }}
                </p>
                <p class="text-[10px] font-mono text-slate-400 mt-0.5 flex items-center gap-1">
                  <ShieldCheck :size="9" class="text-teal-500 shrink-0" />
                  {{ client.no_dossier_leopard || '—' }}
                </p>
              </div>
            </div>
            <span
              :class="getStatusBadgeClass(client)"
              class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-bold shrink-0"
            >
              <span class="w-1.5 h-1.5 rounded-full" :class="getStatusDotClass(client)"></span>
              {{ getStatusText(client) }}
            </span>
          </div>

          <!-- Infos -->
          <div class="space-y-1.5 text-xs text-slate-500 dark:text-slate-400">
            <div v-if="client.date_naissance" class="flex items-center gap-2">
              <Cake :size="12" class="text-slate-400 shrink-0" />
              <span>{{ formatDateCourt(client.date_naissance) }} · <strong class="text-slate-600 dark:text-slate-300">{{ calcAge(client.date_naissance) }} ans</strong> · {{ getSexeLabel(client.sexe) }}</span>
            </div>
            <div v-if="client.telephone" class="flex items-center gap-2">
              <Phone :size="12" class="text-slate-400 shrink-0" />
              <span>{{ formatPhone(client.telephone) }}</span>
            </div>
            <div v-if="client.ville" class="flex items-center gap-2">
              <MapPin :size="12" class="text-slate-400 shrink-0" />
              <span>{{ client.ville }}<span v-if="client.province">, {{ client.province }}</span></span>
            </div>
            <div v-if="getEtablissement(client)" class="flex items-center gap-2">
              <Building2 :size="12" class="text-slate-400 shrink-0" />
              <span>{{ getEtablissement(client) }}</span>
            </div>
          </div>

          <!-- Note d'alerte -->
          <div
            v-if="client.note_fixe"
            class="px-2.5 py-1.5 rounded-lg bg-amber-50 dark:bg-amber-900/20 border-l-2 border-amber-400 flex items-start gap-1.5"
          >
            <AlertTriangle :size="11" class="text-amber-500 shrink-0 mt-0.5" />
            <p class="text-[11px] text-amber-700 dark:text-amber-300 line-clamp-2">{{ client.note_fixe }}</p>
          </div>

          <!-- Actions carte -->
          <div
            class="flex items-center justify-end gap-1 pt-2 border-t border-slate-100 dark:border-slate-800 opacity-0 group-hover:opacity-100 transition-opacity"
            @click.stop
          >
            <button
              @click.stop="viewClient(client.id)"
              class="flex items-center gap-1 px-2.5 py-1 rounded-lg text-[11px] font-semibold bg-teal-700 text-white hover:bg-teal-800 transition-colors"
            >
              <FolderOpen :size="12" />
              Ouvrir
            </button>
            <button
              @click.stop="confirmDelete(client)"
              class="p-1.5 rounded-lg hover:bg-red-100 dark:hover:bg-red-900/20 transition-colors"
            >
              <Trash2 :size="13" class="text-red-400" />
            </button>
          </div>
        </div>

        <!-- Vide -->
        <div
          v-if="filteredClients.length === 0 && !isLoading"
          class="sm:col-span-2 lg:col-span-3 xl:col-span-4 py-16 text-center"
        >
          <UserX :size="40" class="mx-auto text-slate-300 dark:text-slate-700 mb-3" />
          <p class="text-sm font-semibold text-slate-500 mb-1">Aucun dossier trouvé</p>
          <button @click="resetFilters" class="text-xs text-teal-700 font-semibold hover:underline">
            Réinitialiser les filtres
          </button>
        </div>

      </div>

    </main>

    <!-- ══════════════════════════════════════════════════════
         MODAL CONFIRMATION SUPPRESSION
    ══════════════════════════════════════════════════════ -->
    <Teleport to="body">
      <div
        v-if="clientASupprimer"
        class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4"
        @click.self="clientASupprimer = null"
      >
        <div class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-700 shadow-2xl w-full max-w-sm p-6 animate-fadeIn">
          <div class="flex items-center gap-3 mb-4">
            <div class="w-10 h-10 rounded-lg bg-red-100 dark:bg-red-900/30 flex items-center justify-center">
              <Trash2 :size="18" class="text-red-600 dark:text-red-400" />
            </div>
            <div>
              <h3 class="text-base font-bold text-slate-800 dark:text-slate-100">Supprimer le dossier ?</h3>
              <p class="text-xs text-slate-500 dark:text-slate-400">Cette action est irréversible.</p>
            </div>
          </div>
          <div class="bg-slate-50 dark:bg-slate-800 rounded-lg px-4 py-3 mb-5">
            <p class="text-sm font-semibold text-slate-700 dark:text-slate-300">
              {{ clientASupprimer?.prenom }} {{ clientASupprimer?.nom }}
            </p>
            <p class="text-xs font-mono text-slate-400 mt-0.5">{{ clientASupprimer?.no_dossier_leopard }}</p>
          </div>
          <div class="flex gap-2">
            <button
              @click="clientASupprimer = null"
              class="flex-1 px-4 py-2 rounded-lg text-sm font-semibold border border-slate-200 dark:border-slate-700 text-slate-700 dark:text-slate-300 hover:bg-slate-50 dark:hover:bg-slate-800 transition-all"
            >
              Annuler
            </button>
            <button
              @click="executeDelete"
              class="flex-1 px-4 py-2 rounded-lg text-sm font-semibold bg-red-600 hover:bg-red-700 text-white shadow-sm transition-all"
            >
              Supprimer définitivement
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Modal création -->
    <ClientForm
      v-if="showForm"
      @close="showForm = false"
      @created="onClientCreated"
    />

  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { GetClients, DeleteClient } from '../../wailsjs/go/main/App'
import ClientForm from '../components/clients/ClientForm.vue'
import { useClientStore } from '../stores/clientStore'
import {
  Users, UserPlus, UserCheck, UserX, Archive,
  Search, LayoutGrid, List, FolderOpen,
  Phone, Smartphone, MapPin, Building2,
  ShieldCheck, Trash2, RefreshCw, X, ChevronRight,
  AlertTriangle, CalendarPlus, TrendingUp, Cake, Loader2
} from 'lucide-vue-next'

// ─── État ───────────────────────────────────────────────
const router          = useRouter()
const store           = useClientStore()

const clients          = ref([])
const showForm         = ref(false)
const searchQuery      = ref('')
const filterStatut     = ref('tous')
const viewMode         = ref('table')
const sortBy           = ref('nom')
const isLoading        = ref(false)
const clientASupprimer = ref(null)
const searchInput      = ref(null)

// ─── Filtres rapides ────────────────────────────────────
const filtresRapides = [
  { value: 'tous',    label: 'Tous'    },
  { value: 'actif',   label: 'Actifs'  },
  { value: 'inactif', label: 'Archivés'},
  { value: 'dcd',     label: 'DCD'     },
]

// ─── Statistiques ───────────────────────────────────────
const stats = computed(() => {
  const total    = clients.value.length
  const actifs   = clients.value.filter(c => c.actif === 1 && c.dcd === 0).length
  const dcd      = clients.value.filter(c => c.dcd === 1).length
  const archives = clients.value.filter(c => c.actif === 0 && c.dcd === 0).length
  const now      = new Date()
  const nouveaux = clients.value.filter(c => {
    if (!c.created_at) return false
    const d = new Date(c.created_at)
    return d.getMonth() === now.getMonth() && d.getFullYear() === now.getFullYear()
  }).length
  return { total, actifs, archives, nouveaux, dcd }
})

// ─── Liste filtrée + triée ──────────────────────────────
const filteredClients = computed(() => {
  let list = [...clients.value]

  if (searchQuery.value.trim()) {
    const q = searchQuery.value.toLowerCase().trim()
    list = list.filter(c =>
      c.nom?.toLowerCase().includes(q) ||
      c.prenom?.toLowerCase().includes(q) ||
      c.no_dossier_leopard?.toLowerCase().includes(q) ||
      c.ville?.toLowerCase().includes(q) ||
      c.telephone?.includes(q) ||
      c.cellulaire?.includes(q)
    )
  }

  if (filterStatut.value === 'actif')   list = list.filter(c => c.actif === 1 && c.dcd === 0)
  if (filterStatut.value === 'inactif') list = list.filter(c => c.actif === 0 && c.dcd === 0)
  if (filterStatut.value === 'dcd')     list = list.filter(c => c.dcd === 1)

  list.sort((a, b) => {
    if (sortBy.value === 'nom')    return (a.nom || '').localeCompare(b.nom || '', 'fr')
    if (sortBy.value === 'prenom') return (a.prenom || '').localeCompare(b.prenom || '', 'fr')
    if (sortBy.value === 'recent') return new Date(b.created_at || 0) - new Date(a.created_at || 0)
    if (sortBy.value === 'age') {
      const ageA = typeof calcAge(a.date_naissance) === 'number' ? calcAge(a.date_naissance) : 0
      const ageB = typeof calcAge(b.date_naissance) === 'number' ? calcAge(b.date_naissance) : 0
      return ageB - ageA
    }
    return 0
  })

  return list
})

const dateAujourdhui = computed(() =>
  new Date().toLocaleDateString('fr-CA', { year: 'numeric', month: 'long', day: 'numeric' })
)

// ─── Helpers ────────────────────────────────────────────
const calcAge = (dateNaissance) => {
  if (!dateNaissance) return '?'
  const birth = new Date(dateNaissance)
  const today = new Date()
  let age = today.getFullYear() - birth.getFullYear()
  const m = today.getMonth() - birth.getMonth()
  if (m < 0 || (m === 0 && today.getDate() < birth.getDate())) age--
  return age
}

const isBirthday = (client) => {
  if (!client.date_naissance) return false
  const b = new Date(client.date_naissance)
  const t = new Date()
  return b.getMonth() === t.getMonth() && b.getDate() === t.getDate()
}

const formatDateCourt = (dateStr) => {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleDateString('fr-CA')
}

const formatPhone = (phone) => {
  if (!phone) return ''
  const d = phone.replace(/\D/g, '')
  if (d.length === 10) return `(${d.slice(0,3)}) ${d.slice(3,6)}-${d.slice(6)}`
  if (d.length === 11 && d[0] === '1') return `+1 (${d.slice(1,4)}) ${d.slice(4,7)}-${d.slice(7)}`
  return phone
}

const getInitials = (client) => {
  return `${client.prenom?.[0] || ''}${client.nom?.[0] || ''}`.toUpperCase() || '?'
}

const getSexeLabel = (sexe) => {
  if (sexe === 'M') return 'Homme'
  if (sexe === 'F') return 'Femme'
  return ''
}

const getEtablissement = (client) => {
  if (client.rpa_id)   return `RPA #${client.rpa_id}`
  if (client.chsld_id) return `CHSLD #${client.chsld_id}`
  if (client.ri_id)    return `RI #${client.ri_id}`
  return null
}

const getAvatarClass = (client) => {
  if (client.dcd === 1)  return 'bg-slate-400 dark:bg-slate-600'
  if (client.actif === 0) return 'bg-orange-400 dark:bg-orange-600'
  const palette = ['bg-teal-600','bg-blue-600','bg-indigo-600','bg-violet-600','bg-cyan-600','bg-slate-600']
  const idx = ((client.nom?.charCodeAt(0) || 0) + (client.prenom?.charCodeAt(0) || 0)) % palette.length
  return palette[idx]
}

const getStatusText = (client) => {
  if (client.dcd === 1)   return 'DCD'
  if (client.actif === 1) return 'Actif'
  return 'Archivé'
}

const getStatusBadgeClass = (client) => {
  if (client.dcd === 1)   return 'bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400'
  if (client.actif === 1) return 'bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-400'
  return 'bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-400'
}

const getStatusDotClass = (client) => {
  if (client.dcd === 1)   return 'bg-slate-400'
  if (client.actif === 1) return 'bg-emerald-500'
  return 'bg-orange-500'
}

// ─── Actions ────────────────────────────────────────────
const focusSearch = () => searchInput.value?.focus()
const toggleView  = () => { viewMode.value = viewMode.value === 'table' ? 'cards' : 'table' }
const resetFilters = () => { searchQuery.value = ''; filterStatut.value = 'tous' }

const loadClients = async () => {
  isLoading.value = true
  try {
    const result = await GetClients()
    clients.value = result || []
  } catch (err) {
    console.error('Erreur chargement clients:', err)
  } finally {
    isLoading.value = false
  }
}

const viewClient = (id) => {
  store.clearClient()
  router.push(`/app/clients/${id}`)
}

const quickCall = (client) => {
  const tel = client.telephone || client.cellulaire
  if (tel) window.open(`tel:${tel}`)
}

const confirmDelete   = (client) => { clientASupprimer.value = client }

const executeDelete = async () => {
  if (!clientASupprimer.value) return
  try {
    await DeleteClient(clientASupprimer.value.id)
    await loadClients()
  } catch (err) {
    console.error('Erreur suppression:', err)
  } finally {
    clientASupprimer.value = null
  }
}

const onClientCreated = async () => { await loadClients() }

// Raccourci Ctrl+K
const handleKeydown = (e) => {
  if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
    e.preventDefault()
    focusSearch()
  }
}

onMounted(() => {
  loadClients()
  document.addEventListener('keydown', handleKeydown)
})
onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<style scoped>
.animate-fadeIn {
  animation: fadeIn 0.18s ease-out;
}
@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.97) translateY(6px); }
  to   { opacity: 1; transform: scale(1) translateY(0); }
}
</style>