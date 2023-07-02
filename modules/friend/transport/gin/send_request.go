package friendgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)

		senderId := requester.GetId()
		receiverId := c.Param("id")

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		friendRepo := friendrepo.NewFriendRepository(friendStore, requestStore)

		sendRequestBiz := friendbiz.NewSendRequestBiz(friendRepo, appCtx.Notification())
		if err := sendRequestBiz.SendRequest(c.Request.Context(), senderId, receiverId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
