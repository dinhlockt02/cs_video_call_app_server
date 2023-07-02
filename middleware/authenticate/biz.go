package authmiddleware

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	"github.com/pkg/errors"
	"net/http"
)

type Store interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*Device, error)
}

type Biz struct {
	store         Store
	tokenProvider tokenprovider.TokenProvider
}

func NewAuthMiddlewareBiz(
	store Store,
	tokenProvider tokenprovider.TokenProvider,
) *Biz {
	return &Biz{
		store:         store,
		tokenProvider: tokenProvider,
	}
}

func (biz *Biz) Authenticate(ctx context.Context, token string) (*Device, error) {
	tokenPayload, err := biz.tokenProvider.Validate(token)
	if err != nil {
		return nil, common.NewFullErrorResponse(
			http.StatusUnauthorized,
			err,
			"unauthorized",
			err.Error(),
			"UnauthorizedError",
		)
	}

	id, err := common.ToObjectId(tokenPayload.Id)
	if err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid payload token"))
	}

	device, err := biz.store.FindOne(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find device"))
	}

	if device == nil {
		return nil, common.NewFullErrorResponse(http.StatusUnauthorized,
			nil,
			"unauthorized",
			"user not found",
			"UnauthorizedError")
	}

	return device, nil
}
