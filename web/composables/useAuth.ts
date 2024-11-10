import type {
  AuthResponse,
  LoginRequest,
  RegisterRequest,
} from "@/types/auth.type";

export const useAuth = () => {
  const fetch = useNuxtApp().$api;
  const { setUser } = useUserState();

  async function login(loginForm: LoginRequest) {
    const response = await $fetch<AuthResponse>("/api/auth/login", {
      method: "POST",
      body: JSON.stringify(loginForm),
    });

    setUser(response.data.user);

    return response;
  }

  async function register(registerForm: RegisterRequest) {
    const response = await fetch<AuthResponse>("/auth/register", {
      method: "POST",
      body: JSON.stringify(registerForm),
    });

    setUser(response.data.user);

    return response;
  }

  async function validate() {
    const requestFetch = useRequestFetch();

    return requestFetch<AuthResponse>("/api/auth/validate");
  }

  return { login, register, validate };
};
