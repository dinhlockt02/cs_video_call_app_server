package messagegin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	messagebiz "github.com/dinhlockt02/cs_video_call_app_server/modules/message/biz"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	messagerepo "github.com/dinhlockt02/cs_video_call_app_server/modules/message/repository"
	messagestore "github.com/dinhlockt02/cs_video_call_app_server/modules/message/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func PinMessage(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		var data messagemdl.Message

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		msgStore := messagestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := messagerepo.NewMessageRepository(msgStore, userStore)
		biz := messagebiz.NewPinMessageBiz(repo)
		err := biz.Pin(c.Request.Context(), requesterId, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
