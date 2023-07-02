package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authmiddleware "github.com/dinhlockt02/cs_video_call_app_server/middleware/authenticate"
	meetinggin "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitMeetingRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	call := g.Group("/meeting", authmiddleware.Authentication(appCtx))
	{
		groupCall := call.Group("/:groupId")
		{
			groupCall.GET("", meetinggin.ListMeetings(appCtx))
			groupCall.POST("", meetinggin.CreateMeeting(appCtx))
			groupCall.POST("/:meetingId", meetinggin.JoinMeeting(appCtx))
		}
	}
}
