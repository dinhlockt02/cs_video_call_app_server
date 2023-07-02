package appcontext

import (
	fbs "github.com/dinhlockt02/cs_video_call_app_server/components/firebase"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	"github.com/dinhlockt02/cs_video_call_app_server/components/mailer"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	"github.com/dinhlockt02/cs_video_call_app_server/components/pubsub"
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
	FirebaseApp() fbs.App
	Notification() notirepo.Repository
	LiveKitService() lksv.LiveKitService
	PubSub() pubsub.PubSub
}

type AppCtx struct {
	mongoClient    *mongo.Client
	tokenProvider  tokenprovider.TokenProvider
	hasher         hasher.Hasher
	rds            *redis.Client
	fa             fbs.App
	mailer         mailer.Mailer
	notification   notirepo.Repository
	livekitService lksv.LiveKitService
	pubsub         pubsub.PubSub
}

func NewAppContext(
	mongoClient *mongo.Client,
	tokenProvider tokenprovider.TokenProvider,
	hasher hasher.Hasher,
	fa fbs.App,
	mailer mailer.Mailer,
	rds *redis.Client,
	notification notirepo.Repository,
	livekitService lksv.LiveKitService,
	pubsub pubsub.PubSub,
) *AppCtx {
	return &AppCtx{
		mongoClient:    mongoClient,
		tokenProvider:  tokenProvider,
		hasher:         hasher,
		fa:             fa,
		mailer:         mailer,
		rds:            rds,
		notification:   notification,
		livekitService: livekitService,
		pubsub:         pubsub,
	}
}

func (a *AppCtx) MongoClient() *mongo.Client {
	return a.mongoClient
}

func (a *AppCtx) TokenProvider() tokenprovider.TokenProvider {
	return a.tokenProvider
}

func (a *AppCtx) Hasher() hasher.Hasher {
	return a.hasher
}

func (a *AppCtx) Redis() *redis.Client {
	return a.rds
}

func (a *AppCtx) FirebaseApp() fbs.App {
	return a.fa
}

func (a *AppCtx) Mailer() mailer.Mailer {
	return a.mailer
}

func (a *AppCtx) Notification() notirepo.Repository {
	return a.notification
}

func (a *AppCtx) LiveKitService() lksv.LiveKitService {
	return a.livekitService
}

func (a *AppCtx) PubSub() pubsub.PubSub {
	return a.pubsub
}
