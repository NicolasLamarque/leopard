<template>
  <div class="p-2 bg-blue-50 dark:bg-blue-900/20 rounded-lg border border-blue-200 dark:border-blue-800">
    
    <!-- Badge de statut -->
    <div class="flex items-center justify-between mb-2">
      <div class="flex items-center gap-2">
        <FileText :size="14" class="text-blue-600" />
        <span class="text-xs font-semibold text-gray-700 dark:text-gray-300">
          Fiche d'identification
        </span>
      </div>
      
      <!-- Statut: Présente ✓ -->
      <div v-if="ficheExists" class="flex items-center gap-1 text-xs text-green-600 dark:text-green-400">
        <CheckCircle :size="14" />
        <span class="font-semibold">Présente</span>
      </div>
      
      <!-- Statut: Manquante -->
      <div v-else class="flex items-center gap-1 text-xs text-orange-600 dark:text-orange-400">
        <AlertCircle :size="14" />
        <span class="font-semibold">Manquante</span>
      </div>
    </div>

    <!-- Actions -->
    <div v-if="!ficheExists" class="space-y-2">
      
      <!-- Validation des données -->
      <div v-if="!isDonneesCompletes" class="p-2 bg-yellow-50 dark:bg-yellow-900/20 rounded border border-yellow-300 dark:border-yellow-700">
        <div class="flex items-start gap-2">
          <AlertTriangle :size="12" class="text-yellow-600 flex-shrink-0 mt-0.5" />
          <div class="flex-1">
            <p class="text-[10px] font-semibold text-yellow-800 dark:text-yellow-300 mb-1">
              Données incomplètes
            </p>
            <ul class="text-[9px] text-yellow-700 dark:text-yellow-400 space-y-0.5">
              <li v-for="champ in champsManquants" :key="champ">• {{ champ }}</li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Bouton créer -->
      <button
        @click="creerFiche"
        :disabled="!isDonneesCompletes || isGenerating"
        class="w-full flex items-center justify-center gap-2 px-3 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-300 disabled:cursor-not-allowed text-white rounded-lg text-xs font-semibold transition-all"
      >
        <FilePlus v-if="!isGenerating" :size="14" />
        <Loader2 v-else :size="14" class="animate-spin" />
        <span>{{ isGenerating ? 'Génération...' : 'Créer la fiche' }}</span>
      </button>
    </div>

    <!-- Si fiche existe: bouton ouvrir -->
    <button
      v-else
      @click="ouvrirFiche"
      class="w-full flex items-center justify-center gap-2 px-3 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-xs font-semibold transition-all"
    >
      <ExternalLink :size="14" />
      <span>Ouvrir la fiche</span>
    </button>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  FileText, 
  CheckCircle, 
  AlertCircle, 
  AlertTriangle,
  FilePlus,
  ExternalLink,
  Loader2 
} from 'lucide-vue-next'
import { jsPDF } from 'jspdf'
import logoPath from '@/assets/images/LeopardLogo.png'

