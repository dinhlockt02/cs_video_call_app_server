package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/pkg/errors"
)

func UpdateGroupWhenRoomFinished(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicRoomFinished)
	meetingStore := meetingstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))

	go func() {
		for roomId := range ch {
			go func(roomId string) {
				defer common.Recovery()
				filter := map[string]interface{}{}

				err := common.AddIdFilter(filter, roomId)
				if err != nil {
					panic(errors.Wrap(err, "invalid room id"))
				}

				meeting, err := meetingStore.FindMeeting(context.Background(),
					filter)
				if err != nil {
					panic(errors.Wrap(err, "can not find meeting"))
				}

				if meeting == nil {
					panic(errors.Wrap(err, "meeting not found"))
				}

				meetingFilter := common.GetAndFilter(meetingstore.GetGroupFilter(meeting.GroupId), meetingstore.GetStatusFilter(meetingmodel.OnGoing))
				latestMeeting, err := meetingStore.FindMeeting(context.Background(), meetingFilter)
				if err != nil {
					panic(errors.Wrap(err, "can not find meeting"))
				}

				meetingId := ""
				if latestMeeting != nil && latestMeeting.Id != nil {
					meetingId = *meeting.Id
				}

				groupFilter, err := common.GetIdFilter(meeting.GroupId)
				if err != nil {
					panic(errors.Wrap(err, "can not find group"))
				}

				err = groupStore.UpdateGroup(ctx, groupFilter, &groupmdl.Group{
					LatestMeeting: &meetingId,
				})

			}(roomId)
		}
	}()
}
