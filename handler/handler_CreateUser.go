package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/models"
	"github.com/aadityadike/RSS-Aggregator/response"
	"github.com/google/uuid"
)

func (config *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	user, err := config.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error creating user %v", err))
		return
	}

	response.RespondWithJson(w, 200, models.DatabaseUserToUser(user))
}
