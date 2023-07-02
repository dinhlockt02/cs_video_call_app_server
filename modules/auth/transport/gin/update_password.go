package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func UpdatePassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data authmodel.UpdatePasswordUser

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewUpdatePasswordBiz(authStore, appCtx.Hasher())

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		idFilter, _ := common.GetIdFilter(requester.GetId())
		err := biz.Update(context.Request.Context(), idFilter, &data)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
