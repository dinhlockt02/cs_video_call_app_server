package appcontext

import (
	fbs "github.com/dinhlockt02/cs_video_call_app_server/components/firebase"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	"github.com/dinhlockt02/cs_video_call_app_server/components/mailer"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	MongoClient() *mongo.Client
	TokenProvider() tokenprovider.TokenProvider
	Hasher() hasher.Hasher
	Mailer() mailer.Mailer
	Redis() *redis.Client
	FirebaseApp() fbs.FirebaseApp
}

type appContext struct {
	mongoClient   *mongo.Client
	tokenProvider tokenprovider.TokenProvider
	hasher        hasher.Hasher
	rds           *redis.Client
	fa            fbs.FirebaseApp
	mailer        mailer.Mailer
}

func NewAppContext(
	mongoClient *mongo.Client,
	tokenProvider tokenprovider.TokenProvider,
	hasher hasher.Hasher,
	fa fbs.FirebaseApp,
	mailer mailer.Mailer,
	rds *redis.Client,
) *appContext {
	return &appContext{
		mongoClient:   mongoClient,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		fa:            fa,
		mailer:        mailer,
		rds:           rds,
	}
}

func (a *appContext) MongoClient() *mongo.Client {
	return a.mongoClient
}

func (a *appContext) TokenProvider() tokenprovider.TokenProvider {
	return a.tokenProvider
}

func (a *appContext) Hasher() hasher.Hasher {
	return a.hasher
}

func (a *appContext) Redis() *redis.Client {
	return a.rds
}

func (a *appContext) FirebaseApp() fbs.FirebaseApp {
	return a.fa
}

func (a *appContext) Mailer() mailer.Mailer {
	return a.mailer
}
