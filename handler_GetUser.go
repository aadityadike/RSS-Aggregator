package main

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/auth"
)

func (config *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error while getting an ApiKey %v", err))
		return
	}

	user, err := config.DB.GetUserByAPIKey(r.Context(), apikey)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while getting user %v", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}