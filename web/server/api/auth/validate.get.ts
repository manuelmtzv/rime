export default defineEventHandler(async (event) => {
  const cookies = parseCookies(event);

  const runtimeConfig = useRuntimeConfig();
  const response = await $fetch(
    `${runtimeConfig.public.serverUrl}/auth/validate`,
    {
      headers: {
        Authorization: `Bearer ${cookies["token"]}`,
      },
    }
  );

  return response;
});
