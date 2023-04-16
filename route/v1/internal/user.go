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

	auth := g.Group("/user", middleware.Authentication(appCtx, userStore))
	{
		auth.PUT("/self", usergin.UpdateSelf(appCtx))
	}
}
