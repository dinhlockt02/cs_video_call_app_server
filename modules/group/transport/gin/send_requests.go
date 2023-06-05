package groupgin

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sync"
)

func SendGroupRequests(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data struct {
			Friends []string `json:"friends"`
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		groupId := c.Param("groupId")

		err := c.ShouldBind(&data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if !primitive.IsValidObjectID(groupId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		sendGroupRequestBiz := groupbiz.NewSendGroupRequestBiz(groupRepo)

		group, err := groupbiz.NewGetGroupBiz(groupRepo).GetById(c.Request.Context(), groupId)
		if err != nil {
			panic(err)
		}

		members := map[string]struct{}{}
		for _, member := range group.Members {
			members[member] = struct{}{}
		}

		go func() {
			defer common.Recovery()
			wg := sync.WaitGroup{}
			for _, friend := range data.Friends {
				if !primitive.IsValidObjectID(friend) {
					log.Err(common.ErrInvalidObjectId)
					continue
				}
				if _, ok := members[friend]; !ok && friend != requester.GetId() {
					go func(friendId string) {
						wg.Add(1)
						defer wg.Done()
						defer common.Recovery()
						err := sendGroupRequestBiz.SendRequest(context.Background(), requester.GetId(), friendId, group)
						if err != nil {
							log.Error().Msgf("%v\n", err)
						}
					}(friend)
				}
			}
			wg.Wait()
		}()

		c.JSON(http.StatusCreated, gin.H{"data": true})
	}
}
