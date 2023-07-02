package groupgin

import (
	"context"
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

func SendGroupRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		friendId := c.Param("friendId")
		groupId := c.Param("groupId")

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

		getGroupBiz := groupbiz.NewGetGroupBiz(groupRepo, appCtx.Notification())

		group, err := getGroupBiz.GetById(context.Background(), groupId)
		if err != nil {
			panic(err)
		}

		sendGroupRequestBiz := groupbiz.NewSendGroupRequestBiz(groupRepo, appCtx.Notification())
		err = sendGroupRequestBiz.SendRequest(context.Background(), requester.GetId(), friendId, group)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, gin.H{"data": true})
	}
}
