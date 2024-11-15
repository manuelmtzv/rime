import type { $Fetch, NitroFetchRequest } from "nitropack";
import type { ListResponse, PopularUser, User } from "@/types";

export const userRepository = <T>(fetch?: $Fetch<T, NitroFetchRequest>) => {
  if (!fetch) {
    fetch = useNuxtApp().$serverApi;
  }

  return {
    async getUsers(): Promise<ListResponse<User>> {
      return fetch<ListResponse<User>>("/users");
    },

    async getPopularUsers(): Promise<ListResponse<PopularUser>> {
      return fetch<ListResponse<PopularUser>>("/users/popular");
    },
  };
};
