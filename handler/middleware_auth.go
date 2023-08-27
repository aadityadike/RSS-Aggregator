package handler

import (
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/internal/auth"
	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/aadityadike/RSS-Aggregator/response"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (config *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)

		if err != nil {
			response.RespondWithError(w, 403, fmt.Sprintf("Error while getting an ApiKey %v", err))
			return
		}

		user, err := config.DB.GetUserByAPIKey(r.Context(), apikey)

		if err != nil {
			response.RespondWithError(w, 400, fmt.Sprintf("Error while getting user %v", err))
			return
		}

		handler(w, r, user)
	}
}
