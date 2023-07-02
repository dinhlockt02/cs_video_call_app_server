package friendrepo

import (
	"context"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendstore "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/store"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

type Repository interface {
	FindRequest(
		ctx context.Context,
		sender string,
		receiver string,
	) (*requestmdl.Request, error)
	FindUser(
		ctx context.Context,
		filter map[string]interface{},
		options ...FindUserOption,
	) (*friendmodel.User, error)
	UpdateUser(
		ctx context.Context,
		filter map[string]interface{},
		user *friendmodel.User,
	) error
	DeleteRequest(
		ctx context.Context,
		filter map[string]interface{},
	) error
	FindRequests(
		ctx context.Context,
		filter map[string]interface{},
	) ([]requestmdl.Request, error)
	CreateRequest(
		ctx context.Context,
		request *requestmdl.Request,
	) error
}

type FriendRepository struct {
	friendstore  friendstore.Store
	requestStore requeststore.Store
}

func NewFriendRepository(
	friendstore friendstore.Store,
	requestStore requeststore.Store,
) *FriendRepository {
	return &FriendRepository{
		friendstore:  friendstore,
		requestStore: requestStore,
	}
}
