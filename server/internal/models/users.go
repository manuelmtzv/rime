package models

import (
	"database/sql"
	"rime-server/internal/validator"
)

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(user.LastName != "", "last_name", "must be provided")

	v.Check(user.Username != "", "username", "must be provided")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(user.Email != "" && validator.Matches(user.Email, validator.EmailRX), "email", "must be a valid email address")
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Get(id string) (*User, error) {
	return nil, nil
}

func (u UserModel) Insert(user *User) error {
	return nil
}

func (u UserModel) Update(user *User) error {
	return nil
}

func (u UserModel) Delete(id string) error {
	return nil
}
