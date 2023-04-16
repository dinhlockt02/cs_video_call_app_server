package middleware

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthenticateUserStore interface {
	FindById(ctx context.Context, userId string) (*usermodel.User, error)
}

func Authentication(appCtx appcontext.AppContext, store AuthenticateUserStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		authorizationHeader := strings.Split(c.GetHeader("Authorization"), " ")

		if len(authorizationHeader) != 2 || authorizationHeader[0] != "Bearer" {
			var unauthorizedError = common.NewFullErrorResponse(http.StatusUnauthorized,
				nil,
				"unauthorized",
				"Invalid header",
				"UnauthorizedError")

			panic(unauthorizedError)
		}

		tokenProvider := appCtx.TokenProvider()

		tokenPayload, err := tokenProvider.Validate(authorizationHeader[1])
		if err != nil {
			var unauthorizedError = common.NewFullErrorResponse(http.StatusUnauthorized,
				err,
				"unauthorized",
				err.Error(),
				"UnauthorizedError")

			panic(unauthorizedError)
		}

		user, err := store.FindById(c.Request.Context(), tokenPayload.UserId)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		if user == nil {
			var unauthorizedError = common.NewFullErrorResponse(http.StatusUnauthorized,
				nil,
				"unauthorized",
				"user not found",
				"UnauthorizedError")

			panic(unauthorizedError)
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
