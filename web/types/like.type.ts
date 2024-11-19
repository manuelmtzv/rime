export type WritingLike = {
    authorId: string; 
    writingId: string;
    createdAt: string
}

export const likeEntities = ['writings', 'comments'];

export type LikeEntity = typeof likeEntities[number];