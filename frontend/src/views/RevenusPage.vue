<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 flex flex-col">

    <!-- ══════════════════════════════════════════════
         HEADER
    ══════════════════════════════════════════════ -->
    <header class="border-b border-slate-700/60 bg-slate-900/80 backdrop-blur-md sticky top-0 z-30">
      <div class="flex items-center justify-between px-6 py-3 gap-4">

        <div class="flex items-center gap-3">
          <div class="w-8 h-8 rounded-lg bg-teal-500/20 border border-teal-500/30 flex items-center justify-center flex-shrink-0">
            <TrendingUp class="w-4 h-4 text-teal-400" />
          </div>
          <div>
            <h1 class="text-sm font-semibold text-slate-100 leading-none">Revenus & Séances</h1>
            <p class="text-xs text-slate-500 mt-0.5">{{ revenusFiltres.length }} entrée{{ revenusFiltres.length > 1 ? 's' : '' }}</p>
          </div>
        </div>

        <!-- Filtres rapides -->
        <div class="flex-1 flex items-center gap-2 overflow-x-auto scrollbar-none">
          <button
            v-for="f in filtresRapides"
            :key="f.value"
            @click="filtreStatut = filtreStatut === f.value ? '' : f.value"
            :class="[
              'flex-shrink-0 flex items-center gap-1.5 px-3 py-1.5 rounded-full text-xs font-medium transition-all',
              filtreStatut === f.value
                ? 'bg-teal-500 text-white shadow-lg shadow-teal-500/25'
                : 'bg-slate-800 text-slate-400 border border-slate-700 hover:border-teal-500/40 hover:text-teal-300'
            ]"
          >
            <component :is="f.icon" class="w-3 h-3" />
            {{ f.label }}
          </button>
        </div>

        <div class="flex items-center gap-2 flex-shrink-0">
          <!-- Sélecteur mois/année -->
          <div class="flex items-center gap-1 bg-slate-800 border border-slate-700 rounded-lg px-2 py-1.5">
            <button @click="moisPrecedent" class="text-slate-400 hover:text-slate-200 transition-colors">
              <ChevronLeft class="w-3.5 h-3.5" />
            </button>
            <span class="text-xs font-medium text-slate-200 px-2 min-w-24 text-center">{{ moisAffichage }}</span>
            <button @click="moisSuivant" class="text-slate-400 hover:text-slate-200 transition-colors">
              <ChevronRight class="w-3.5 h-3.5" />
            </button>
          </div>
          <button
            @click="ouvrirFormulaire"
            class="flex items-center gap-1.5 px-4 py-2 bg-teal-500 hover:bg-teal-400 text-white rounded-lg text-sm font-medium transition-all shadow-lg shadow-teal-500/25"
          >
            <Plus class="w-4 h-4" />
            Nouvelle séance
          </button>
        </div>
      </div>
    </header>

    <!-- ══════════════════════════════════════════════
         BODY
    ══════════════════════════════════════════════ -->
    <div class="flex flex-1 overflow-hidden">

      <!-- SIDEBAR — liste des revenus -->
      <aside class="w-80 border-r border-slate-700/60 flex flex-col bg-slate-900/40 overflow-hidden flex-shrink-0">

        <!-- Recherche -->
        <div class="p-3 border-b border-slate-700/40">
          <div class="relative">
            <Search class="absolute left-2.5 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-slate-500" />
            <input
              v-model="recherche"
              type="text"
              placeholder="Client, service, description..."
              class="w-full pl-8 pr-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-xs text-slate-200 placeholder-slate-500 focus:outline-none focus:border-teal-500/60 transition-all"
            />
          </div>
        </div>

        <!-- Stats mois courant -->
        <div class="grid grid-cols-2 gap-2 p-3 border-b border-slate-700/40">
          <div class="bg-slate-800/50 rounded-lg p-2.5">
            <p class="text-xs text-slate-500">Revenus</p>
            <p class="text-lg font-bold text-teal-400">{{ formatCAD(totalMois) }}</p>
          </div>
          <div class="bg-slate-800/50 rounded-lg p-2.5">
            <p class="text-xs text-slate-500">En attente</p>
            <p class="text-lg font-bold text-amber-400">{{ formatCAD(totalEnAttente) }}</p>
          </div>
        </div>

        <!-- Liste revenus -->
        <div class="flex-1 overflow-y-auto py-2 space-y-0.5 px-2">
          <button
            v-for="rev in revenusFiltres"
            :key="rev.ID"
            @click="selectionner(rev)"
            :class="[
              'w-full text-left flex items-start gap-3 px-3 py-2.5 rounded-lg transition-all group',
              revenuSelectionne?.ID === rev.ID
                ? 'bg-teal-500/15 border border-teal-500/30'
                : 'hover:bg-slate-800/60 border border-transparent'
            ]"
          >
            <!-- Indicateur statut -->
            <div class="flex-shrink-0 mt-1">
              <div :class="['w-2 h-2 rounded-full', getStatutColor(rev.StatutPaiement)]" />
            </div>

            <div class="min-w-0 flex-1">
              <div class="flex items-start justify-between gap-1">
                <p class="text-xs font-semibold text-teal-400">{{ formatCAD(rev.MontantTotal) }}</p>
                <p class="text-xs text-slate-500 flex-shrink-0">{{ formatDate(rev.DateService) }}</p>
              </div>
              <p class="text-xs text-slate-300 truncate mt-0.5">{{ rev.Description || rev.Categorie }}</p>
              <div class="flex items-center gap-2 mt-1">
                <span :class="['text-xs px-1.5 py-0.5 rounded', getStatutBadge(rev.StatutPaiement)]">
                  {{ getStatutLabel(rev.StatutPaiement) }}
                </span>
                <span class="text-xs text-slate-600">{{ rev.ModePaiement || '—' }}</span>
              </div>
            </div>
          </button>

          <div v-if="revenusFiltres.length === 0" class="flex flex-col items-center justify-center py-16 text-slate-600">
            <Receipt class="w-8 h-8 mb-2 opacity-40" />
            <p class="text-xs">Aucun revenu ce mois</p>
          </div>
        </div>
      </aside>

      <!-- VIEWER / FORMULAIRE -->
      <main class="flex-1 overflow-y-auto">

        <!-- État vide -->
        <div v-if="!showForm && !revenuSelectionne" class="h-full flex flex-col items-center justify-center text-slate-600">
          <div class="w-16 h-16 rounded-2xl bg-slate-800/60 border border-slate-700/40 flex items-center justify-center mb-4">
            <TrendingUp class="w-7 h-7 opacity-40" />
          </div>
          <p class="text-sm font-medium text-slate-500">Sélectionnez une séance</p>
          <p class="text-xs text-slate-600 mt-1">ou enregistrez-en une nouvelle</p>
          <button
            @click="ouvrirFormulaire"
            class="mt-6 flex items-center gap-2 px-4 py-2 bg-teal-500/10 hover:bg-teal-500/20 border border-teal-500/30 text-teal-400 rounded-lg text-sm transition-all"
          >
            <Plus class="w-4 h-4" />
            Nouvelle séance
          </button>
        </div>

        <!-- ════════════════════════════════════════
             FORMULAIRE
        ════════════════════════════════════════ -->
        <div v-else-if="showForm" class="max-w-2xl mx-auto p-6">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-lg font-semibold text-slate-100">
              {{ formMode === 'create' ? 'Nouvelle séance / revenu' : 'Modifier' }}
            </h2>
            <button @click="fermerFormulaire" class="text-slate-500 hover:text-slate-300">
              <X class="w-5 h-5" />
            </button>
          </div>

          <div class="space-y-5">

            <!-- ── Qui et quand ── -->
            <section class="bg-slate-800/40 border border-slate-700/40 rounded-xl p-4">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <User class="w-3.5 h-3.5" /> Client & Date
              </h3>
              <div class="grid grid-cols-2 gap-3">
                <!-- Sélection client -->
                <div class="col-span-2">
                  <label class="block text-xs font-medium text-slate-400 mb-1">Client</label>
                  <div class="relative">
                    <Search class="absolute left-2.5 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-slate-500 pointer-events-none" />
                    <input
                      v-model="rechercheClient"
                      type="text"
                      placeholder="Rechercher un client..."
                      @focus="showClientDropdown = true"
                      @blur="onBlurClient"
                      class="w-full pl-8 pr-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 placeholder-slate-500 focus:outline-none focus:border-teal-500/60 transition-all"
                    />
                    <div
                      v-if="showClientDropdown && clientsFiltres.length > 0"
                      class="absolute z-20 top-full left-0 right-0 mt-1 bg-slate-800 border border-slate-700 rounded-lg shadow-2xl max-h-48 overflow-y-auto"
                    >
                      <button
                        v-for="client in clientsFiltres"
                        :key="client.ID"
                        @mousedown.prevent="selectionnerClient(client)"
                        class="w-full text-left px-3 py-2 text-sm text-slate-200 hover:bg-slate-700 transition-colors"
                      >
                        {{ client.Prenom }} {{ client.Nom }}
                      </button>
                    </div>
                    <!-- Client sélectionné -->
                    <div v-if="clientSelectionne" class="mt-1.5 flex items-center gap-2">
                      <div class="flex items-center gap-2 bg-teal-500/10 border border-teal-500/20 rounded-lg px-2.5 py-1">
                        <UserCheck class="w-3.5 h-3.5 text-teal-400" />
                        <span class="text-xs text-teal-300">{{ clientSelectionne.Prenom }} {{ clientSelectionne.Nom }}</span>
                        <button @click="deselectionnerClient" class="text-teal-500 hover:text-teal-300 ml-1">
                          <X class="w-3 h-3" />
                        </button>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Date service -->
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Date de service</label>
                  <input
                    type="date"
                    v-model="form.date_service"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  />
                </div>

                <!-- Catégorie -->
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Catégorie</label>
                  <select
                    v-model="form.categorie"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  >
                    <option v-for="c in categories" :key="c.value" :value="c.value">{{ c.label }}</option>
                  </select>
                </div>
              </div>
            </section>

            <!-- ── Service ── -->
            <section class="bg-slate-800/40 border border-slate-700/40 rounded-xl p-4">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <LayoutGrid class="w-3.5 h-3.5" /> Service
              </h3>

              <!-- Sélection depuis le catalogue -->
              <div class="grid grid-cols-2 gap-2 mb-3">
                <button
                  v-for="svc in servicesActifs"
                  :key="svc.ID"
                  @click="appliquerService(svc)"
                  :class="[
                    'text-left px-3 py-2.5 rounded-lg border text-xs transition-all',
                    form.service_id === svc.ID
                      ? 'bg-teal-500/20 border-teal-500/40 text-teal-300'
                      : 'bg-slate-800/50 border-slate-700/40 text-slate-400 hover:border-teal-500/30 hover:text-slate-200'
                  ]"
                >
                  <p class="font-medium truncate">{{ svc.Nom }}</p>
                  <p class="text-slate-500 mt-0.5">{{ formatTarifService(svc) }}</p>
                </button>
              </div>

              <!-- Ou saisie manuelle -->
              <div class="border-t border-slate-700/40 pt-3">
                <p class="text-xs text-slate-600 mb-2">ou saisie libre :</p>
                <div class="grid grid-cols-2 gap-3">
                  <div>
                    <label class="block text-xs font-medium text-slate-400 mb-1">Mode facturation</label>
                    <select
                      v-model="form.mode_facturation"
                      @change="onModeFactChange"
                      class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                    >
                      <option value="horaire">Horaire</option>
                      <option value="acte">À l'acte</option>
                      <option value="forfait">Forfait</option>
                      <option value="kilometrage">Kilométrage</option>
                      <option value="autre">Autre</option>
                    </select>
                  </div>

                  <!-- Durée (si horaire) -->
                  <div v-if="form.mode_facturation === 'horaire'">
                    <label class="block text-xs font-medium text-slate-400 mb-1">Durée (heures)</label>
                    <input
                      type="number"
                      v-model.number="form.duree_heures"
                      step="0.25"
                      min="0"
                      @input="calculerMontant"
                      class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                    />
                  </div>

                  <!-- KM (si kilométrage) -->
                  <div v-if="form.mode_facturation === 'kilometrage'">
                    <label class="block text-xs font-medium text-slate-400 mb-1">Kilomètres</label>
                    <input
                      type="number"
                      v-model.number="form.km"
                      step="1"
                      min="0"
                      @input="calculerMontant"
                      class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                    />
                  </div>

                  <!-- Taux horaire -->
                  <div v-if="form.mode_facturation === 'horaire'">
                    <label class="block text-xs font-medium text-slate-400 mb-1">Taux horaire ($)</label>
                    <input
                      type="number"
                      v-model.number="form.taux_horaire_applique"
                      step="0.01"
                      @input="calculerMontant"
                      class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                    />
                  </div>
                </div>
              </div>
            </section>

            <!-- ── Montants ── -->
            <section class="bg-slate-800/40 border border-slate-700/40 rounded-xl p-4">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <DollarSign class="w-3.5 h-3.5" /> Montants
              </h3>
              <div class="grid grid-cols-3 gap-3">
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Montant brut ($)</label>
                  <input
                    type="number"
                    v-model.number="form.montant_brut"
                    step="0.01"
                    @input="calculerTaxes"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  />
                </div>
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">TPS ($)</label>
                  <input
                    type="number"
                    v-model.number="form.montant_tps"
                    step="0.01"
                    @input="calculerTotal"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  />
                </div>
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">TVQ ($)</label>
                  <input
                    type="number"
                    v-model.number="form.montant_tvq"
                    step="0.01"
                    @input="calculerTotal"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  />
                </div>
              </div>

              <!-- Total -->
              <div class="mt-3 flex items-center justify-between bg-teal-500/10 border border-teal-500/20 rounded-lg px-4 py-3">
                <span class="text-sm text-teal-300 font-medium">Total</span>
                <span class="text-2xl font-bold text-teal-400">{{ formatCAD(form.montant_total) }}</span>
              </div>
            </section>

            <!-- ── Paiement ── -->
            <section class="bg-slate-800/40 border border-slate-700/40 rounded-xl p-4">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <CreditCard class="w-3.5 h-3.5" /> Paiement
              </h3>
              <div class="grid grid-cols-3 gap-3">
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Statut</label>
                  <select
                    v-model="form.statut_paiement"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  >
                    <option value="en_attente">En attente</option>
                    <option value="paye">Payé</option>
                    <option value="partiel">Partiel</option>
                    <option value="annule">Annulé</option>
                  </select>
                </div>
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Mode</label>
                  <select
                    v-model="form.mode_paiement"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  >
                    <option value="">—</option>
                    <option value="interac">Interac</option>
                    <option value="especes">Espèces</option>
                    <option value="cheque">Chèque</option>
                    <option value="virement">Virement</option>
                    <option value="carte">Carte</option>
                  </select>
                </div>
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Date paiement</label>
                  <input
                    type="date"
                    v-model="form.date_paiement"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 focus:outline-none focus:border-teal-500/60 transition-all"
                  />
                </div>
              </div>
            </section>

            <!-- ── Description & Notes ── -->
            <section class="bg-slate-800/40 border border-slate-700/40 rounded-xl p-4">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <FileText class="w-3.5 h-3.5" /> Description
              </h3>
              <div class="space-y-3">
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Description (ligne de facture)</label>
                  <input
                    type="text"
                    v-model="form.description"
                    placeholder="Ex: Consultation individuelle — suivi semaine 3"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 placeholder-slate-500 focus:outline-none focus:border-teal-500/60 transition-all"
                  />
                </div>
                <div>
                  <label class="block text-xs font-medium text-slate-400 mb-1">Notes internes</label>
                  <textarea
                    v-model="form.notes"
                    rows="2"
                    placeholder="Notes privées (non visibles sur la facture)"
                    class="w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 placeholder-slate-500 focus:outline-none focus:border-teal-500/60 transition-all resize-none"
                  />
                </div>
              </div>
            </section>

            <!-- Actions -->
            <div class="flex items-center justify-between pt-2">
              <button @click="fermerFormulaire" class="px-4 py-2 text-sm text-slate-400 hover:text-slate-200 transition-colors">
                Annuler
              </button>
              <div class="flex items-center gap-3">
                <span v-if="erreur" class="text-xs text-red-400">{{ erreur }}</span>
                <button
                  @click="sauvegarder"
                  :disabled="saving || !clientSelectionne"
                  class="flex items-center gap-2 px-5 py-2 bg-teal-500 hover:bg-teal-400 disabled:opacity-50 disabled:cursor-not-allowed text-white rounded-lg text-sm font-medium transition-all shadow-lg shadow-teal-500/20"
                >
                  <Loader2 v-if="saving" class="w-4 h-4 animate-spin" />
                  <Check v-else class="w-4 h-4" />
                  {{ formMode === 'create' ? 'Enregistrer' : 'Mettre à jour' }}
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- ════════════════════════════════════════
             FICHE DÉTAIL
        ════════════════════════════════════════ -->
        <div v-else-if="revenuSelectionne" class="max-w-2xl mx-auto p-6">

          <div class="flex items-start justify-between mb-6">
            <div>
              <p class="text-xs text-slate-500 mb-1">{{ formatDateLong(revenuSelectionne.DateService) }}</p>
              <h2 class="text-2xl font-bold text-teal-400">{{ formatCAD(revenuSelectionne.MontantTotal) }}</h2>
              <p class="text-sm text-slate-300 mt-1">{{ revenuSelectionne.Description || revenuSelectionne.Categorie }}</p>
            </div>
            <div class="flex items-center gap-2">
              <button
                @click="ouvrirEdition(revenuSelectionne)"
                class="flex items-center gap-1.5 px-3 py-1.5 bg-slate-700/60 hover:bg-slate-700 border border-slate-600/60 text-slate-300 rounded-lg text-xs transition-all"
              >
                <Pencil class="w-3.5 h-3.5" /> Modifier
              </button>
              <button
                @click="confirmerSuppression(revenuSelectionne.ID)"
                class="flex items-center gap-1.5 px-3 py-1.5 bg-red-500/10 hover:bg-red-500/20 border border-red-500/30 text-red-400 rounded-lg text-xs transition-all"
              >
                <Trash2 class="w-3.5 h-3.5" />
              </button>
            </div>
          </div>

          <!-- Grille infos -->
          <div class="grid grid-cols-2 gap-3 mb-5">
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-3">
              <p class="text-xs text-slate-500 mb-1">Statut paiement</p>
              <span :class="['text-sm font-medium px-2 py-1 rounded-lg', getStatutBadge(revenuSelectionne.StatutPaiement)]">
                {{ getStatutLabel(revenuSelectionne.StatutPaiement) }}
              </span>
            </div>
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-3">
              <p class="text-xs text-slate-500 mb-1">Mode paiement</p>
              <p class="text-sm font-medium text-slate-200">{{ revenuSelectionne.ModePaiement || '—' }}</p>
            </div>
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-3">
              <p class="text-xs text-slate-500 mb-1">Catégorie</p>
              <p class="text-sm font-medium text-slate-200">{{ revenuSelectionne.Categorie }}</p>
            </div>
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-3">
              <p class="text-xs text-slate-500 mb-1">Mode facturation</p>
              <p class="text-sm font-medium text-slate-200">{{ revenuSelectionne.ModeFacturation }}</p>
            </div>
          </div>

          <!-- Détail montants -->
          <div class="bg-slate-800/30 border border-slate-700/30 rounded-xl p-4 mb-5">
            <p class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3">Détail montants</p>
            <div class="space-y-2">
              <div class="flex justify-between text-sm">
                <span class="text-slate-400">Montant brut</span>
                <span class="text-slate-200">{{ formatCAD(revenuSelectionne.MontantBrut) }}</span>
              </div>
              <div v-if="revenuSelectionne.MontantTPS > 0" class="flex justify-between text-sm">
                <span class="text-slate-400">TPS (5%)</span>
                <span class="text-slate-200">{{ formatCAD(revenuSelectionne.MontantTPS) }}</span>
              </div>
              <div v-if="revenuSelectionne.MontantTVQ > 0" class="flex justify-between text-sm">
                <span class="text-slate-400">TVQ (9.975%)</span>
                <span class="text-slate-200">{{ formatCAD(revenuSelectionne.MontantTVQ) }}</span>
              </div>
              <div class="flex justify-between text-sm font-semibold pt-2 border-t border-slate-700/40">
                <span class="text-teal-300">Total</span>
                <span class="text-teal-400">{{ formatCAD(revenuSelectionne.MontantTotal) }}</span>
              </div>
            </div>
          </div>

          <!-- Durée si horaire -->
          <div v-if="revenuSelectionne.DureeHeures > 0" class="bg-slate-800/30 border border-slate-700/30 rounded-xl p-4 mb-5">
            <div class="flex justify-between text-sm">
              <span class="text-slate-400">Durée</span>
              <span class="text-slate-200">{{ revenuSelectionne.DureeHeures }}h à {{ formatCAD(revenuSelectionne.TauxHoraireApplique) }}/h</span>
            </div>
          </div>

          <!-- Notes -->
          <div v-if="revenuSelectionne.Notes" class="bg-sky-500/5 border border-sky-500/20 rounded-xl p-4">
            <p class="text-xs text-sky-400 mb-2 flex items-center gap-1.5">
              <StickyNote class="w-3.5 h-3.5" /> Notes internes
            </p>
            <p class="text-sm text-slate-300">{{ revenuSelectionne.Notes }}</p>
          </div>

          <!-- Bouton générer facture -->
          <div class="mt-6 pt-6 border-t border-slate-700/40">
            <button
              class="w-full flex items-center justify-center gap-2 px-4 py-3 bg-slate-700/50 hover:bg-slate-700 border border-slate-600/60 text-slate-300 hover:text-slate-100 rounded-xl text-sm font-medium transition-all"
            >
              <FileText class="w-4 h-4" />
              Générer une facture depuis cette séance
            </button>
          </div>
        </div>
      </main>
    </div>

    <!-- ══════════════════════════════════════════════
         FOOTER — résumé mois
    ══════════════════════════════════════════════ -->
    <footer class="border-t border-slate-700/60 bg-slate-900/80 backdrop-blur-md px-6 py-3">
      <div class="flex items-center gap-6 overflow-x-auto scrollbar-none">
        <div v-for="stat in statsFooter" :key="stat.label" class="flex-shrink-0 flex items-center gap-2">
          <div :class="['w-1.5 h-1.5 rounded-full', stat.color]" />
          <span class="text-xs text-slate-500">{{ stat.label }}</span>
          <span class="text-xs font-semibold text-slate-300">{{ stat.valeur }}</span>
        </div>
      </div>
    </footer>

    <!-- Modal confirmation suppression -->
    <Teleport to="body">
      <div v-if="showConfirmDelete" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
        <div class="bg-slate-800 border border-slate-700 rounded-2xl p-6 w-80 shadow-2xl">
          <p class="text-sm font-semibold text-slate-100 mb-2">Supprimer ce revenu ?</p>
          <p class="text-xs text-slate-400 mb-5">Cette action est irréversible.</p>
          <div class="flex gap-3">
            <button @click="showConfirmDelete = false" class="flex-1 py-2 text-sm text-slate-400 hover:text-slate-200">Annuler</button>
            <button @click="supprimerConfirme" class="flex-1 py-2 bg-red-500 hover:bg-red-400 text-white rounded-lg text-sm font-medium transition-all">Supprimer</button>
          </div>
        </div>
      </div>
    </Teleport>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import {
  TrendingUp, Plus, Search, X, ChevronLeft, ChevronRight,
  User, UserCheck, LayoutGrid, DollarSign, CreditCard, FileText,
  Check, Loader2, Pencil, Trash2, Receipt, StickyNote
} from 'lucide-vue-next'
import {
  GetAllRevenus, CreateRevenu, UpdateRevenu, DeleteRevenu,
  GetClients, GetAllServices, GetParametresFinance
} from '../../wailsjs/go/main/App'

