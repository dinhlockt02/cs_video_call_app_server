package internal

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	"github.com/gin-gonic/gin"
	lkwebhook "github.com/livekit/protocol/webhook"
	"net/http"
)

func InitLiveKitWebhookRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	webhook := g.Group("/livekit-webhook")
	{
		webhook.POST("", func(c *gin.Context) {

			event, err := lkwebhook.ReceiveWebhookEvent(c.Request, appCtx.LiveKitService().AuthProvider())
			if err != nil {
				panic(err)
				return
			}
			if event.Event == "room_finished" {
				err = appCtx.PubSub().Publish(c.Request.Context(), common.TopicRoomFinished, event.Room.Name)
				if err != nil {
					panic(err)
					return
				}
			}
			c.Status(http.StatusOK)
		})
	}
}
