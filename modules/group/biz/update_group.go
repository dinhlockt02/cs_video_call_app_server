package groupbiz

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/pubsub"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UpdateGroupBiz struct {
	groupRepo grouprepo.Repository
	pubsub    pubsub.PubSub
}

func NewUpdateGroupBiz(groupRepo grouprepo.Repository, pubsub pubsub.PubSub) *UpdateGroupBiz {
	return &UpdateGroupBiz{groupRepo: groupRepo, pubsub: pubsub}
}

func (biz *UpdateGroupBiz) Update(ctx context.Context, filter map[string]interface{}, data *groupmdl.Group) error {
	data.Members = nil

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid update data"))
	}

	if err := biz.groupRepo.UpdateGroup(ctx, filter, data); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "can not update group"))
	}

	go func() {
		defer common.Recovery()
		biz.publishEvent(context.Background(), filter)
	}()
	return nil
}

func (biz *UpdateGroupBiz) publishEvent(ctx context.Context, groupFilter map[string]interface{}) {
	group, err := biz.groupRepo.FindGroup(ctx, groupFilter)
	if err != nil {
		log.Error().Err(err).Msg("can not find group")
		return
	}

	if group == nil {
		log.Error().Err(err).Msg("can not group not found")
		return
	}

	groupName := ""
	groupImageURL := ""
	if group.Name != nil {
		groupName = *group.Name
	}

	if group.ImageURL != nil {
		groupImageURL = *group.ImageURL
	}

	marshaledGroup := &common.Group{
		Id:       *group.Id,
		Name:     groupName,
		ImageURL: groupImageURL,
	}

	marshaledData, err := json.Marshal(marshaledGroup)
	if err != nil {
		log.Error().Err(err).Msg("can not marshaled group")
		return
	}

	err = biz.pubsub.Publish(ctx, common.TopicGroupUpdated, string(marshaledData))
	if err != nil {
		log.Error().Err(err).Msg("can not publish event")
	}
}
