package subscriber

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func UpdateNotificationWhenUserUpdateProfile(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)

	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {

		for data := range ch {
			var user common.User
			err := json.Unmarshal([]byte(data), &user)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}
			go func(user common.User) {
				defer common.Recovery()

				// Update subject
				go func(user common.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetSubjectFilter(user.Id, notimodel.User),
						&notimodel.UpdateNotification{Subject: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    user.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

				// Update direct
				go func(user common.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetDirectFilter(user.Id, notimodel.User),
						&notimodel.UpdateNotification{Direct: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    user.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

				// Update indirect
				go func(user common.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetIndirectFilter(user.Id, notimodel.User),
						&notimodel.UpdateNotification{Indirect: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    user.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

				// Update prep
				go func(user common.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetPrepFilter(user.Id, notimodel.User),
						&notimodel.UpdateNotification{Prep: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    user.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

			}(user)
		}
	}()
}
