import type { $Fetch, NitroFetchRequest } from "nitropack";
import type { ListResponse, Writting } from "@/types";

export const writingRepository = <T>(fetch?: $Fetch<T, NitroFetchRequest>) => {
  if (!fetch) {
    fetch = useNuxtApp().$api;
  }

  return {
    async getWriting(id: string): Promise<Writting> {
      return fetch<Writting>(`/writings/${id}`);
    },

    async getWritings(): Promise<ListResponse<Writting>> {
      return fetch<ListResponse<Writting>>("/writings");
    },
  };
};
