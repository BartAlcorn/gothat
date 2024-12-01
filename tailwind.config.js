/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pkg/**/*.templ"],
  theme: {
    extend: {},
  },
  darkMode: "media",
  plugins: [
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography"),
    require("tailwindcss-view-transitions"),
    require("daisyui"),
  ],
  daisyui: {
    themes: ["light", "dark", "cupcake"],
  },
};
