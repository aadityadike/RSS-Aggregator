package main

import (
	// standard libraries
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	// non-standard libraries
	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/routers"
	"github.com/joho/godotenv"
)

var DB *database.Queries

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	connection, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}

	DB = database.New(connection)
	routers.Routes(DB)

	go scraper(DB, 10, time.Minute)

	srv := &http.Server{
		Handler: routers.Router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is starting on %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
