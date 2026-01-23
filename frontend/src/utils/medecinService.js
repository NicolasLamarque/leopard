// frontend/src/services/medecinService.js
import { GetMedecins, GetMedecinByID, CreateMedecin, UpdateMedecin, DeleteMedecin, SearchMedecins, GetMedecinClients, GetMedecinClientsCount } from '../../wailsjs/go/main/App'

export default {
  // Récupérer tous les médecins
  async getAll() {
    try {
      return await GetMedecins()
    } catch (error) {
      console.error('Erreur récupération médecins:', error)
      throw error
    }
  },

  // Récupérer un médecin par ID
  async getById(id) {
    try {
      return await GetMedecinByID(id)
    } catch (error) {
      console.error('Erreur récupération médecin:', error)
      throw error
    }
  },

  // Créer un médecin
  async create(medecin) {
    try {
      return await CreateMedecin(medecin)
    } catch (error) {
      console.error('Erreur création médecin:', error)
      throw error
    }
  },

  // Mettre à jour un médecin
  async update(medecin) {
    try {
      return await UpdateMedecin(medecin)
    } catch (error) {
      console.error('Erreur mise à jour médecin:', error)
      throw error
    }
  },

  // Supprimer un médecin
  async delete(id) {
    try {
      return await DeleteMedecin(id)
    } catch (error) {
      console.error('Erreur suppression médecin:', error)
      throw error
    }
  },

  // Rechercher des médecins
  async search(query) {
    try {
      return await SearchMedecins(query)
    } catch (error) {
      console.error('Erreur recherche médecins:', error)
      throw error
    }
  },

  // 🆕 Récupérer les clients d'un médecin
  async getClients(licence) {
    try {
      return await GetMedecinClients(licence)
    } catch (error) {
      console.error('Erreur récupération clients du médecin:', error)
      throw error
    }
  },

  // 🆕 Récupérer le nombre de clients d'un médecin
  async getClientsCount(licence) {
    try {
      return await GetMedecinClientsCount(licence)
    } catch (error) {
      console.error('Erreur comptage clients:', error)
      throw error
    }
  }
}