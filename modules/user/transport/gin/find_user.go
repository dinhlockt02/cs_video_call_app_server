package usergin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	userbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/user/biz"
	userrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/user/repository"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func FindUser(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		type filterBody struct {
			Email string `form:"email,omitempty"`
		}

		var filter filterBody

		err := context.ShouldBind(&filter)

		if err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))

		friendRepo := friendrepo.NewFriendRepository(friendStore, requestStore)
		findUserRepo := userrepo.NewFindUserRepo(userStore, friendRepo)

		findUserBiz := userbiz.NewFindUserBiz(findUserRepo)

		user, err := findUserBiz.FindUser(context.Request.Context(), requester.GetId(), userstore.GetEmailFilter(filter.Email))
		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": user})
	}
}
