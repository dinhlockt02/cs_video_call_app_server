package notimodel

type NotificationObject struct {
	Id    string                 `json:"id" bson:"id"`
	Name  string                 `bson:"name" bson:"name"`
	Image *string                `bson:"image,omitempty" json:"image,omitempty"`
	Type  NotificationObjectType `json:"type" bson:"type"`
}
