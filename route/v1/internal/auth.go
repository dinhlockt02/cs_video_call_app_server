package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authgin "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	auth := g.Group("/auth")
	{
		auth.POST("/register", authgin.Register(appCtx))
		auth.POST("/login", authgin.Login(appCtx))
		auth.POST("/login-with-firebase", authgin.LoginWithFirebase(appCtx))
	}
}
