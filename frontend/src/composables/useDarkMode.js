// composables/useDarkMode.js
import { ref } from 'vue'

// État global partagé entre tous les composants
const theme = ref('light')
const isDark = ref(false)

export function useDarkMode() {
  const applyTheme = (newTheme) => {
    if (newTheme === 'dark') {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }

  const setTheme = (newTheme) => {
    theme.value = newTheme
    isDark.value = newTheme === 'dark'
    applyTheme(newTheme)
  }

  return { 
    theme, 
    isDark, 
    setTheme
  }
}