package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	messagegin "github.com/dinhlockt02/cs_video_call_app_server/modules/message/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitMessageRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	user := g.Group("/message", authmiddleware.Authentication(appCtx))
	{
		user.POST("", messagegin.PinMessage(appCtx))
		user.DELETE("/:messageId", messagegin.UnpinMessage(appCtx))
		user.GET("/:groupId", messagegin.ListMessages(appCtx))
	}
}
