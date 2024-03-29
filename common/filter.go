package common

import "go.mongodb.org/mongo-driver/bson/primitive"

// AddIdFilter is an method that will add id filter to provied filter
//
// It required id as string, and has the mongoid format.
//
// It will return [ErrInvalidRequest] if the id is not has the right format.
func AddIdFilter(filter map[string]interface{}, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidRequest(err)
	}
	filter["_id"] = _id
	return nil
}

func GetIdFilter(id string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := AddIdFilter(m, id)
	return m, err
}

func MustGetIdFilter(id string) map[string]interface{} {
	m := make(map[string]interface{})
	err := AddIdFilter(m, id)
	if err != nil {
		panic(err)
	}
	return m
}
