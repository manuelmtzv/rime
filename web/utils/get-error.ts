import { H3Error } from "h3";
import type { H3ErrorPayload } from "@/types/api.type";

type config = {
  defaultErrorMessage?: string;
  internal?: boolean;
};

export function mapError(error: unknown): Error {
  if (error instanceof Error) {
    return error;
  }

  return new Error("An error occurred");
}

export function mapH3Error(error: unknown): H3Error {
  if (isH3Error(error)) {
    return error as H3Error;
  }

  return new H3Error("An error occurred", {
    cause: error instanceof Error ? error.message : JSON.stringify(error),
  });
}

export function getErrorMessage(
  error: any,
  { internal = false, defaultErrorMessage = "An error occurred" }: config = {}
): string {
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
