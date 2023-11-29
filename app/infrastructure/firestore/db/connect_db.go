package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"log"
)

// NewFirestoreClient is
func NewFirestoreClient(ctx context.Context, app *firebase.App) (*firestore.Client, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("error initializing firestore: %v", err)
		return nil, err
	}

	return client, nil
}
