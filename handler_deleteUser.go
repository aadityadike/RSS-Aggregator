package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (config *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	err = config.DB.DeleteUser(r.Context(), params.Name)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting user %v", err))
	}

	respondWithJson(w, 200, struct{}{})
}
