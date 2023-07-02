package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) CreateEmailAndPasswordUser(ctx context.Context, data *authmodel.RegisterUser) (*authmodel.User, error) {
	log.Debug().Msg("inserting a firebase user")
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return nil, errors.Wrap(err, "cant not insert user")
	}
	id := result.InsertedID.(primitive.ObjectID)
	log.Debug().Str("user_id", id.Hex()).Msg("inserted a firebase user")
	return &authmodel.User{
		Id:             id.Hex(),
		Email:          data.Email,
		Password:       data.Password,
		EmailVerified:  data.EmailVerified,
		ProfileUpdated: data.ProfileUpdated,
	}, nil
}

func (s *mongoStore) CreateFirebaseUser(ctx context.Context, data *authmodel.RegisterFirebaseUser) (*authmodel.User, error) {
	log.Debug().Msg("inserting a firebase user")
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return nil, errors.Wrap(err, "can not insert user")
	}
	id := result.InsertedID.(primitive.ObjectID)
	log.Debug().Str("user_id", id.Hex()).Msg("inserted a firebase user")
	return &authmodel.User{
		Id:             id.Hex(),
		Email:          data.Email,
		EmailVerified:  data.EmailVerified,
		ProfileUpdated: data.ProfileUpdated,
	}, nil
}