// ── État ─────────────────────────────────────────────────────────────────────
const revenus      = ref([])
const clients      = ref([])
const services     = ref([])
const parametres   = ref(null)

const revenuSelectionne  = ref(null)
const clientSelectionne  = ref(null)
const showForm           = ref(false)
const formMode           = ref('create')
const saving             = ref(false)
const erreur             = ref('')
const showConfirmDelete  = ref(false)
const idASupprimer       = ref(null)

// Filtres
const recherche      = ref('')
const rechercheClient = ref('')
const showClientDropdown = ref(false)
const filtreStatut   = ref('')
const moisCourant    = ref(new Date().getMonth() + 1)
const anneeCourante  = ref(new Date().getFullYear())

// ── Formulaire vide ───────────────────────────────────────────────────────────
const formVide = () => ({
  client_id:            null,
  service_id:           null,
  date_service:         new Date().toISOString().split('T')[0],
  categorie:            'consultation',
  mode_facturation:     'horaire',
  statut_paiement:      'en_attente',
  mode_paiement:        '',
  date_paiement:        '',
  duree_heures:         1,
  taux_horaire_applique: 0,
  montant_brut:         0,
  montant_tps:          0,
  montant_tvq:          0,
  montant_total:        0,
  description:          '',
  notes:                '',
  reference_paiement:   null,
  km:                   0,
})
const form = ref(formVide())

