package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (config *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string `json:"apiKey"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	user, err := config.DB.GetUserByAPIKey(r.Context(), params.ApiKey)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}
