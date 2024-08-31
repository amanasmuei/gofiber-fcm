package gofiberfcm

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client

// Init initializes Firebase and sets up the FCM client.
func Init() error {
	ctx := context.Background()
	databaseURL := os.Getenv("DatabaseUrl")
	conf := &firebase.Config{ProjectID: "simi-digital", DatabaseURL: databaseURL}

	app, err := firebase.NewApp(ctx, conf, option.WithCredentialsFile("simi-digital-firebase-adminsdk.json"))
	if err != nil {
		return err
	}

	fcmClient, err = app.Messaging(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetClient returns the FCM client.
func GetClient() *messaging.Client {
	return fcmClient
}
