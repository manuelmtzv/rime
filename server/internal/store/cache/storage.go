package cache

import (
	"context"
	"rime-api/internal/models"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Users interface {
		Get(context.Context, string) (*models.User, error)
		Set(context.Context, *models.User) error
		Delete(context.Context, string)
	}
}

func NewRedisStorage(rbd *redis.Client) Storage {
	return Storage{
		Users: &UserStore{rdb: rbd},
	}
}
