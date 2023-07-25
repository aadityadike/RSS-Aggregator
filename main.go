// Importing required packages
package main

import (
	"log"
	"net/http"
	"os"

	// Importing the "chi" router library
	"github.com/go-chi/chi"
	// Importing the "cors" middleware for handling CORS (Cross-Origin Resource Sharing)
	"github.com/go-chi/cors"
	// Importing the "godotenv" package to load environment variables from a .env file
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables from .env file (if present)
	godotenv.Load()

	// Get the port number from the environment variables
	portString := os.Getenv("PORT")

	// Check if the PORT is defined or not. If not, the server cannot start, so we log a fatal error and exit the program.
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	// Create a new router instance using "chi"
	router := chi.NewRouter()

	// Configuration for Cross-Origin Resource Sharing (CORS)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},                   // Allow requests from any HTTP or HTTPS origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow specific HTTP methods
		AllowedHeaders:   []string{"*"},                                       // Allow any headers in the request
		ExposedHeaders:   []string{"Link"},                                    // Expose specific headers in the response
		AllowCredentials: false,                                               // Do not allow sending credentials (e.g., cookies) with the requests
		MaxAge:           300,                                                 // Cache preflight request results for 5 minutes
	}))

	// Create a new sub-router for version 1 of the API
	v1Router := chi.NewRouter()

	// Define a route for "/v1/healthz" and associate it with the "handlerReadiness" function
	v1Router.Get("/healthz", handlerReadiness)
	// Define a route for "/v1/err" and associate it with the "handlerError" function
	v1Router.Get("/err", handlerError)

	// Mount the v1Router under the "/v1" prefix in the main router
	router.Mount("/v1", v1Router)

	// Create an HTTP server instance
	srv := &http.Server{
		Handler: router,           // Set the router as the handler for the server
		Addr:    ":" + portString, // Set the address to listen on with the provided port
	}

	// Log a message indicating that the server is starting on the specified port
	log.Printf("Server is starting on %v", portString)

	// Start listening and serving HTTP requests
	err := srv.ListenAndServe()

	/*
	 * This will return err after the server stops listening. The point is that our server will run forever,
	 * but if an error occurs during listening (e.g., the port is already in use), it will be caught by the following if statement.
	 */
	if err != nil {
		log.Fatal((err))
	}
}
