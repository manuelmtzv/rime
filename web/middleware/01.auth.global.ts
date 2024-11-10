import { LANG_PATTERN, WHITELIST, PROTECTED_ROUTES } from "@/constants";

export default defineNuxtRouteMiddleware(async (to, from) => {
  if (import.meta.client || !to.name) return;

  const localePath = useLocalePath();

  let path = to.path;
  if (LANG_PATTERN.test(path)) {
    path = path.replace(LANG_PATTERN, "/");
  }

  if (WHITELIST.includes(path)) return;

  const { user, setUser } = useUserState();
  const { validate } = useAuth();
  const toast = useToast();

  try {
    const { data } = await validate();

    if (data.user) {
      setUser(data.user);
    }
  } catch (error) {
    if (PROTECTED_ROUTES.includes(path)) {
      console.error("Validation error:", error);
      return navigateTo(localePath("/auth/login"));
    }
  }
});
