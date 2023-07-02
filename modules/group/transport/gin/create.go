package groupgin

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

func CreateGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *groupmdl.Group

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid request body")))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		invitedMembers := data.Members

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		createGroupBiz := groupbiz.NewCreateGroupBiz(groupRepo)

		if err := createGroupBiz.Create(c.Request.Context(), requester.GetId(), data); err != nil {
			panic(err)
		}

		go func() {
			defer common.Recovery()
			wg := sync.WaitGroup{}
			sendGroupRequestBiz := groupbiz.NewSendGroupRequestBiz(groupRepo, appCtx.Notification())
			for _, member := range invitedMembers {
				if member != requester.GetId() {
					wg.Add(1)
					go func(mem string) {
						defer wg.Done()
						defer common.Recovery()
						err := sendGroupRequestBiz.SendRequest(context.Background(), requester.GetId(), mem, data)
						if err != nil {
							log.Error().Err(err).Msg("send request failed")
						}
					}(member)
				}
			}
			wg.Wait()
		}()

		c.JSON(http.StatusCreated, gin.H{"data": data.Id})
	}
}
