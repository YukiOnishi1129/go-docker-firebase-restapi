package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("connect to http://localhost:%s/ for rest api", "4000")
	http.ListenAndServe(":4000", nil)
}
