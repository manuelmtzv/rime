export type AuthAction = () => void | Promise<void>;

export const useAuthAction = () => {
  const { user } = useUserState();
  const { setShowAuthRequiredModal } = useAppState();

  async function requireAuthAndRun(callback?: AuthAction) {
    if (!user.value) {
      setShowAuthRequiredModal(true);
      return;
    }

    if (callback) {
      await callback();
    }
  }

  return { requireAuthAndRun };
};
