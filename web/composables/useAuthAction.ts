export type AuthAction = () => void | Promise<void>;

export const useAuthAction = () => {
  const { user } = useUserState();
  const { setShowAuthRequiredModal } = useAppState();

  async function run(callback: AuthAction) {
    if (!user.value) {
      setShowAuthRequiredModal(true);
      return;
    }
    await callback();
  }

  return { run };
};
