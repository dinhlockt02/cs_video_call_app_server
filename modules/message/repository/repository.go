package messagerepo

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	messagestore "github.com/dinhlockt02/cs_video_call_app_server/modules/message/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
)

type Repository interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*messagemdl.User, error)
	CreateMessage(ctx context.Context, message *messagemdl.Message) error
	DeleteMessage(ctx context.Context, filter map[string]interface{}) error
	ListMessages(ctx context.Context, filter map[string]interface{}) ([]messagemdl.Message, error)
}

type MessageRepository struct {
	messageStore messagestore.Store
	userStore    userstore.Store
}

func NewMessageRepository(
	messageStore messagestore.Store,
	userStore userstore.Store,
) *MessageRepository {
	return &MessageRepository{
		messageStore: messageStore,
		userStore:    userStore,
	}
}
