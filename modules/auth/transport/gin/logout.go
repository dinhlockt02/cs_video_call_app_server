package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		deviteStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewLogoutBiz(deviteStore)

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		idFilter, _ := common.GetIdFilter(requester.GetDeviceId())
		err := biz.Logout(context.Request.Context(), idFilter)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
