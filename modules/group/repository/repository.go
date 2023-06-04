package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

type Repository interface {
	CreateGroup(
		ctx context.Context,
		group *groupmdl.Group,
	) error
	CreateRequest(
		ctx context.Context,
		req *requestmdl.Request,
	) error
	List(
		ctx context.Context,
		filter map[string]interface{},
	) ([]groupmdl.Group, error)
	FindUser(
		ctx context.Context,
		filter map[string]interface{},
	) (*groupmdl.User, error)
	UpdateUser(
		ctx context.Context,
		filter map[string]interface{},
		updatedUser *groupmdl.User,
	) error
	FindRequest(
		ctx context.Context,
		filter map[string]interface{},
	) (*requestmdl.Request, error)
	FindGroup(
		ctx context.Context,
		filter map[string]interface{},
	) (*groupmdl.Group, error)
	UpdateGroup(
		ctx context.Context,
		filter map[string]interface{},
		updatedGroup *groupmdl.Group,
	) error
	DeleteRequest(
		ctx context.Context,
		filter map[string]interface{},
	) error
	FindRequests(
		ctx context.Context,
		filter map[string]interface{},
	) ([]requestmdl.Request, error)
	FindUsers(
		ctx context.Context,
		filter map[string]interface{},
	) ([]groupmdl.User, error)
}

type groupRepository struct {
	groupStore   groupstore.Store
	requestStore requeststore.Store
}

func NewGroupRepository(
	groupStore groupstore.Store,
	requestStore requeststore.Store,
) *groupRepository {
	return &groupRepository{
		groupStore:   groupStore,
		requestStore: requestStore,
	}
}
