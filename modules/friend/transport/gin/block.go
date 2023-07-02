package friendgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Block(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		userId := requester.GetId()
		blockedId := context.Param("id")

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		blockBiz := friendbiz.NewBlockBiz(friendStore)
		if err := blockBiz.Block(context.Request.Context(), userId, blockedId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
