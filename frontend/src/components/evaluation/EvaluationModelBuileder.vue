<template>
  <!-- Layout plein écran fixe : header + corps scrollable -->
  <div class="h-screen flex flex-col bg-gray-950 text-white overflow-hidden">
    <!-- ══════════ HEADER FIXE ══════════ -->
    <div class="shrink-0 border-b border-gray-800 bg-gray-950/95">
      <div class="px-6 py-3 flex items-center justify-between gap-4">
        <div class="flex items-center gap-3">
          <button
            @click="$emit('back')"
            class="p-2 hover:bg-gray-800 rounded-lg transition-colors"
          >
            <ArrowLeft :size="18" class="text-gray-400" />
          </button>
          <div class="w-px h-6 bg-gray-800" />
          <div class="p-2 bg-teal-500/20 rounded-lg">
            <Layers :size="18" class="text-teal-400" />
          </div>
          <div>
            <h1 class="text-sm font-bold leading-tight">
              Constructeur de modèles
            </h1>
            <p class="text-[10px] text-gray-500">Évaluations OTSTCFQ</p>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <button
            @click="openImportModal"
            class="flex items-center gap-2 px-3 py-1.5 bg-gray-800 hover:bg-gray-700 text-gray-300 rounded-lg text-xs font-bold transition-all border border-gray-700"
          >
            <FolderOpen :size="14" />
            Importer une base
          </button>
        </div>
        <!-- Nom + icône + couleur du modèle -->
        <div class="flex items-center gap-3 flex-1 max-w-2xl">
          <input
            v-model="model.nom"
            placeholder="Nom du modèle..."
            class="flex-1 bg-gray-900 border border-gray-700 rounded-xl px-4 py-2 text-sm font-semibold focus:outline-none focus:border-teal-500 placeholder-gray-600 text-white"
          />
          <select
            v-model="model.icone"
            class="bg-gray-900 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white"
          >
            <option v-for="ic in icones" :key="ic.val" :value="ic.val">
              {{ ic.label }}
            </option>
          </select>
          <div class="flex gap-1.5 shrink-0">
            <button
              v-for="c in couleurs"
              :key="c.val"
              @click="model.couleur = c.val"
              :class="[
                'w-5 h-5 rounded-full border-2 transition-all',
                c.bg,
                model.couleur === c.val
                  ? 'border-white scale-125'
                  : 'border-transparent opacity-50 hover:opacity-80',
              ]"
            />
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center gap-2 shrink-0">
          <button
            @click="previewMode = !previewMode"
            :class="[
              'flex items-center gap-2 px-4 py-2 rounded-xl text-sm font-semibold border transition-all',
              previewMode
                ? 'bg-teal-500/20 border-teal-500/40 text-teal-300'
                : 'border-gray-700 text-gray-400 hover:border-gray-600 hover:text-gray-200',
            ]"
          >
            <Eye :size="15" />
            {{ previewMode ? "Édition" : "Aperçu" }}
          </button>
          <button
            @click="saveModel"
            :disabled="isSaving || !model.nom.trim() || !model.sections.length"
            class="flex items-center gap-2 bg-teal-600 hover:bg-teal-500 disabled:bg-gray-800 disabled:text-gray-600 text-white px-5 py-2 rounded-xl text-sm font-bold shadow transition-all active:scale-95"
          >
            <Loader2 v-if="isSaving" :size="15" class="animate-spin" />
            <Save v-else :size="15" />
            {{ isSaving ? "Sauvegarde..." : "Sauvegarder" }}
          </button>
        </div>
      </div>
    </div>

    <!-- ══════════ CORPS : sidebar + éditeur ══════════ -->
    <div class="flex-1 flex overflow-hidden">
      <!-- ── PANNEAU GAUCHE : Sections ── -->
      <div
        class="w-64 shrink-0 border-r border-gray-800 flex flex-col bg-gray-900/40"
      >
        <!-- Header panneau -->
        <div
          class="px-4 py-3 border-b border-gray-800 flex items-center justify-between"
        >
          <span
            class="text-[10px] font-bold text-gray-500 uppercase tracking-widest"
            >Sections</span
          >
          <span class="text-[10px] text-gray-600 font-mono">{{
            model.sections.length
          }}</span>
        </div>

        <!-- Liste scrollable -->
        <div class="flex-1 overflow-y-auto p-3 space-y-1.5">
          <div
            v-for="(section, sIdx) in model.sections"
            :key="section._uid"
            :data-section-idx="sIdx"
            :class="[
              'group flex items-center gap-2 p-2.5 rounded-xl border cursor-pointer transition-all select-none',
              activeSectionIdx === sIdx
                ? 'bg-teal-600/20 border-teal-500/40 text-teal-300'
                : 'bg-gray-900 border-gray-800 text-gray-400 hover:border-gray-700 hover:text-gray-200',
              dragOverSectionIdx === sIdx && draggingSectionIdx !== sIdx
                ? 'border-teal-500 bg-teal-950/40'
                : '',
            ]"
            @click="activeSectionIdx = sIdx"
            @mousedown="onSectionMouseDown($event, sIdx)"
          >
            <GripVertical
              :size="13"
              class="text-gray-700 group-hover:text-gray-500 shrink-0"
            />
            <component
              :is="iconComponent(section.icone)"
              :size="13"
              :class="sectionIconColor(section.couleur)"
              class="shrink-0"
            />
            <span class="flex-1 text-xs font-medium truncate">{{
              section.label || "Sans titre"
            }}</span>
            <span class="text-[10px] text-gray-600 shrink-0">{{
              section.fields?.length || 0
            }}</span>
            <button
              @click.stop="removeSection(sIdx)"
              class="opacity-0 group-hover:opacity-100 p-0.5 hover:text-red-400 transition-all shrink-0"
            >
              <X :size="11" />
            </button>
          </div>

          <button
            @click="addSection"
            class="w-full flex items-center justify-center gap-1.5 p-2.5 border border-dashed border-gray-700 rounded-xl text-gray-600 hover:border-teal-700 hover:text-teal-400 text-xs font-medium transition-all mt-1"
          >
            <Plus :size="13" />
            Nouvelle section
          </button>
        </div>

        <!-- Stats -->
        <div class="shrink-0 p-3 border-t border-gray-800 space-y-1.5">
          <div class="flex justify-between text-xs">
            <span class="text-gray-600">Champs total</span>
            <span class="text-white font-bold">{{ totalFields }}</span>
          </div>
          <div class="flex justify-between text-xs">
            <span class="text-gray-600">Obligatoires</span>
            <span class="text-amber-400 font-bold">{{ requiredFields }}</span>
          </div>
        </div>
      </div>

      <!-- ── ZONE CENTRALE : Éditeur ── -->
      <div class="flex-1 overflow-y-auto">
        <div class="max-w-3xl mx-auto p-6 space-y-4">
          <!-- Aucune section -->
          <div
            v-if="!model.sections.length"
            class="flex flex-col items-center justify-center h-64 text-gray-600"
          >
            <Layers :size="48" class="mb-3 opacity-20" />
            <p class="text-sm font-medium text-gray-500">Aucune section</p>
            <p class="text-xs mt-1 text-gray-600">
              Créez votre première section
            </p>
            <button
              @click="addSection"
              class="mt-5 flex items-center gap-2 bg-teal-600 hover:bg-teal-500 text-white px-4 py-2 rounded-xl font-bold text-sm transition-all"
            >
              <Plus :size="15" /> Créer une section
            </button>
          </div>

          <template v-else-if="activeSection">
            <!-- ══ MODE APERÇU ══ -->
            <div v-if="previewMode" class="space-y-5">
              <div
                :class="[
                  'flex items-center gap-3 px-5 py-3 rounded-xl border',
                  sectionBannerClass(activeSection.couleur),
                ]"
              >
                <component
                  :is="iconComponent(activeSection.icone)"
                  :size="17"
                />
                <span class="font-bold text-sm">{{
                  activeSection.label || "Section sans titre"
                }}</span>
              </div>
              <div
                v-for="field in activeSection.fields"
                :key="field.id"
                class="space-y-1.5"
              >
                <div
                  v-if="field.type === 'info'"
                  class="p-4 bg-slate-800/50 border border-slate-700 rounded-xl text-sm text-slate-400"
                >
                  {{ field.contenu || "Bloc info..." }}
                </div>
                <template v-else>
                  <label class="block text-sm font-semibold text-gray-300"
                    >{{ field.label || "Champ sans titre"
                    }}<span v-if="field.required" class="text-red-400 ml-1"
                      >*</span
                    ></label
                  >
                  <p v-if="field.help" class="text-xs text-gray-500">
                    {{ field.help }}
                  </p>
                  <textarea
                    v-if="field.type === 'textarea'"
                    disabled
                    :rows="field.rows || 4"
                    :placeholder="field.placeholder || ''"
                    class="w-full bg-gray-900 border border-gray-700 rounded-xl px-4 py-3 text-sm text-gray-600 resize-none"
                  />
                  <input
                    v-else-if="field.type === 'text'"
                    disabled
                    type="text"
                    :placeholder="field.placeholder || ''"
                    class="w-full bg-gray-900 border border-gray-700 rounded-xl px-4 py-2.5 text-sm text-gray-600"
                  />
                  <input
                    v-else-if="field.type === 'date'"
                    disabled
                    type="date"
                    class="bg-gray-900 border border-gray-700 rounded-xl px-4 py-2.5 text-sm text-gray-600"
                  />
                  <select
                    v-else-if="field.type === 'select'"
                    disabled
                    class="w-full bg-gray-900 border border-gray-700 rounded-xl px-4 py-2.5 text-sm text-gray-600"
                  >
                    <option value="">— Sélectionner —</option>
                    <option v-for="opt in field.options || []" :key="opt">
                      {{ opt }}
                    </option>
                  </select>
                  <div
                    v-else-if="field.type === 'checkboxes'"
                    class="grid grid-cols-2 gap-2"
                  >
                    <div
                      v-for="opt in field.options || []"
                      :key="opt"
                      class="flex items-center gap-2 p-2.5 bg-gray-900 border border-gray-700 rounded-lg text-xs text-gray-500"
                    >
                      <div
                        class="w-3.5 h-3.5 rounded border border-gray-600 shrink-0"
                      />
                      {{ opt }}
                    </div>
                  </div>
                  <div
                    v-else-if="field.type === 'table'"
                    class="border border-gray-700 rounded-xl overflow-hidden"
                  >
                    <div
                      v-if="(field.columns || []).length"
                      class="grid bg-gray-900 border-b border-gray-700 px-3 py-2"
                      :style="{
                        gridTemplateColumns: `repeat(${field.columns.length}, 1fr)`,
                      }"
                    >
                      <span
                        v-for="col in field.columns"
                        :key="col.id"
                        class="text-[10px] font-bold text-gray-500 uppercase"
                        >{{ col.label }}</span
                      >
                    </div>
                    <div class="px-4 py-5 text-center text-xs text-gray-600">
                      Tableau — {{ (field.columns || []).length }} colonne(s)
                    </div>
                  </div>
                </template>
              </div>
            </div>

            <!-- ══ MODE ÉDITION ══ -->
            <div v-else class="space-y-4">
              <!-- Config section -->
              <div
                class="bg-gray-900 border border-gray-800 rounded-2xl p-4 space-y-4"
              >
                <p
                  class="text-[10px] font-bold text-gray-500 uppercase tracking-widest"
                >
                  Section active
                </p>
                <div class="grid grid-cols-2 gap-3">
                  <div>
                    <label
                      class="block text-xs font-semibold text-gray-400 mb-1.5"
                      >Titre *</label
                    >
                    <input
                      v-model="activeSection.label"
                      placeholder="Ex: Contexte et Référence"
                      class="w-full bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                    />
                  </div>
                  <div>
                    <label
                      class="block text-xs font-semibold text-gray-400 mb-1.5"
                      >Identifiant (auto)</label
                    >
                    <input
                      :value="activeSection.id"
                      disabled
                      class="w-full bg-gray-800/40 border border-gray-800 rounded-xl px-3 py-2 text-xs text-gray-600 font-mono"
                    />
                  </div>
                </div>
                <div class="flex items-center gap-6">
                  <div>
                    <label
                      class="block text-xs font-semibold text-gray-400 mb-1.5"
                      >Icône</label
                    >
                    <select
                      v-model="activeSection.icone"
                      class="bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white"
                    >
                      <option
                        v-for="ic in icones"
                        :key="ic.val"
                        :value="ic.val"
                      >
                        {{ ic.label }}
                      </option>
                    </select>
                  </div>
                  <div>
                    <label
                      class="block text-xs font-semibold text-gray-400 mb-2"
                      >Couleur</label
                    >
                    <div class="flex gap-2">
                      <button
                        v-for="c in couleurs"
                        :key="c.val"
                        @click="activeSection.couleur = c.val"
                        :class="[
                          'w-5 h-5 rounded-full border-2 transition-all',
                          c.bg,
                          activeSection.couleur === c.val
                            ? 'border-white scale-125'
                            : 'border-transparent opacity-50',
                        ]"
                      />
                    </div>
                  </div>
                </div>
              </div>

              <!-- Champs -->
              <div class="space-y-3">
                <span
                  class="text-[10px] font-bold text-gray-500 uppercase tracking-widest"
                  >Champs — {{ activeSection.fields?.length || 0 }}</span
                >

                <div class="space-y-2 min-h-[20px]">
                  <div
                    v-for="(field, fIdx) in activeSection.fields"
                    :key="field.id"
                    :data-field-idx="fIdx"
                    :class="[
                      'bg-gray-900 border rounded-2xl overflow-hidden transition-all',
                      draggingFieldIdx === fIdx
                        ? 'opacity-30 border-teal-500'
                        : 'border-gray-800 hover:border-gray-700',
                      dragOverFieldIdx === fIdx && draggingFieldIdx !== fIdx
                        ? 'border-teal-500 bg-teal-950/30'
                        : '',
                    ]"
                  >
                    <!-- Header champ -->
                    <div
                      class="flex items-center gap-2 px-4 py-2.5 border-b border-gray-800/60 cursor-grab"
                      @mousedown="onFieldMouseDown($event, fIdx)"
                    >
                      <GripVertical :size="13" class="text-gray-600 shrink-0" />
                      <span
                        :class="[
                          'text-[9px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wide shrink-0',
                          fieldTypeBadge(field.type),
                        ]"
                        >{{ fieldTypeLabel(field.type) }}</span
                      >
                      <span
                        class="flex-1 text-xs font-medium text-gray-300 truncate"
                        >{{
                          field.label || field.contenu || "Sans titre"
                        }}</span
                      >
                      <div class="flex items-center gap-0.5 shrink-0">
                        <button
                          @click="toggleFieldExpand(field.id)"
                          class="p-1 text-gray-500 hover:text-gray-300"
                        >
                          <ChevronDown
                            :size="13"
                            :class="
                              expandedFields.has(field.id) ? 'rotate-180' : ''
                            "
                            class="transition-transform"
                          />
                        </button>
                        <button
                          @click="moveFieldUp(fIdx)"
                          :disabled="fIdx === 0"
                          class="p-1 text-gray-600 hover:text-gray-300 disabled:opacity-20"
                        >
                          <ChevronUp :size="13" />
                        </button>
                        <button
                          @click="moveFieldDown(fIdx)"
                          :disabled="fIdx === activeSection.fields.length - 1"
                          class="p-1 text-gray-600 hover:text-gray-300 disabled:opacity-20"
                        >
                          <ChevronDown :size="13" />
                        </button>
                        <button
                          @click="removeField(fIdx)"
                          class="p-1 text-gray-700 hover:text-red-400"
                        >
                          <Trash2 :size="13" />
                        </button>
                      </div>
                    </div>

                    <!-- Corps champ -->
                    <div
                      v-if="expandedFields.has(field.id)"
                      class="p-4 space-y-3"
                    >
                      <template v-if="field.type === 'info'">
                        <label
                          class="block text-xs font-semibold text-gray-400 mb-1"
                          >Contenu du bloc</label
                        >
                        <textarea
                          v-model="field.contenu"
                          rows="3"
                          placeholder="Note légale, instruction..."
                          class="w-full bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600 resize-none"
                        />
                      </template>
                      <template v-else>
                        <div class="grid grid-cols-2 gap-3">
                          <div>
                            <label
                              class="block text-xs font-semibold text-gray-400 mb-1"
                              >Label *</label
                            ><input
                              v-model="field.label"
                              placeholder="Libellé..."
                              class="w-full bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                            />
                          </div>
                          <div>
                            <label
                              class="block text-xs font-semibold text-gray-400 mb-1"
                              >Identifiant</label
                            ><input
                              v-model="field.id"
                              placeholder="snake_case..."
                              class="w-full bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-xs font-mono focus:outline-none focus:border-teal-500 text-gray-400 placeholder-gray-600"
                            />
                          </div>
                        </div>
                        <div class="grid grid-cols-2 gap-3">
                          <div>
                            <label
                              class="block text-xs font-semibold text-gray-400 mb-1"
                              >Texte d'aide</label
                            ><input
                              v-model="field.help"
                              placeholder="Description..."
                              class="w-full bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                            />
                          </div>
                          <div>
                            <label
                              class="block text-xs font-semibold text-gray-400 mb-1"
                              >Placeholder</label
                            ><input
                              v-model="field.placeholder"
                              placeholder="Texte indicatif..."
                              class="w-full bg-gray-800 border border-gray-700 rounded-xl px-3 py-2 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                            />
                          </div>
                        </div>
                        <!-- TEXTAREA -->
                        <div
                          v-if="field.type === 'textarea'"
                          class="flex items-center gap-4"
                        >
                          <div class="flex items-center gap-2">
                            <label class="text-xs text-gray-400">Lignes</label
                            ><input
                              v-model.number="field.rows"
                              type="number"
                              min="2"
                              max="20"
                              class="w-16 bg-gray-800 border border-gray-700 rounded-lg px-2 py-1.5 text-sm text-center focus:outline-none focus:border-teal-500 text-white"
                            />
                          </div>
                          <label
                            class="flex items-center gap-2 text-xs text-gray-400 cursor-pointer"
                            ><input
                              type="checkbox"
                              v-model="field.required"
                              class="accent-teal-500"
                            />
                            Obligatoire</label
                          >
                        </div>
                        <!-- TEXT / DATE -->
                        <div
                          v-if="field.type === 'text' || field.type === 'date'"
                        >
                          <label
                            class="flex items-center gap-2 text-xs text-gray-400 cursor-pointer"
                            ><input
                              type="checkbox"
                              v-model="field.required"
                              class="accent-teal-500"
                            />
                            Obligatoire</label
                          >
                        </div>
                        <!-- SELECT -->
                        <div v-if="field.type === 'select'" class="space-y-2">
                          <div class="flex items-center justify-between">
                            <label class="text-xs font-semibold text-gray-400"
                              >Options</label
                            ><button
                              @click="addOption(field)"
                              class="text-xs text-teal-400 hover:text-teal-300 flex items-center gap-1"
                            >
                              <Plus :size="11" /> Ajouter
                            </button>
                          </div>
                          <div class="space-y-1.5">
                            <div
                              v-for="(opt, oIdx) in field.options || []"
                              :key="oIdx"
                              class="flex items-center gap-2"
                            >
                              <input
                                :value="opt"
                                @input="
                                  field.options[oIdx] = $event.target.value
                                "
                                placeholder="Option..."
                                class="flex-1 bg-gray-800 border border-gray-700 rounded-lg px-3 py-1.5 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                              />
                              <button
                                @click="removeOption(field, oIdx)"
                                class="text-gray-600 hover:text-red-400"
                              >
                                <X :size="12" />
                              </button>
                            </div>
                            <p
                              v-if="!field.options?.length"
                              class="text-xs text-gray-700 text-center py-2"
                            >
                              Aucune option
                            </p>
                          </div>
                          <label
                            class="flex items-center gap-2 text-xs text-gray-400 cursor-pointer"
                            ><input
                              type="checkbox"
                              v-model="field.required"
                              class="accent-teal-500"
                            />
                            Obligatoire</label
                          >
                        </div>
                        <!-- CHECKBOXES -->
                        <div
                          v-if="field.type === 'checkboxes'"
                          class="space-y-2"
                        >
                          <div class="flex items-center justify-between">
                            <label class="text-xs font-semibold text-gray-400"
                              >Cases</label
                            ><button
                              @click="addOption(field)"
                              class="text-xs text-teal-400 hover:text-teal-300 flex items-center gap-1"
                            >
                              <Plus :size="11" /> Ajouter
                            </button>
                          </div>
                          <div class="grid grid-cols-2 gap-1.5">
                            <div
                              v-for="(opt, oIdx) in field.options || []"
                              :key="oIdx"
                              class="flex items-center gap-1.5"
                            >
                              <input
                                :value="opt"
                                @input="
                                  field.options[oIdx] = $event.target.value
                                "
                                placeholder="Option..."
                                class="flex-1 bg-gray-800 border border-gray-700 rounded-lg px-2 py-1.5 text-xs focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                              />
                              <button
                                @click="removeOption(field, oIdx)"
                                class="text-gray-600 hover:text-red-400"
                              >
                                <X :size="11" />
                              </button>
                            </div>
                          </div>
                        </div>
                        <!-- TABLE -->
                        <div v-if="field.type === 'table'" class="space-y-2">
                          <div class="flex items-center justify-between">
                            <label class="text-xs font-semibold text-gray-400"
                              >Colonnes</label
                            ><button
                              @click="addColumn(field)"
                              class="text-xs text-teal-400 hover:text-teal-300 flex items-center gap-1"
                            >
                              <Plus :size="11" /> Colonne
                            </button>
                          </div>
                          <div class="space-y-1.5">
                            <div
                              v-for="(col, cIdx) in field.columns || []"
                              :key="cIdx"
                              class="flex items-center gap-2"
                            >
                              <input
                                v-model="col.label"
                                placeholder="Entête..."
                                class="flex-1 bg-gray-800 border border-gray-700 rounded-lg px-3 py-1.5 text-sm focus:outline-none focus:border-teal-500 text-white placeholder-gray-600"
                              />
                              <select
                                v-model="col.type"
                                class="bg-gray-800 border border-gray-700 rounded-lg px-2 py-1.5 text-xs focus:outline-none focus:border-teal-500 text-white"
                              >
                                <option value="text">Texte</option>
                                <option value="date">Date</option>
                              </select>
                              <input
                                v-model="col.id"
                                placeholder="id..."
                                class="w-20 bg-gray-800 border border-gray-700 rounded-lg px-2 py-1.5 text-xs font-mono focus:outline-none focus:border-teal-500 text-gray-500 placeholder-gray-700"
                              />
                              <button
                                @click="removeColumn(field, cIdx)"
                                class="text-gray-600 hover:text-red-400"
                              >
                                <X :size="12" />
                              </button>
                            </div>
                            <p
                              v-if="!field.columns?.length"
                              class="text-xs text-gray-700 text-center py-2"
                            >
                              Aucune colonne
                            </p>
                          </div>
                        </div>
                      </template>
                    </div>
                  </div>
                </div>

                <!-- Palette ajout champs -->
                <div
                  class="border border-dashed border-gray-800 rounded-2xl p-4"
                >
                  <p
                    class="text-[10px] font-bold text-gray-700 uppercase tracking-widest mb-3"
                  >
                    Ajouter un champ
                  </p>
                  <div class="flex flex-wrap gap-2">
                    <button
                      v-for="ft in fieldTypes"
                      :key="ft.type"
                      @click="addField(ft.type)"
                      :class="[
                        'flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-semibold border transition-all',
                        ft.btnClass,
                      ]"
                    >
                      <component :is="ft.icon" :size="11" />{{ ft.label }}
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </div>
      </div>
    </div>

    <!-- Ghost drag visuel -->
    <div
      v-if="dragGhost.visible"
      class="fixed pointer-events-none z-50 bg-gray-800 border border-teal-500 rounded-xl px-4 py-2 text-xs font-semibold text-teal-300 shadow-2xl"
      :style="{
        left: dragGhost.x + 'px',
        top: dragGhost.y + 'px',
        transform: 'translate(-50%, -50%)',
      }"
    >
      {{ dragGhost.label }}
    </div>

    <!-- Toast -->
    <Transition name="toast">
      <div
        v-if="toast.show"
        :class="[
          'fixed bottom-6 right-6 flex items-center gap-3 px-5 py-3 rounded-2xl shadow-2xl border text-sm font-semibold z-50',
          toast.type === 'success'
            ? 'bg-emerald-900 border-emerald-700 text-emerald-200'
            : 'bg-red-900 border-red-700 text-red-200',
        ]"
      >
        <CheckCircle v-if="toast.type === 'success'" :size="16" />
        <AlertCircle v-else :size="16" />
        {{ toast.message }}
      </div>
    </Transition>

