package main

import (
	"fmt"
	"net/http"
)

func (config *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := config.DB.GetAllTheUserData(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error with getting Feeds : %v", err))
		return
	}

	respondWithJson(w, 201, databaseUsersToUsers(users))
}
