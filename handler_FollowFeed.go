package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	fmt.Print(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	feedFollowConnection, err := config.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds that you want %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowsToFollows(feedFollowConnection))
}
