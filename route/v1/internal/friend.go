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
		friend.GET("/", friendgin.FindFriend(appCtx))
		friendRequest := g.Group("/request")
		{
			friendRequest.GET("/sent", friendgin.GetSentRequest(appCtx))
			friendRequest.GET("/received", friendgin.GetReceivedRequest(appCtx))
			friendRequest.POST("/:id", friendgin.SendRequest(appCtx))
			friendRequest.DELETE("/:id", friendgin.RecallRequest(appCtx))
			friendRequest.POST("/:id/accept", friendgin.AcceptRequest(appCtx))
			friendRequest.DELETE("/:id/reject", friendgin.RejectRequest(appCtx))
		}

	}
}
