package cache

import (
	"context"
	"rime-api/internal/models"

	"github.com/go-redis/redis/v8"
)

type UserStore struct {
	rdb *redis.Client
}

func (s *UserStore) Get(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}

func (s *UserStore) Set(ctx context.Context, user *models.User) error {
	return nil
}

func (s *UserStore) Delete(ctx context.Context, id int64) {

}
