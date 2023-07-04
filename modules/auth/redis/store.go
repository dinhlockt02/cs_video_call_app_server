package authredis

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"time"
)

const prefix = "verify-email:"
const forgetPasswordPrefix = "forget-password:"

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client}
}

func (s *RedisStore) SetVerifyEmailCode(ctx context.Context, code string, userId string) error {
	log.Debug().Str("code", code).Str("userId", userId).Msg("set verify email code")
	err := s.client.Set(ctx, prefix+code, userId, 10*time.Minute).Err()
	if err != nil {
		return errors.Wrap(err, "can not set verify email code")
	}
	return nil
}

func (s *RedisStore) GetVerifyEmailCode(ctx context.Context, code string) (string, error) {
	log.Debug().Str("code", code).Msg("get verify email code")
	val, err := s.client.Get(ctx, prefix+code).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", errors.Wrap(err, "can not get verify email code")
	}
	return val, nil
}

func (s *RedisStore) SetForgetPasswordCode(ctx context.Context, code string, email string) error {
	err := s.client.Set(ctx, forgetPasswordPrefix+code, email, 10*time.Minute).Err()
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

func (s *RedisStore) GetForgetPasswordEmail(ctx context.Context, code string) (string, error) {
	val, err := s.client.Get(ctx, forgetPasswordPrefix+code).Result()
	return val, err
}
