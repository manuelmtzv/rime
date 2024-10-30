package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
