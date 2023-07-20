package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	fbs "github.com/dinhlockt02/cs_video_call_app_server/components/firebase"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	"github.com/dinhlockt02/cs_video_call_app_server/components/mailer"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	notiservice "github.com/dinhlockt02/cs_video_call_app_server/components/notification/service"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	redispubsub "github.com/dinhlockt02/cs_video_call_app_server/components/pubsub/redis"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider/jwt"
	"github.com/dinhlockt02/cs_video_call_app_server/middleware"
	v1 "github.com/dinhlockt02/cs_video_call_app_server/route/v1"
	"github.com/dinhlockt02/cs_video_call_app_server/subscriber"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init() {
	setupLogger()

	var err error

	common.AppDatabase = os.Getenv("MONGO_DB")
	common.AccessTokenExpiry, err = strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY"))

	if err != nil {
		log.Panic().Msg(err.Error())
	}
}

func main() {
	// Get mongo client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := connectMongoDB(ctx)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Get token provider

	tokenProvider := jwt.NewJwtTokenProvider(os.Getenv("SECRET"))

	// Create bcrypt hasher

	bcryptHasher := hasher.NewBcryptHasher()

	// Create mailer

	sendgridMailer := mailer.NewSendGridMailer(
		os.Getenv("SENDGRID_SENDER_NAME"),
		os.Getenv("SENDGRID_SENDER_EMAIL"),
		os.Getenv("SENDGRID_API_KEY"),
	)

	// Create Firebase App
	opt := option.WithCredentialsFile("./service-account-key.json")
	fa, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	app := fbs.NewFirebaseApp(fa)

	// Create redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	// Create pubsub

	pubsub := redispubsub.NewRedisPubSub(redisClient)

	// Create notification service
	firebaseNotificationClient, err := fa.Messaging(context.Background())
	if err != nil {
		log.Panic().Err(err)
	}
	ntsv := notiservice.NewFirebaseNotificationService(firebaseNotificationClient)
	store := notistore.NewMongoStore(client.Database(common.AppDatabase))
	notification := notirepo.NewNotificationService(ntsv, store)

	// Create LiveKit Service

	timeout, err := strconv.ParseUint(os.Getenv("LK_ROOM_TIMEOUT"), 10, 64)
	maximumParticipant, err := strconv.ParseUint(os.Getenv("LK_MAXIMUM_PARTICIPANT"), 10, 64)

	lkservice := lksv.NewLiveKitService(
		os.Getenv("LK_API_KEY"),
		os.Getenv("LK_API_SECRET"),
		os.Getenv("LK_HOST"),
		uint32(timeout),
		uint32(maximumParticipant),
	)

	// Create app context

	appCtx := appcontext.NewAppContext(
		client,
		tokenProvider,
		bcryptHasher,
		app,
		sendgridMailer,
		redisClient,
		notification,
		lkservice,
		pubsub,
	)

	subscriber.Setup(context.Background(), appCtx)

	envport := os.Getenv("SERVER_PORT")
	if envport == "" {
		envport = "8080"
	}
	port := fmt.Sprintf(":%v", envport)

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(middleware.Recover(appCtx))

	v1.InitRoute(r, appCtx)
	r.StaticFile("/delete-your-account", "./templates/delete-account.html")
	r.StaticFile("/reset-password", "./templates/reset-password.html")
	r.StaticFile("/verify/success", "./templates/verify-success.html")
	r.StaticFile("/verify/failure", "./templates/verify-failed.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(port); err != nil {
		log.Panic().Msg(err.Error())
	}
}

func connectMongoDB(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Connect to mongodb successful")
	return client, nil
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.Mode() == gin.DebugMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Logger = log.With().Caller().Logger()
	if gin.Mode() == gin.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
