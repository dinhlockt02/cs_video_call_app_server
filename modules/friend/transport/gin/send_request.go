package friendgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func SendRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		senderId := requester.GetId()
		receiverId := context.Param("id")

		if !primitive.IsValidObjectID(senderId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}
		if !primitive.IsValidObjectID(receiverId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		sendRequestBiz := friendbiz.NewSendRequestBiz(friendStore)
		if err := sendRequestBiz.SendRequest(context.Request.Context(), senderId, receiverId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
