package grouprepo

import (
	"context"
)

func (repo *GroupRepository) DeleteOne(
	ctx context.Context,
	filter map[string]interface{},
) error {
	return repo.groupStore.DeleteOne(ctx, filter)
}
