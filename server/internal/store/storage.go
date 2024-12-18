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
		FindPopular(context.Context) ([]*models.PopularUser, error)
		FindByIdentifier(context.Context, string) (*models.User, error)
		Update(context.Context, *models.User) error
		Delete(context.Context, string) error
	}
	Writings interface {
		Create(context.Context, *models.Writing) error
		FindAll(context.Context) ([]*models.Writing, error)
		FindOne(context.Context, string) (*models.Writing, error)
		FindDetails(context.Context, string) (*models.WritingDetails, error)
		ComposeFeed(context.Context, *string) ([]*models.Writing, error)
		Update(context.Context, *models.Writing) error
		Delete(context.Context, string) error
	}
	Tags interface {
		Create(context.Context, *models.Tag) error
		FindAll(context.Context) ([]*models.Tag, error)
		FindOneByName(context.Context, string) (*models.Tag, error)
		FindPopular(context.Context) ([]*models.Tag, error)
	}
	Likes interface {
		CreateWritingLike(context.Context, *models.WritingLike) error
		CreateCommentLike(context.Context, *models.CommentLike) error
		DeleteWritingLike(context.Context, string, string) error
		DeleteCommentLike(context.Context, string, string) error
	}
	Comments interface {
		Create(context.Context, *models.Comment) error
		FindAll(context.Context, string) ([]*models.Comment, error)
		FindOne(context.Context, string) (*models.Comment, error)
		Update(context.Context, *models.Comment) error
		Delete(context.Context, string) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:    &UserStore{db},
		Writings: &WritingStore{db},
		Tags:     &TagStore{db},
		Likes:    &LikeStore{db},
		Comments: &CommentStore{db},
	}
}
