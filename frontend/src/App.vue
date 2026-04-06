<<template>
  <router-view />
   <Toast /> 
  
</template>

<script setup lang="ts">
// On ne gère plus le "v-if" ici !
// Détection automatique du thème système

import Toast from './components/shared/ToastContainer.vue'


</script>