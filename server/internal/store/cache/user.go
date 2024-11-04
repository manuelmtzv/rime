package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"rime-api/internal/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserStore struct {
	rdb *redis.Client
}

const UserExpTime = time.Hour * 2

func (s *UserStore) Get(ctx context.Context, id string) (*models.User, error) {
	userKey := fmt.Sprintf("user-%s", id)

	dataString, err := s.rdb.Get(ctx, userKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var u models.User
	if dataString != "" {
		err := json.Unmarshal([]byte(dataString), &u)
		if err != nil {
			return nil, err
		}
	}

	return &u, nil
}

func (s *UserStore) Set(ctx context.Context, user *models.User) error {
	userKey := fmt.Sprintf("user-%s", user.ID)

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, userKey, json, UserExpTime).Err()
}

func (s *UserStore) Delete(ctx context.Context, id string) {
	userKey := fmt.Sprintf("user-%s", id)
	s.rdb.Del(ctx, userKey)
}
