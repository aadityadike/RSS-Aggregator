package handler

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/models"
	"github.com/aadityadike/RSS-Aggregator/response"
)

func (config *ApiConfig) HandlerGetFeedFollowing(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollowing, err := config.DB.GetFeedFollowing(r.Context(), user.ID)

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error while getting Feed Following: %s", err))
	}

	response.RespondWithJson(w, 200, models.RangeOfFeedFollowing(feedsFollowing))
}
