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

func UpdateNotificationWhenGroupUpdated(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicGroupUpdated)

	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {

		for data := range ch {
			var group common.Group
			err := json.Unmarshal([]byte(data), &group)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}
			go func(group common.Group) {
				defer common.Recovery()

				// Update subject
				go func(group common.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetSubjectFilter(group.Id, notimodel.Group),
						&notimodel.UpdateNotification{Subject: &notimodel.NotificationObject{
							Name:  group.Name,
							Image: &group.ImageURL,
							Id:    group.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

				// Update direct
				go func(group common.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetDirectFilter(group.Id, notimodel.Group),
						&notimodel.UpdateNotification{Direct: &notimodel.NotificationObject{
							Name:  group.Name,
							Image: &group.ImageURL,
							Id:    group.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

				// Update indirect
				go func(group common.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetIndirectFilter(group.Id, notimodel.Group),
						&notimodel.UpdateNotification{Indirect: &notimodel.NotificationObject{
							Name:  group.Name,
							Image: &group.ImageURL,
							Id:    group.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

				// Update prep
				go func(group common.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetPrepFilter(group.Id, notimodel.Group),
						&notimodel.UpdateNotification{Prep: &notimodel.NotificationObject{
							Name:  group.Name,
							Image: &group.ImageURL,
							Id:    group.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

			}(group)
		}
	}()
}
