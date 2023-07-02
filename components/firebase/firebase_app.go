package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/pkg/errors"
)

type App interface {
	VerifyToken(ctx context.Context, idToken string) (*string, error)
	ExtractEmailFromUID(ctx context.Context, uid string) (*string, error)
}

type AppImpl struct {
	app *firebase.App
}

func NewFirebaseApp(app *firebase.App) *AppImpl {
	return &AppImpl{app: app}
}

func (fa *AppImpl) VerifyToken(ctx context.Context, idToken string) (*string, error) {
	client, err := fa.app.Auth(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can not get auth client")
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, errors.Wrap(err, "can not verify id token")
	}
	return &token.UID, nil
}

func (fa *AppImpl) ExtractEmailFromUID(ctx context.Context, uid string) (*string, error) {
	client, err := fa.app.Auth(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can not get auth client")
	}

	firebaseUser, err := client.GetUser(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "can not get user from uid")
	}
	return &firebaseUser.Email, nil
}
