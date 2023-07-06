package subscriber

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	messagestore "github.com/dinhlockt02/cs_video_call_app_server/modules/message/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func UpdateMessagesWhenUserUpdateProfile(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)
	messageStore := messagestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for data := range ch {

			var user common.User
			err := json.Unmarshal([]byte(data), &user)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}

			go func(user common.User) {
				defer common.Recovery()
				err = messageStore.UpdateMany(ctx, messagestore.GetSenderIdFilter(user.Id),
					&messagemdl.UpdateMessage{Sender: &messagemdl.User{
						Id:     user.Id,
						Name:   user.Name,
						Avatar: user.Avatar,
					}})
				if err != nil {
					log.Error().Err(err).Msg("can not update message")
				}
			}(user)
		}
	}()
}
