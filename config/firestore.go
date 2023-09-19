package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func SetupFirestore() *firestore.Client {

	//Using a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		println("err 20", err)
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		println("err 26", err)
		log.Fatalln(err)
	}

	return client
}
