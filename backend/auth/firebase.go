package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	FirebaseApp = app
}
