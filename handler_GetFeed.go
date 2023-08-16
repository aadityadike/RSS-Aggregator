package main

import (
	"fmt"
	"net/http"
)

func (config *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feed, err := config.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error with getting Feeds : %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsToFeeds(feed))
}
