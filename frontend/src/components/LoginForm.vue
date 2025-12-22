<template>
<div class="form-container">

  <form @submit.prevent="login" class="form">
  

   <h1 class="h1-title">Connexion</h1>  
   <div class="overlay" style="text-align: center; margin-bottom: 1rem;">
    <LeopardLogo size="400" class="brand-icon" /> 
  </div>
    <input v-model="username" placeholder="Nom d'utilisateur" required />
    <input v-model="password" type="password" placeholder="Mot de passe" required />
    <button type="submit">Se connecter</button>
    <p v-if="error" class="error">{{ error }}</p>
  </form>
  
</div>
</template>

<script>
import { ref } from 'vue';
import { Login } from '../../wailsjs/go/main/App';
import LeopardLogo from './LeopardLogo.vue'; 
import { useRouter } from 'vue-router'; // 1. IMPORTER useRouter

export default {
  components: {
    LeopardLogo
  },
  emits: ['success'],
  setup(props, { emit }) {
    const router = useRouter(); //  1. INITIALISER useRouter
    const username = ref('')
    const password = ref('')
    const error = ref('')
    
    const login = async () => {
      try {
        const user = await Login(username.value, password.value)
        
        // 2. IMPORTANT : On enregistre le succès
        emit('success', user)
        
        //  On change de page vers le dashboard
        // Selon ton router/index.ts, le chemin est '/app/clients'
        router.push('/app/clients') 
        
      } catch (e) {
        error.value = 'Identifiants invalides'
      }
    }
    
    return { username, password, error, login }
  }
}
</script>

<style scoped>
.form {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 400px;
}
input {
  width: 100%;
  padding: 0.75rem;
  margin: 0.5rem 0;
  border: 1px solid #ddd;
  border-radius: 4px;
}
button {
  width: 100%;
  padding: 0.75rem;
  background: #667eea;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.error { color: red; }

.h1-title {
  text-align: center;
  margin-bottom: 1rem;
  
  color: #0e6b6b;
  font-size: 30px;
}
</style>