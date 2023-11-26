package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("server start")

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "4000"
	}
	log.Printf("connect to http://localhost:%s/ for rest api", port)

	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_CREDENTIALS_JSON")))
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
	}
	log.Printf("ok initializing app")

	http.ListenAndServe(":4000", nil)
}