// ── Options ───────────────────────────────────────────────────────────────────
const categories = [
  { value: 'consultation', label: 'Consultation' },
  { value: 'evaluation',   label: 'Évaluation' },
  { value: 'rapport',      label: 'Rapport' },
  { value: 'forfait',      label: 'Forfait' },
  { value: 'atelier',      label: 'Atelier / Groupe' },
  { value: 'deplacement',  label: 'Déplacement' },
  { value: 'autre',        label: 'Autre' },
]

const filtresRapides = [
  { value: 'en_attente', label: 'En attente', icon: Receipt },
  { value: 'paye',       label: 'Payés',      icon: Check },
  { value: 'partiel',    label: 'Partiels',   icon: CreditCard },
]

// ── Computed ──────────────────────────────────────────────────────────────────
const moisAffichage = computed(() => {
  const d = new Date(anneeCourante.value, moisCourant.value - 1, 1)
  return d.toLocaleDateString('fr-CA', { month: 'long', year: 'numeric' })
})

const revenusFiltres = computed(() => {
  const prefix = `${anneeCourante.value}-${String(moisCourant.value).padStart(2, '0')}`
  return revenus.value.filter(r => {
    if (!r.DateService?.startsWith(prefix)) return false
    if (filtreStatut.value && r.StatutPaiement !== filtreStatut.value) return false
    if (recherche.value) {
      const q = recherche.value.toLowerCase()
      return r.Description?.toLowerCase().includes(q) ||
             r.Categorie?.toLowerCase().includes(q)
    }
    return true
  })
})

