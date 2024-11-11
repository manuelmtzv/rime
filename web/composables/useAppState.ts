export const useAppState = () => {
  const store = useAppStore();
  const state = storeToRefs(store);
  const { setShowAuthRequiredModal } = store;

  return {
    ...state,
    setShowAuthRequiredModal,
  };
};
