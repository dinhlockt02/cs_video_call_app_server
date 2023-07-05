package messagerepo

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
)

func (repo *MessageRepository) ListMessages(ctx context.Context,
	filter map[string]interface{}) ([]messagemdl.Message, error) {
	return repo.messageStore.List(ctx, filter)
}
