package models

type Writing struct {
	ID        string `json:"id"`
	Type      string `json:"type,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"text,omitempty"`
	Author    *User  `json:"author,omitempty"`
	AuthorID  string `json:"authorId,omitempty"`
	Tags      []*Tag `json:"tags,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type WritingDetails struct {
	ID   string `json:"id"`
	Tags []*Tag `json:"tags"`
}
