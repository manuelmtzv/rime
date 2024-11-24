import { loginSchema } from "@/schemas/auth.schemas";
import { AuthResponse } from "@/types";
import { mapH3Error } from "@/utils/get-error";

export default defineEventHandler(async (event) => {
  const body = await readValidatedBody(event, loginSchema.parse);
  const runtimeConfig = useRuntimeConfig();

  try {
    const response = await event.$fetch<AuthResponse>(
      `${runtimeConfig.public.serverUrl}/auth/login`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      }
    );

    setCookie(event, "token", response.data.token, {
      maxAge: 60 * 60, // TODO: change this for a value from the server
      httpOnly: true,
      sameSite: "strict",
    });

    setCookie(event, "refresh-token", response.data.user.id, {
      maxAge: 60 * 60 * 24 * 14, // TODO: change this for a value from the server
      httpOnly: true,
      sameSite: "strict",
    });

    return response;
  } catch (err) {
    sendError(event, mapH3Error(err));
  }
});
