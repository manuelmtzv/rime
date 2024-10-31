package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.User) error
		FindAll(context.Context) ([]*models.User, error)
		FindOne(context.Context, string) (*models.User, error)
		FindByEmail(context.Context, string) (*models.User, error)
		Update(context.Context, *models.User) error
		Delete(context.Context, string) error
	}
	Writtings interface {
		Create(context.Context, *models.Writting) error
		FindAll(context.Context) ([]*models.Writting, error)
		FindOne(context.Context, string) (*models.Writting, error)
		Update(context.Context, *models.Writting) error
		Delete(context.Context, string) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:     &UserStore{db},
		Writtings: &WrittingStore{db},
	}
}
