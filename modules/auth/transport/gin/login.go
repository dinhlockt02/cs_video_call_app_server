package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		type Body struct {
			Data   authmodel.LoginUser `json:"data"`
			Device devicemodel.Device  `json:"device"`
		}

		var body = Body{
			Data:   authmodel.LoginUser{},
			Device: devicemodel.Device{},
		}

		if err := context.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		log.Debug().Str("email", body.Data.Email).Str("password", body.Data.Password).Msg("data")

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewLoginBiz(appCtx.TokenProvider(), authStore, deviceStore, appCtx.Hasher())
		result, err := biz.Login(context.Request.Context(), &body.Data, &body.Device)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
