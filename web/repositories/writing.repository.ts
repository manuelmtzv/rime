import type {
  DataResponse,
  ListResponse,
  Writing,
  WritingCreate,
} from "@/types";

export const writingRepository = <T>() => {
  const fetch = useNuxtApp().$serverApi;
  const ifetch = useNuxtApp().$internalApi;

  return {
    async createWriting(data: WritingCreate): Promise<DataResponse<Writing>> {
      return ifetch<DataResponse<Writing>>("/writings", {
        method: "POST",
        body: JSON.stringify(data),
      });
    },
    async getWriting(id: string): Promise<DataResponse<Writing>> {
      return fetch<DataResponse<Writing>>(`/writings/${id}`);
    },
    async getWritings(): Promise<ListResponse<Writing>> {
      return fetch<ListResponse<Writing>>("/writings");
    },
  };
};
