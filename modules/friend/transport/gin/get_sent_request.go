package friendgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSentRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		senderId := requester.GetId()

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		getSentRequestBiz := friendbiz.NewGetSentRequestBiz(friendStore)
		result, err := getSentRequestBiz.GetSentRequest(context.Request.Context(), senderId)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
