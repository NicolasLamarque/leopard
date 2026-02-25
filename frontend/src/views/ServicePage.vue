<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 flex flex-col">

    <!-- ══════════════════════════════════════════════
         HEADER — Actions CRUD + Export
    ══════════════════════════════════════════════ -->
    <header class="border-b border-slate-700/60 bg-slate-900/80 backdrop-blur-md sticky top-0 z-30">
      <div class="flex items-center justify-between px-6 py-3 gap-4">

        <!-- Titre + breadcrumb -->
        <div class="flex items-center gap-3 min-w-0">
          <div class="w-8 h-8 rounded-lg bg-teal-500/20 border border-teal-500/30 flex items-center justify-center flex-shrink-0">
            <LayoutGrid class="w-4 h-4 text-teal-400" />
          </div>
          <div class="min-w-0">
            <h1 class="text-sm font-semibold text-slate-100 leading-none">Catalogue de services</h1>
            <p class="text-xs text-slate-500 mt-0.5">{{ servicesActifs.length }} service{{ servicesActifs.length > 1 ? 's' : '' }} actif{{ servicesActifs.length > 1 ? 's' : '' }}</p>
          </div>
        </div>

        <!-- Filtre catégorie rapide -->
        <div class="flex-1 flex items-center gap-2 overflow-x-auto scrollbar-none">
          <button
            v-for="cat in categoriesDisponibles"
            :key="cat.value"
            @click="filtreCategorie = filtreCategorie === cat.value ? '' : cat.value"
            :class="[
              'flex-shrink-0 flex items-center gap-1.5 px-3 py-1.5 rounded-full text-xs font-medium transition-all duration-200',
              filtreCategorie === cat.value
                ? 'bg-teal-500 text-white shadow-lg shadow-teal-500/25'
                : 'bg-slate-800 text-slate-400 border border-slate-700 hover:border-teal-500/40 hover:text-teal-300'
            ]"
          >
            <component :is="cat.icon" class="w-3 h-3" />
            {{ cat.label }}
          </button>
        </div>

        <!-- Actions -->
        <div class="flex items-center gap-2 flex-shrink-0">
          <button
            @click="toggleArchives"
            :class="[
              'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium transition-all',
              showArchives ? 'bg-slate-600 text-slate-200' : 'text-slate-400 hover:text-slate-200'
            ]"
          >
            <Archive class="w-3.5 h-3.5" />
            Archives
          </button>
          <button
            @click="ouvrirFormulaire(null)"
            class="flex items-center gap-1.5 px-4 py-2 bg-teal-500 hover:bg-teal-400 text-white rounded-lg text-sm font-medium transition-all shadow-lg shadow-teal-500/25 hover:shadow-teal-400/30"
          >
            <Plus class="w-4 h-4" />
            Nouveau service
          </button>
        </div>
      </div>
    </header>

    <!-- ══════════════════════════════════════════════
         BODY — Sidebar + Viewer
    ══════════════════════════════════════════════ -->
    <div class="flex flex-1 overflow-hidden">

      <!-- SIDEBAR — liste des services -->
      <aside class="w-72 border-r border-slate-700/60 flex flex-col bg-slate-900/40 overflow-hidden flex-shrink-0">

        <!-- Recherche -->
        <div class="p-3 border-b border-slate-700/40">
          <div class="relative">
            <Search class="absolute left-2.5 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-slate-500" />
            <input
              v-model="recherche"
              type="text"
              placeholder="Rechercher..."
              class="w-full pl-8 pr-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-xs text-slate-200 placeholder-slate-500 focus:outline-none focus:border-teal-500/60 focus:ring-1 focus:ring-teal-500/20 transition-all"
            />
          </div>
        </div>

        <!-- Groupes par catégorie -->
        <div class="flex-1 overflow-y-auto py-2 space-y-1 px-2">
          <template v-for="groupe in servicesGroupes" :key="groupe.Categorie">
            <div v-if="groupe.services.length > 0">
              <!-- En-tête groupe -->
              <div class="flex items-center gap-2 px-2 py-1.5 mb-1">
                <component :is="getCatIcon(groupe.Categorie)" class="w-3 h-3 text-slate-500" />
                <span class="text-xs font-semibold text-slate-500 uppercase tracking-wider">{{ getCatLabel(groupe.Categorie) }}</span>
                <span class="ml-auto text-xs text-slate-600">{{ groupe.services.length }}</span>
              </div>

              <!-- Items -->
              <button
                v-for="svc in groupe.services"
                :key="svc.ID"
                @click="selectionner(svc)"
                :class="[
                  'w-full text-left flex items-start gap-3 px-3 py-2.5 rounded-lg transition-all group',
                  serviceSelectionne?.ID === svc.ID
                    ? 'bg-teal-500/15 border border-teal-500/30'
                    : 'hover:bg-slate-800/60 border border-transparent'
                ]"
              >
                <div :class="['w-7 h-7 rounded-md flex-shrink-0 flex items-center justify-center mt-0.5', getCatBg(svc.Categorie)]">
                  <component :is="getCatIcon(svc.Categorie)" :class="['w-3.5 h-3.5', getCatColor(svc.Categorie)]" />
                </div>
                <div class="min-w-0 flex-1">
                  <div class="flex items-start justify-between gap-1">
                    <p class="text-xs font-medium text-slate-200 leading-snug truncate">{{ svc.Nom }}</p>
                    <span v-if="!svc.Actif" class="flex-shrink-0 text-xs text-slate-600 italic">archivé</span>
                  </div>
                  <p class="text-xs text-slate-500 mt-0.5">
                    {{ formatTarif(svc) }}
                  </p>
                  <span class="inline-block mt-1 px-1.5 py-0.5 rounded text-xs font-mono text-slate-600 bg-slate-800/80">{{ svc.Code }}</span>
                </div>
              </button>
            </div>
          </template>

          <!-- Vide -->
          <div v-if="servicesFiltres.length === 0" class="flex flex-col items-center justify-center py-16 text-slate-600">
            <PackageSearch class="w-8 h-8 mb-2 opacity-40" />
            <p class="text-xs">Aucun service trouvé</p>
          </div>
        </div>
      </aside>

      <!-- VIEWER — détail du service sélectionné -->
      <main class="flex-1 overflow-y-auto">

        <!-- État vide -->
        <div v-if="!serviceSelectionne && !showForm" class="h-full flex flex-col items-center justify-center text-slate-600">
          <div class="w-16 h-16 rounded-2xl bg-slate-800/60 border border-slate-700/40 flex items-center justify-center mb-4">
            <LayoutGrid class="w-7 h-7 opacity-40" />
          </div>
          <p class="text-sm font-medium text-slate-500">Sélectionnez un service</p>
          <p class="text-xs text-slate-600 mt-1">ou créez-en un nouveau</p>
          <button
            @click="ouvrirFormulaire(null)"
            class="mt-6 flex items-center gap-2 px-4 py-2 bg-teal-500/10 hover:bg-teal-500/20 border border-teal-500/30 text-teal-400 rounded-lg text-sm transition-all"
          >
            <Plus class="w-4 h-4" />
            Nouveau service
          </button>
        </div>

        <!-- FORMULAIRE CRÉATION / ÉDITION -->
        <div v-else-if="showForm" class="max-w-2xl mx-auto p-6">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-lg font-semibold text-slate-100">
              {{ formMode === 'create' ? 'Nouveau service' : 'Modifier le service' }}
            </h2>
            <button @click="fermerFormulaire" class="text-slate-500 hover:text-slate-300 transition-colors">
              <X class="w-5 h-5" />
            </button>
          </div>

          <FormKit type="form" @submit="sauvegarder" :actions="false">

            <!-- Identification -->
            <section class="mb-6">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <Tag class="w-3.5 h-3.5" /> Identification
              </h3>
              <div class="grid grid-cols-2 gap-3">
                <FormKit
                  type="text"
                  name="code"
                  label="Code"
                  placeholder="CONS-IND-60"
                  v-model="form.code"
                  validation="required"
                  outer-class="col-span-1"
                  :classes="fkClasses"
                />
                <FormKit
                  type="select"
                  name="categorie"
                  label="Catégorie"
                  v-model="form.categorie"
                  :options="categoriesOptions"
                  :classes="fkClasses"
                />
                <FormKit
                  type="text"
                  name="nom"
                  label="Nom du service"
                  placeholder="Consultation individuelle — 60 min"
                  v-model="form.nom"
                  validation="required"
                  outer-class="col-span-2"
                  :classes="fkClasses"
                />
                <FormKit
                  type="select"
                  name="type_livraison"
                  label="Type de livraison"
                  v-model="form.type_livraison"
                  :options="typesLivraison"
                  :classes="fkClasses"
                />
                <FormKit
                  type="select"
                  name="icone"
                  label="Icône"
                  v-model="form.icone"
                  :options="iconesOptions"
                  :classes="fkClasses"
                />
              </div>
            </section>

            <!-- Tarification -->
            <section class="mb-6">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <DollarSign class="w-3.5 h-3.5" /> Tarification
              </h3>
              <div class="grid grid-cols-2 gap-3">
                <FormKit
                  type="select"
                  name="mode_tarification"
                  label="Mode"
                  v-model="form.mode_tarification"
                  :options="modesTarification"
                  :classes="fkClasses"
                  @change="onModeTarifChange"
                />
                <FormKit
                  v-if="form.mode_tarification === 'horaire'"
                  type="number"
                  name="duree_minutes"
                  label="Durée (min)"
                  placeholder="60"
                  v-model.number="form.duree_minutes"
                  :classes="fkClasses"
                />
                <FormKit
                  v-if="form.mode_tarification === 'horaire'"
                  type="number"
                  name="taux_horaire"
                  label="Taux horaire ($)"
                  placeholder="120.00"
                  step="0.01"
                  v-model.number="form.taux_horaire"
                  :classes="fkClasses"
                />
                <FormKit
                  v-if="['acte','forfait','abonnement'].includes(form.mode_tarification)"
                  type="number"
                  name="tarif_unitaire"
                  label="Tarif ($)"
                  placeholder="350.00"
                  step="0.01"
                  v-model.number="form.tarif_unitaire"
                  outer-class="col-span-2"
                  :classes="fkClasses"
                />
                <FormKit
                  v-if="form.mode_tarification === 'prix_libre'"
                  type="number"
                  name="tarif_min"
                  label="Tarif min ($)"
                  step="0.01"
                  v-model.number="form.tarif_min"
                  :classes="fkClasses"
                />
                <FormKit
                  v-if="form.mode_tarification === 'prix_libre'"
                  type="number"
                  name="tarif_max"
                  label="Tarif max ($)"
                  step="0.01"
                  v-model.number="form.tarif_max"
                  :classes="fkClasses"
                />
              </div>
            </section>

            <!-- Fiscal -->
            <section class="mb-6">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <Receipt class="w-3.5 h-3.5" /> Fiscal
              </h3>
              <div class="flex items-center gap-6">
                <FormKit
                  type="checkbox"
                  name="exempte_taxes"
                  label="Exempté TPS/TVQ (soins de santé)"
                  v-model="form.exempte_taxes"
                  :classes="fkCheckboxClasses"
                />
                <template v-if="!form.exempte_taxes">
                  <FormKit type="checkbox" name="tps_applicable" label="TPS" v-model="form.tps_applicable" :classes="fkCheckboxClasses" />
                  <FormKit type="checkbox" name="tvq_applicable" label="TVQ" v-model="form.tvq_applicable" :classes="fkCheckboxClasses" />
                </template>
              </div>
            </section>

            <!-- Description -->
            <section class="mb-6">
              <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-2">
                <FileText class="w-3.5 h-3.5" /> Description
              </h3>
              <div class="space-y-3">
                <FormKit
                  type="text"
                  name="description_courte"
                  label="Description courte (factures)"
                  placeholder="Consultation individuelle"
                  v-model="form.description_courte"
                  :classes="fkClasses"
                />
                <FormKit
                  type="textarea"
                  name="description_longue"
                  label="Description longue (contrats)"
                  v-model="form.description_longue"
                  rows="3"
                  :classes="fkClasses"
                />
                <FormKit
                  type="text"
                  name="public_cible"
                  label="Public cible"
                  placeholder="Adultes 18+, troubles anxieux"
                  v-model="form.public_cible"
                  :classes="fkClasses"
                />
              </div>
            </section>

            <!-- Extensibilité future (collapsible) -->
            <section class="mb-6">
              <button
                type="button"
                @click="showFutur = !showFutur"
                class="flex items-center gap-2 text-xs font-semibold text-slate-500 uppercase tracking-wider hover:text-slate-300 transition-colors mb-3"
              >
                <Sparkles class="w-3.5 h-3.5" />
                Options avancées (ateliers, tracts, capsules...)
                <ChevronDown :class="['w-3.5 h-3.5 transition-transform', showFutur ? 'rotate-180' : '']" />
              </button>
              <div v-if="showFutur" class="grid grid-cols-2 gap-3 pt-1">
                <FormKit type="number" name="nb_places_max" label="Places max (ateliers)" v-model.number="form.nb_places_max" :classes="fkClasses" />
                <FormKit type="number" name="nb_seances_forfait" label="Séances (forfait)" v-model.number="form.nb_seances_forfait" :classes="fkClasses" />
                <FormKit type="text" name="format_tract" label="Format tract" placeholder="Trifold A4" v-model="form.format_tract" :classes="fkClasses" />
                <FormKit type="text" name="couleur_tract" label="Couleur tract" placeholder="N&B" v-model="form.couleur_tract" :classes="fkClasses" />
                <FormKit type="number" name="duree_video_minutes" label="Durée vidéo (min)" v-model.number="form.duree_video_minutes" :classes="fkClasses" />
                <FormKit type="text" name="url_ressource" label="URL ressource" v-model="form.url_ressource" :classes="fkClasses" />
                <FormKit type="text" name="tags" label="Tags (JSON)" placeholder='["TCC","adulte"]' v-model="form.tags" outer-class="col-span-2" :classes="fkClasses" />
              </div>
            </section>

            <!-- Notes internes -->
            <section class="mb-8">
              <FormKit
                type="textarea"
                name="notes_internes"
                label="Notes internes"
                v-model="form.notes_internes"
                rows="2"
                :classes="fkClasses"
              />
            </section>

            <!-- Actions formulaire -->
            <div class="flex items-center justify-between pt-4 border-t border-slate-700/60">
              <button
                type="button"
                @click="fermerFormulaire"
                class="px-4 py-2 text-sm text-slate-400 hover:text-slate-200 transition-colors"
              >
                Annuler
              </button>
              <div class="flex items-center gap-3">
                <span v-if="erreurSauvegarde" class="text-xs text-red-400">{{ erreurSauvegarde }}</span>
                <button
                  type="submit"
                  :disabled="saving"
                  class="flex items-center gap-2 px-5 py-2 bg-teal-500 hover:bg-teal-400 disabled:opacity-50 text-white rounded-lg text-sm font-medium transition-all shadow-lg shadow-teal-500/20"
                >
                  <Loader2 v-if="saving" class="w-4 h-4 animate-spin" />
                  <Check v-else class="w-4 h-4" />
                  {{ formMode === 'create' ? 'Créer' : 'Enregistrer' }}
                </button>
              </div>
            </div>
          </FormKit>
        </div>

        <!-- FICHE DÉTAIL du service sélectionné -->
        <div v-else-if="serviceSelectionne" class="p-6 max-w-2xl mx-auto">

          <!-- En-tête fiche -->
          <div class="flex items-start justify-between mb-6">
            <div class="flex items-start gap-4">
              <div :class="['w-12 h-12 rounded-xl flex items-center justify-center flex-shrink-0', getCatBg(serviceSelectionne.Categorie)]">
                <component :is="getCatIcon(serviceSelectionne.Categorie)" :class="['w-6 h-6', getCatColor(serviceSelectionne.Categorie)]" />
              </div>
              <div>
                <h2 class="text-xl font-semibold text-slate-100">{{ serviceSelectionne.Nom }}</h2>
                <div class="flex items-center gap-2 mt-1">
                  <span class="font-mono text-xs text-slate-500 bg-slate-800 px-2 py-0.5 rounded">{{ serviceSelectionne.Code }}</span>
                  <span :class="['text-xs px-2 py-0.5 rounded-full', getCatBg(serviceSelectionne.Categorie), getCatColor(serviceSelectionne.Categorie)]">
                    {{ getCatLabel(serviceSelectionne.Categorie) }}
                  </span>
                  <span v-if="!serviceSelectionne.Actif" class="text-xs px-2 py-0.5 rounded-full bg-slate-700 text-slate-400">Archivé</span>
                </div>
              </div>
            </div>
            <div class="flex items-center gap-2">
              <button
                @click="ouvrirFormulaire(serviceSelectionne)"
                class="flex items-center gap-1.5 px-3 py-1.5 bg-slate-700/60 hover:bg-slate-700 border border-slate-600/60 text-slate-300 rounded-lg text-xs transition-all"
              >
                <Pencil class="w-3.5 h-3.5" />
                Modifier
              </button>
              <button
                v-if="serviceSelectionne.Actif"
                @click="archiver(serviceSelectionne.ID)"
                class="flex items-center gap-1.5 px-3 py-1.5 bg-slate-700/60 hover:bg-slate-700 border border-slate-600/60 text-slate-400 rounded-lg text-xs transition-all"
              >
                <Archive class="w-3.5 h-3.5" />
                Archiver
              </button>
              <button
                @click="confirmerSuppression(serviceSelectionne.ID)"
                class="flex items-center gap-1.5 px-3 py-1.5 bg-red-500/10 hover:bg-red-500/20 border border-red-500/30 text-red-400 rounded-lg text-xs transition-all"
              >
                <Trash2 class="w-3.5 h-3.5" />
              </button>
            </div>
          </div>

          <!-- Grille infos -->
          <div class="grid grid-cols-2 gap-4 mb-6">

            <!-- Tarif -->
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-4">
              <p class="text-xs text-slate-500 mb-1 flex items-center gap-1.5"><DollarSign class="w-3 h-3" /> Tarification</p>
              <p class="text-2xl font-bold text-teal-400">{{ formatTarifDetail(serviceSelectionne) }}</p>
              <p class="text-xs text-slate-500 mt-1">{{ getModeLabel(serviceSelectionne.ModeTarification) }}</p>
            </div>

            <!-- Durée -->
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-4">
              <p class="text-xs text-slate-500 mb-1 flex items-center gap-1.5"><Clock class="w-3 h-3" /> Durée</p>
              <p class="text-2xl font-bold text-slate-200">
                {{ serviceSelectionne.DureeMinutes ? serviceSelectionne.DureeMinutes + ' min' : '—' }}
              </p>
              <p class="text-xs text-slate-500 mt-1">{{ serviceSelectionne.TypeLivraison }}</p>
            </div>

            <!-- Fiscal -->
            <div class="bg-slate-800/50 border border-slate-700/40 rounded-xl p-4 col-span-2">
              <p class="text-xs text-slate-500 mb-2 flex items-center gap-1.5"><Receipt class="w-3 h-3" /> Statut fiscal</p>
              <div class="flex items-center gap-3">
                <div :class="['flex items-center gap-1.5 px-2.5 py-1 rounded-lg text-xs font-medium',
                  serviceSelectionne.ExempteTaxes
                    ? 'bg-emerald-500/15 border border-emerald-500/30 text-emerald-400'
                    : 'bg-amber-500/15 border border-amber-500/30 text-amber-400'
                ]">
                  <component :is="serviceSelectionne.ExempteTaxes ? ShieldCheck : AlertTriangle" class="w-3.5 h-3.5" />
                  {{ serviceSelectionne.ExempteTaxes ? 'Exempté TPS/TVQ' : 'Taxable' }}
                </div>
                <span v-if="!serviceSelectionne.ExempteTaxes && serviceSelectionne.TPSApplicable" class="text-xs text-slate-400">TPS 5%</span>
                <span v-if="!serviceSelectionne.ExempteTaxes && serviceSelectionne.TVQApplicable" class="text-xs text-slate-400">TVQ 9.975%</span>
              </div>
            </div>
          </div>

          <!-- Description -->
          <div v-if="serviceSelectionne.DescriptionCourte || serviceSelectionne.DescriptionLongue" class="mb-6 space-y-3">
            <div v-if="serviceSelectionne.DescriptionCourte" class="bg-slate-800/30 border border-slate-700/30 rounded-lg p-3">
              <p class="text-xs text-slate-500 mb-1">Description courte (factures)</p>
              <p class="text-sm text-slate-300">{{ serviceSelectionne.DescriptionCourte }}</p>
            </div>
            <div v-if="serviceSelectionne.DescriptionLongue" class="bg-slate-800/30 border border-slate-700/30 rounded-lg p-3">
              <p class="text-xs text-slate-500 mb-1">Description longue (contrats)</p>
              <p class="text-sm text-slate-300 whitespace-pre-wrap">{{ serviceSelectionne.DescriptionLongue }}</p>
            </div>
          </div>

          <!-- Options avancées (si remplies) -->
          <div v-if="hasOptionsAvancees(serviceSelectionne)" class="bg-slate-800/30 border border-slate-700/30 rounded-xl p-4 mb-6">
            <p class="text-xs font-semibold text-slate-500 uppercase tracking-wider mb-3 flex items-center gap-1.5">
              <Sparkles class="w-3.5 h-3.5" /> Options avancées
            </p>
            <div class="grid grid-cols-2 gap-2 text-xs">
              <div v-if="serviceSelectionne.NbPlacesMax" class="flex justify-between text-slate-400">
                <span>Places max</span><span class="text-slate-200">{{ serviceSelectionne.NbPlacesMax }}</span>
              </div>
              <div v-if="serviceSelectionne.NbSeancesForfait" class="flex justify-between text-slate-400">
                <span>Séances forfait</span><span class="text-slate-200">{{ serviceSelectionne.NbSeancesForfait }}</span>
              </div>
              <div v-if="serviceSelectionne.FormatTract" class="flex justify-between text-slate-400">
                <span>Format tract</span><span class="text-slate-200">{{ serviceSelectionne.FormatTract }}</span>
              </div>
              <div v-if="serviceSelectionne.URLRessource" class="flex justify-between text-slate-400">
                <span>URL</span><span class="text-teal-400 truncate max-w-32">{{ serviceSelectionne.URLRessource }}</span>
              </div>
            </div>
          </div>

          <!-- Notes internes -->
          <div v-if="serviceSelectionne.NotesInternes" class="bg-sky-500/5 border border-sky-500/20 rounded-xl p-4">
            <p class="text-xs text-sky-400 mb-2 flex items-center gap-1.5"><StickyNote class="w-3.5 h-3.5" /> Notes internes</p>
            <p class="text-sm text-slate-300 whitespace-pre-wrap">{{ serviceSelectionne.NotesInternes }}</p>
          </div>
        </div>
      </main>
    </div>

    <!-- ══════════════════════════════════════════════
         FOOTER — services en attente / résumé rapide
    ══════════════════════════════════════════════ -->
    <footer class="border-t border-slate-700/60 bg-slate-900/80 backdrop-blur-md px-6 py-3">
      <div class="flex items-center gap-4 overflow-x-auto scrollbar-none">
        <span class="text-xs text-slate-600 flex-shrink-0">Résumé :</span>

        <div
          v-for="stat in statsCategories"
          :key="stat.Categorie"
          class="flex-shrink-0 flex items-center gap-2 bg-slate-800/50 border border-slate-700/40 rounded-lg px-3 py-1.5"
        >
          <component :is="getCatIcon(stat.Categorie)" :class="['w-3 h-3', getCatColor(stat.Categorie)]" />
          <span class="text-xs text-slate-400">{{ getCatLabel(stat.Categorie) }}</span>
          <span class="text-xs font-semibold text-slate-200">{{ stat.count }}</span>
          <span v-if="stat.tarifMoyen > 0" class="text-xs text-teal-400">{{ formatCAD(stat.tarifMoyen) }}/moy.</span>
        </div>

        <!-- Alerte services sans tarif -->
        <div v-if="servicesSansTarif.length > 0" class="flex-shrink-0 flex items-center gap-1.5 text-xs text-amber-400 bg-amber-500/10 border border-amber-500/20 rounded-lg px-3 py-1.5">
          <AlertTriangle class="w-3 h-3" />
          {{ servicesSansTarif.length }} service{{ servicesSansTarif.length > 1 ? 's' : '' }} sans tarif
        </div>
      </div>
    </footer>

    <!-- ══════════════════════════════════════════════
         MODAL confirmation suppression
    ══════════════════════════════════════════════ -->
    <Teleport to="body">
      <div v-if="showConfirmDelete" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
        <div class="bg-slate-800 border border-slate-700 rounded-2xl p-6 w-80 shadow-2xl">
          <div class="flex items-center gap-3 mb-4">
            <div class="w-10 h-10 rounded-full bg-red-500/15 border border-red-500/30 flex items-center justify-center">
              <Trash2 class="w-5 h-5 text-red-400" />
            </div>
            <div>
              <p class="text-sm font-semibold text-slate-100">Supprimer ce service ?</p>
              <p class="text-xs text-slate-400 mt-0.5">Cette action est irréversible</p>
            </div>
          </div>
          <p class="text-xs text-slate-400 bg-slate-700/50 rounded-lg p-3 mb-4">
            Si ce service est utilisé dans des revenus, il sera automatiquement archivé plutôt que supprimé.
          </p>
          <div class="flex items-center gap-3">
            <button @click="showConfirmDelete = false" class="flex-1 py-2 text-sm text-slate-400 hover:text-slate-200 transition-colors">
              Annuler
            </button>
            <button
              @click="supprimerConfirme"
              class="flex-1 py-2 bg-red-500 hover:bg-red-400 text-white rounded-lg text-sm font-medium transition-all"
            >
              Supprimer
            </button>
          </div>
        </div>
      </div>
    </Teleport>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import {
  LayoutGrid, Plus, Search, Archive, X, Tag, DollarSign, Receipt,
  FileText, Sparkles, ChevronDown, Check, Loader2, Pencil, Trash2,
  Clock, ShieldCheck, AlertTriangle, StickyNote, PackageSearch,
  Users, ClipboardList, BookOpen, Package, Presentation, BookMarked,
  Printer, Video, Globe, HelpCircle
} from 'lucide-vue-next'
import {
  GetAllServices, CreateService, UpdateService,
  ArchiverService, DeleteService
} from '../../wailsjs/go/main/App'

