package groupgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		listGroupBiz := groupbiz.NewListGroupBiz(groupRepo)
		groups, err := listGroupBiz.List(c.Request.Context(), requester.GetId())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, gin.H{"data": groups})
	}
}
