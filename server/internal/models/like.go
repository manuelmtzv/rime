package models

type Like struct {
	ID        string   `json:"id"`
	UserID    string   `json:"useId"`
	User      *User    `json:"user"`
	WritingID string   `json:"postId"`
	Writing   *Writing `json:"post"`
	CreatedAt string   `json:"createdAt"`
}
