export function cleanTiptapOutput(content: string): string {
    const result = content
      .replace(/(<p><\/p>)+/g, "<br />")
      .replace(/(<p><br><\/p>)+/g, "<br />")
      .replace(/(<br \/>)+/g, "<br />")
      .replace(/(<br>)+/g, "<br><br>");
  
    return result;
  }