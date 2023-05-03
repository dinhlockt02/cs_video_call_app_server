package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	friendgin "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitFriendRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	friend := g.Group("/friend", authmiddleware.Authentication(appCtx))
	{
		friend.GET("/", friendgin.FindFriend(appCtx))
		friendRequest := friend.Group("/request")
		{
			friendRequest.GET("/sent", friendgin.GetSentRequest(appCtx))
			friendRequest.GET("/received", friendgin.GetReceivedRequest(appCtx))
			friendRequest.POST("/:id", friendgin.SendRequest(appCtx))
			friendRequest.DELETE("/:id", friendgin.RecallRequest(appCtx))
			friendRequest.POST("/:id/accept", friendgin.AcceptRequest(appCtx))
			friendRequest.DELETE("/:id/reject", friendgin.RejectRequest(appCtx))
		}
		friend.DELETE("/:id", friendgin.Unfriend(appCtx))
		friend.PUT("/:id/block", friendgin.Block(appCtx))
		friend.PUT("/:id/unblock", friendgin.Unblock(appCtx))
	}
}
