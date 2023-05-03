package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authredis "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/redis"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendVerifyEmail(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err := authbiz.
			NewSendVerifyEmail(
				appCtx.Mailer(),
				authStore,
				authredis.NewRedisStore(
					appCtx.Redis(),
				),
			).
			Send(context.Request.Context(), requester.GetId(), false)
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
