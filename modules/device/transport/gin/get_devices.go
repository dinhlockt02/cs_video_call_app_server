package devicegin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	devicebiz "github.com/dinhlockt02/cs_video_call_app_server/modules/device/biz"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDevices(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		devices, err := devicebiz.NewGetDevicesBiz(deviceStore).
			Get(context.Request.Context(), devicestore.GetUserIdFilter(requester.GetId()))
		if err != nil {
			panic(err)
		}
		for i := range devices {
			if *devices[i].Id == requester.GetDeviceId() {
				devices[i].IsCurrentDevice = common.Ptr(true)
			} else {
				devices[i].IsCurrentDevice = common.Ptr(false)
			}
		}
		context.JSON(http.StatusOK, gin.H{"data": devices})
	}
}
