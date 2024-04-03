package fcm

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"log"
	"fmt"
)

func sendToTopic(ctx context.Context, client *messaging.Client) {
	topic := "highScores"

	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Topic: topic,
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
