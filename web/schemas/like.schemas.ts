import { z } from "zod";
import { likeEntities } from "@/types/like.type";

export const likeEntitySchema = z.object({
  entity: z.enum(likeEntities),
  id: z.string().uuid(),
});
