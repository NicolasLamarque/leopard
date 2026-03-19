// frontend/src/composables/useRefListes.js
import { ref } from 'vue'
import {
  GetRefListeByCategorie,
  GetAllRefCategories,
  CreateRefListe,
  UpdateRefListe,
  DeleteRefListe,
  UpdateRefListeOrdre,
  OuvrirDialogCSVRefListes
} from '../../wailsjs/go/main/App'

// ─── État global partagé (chargé une seule fois pour toute l'app) ───
const cache = ref({})         // { type_intervenant: [...], statut_dossier: [...] }
const categories = ref([])
const loading = ref(false)
const error = ref(null)

export function useRefListes() {

  // ─── Charger une catégorie (avec cache) ──────────────────────────
  const getCategorie = async (categorie) => {
    if (cache.value[categorie]) return cache.value[categorie]
    try {
      const items = await GetRefListeByCategorie(categorie)
      cache.value[categorie] = items || []
      return cache.value[categorie]
    } catch (err) {
      console.error(`Erreur chargement liste '${categorie}':`, err)
      return []
    }
  }

  // ─── Charger toutes les catégories (page Settings) ───────────────
  const loadCategories = async () => {
    loading.value = true
    error.value = null
    try {
      categories.value = await GetAllRefCategories() || []
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  // ─── Vider le cache d'une catégorie (après ajout/modif/suppression) ─
  const invaliderCache = (categorie = null) => {
    if (categorie) {
      delete cache.value[categorie]
    } else {
      cache.value = {} // Vider tout
    }
  }

  // ─── CRUD ────────────────────────────────────────────────────────
  const creerEntree = async (req) => {
    loading.value = true
    try {
      await CreateRefListe(req)
      invaliderCache(req.categorie)
    } catch (err) {
      error.value = err.message
      throw err
    } finally {
      loading.value = false
    }
  }

  const modifierEntree = async (id, req) => {
    loading.value = true
    try {
      await UpdateRefListe(id, req)
      invaliderCache(req.categorie)
    } catch (err) {
      error.value = err.message
      throw err
    } finally {
      loading.value = false
    }
  }

  const supprimerEntree = async (id, categorie) => {
    loading.value = true
    try {
      await DeleteRefListe(id)
      invaliderCache(categorie)
    } catch (err) {
      error.value = err.message
      throw err
    } finally {
      loading.value = false
    }
  }

  const mettreAJourOrdre = async (id, ordre) => {
    try {
      await UpdateRefListeOrdre(id, ordre)
    } catch (err) {
      console.error('Erreur mise à jour ordre:', err)
    }
  }

  // ─── Import CSV ──────────────────────────────────────────────────
  const importerCSV = async () => {
    loading.value = true
    error.value = null
    try {
      const rapport = await OuvrirDialogCSVRefListes()
      invaliderCache() // Vider tout le cache après import
      return rapport
    } catch (err) {
      error.value = err.message
      throw err
    } finally {
      loading.value = false
    }
  }

  // ─── Helpers pour les dropdowns FormKit ──────────────────────────
  // Retourne les options au format { label, value } pour FormKit
  const optionsFormKit = async (categorie) => {
  const items = await getCategorie(categorie)
  console.log('ITEM[0]:', JSON.stringify(items[0])) // ← on regarde la structure
  return items.map(item => ({
    label: item.Libelle || item.libelle,
    value: item.Libelle || item.libelle
  }))
}
  // ─── Helper badge couleur (remplace getTypeColor codé en dur) ────
  const badgeClasses = (couleur) => {
    const c = couleur || 'slate'
    return `bg-${c}-100 dark:bg-${c}-900/30 text-${c}-700 dark:text-${c}-400 border border-${c}-200 dark:border-${c}-800`
  }

  const badgeDot = (couleur) => {
    const c = couleur || 'slate'
    return `bg-${c}-400`
  }

  return {
    // État
    cache,
    categories,
    loading,
    error,
    // Lecture
    getCategorie,
    loadCategories,
    optionsFormKit,
    // CRUD
    creerEntree,
    modifierEntree,
    supprimerEntree,
    mettreAJourOrdre,
    // Import
    importerCSV,
    // Helpers UI
    badgeClasses,
    badgeDot
  }
}