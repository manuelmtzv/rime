package models

type Comment struct {
	ID        string   `json:"id"`
	Content   string   `json:"content"`
	UserID    string   `json:"userId"`
	User      *User    `json:"user"`
	WritingID string   `json:"postId"`
	Writing   *Writing `json:"post"`
	ReplyTo   string   `json:"replyTo"`
	Comment   *Comment `json:"comment"`
	CreatedAt string   `json:"createdAt"`
}
