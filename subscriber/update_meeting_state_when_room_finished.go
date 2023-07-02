package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/pkg/errors"
	"time"
)

func UpdateMeetingStateWhenRoomFinished(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicRoomFinished)

	meetingStore := meetingstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for roomId := range ch {
			go func(roomId string) {
				defer common.Recovery()
				filter := map[string]interface{}{}

				err := common.AddIdFilter(filter, roomId)
				if err != nil {
					panic(errors.Wrap(err, "invalid room id"))
				}

				now := time.Now()
				err = meetingStore.UpdateMeeting(context.Background(), filter, &meetingmodel.UpdateMeeting{
					Status:  meetingmodel.Ended,
					TimeEnd: &now,
				})
				if err != nil {
					panic(errors.Wrap(err, "can not update meeting"))
				}
			}(roomId)
		}
	}()
}
