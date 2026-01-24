<template>
<div class="form-container">
  <form @submit.prevent="isSetupMode ? createFirstUser() : login()" class="form">
    <h1 class="h1-title">{{ isSetupMode ? 'Premier utilisateur' : 'Connexion' }}</h1>
    
    <div class="overlay" style="text-align: center; margin-bottom: 1rem;">
      <LeopardLogo size="400" class="brand-icon" /> 
    </div>

    <!-- Message de bienvenue en mode Setup -->
    <div v-if="isSetupMode" class="setup-message">
      <p>👋 Bienvenue ! Créez le premier compte administrateur.</p>
    </div>

    <!-- Champs communs -->
    <input 
      v-model="username" 
      placeholder="Nom d'utilisateur" 
      required 
      :disabled="loading"
    />
    <input 
      v-model="password" 
      type="password" 
      placeholder="Mot de passe" 
      required 
      :disabled="loading"
    />

    <!-- Champs supplémentaires en mode Setup -->
    <template v-if="isSetupMode">
      <input 
        v-model="confirmPassword" 
        type="password" 
        placeholder="Confirmer le mot de passe" 
        required 
        :disabled="loading"
      />
      <input 
        v-model="fullName" 
        placeholder="Nom complet" 
        required 
        :disabled="loading"
      />
    </template>

    <!-- Bouton -->
    <button type="submit" :disabled="loading">
      {{ loading ? '⏳ Chargement...' : (isSetupMode ? 'Créer le compte' : 'Se connecter') }}
    </button>

    <!-- Erreur -->
    <p v-if="error" class="error">{{ error }}</p>

    <!-- Toggle manuel (optionnel, si tu veux forcer le mode) -->
    <button 
      v-if="!loading && userCount > 0" 
      type="button" 
      @click="toggleMode" 
      class="toggle-btn"
    >
      {{ isSetupMode ? '← Retour à la connexion' : 'Créer un nouveau compte' }}
    </button>
  </form>
</div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { Login } from '../../wailsjs/go/main/App';
import LeopardLogo from './LeopardLogo.vue'; 
import { useRouter } from 'vue-router';

import { GetUserCount } from '../../wailsjs/go/main/App';


export default {
  components: { LeopardLogo },
  emits: ['success'],
  setup(props, { emit }) {
    const router = useRouter();
    const username = ref('');
    const password = ref('');
    const confirmPassword = ref('');
    const fullName = ref('');
    const error = ref('');
    const loading = ref(false);
    const isSetupMode = ref(false);
    const userCount = ref(0);

    // Vérifier s'il existe déjà des utilisateurs
    const checkUserCount = async () => {
      try {
        // Appel à une fonction Go qui vérifie si des users existent
        const count = await window.go.main.App.GetUserCount();
        userCount.value = count;
        
        // Si aucun user, activer le mode Setup automatiquement
        if (count === 0) {
          isSetupMode.value = true;
        }
      } catch (e) {
        console.error('Erreur vérification users:', e);
        // En cas d'erreur, on suppose qu'il faut créer le premier user
        isSetupMode.value = true;
      }
    };

    // Login classique
    const login = async () => {
      loading.value = true;
      error.value = '';
      
      try {
        const user = await Login(username.value, password.value);
        emit('success', user);
        router.push({ name: 'home' });
      } catch (e) {
        error.value = 'Identifiants invalides';
      } finally {
        loading.value = false;
      }
    };

    // Créer le premier utilisateur
    const createFirstUser = async () => {
      error.value = '';

      // Validations
      if (password.value !== confirmPassword.value) {
        error.value = 'Les mots de passe ne correspondent pas';
        return;
      }

      if (password.value.length < 5) {
        error.value = 'Le mot de passe doit contenir au moins 5 caractères';
        return;
      }

      if (!fullName.value.trim()) {
        error.value = 'Le nom complet est requis';
        return;
      }

      loading.value = true;

      try {
        // Appel à la fonction Go pour créer le premier user (admin)
        await window.go.main.App.CreateFirstUser({
          username: username.value,
          password: password.value,
          fullName: fullName.value,
          role: 'admin'
        });

        // Connexion automatique après création
        const user = await Login(username.value, password.value);
        emit('success', user);
        router.push({ name: 'home' });
      } catch (e) {
        error.value = e.message || 'Erreur lors de la création du compte';
      } finally {
        loading.value = false;
      }
    };

    // Toggle entre login et setup
    const toggleMode = () => {
      isSetupMode.value = !isSetupMode.value;
      error.value = '';
      password.value = '';
      confirmPassword.value = '';
    };

    // Vérifier au montage
    onMounted(() => {
      checkUserCount();
    });

    return { 
      username, 
      password, 
      confirmPassword,
      fullName,
      error, 
      loading,
      isSetupMode,
      userCount,
      login,
      createFirstUser,
      toggleMode
    };
  }
}
</script>

<style scoped>
.form-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.form {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  width: 450px;
  box-shadow: 0 10px 40px rgba(0,0,0,0.1);
}

.h1-title {
  text-align: center;
  margin-bottom: 0.5rem;
  color: #0e6b6b;
  font-size: 30px;
}

.setup-message {
  background: #e6f7ff;
  border: 1px solid #91d5ff;
  border-radius: 6px;
  padding: 0.75rem;
  margin-bottom: 1rem;
  text-align: center;
}

.setup-message p {
  margin: 0;
  color: #0050b3;
  font-size: 0.9rem;
}

input {
  width: 100%;
  padding: 0.75rem;
  margin: 0.5rem 0;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s;
  box-sizing: border-box;
}

input:focus {
  outline: none;
  border-color: #667eea;
}

input:disabled {
  background: #f5f5f5;
  cursor: not-allowed;
}

button {
  width: 100%;
  padding: 0.75rem;
  background: #667eea;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: background 0.3s;
  margin-top: 0.5rem;
}

button:hover:not(:disabled) {
  background: #5568d3;
}

button:disabled {
  background: #a0aec0;
  cursor: not-allowed;
}

.toggle-btn {
  background: transparent;
  color: #667eea;
  border: 1px solid #667eea;
  margin-top: 1rem;
  font-weight: normal;
}

.toggle-btn:hover {
  background: #f7fafc;
}

.error { 
  color: #e53e3e;
  text-align: center;
  margin-top: 1rem;
  font-size: 0.9rem;
}
</style>