import type { $Fetch, NitroFetchRequest } from "nitropack";
import { type DataResponse, type LikeEntity, type WritingLike } from "@/types";

export const likeRepository = <T>(fetch?: $Fetch<T, NitroFetchRequest>) => {
  if (!fetch) {
    fetch = useNuxtApp().$serverApi;
  }

  const ifetch = useNuxtApp().$internalApi;

  return {
    async likeEntity(
      entity: LikeEntity,
      id: string
    ): Promise<DataResponse<WritingLike>> {
      return ifetch<DataResponse<WritingLike>>(`/likes`, {
        method: "POST",
        body: JSON.stringify({ entity, id }),
      });
    },
    async unlikeEntity(
      entity: LikeEntity,
      id: string
    ): Promise<DataResponse<WritingLike>> {
      return fetch<DataResponse<WritingLike>>(`/likes/${entity}/${id}`, {
        method: "DELETE",
      });
    },
  };
};
