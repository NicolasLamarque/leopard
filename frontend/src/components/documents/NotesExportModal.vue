<template>
  <Transition name="fade">
    <div v-if="isOpen" class="fixed inset-0 z-[100] flex items-center justify-center bg-slate-900/80 backdrop-blur-sm p-4 md:p-10">
      
      <div class="bg-slate-100 w-full max-w-5xl h-full rounded-2xl shadow-2xl flex flex-col overflow-hidden border border-white/20">
        
        <div class="px-6 py-4 bg-white border-b flex justify-between items-center print:hidden">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-teal-600 rounded-lg text-white">
              <Eye :size="20" />
            </div>
            <div>
              <h3 class="font-bold text-slate-800 text-lg">Prévisualisation du rapport</h3>
              <p class="text-xs text-slate-500">{{ selectedNotes.length }} note(s) sélectionnée(s)</p>
            </div>
          </div>
          <button @click="$emit('close')" class="p-2 hover:bg-slate-100 rounded-full transition-colors">
            <X :size="24" class="text-slate-400" />
          </button>
        </div>

        <div class="flex-1 overflow-y-auto p-8 flex justify-center bg-slate-200/50 print:bg-white print:p-0 print:overflow-visible">
          <div id="print-area" class="bg-white shadow-xl w-[210mm] min-h-[297mm] p-[2cm] print:shadow-none print:w-full">
            
            <div v-for="(note, index) in selectedNotes" :key="note.id" 
                 :class="{'break-after-page': index < selectedNotes.length - 1}"
                 class="mb-12 last:mb-0">
              <RapportNoteContent :note="note" :client="client" />
            </div>

          </div>
        </div>

        <div class="absolute bottom-10 right-16 flex flex-col gap-4 print:hidden">
          
          <button @click="handleArchive" 
                :disabled="isArchiving"
                class="flex items-center justify-center gap-3 bg-slate-800 text-white px-6 py-4 rounded-2xl shadow-xl hover:bg-slate-700 transition-all hover:-translate-y-1 active:translate-y-0 disabled:opacity-50 border border-white/10 group">
          
          <Loader2 v-if="isArchiving" :size="20" class="animate-spin text-teal-400" />
          <FolderPlus v-else :size="20" class="text-teal-400 group-hover:scale-110 transition-transform" />
          
          <span class="font-bold text-sm tracking-wide">
            {{ isArchiving ? 'Archivage Go en cours...' : 'Ajouter au dossier (Maroto)' }}
          </span>
        </button>

          <button @click="handlePrint" 
                  class="flex items-center gap-3 bg-teal-600 text-white px-8 py-4 rounded-full shadow-2xl hover:bg-teal-500 transition-all hover:scale-110 active:scale-95">
            <Printer :size="24" />
            <span class="font-bold">Imprimer / Sauvegarder PDF</span>
          </button>
        </div>

      </div>
    </div>
  </Transition>
</template>

<script setup>
import { Printer, FolderPlus, X, Eye } from 'lucide-vue-next'

const props = defineProps({
  isOpen: Boolean,
  selectedNotes: Array,
  client: Object
})
const isArchiving = ref(false) // Pour animer le bouton "Ajouter au dossier"

const handleArchive = async () => {
  if (isArchiving.value) return // Évite les doubles clics compulsifs ;)

  isArchiving.value = true
  try {
    // Appel à ton backend Go (Maroto)
    // On passe l'ID du dossier et la liste des IDs de notes
    const path = await window.go.main.App.ExportPDFBackend(
      props.client.no_dossier_leopard, 
      props.selectedNotes.map(n => n.id)
    )
    
    alert("📁 Archivage réussi !\nLe document a été sécurisé ici :\n" + path)
  } catch (err) {
    console.error("Erreur archivage Go:", err)
    alert("Erreur lors de l'archivage. Vérifiez si le dossier est accessible.")
  } finally {
    isArchiving.value = false
  }
}
const handlePrint = () => {
  window.print()
}

</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }

@media print {
  @page { size: portrait; margin: 0; }
  body { background: white; }
  /* On cache tout ce qui n'est pas le print-area */
  :not(#print-area, #print-area *) {
    display: none !important;
  }
  #print-area {
    position: absolute;
    top: 0;
    left: 0;
    width: 100% !important;
    margin: 0 !important;
    padding: 2cm !important;
  }
}
</style>