package fcm

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/Dev79844/email-fcm-package/config"
	"google.golang.org/api/option"
	"firebase.google.com/go/messaging"
)

type FCMMessenger struct{
	config config.FCMSenderConfig
	message string
	app		*firebase.App
	client 	*messaging.Client
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
		config: config,
		app: app,
		client: client,
	}, nil
}

