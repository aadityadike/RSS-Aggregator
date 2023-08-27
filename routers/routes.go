package routers

import (
	"github.com/aadityadike/RSS-Aggregator/handler"
	"github.com/aadityadike/RSS-Aggregator/handler/Test"
	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var Router = chi.NewRouter()

func Routes(db *database.Queries) {
	config := handler.ApiConfig{
		DB: db,
	}

	// Configuration Cross-Origin Resource Sharing (CORS).
	Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"}, // Expose specific headers in the response.
		AllowCredentials: false,            // Do not allow sending credentials (e.g., cookies) with requests.
		MaxAge:           300,              // Cache preflight request results for 5 minutes.
	}))

	v1Router := chi.NewRouter()
	Router.Mount("/v1", v1Router)

	/* GET REQUEST */
	// TESTS
	v1Router.Get("/valid", Test.HandlerError)
	v1Router.Get("/err", Test.HandlerResponse)

	v1Router.Get("/user", config.MiddlewareAuth(config.HandlerGetUser))
	v1Router.Get("/feed", config.HandlerGetFeed)
	v1Router.Get("/feedFollows", config.MiddlewareAuth(config.HandlerGetFeedFollowing))
	v1Router.Get("/posts", config.MiddlewareAuth(config.HandlerGetPosts))

	// VALIDATING & FOR DEVELOPMENT PURPOSE (CHECK WHETHER DATABASE CONTAINS USERS OR NOT).
	v1Router.Get("/users", config.HandlerGetUsers)

	/* POST REQUEST */
	v1Router.Post("/feed", config.MiddlewareAuth(config.HandlerCreateFeed))
	v1Router.Post("/user", config.HandlerCreateUser)
	v1Router.Post("/feedFollows", config.MiddlewareAuth(config.HandlerFollowFeed))

	/* DELETE REQUEST */
	v1Router.Delete("/feedFollows/{feedFollowsID}", config.MiddlewareAuth(config.HandlerDeleteFeedFollowing))
	v1Router.Delete("/delete", config.HandlerDeleteUser)
}
