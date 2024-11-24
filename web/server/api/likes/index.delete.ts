import { likeEntitySchema } from "@/schemas/like.schemas";
import { mapH3Error } from "@/utils/get-error";
import { DataResponse, WritingLike } from "@/types";

export default defineEventHandler(async (event) => {
  const serverUrl = useRuntimeConfig().public.serverUrl;

  const validBody = await readValidatedBody(event, (body) =>
    likeEntitySchema.safeParse(body)
  );

  if (!validBody.success) {
    throw validBody.error.issues;
  }

  try {
    const response = await event.$fetch<DataResponse<WritingLike>>(
      `${serverUrl}/likes/${validBody.data.entity}/${validBody.data.id}`,
      {
        method: "DELETE",
      }
    );

    return response;
  } catch (err) {
    sendError(event, mapH3Error(err));
  }
});
