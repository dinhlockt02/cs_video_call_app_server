package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	fbs "github.com/dinhlockt02/cs_video_call_app_server/components/firebase"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider/jwt"
	"github.com/dinhlockt02/cs_video_call_app_server/middleware"
	v1 "github.com/dinhlockt02/cs_video_call_app_server/route/v1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	// Create Firebase App
	opt := option.WithCredentialsFile("./service-account-key.json")
	fa, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	app := fbs.NewFirebaseApp(fa)

	// Create app context

	appCtx := appcontext.NewAppContext(client, tokenProvider, bcryptHasher, app)

	envport := os.Getenv("SERVER_PORT")
	if envport == "" {
		envport = "8080"
	}
	port := fmt.Sprintf(":%v", envport)

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(middleware.Recover(appCtx))

	v1.InitRoute(r, appCtx)

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
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.Mode() == gin.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