const servicesActifs = computed(() => services.value.filter(s => s.Actif))

const clientsFiltres = computed(() => {
  if (!rechercheClient.value) return clients.value.slice(0, 8)
  const q = rechercheClient.value.toLowerCase()
  return clients.value.filter(c =>
    c.Nom?.toLowerCase().includes(q) || c.Prenom?.toLowerCase().includes(q)
  ).slice(0, 8)
})

const totalMois = computed(() =>
  revenusFiltres.value
    .filter(r => r.StatutPaiement !== 'annule')
    .reduce((s, r) => s + (r.MontantTotal || 0), 0)
)

const totalEnAttente = computed(() =>
  revenusFiltres.value
    .filter(r => r.StatutPaiement === 'en_attente')
    .reduce((s, r) => s + (r.MontantTotal || 0), 0)
)

const statsFooter = computed(() => [
  { label: 'Séances',    valeur: revenusFiltres.value.length,                            color: 'bg-slate-500' },
  { label: 'Facturé',   valeur: formatCAD(totalMois.value),                              color: 'bg-teal-500' },
  { label: 'En attente',valeur: formatCAD(totalEnAttente.value),                         color: 'bg-amber-500' },
  { label: 'Payé',      valeur: formatCAD(totalMois.value - totalEnAttente.value),       color: 'bg-emerald-500' },
])