<Transition name="modal-fade">
  <div v-if="showImportModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6">
    <div class="absolute inset-0 bg-gray-950/80 backdrop-blur-sm" @click="showImportModal = false" />
    
    <div class="relative bg-gray-900 border border-gray-800 w-full max-w-2xl max-h-[80vh] rounded-2xl shadow-2xl flex flex-col overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-800 flex justify-between items-center bg-gray-900/50">
        <h3 class="font-bold text-teal-400 flex items-center gap-2">
          <FolderOpen :size="18" />
          Choisir un modèle comme base
        </h3>
        <button @click="showImportModal = false" class="text-gray-500 hover:text-white">
          <X :size="20" />
        </button>
      </div>

      <div class="flex-1 overflow-y-auto p-4 space-y-2">
        <div v-if="availableModels.length === 0" class="text-center py-10 text-gray-500">
          Aucun modèle trouvé en base de données.
        </div>
        
        <button 
          v-for="m in availableModels" 
          :key="m.id"
          @click="confirmImport(m)"
          class="w-full flex items-center gap-4 p-3 rounded-xl bg-gray-800/50 border border-gray-700/50 hover:border-teal-500/50 hover:bg-gray-800 transition-all group text-left"
        >
          <div 
  :class="[
    `p-2 rounded-lg bg-${m.couleur || 'teal'}-500/20 text-${m.couleur || 'teal'}-400`,
    'group-hover:scale-110 transition-transform'
  ]"
