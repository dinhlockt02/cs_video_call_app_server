package mailer

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	senderName  string
	senderEmail string
	apiKey      string
}

func NewSendGridMailer(
	senderName string,
	senderEmail string,
	apiKey string,
) *SendGridMailer {
	return &SendGridMailer{
		senderName:  senderName,
		senderEmail: senderEmail,
		apiKey:      apiKey,
	}
}

func (s *SendGridMailer) Send(subject string, receiverEmail string, receiverName string, htmlContent string) error {
	from := mail.NewEmail(s.senderName, s.senderEmail)
	to := mail.NewEmail(receiverName, receiverEmail)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(s.apiKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
