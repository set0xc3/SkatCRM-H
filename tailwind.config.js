module.exports = {
  content: ["internal/frontend/templates/**/*.templ"],
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["light", "dark"],
  },
};