// ── État ─────────────────────────────────────────────────────────────────────
const services = ref([])
const serviceSelectionne = ref(null)
const showForm = ref(false)
const formMode = ref('create') // 'create' | 'edit'
const saving = ref(false)
const erreurSauvegarde = ref('')
const recherche = ref('')
const filtreCategorie = ref('')
const showArchives = ref(false)
const showFutur = ref(false)
const showConfirmDelete = ref(false)
const idASupprimer = ref(null)

// ── Formulaire vide ───────────────────────────────────────────────────────────
const formVide = () => ({
  code: '', categorie: 'consultation', type_livraison: 'presentiel',
  actif: 1, ordre_affichage: 0,
  nom: '', description_courte: '', description_longue: '',
  public_cible: '', notes_internes: '',
  mode_tarification: 'horaire', duree_minutes: 60,
  taux_horaire: 0, tarif_unitaire: 0, tarif_min: 0, tarif_max: 0,
  exempte_taxes: true, tps_applicable: false, tvq_applicable: false,
  nb_places_max: null, nb_seances_forfait: null, duree_programme_semaines: null,
  format_tract: '', couleur_tract: '', duree_video_minutes: null,
  url_ressource: '', tags: '', icone: 'users'
})
const form = ref(formVide())

// ── Options de select ─────────────────────────────────────────────────────────
const categoriesDisponibles = [
  { value: 'consultation', label: 'Consultation', icon: Users },
  { value: 'evaluation',   label: 'Évaluation',   icon: ClipboardList },
  { value: 'rapport',      label: 'Rapport',      icon: BookOpen },
  { value: 'forfait',      label: 'Forfait',      icon: Package },
  { value: 'atelier',      label: 'Atelier',      icon: Presentation },
  { value: 'formation',    label: 'Formation',    icon: BookMarked },
  { value: 'tract',        label: 'Tract',        icon: Printer },
  { value: 'capsule',      label: 'Capsule',      icon: Video },
  { value: 'ressource_num',label: 'Ressource',    icon: Globe },
  { value: 'autre',        label: 'Autre',        icon: HelpCircle },
]

