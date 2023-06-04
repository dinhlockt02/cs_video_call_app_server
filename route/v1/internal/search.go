package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	searchgin "github.com/dinhlockt02/cs_video_call_app_server/modules/search/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitSearchRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	g.GET("/search", authmiddleware.Authentication(appCtx), searchgin.Search(appCtx))
}
