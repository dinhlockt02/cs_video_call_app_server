package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) CreateEmailAndPasswordUser(ctx context.Context, data *authmodel.RegisterUser) (*authmodel.User, error) {
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return &authmodel.User{
		Id:             id.Hex(),
		Email:          data.Email,
		Password:       data.Password,
		EmailVerified:  data.EmailVerified,
		ProfileUpdated: data.ProfileUpdated,
	}, nil
}

func (s *mongoStore) CreateFirebaseUser(ctx context.Context, data *authmodel.RegisterFirebaseUser) (*authmodel.User, error) {
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return &authmodel.User{
		Id:             id.Hex(),
		Email:          data.Email,
		EmailVerified:  data.EmailVerified,
		ProfileUpdated: data.ProfileUpdated,
	}, nil
}