>
             <component :is="iconComponent(m.icone)" :size="20" />
          </div>
          <div class="flex-1">
            <div class="font-bold text-sm text-gray-200">{{ m.nom }}</div>
            <div class="text-[10px] text-gray-500 uppercase tracking-wider">Version {{ m.version }}</div>
          </div>
          <ChevronRight :size="16" class="text-gray-600 group-hover:text-teal-400 group-hover:translate-x-1 transition-all" />
        </button>
      </div>
    </div>
  </div>
</Transition>


  </div>
</template>

<script setup>
import { ref, computed, reactive, watch } from "vue";
import {
  ArrowLeft,
  Layers,
  Eye,
  Save,
  Loader2,
  Plus,
  X,
  Trash2,
  GripVertical,
  ChevronDown,
  ChevronUp,
  ChevronRight,
  CheckCircle,
  AlertCircle,
  FileText,
  Brain,
  Heart,
  Activity,
  Briefcase,
  Users,
  Target,
  Network,
  ClipboardList,
  Scale,
  ShieldCheck,
  AlignLeft,
  Type,
  Calendar,
  List,
  CheckSquare,
  Table,
  Info,
} from "lucide-vue-next";

defineEmits(["back"]);

const model = reactive({
  id: "",
  nom: "",
  icone: "ClipboardList",
  couleur: "teal",
  sections: [],
});
const activeSectionIdx = ref(0);
const previewMode = ref(false);
const isSaving = ref(false);
const expandedFields = ref(new Set());
const toast = reactive({ show: false, message: "", type: "success" });

