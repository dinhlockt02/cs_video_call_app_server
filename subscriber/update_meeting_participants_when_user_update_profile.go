package subscriber

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func UpdateMeetingParticipantsWhenUserUpdateProfile(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)
	meetingStore := meetingstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for data := range ch {
			var user common.User
			err := json.Unmarshal([]byte(data), &user)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}

			go func(user common.User) {
				defer common.Recovery()
				err = meetingStore.UpdateParticipants(ctx, meetingstore.GetParticipantIdFilter(user.Id),
					&meetingmodel.Participant{
						Id:     user.Id,
						Name:   user.Name,
						Avatar: user.Avatar,
					})
				if err != nil {
					return
				}
			}(user)
		}
	}()
}
