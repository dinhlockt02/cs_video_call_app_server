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

func LeaveGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		requesterId := requester.GetId()
		groupId := context.Param("groupId")

		if !primitive.IsValidObjectID(groupId) {
			panic(common.ErrInvalidRequest(errors.New("invalid group id")))
		}

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		userFilter := map[string]interface{}{}
		_ = common.AddIdFilter(userFilter, requesterId)

		groupFilter := map[string]interface{}{}
		_ = common.AddIdFilter(groupFilter, groupId)

		session, err := appCtx.MongoClient().StartSession()
		if err != nil {
			panic(common.ErrInternal(err))
		}
		defer session.EndSession(context.Request.Context())

		err = session.StartTransaction()
		if err != nil {
			panic(common.ErrInternal(err))
		}

		defer func() {
			session.EndSession(context.Request.Context())
		}()

		if err = groupbiz.NewLeaveGroupBiz(groupRepo).Leave(context.Request.Context(), userFilter, groupFilter); err != nil {
			panic(err)
		}

		err = session.CommitTransaction(context.Request.Context())
		if err != nil {
			panic(common.ErrInternal(err))
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
