export const useAppStore = defineStore("app", () => {
  const showAuthRequiredModal = ref(false);

  const setShowAuthRequiredModal = (value: boolean) => {
    showAuthRequiredModal.value = value;
  };

  return {
    showAuthRequiredModal,
    setShowAuthRequiredModal,
  };
});
