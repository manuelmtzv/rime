import type { User } from "@/types";

export const useUserStore = defineStore("user", () => {
  const user = ref<User | null>(null);

  const fullName = computed(() => {
    return user.value ? `${user.value.name} ${user.value.lastname}` : "";
  });

  const setUser = (newUser: User) => {
    user.value = newUser;
  };

  const logout = () => {
    user.value = null;
  };

  return {
    user,
    fullName,
    setUser,
    logout,
  };
});
