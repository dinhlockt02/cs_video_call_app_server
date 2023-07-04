package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authredis "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/redis"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

func VerifyEmail(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		code, ok := context.GetQuery("code")
		if !ok {
			panic(common.ErrInvalidRequest(errors.New("invalid verify url")))
		}
		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err := authbiz.NewVerifyEmail(authStore, authredis.NewRedisStore(
			appCtx.Redis(),
		)).Verify(context.Request.Context(), code)
		if err != nil {
			log.Error().Stack().Err(err).Msg("verify email failed")
			context.Redirect(http.StatusBadRequest, "/verify/failure")
			return
		}
		context.Redirect(http.StatusFound, "/verify/success")
	}
}
