package handler

import (
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/models"
	"github.com/aadityadike/RSS-Aggregator/response"
)

func (config *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	response.RespondWithJson(w, 201, models.DatabaseUserToUser(user))
}