// ── Calculs automatiques ──────────────────────────────────────────────────────
const calculerMontant = () => {
  if (form.value.mode_facturation === 'horaire') {
    form.value.montant_brut = Math.round(
      form.value.duree_heures * form.value.taux_horaire_applique * 100
    ) / 100
  } else if (form.value.mode_facturation === 'kilometrage' && parametres.value) {
    form.value.montant_brut = Math.round(
      form.value.km * parametres.value.TauxKM * 100
    ) / 100
  }
  calculerTaxes()
}

const calculerTaxes = () => {
  // Par défaut exempte (services de santé)
  // Les taxes s'activent manuellement si applicable
  calculerTotal()
}

const calculerTotal = () => {
  form.value.montant_total = Math.round(
    (form.value.montant_brut + form.value.montant_tps + form.value.montant_tvq) * 100
  ) / 100
}

// ── Application d'un service du catalogue ────────────────────────────────────
const appliquerService = (svc) => {
  form.value.service_id = svc.ID
  form.value.categorie  = svc.Categorie

  if (svc.ModeTarification === 'horaire') {
    form.value.mode_facturation      = 'horaire'
    form.value.taux_horaire_applique = svc.TauxHoraire || 0
    form.value.duree_heures          = (svc.DureeMinutes || 60) / 60
    form.value.description           = svc.DescriptionCourte || svc.Nom
    calculerMontant()
  } else if (['acte', 'forfait'].includes(svc.ModeTarification)) {
    form.value.mode_facturation = svc.ModeTarification
    form.value.montant_brut     = svc.TarifUnitaire || 0
    form.value.description      = svc.DescriptionCourte || svc.Nom
    calculerTaxes()
  }
}

