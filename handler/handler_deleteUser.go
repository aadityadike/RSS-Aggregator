package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aadityadike/RSS-Aggregator/response"
)

func (config *ApiConfig) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	err = config.DB.DeleteUser(r.Context(), params.Name)

	if err != nil {
		response.RespondWithError(w, 400, fmt.Sprintf("Error deleting user %v", err))
	}

	response.RespondWithJson(w, 200, struct{}{})
}
