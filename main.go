package main

import (
	// standard libraries
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	// non-standard libraries
	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

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

	config := apiConfig{
		DB: database.New(connection),
	}

	router := chi.NewRouter()

	// Configuration Cross-Origin Resource Sharing (CORS).
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"}, // Expose specific headers in the response.
		AllowCredentials: false,            // Do not allow sending credentials (e.g., cookies) with requests.
		MaxAge:           300,              // Cache preflight request results for 5 minutes.
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/user", config.middlewareAuth(config.handlerGetUser))
	v1Router.Post("/feed", config.middlewareAuth(config.handlerCreateFeed))
	v1Router.Post("/user", config.handlerCreateUser)
	v1Router.Get("/err", handlerError)
	v1Router.Get("/feed", config.handlerGetFeed)
	v1Router.Delete("/delete", config.handlerDeleteUser)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is starting on %v", portString)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