// ── Drag & Drop SOURIS (fiable WebView2) ────────────────────
const draggingSectionIdx = ref(null);
const dragOverSectionIdx = ref(null);
const draggingFieldIdx = ref(null);
const dragOverFieldIdx = ref(null);
const dragGhost = reactive({ visible: false, x: 0, y: 0, label: "" });

const onSectionMouseDown = (e, idx) => {
  if (e.button !== 0 || e.target.closest("button")) return;
  draggingSectionIdx.value = idx;
  dragGhost.label = model.sections[idx]?.label || "Section";
  dragGhost.x = e.clientX;
  dragGhost.y = e.clientY;
  dragGhost.visible = false;
  const onMove = (ev) => {
    dragGhost.x = ev.clientX;
    dragGhost.y = ev.clientY;
    dragGhost.visible = true;
    let found = null;
    for (const el of document.querySelectorAll("[data-section-idx]")) {
      const r = el.getBoundingClientRect();
      if (ev.clientY >= r.top && ev.clientY <= r.bottom) {
        found = parseInt(el.getAttribute("data-section-idx"));
        break;
      }
    }
    dragOverSectionIdx.value = found;
  };
  const onUp = () => {
    if (
      draggingSectionIdx.value !== null &&
      dragOverSectionIdx.value !== null &&
      draggingSectionIdx.value !== dragOverSectionIdx.value
    ) {
      const [from, to] = [draggingSectionIdx.value, dragOverSectionIdx.value];
      const item = model.sections.splice(from, 1)[0];
      model.sections.splice(to, 0, item);
      activeSectionIdx.value = to;
    }
    draggingSectionIdx.value = null;
    dragOverSectionIdx.value = null;
    dragGhost.visible = false;
    window.removeEventListener("mousemove", onMove);
    window.removeEventListener("mouseup", onUp);
  };
  window.addEventListener("mousemove", onMove);
  window.addEventListener("mouseup", onUp);
};

