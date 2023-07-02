package authbiz

import (
	"context"
	"fmt"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/mailer"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type SendVerifyEmailAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type SendVerifyEmailRedisStore interface {
	SetVerifyEmailCode(ctx context.Context, code string, user_id string) error
}

type SendVerifyEmail struct {
	mailer     mailer.Mailer
	authstore  SendVerifyEmailAuthStore
	redisStore SendVerifyEmailRedisStore
}

func NewSendVerifyEmail(
	mailer mailer.Mailer,
	authstore SendVerifyEmailAuthStore,
	redisStore SendVerifyEmailRedisStore,
) *SendVerifyEmail {
	return &SendVerifyEmail{
		mailer:     mailer,
		authstore:  authstore,
		redisStore: redisStore,
	}
}

func (biz *SendVerifyEmail) Send(ctx context.Context, receiverId string, isConcurrent bool) error {
	idFilter, err := common.GetIdFilter(receiverId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid receiver id"))
	}
	receiver, err := biz.authstore.Find(ctx, idFilter)

	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if receiver == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(authmodel.UserNotFound))
	}

	code := biz.getCode(receiverId)
	// TODO: Create a config object
	link := os.Getenv("VERIFY_EMAIL_URL") + code

	err = biz.redisStore.SetVerifyEmailCode(ctx, code, receiverId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not set verify email code"))
	}

	if isConcurrent {
		// TODO: do something about receiver name
		go func() {
			err = biz.mailer.Send(authmodel.VerifyEmailTitle, receiver.Email, "", authmodel.VerifyEmailBody(link))
			if err != nil {
				log.Error().Err(err).Stack().Msg("send email failed")
			}
		}()
	} else {
		err = biz.mailer.Send(authmodel.VerifyEmailTitle, receiver.Email, "", authmodel.VerifyEmailBody(link))
		if err != nil {
			return common.ErrInternal(errors.Wrap(err, "send email failed"))
		}
	}
	return nil
}

func (biz *SendVerifyEmail) getCode(userId string) string {
	return fmt.Sprintf("%v:%v", time.Now().UnixNano(), userId)
}
