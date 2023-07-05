package messagegin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	messagebiz "github.com/dinhlockt02/cs_video_call_app_server/modules/message/biz"
	messagerepo "github.com/dinhlockt02/cs_video_call_app_server/modules/message/repository"
	messagestore "github.com/dinhlockt02/cs_video_call_app_server/modules/message/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnpinMessage(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		messageId := c.Param("messageId")

		msgStore := messagestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := messagerepo.NewMessageRepository(msgStore, userStore)
		biz := messagebiz.NewUnpinMessageBiz(repo)
		err := biz.Unpin(c.Request.Context(), requesterId, messageId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
