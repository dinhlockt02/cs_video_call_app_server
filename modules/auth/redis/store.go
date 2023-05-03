package authredis

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/redis/go-redis/v9"
	"time"
)

const prefix = "verify-email:"

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *redisStore {
	return &redisStore{client: client}
}

func (s *redisStore) SetVerifyEmailCode(ctx context.Context, code string, user_id string) error {
	err := s.client.Set(ctx, prefix+code, user_id, 10*time.Minute).Err()
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

func (s *redisStore) GetVerifyEmailCode(ctx context.Context, code string) string {
	val, _ := s.client.Get(ctx, prefix+code).Result()
	return val
}
