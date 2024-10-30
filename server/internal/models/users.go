package models

import "rime-api/internal/validator"

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"hashed_password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(user.LastName != "", "last_name", "must be provided")

	v.Check(user.Username != "", "username", "must be provided")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(user.Email != "" && validator.Matches(user.Email, validator.EmailRX), "email", "must be a valid email address")
}
