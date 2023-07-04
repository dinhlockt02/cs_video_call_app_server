package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authredis "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/redis"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResetPassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		body := &authmodel.ResetPasswordBody{}

		err := c.ShouldBind(body)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewResetPasswordBiz(
			authStore,
			authredis.NewRedisStore(
				appCtx.Redis(),
			),
			appCtx.Hasher())
		err = biz.Execute(c.Request.Context(), body)
		if err != nil {
			panic(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
