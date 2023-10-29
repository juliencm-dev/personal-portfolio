/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.tmpl"], 
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

