import { H3Error } from "h3";
import type { H3ErrorPayload } from "@/types/api.type";

type config = {
  defaultErrorMessage?: string;
  internal?: boolean;
};

export function getError(
  error: any,
  { internal = false, defaultErrorMessage = "An error occurred" }: config = {}
) {
  if (isH3Error(error)) {
    return (
      (error as H3Error<H3ErrorPayload>).data?.error || defaultErrorMessage
    );
  }

  if (error instanceof Error) {
    return error.message;
  }

  if (typeof error === "string") {
    return error;
  }

  return defaultErrorMessage;
}

export function isH3Error(error: unknown): error is H3Error {
  return (error as H3Error).data !== undefined;
}
