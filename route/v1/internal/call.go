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
		call.POST("/:friendId/:callRoomId", callgin.CreateNewCall(appCtx))
		call.DELETE("/:friendId/:callRoomId/reject", callgin.RejectCall(appCtx))
		call.DELETE("/:friendId/:callRoomId/abandon", callgin.AbandonCall(appCtx))
	}
}
