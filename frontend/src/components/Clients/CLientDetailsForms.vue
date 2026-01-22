<template>
  <div class="space-y-6">
 <!-- Menu d'actions contextuel -->
    <div class="bg-white dark:bg-gray-900 rounded-xl shadow-sm border border-gray-200 dark:border-gray-800 p-4">

      
      <div class="flex flex-wrap items-center gap-3">
        <!-- Bouton Contacts -->
        <button
          @click="showContacts = !showContacts"
          :class="[
            'flex items-center gap-2 px-4 py-2.5 rounded-lg font-medium transition-all',
            showContacts 
              ? 'bg-blue-600 text-white shadow-lg' 
              : 'bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700'
          ]"
        >
          <UsersIcon :size="18" />
          <span>Contacts</span>
          <span v-if="contactsCount > 0" class="px-2 py-0.5 bg-white/20 rounded-full text-xs">
            {{ contactsCount }}
          </span>
        </button>

        <!-- Bouton Notes -->
        <button
          @click="$emit('show-notes')"
          class="flex items-center gap-2 px-4 py-2.5 rounded-lg font-medium bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all"
        >
          <FileText :size="18" />
          <span>Notes cliniques</span>
        </button>

        <!-- Bouton Dossier -->
        <button
          v-if="folderExists"
          @click="handleOpenFolder"
          class="flex items-center gap-2 px-4 py-2.5 rounded-lg font-medium bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all"
        >
          <FolderOpen :size="18" />
          <span>Ouvrir dossier</span>
        </button>


        <!-- Statuts rapides -->
        <template>
  <div class="ml-auto flex items-center gap-2">
    <span :class="getStatusBadgeClass()">
      {{ getStatusText() }}
    </span>
  </div>
</template>
      </div>

      <ClientFolderWidget
  v-if="formData.id"
  :leopard-number="leopardNumber"
  :client-name="`${formData.nom} ${formData.prenom}`"
  :folder-exists="folderExists"
  :folder-info="folderInfo"
  @folder-created="onFolderCreated"
  @folder-opened="onFolderOpened"
  @refresh="checkFolderExists"
  class="mb-6"
