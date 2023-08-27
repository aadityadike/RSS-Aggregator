package handler

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/response"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (config *ApiConfig) HandlerDeleteFeedFollowing(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := chi.URLParam(r, "feedFollowsID")

	id, err := uuid.Parse(feedFollowID)
	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("You have messed in FeedFollowID: %s", err))
		return
	}

	err = config.DB.DeleteFeedFollowing(r.Context(), database.DeleteFeedFollowingParams{
		ID:     id,
		UserID: user.ID,
	})

	if err != nil {
		response.RespondWithError(w, 404, fmt.Sprintf("Error in removing following : %s", err))
		return
	}

	response.RespondWithJson(w, 200, struct{}{})
}
