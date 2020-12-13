package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	port = ":8080"
)

func main() {
	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("Starting apiserver on port %s", port)

	log.Fatal(
		http.ListenAndServe(
			port,
			handlers.CORS(
				originsOk,
				headersOk,
				methodsOk,
			)(
				router,
			),
		),
	)
}
