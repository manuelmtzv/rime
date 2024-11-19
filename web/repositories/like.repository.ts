import type { $Fetch, NitroFetchRequest } from "nitropack";
import {  type DataResponse, type LikeEntity, type WritingLike } from "@/types";

export const writingRepository = <T>(fetch?: $Fetch<T, NitroFetchRequest>) => {
  if (!fetch) {
    fetch = useNuxtApp().$serverApi;
  }

  return {
    async likeEntity(entity: LikeEntity, id: string): Promise<DataResponse<WritingLike>> {;
        return fetch<DataResponse<WritingLike>>(`/${entity}/${id}/like`, {
            method: "POST",
        });
    },
    async unlikeEntity(entity: LikeEntity, id: string): Promise<DataResponse<WritingLike>> {
        return fetch<DataResponse<WritingLike>>(`/${entity}/${id}/unlike`, {
            method: "DELETE",
        });
    }
  };
};
