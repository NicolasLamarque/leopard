<template>
  <img :src="LogoLeopard" :width="size" :height="size" alt="Leopard Logo" />
</template>

<script setup>
    import LogoLeopard from '../assets/images/LeopardLogo.png'
defineProps({
    size: {
        type: Number,
        default: 48

    }
    
})
</script>