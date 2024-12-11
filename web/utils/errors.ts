import { H3Error } from "h3";
import type { ApiErrorPayload, H3ErrorPayload } from "@/types/api.type";
import { map } from "zod";

type config = {
  defaultErrorMessage?: string;
  internal?: boolean;
};

type ServerError = {
  error: string;
  status: number;
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

export function getApiErrorPayload(error: any): ApiErrorPayload {
  if (isH3Error(error)) {
    return (error as H3Error).data as ApiErrorPayload;
  }

  return {
    error: getErrorMessage(error),
  };
}

export function composeError(error: any): ServerError {
  const { error: errorMessage } = getApiErrorPayload(error);
  const { statusCode = 500 } = mapH3Error(error);

  return {
    error: errorMessage,
    status: statusCode,
  };
}

export function getErrorMessage(
  error: any,
  { defaultErrorMessage = "An error occurred" }: config = {}
): string {
  if (isH3Error(error)) {
    return (error as H3Error<Error>).data?.message || defaultErrorMessage;
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
