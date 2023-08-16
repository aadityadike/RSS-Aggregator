package main

import (
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
)

func (config *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 201, databaseUserToUser(user))
}
