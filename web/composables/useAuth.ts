import type {
  AuthResponse,
  LoginRequest,
  RegisterRequest,
} from "@/types/auth.type";

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

  async function register(registerForm: RegisterRequest) {
    const response = await fetch<AuthResponse>(
      `${runtimeConfig.public.API_BASE_URL}/auth/register`,
      {
        method: "POST",
        body: JSON.stringify(registerForm),
      }
    );

    setUser(response.data.user);

    return response;
  }

  return { login };
};
