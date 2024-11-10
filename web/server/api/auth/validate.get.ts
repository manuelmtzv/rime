export default defineEventHandler(async (event) => {
  const cookie = getCookie(event, "token");

  console.log("cookie", cookie);

  const runtimeConfig = useRuntimeConfig();
  const response = await $fetch(
    `${runtimeConfig.public.serverUrl}/auth/validate`,
    {
      headers: {
        Authorization: `Bearer ${cookie}`,
      },
    }
  );

  return response;
});
