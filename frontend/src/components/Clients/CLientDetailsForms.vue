<template>
  <div class="space-y-4">

    <!-- ══ BARRE D'ACTIONS ══ -->
    <div class="bg-white dark:bg-gray-900 rounded-xl border border-gray-200 dark:border-gray-800 px-4 py-2.5 flex flex-wrap items-center gap-2">

      <button @click="showContacts = !showContacts"
        :class="showContacts
          ? 'bg-blue-600 text-white'
          : 'bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700'"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium transition-all">
        <UsersIcon :size="14" />
        Contacts
        <span v-if="contactsCount > 0"
          class="px-1.5 py-0.5 rounded-full text-[10px] font-bold"
          :class="showContacts ? 'bg-white/25 text-white' : 'bg-blue-100 dark:bg-blue-900/40 text-blue-700 dark:text-blue-300'">
          {{ contactsCount }}
        </span>
      </button>

      <button @click="$emit('show-notes')"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all">
        <FileText :size="14" />
        Notes cliniques
      </button>

      <button v-if="folderExists" @click="handleOpenFolder"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all">
        <FolderOpen :size="14" />
        Dossier fichiers
      </button>

      <button @click="$emit('show-evaluations')"
        class="flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all">
        <ShieldCheck :size="14" />
        Évaluations
      </button>

      <div class="ml-auto">
        <span :class="getStatusBadgeClass()" class="text-xs">{{ getStatusText() }}</span>
      </div>
    </div>

    <!-- Widget dossier -->
    <ClientFolderWidget
      v-if="formData.id"
      :leopard-number="leopardNumber"
      :client-name="`${formData.nom} ${formData.prenom}`"
      :folder-exists="folderExists"
      :folder-info="folderInfo"
      @folder-created="onFolderCreated"
      @folder-opened="onFolderOpened"
      @refresh="checkFolderExists"
    />

    <!-- Fiche rapide -->
    <ClientQuickCard
      v-if="formData.id"
      :client="formData"
      @view-medecin="viewMedecinDetails"
    />

    <!-- ══ FORMULAIRE AVEC NAV LATÉRALE ══ -->
    <div class="bg-white dark:bg-gray-900 rounded-xl border border-gray-200 dark:border-gray-800 overflow-hidden">
      <div class="flex" style="min-height: 520px;">

        <!-- Navigation latérale -->
        <div class="w-44 shrink-0 border-r border-gray-100 dark:border-gray-800 bg-gray-50 dark:bg-gray-900/60 py-3 flex flex-col">
          <nav class="flex flex-col gap-0.5 px-2 flex-1">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="activeTab = tab.id"
              :class="activeTab === tab.id
                ? 'bg-teal-700 text-white shadow-sm'
                : 'text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800 hover:text-gray-700 dark:hover:text-gray-200'"
              class="flex items-center gap-2 px-3 py-2 rounded-lg text-xs font-medium transition-all text-left w-full whitespace-nowrap">
              <component :is="tab.icon" :size="14" class="shrink-0" />
              {{ tab.label }}
            </button>
          </nav>
          <!-- Sauvegarder -->
          <div class="px-2 pt-3 border-t border-gray-100 dark:border-gray-800 mt-2">
            <button @click="submitHandler" type="button"
              class="flex items-center justify-center gap-1.5 w-full px-3 py-2 rounded-lg text-xs font-semibold bg-teal-700 hover:bg-teal-800 text-white shadow-sm transition-all">
              <Save :size="13" />
              Sauvegarder
            </button>
          </div>
        </div>

        <!-- Contenu onglet -->
        <div class="flex-1 overflow-auto p-5">
          <form @submit.prevent="submitHandler">

            <!-- ══ IDENTITÉ ══ -->
            <div v-show="activeTab === 'identity'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Identité du client
              </p>
              <div class="grid grid-cols-3 gap-3">

                <div class="col-span-2">
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">
                    Nom <span class="text-red-500">*</span>
                  </label>
                  <input v-model="formData.nom" required
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div class="col-span-1">
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">
                    Prénom <span class="text-red-500">*</span>
                  </label>
                  <input v-model="formData.prenom" required
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Date de naissance</label>
                  <input v-model="formData.date_naissance" type="date"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Sexe</label>
                  <select v-model="formData.sexe"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all">
                    <option value="">—</option>
                    <option value="M">Masculin</option>
                    <option value="F">Féminin</option>
                    <option value="X">Non-binaire</option>
                  </select>
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Statut marital</label>
                  <select v-model="formData.statut_marital"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all">
                    <option value="">—</option>
                    <option value="Célibataire">Célibataire</option>
                    <option value="Marié(e)">Marié(e)</option>
                    <option value="Conjoint(e) de fait">Conjoint(e) de fait</option>
                    <option value="Divorcé(e)">Divorcé(e)</option>
                    <option value="Veuf(ve)">Veuf(ve)</option>
                    <option value="Séparé(e)">Séparé(e)</option>
                  </select>
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Lieu de naissance</label>
                  <input v-model="formData.lieu_naissance" placeholder="Ville, Pays"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Identité de genre</label>
                  <input v-model="formData.identite_genre"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Orientation sexuelle</label>
                  <input v-model="formData.orientation_sexuelle"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div class="col-span-3">
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">N° Dossier Léopard</label>
                  <div class="flex items-center gap-2 px-3 py-1.5 bg-teal-50 dark:bg-teal-900/20 border border-teal-200 dark:border-teal-800 rounded-md">
                    <Shield :size="13" class="text-teal-600 dark:text-teal-400 shrink-0" />
                    <span class="font-mono text-xs font-semibold text-teal-700 dark:text-teal-300">
                      {{ leopardNumber || 'En attente de génération' }}
                    </span>
                  </div>
                </div>

              </div>
            </div>

            <!-- ══ COORDONNÉES ══ -->
            <div v-show="activeTab === 'contact'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Coordonnées
              </p>
              <div class="grid grid-cols-3 gap-3">
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Téléphone</label>
                  <input v-model="formData.telephone" type="tel" placeholder="(819) 555-0123"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Cellulaire</label>
                  <input v-model="formData.cellulaire" type="tel" placeholder="(819) 555-0456" 
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Courriel</label>
                  <input v-model="formData.email" type="email" placeholder="exemple@email.com"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>
              </div>
              <!-- Appel rapide -->
              <div v-if="formData.telephone || formData.cellulaire || formData.email"
                class="mt-4 flex flex-wrap gap-2 p-3 bg-gray-50 dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700">
                <span class="text-[10px] font-bold uppercase tracking-wider text-gray-400 self-center mr-1">Appel rapide</span>
                <a v-if="formData.telephone" :href="`tel:${formData.telephone}`"
                  class="flex items-center gap-1.5 px-2.5 py-1 rounded-lg text-xs font-medium bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700 text-gray-700 dark:text-gray-300 hover:border-teal-400 hover:text-teal-700 transition-all">
                  <Phone :size="11" /> Domicile
                </a>
                <a v-if="formData.cellulaire" :href="`tel:${formData.cellulaire}`"
                  class="flex items-center gap-1.5 px-2.5 py-1 rounded-lg text-xs font-medium bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700 text-gray-700 dark:text-gray-300 hover:border-blue-400 hover:text-blue-700 transition-all">
                  <Smartphone :size="11" /> Cellulaire
                </a>
                <a v-if="formData.email" :href="`mailto:${formData.email}`"
                  class="flex items-center gap-1.5 px-2.5 py-1 rounded-lg text-xs font-medium bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-700 text-gray-700 dark:text-gray-300 hover:border-violet-400 hover:text-violet-700 transition-all">
                  <Mail :size="11" /> Courriel
                </a>
              </div>
            </div>

            <!-- ══ ADRESSE ══ -->
            <div v-show="activeTab === 'address'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Adresse
              </p>
              <div class="grid grid-cols-3 gap-3">

                <div class="col-span-2">
                  <label class="flex items-center justify-between text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">
                    <span>Adresse</span>
                    <button v-if="formData.adresse" @click.prevent="openGoogleMaps" type="button"
                      class="flex items-center gap-1 text-[10px] text-blue-600 dark:text-blue-400 hover:underline font-medium normal-case tracking-normal">
                      <MapPin :size="10" /> Google Maps
                    </button>
                  </label>
                  <input v-model="formData.adresse" placeholder="123 Rue Principale"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Appartement</label>
                  <input v-model="formData.appartement" placeholder="Apt 101"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Ville</label>
                  <input v-model="formData.ville" placeholder="Gatineau"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Code postal</label>
                  <input v-model="formData.code_postal" placeholder="J8V 1A1"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono uppercase focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Province</label>
                  <select v-model="formData.province"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all">
                    <option value="">—</option>
                    <option value="QC">Québec</option>
                    <option value="ON">Ontario</option>
                    <option value="BC">Colombie-Britannique</option>
                    <option value="AB">Alberta</option>
                    <option value="MB">Manitoba</option>
                    <option value="SK">Saskatchewan</option>
                    <option value="NS">Nouvelle-Écosse</option>
                    <option value="NB">Nouveau-Brunswick</option>
                    <option value="PE">Île-du-Prince-Édouard</option>
                    <option value="NL">Terre-Neuve-et-Labrador</option>
                  </select>
                </div>

                <div class="col-span-3">
                  <PaysSelect v-model="formData.pays" />
                </div>
              </div>
            </div>

            <!-- ══ SOCIO-CULTUREL ══ -->
            <div v-show="activeTab === 'sociocultural'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Profil socio-culturel
              </p>
              <div class="grid grid-cols-3 gap-3">

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Occupation</label>
                  <input v-model="formData.occupation"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Profession</label>
                  <input v-model="formData.profession"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Employeur</label>
                  <input v-model="formData.employeur"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Niveau scolaire</label>
                  <select v-model="formData.niveau_scolaire"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all">
                    <option value="">—</option>
                    <option value="Primaire">Primaire</option>
                    <option value="Secondaire">Secondaire</option>
                    <option value="Collégial">Collégial</option>
                    <option value="Universitaire">Universitaire</option>
                    <option value="Autre">Autre</option>
                  </select>
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Langue préférée</label>
                  <select v-model="formData.langue_preferee"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all">
                    <option value="Français">Français</option>
                    <option value="Anglais">Anglais</option>
                    <option value="Autre">Autre</option>
                  </select>
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Première langue</label>
                  <input v-model="formData.premiere_langue"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>

                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Origine ethnique</label>
                  <input v-model="formData.origine_ethnique"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Religion</label>
                  <input v-model="formData.religion"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                </div>

                <div class="col-span-3">
                  <label class="flex items-center gap-2 cursor-pointer group">
                    <input v-model="formData.premiere_nation" type="checkbox" :true-value="1" :false-value="0"
                      class="w-3.5 h-3.5 text-teal-600 border-gray-300 rounded focus:ring-teal-500" />
                    <span class="text-xs font-medium text-gray-600 dark:text-gray-400 group-hover:text-gray-800 dark:group-hover:text-gray-200 transition-colors">
                      Personne des Premières Nations
                    </span>
                  </label>
                </div>

              </div>
            </div>

            <!-- ══ MÉDICAL ══ -->
            <div v-show="activeTab === 'medical'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Dossier médical
              </p>

              <!-- Numéros -->
              <div class="grid grid-cols-3 gap-3 mb-5">
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">RAMQ</label>
                  <input v-model="formData.numero_assurance_maladie" placeholder="TREJ 4503 1501"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono tracking-wider focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">NAS</label>
                  <input v-model="formData.numero_securite_sociale" placeholder="123-456-789"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" />
                </div>
                <div class="grid grid-cols-2 gap-2">
                  <div>
                    <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">HCM</label>
                    <input v-model="formData.no_hcm"
                      class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                  </div>
                  <div>
                    <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">CHAUR</label>
                    <input v-model="formData.no_chaur"
                      class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 font-mono focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
                  </div>
                </div>
              </div>

              <!-- Médecin -->
              <div class="flex items-center gap-3 mb-3">
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
                <span class="text-[10px] font-bold uppercase tracking-wider text-gray-400 dark:text-gray-600">Médecin de famille</span>
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
              </div>
              <MedecinSelector
                v-model="formData.medecin_famille_No_Licence"
                @medecin-selected="handleMedecinSelected"
                @view-details="openMedecinModal"
              />
              <MedecinDetailsModal
                v-if="showMedecinModal && selectedMedecinForModal"
                :medecin="selectedMedecinForModal"
                @close="showMedecinModal = false"
                @edit="editMedecin"
                @delete="deleteMedecin"
              />
            </div>

            <!-- ══ ÉTABLISSEMENTS ══ -->
            <div v-show="activeTab === 'establishments'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Établissements & intervenants
              </p>

              <!-- Milieu de vie -->
              <div class="flex items-center gap-3 mb-3">
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
                <span class="text-[10px] font-bold uppercase tracking-wider text-gray-400 dark:text-gray-600">Milieu de vie</span>
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
              </div>
              <div class="grid grid-cols-3 gap-3 mb-5">
                <div class="col-span-2">
                  <RPAselector v-model="formData.rpa_id" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">CHSLD (ID)</label>
                  <input v-model="formData.chsld_id" type="number"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" placeholder="—" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">RI (ID)</label>
                  <input v-model="formData.ri_id" type="number"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" placeholder="—" />
                </div>
                <div>
                  <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">PIVOT (ID)</label>
                  <input v-model="formData.pivot_id" type="number"
                    class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all placeholder-gray-400" placeholder="—" />
                </div>
              </div>

              <!-- Notaire -->
              <div class="flex items-center gap-3 mb-3">
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
                <span class="text-[10px] font-bold uppercase tracking-wider text-gray-400 dark:text-gray-600">Notaire & procurations</span>
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
              </div>
              <div class="mb-4">
                <NotaireSelector
                  v-model="formData.notaire_id"
                  @notaire-selected="handleNotaireSelected"
                  @view-details="viewNotaireDetails"
                />
              </div>
              <div class="flex flex-wrap gap-6">
                <label class="flex items-center gap-2 cursor-pointer group">
                  <input v-model="formData.procuration_bancaire" type="checkbox" :true-value="1" :false-value="0"
                    class="w-3.5 h-3.5 text-teal-600 border-gray-300 rounded focus:ring-teal-500" />
                  <span class="text-xs font-medium text-gray-600 dark:text-gray-400 group-hover:text-gray-800 dark:group-hover:text-gray-200 transition-colors">
                    Procuration bancaire
                  </span>
                </label>
                <label class="flex items-center gap-2 cursor-pointer group">
                  <input v-model="formData.procuration_notariee" type="checkbox" :true-value="1" :false-value="0"
                    class="w-3.5 h-3.5 text-teal-600 border-gray-300 rounded focus:ring-teal-500" />
                  <span class="text-xs font-medium text-gray-600 dark:text-gray-400 group-hover:text-gray-800 dark:group-hover:text-gray-200 transition-colors">
                    Procuration notariée
                  </span>
                </label>
              </div>
            </div>

            <!-- ══ NOTES & STATUT ══ -->
            <div v-show="activeTab === 'notes'">
              <p class="text-[10px] font-bold uppercase tracking-widest text-gray-400 dark:text-gray-500 mb-4 pb-2 border-b border-gray-100 dark:border-gray-800">
                Notes & statut du dossier
              </p>

              <div class="mb-4">
                <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">
                  Note fixe (visible sur la fiche)
                </label>
                <textarea v-model="formData.note_fixe" rows="4"
                  placeholder="Notes importantes : allergies, directives particulières, comportements…"
                  class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all resize-none placeholder-gray-400">
                </textarea>
                <div v-if="formData.note_fixe"
                  class="mt-1.5 flex items-start gap-2 px-3 py-2 rounded-lg bg-amber-50 dark:bg-amber-900/20 border-l-2 border-amber-400">
                  <Eye :size="12" class="text-amber-500 shrink-0 mt-0.5" />
                  <p class="text-xs text-amber-700 dark:text-amber-300 leading-relaxed">{{ formData.note_fixe }}</p>
                </div>
              </div>

              <div class="flex items-center gap-3 mb-3">
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
                <span class="text-[10px] font-bold uppercase tracking-wider text-gray-400 dark:text-gray-600">Statut du dossier</span>
                <div class="h-px flex-1 bg-gray-100 dark:bg-gray-800"></div>
              </div>

              <div class="grid grid-cols-2 gap-3 mb-4">
                <label class="flex items-center gap-2.5 px-3 py-2.5 border rounded-lg cursor-pointer transition-all"
                  :class="formData.actif === 1
                    ? 'border-emerald-400 bg-emerald-50 dark:bg-emerald-900/10 dark:border-emerald-700'
                    : 'border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800'">
                  <input v-model="formData.actif" type="checkbox" :true-value="1" :false-value="0"
                    class="w-3.5 h-3.5 text-emerald-600 border-gray-300 rounded" />
                  <div>
                    <p class="text-xs font-semibold text-gray-700 dark:text-gray-300">Dossier actif</p>
                    <p class="text-[10px] text-gray-400">Suivi en cours</p>
                  </div>
                </label>
                <label class="flex items-center gap-2.5 px-3 py-2.5 border rounded-lg cursor-pointer transition-all"
                  :class="formData.dcd === 1
                    ? 'border-gray-500 bg-gray-100 dark:bg-gray-800 dark:border-gray-600'
                    : 'border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800'">
                  <input v-model="formData.dcd" type="checkbox" :true-value="1" :false-value="0"
                    class="w-3.5 h-3.5 text-gray-600 border-gray-300 rounded" />
                  <div>
                    <p class="text-xs font-semibold text-gray-700 dark:text-gray-300">Décédé (DCD)</p>
                    <p class="text-[10px] text-gray-400">Ferme le dossier</p>
                  </div>
                </label>
              </div>

              <div v-if="formData.dcd === 1">
                <label class="block text-[11px] font-semibold text-gray-500 dark:text-gray-400 mb-1 uppercase tracking-wide">Date de décès</label>
                <input v-model="formData.date_deces" type="date"
                  class="w-full px-2.5 py-1.5 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 text-sm text-gray-900 dark:text-gray-100 focus:border-teal-500 focus:outline-none focus:ring-1 focus:ring-teal-500/30 transition-all" />
              </div>
            </div>

          </form>
        </div>
      </div>
    </div>

    <!-- Évaluations -->
    <EvaluationView
      :is-open="showEvaluations"
      :client="formData"
      @close="showEvaluations = false"
    />

    <!-- Sidebar Contacts -->
    <Transition name="slide">
      <div v-if="showContacts"
        class="fixed inset-y-0 right-0 w-full md:w-[500px] bg-white dark:bg-gray-900 shadow-2xl z-50 overflow-y-auto border-l border-gray-200 dark:border-gray-800">
        <ContactsSidebar
          :client-id="formData.id"
          @close="showContacts = false"
          @contact-added="contactsCount++"
          @contact-deleted="contactsCount--"
        />
      </div>
    </Transition>
    <Transition name="fade">
      <div v-if="showContacts" @click="showContacts = false"
        class="fixed inset-0 bg-black/40 backdrop-blur-sm z-40">
      </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { 
  User, Phone, Mail, MapPin, Heart, Building, ClipboardList, 
  FileText, UsersIcon, FolderOpen, FolderPlus, Shield, Eye,
  Save, Loader2, Smartphone, Globe, ShieldCheck
} from 'lucide-vue-next'
import ContactsSidebar from '../../components/Contacts/ContactSidebar.vue'
import ClientQuickCard from '../Clients/ClientQuickCard.vue'
import MedecinSelector from '@/components/Medecins/MedecinSelector.vue'
import MedecinDetailsModal from '@/components/Medecins/MedecinDetailsModal.vue'
import {
  generateLeopardNumber, createClientFolder, openClientFolder, clientFolderExists,
} from "@/utils/clientFolderManager";
import ClientFolderWidget from './ClientFolderWidget.vue';
import RPAselector from '../RPA/RPAselector.vue'
import NotaireSelector from '@/components/Notaires/NotaireSelector.vue'
import EvaluationView from '../evaluation/EvaluationView.vue'
import PaysSelect from '../Selectors/PaysSelector.vue'
import { GetAllPays } from '../../../wailsjs/go/main/App'

