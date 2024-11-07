export function getHtmlText(value: string): string {
  return value.replace(/<[^>]+>/g, "");
}
