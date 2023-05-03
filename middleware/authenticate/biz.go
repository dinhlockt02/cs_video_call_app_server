package authmiddleware

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	"net/http"
)

type AuthMiddlewareMongoStore interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*Device, error)
}

type authMiddlewareBiz struct {
	store         AuthMiddlewareMongoStore
	tokenProvider tokenprovider.TokenProvider
}

func NewAuthMiddlewareBiz(
	store AuthMiddlewareMongoStore,
	tokenProvider tokenprovider.TokenProvider,
) *authMiddlewareBiz {
	return &authMiddlewareBiz{
		store:         store,
		tokenProvider: tokenProvider,
	}
}

func (biz *authMiddlewareBiz) Authenticate(ctx context.Context, token string) (*Device, error) {
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
		return nil, common.ErrInvalidRequest(err)
	}

	device, err := biz.store.FindOne(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, common.ErrInternal(err)
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