const onFieldMouseDown = (e, idx) => {
  if (e.button !== 0 || e.target.closest("button")) return;
  draggingFieldIdx.value = idx;
  dragGhost.label = activeSection.value?.fields[idx]?.label || "Champ";
  dragGhost.x = e.clientX;
  dragGhost.y = e.clientY;
  dragGhost.visible = false;
  const onMove = (ev) => {
    dragGhost.x = ev.clientX;
    dragGhost.y = ev.clientY;
    dragGhost.visible = true;
    let found = null;
    for (const el of document.querySelectorAll("[data-field-idx]")) {
      const r = el.getBoundingClientRect();
      if (ev.clientY >= r.top && ev.clientY <= r.bottom) {
        found = parseInt(el.getAttribute("data-field-idx"));
        break;
      }
    }
    dragOverFieldIdx.value = found;
  };
  const onUp = () => {
    if (
      draggingFieldIdx.value !== null &&
      dragOverFieldIdx.value !== null &&
      draggingFieldIdx.value !== dragOverFieldIdx.value &&
      activeSection.value
    ) {
      const [from, to] = [draggingFieldIdx.value, dragOverFieldIdx.value];
      const fields = activeSection.value.fields;
      const item = fields.splice(from, 1)[0];
      fields.splice(to, 0, item);
    }
    draggingFieldIdx.value = null;
    dragOverFieldIdx.value = null;
    dragGhost.visible = false;
    window.removeEventListener("mousemove", onMove);
    window.removeEventListener("mouseup", onUp);
  };
  window.addEventListener("mousemove", onMove);
  window.addEventListener("mouseup", onUp);
};

