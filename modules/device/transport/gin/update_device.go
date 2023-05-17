package devicegin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	devicebiz "github.com/dinhlockt02/cs_video_call_app_server/modules/device/biz"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateDevice(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		var deviceData devicemodel.UpdateDevice

		err := context.ShouldBind(&deviceData)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		updateDeviceBiz := devicebiz.NewUpdateDeviceBiz(deviceStore)

		id, err := primitive.ObjectIDFromHex(requester.GetDeviceId())
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = updateDeviceBiz.Update(context.Request.Context(), map[string]interface{}{
			"_id": id,
		}, &deviceData); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
