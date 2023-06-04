package meetingmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type UpdateMeeting struct {
	Title        *string       `bson:"title,omitempty" json:"title,omitempty"`
	Description  *string       `json:"description,omitempty" bson:"description,omitempty"`
	TimeEnd      *time.Time    `json:"time_end,omitempty" bson:"time_end,omitempty"`
	Participants []string      `bson:"participants,omitempty" json:"participants,omitempty"`
	Status       MeetingStatus `bson:"status" json:"status"`
}

func (UpdateMeeting) CollectionName() string {
	return common.MeetingCollectionName
}
