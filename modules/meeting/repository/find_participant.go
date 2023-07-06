package meetingrepo

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *meetingRepository) FindParticipant(ctx context.Context,
	filter map[string]interface{}) (*meetingmodel.Participant, error) {
	log.Debug().Any("filter", filter).Msg("find participant")
	user, err := r.userStore.Find(ctx, filter)

	if err != nil {
		return nil, errors.Wrap(err, "can not find participant")
	}
	if user == nil {
		return nil, nil
	}

	return &meetingmodel.Participant{
		Id:     *user.Id,
		Name:   user.Name,
		Avatar: user.Avatar,
	}, nil
}