const showEvaluations = ref(false)

const props = defineProps({
  clientData: { type: Object, required: true }
})

const emit = defineEmits(['save', 'folderCreated', 'show-notes', 'show-evaluations'])

const formData = ref({ ...props.clientData })
const leopardNumber = ref('')
const folderExists = ref(false)
const folderInfo = ref(null)
const isCreatingFolder = ref(false)
const showContacts = ref(false)
const contactsCount = ref(0)
const activeTab = ref('identity')

const clientStatus = ref(null)

const loadClientStatus = async () => {
  try {
    clientStatus.value = await GetClientStatus(clientId)
  } catch (error) {
    console.error('Erreur chargement statut:', error)
  }
}

const tabs = [
  { id: 'identity',       label: 'Identité',      icon: User          },
  { id: 'contact',        label: 'Coordonnées',    icon: Phone         },
  { id: 'address',        label: 'Adresse',        icon: MapPin        },
  { id: 'sociocultural',  label: 'Socio-culturel', icon: Globe         },
  { id: 'medical',        label: 'Médical',        icon: Heart         },
  { id: 'establishments', label: 'Établissements', icon: Building      },
  { id: 'notes',          label: 'Notes & Statut', icon: ClipboardList },
]

const getStatusBadgeClass = () => {
  if (!clientStatus.value) return ''
  const baseClass = 'px-3 py-1 rounded-full text-sm font-medium'
  switch (clientStatus.value.color) {
    case 'gray':   return `${baseClass} bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300`
    case 'green':  return `${baseClass} bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300`
    case 'orange': return `${baseClass} bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300`
    default:       return baseClass
  }
}