// ── Computed ────────────────────────────────────────────────
const activeSection = computed(
  () => model.sections[activeSectionIdx.value] || null,
);
const totalFields = computed(() =>
  model.sections.reduce((a, s) => a + (s.fields?.length || 0), 0),
);
const requiredFields = computed(() =>
  model.sections.reduce(
    (a, s) => a + (s.fields?.filter((f) => f.required)?.length || 0),
    0,
  ),
);

// ── Constantes ──────────────────────────────────────────────
const icones = [
  { val: "ClipboardList", label: "📋 Évaluation" },
  { val: "FileText", label: "📄 Document" },
  { val: "Brain", label: "🧠 Cognitif" },
  { val: "Heart", label: "❤️ Santé" },
  { val: "Users", label: "👥 Réseau" },
  { val: "Scale", label: "⚖️ Juridique" },
  { val: "Target", label: "🎯 Objectifs" },
  { val: "ShieldCheck", label: "🛡️ Protection" },
];
const couleurs = [
  { val: "teal", bg: "bg-teal-500" },
  { val: "blue", bg: "bg-blue-500" },
  { val: "purple", bg: "bg-purple-500" },
  { val: "amber", bg: "bg-amber-500" },
  { val: "emerald", bg: "bg-emerald-500" },
  { val: "red", bg: "bg-red-500" },
  { val: "slate", bg: "bg-slate-500" },
];
const fieldTypes = [
  {
    type: "textarea",
    label: "Texte long",
    icon: AlignLeft,
    btnClass:
      "bg-blue-500/10 border-blue-500/30 text-blue-300 hover:bg-blue-500/20",
  },
  {
    type: "text",
    label: "Texte court",
    icon: Type,
    btnClass:
      "bg-teal-500/10 border-teal-500/30 text-teal-300 hover:bg-teal-500/20",
  },
  {
    type: "date",
    label: "Date",
    icon: Calendar,
    btnClass:
      "bg-purple-500/10 border-purple-500/30 text-purple-300 hover:bg-purple-500/20",
  },
  {
    type: "select",
    label: "Liste",
    icon: List,
    btnClass:
      "bg-amber-500/10 border-amber-500/30 text-amber-300 hover:bg-amber-500/20",
  },
  {
    type: "checkboxes",
    label: "Cases",
    icon: CheckSquare,
    btnClass:
      "bg-emerald-500/10 border-emerald-500/30 text-emerald-300 hover:bg-emerald-500/20",
  },
  {
    type: "table",
    label: "Tableau",
    icon: Table,
    btnClass:
      "bg-pink-500/10 border-pink-500/30 text-pink-300 hover:bg-pink-500/20",
  },
  {
    type: "info",
    label: "Bloc info",
    icon: Info,
    btnClass:
      "bg-slate-500/10 border-slate-500/30 text-slate-300 hover:bg-slate-500/20",
  },
];

