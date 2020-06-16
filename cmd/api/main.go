package main

import (
	"github.com/alimnastaev/testapi/internal/api/router"

	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func main() {

	r := router.New("1.0.0")

	router := mux.NewRouter()
	router.HandleFunc("/", r.Healthcheck).Methods(http.MethodGet)
	router.HandleFunc("/healthcheck", r.Healthcheck).Methods(http.MethodGet)

	s := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Starting the api")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("There was an error: %v", err)
	}
}
