/** @type {import('tailwindcss').Config} */
export default {
  content: ["./**/*.templ"], // this is where our templates are located
  theme: {
    extend: {
      colors: {
        primary: colors.lime, 
      }
    },
  },
  plugins: [],
}
