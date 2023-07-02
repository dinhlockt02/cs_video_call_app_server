package devicegin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	devicebiz "github.com/dinhlockt02/cs_video_call_app_server/modules/device/biz"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func UpdateDevice(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		var deviceData devicemodel.UpdateDevice

		err := context.ShouldBind(&deviceData)

		if err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		updateDeviceBiz := devicebiz.NewUpdateDeviceBiz(deviceStore)

		idFilter, err := common.GetIdFilter(requester.GetDeviceId())
		if err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid device id")))
		}

		if err = updateDeviceBiz.Update(context.Request.Context(), idFilter, &deviceData); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
