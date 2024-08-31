package gofiberfcm

import (
	"context"

	"firebase.google.com/go/messaging"
)

// SendNotification sends a push notification to the specified token with the given title and body.
func SendNotification(token, title, body string) (string, error) {
	client := GetClient()
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	// Send the notification
	return client.Send(context.Background(), message)
}
