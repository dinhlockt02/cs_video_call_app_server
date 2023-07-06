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

func UpdateRequestsWhenGroupUpdated(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicGroupUpdated)

	requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for data := range ch {
			var group common.Group
			err := json.Unmarshal([]byte(data), &group)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}
			go func(group common.Group) {
				defer common.Recovery()

				err = requestStore.UpdateRequests(
					context.Background(),
					requeststore.GetRequestGroupIdFilter(group.Id),
					&requestmdl.UpdateRequest{
						Group: &requestmdl.RequestGroup{
							Id:       group.Id,
							Name:     group.Name,
							ImageUrl: group.ImageURL,
						},
					})
				if err != nil {
					return
				}

			}(group)
		}
	}()
}
