package meetingmodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type MeetingStatus string

const (
	OnGoing MeetingStatus = "on-going"
	Ended                 = "ended"
)

type Meeting struct {
	common.MongoId `json:",inline" bson:",inline"`
	Title          string        `bson:"title" json:"title"`
	Description    string        `json:"description" bson:"description"`
	TimeStart      *time.Time    `bson:"time_start" json:"time_start"`
	TimeEnd        *time.Time    `json:"time_end,omitempty" bson:"time_end,omitempty"`
	Participants   []string      `bson:"participants,omitempty" json:"participants,omitempty"`
	GroupId        string        `json:"group" bson:"group"`
	Status         MeetingStatus `bson:"status" json:"status"`
}

func (Meeting) CollectionName() string {
	return common.MeetingCollectionName
}

var (
	ErrMeetingEnded    = errors.New("meeting has ended")
	ErrMeetingNotFound = errors.New("meeting not found")
)
