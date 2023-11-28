package main

import (
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/server"
)

func main() {
	server.NewServerSetting()
	//e := echo.New()
	//
	//port := os.Getenv("BACKEND_PORT")
	//if port == "" {
	//	port = "4000"
	//}
	//
	//// Firebaseの初期化
	//opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_CREDENTIALS_JSON")))
	//app, err := firebase.NewApp(context.Background(), nil, opt)
	//if err != nil {
	//	log.Printf("error initializing app: %v", err)
	//}
	//log.Printf("ok initializing app")
	//
	//// Firestoreの初期化
	//client, err := db.NewFirestoreClient(ctx, app)
	//if err != nil {
	//	log.Printf("error initializing firestore: %v", err)
	//}
	//defer func(client *firestore.Client) {
	//	err := client.Close()
	//	if err != nil {
	//		log.Printf("error closing firestore: %v", err)
	//	}
	//}(client)
	//
	//// CORSの設定
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	//AllowOrigins: []string{"http://localhost:80"},
	//	AllowMethods: []string{
	//		http.MethodGet,
	//		http.MethodPut,
	//		http.MethodPost,
	//		http.MethodDelete,
	//	},
	//}))
	//
	//// ルーティングの設定
	//route.InitRoute(e, client)
	//
	//// サーバーの起動
	//log.Printf("connect to http://localhost:%s/ for rest api", port)
	//e.Logger.Fatal(e.Start(":" + port))

}
