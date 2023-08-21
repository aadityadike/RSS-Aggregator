package main

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
)

func (config *apiConfig) handlerGetFeedFollowing(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollowing, err := config.DB.GetFeedFollowing(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while getting Feed Following: %s", err))
	}

	respondWithJson(w, 200, databaseFeedsFollowingToFeedsFollowing(feedsFollowing))
}
