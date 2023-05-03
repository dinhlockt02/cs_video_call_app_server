package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	usergin "github.com/dinhlockt02/cs_video_call_app_server/modules/user/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	user := g.Group("/user", authmiddleware.Authentication(appCtx))
	{
		user.PUT("/self", usergin.UpdateSelf(appCtx))
		user.GET("/self", usergin.GetSelfDetail(appCtx))
		user.GET("/:id", usergin.GetUserDetail(appCtx))

	}
}
