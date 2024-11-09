import type { $Fetch, NitroFetchRequest } from "nitropack";
import type { ListResponse, Tag } from "@/types";

export const tagRepository = <T>(fetch?: $Fetch<T, NitroFetchRequest>) => {
  if (!fetch) {
    fetch = useNuxtApp().$api;
  }

  return {
    async getTags(): Promise<ListResponse<Tag>> {
      return fetch<ListResponse<Tag>>("/tags");
    },
    async getPopularTags(): Promise<ListResponse<Tag>> {
      return fetch<ListResponse<Tag>>("/tags/popular");
    },
  };
};
