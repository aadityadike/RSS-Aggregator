package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// get port
	portString := os.Getenv("PORT")
	// check port is there or not.
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	fmt.Println("PORT", portString)
	// new Route
	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is starting on %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal((err))
	}
}
