package authbiz

import (
	"context"
	"fmt"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/mailer"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type ForgetPasswordAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type ForgetPasswordRedisStore interface {
	SetForgetPasswordCode(ctx context.Context, code string, email string) error
}

type ForgetPasswordBiz struct {
	mailer     mailer.Mailer
	authstore  ForgetPasswordAuthStore
	redisStore ForgetPasswordRedisStore
}

func NewForgetPasswordBiz(
	mailer mailer.Mailer,
	authstore ForgetPasswordAuthStore,
	redisStore ForgetPasswordRedisStore,
) *ForgetPasswordBiz {
	return &ForgetPasswordBiz{
		mailer:     mailer,
		authstore:  authstore,
		redisStore: redisStore,
	}
}

func (biz *ForgetPasswordBiz) Execute(ctx context.Context, email string) error {
	log.Debug().Any("email", email).Msg("forget password")
	receiver, err := biz.authstore.Find(ctx, authstore.GetEmailFilter(email))

	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find receiver"))
	}

	if receiver == nil {
		log.Debug().Msgf("user %s not found", email)
		return nil
	}

	code := biz.getCode(email)
	link := os.Getenv("RESET_PASSWORD_URL") + code

	err = biz.redisStore.SetForgetPasswordCode(ctx, code, email)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "can not set forget password code"))
	}

	go func(link string) {
		log.Info().Str("link", link).Msg("send ForgetPasswordEmail")
		err = biz.mailer.Send(authmodel.ForgetPasswordEmail, receiver.Email, "", authmodel.ForgetPasswordEmailBody(link))
		if err != nil {
			log.Error().Stack().Err(err).Msg("send email failed")
		}
	}(link)
	return nil
}

func (biz *ForgetPasswordBiz) getCode(user_id string) string {
	return fmt.Sprintf("%v:%v", time.Now().UnixNano(), user_id)
}
