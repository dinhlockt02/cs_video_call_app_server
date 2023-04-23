package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	"github.com/dinhlockt02/cs_video_call_app_server/middleware"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	usergin "github.com/dinhlockt02/cs_video_call_app_server/modules/user/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))

	user := g.Group("/user", middleware.Authentication(appCtx, userStore))
	{
		user.PUT("/self", usergin.UpdateSelf(appCtx))
		user.GET("/:id", usergin.GetUserDetail(appCtx))
	}
}
