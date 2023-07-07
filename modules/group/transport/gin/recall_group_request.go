package groupgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func RecallRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		requesterId := requester.GetId()
		friendId := context.Param("friendId")
		groupId := context.Param("groupId")

		if !primitive.IsValidObjectID(friendId) {
			panic(common.ErrInvalidRequest(errors.New("invalid friend id")))
		}
		if !primitive.IsValidObjectID(groupId) {
			panic(common.ErrInvalidRequest(errors.New("invalid group id")))
		}

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		recallRequestBiz := groupbiz.NewRecallGroupRequestBiz(groupRepo, appCtx.Notification(), appCtx.PubSub())
		if err := recallRequestBiz.RecallRequest(context.Request.Context(), requesterId, groupId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
