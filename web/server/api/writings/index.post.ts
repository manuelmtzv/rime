import { writingSchema } from "@/schemas";
import { mapH3Error } from "@/utils/errors";
import { DataResponse, Writing } from "@/types";

export default defineEventHandler(async (event) => {
  const serverUrl = useRuntimeConfig().public.serverUrl;

  const validBody = await readValidatedBody(event, (body) =>
    writingSchema.safeParse(body)
  );

  if (!validBody.success) {
    console.log(validBody.error.issues);
    throw validBody.error.issues;
  }

  try {
    const response = await event.$fetch<DataResponse<Writing>>(
      `${serverUrl}/writings`,
      {
        method: "POST",
        body: JSON.stringify(validBody.data),
      }
    );

    return response;
  } catch (err) {
    sendError(event, mapH3Error(err));
  }
});
