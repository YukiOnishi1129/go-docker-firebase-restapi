package main

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/server/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("server start")

	ctx := context.Background()

	e := echo.New()

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "4000"
	}
	log.Printf("connect to http://localhost:%s/ for rest api", port)

	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_CREDENTIALS_JSON")))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
	}
	log.Printf("ok initializing app")

	_, err = app.Firestore(ctx)
	if err != nil {
		log.Printf("error initializing firestore: %v", err)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://localhost:80"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	route.InitRoute(e)

	//vi := e.Group("/api/v1")
	//
	//vi.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World")
	//})
	e.Logger.Fatal(e.Start(":" + port))
}
