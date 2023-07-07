package common

const (
	TopicRoomFinished      = "TopicRoomFinished"
	TopicRoomCreated       = "TopicRoomCreated"
	TopicUserUpdateProfile = "TopicUserUpdateProfile"
	TopicGroupUpdated      = "TopicGroupUpdated"
	TopicRequestDeleted    = "TopicRequestDeleted"
)

// Defined MarshaledType sent to pubsub

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Group struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}
