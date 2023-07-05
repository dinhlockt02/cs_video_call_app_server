package messagerepo

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
)

func (repo *MessageRepository) CreateMessage(ctx context.Context, message *messagemdl.Message) error {
	return repo.messageStore.Create(ctx, message)
}
