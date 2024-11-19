package models

type WritingLike struct {
	AuthorID  string   `json:"authorId"`
	Author    *User    `json:"author,omitempty"`
	WritingID string   `json:"writingId"`
	Writing   *Writing `json:"writing,omitempty"`
	CreatedAt string   `json:"createdAt"`
}

type CommentLike struct {
	AuthorID  string   `json:"authorId"`
	Author    *User    `json:"author,omitempty"`
	CommentID string   `json:"commentId"`
	Comment   *Comment `json:"comment,omitempty"`
	CreatedAt string   `json:"createdAt"`
}
