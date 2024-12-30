import { z } from "zod";

export const writingSchema = z.object({
  type: z.string(),
  title: z.string(),
  content: z.string(),
});
