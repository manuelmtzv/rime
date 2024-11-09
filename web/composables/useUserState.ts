export const useUserState = () => {
  const store = useUserStore();
  const state = storeToRefs(store);
  const { setUser, logout } = store;

  return {
    ...state,
    setUser,
    logout,
  };
};
