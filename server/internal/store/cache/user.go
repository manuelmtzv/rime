package cache

import (
	"context"
	"encoding/json"
	"rime-api/internal/models"

	"github.com/go-redis/redis/v8"
)

type UserStore struct {
	rdb *redis.Client
}

func (s *UserStore) Get(ctx context.Context, id string) (*models.User, error) {
	cmd := s.rdb.Get(ctx, id)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	var u models.User

	err := json.Unmarshal([]byte(cmd.Val()), &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *UserStore) Set(ctx context.Context, user *models.User) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}

	cmd := s.rdb.Set(ctx, user.ID, b, 0)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (s *UserStore) Delete(ctx context.Context, id string) {
	s.rdb.Del(ctx, id)
}
