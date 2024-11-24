import { type DataResponse, type LikeEntity, type WritingLike } from "@/types";

export const likeRepository = <T>() => {
  const ifetch = useNuxtApp().$internalApi;

  return {
    async likeEntity(
      entity: LikeEntity,
      id: string
    ): Promise<DataResponse<WritingLike>> {
      return ifetch<DataResponse<WritingLike>>("/likes", {
        method: "POST",
        body: JSON.stringify({ entity, id }),
      });
    },
    async unlikeEntity(
      entity: LikeEntity,
      id: string
    ): Promise<DataResponse<WritingLike>> {
      return ifetch<DataResponse<WritingLike>>("/likes", {
        method: "DELETE",
        body: JSON.stringify({ entity, id }),
      });
    },
  };
};
