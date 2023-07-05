package messagerepo

import (
	"context"
)

func (repo *MessageRepository) DeleteMessage(ctx context.Context, filter map[string]interface{}) error {
	return repo.messageStore.DeleteOne(ctx, filter)
}
