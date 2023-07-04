package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	"github.com/dinhlockt02/cs_video_call_app_server/components/mailer"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ResetPasswordAuthStore interface {
	ResetPassword(ctx context.Context, filter map[string]interface{}, data *authmodel.ResetPasswordBody) error
}

type ResetPasswordRedisStore interface {
	GetForgetPasswordEmail(ctx context.Context, code string) (string, error)
}

type ResetPasswordBiz struct {
	mailer         mailer.Mailer
	authstore      ResetPasswordAuthStore
	redisStore     ResetPasswordRedisStore
	passwordHasher hasher.Hasher
}

func NewResetPasswordBiz(
	authstore ResetPasswordAuthStore,
	redisStore ResetPasswordRedisStore,
	passwordHasher hasher.Hasher,
) *ResetPasswordBiz {
	return &ResetPasswordBiz{
		authstore:      authstore,
		redisStore:     redisStore,
		passwordHasher: passwordHasher,
	}
}

func (biz *ResetPasswordBiz) Execute(ctx context.Context, data *authmodel.ResetPasswordBody) error {
	log.Debug().Any("data", data).Msg("reset password")
	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid data"))
	}

	email, err := biz.redisStore.GetForgetPasswordEmail(ctx, data.Code)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not get forget password email"))
	}

	if m := common.EmailRegexp.Match([]byte(email)); !m {
		return common.ErrInvalidRequest(errors.New(authmodel.InvalidCode))
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.Password)
	if err != nil {
		return common.ErrInternal(err)
	}
	data.Password = hashedPassword

	err = biz.authstore.ResetPassword(ctx, map[string]interface{}{
		"email": email,
	}, data)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not reset password"))
	}
	return nil
}