const categoriesOptions = categoriesDisponibles.map(c => ({ label: c.label, value: c.value }))

const modesTarification = [
  { label: 'Horaire', value: 'horaire' },
  { label: 'À l\'acte', value: 'acte' },
  { label: 'Forfait', value: 'forfait' },
  { label: 'Prix libre', value: 'prix_libre' },
  { label: 'Abonnement', value: 'abonnement' },
  { label: 'Gratuit', value: 'gratuit' },
]

const typesLivraison = [
  { label: 'Présentiel', value: 'presentiel' },
  { label: 'Virtuel', value: 'virtuel' },
  { label: 'Hybride', value: 'hybride' },
  { label: 'Asynchrone', value: 'asynchrone' },
  { label: 'Imprimé', value: 'imprime' },
  { label: 'Numérique', value: 'numerique' },
]

const iconesOptions = [
  { label: 'Personnes', value: 'users' },
  { label: 'Presse-papiers', value: 'clipboard-list' },
  { label: 'Livre', value: 'book-open' },
  { label: 'Colis', value: 'package' },
  { label: 'Présentation', value: 'presentation' },
  { label: 'Vidéo', value: 'video' },
  { label: 'Imprimante', value: 'printer' },
  { label: 'Globe', value: 'globe' },
]

// ── Classes FormKit (dark/slate) ──────────────────────────────────────────────
const fkClasses = {
  outer: 'mb-0',
  label: 'block text-xs font-medium text-slate-400 mb-1',
  inner: 'w-full',
  input: 'w-full px-3 py-2 bg-slate-800/70 border border-slate-700/60 rounded-lg text-sm text-slate-200 placeholder-slate-500 focus:outline-none focus:border-teal-500/60 focus:ring-1 focus:ring-teal-500/20 transition-all',
  message: 'text-xs text-red-400 mt-1',
}