const getStatusText = () => clientStatus.value?.label || 'Inconnu'

const openGoogleMaps = () => {
  const address = `${formData.value.adresse}, ${formData.value.ville}, ${formData.value.code_postal}`
  window.open(`https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(address)}`, '_blank')
}

const viewMedecinDetails = () => {
  alert('Ouverture de la fiche médecin (à implémenter)')
}

const showMedecinModal = ref(false)
const selectedMedecinForModal = ref(null)

const openMedecinModal = (medecin) => {
  selectedMedecinForModal.value = medecin
  showMedecinModal.value = true
}

const handleMedecinSelected = (medecin) => {
  console.log('Médecin sélectionné:', medecin)
}

const editMedecin = (medecin) => {
  console.log('Éditer médecin:', medecin)
  showMedecinModal.value = false
}

const deleteMedecin = async (medecinId) => {
  try {
    await window.go.main.App.DeleteMedecin(medecinId)
    alert('✅ Médecin supprimé')
    showMedecinModal.value = false
    if (formData.value.medecin_famille_No_Licence === selectedMedecinForModal.value.licence) {
      formData.value.medecin_famille_No_Licence = ''
    }
  } catch (error) {
    console.error('Erreur suppression:', error)
    alert('❌ Erreur lors de la suppression')
  }
}