/>
    </div>

    <!-- Fiche Rapide Client -->
    <ClientQuickCard 
      v-if="formData.id"
      :client="formData"
      @view-medecin="viewMedecinDetails"
    />

   
    <!-- Sidebar Contacts -->
    <Transition name="slide">
      <div v-if="showContacts" class="fixed inset-y-0 right-0 w-full md:w-[500px] bg-white dark:bg-gray-900 shadow-2xl z-50 overflow-y-auto">
        <ContactsSidebar 
          :client-id="formData.id"
          @close="showContacts = false"
          @contact-added="contactsCount++"
          @contact-deleted="contactsCount--"
        />
      </div>
    </Transition>

    <!-- Overlay -->
    <Transition name="fade">
      <div 
        v-if="showContacts"
        @click="showContacts = false"
        class="fixed inset-0 bg-black/50 backdrop-blur-sm z-40"
      ></div>
    </Transition>

    <!-- Formulaire principal avec onglets -->
    <div class="bg-white dark:bg-gray-900 rounded-xl shadow-sm border border-gray-200 dark:border-gray-800">
      
      <!-- Onglets -->
      <div class="border-b border-gray-200 dark:border-gray-800">
        <nav class="flex gap-1 p-2 overflow-x-auto" aria-label="Tabs">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'flex items-center gap-2 px-4 py-2.5 rounded-lg font-medium transition-all whitespace-nowrap',
              activeTab === tab.id
                ? 'bg-blue-50 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300'
                : 'text-gray-600 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800'
            ]"
          >
            <component :is="tab.icon" :size="18" />
            <span>{{ tab.label }}</span>
          </button>
        </nav>
      </div>

      <!-- Contenu des onglets -->
      <div class="p-6">
        <form @submit.prevent="submitHandler">
          
          <!-- Onglet Identité -->
          <div v-show="activeTab === 'identity'" class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Nom <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="formData.nom"
                  required
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Prénom <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="formData.prenom"
                  required
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Date de naissance
                </label>
                <input 
                  v-model="formData.date_naissance"
                  type="date"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Sexe
                </label>
                <select 
                  v-model="formData.sexe"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                >
                  <option value="">Non spécifié</option>
                  <option value="M">Masculin</option>
                  <option value="F">Féminin</option>
                  <option value="X">Non-binaire</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Lieu de naissance
                </label>
                <input 
                  v-model="formData.lieu_naissance"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Statut marital
                </label>
                <select 
                  v-model="formData.statut_marital"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                >
                  <option value="">Non spécifié</option>
                  <option value="Célibataire">Célibataire</option>
                  <option value="Marié(e)">Marié(e)</option>
                  <option value="Conjoint(e) de fait">Conjoint(e) de fait</option>
                  <option value="Divorcé(e)">Divorcé(e)</option>
                  <option value="Veuf(ve)">Veuf(ve)</option>
                  <option value="Séparé(e)">Séparé(e)</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Identité de genre
                </label>
                <input 
                  v-model="formData.identite_genre"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Orientation sexuelle
                </label>
                <input 
                  v-model="formData.orientation_sexuelle"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <!-- N° Leopard -->
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° Dossier Léopard
                </label>
                <div class="w-full px-4 py-2.5 bg-gradient-to-r from-purple-50 to-indigo-50 dark:from-purple-900/20 dark:to-indigo-900/20 border-2 border-purple-200 dark:border-purple-800 rounded-lg flex items-center gap-2">
                  <Shield :size="16" class="text-purple-600 dark:text-purple-400" />
                  <span class="font-mono text-sm text-purple-700 dark:text-purple-300 font-semibold">
                    {{ leopardNumber || 'En attente' }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Onglet Coordonnées -->
          <div v-show="activeTab === 'contact'" class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="flex items-center gap-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  <Phone :size="16" />
                  Téléphone
                </label>
                <input 
                  v-model="formData.telephone"
                  type="tel"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="819-555-0123"
                />
              </div>

              <div>
                <label class="flex items-center gap-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  <Smartphone :size="16" />
                  Cellulaire
                </label>
                <input 
                  v-model="formData.cellulaire"
                  type="tel"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="819-555-0456"
                />
              </div>

              <div class="md:col-span-2">
                <label class="flex items-center gap-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  <Mail :size="16" />
                  Courriel
                </label>
                <input 
                  v-model="formData.email"
                  type="email"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="exemple@email.com"
                />
              </div>
            </div>
          </div>

          <!-- Onglet Adresse -->
          <div v-show="activeTab === 'address'" class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="md:col-span-2">
                <label class="flex items-center justify-between text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  <span>Adresse complète</span>
                  <button
                    v-if="formData.adresse"
                    @click.prevent="openGoogleMaps"
                    type="button"
                    class="flex items-center gap-1 text-xs text-blue-600 dark:text-blue-400 hover:underline"
                  >
                    <MapPin :size="14" />
                    <span>Voir sur Google Maps</span>
                  </button>
                </label>
                <textarea 
                  v-model="formData.adresse"
                  rows="2"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all resize-none"
                  placeholder="123 Rue Principale"
                ></textarea>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Appartement
                </label>
                <input 
                  v-model="formData.appartement"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="Apt 101"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Code postal
                </label>
                <input 
                  v-model="formData.code_postal"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="J8V 1A1"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Ville
                </label>
                <input 
                  v-model="formData.ville"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="Gatineau"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Province
                </label>
                <select 
                  v-model="formData.province"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                >
                  <option value="">Sélectionner...</option>
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

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Pays
                </label>
                <input 
                  v-model="formData.pays"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                  placeholder="Canada"
                />
              </div>
            </div>
          </div>

          <!-- Onglet Socio-culturel -->
          <div v-show="activeTab === 'sociocultural'" class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Occupation
                </label>
                <input 
                  v-model="formData.occupation"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Profession
                </label>
                <input 
                  v-model="formData.profession"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Employeur
                </label>
                <input 
                  v-model="formData.employeur"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Niveau scolaire
                </label>
                <select 
                  v-model="formData.niveau_scolaire"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                >
                  <option value="">Non spécifié</option>
                  <option value="Primaire">Primaire</option>
                  <option value="Secondaire">Secondaire</option>
                  <option value="Collégial">Collégial</option>
                  <option value="Universitaire">Universitaire</option>
                  <option value="Autre">Autre</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Langue préférée
                </label>
                <select 
                  v-model="formData.langue_preferee"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                >
                  <option value="Français">Français</option>
                  <option value="Anglais">Anglais</option>
                  <option value="Autre">Autre</option>
                </select>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Première langue
                </label>
                <input 
                  v-model="formData.premiere_langue"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Origine ethnique
                </label>
                <input 
                  v-model="formData.origine_ethnique"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Religion
                </label>
                <input 
                  v-model="formData.religion"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div class="md:col-span-2">
                <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-lg cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                  <input 
                    v-model="formData.premiere_nation"
                    type="checkbox"
                    :true-value="1"
                    :false-value="0"
                    class="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                  <span class="text-sm font-medium text-gray-700 dark:text-gray-300">Personne des Premières Nations</span>
                </label>
              </div>
            </div>
          </div>

          <!-- Onglet Médical -->
          <div v-show="activeTab === 'medical'" class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° Assurance maladie (RAMQ)
                </label>
                <input 
                  v-model="formData.numero_assurance_maladie"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all font-mono"
                  placeholder="TREJ 4503 1501"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° Sécurité sociale
                </label>
                <input 
                  v-model="formData.numero_securite_sociale"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all font-mono"
                  placeholder="123-456-789"
                />
              </div>

             <!-- Dans l'onglet médical, remplacer le champ médecin par : -->
  <MedecinSelector 
    v-model="formData.medecin_famille_No_Licence"
    @medecin-selected="handleMedecinSelected"
    @view-details="openMedecinModal"
  />

   <!-- Modale médecin -->
  <MedecinDetailsModal
    v-if="showMedecinModal && selectedMedecinForModal"
    :medecin="selectedMedecinForModal"
    @close="showMedecinModal = false"
    @edit="editMedecin"
    @delete="deleteMedecin"
    
  />

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° HCM
                </label>
                <input 
                  v-model="formData.no_hcm"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  N° CHAUR
                </label>
                <input 
                  v-model="formData.no_chaur"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>
            </div>
          </div>

          <!-- Onglet Établissements -->
          <div v-show="activeTab === 'establishments'" class="space-y-4">
            <div class="grid grid-cols-2 md:grid-cols-2 lg:grid-cols-2 gap-4"> 
              <div>

                <RPAselector 
      v-model="formData.rpa_id" 
    />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  CHSLD
                </label>
                <input 
                  v-model="formData.chsld_id"
                  type="number"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  RI
                </label>
                <input 
                  v-model="formData.ri_id"
                  type="number"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  Notaire
                </label>
                <input 
                  v-model="formData.notaire_id"
                  type="number"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  PIVOT
                </label>
                <input 
                  v-model="formData.pivot_id"
                  type="number"
                  class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
                />
              </div>
            </div>
          </div>

          <!-- Onglet Notes & Statut -->
          <div v-show="activeTab === 'notes'" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                Note fixe
              </label>
              <textarea 
                v-model="formData.note_fixe"
                rows="6"
                class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all resize-none"
                placeholder="Notes importantes concernant ce client..."
              ></textarea>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-lg cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                <input 
                  v-model="formData.actif"
                  type="checkbox"
                  :true-value="1"
                  :false-value="0"
                  class="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="text-sm font-medium text-gray-700 dark:text-gray-300">Client actif</span>
              </label>

              <label class="flex items-center gap-3 p-4 border-2 border-gray-300 dark:border-gray-700 rounded-lg cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-800 transition-all">
                <input 
                  v-model="formData.dcd"
                  type="checkbox"
                  :true-value="1"
                  :false-value="0"
                  class="w-5 h-5 text-gray-600 border-gray-300 rounded focus:ring-gray-500"
                />
                <span class="text-sm font-medium text-gray-700 dark:text-gray-300">Décédé</span>
              </label>
            </div>

            <div v-if="formData.dcd === 1">
              <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                Date de décès
              </label>
              <input 
                v-model="formData.date_deces"
                type="date"
                class="w-full px-4 py-2.5 border-2 border-gray-300 dark:border-gray-700 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition-all"
              />
            </div>
          </div>

          <!-- Bouton sauvegarder -->
          <div class="mt-8 flex justify-end">
            <button
              type="submit"
              class="flex items-center gap-2 bg-gradient-to-r from-teal-600 to-teal-700 hover:from-teal-700 hover:to-teal-800 text-white px-8 py-3 rounded-xl font-semibold shadow-lg hover:shadow-xl transition-all"
            >
              <Save :size="20" />
              <span>Sauvegarder les modifications</span>
            </button>
          </div>
        </form>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { 
  User, Phone, Mail, MapPin, Heart, Building, ClipboardList, 
  FileText, UsersIcon, FolderOpen, FolderPlus, Shield, Eye,
  Save, Loader2, Smartphone, Globe
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


const props = defineProps({
  clientData: { type: Object, required: true }
})

const emit = defineEmits(['save', 'folderCreated', 'show-notes'])

const formData = ref({ ...props.clientData })
const leopardNumber = ref('')
const folderExists = ref(false)
const isCreatingFolder = ref(false)
const showContacts = ref(false)
const contactsCount = ref(0)
const activeTab = ref('identity')

// Au chargement du client
const clientStatus = ref(null)

const loadClientStatus = async () => {
  try {
    clientStatus.value = await GetClientStatus(clientId)
  } catch (error) {
    console.error('Erreur chargement statut:', error)
  }
}

const tabs = [
  { id: 'identity', label: 'Identité', icon: User },
  { id: 'contact', label: 'Coordonnées', icon: Phone },
  { id: 'address', label: 'Adresse', icon: MapPin },
  { id: 'sociocultural', label: 'Socio-culturel', icon: Globe },
  { id: 'medical', label: 'Médical', icon: Heart },
  { id: 'establishments', label: 'Établissements', icon: Building },
  { id: 'notes', label: 'Notes & Statut', icon: ClipboardList }
]

// Classes CSS basées sur la couleur
const getStatusBadgeClass = () => {
  if (!clientStatus.value) return ''
  
  const baseClass = 'px-3 py-1 rounded-full text-sm font-medium'
  
  switch (clientStatus.value.color) {
    case 'gray':
      return `${baseClass} bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300`
    case 'green':
      return `${baseClass} bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-300`
    case 'orange':
      return `${baseClass} bg-orange-100 dark:bg-orange-900/30 text-orange-700 dark:text-orange-300`
    default:
      return baseClass
  }
}
const getStatusText = () => {
  return clientStatus.value?.label || 'Inconnu'
}

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

const editMedecin = (medecin) => {
  // TODO: Ouvrir formulaire d'édition
  console.log('Éditer médecin:', medecin)
  showMedecinModal.value = false
}

const deleteMedecin = async (medecinId) => {
  try {
    await window.go.main.App.DeleteMedecin(medecinId)
    alert('✅ Médecin supprimé')
    showMedecinModal.value = false
    // Optionnel: vider la sélection si c'était le médecin sélectionné
    if (formData.value.medecin_famille_No_Licence === selectedMedecinForModal.value.licence) {
      formData.value.medecin_famille_No_Licence = ''
    }
  } catch (error) {
    console.error('Erreur suppression:', error)
    alert('❌ Erreur lors de la suppression')
  }
}

const onFolderCreated = (data) => {
  folderExists.value = true;
  leopardNumber.value = data.leopardNumber;
  formData.value.no_dossier_leopard = data.leopardNumber;
  loadFolderInfo();
  alert(`✅ Dossier créé: ${data.leopardNumber}`);
  emit("folderCreated", data);
};

const onFolderOpened = () => {
  alert("📂 Dossier ouvert dans l'explorateur");
};

const handleOpenFolder = () => {
  console.log('Ouvrir dossier:', leopardNumber.value)
}

const handleCreateFolder = () => {
  console.log('Créer dossier:', leopardNumber.value)
}
const loadFolderInfo = async () => {
  try {
    const info = await window.go.main.App.GetClientFolderInfo(leopardNumber.value);
    if (info) folderInfo.value = info;
  } catch (error) {
    console.error("❌ Erreur chargement infos dossier:", error);
  }
};

const submitHandler = () => {
  emit('save', formData.value)
}

watch(() => props.clientData, (newVal) => {
  formData.value = { ...newVal }
  leopardNumber.value = newVal.no_dossier_leopard || ''
}, { deep: true })

onMounted(() => {
  leopardNumber.value = formData.value.no_dossier_leopard || ''
})
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}

.slide-enter-from {
  transform: translateX(100%);
}

.slide-leave-to {
  transform: translateX(100%);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>