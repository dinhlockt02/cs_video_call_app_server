package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	"github.com/dinhlockt02/cs_video_call_app_server/middleware"
	friendgin "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/transport/gin"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
)

func InitFriendRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))

	friend := g.Group("/friend", middleware.Authentication(appCtx, userStore))
	{
		friend.GET("/request/sent", friendgin.GetSentRequest(appCtx))
		friend.GET("/request/received", friendgin.GetReceivedRequest(appCtx))
		friend.POST("/request/:id", friendgin.SendRequest(appCtx))
		friend.DELETE("/request/:id", friendgin.RecallRequest(appCtx))
		friend.POST("/request/:id/accept", friendgin.AcceptRequest(appCtx))
		friend.DELETE("/request/:id/reject", friendgin.RejectRequest(appCtx))

	}
}
