package friendgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindFriend(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		findFriendBiz := friendbiz.NewFindFriendBiz(friendStore)

		filter, _ := common.GetIdFilter(requester.GetId())
		friends, err := findFriendBiz.FindFriend(context.Request.Context(), filter, map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": friends})
	}
}