// ── Icônes & couleurs ────────────────────────────────────────
const iconMap = {
  FileText,
  Brain,
  Heart,
  Activity,
  Briefcase,
  Users,
  Target,
  Network,
  ClipboardList,
  Scale,
  ShieldCheck,
};
const iconComponent = (n) => iconMap[n] || FileText;
const colorIconMap = {
  teal: "text-teal-400",
  blue: "text-blue-400",
  purple: "text-purple-400",
  amber: "text-amber-400",
  emerald: "text-emerald-400",
  red: "text-red-400",
  slate: "text-slate-400",
};
const sectionIconColor = (c) => colorIconMap[c] || "text-gray-400";
const sectionBannerMap = {
  teal: "bg-teal-900/30 border-teal-700/40 text-teal-300",
  blue: "bg-blue-900/30 border-blue-700/40 text-blue-300",
  purple: "bg-purple-900/30 border-purple-700/40 text-purple-300",
  amber: "bg-amber-900/30 border-amber-700/40 text-amber-300",
  emerald: "bg-emerald-900/30 border-emerald-700/40 text-emerald-300",
  red: "bg-red-900/30 border-red-700/40 text-red-300",
  slate: "bg-slate-800/50 border-slate-700/40 text-slate-300",
};
const sectionBannerClass = (c) => sectionBannerMap[c] || sectionBannerMap.teal;
const fieldTypeBadge = (t) =>
  ({
    textarea: "bg-blue-500/20 text-blue-300",
    text: "bg-teal-500/20 text-teal-300",
    date: "bg-purple-500/20 text-purple-300",
    select: "bg-amber-500/20 text-amber-300",
    checkboxes: "bg-emerald-500/20 text-emerald-300",
    table: "bg-pink-500/20 text-pink-300",
    info: "bg-slate-500/20 text-slate-300",
  })[t] || "bg-gray-500/20 text-gray-300";
const fieldTypeLabel = (t) =>
  ({
    textarea: "Texte long",
    text: "Texte",
    date: "Date",
    select: "Liste",
    checkboxes: "Cases",
    table: "Tableau",
    info: "Info",
  })[t] || t;

// ── Helpers ──────────────────────────────────────────────────
const uid = () => Math.random().toString(36).slice(2, 9);
const slug = (s) =>
  (s || "")
    .toLowerCase()
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/[^a-z0-9]+/g, "_")
    .replace(/^_|_$/g, "");
const showToast = (message, type = "success") => {
  Object.assign(toast, { show: true, message, type });
  setTimeout(() => {
    toast.show = false;
  }, 3000);
};

