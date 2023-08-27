package handler

import (
	"log"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/models"
	"github.com/aadityadike/RSS-Aggregator/response"
)

func (config *ApiConfig) HandlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := config.DB.GetPosts(r.Context(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		log.Printf("Error in Getting posts that user is following: %s", err)
		return
	}

	response.RespondWithJson(w, 200, models.RangeOfPosts(posts))
}