const fkCheckboxClasses = {
  outer: 'mb-0',
  label: 'text-xs text-slate-400 ml-2 cursor-pointer',
  input: 'accent-teal-500 cursor-pointer',
  wrapper: 'flex items-center',
}

// ── Computed ──────────────────────────────────────────────────────────────────
const servicesFiltres = computed(() => {
  return services.value.filter(s => {
    if (!showArchives.value && !s.Actif) return false
    if (filtreCategorie.value && s.Categorie !== filtreCategorie.value) return false
    if (recherche.value) {
      const q = recherche.value.toLowerCase()
      return s.Nom?.toLowerCase().includes(q) || s.Code?.toLowerCase().includes(q)
    }
    return true
  })
})

const servicesActifs = computed(() => services.value.filter(s => s.Actif))

const servicesGroupes = computed(() => {
  const groupes = {}
  servicesFiltres.value.forEach(s => {
    if (!groupes[s.Categorie]) groupes[s.Categorie] = []
    groupes[s.Categorie].push(s)
  })
  return Object.entries(groupes).map(([categorie, svcs]) => ({ categorie, services: svcs }))
})

const statsCategories = computed(() => {
  const stats = {}
  servicesActifs.value.forEach(s => {
    if (!stats[s.Categorie]) stats[s.Categorie] = { categorie: s.Categorie, count: 0, totalTarif: 0, countTarif: 0 }
    stats[s.Categorie].count++
    const tarif = s.taux_horaire || s.tarif_unitaire
    if (tarif > 0) {
      stats[s.Categorie].totalTarif += tarif
      stats[s.Categorie].countTarif++
    }
  })
  return Object.values(stats).map(s => ({
    ...s,
    tarifMoyen: s.countTarif > 0 ? s.totalTarif / s.countTarif : 0
  }))
})