const onModeFactChange = () => {
  form.value.montant_brut = 0
  form.value.montant_total = 0
  form.value.service_id = null
}

// ── Clients ───────────────────────────────────────────────────────────────────
const selectionnerClient = (client) => {
  clientSelectionne.value  = client
  rechercheClient.value    = ''
  showClientDropdown.value = false
  form.value.client_id     = client.ID
}

const deselectionnerClient = () => {
  clientSelectionne.value  = null
  form.value.client_id     = null
  rechercheClient.value    = ''
}

const onBlurClient = () => {
  setTimeout(() => { showClientDropdown.value = false }, 150)
}

// ── Navigation mois ───────────────────────────────────────────────────────────
const moisPrecedent = () => {
  if (moisCourant.value === 1) { moisCourant.value = 12; anneeCourante.value-- }
  else moisCourant.value--
  charger()
}

const moisSuivant = () => {
  if (moisCourant.value === 12) { moisCourant.value = 1; anneeCourante.value++ }
  else moisCourant.value++
  charger()
}

// ── Helpers visuels ───────────────────────────────────────────────────────────
const formatCAD = (n) => {
  if (!n) return '0,00 $'
  return n.toLocaleString('fr-CA', { style: 'currency', currency: 'CAD' })
}

const formatDate = (d) => {
  if (!d) return '—'
  return new Date(d).toLocaleDateString('fr-CA', { month: 'short', day: 'numeric' })
}

