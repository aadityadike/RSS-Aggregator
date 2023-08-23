package main

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerDeleteFeedFollowing(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := chi.URLParam(r, "feedFollowsID")

	id, err := uuid.Parse(feedFollowID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("You have messed in FeedFollowID: %s", err))
		return
	}

	err = config.DB.DeleteFeedFollowing(r.Context(), database.DeleteFeedFollowingParams{
		ID:     id,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Error in removing following : %s", err))
		return
	}

	respondWithJson(w, 200, struct{}{})
}
