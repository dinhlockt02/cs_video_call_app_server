package userstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindById(ctx context.Context, userId string) (*usermodel.User, error) {

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	result := s.database.
		Collection(usermodel.User{}.CollectionName()).
		FindOne(ctx, bson.D{{"_id", id}})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	var user = new(usermodel.User)
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *mongoStore) Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error) {
	result := s.database.
		Collection(usermodel.User{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	var user = new(usermodel.User)
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
