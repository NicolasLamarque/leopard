// frontend/src/stores/ClientStore.js

import { defineStore } from 'pinia'
import { GetClientByID } from '../../wailsjs/go/main/App'

export const useClientStore = defineStore('client', {
  state: () => ({
    currentClient: null,
    loading: false,
    error: null // Ajout d'un état d'erreur
  }),
  actions: {
    async selectClient(id) {
      // Sécurité : s'assurer que l'ID est bien un nombre entier
      const clientId = typeof id === 'string' ? parseInt(id, 10) : id
      
      this.loading = true
      this.error = null
      this.currentClient = null // On vide immédiatement avant de charger

      try {
        console.log(`📡 Appel Go: GetClientByID(${clientId})`)
        const result = await GetClientByID(clientId)
        
        if (result) {
          this.currentClient = result
          console.log("✅ Client chargé dans le store:", result.nom)
        } else {
          this.error = "Client non trouvé"
        }
      } catch (err) {
        this.error = "Erreur de connexion à la base de données"
        console.error("❌ Erreur store:", err)
      } finally {
        this.loading = false
      }
    },
    
    clearClient() {
      this.currentClient = null
      this.loading = false
      this.error = null
    }
  }
})