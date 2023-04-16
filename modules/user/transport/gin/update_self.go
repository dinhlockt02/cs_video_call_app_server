package usergin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	userbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/user/biz"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateSelf(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		var updateData usermodel.UpdateUser

		err := context.ShouldBind(&updateData)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		updateUserBiz := userbiz.NewUpdateUserBiz(userStore)

		id, err := primitive.ObjectIDFromHex(requester.GetId())
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = updateUserBiz.Update(context.Request.Context(), map[string]interface{}{
			"_id": id,
		}, &updateData); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
