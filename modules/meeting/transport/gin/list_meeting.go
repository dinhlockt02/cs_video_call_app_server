package meetinggin

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	meetingbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/biz"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListMeetings(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupId := c.Param("groupId")

		if _, err := common.ToObjectId(groupId); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		getGroupBiz := groupbiz.NewGetGroupBiz(groupRepo)

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

		meetingStore := meetingstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		meetingRepo := meetingrepo.NewMeetingRepository(meetingStore)
		listMeetingsbiz := meetingbiz.NewListMeetingsBiz(meetingRepo)

		data, err := listMeetingsbiz.List(c.Request.Context(), groupId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
