package models

type Like struct {
	ID        string `json:"id"`
	UserID    string `json:"useId"`
	User      *User  `json:"user,omitempty"`
	CreatedAt string `json:"createdAt"`
}
