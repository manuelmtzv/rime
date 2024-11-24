import type { User } from "./user.type";

export type WritingLike = {
  authorId: string;
  author?: User;
  writingId: string;
  createdAt: string;
};

export type CommentLike = {
  authorId: string;
  author?: User;
  commentId: string;
  createdAt: string;
};

export const likeEntities = ["writings", "comments"] as const;
export type LikeEntity = (typeof likeEntities)[number];
