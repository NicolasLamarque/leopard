// tailwind.config.js
/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class', // Mode sombre basé sur la classe 'dark'
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class', // Active le mode sombre basé sur la classe 'dark'
  theme: {
    extend: {},
  },
  plugins: [],
}