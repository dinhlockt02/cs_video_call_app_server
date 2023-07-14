package callrepo

import (
	"context"
)

func (repo *CallRepository) RemoveOwner(ctx context.Context, filter map[string]interface{}, owner string) error {
	return repo.callStore.RemoveOwner(ctx, filter, owner)
}
