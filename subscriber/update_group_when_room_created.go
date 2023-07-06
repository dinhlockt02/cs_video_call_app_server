package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/pkg/errors"
)

func UpdateGroupWhenRoomCreated(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicRoomCreated)
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

				groupFilter, err := common.GetIdFilter(meeting.GroupId)
				if err != nil {
					panic(errors.Wrap(err, "can not find group"))
				}

				err = groupStore.UpdateGroup(ctx, groupFilter, &groupmdl.Group{
					LatestMeeting: meeting.Id,
				})

			}(roomId)
		}
	}()
}
