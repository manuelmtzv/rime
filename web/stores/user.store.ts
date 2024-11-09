import type { User } from "@/types";

export const useUserStore = defineStore("user", () => {
  const user = ref<User | null>(null);

  const setUser = (newUser: User) => {
    user.value = newUser;
  };

  const logout = () => {
    user.value = null;
  };

  return {
    user,
    setUser,
    logout,
  };
});
