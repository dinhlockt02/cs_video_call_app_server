package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func IsEmailVerified(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		id, _ := primitive.ObjectIDFromHex(requester.GetId())

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewIsEmailVerifiedBiz(authStore)
		isVerfied, err := biz.IsEmailVerified(context.Request.Context(), map[string]interface{}{
			"_id": id,
		})
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": isVerfied})
	}
}
