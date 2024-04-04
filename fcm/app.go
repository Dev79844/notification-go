package fcm

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/Dev79844/email-fcm-package/config"
	"google.golang.org/api/option"
	"firebase.google.com/go/messaging"
)

type FCMMessenger struct{
	Config config.FCMSenderConfig
	Message string
	App		*firebase.App
	Client 	*messaging.Client
}

func NewFCMMessenger(config config.FCMSenderConfig) (*FCMMessenger, error) {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &FCMMessenger{
		Config: config,
		App: app,
		Client: client,
	}, nil
}

func (m *FCMMessenger) SendMessage() error {
	response, err := m.Client.Send(context.Background(), &messaging.Message{
		Data: map[string]string{
			"message": m.Message,
		},
		Topic: m.Config.Topic,
	})
	if err != nil {
		return err
	}
	
	fmt.Println("Successfully sent message:", response)
	return nil
}