// ── Sections ─────────────────────────────────────────────────
const addSection = () => {
  model.sections.push({
    _uid: uid(),
    id: `section_${uid()}`,
    label: "",
    icone: "FileText",
    couleur: "teal",
    fields: [],
  });
  activeSectionIdx.value = model.sections.length - 1;
};
const removeSection = (idx) => {
  model.sections.splice(idx, 1);
  if (activeSectionIdx.value >= model.sections.length)
    activeSectionIdx.value = Math.max(0, model.sections.length - 1);
};

// ── Champs ────────────────────────────────────────────────────
const addField = (type) => {
  if (!activeSection.value) return;
  const field = {
    id: `${type}_${uid()}`,
    type,
    label: "",
    help: "",
    placeholder: "",
    required: false,
  };
  if (type === "textarea") field.rows = 5;
  if (["select", "checkboxes"].includes(type)) field.options = [];
  if (type === "table") field.columns = [];
  if (type === "info") field.contenu = "";
  activeSection.value.fields.push(field);
  expandedFields.value = new Set([...expandedFields.value, field.id]);
};
const removeField = (idx) => {
  activeSection.value.fields.splice(idx, 1);
};
const moveFieldUp = (idx) => {
  if (idx === 0) return;
  const f = activeSection.value.fields;
  [f[idx - 1], f[idx]] = [f[idx], f[idx - 1]];
};
const moveFieldDown = (idx) => {
  const f = activeSection.value.fields;
  if (idx >= f.length - 1) return;
  [f[idx], f[idx + 1]] = [f[idx + 1], f[idx]];
};
const toggleFieldExpand = (id) => {
  const s = new Set(expandedFields.value);
  s.has(id) ? s.delete(id) : s.add(id);
  expandedFields.value = s;
};

// ── Options & colonnes ────────────────────────────────────────
const addOption = (f) => {
  if (!f.options) f.options = [];
  f.options.push("");
};
const removeOption = (f, i) => {
  f.options.splice(i, 1);
};
const addColumn = (f) => {
  if (!f.columns) f.columns = [];
  f.columns.push({ id: `col_${uid()}`, label: "", type: "text" });
};
const removeColumn = (f, i) => {
  f.columns.splice(i, 1);
};

// ── Auto-slug ─────────────────────────────────────────────────
watch(
  () => activeSection.value?.label,
  (newLabel) => {
    if (activeSection.value && !activeSection.value.id_manual_edited)
      activeSection.value.id = slug(newLabel);
  },
);

// ── Sauvegarde ────────────────────────────────────────────────
const saveModel = async () => {
  if (!model.nom?.trim()) {
    showToast("Le nom du modèle est requis", "error");
    return;
  }
  if (!model.sections.length) {
    showToast("Ajoutez au moins une section", "error");
    return;
  }
  isSaving.value = true;
  try {
    await window.go.main.EvalHandler.SaveDefinition({
      id: model.id || slug(model.nom) || `model_${uid()}`,
      nom: model.nom,
      icone: model.icone,
      couleur: model.couleur,
      schema_json: JSON.stringify({ sections: model.sections }),
      version: 1,
      active: 1,
    });
    showToast("Modèle sauvegardé !");
  } catch (err) {
    console.error(err);
    showToast(err?.message || "Erreur sauvegarde", "error");
  } finally {
    isSaving.value = false;
  }
};

// ── Chargement ───────────────────────────────────────────────d'un modèle existant
/**
 * Charge un modèle existant pour s'en servir de base
 * @param {Object} templateModel - Le modèle sélectionné dans une liste
 */
const useAsTemplate = (templateModel) => {
  // 1. On parse le schéma JSON du modèle source
  const sourceSchema = JSON.parse(templateModel.schema_json);

  // 2. On remplit notre état réactif 'model' avec les données sources
  model.nom = `${templateModel.nom} (Copie)`; // On change le nom pour éviter la confusion
  model.icone = templateModel.icone;
  model.couleur = templateModel.couleur;

  // 3. IMPORTANT : On copie les sections mais on laisse model.id VIDE
  // pour que window.go.main.EvalHandler.SaveDefinition crée un NOUVEL enregistrement
  model.id = "";
  model.sections = JSON.parse(JSON.stringify(sourceSchema.sections)); // Deep copy des sections

  showToast(
    "Modèle chargé comme base. Modifiez-le et sauvegardez pour créer un nouveau modèle.",
  );
};

// Dans <script setup> de EvaluationModelBuilder.vue

const availableModels = ref([])
const showImportModal = ref(false)

const openImportModal = async () => {
  try {
    // On appelle ton backend Go pour lister les modèles existants
    const defs = await window.go.main.EvalHandler.GetDefinitions()
    availableModels.value = defs
    showImportModal.value = true
  } catch (err) {
    showToast("Impossible de charger les modèles", "error")
  }
}

const confirmImport = (selectedTemplate) => {
  // On utilise la logique de "Deep Copy" qu'on a vue
  const schema = JSON.parse(selectedTemplate.schema_json)
  
  model.nom = `${selectedTemplate.nom} (Copie)`
  model.icone = selectedTemplate.icone
  model.couleur = selectedTemplate.couleur
  model.sections = JSON.parse(JSON.stringify(schema.sections))
  
  // CRUCIAL : On reset l'ID pour que ce soit un NOUVEAU modèle
  model.id = "" 
  
  showImportModal.value = false
  showToast("Base importée avec succès !")
}
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}
.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translateY(12px);
}
.rotate-180 {
  transform: rotate(180deg);
}
</style>
