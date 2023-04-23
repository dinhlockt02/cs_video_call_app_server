package usergin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	userbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/user/biz"
	userrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/user/repository"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetUserDetail(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)

		userId := c.Param("id")
		if !primitive.IsValidObjectID(userId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		userRepo := userrepo.NewUserRepository(userStore, friendStore)
		userDetailBiz := userbiz.NewUserDetailBiz(userRepo)

		user, err := userDetailBiz.GetUserDetail(c.Request.Context(), userId, requester.GetId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
