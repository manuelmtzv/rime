import type { User } from "./user.type";
import type { WritingLike } from "./like.type";

export type Writing = {
  id: string;
  title: string;
  text: string;
  type: WritingType;
  createdAt: string;
  updatedAt: string;
  author?: Partial<User>;
  likes?: WritingLike[];
};

export type WritingType = "poetry";
