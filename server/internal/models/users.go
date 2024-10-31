package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"hashed_password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
