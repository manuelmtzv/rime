export const useFormState = () => {
  const loading = ref(false);
  const error = ref<string | null>(null);
  const message = ref<string | null>(null);

  return {
    loading,
    error,
    message,
  };
};
