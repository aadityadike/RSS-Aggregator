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

func (config *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	feed, err := config.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error creating feed %v", err))
		return
	}

	response.RespondWithJson(w, 201, models.DatabaseFeedToFeed(feed))
}
