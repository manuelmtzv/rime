import type { User } from "./user.type";

export type Writing = {
  id: string;
  title: string;
  text: string;
  type: WritingType;
  createdAt: string;
  updatedAt: string;
  author?: Partial<User>;
};

export type WritingType = "poetry";
