package models

type Like struct {
	ID        string `json:"id"`
	AuthorID  string `json:"authorId"`
	Author    *User  `json:"author,omitempty"`
	CreatedAt string `json:"createdAt"`
}
