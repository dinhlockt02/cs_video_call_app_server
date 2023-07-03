package callgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	callbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/call/biz"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func ListCalls(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		requesterId := requester.GetId()
		callId := context.Param("callId")

		filter := make(map[string]interface{})
		if status, ok := context.GetQuery("status"); ok {
			filter["status"] = status
		}
		if callee, ok := context.GetQuery("callee"); ok {
			filter["callee"] = callee
		}

		if !primitive.IsValidObjectID(callId) {
			panic(common.ErrInvalidRequest(errors.Wrap(common.ErrInvalidObjectId, "invalid friend id")))
		}
		callStore := callstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		callRepo := callrepo.NewCallRepository(
			userStore,
			callStore,
		)
		requesterFilter, _ := common.GetIdFilter(requesterId)
		token, err := callbiz.NewListCallsBiz(callRepo).
			List(context.Request.Context(), requesterFilter)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": token})
	}
}
