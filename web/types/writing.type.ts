import type { User } from "./user.type";
import type { WritingLike } from "./like.type";

export type Writing = {
  id: string;
  title: string;
  content: string;
  type: WritingType;
  createdAt: string;
  updatedAt: string;
  author?: Partial<User>;
  likes?: WritingLike[];
};

export type WritingCreate = {
  title: string;
  content: string;
  type: WritingType;
};

export type WritingType = "poetry";
