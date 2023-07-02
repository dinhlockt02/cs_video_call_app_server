package devicegin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	devicebiz "github.com/dinhlockt02/cs_video_call_app_server/modules/device/biz"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func DeleteDevice(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		deviceId := context.Param("deviceId")
		deviceFilter, err := common.GetIdFilter(deviceId)
		if err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid device id: "+deviceId)))
		}

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err = devicebiz.NewDeleteDevicesBiz(deviceStore).Delete(context.Request.Context(), common.GetAndFilter(
			deviceFilter,
			devicestore.GetUserIdFilter(requester.GetId()),
		),
		)

		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
