package subscriber

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func UpdateRequestsWhenUserUpdateProfile(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)

	requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for data := range ch {
			var user common.User
			err := json.Unmarshal([]byte(data), &user)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}
			go func(user common.User) {
				defer common.Recovery()

				err = requestStore.UpdateRequests(
					context.Background(),
					requeststore.GetRequestSenderIdFilter(user.Id),
					&requestmdl.UpdateRequest{
						Sender: &requestmdl.RequestUser{
							Id:     user.Id,
							Name:   user.Name,
							Avatar: user.Avatar,
						},
					})
				if err != nil {
					return
				}

				err = requestStore.UpdateRequests(
					context.Background(),
					requeststore.GetRequestReceiverIdFilter(user.Id),
					&requestmdl.UpdateRequest{
						Receiver: &requestmdl.RequestUser{
							Id:     user.Id,
							Name:   user.Name,
							Avatar: user.Avatar,
						},
					})
				if err != nil {
					return
				}

			}(user)
		}
	}()
}
