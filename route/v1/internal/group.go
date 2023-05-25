package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	groupgin "github.com/dinhlockt02/cs_video_call_app_server/modules/group/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitGroupRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	group := g.Group("/group", authmiddleware.Authentication(appCtx))
	{
		group.POST("", groupgin.CreateGroup(appCtx))
	}
}
