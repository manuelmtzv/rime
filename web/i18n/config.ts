import { messages } from "./index";

export default defineI18nConfig(() => ({
  legacy: false,
  locale: "en",
  messages,
  warnHtmlMessage: "off",
  warnHtmlInMessage: "off",
}));
