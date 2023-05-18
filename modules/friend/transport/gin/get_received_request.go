package friendgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetReceivedRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)

		receivedId := requester.GetId()

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		getReceivedRequestsBiz := friendbiz.NewGetReceivedRequestBiz(friendStore)
		result, err := getReceivedRequestsBiz.GetReceivedRequest(c.Request.Context(), receivedId)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
