package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	"github.com/dinhlockt02/cs_video_call_app_server/middleware"
	authgin "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/transport/gin"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))

	auth := g.Group("/auth")
	{
		auth.POST("/register", authgin.Register(appCtx))
		auth.POST("/login", authgin.Login(appCtx))
		auth.POST("/login-with-firebase", authgin.LoginWithFirebase(appCtx))
		auth.POST("/send-verify-email", middleware.Authentication(appCtx, userStore), authgin.SendVerifyEmail(appCtx))
		auth.GET("/verify-email", authgin.VerifyEmail(appCtx))
	}
}