const servicesSansTarif = computed(() =>
  servicesActifs.value.filter(s =>
    s.ModeTarification !== 'gratuit' &&
    !s.TauxHoraire && !s.TarifUnitaire && !s.TarifMin
  )
)

// ── Helpers visuels ───────────────────────────────────────────────────────────
const catMap = {
  consultation: { label: 'Consultation', icon: Users,       bg: 'bg-teal-500/15',   color: 'text-teal-400' },
  evaluation:   { label: 'Évaluation',   icon: ClipboardList,bg: 'bg-sky-500/15',    color: 'text-sky-400' },
  rapport:      { label: 'Rapport',      icon: BookOpen,     bg: 'bg-indigo-500/15', color: 'text-indigo-400' },
  forfait:      { label: 'Forfait',      icon: Package,      bg: 'bg-violet-500/15', color: 'text-violet-400' },
  atelier:      { label: 'Atelier',      icon: Presentation, bg: 'bg-amber-500/15',  color: 'text-amber-400' },
  groupe:       { label: 'Groupe',       icon: Users,        bg: 'bg-orange-500/15', color: 'text-orange-400' },
  formation:    { label: 'Formation',    icon: BookMarked,   bg: 'bg-emerald-500/15',color: 'text-emerald-400' },
  tract:        { label: 'Tract',        icon: Printer,      bg: 'bg-pink-500/15',   color: 'text-pink-400' },
  capsule:      { label: 'Capsule',      icon: Video,        bg: 'bg-red-500/15',    color: 'text-red-400' },
  ressource_num:{ label: 'Ressource',    icon: Globe,        bg: 'bg-cyan-500/15',   color: 'text-cyan-400' },
  autre:        { label: 'Autre',        icon: HelpCircle,   bg: 'bg-slate-500/15',  color: 'text-slate-400' },
}

