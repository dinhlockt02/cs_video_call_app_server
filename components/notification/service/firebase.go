package notiservice

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4/messaging"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/rs/zerolog/log"
	"time"
)

type firebaseNotificationService struct {
	client *messaging.Client
}

func NewFirebaseNotificationService(client *messaging.Client) firebaseNotificationService {
	return firebaseNotificationService{
		client: client,
	}
}

func (service firebaseNotificationService) Push(ctx context.Context, token []string, notification *notimodel.Notification) error {
	title, body := notification.GetMessage()

	marshaledNotification, err := json.Marshal(notification)
	if err != nil {
		return common.ErrInternal(err)
	}
	zeroDuration := time.Duration(0)

	msg := &messaging.MulticastMessage{
		Tokens: token,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			TTL:      &zeroDuration,
			Notification: &messaging.AndroidNotification{
				Title:    title,
				Body:     body,
				ImageURL: *notification.Prep.Image,
			},
		},
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: *notification.Prep.Image,
		},
		Data: map[string]string{
			"notification": string(marshaledNotification),
		},
	}

	br, err := service.client.SendMulticast(ctx, msg)
	if err != nil {
		return common.ErrInternal(err)
	}
	log.Info().Msgf("%v sent successfully", br.SuccessCount)
	log.Info().Msgf("%v sent failed", br.FailureCount)
	return nil
}
