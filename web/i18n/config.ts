import { messages } from "./index";

export default defineI18nConfig(() => ({
  messages,
  legacy: false,
  locale: "en",
  warnHtmlMessage: "off",
  warnHtmlInMessage: "off",
}));