const getCatIcon  = (cat) => catMap[cat]?.icon   || HelpCircle
const getCatLabel = (cat) => catMap[cat]?.label  || cat
const getCatBg    = (cat) => catMap[cat]?.bg     || 'bg-slate-500/15'
const getCatColor = (cat) => catMap[cat]?.color  || 'text-slate-400'

const getModeLabel = (mode) => ({
  horaire: 'À l\'heure', acte: 'À l\'acte', forfait: 'Forfait fixe',
  prix_libre: 'Prix libre', abonnement: 'Abonnement', gratuit: 'Gratuit'
}[mode] || mode)

const formatCAD = (n) => n ? n.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ' ') + ' $' : '—'

const formatTarif = (svc) => {
  if (svc.ModeTarification === 'horaire' && svc.TauxHoraire)
    return `${svc.TauxHoraire} $/h`
  if (svc.TarifUnitaire) return formatCAD(svc.TarifUnitaire)
  if (svc.ModeTarification === 'gratuit') return 'Gratuit'
  return 'Sans tarif'
}

const formatTarifDetail = (svc) => {
  if (svc.ModeTarification === 'horaire' && svc.TauxHoraire)
    return `${svc.TauxHoraire} $/h`
  if (svc.TarifUnitaire) return formatCAD(svc.TarifUnitaire)
  if (svc.TarifMin) return `${formatCAD(svc.TarifMin)} – ${formatCAD(svc.TarifMax)}`
  if (svc.ModeTarification === 'gratuit') return 'Gratuit'
  return '—'
}