const props = defineProps({
  client: {
    type: Object,
    required: true
  },
  leopardNumber: {
    type: String,
    required: true
  },
  subfolderPath: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['fiche-created', 'refresh'])

const ficheExists = ref(false)
const isGenerating = ref(false)

// Champs obligatoires pour la fiche
const champsRequis = {
  nom: 'Nom',
  prenom: 'Prénom',
  date_naissance: 'Date de naissance',
  numero_assurance_maladie: 'No. RAMQ',
  adresse: 'Adresse',
  ville: 'Ville',
  code_postal: 'Code postal'
}

// Validation des données
const champsManquants = computed(() => {
  const manquants = []
  
  Object.entries(champsRequis).forEach(([key, label]) => {
    const valeur = props.client[key]
    if (!valeur || valeur.trim() === '') {
      manquants.push(label)
    }
  })
  
  return manquants
})

const isDonneesCompletes = computed(() => {
  return champsManquants.value.length === 0
})

// Vérifier l'existence de la fiche
const verifierExistence = async () => {
  try {
    const nomFichier = `Fiche_Identification_${props.leopardNumber}.pdf`
    const cheminComplet = `${props.subfolderPath}/${nomFichier}`
    
    const result = await window.go.main.App.FileExists(cheminComplet)
    ficheExists.value = result
    
    console.log('📄 Vérification fiche:', { nomFichier, existe: result })
  } catch (err) {
    console.error('❌ Erreur vérification fiche:', err)
    ficheExists.value = false
  }
}

// Générer la fiche PDF
const creerFiche = async () => {
  if (!isDonneesCompletes.value) {
    alert('⚠️ Veuillez compléter les informations client obligatoires avant de créer la fiche.')
    return
  }

  isGenerating.value = true

  try {
    const doc = new jsPDF()
    let yPos = 20

    // Logo
    const img = new Image()
    img.src = logoPath
    await new Promise((resolve) => {
      img.onload = () => {
        doc.addImage(img, 'PNG', 15, 10, 30, 30)
        resolve()
      }
    })

    // Titre principal
    doc.setFontSize(18)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(0, 51, 102) // Bleu foncé
    doc.text('FICHE D\'IDENTIFICATION', 50, 20)

    doc.setFontSize(10)
    doc.setFont('helvetica', 'normal')
    doc.setTextColor(100, 100, 100)
    doc.text(`No. Dossier Leopard: ${props.leopardNumber}`, 50, 27)

    yPos = 50

    // Ligne de séparation
    doc.setDrawColor(0, 51, 102)
    doc.setLineWidth(0.5)
    doc.line(15, yPos, 195, yPos)
    yPos += 10

    // === SECTION 1: INFORMATIONS PERSONNELLES ===
    doc.setFillColor(240, 248, 255)
    doc.rect(15, yPos, 180, 8, 'F')
    doc.setFontSize(12)
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(0, 51, 102)
    doc.text('INFORMATIONS PERSONNELLES', 17, yPos + 5)
    yPos += 12

    doc.setFontSize(10)
    doc.setFont('helvetica', 'normal')
    doc.setTextColor(0, 0, 0)

    const infosPerso = [
      ['Nom:', props.client.nom || '-'],
      ['Prénom:', props.client.prenom || '-'],
      ['Date de naissance:', props.client.date_naissance || '-'],
      ['Sexe:', props.client.sexe || '-'],
      ['Lieu de naissance:', props.client.lieu_naissance || '-'],
      ['Statut marital:', props.client.statut_marital || '-'],
      ['Langue préférée:', props.client.langue_preferee || '-'],
      ['Première langue:', props.client.premiere_langue || '-']
    ]

    infosPerso.forEach(([label, value]) => {
      doc.setFont('helvetica', 'bold')
      doc.text(label, 20, yPos)
      doc.setFont('helvetica', 'normal')
      doc.text(value, 70, yPos)
      yPos += 6
    })

    yPos += 5

    // === SECTION 2: COORDONNÉES ===
    doc.setFillColor(240, 248, 255)
    doc.rect(15, yPos, 180, 8, 'F')
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(0, 51, 102)
    doc.text('COORDONNÉES', 17, yPos + 5)
    yPos += 12

    doc.setTextColor(0, 0, 0)

    const coordonnees = [
      ['Adresse:', props.client.adresse || '-'],
      ['Appartement:', props.client.appartement || '-'],
      ['Ville:', props.client.ville || '-'],
      ['Province:', props.client.province || 'Québec'],
      ['Code postal:', props.client.code_postal || '-'],
      ['Pays:', props.client.pays || 'Canada'],
      ['Téléphone:', props.client.telephone || '-'],
      ['Cellulaire:', props.client.cellulaire || '-'],
      ['Email:', props.client.email || '-']
    ]

    coordonnees.forEach(([label, value]) => {
      doc.setFont('helvetica', 'bold')
      doc.text(label, 20, yPos)
      doc.setFont('helvetica', 'normal')
      doc.text(value, 70, yPos)
      yPos += 6
    })

    yPos += 5

    // === SECTION 3: NUMÉROS D'IDENTIFICATION ===
    doc.setFillColor(240, 248, 255)
    doc.rect(15, yPos, 180, 8, 'F')
    doc.setFont('helvetica', 'bold')
    doc.setTextColor(0, 51, 102)
    doc.text('NUMÉROS D\'IDENTIFICATION', 17, yPos + 5)
    yPos += 12

    doc.setTextColor(0, 0, 0)

    const numeros = [
      ['No. Assurance Maladie (RAMQ):', props.client.numero_assurance_maladie || '-'],
      ['No. Sécurité Sociale:', props.client.numero_securite_sociale || '-'],
      ['No. HCM:', props.client.no_hcm || '-'],
      ['No. CHAUR:', props.client.no_chaur || '-']
    ]

    numeros.forEach(([label, value]) => {
      doc.setFont('helvetica', 'bold')
      doc.text(label, 20, yPos)
      doc.setFont('helvetica', 'normal')
      doc.text(value, 80, yPos)
      yPos += 6
    })

    yPos += 5

    // === SECTION 4: INFORMATIONS FAMILIALES (SI DISPONIBLES) ===
    if (props.client.nom_pere || props.client.nom_mere) {
      doc.setFillColor(240, 248, 255)
      doc.rect(15, yPos, 180, 8, 'F')
      doc.setFont('helvetica', 'bold')
      doc.setTextColor(0, 51, 102)
      doc.text('INFORMATIONS FAMILIALES', 17, yPos + 5)
      yPos += 12

      doc.setTextColor(0, 0, 0)

      if (props.client.nom_pere) {
        doc.setFont('helvetica', 'bold')
        doc.text('Père:', 20, yPos)
        doc.setFont('helvetica', 'normal')
        doc.text(props.client.nom_pere, 70, yPos)
        yPos += 6

        if (props.client.telephone_pere) {
          doc.setFont('helvetica', 'bold')
          doc.text('Tél. père:', 20, yPos)
          doc.setFont('helvetica', 'normal')
          doc.text(props.client.telephone_pere, 70, yPos)
          yPos += 6
        }
      }

      if (props.client.nom_mere) {
        doc.setFont('helvetica', 'bold')
        doc.text('Mère:', 20, yPos)
        doc.setFont('helvetica', 'normal')
        doc.text(props.client.nom_mere, 70, yPos)
        yPos += 6

        if (props.client.telephone_mere) {
          doc.setFont('helvetica', 'bold')
          doc.text('Tél. mère:', 20, yPos)
          doc.setFont('helvetica', 'normal')
          doc.text(props.client.telephone_mere, 70, yPos)
          yPos += 6
        }
      }

      yPos += 5
    }

    // === PIED DE PAGE ===
    doc.setFontSize(8)
    doc.setFont('helvetica', 'italic')
    doc.setTextColor(150, 150, 150)
    doc.text(`Date de création: ${new Date().toLocaleString('fr-CA')}`, 15, 280)
    doc.text(`Créé par: Utilisateur #${props.client.created_by || 'N/A'}`, 15, 285)

    // Sauvegarder
    const nomFichier = `Fiche_Identification_${props.leopardNumber}.pdf`
    const pdfBlob = doc.output('blob')
    
    // Convertir en base64 pour passer au backend
    const reader = new FileReader()
    reader.readAsDataURL(pdfBlob)
    
    reader.onloadend = async () => {
      const base64data = reader.result.split(',')[1]
      
      // Appeler le backend pour sauvegarder
      const result = await window.go.main.App.SavePDFToFolder(
        props.subfolderPath,
        nomFichier,
        base64data
      )

      if (result.success) {
        alert(`✅ Fiche d'identification créée avec succès!\n\nEmplacement:\n${result.path}`)
        ficheExists.value = true
        emit('fiche-created', result)
        emit('refresh')
      } else {
        alert(`❌ Erreur: ${result.error}`)
      }
    }

  } catch (err) {
    console.error('❌ Erreur génération PDF:', err)
    alert(`❌ Erreur: ${err.message}`)
  } finally {
    isGenerating.value = false
  }
}

// Ouvrir la fiche existante
const ouvrirFiche = async () => {
  try {
    const nomFichier = `Fiche_Identification_${props.leopardNumber}.pdf`
    const cheminComplet = `${props.subfolderPath}/${nomFichier}`
    
    const result = await window.go.main.App.OpenFile(cheminComplet)
    
    if (!result.success) {
      alert(`❌ Impossible d'ouvrir la fiche: ${result.error}`)
    }
  } catch (err) {
    console.error('❌ Erreur ouverture:', err)
    alert(`❌ Erreur: ${err.message}`)
  }
}

onMounted(() => {
  verifierExistence()
})
</script>

<style scoped>
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>