const handleNotaireSelected = (notaire) => {
  console.log('Notaire sélectionné:', notaire)
}

const viewNotaireDetails = (notaire) => {
  console.log('Voir notaire:', notaire)
}

const onFolderCreated = (data) => {
  folderExists.value = true
  leopardNumber.value = data.leopardNumber
  formData.value.no_dossier_leopard = data.leopardNumber
  loadFolderInfo()
  alert(`✅ Dossier créé: ${data.leopardNumber}`)
  emit('folderCreated', data)
}

const onFolderOpened = () => {
  alert("📂 Dossier ouvert dans l'explorateur")
}

const handleOpenFolder = () => {
  console.log('Ouvrir dossier:', leopardNumber.value)
}

const handleCreateFolder = () => {
  console.log('Créer dossier:', leopardNumber.value)
}

const checkFolderExists = async () => {
  try {
    const exists = await window.go.main.App.ClientFolderExists(leopardNumber.value)
    folderExists.value = exists
  } catch {}
}

const loadFolderInfo = async () => {
  try {
    const info = await window.go.main.App.GetClientFolderInfo(leopardNumber.value)
    if (info) folderInfo.value = info
  } catch (error) {
    console.error('❌ Erreur chargement infos dossier:', error)
  }
}
const formatPhone = (phone) => {
  if (!phone) return "";
  const cleaned = phone.replace(/\D/g, "");

  if (cleaned.length === 10) {
    return `(${cleaned.slice(0, 3)}) ${cleaned.slice(3, 6)}-${cleaned.slice(6)}`;
  }

  if (cleaned.length === 11 && cleaned[0] === "1") {
    return `+1 (${cleaned.slice(1, 4)}) ${cleaned.slice(4, 7)}-${cleaned.slice(7)}`;
  }

  return phone;
};

const submitHandler = () => {
  emit('save', formData.value)
}

watch(() => props.clientData, (newVal) => {
  formData.value = { ...newVal }
  leopardNumber.value = newVal.no_dossier_leopard || ''
}, { deep: true })

const listePays = ref([])

onMounted(async () => {
  try { listePays.value = await GetAllPays() } catch {}
  leopardNumber.value = formData.value.no_dossier_leopard || ''
})

const paysAlpha2 = computed(() => {
  if (!formData.value.pays || !listePays.value.length) return ''
  const trouve = listePays.value.find(p =>
    p.pays.toLowerCase().trim() === formData.value.pays.toLowerCase().trim()
  )
  return trouve ? trouve.alpha2 : ''
})
</script>

<style scoped>
.slide-enter-active, .slide-leave-active { transition: transform 0.28s ease; }
.slide-enter-from, .slide-leave-to { transform: translateX(100%); }
.fade-enter-active, .fade-leave-active { transition: opacity 0.25s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>