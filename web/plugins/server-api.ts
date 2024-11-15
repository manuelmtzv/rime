export default defineNuxtPlugin({
  name: "Rime Server Api Plugin",
  setup() {
    const serverApi = defineApi();

    return {
      provide: {
        serverApi,
      }
    }
  },
});

