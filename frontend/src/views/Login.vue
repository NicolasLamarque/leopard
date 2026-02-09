<template>
  <div class="login-wrapper">
    <div class="split-layout">
      <div class="brand-side">
        <LeopardLogo size="425" />
        <div class="tagline">
  
        </div>
      </div>

      <div class="form-side">
        <LoginForm @success="$emit('login', $event)" />
      </div>
    </div>
  </div>
</template>

<script>
import LoginForm from '../components/LoginForm.vue'
import LeopardLogo from '../components/LeopardLogo.vue'

export default {
  components: { LoginForm, LeopardLogo },
  emits: ['login']
  
}
</script>

<style scoped>
.login-wrapper {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(circle at center, #0a0c10 100%, #0a0c10 100%);
  color: #e2e8f0;
  overflow: hidden;
}

.split-layout {
  display: flex;
  align-items: center;
  width: 95%;
  max-width: 1300px;
  gap: 2rem;
}

.brand-side {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.tagline {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-top: -1rem;
    opacity: 0.5;
}

.line { width: 30px; height: 1px; background: #386f86; }

.form-side {
  flex: 1;
  display: flex;
  justify-content: center;
  max-height: 90vh;
  overflow-y: auto; /* Pour scroller si le formulaire setup est long */
  padding: 20px;
}

@media (max-width: 1000px) {
  .split-layout { flex-direction: column; overflow-y: auto; }
  .brand-side { margin-bottom: 2rem; }
}
</style>