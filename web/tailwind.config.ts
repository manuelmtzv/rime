/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      colors: {
        grey: {
          50: "#f9f9f9",
          100: "#f0f0f0",
          200: "#e0e0e0",
          300: "#d0d0d0",
          400: "#b0b0b0",
          500: "#909090",
          600: "#707070",
          700: "#505050",
          800: "#303030",
          900: "#101010",
        },
      },
    },
    fontFamily: {
      general: ["Work Sans", "sans-serif"],
      poetry: ["Merriweather", "serif"],
    },
  },
  plugins: [],
};
