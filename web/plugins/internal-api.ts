export default defineNuxtPlugin({
    name: "Rime Internal Api Plugin",
    setup() {
      const internalApi = defineApi(
        { internal: true }
      );
  
      return {
        provide: {
          internalApi,
        }
      }
    },
  });
  
  