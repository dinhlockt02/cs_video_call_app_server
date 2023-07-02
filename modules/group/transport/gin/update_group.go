package groupgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func UpdateGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *groupmdl.Group

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid request body")))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupId := c.Param("groupId")

		if _, err := common.ToObjectId(groupId); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid group id")))
		}

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		getGroupBiz := groupbiz.NewGetGroupBiz(groupRepo, appCtx.Notification())

		group, err := getGroupBiz.GetById(c.Request.Context(), groupId)
		if err != nil {
			panic(err)
		}

		isMember := false
		for _, member := range group.Members {
			if member == requester.GetId() {
				isMember = true
			}
		}

		if !isMember {
			panic(common.ErrForbidden(errors.New("user is not a member of group")))
		}

		groupFilter, err := common.GetIdFilter(groupId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err = groupbiz.NewUpdateGroupBiz(groupRepo).Update(c.Request.Context(), groupFilter, data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
