import { createI18nMessage } from "@vuelidate/validators";
import * as validators from "@vuelidate/validators";

export const useValidation = () => {
  const i18n = useI18n();
  const withI18nMessage = createI18nMessage({ t: i18n.t.bind(i18n) });

  const required = withI18nMessage(validators.required);
  const email = withI18nMessage(validators.email);
  const minLength = withI18nMessage(validators.minLength);
  const maxLength = withI18nMessage(validators.maxLength);

  return {
    required,
    email,
    minLength,
    maxLength,
  };
};