const formatDateLong = (d) => {
  if (!d) return '—'
  return new Date(d).toLocaleDateString('fr-CA', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
}

const formatTarifService = (svc) => {
  if (svc.ModeTarification === 'horaire' && svc.TauxHoraire)
    return `${svc.TauxHoraire} $/h · ${svc.DureeMinutes || '?'} min`
  if (svc.TarifUnitaire) return formatCAD(svc.TarifUnitaire)
  if (svc.ModeTarification === 'gratuit') return 'Gratuit'
  return '—'
}

const getStatutColor = (s) => ({
  en_attente: 'bg-amber-400',
  paye:       'bg-emerald-400',
  partiel:    'bg-sky-400',
  annule:     'bg-slate-500',
}[s] || 'bg-slate-500')

const getStatutBadge = (s) => ({
  en_attente: 'bg-amber-500/15 text-amber-400',
  paye:       'bg-emerald-500/15 text-emerald-400',
  partiel:    'bg-sky-500/15 text-sky-400',
  annule:     'bg-slate-500/15 text-slate-400',
}[s] || 'bg-slate-500/15 text-slate-400')

const getStatutLabel = (s) => ({
  en_attente: 'En attente',
  paye:       'Payé',
  partiel:    'Partiel',
  annule:     'Annulé',
}[s] || s)

// ── CRUD ──────────────────────────────────────────────────────────────────────
const charger = async () => {
  try {
    const debut = `${anneeCourante.value}-${String(moisCourant.value).padStart(2, '0')}-01`
    const fin   = new Date(anneeCourante.value, moisCourant.value, 0)
      .toISOString().split('T')[0]
    const res = await GetAllRevenus(debut, fin, '', '')
    revenus.value = res || []
  } catch (e) {
    console.error('❌ Revenus:', e)
  }
}

const ouvrirFormulaire = () => {
  form.value      = formVide()
  clientSelectionne.value = null
  rechercheClient.value   = ''
  formMode.value  = 'create'
  erreur.value    = ''
  revenuSelectionne.value = null
  showForm.value  = true
}

const ouvrirEdition = (rev) => {
  form.value = {
    client_id:            rev.ClientID,
    service_id:           rev.ServiceID,
    date_service:         rev.DateService,
    categorie:            rev.Categorie,
    mode_facturation:     rev.ModeFacturation,
    statut_paiement:      rev.StatutPaiement,
    mode_paiement:        rev.ModePaiement || '',
    date_paiement:        rev.DatePaiement || '',
    duree_heures:         rev.DureeHeures,
    taux_horaire_applique: rev.TauxHoraireApplique,
    montant_brut:         rev.MontantBrut,
    montant_tps:          rev.MontantTPS,
    montant_tvq:          rev.MontantTVQ,
    montant_total:        rev.MontantTotal,
    description:          rev.Description || '',
    notes:                rev.Notes || '',
    reference_paiement:   null,
    km:                   0,
    id:                   rev.ID,
  }
  formMode.value  = 'edit'
  showForm.value  = true
  revenuSelectionne.value = null
}

const fermerFormulaire = () => {
  showForm.value = false
  erreur.value   = ''
}

const selectionner = (rev) => {
  revenuSelectionne.value = rev
  showForm.value = false
}

const sauvegarder = async () => {
  if (!form.value.client_id) { erreur.value = 'Veuillez sélectionner un client.'; return }
  saving.value = true
  erreur.value = ''
  try {
    const payload = {
      client_id:             form.value.client_id,
      service_id:            form.value.service_id,
      date_service:          form.value.date_service,
      categorie:             form.value.categorie,
      mode_facturation:      form.value.mode_facturation,
      statut_paiement:       form.value.statut_paiement,
      mode_paiement:         form.value.mode_paiement,
      date_paiement:         form.value.date_paiement,
      duree_heures:          form.value.duree_heures,
      taux_horaire_applique: form.value.taux_horaire_applique,
      montant_brut:          form.value.montant_brut,
      montant_tps:           form.value.montant_tps,
      montant_tvq:           form.value.montant_tvq,
      montant_total:         form.value.montant_total,
      description:           form.value.description,
      notes:                 form.value.notes,
      reference_paiement:    null,
    }
    if (formMode.value === 'create') {
      await CreateRevenu(payload)
    } else {
      await UpdateRevenu({ id: form.value.id, ...payload })
    }
    await charger()
    fermerFormulaire()
  } catch (e) {
    erreur.value = e.toString()
  } finally {
    saving.value = false
  }
}

const confirmerSuppression = (id) => {
  idASupprimer.value  = id
  showConfirmDelete.value = true
}

const supprimerConfirme = async () => {
  try {
    await DeleteRevenu(idASupprimer.value)
    await charger()
    revenuSelectionne.value = null
  } catch (e) {
    console.error(e)
  } finally {
    showConfirmDelete.value = false
    idASupprimer.value = null
  }
}

// ── Init ──────────────────────────────────────────────────────────────────────
onMounted(async () => {
  await charger()
  try {
    const [c, s, p] = await Promise.all([
      GetClients(),
      GetAllServices('', true),
      GetParametresFinance(),
    ])
    clients.value    = c || []
    services.value   = s || []
    parametres.value = p
    // Pré-remplir le taux horaire par défaut
    if (p?.TauxHoraireDefaut) {
      form.value.taux_horaire_applique = p.TauxHoraireDefaut
    }
  } catch (e) {
    console.error('❌ Init RevenuPage:', e)
  }
})
</script>

<style scoped>
.scrollbar-none::-webkit-scrollbar { display: none; }
.scrollbar-none { -ms-overflow-style: none; scrollbar-width: none; }
</style>