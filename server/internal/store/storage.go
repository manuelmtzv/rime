package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type Storage struct {
	Users interface {
		FindAll(context.Context) ([]*models.User, error)
		Create(context.Context, *models.User) error
		FindOne(context.Context, string) (*models.User, error)
		Update(context.Context, *models.User) error
		Delete(context.Context, string) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
