package subscriber

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func UpdateCallsWhenUserUpdateProfile(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)
	callStore := callstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for data := range ch {

			var user common.User
			err := json.Unmarshal([]byte(data), &user)
			if err != nil {
				log.Error().Stack().Err(errors.Wrap(err, "can not unmarshal json")).Send()
			}

			go func(user common.User) {
				defer common.Recovery()
				calleeFilter := callstore.GetCalleeIdFilter(user.Id)

				err = callStore.UpdateMany(ctx, calleeFilter, &callmdl.UpdateCall{
					Callee: &callmdl.User{
						Id:     user.Id,
						Name:   user.Name,
						Avatar: user.Avatar,
					},
				})
				if err != nil {
					log.Error().Stack().Err(errors.Wrap(err, "can not update call's callee")).Send()
				}

				callerFilter := callstore.GetCallerIdFilter(user.Id)

				err = callStore.UpdateMany(ctx, callerFilter, &callmdl.UpdateCall{
					Caller: &callmdl.User{
						Id:     user.Id,
						Name:   user.Name,
						Avatar: user.Avatar,
					},
				})
				if err != nil {
					log.Error().Stack().Err(errors.Wrap(err, "can not update call's caller")).Send()
				}

			}(user)
		}
	}()
}
