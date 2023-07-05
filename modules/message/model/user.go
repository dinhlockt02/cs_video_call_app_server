package messagemdl

type User struct {
	Id     string `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Avatar string `bson:"avatar" json:"avatar"`
}
