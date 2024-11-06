package models

type Writting struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"tilte"`
	Content   string `json:"text"`
	Author    User   `json:"author"`
	AuthorID  string `json:"author_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
