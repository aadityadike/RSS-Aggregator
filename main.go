package main

import (
	// standard libraries
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	// non-standard libraries
	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// apiConfig stores the configuration for the API.
type apiConfig struct {
	DB *database.Queries // Instance of the database queries.
}

func main() {
	// Load environment variables from .env file (if present).
	godotenv.Load()

	// Get the port number from the environment variables.
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	// Get the database connection URL from the environment variables.
	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	// Open a connection to the PostgreSQL database.
	connection, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}

	// Create a new instance of the API configuration.
	apiconf := apiConfig{
		DB: database.New(connection),
	}

	// Create a new router instance using "chi".
	router := chi.NewRouter()

	// Configure Cross-Origin Resource Sharing (CORS) settings.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},                   // Allow requests from any HTTP or HTTPS origin.
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow specific HTTP methods.
		AllowedHeaders:   []string{"*"},                                       // Allow any headers in the request.
		ExposedHeaders:   []string{"Link"},                                    // Expose specific headers in the response.
		AllowCredentials: false,                                               // Do not allow sending credentials (e.g., cookies) with requests.
		MaxAge:           300,                                                 // Cache preflight request results for 5 minutes.
	}))

	// Create a sub-router for version 1 of the API.
	v1Router := chi.NewRouter()

	// Define routes and associate them with handler functions.
	v1Router.Get("/healthz", handlerReadiness)         // Route for checking API readiness.
	v1Router.Get("/err", handlerError)                 // Route for triggering an error.
	v1Router.Post("/users", apiconf.handlerCreateUser) // Route for creating users.

	// Mount the v1Router under the "/v1" prefix in the main router.
	router.Mount("/v1", v1Router)

	// Create an HTTP server instance.
	srv := &http.Server{
		Handler: router,           // Set the router as the handler for the server.
		Addr:    ":" + portString, // Set the address to listen on with the provided port.
	}

	// Log a message indicating that the server is starting on the specified port.
	log.Printf("Server is starting on %v", portString)

	// Start listening and serving HTTP requests.
	err = srv.ListenAndServe()

	/*
	 * If an error occurs during listening (e.g., the port is already in use),
	 * it will be caught by the following if statement.
	 */
	if err != nil {
		log.Fatal(err)
	}
}
