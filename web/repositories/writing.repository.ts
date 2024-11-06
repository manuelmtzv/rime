import type { $Fetch, NitroFetchRequest } from "nitropack";
import type { ListResponse, Writing } from "@/types";

export const writingRepository = <T>(fetch?: $Fetch<T, NitroFetchRequest>) => {
  if (!fetch) {
    fetch = useNuxtApp().$api;
  }

  return {
    async getWriting(id: string): Promise<Writing> {
      return fetch<Writing>(`/writings/${id}`);
    },

    async getWritings(): Promise<ListResponse<Writing>> {
      return fetch<ListResponse<Writing>>("/writings");
    },
  };
};
