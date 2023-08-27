package handler

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/models"
	"github.com/aadityadike/RSS-Aggregator/response"
)

func (config *ApiConfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feed, err := config.DB.GetFeeds(r.Context())
	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error with getting Feeds : %v", err))
		return
	}

	response.RespondWithJson(w, 201, models.RangeOfFeeds(feed))
}
