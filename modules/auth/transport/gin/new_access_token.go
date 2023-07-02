package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func NewAccessToken(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		type Body struct {
			RefreshToken string `json:"refresh_token"`
		}

		var body Body

		if err := context.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		biz := authbiz.NewAccessTokenBiz(appCtx.TokenProvider())
		result, err := biz.New(context.Request.Context(), body.RefreshToken)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
