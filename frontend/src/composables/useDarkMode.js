import { ref, watch } from 'vue'

export function useDarkMode() {
  // Initialiser depuis localStorage ou préférence système
  const getInitialTheme = () => {
    const stored = localStorage.getItem('theme')
    if (stored) return stored
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }

  const theme = ref(getInitialTheme())
  const isDark = ref(theme.value === 'dark')

  const setTheme = (newTheme) => {
    theme.value = newTheme
    isDark.value = newTheme === 'dark'
  }

  const toggle = () => {
    setTheme(isDark.value ? 'light' : 'dark')
  }

  // Appliquer le thème au DOM
  watch(theme, (newTheme) => {
    if (newTheme === 'dark') {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
    localStorage.setItem('theme', newTheme)
  }, { immediate: true })

  return { theme, isDark, setTheme, toggle }
}