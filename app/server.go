package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func main() {
	log.Printf("connect to http://localhost:%s/ for rest api", "4000")

	opt := option.WithCredentialsFile("firebase/serviceAccountKey.json")
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		//return nil, fmt.Errorf("error initializing app: %v", err)
		log.Printf("error initializing app: %v", err)
	}
	log.Printf("ok initializing app")

	http.ListenAndServe(":4000", nil)
}
