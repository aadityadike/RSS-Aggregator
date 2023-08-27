package handler

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/models"
	"github.com/aadityadike/RSS-Aggregator/response"
)

func (config *ApiConfig) HandlerGetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := config.DB.GetAllTheUserData(r.Context())
	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error with getting Feeds : %v", err))
		return
	}

	response.RespondWithJson(w, 201, models.RangeOfUsers(users))
}
