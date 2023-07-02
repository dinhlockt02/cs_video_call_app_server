package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"strings"
)

type VerifyEmailAuthStore interface {
	UpdateEmailVerified(ctx context.Context, filter map[string]interface{}) error
}

type VerifyEmailRedisStore interface {
	GetVerifyEmailCode(ctx context.Context, code string) (string, error)
}

type VerifyEmailBiz struct {
	authstore  VerifyEmailAuthStore
	redisStore VerifyEmailRedisStore
}

func NewVerifyEmail(
	authstore VerifyEmailAuthStore,
	redisStore VerifyEmailRedisStore,
) *VerifyEmailBiz {
	return &VerifyEmailBiz{
		authstore:  authstore,
		redisStore: redisStore,
	}
}

func (biz *VerifyEmailBiz) Verify(ctx context.Context, code string) error {
	userId, err := biz.redisStore.GetVerifyEmailCode(ctx, code)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not get code from redis"))
	}
	if strings.TrimSpace(userId) == "" {
		return common.ErrInvalidRequest(errors.New(authmodel.InvalidVerifyCode))
	}

	idFilter, err := common.GetIdFilter(userId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "invalid id"))
	}
	err = biz.authstore.UpdateEmailVerified(ctx, idFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update email verified status"))
	}
	return nil
}
