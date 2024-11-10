const whitelist = ["/auth/login", "/auth/register"];

export default defineNuxtRouteMiddleware(async (to, from) => {
  const localePath = useLocalePath();

  const langPattern = /^\/[a-z]{2}(\/|$)/;
  let path = to.path;

  if (langPattern.test(path)) {
    path = path.replace(langPattern, "/");
  }

  if (whitelist.includes(path)) {
    return;
  }

  const { setUser } = useUserState();
  const { validate } = useAuth();
  const toast = useToast();

  try {
    const response = await validate();

    if (response.data.user) {
      setUser(response?.data.user);
    }
  } catch (validationError) {
    console.error(validationError);

    // try {
    //   await refresh();
    // } catch (refreshError) {
    //   toast.warning("Your session has expired. Please log in again.");
    //   return navigateTo(localePath("/auth/login"));
    // }
  }
});
