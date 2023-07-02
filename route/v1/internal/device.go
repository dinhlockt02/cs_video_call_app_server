package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	devicegin "github.com/dinhlockt02/cs_video_call_app_server/modules/device/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitDeviceRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	user := g.Group("/device", authmiddleware.Authentication(appCtx))
	{
		user.PUT("", devicegin.UpdateDevice(appCtx))
		user.GET("", devicegin.GetDevices(appCtx))
		user.DELETE("/:deviceId", devicegin.DeleteDevice(appCtx))
	}
}
