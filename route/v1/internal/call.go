package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	callgin "github.com/dinhlockt02/cs_video_call_app_server/modules/call/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitCallRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	call := g.Group("/call", authmiddleware.Authentication(appCtx))
	{
		call.POST("/:friendId", callgin.CreateNewCall(appCtx))
		call.GET("/:callRoomId", callgin.JoinCall(appCtx))
		call.GET("", callgin.ListCalls(appCtx))
	}
}
