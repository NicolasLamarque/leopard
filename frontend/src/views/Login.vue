<template>

  <div class="login-page">
    <LoginForm @success="$emit('login', $event)" />
  </div>
  
</template>



<script>
import LoginForm from '../components/LoginForm.vue'

export default {
  components: { LoginForm },
  emits: ['login']
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #067ea3 0%, #86dfbe 100%);
}
</style>