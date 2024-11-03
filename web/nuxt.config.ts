// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxtjs/google-fonts", "@nuxt/icon"],
  googleFonts: {
    families: {
      Merriweather: true,
      "Work Sans": true,
    },
  },
  colorMode: {
    preference: "light",
    fallback: "light",
  },
  icon: {
    size: "24",
  },
});
