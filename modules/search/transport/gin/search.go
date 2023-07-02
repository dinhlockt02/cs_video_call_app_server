package searchgin

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	friendbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/biz"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

func Search(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		searchTerm := c.Query("term")

		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)

		wg := sync.WaitGroup{}
		wg.Add(2)
		rs := map[string]interface{}{}
		go func() {
			defer common.Recovery()
			defer wg.Done()
			friends, err := searchFriend(c.Request.Context(), appCtx, requester, searchTerm)
			if err != nil {
				panic(err)
			}
			rs["friends"] = friends
		}()
		go func() {
			defer common.Recovery()
			defer wg.Done()
			groups, err := searchGroup(c.Request.Context(), appCtx, requester, searchTerm)
			if err != nil {
				panic(err)
			}
			rs["groups"] = groups
		}()

		wg.Wait()

		c.JSON(http.StatusOK, gin.H{"data": rs})
	}
}

func searchFriend(ctx context.Context, appCtx appcontext.AppContext, requester common.Requester, searchTerm string) ([]friendmodel.FriendUser, error) {
	friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	findFriendBiz := friendbiz.NewFindFriendBiz(friendStore)

	filter := map[string]interface{}{}
	err := common.AddIdFilter(map[string]interface{}{}, requester.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "invalid requester id")
	}
	friends, err := findFriendBiz.FindFriend(ctx, filter, common.GetTextSearch(searchTerm, false, false))
	if err != nil {
		return nil, errors.Wrap(err, "can not find friends")
	}
	return friends, nil
}

func searchGroup(ctx context.Context, appCtx appcontext.AppContext, requester common.Requester, searchTerm string) ([]groupmdl.Group, error) {
	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupRepo := grouprepo.NewGroupRepository(
		groupStore,
		requestStore,
	)
	listGroupBiz := groupbiz.NewListGroupBiz(groupRepo)
	groups, err := listGroupBiz.List(ctx, requester.GetId(), common.GetTextSearch(searchTerm, false, true))
	if err != nil {
		return nil, errors.Wrap(err, "can not find groups")
	}
	return groups, nil
}