const hasOptionsAvancees = (svc) =>
  svc.NbPlacesMax || svc.NbSeancesForfait || svc.FormatTract ||
  svc.URLRessource || svc.DureeVideoMinutes

// ── Actions ───────────────────────────────────────────────────────────────────
const charger = async () => {
  try {
    const res = await GetAllServices('', false)
    services.value = res || []
    console.log('✅ Services chargés:', services.value.length)
  } catch (e) {
    console.error('❌ Erreur chargement services:', e)
  }
}

const selectionner = (svc) => {
  serviceSelectionne.value = svc
  showForm.value = false
}

const ouvrirFormulaire = (svc) => {
  if (svc) {
    formMode.value = 'edit'
    form.value = {
      ...formVide(),
      ID: svc.ID,
      code: svc.Code, categorie: svc.Categorie, type_livraison: svc.TypeLivraison,
      actif: svc.Actif, ordre_affichage: svc.OrdreAffichage,
      nom: svc.Nom, description_courte: svc.DescriptionCourte,
      description_longue: svc.DescriptionLongue, public_cible: svc.PublicCible,
      notes_internes: svc.NotesInternes, mode_tarification: svc.ModeTarification,
      duree_minutes: svc.DureeMinutes, taux_horaire: svc.TauxHoraire,
      tarif_unitaire: svc.TarifUnitaire, tarif_min: svc.TarifMin, tarif_max: svc.TarifMax,
      exempte_taxes: !!svc.ExempteTaxes, tps_applicable: !!svc.TPSApplicable,
      tvq_applicable: !!svc.TVQApplicable, nb_places_max: svc.NbPlacesMax,
      nb_seances_forfait: svc.NbSeancesForfait, format_tract: svc.FormatTract,
      couleur_tract: svc.CouleurTract, duree_video_minutes: svc.DureeVideoMinutes,
      url_ressource: svc.URLRessource, tags: svc.Tags, icone: svc.Icone,
    }
  } else {
    formMode.value = 'create'
    form.value = formVide()
  }
  showForm.value = true
  serviceSelectionne.value = null
}

