import type { AuthResponse, LoginRequest } from "@/types/auth.type";

export const useAuth = () => {
  const fetch = useNuxtApp().$api;
  const runtimeConfig = useRuntimeConfig();
  const { setUser } = useUserState();

  async function login(loginForm: LoginRequest) {
    const response = await fetch<AuthResponse>(
      `${runtimeConfig.public.API_BASE_URL}/auth/login`,
      {
        method: "POST",
        body: JSON.stringify(loginForm),
      }
    );

    setUser(response.data.user);

    return response;
  }

  return { login };
};
