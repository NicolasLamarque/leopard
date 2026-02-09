<template>
  <div class="glass-form">
    <div class="header-text">
      <h2>{{ isSetupMode ? 'Initialisation Leopard' : 'Connexion' }}</h2>
      <p v-if="isSetupMode" class="setup-hint">👋 Créez le premier compte administrateur (Loi 25)</p>
    </div>

    <form @submit.prevent="isSetupMode ? createFirstUser() : login()" class="form-grid">
      
      <div class="input-field full">
        <input v-model="username" placeholder="Nom d'utilisateur" required :disabled="loading" />
      </div>
      <div class="input-field full">
        <input v-model="password" type="password" placeholder="Mot de passe" required :disabled="loading" />
      </div>

      <template v-if="isSetupMode">
        <div class="input-field full">
          <input v-model="confirmPassword" type="password" placeholder="Confirmer le mot de passe" required />
        </div>
        <div class="input-field full">
          <input v-model="fullName" placeholder="Nom complet" required />
        </div>
        
        <div class="input-field half">
          <input v-model="email" type="email" placeholder="Adresse e-mail" />
        </div>
        <div class="input-field half">
          <input v-model="phone" type="tel" placeholder="Téléphone" />
        </div>

        <div class="input-field half">
          <input v-model="cellPhone" type="tel" placeholder="Cellulaire" />
        </div>
        <div class="input-field half">
          <input v-model="address" placeholder="Adresse" />
        </div>

        <div class="input-field third">
          <input v-model="city" placeholder="Ville" />
        </div>
        <div class="input-field third">
          <input v-model="province" placeholder="Province" />
        </div>
        <div class="input-field third">
          <input v-model="postalCode" placeholder="Code postal" />
        </div>
        
        <div class="input-field full">
          <input v-model="country" placeholder="Pays" />
        </div>
      </template>

      <button type="submit" class="submit-btn" :disabled="loading">
        {{ loading ? '⏳ Chargement...' : (isSetupMode ? 'Lancer le système' : 'Entrer') }}
      </button>

      <p v-if="error" class="error-msg">{{ error }}</p>

      <button v-if="!loading && userCount > 0" type="button" @click="toggleMode" class="toggle-link">
        {{ isSetupMode ? '← Retour à la connexion' : 'Créer un compte' }}
      </button>
    </form>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { Login } from '../../wailsjs/go/main/App';
import { useRouter } from 'vue-router';

export default {
  emits: ['success'],
  setup(props, { emit }) {
    const router = useRouter();
    const username = ref('');
    const password = ref('');
    const confirmPassword = ref('');
    const fullName = ref('');
    // Tes variables setup
    const email = ref('');
    const phone = ref('');
    const cellPhone = ref('');
    const address = ref('');
    const city = ref('');
    const province = ref('');
    const postalCode = ref('');
    const country = ref('');

    const error = ref('');
    const loading = ref(false);
    const isSetupMode = ref(false);
    const userCount = ref(0);

    const checkUserCount = async () => {
      try {
        const count = await window.go.main.App.GetUserCount();
        userCount.value = count;
        if (count === 0) isSetupMode.value = true;
      } catch (e) { isSetupMode.value = true; }
    };

    const login = async () => {
      loading.value = true;
      error.value = '';
      try {
        const user = await Login(username.value, password.value);
        emit('success', user);
        router.push({ name: 'home' });
      } catch (e) { error.value = 'Identifiants invalides'; }
      finally { loading.value = false; }
    };

    const createFirstUser = async () => {
      error.value = '';
      if (password.value !== confirmPassword.value) { error.value = 'Mots de passe différents'; return; }
      loading.value = true;
      try {
        await window.go.main.App.CreateFirstUser({
          username: username.value,
          password: password.value,
          fullName: fullName.value,
          role: 'admin',
          email: email.value,
          telephone: phone.value,
          cellulaire: cellPhone.value,
          adresse: address.value,
          ville: city.value,
          province: province.value,
          code_postal: postalCode.value,
          pays: country.value
        });
        const user = await Login(username.value, password.value);
        emit('success', user);
        router.push({ name: 'home' });
      } catch (e) { error.value = e.message || 'Erreur création'; }
      finally { loading.value = false; }
    };

    const toggleMode = () => {
      isSetupMode.value = !isSetupMode.value;
      error.value = '';
    };

    onMounted(checkUserCount);

    return { 
      username, password, confirmPassword, fullName, email, phone, cellPhone, 
      address, city, province, postalCode, country,
      error, loading, isSetupMode, userCount, login, createFirstUser, toggleMode 
    };
  }
}
</script>

<style scoped>
.glass-form {
  background: rgba(30, 41, 59, 0.5);
  backdrop-filter: blur(15px);
  padding: 3rem;
  border-radius: 28px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
  width: 100%;
}

.header-text h2 { color: #f8fafc; font-weight: 300; font-size: 1.8rem; margin-bottom: 0.5rem; }
.setup-hint { color: #38bdf8; font-size: 0.9rem; margin-bottom: 2rem; }

.form-grid { display: flex; flex-wrap: wrap; gap: 12px; }

.input-field { display: flex; flex-direction: column; }
.full { width: 100%; }
.half { width: calc(50% - 6px); }
.third { width: calc(33.33% - 8px); }

input {
  background: rgba(15, 23, 42, 0.7);
  border: 1px solid rgba(255, 255, 255, 0.1);
  padding: 0.9rem;
  border-radius: 10px;
  color: white;
  width: 100%;
}

input:focus { border-color: #38bdf8; outline: none; background: rgba(15, 23, 42, 0.9); }

.submit-btn {
  width: 100%;
  background: #0ea5e9;
  color: white;
  padding: 1rem;
  border-radius: 12px;
  border: none;
  font-weight: 600;
  cursor: pointer;
  margin-top: 1rem;
}

.toggle-link {
  background: none; border: none; color: #94a3b8; width: 100%; 
  margin-top: 1rem; cursor: pointer; font-size: 0.8rem;
}

.error-msg { color: #fb7185; width: 100%; text-align: center; margin-top: 1rem; }
</style>