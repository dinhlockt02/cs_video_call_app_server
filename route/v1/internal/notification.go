package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	notigin "github.com/dinhlockt02/cs_video_call_app_server/modules/notification/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitNotificationRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	notification := g.Group("/notification", authmiddleware.Authentication(appCtx))
	{
		notification.GET("", notigin.ListNotification(appCtx))
		notification.DELETE("", notigin.DeleteAllNotifications(appCtx))
		notification.DELETE("/:notificationId", notigin.DeleteNotificationById(appCtx))
		notification.GET("/notification-setting", notigin.GetNotificationSetting(appCtx))
		notification.PUT("/notification-setting", notigin.UpdateNotificationSetting(appCtx))
	}
}
