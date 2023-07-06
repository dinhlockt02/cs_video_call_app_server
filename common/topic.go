package common

const (
	TopicRoomFinished      = "TopicRoomFinished"
	TopicRoomCreated       = "TopicRoomCreated"
	TopicUserUpdateProfile = "TopicUserUpdateProfile"
)

// Defined MarshaledType sent to pubsub

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
