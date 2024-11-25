package models

type Comment struct {
	ID        string   `json:"id"`
	Content   string   `json:"content"`
	AuthorID  string   `json:"authorId"`
	Author    *User    `json:"author,omitempty"`
	WritingID string   `json:"postId"`
	Writing   *Writing `json:"post,omitempty"`
	ReplyTo   string   `json:"replyTo"`
	Comment   *Comment `json:"comment,omitempty"`
	CreatedAt string   `json:"createdAt"`
}
