package meetinggin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	meetingbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/biz"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func GetMeeting(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupId := c.Param("groupId")

		if _, err := common.ToObjectId(groupId); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid group id")))
		}

		meetingId := c.Param("meetingId")

		if _, err := common.ToObjectId(meetingId); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid meeting id")))
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

		meetingStore := meetingstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		meetingRepo := meetingrepo.NewMeetingRepository(meetingStore, userStore)
		getMeetingsbiz := meetingbiz.NewGetMeetingBiz(meetingRepo, groupRepo)

		data, err := getMeetingsbiz.Get(c.Request.Context(), requester.GetId(), groupId, meetingId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
