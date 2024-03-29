package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsEmailVerified(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewIsEmailVerifiedBiz(authStore)

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		idFilter, _ := common.GetIdFilter(requester.GetId())
		isVerfied, err := biz.IsEmailVerified(context.Request.Context(), idFilter)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": isVerfied})
	}
}
