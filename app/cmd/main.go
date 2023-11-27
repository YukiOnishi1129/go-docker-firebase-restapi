package main

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("server start")

	e := echo.New()

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
