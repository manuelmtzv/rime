// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  modules: [
    "@nuxt/ui",
    "@nuxtjs/google-fonts",
    "@nuxt/icon",
    "@pinia/nuxt",
    "@nuxtjs/i18n",
  ],
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
  app: {
    pageTransition: { name: "page", mode: "out-in" },
    head: {
      title: "Rime app",
    },
  },
  runtimeConfig: {
    public: {
      serverUrl: process.env.SERVER_API,
    },
  },
  i18n: {
    locales: [
      {
        code: "en",
        iso: "en-US",
        name: "English",
      },
      {
        code: "es",
        iso: "es-ES",
        name: "Español",
      },
    ],
    defaultLocale: "en",
    detectBrowserLanguage: {
      useCookie: true,
      cookieKey: "i18n_redirected",
      redirectOn: "root",
    },
    vueI18n: "./i18n/config.ts",
  },
});