const fermerFormulaire = () => {
  showForm.value = false
  erreurSauvegarde.value = ''
}

const sauvegarder = async () => {
  saving.value = true
  erreurSauvegarde.value = ''
  try {
    const payload = {
      ...form.value,
      exempte_taxes: form.value.exempte_taxes ? 1 : 0,
      tps_applicable: form.value.tps_applicable ? 1 : 0,
      tvq_applicable: form.value.tvq_applicable ? 1 : 0,
    }
    if (formMode.value === 'create') {
      await CreateService(payload)
    } else {
      await UpdateService({ id: form.value.ID, ...payload })
    }
    await charger()
    fermerFormulaire()
  } catch (e) {
    erreurSauvegarde.value = e.toString()
  } finally {
    saving.value = false
  }
}

const archiver = async (id) => {
  try {
    await ArchiverService(id)
    await charger()
    serviceSelectionne.value = null
  } catch (e) {
    console.error(e)
  }
}

const confirmerSuppression = (id) => {
  idASupprimer.value = id
  showConfirmDelete.value = true
}

const supprimerConfirme = async () => {
  try {
    await DeleteService(idASupprimer.value)
    await charger()
    serviceSelectionne.value = null
  } catch (e) {
    // Si utilisé → archivage automatique côté Go, on recharge
    await charger()
  } finally {
    showConfirmDelete.value = false
    idASupprimer.value = null
  }
}

const toggleArchives = () => {
  showArchives.value = !showArchives.value
}

const onModeTarifChange = () => {
  form.value.taux_horaire = 0
  form.value.tarif_unitaire = 0
  form.value.tarif_min = 0
  form.value.tarif_max = 0
}

// ── Init ──────────────────────────────────────────────────────────────────────
onMounted(charger)
</script>

<style scoped>
.scrollbar-none::-webkit-scrollbar { display: none; }
.scrollbar-none { -ms-overflow-style: none; scrollbar-width: none; }
</style>