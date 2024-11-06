package models

type Writing struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Content   string `json:"text"`
	Author    *User  `json:"-"`
	AuthorID  string `json:"author_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
