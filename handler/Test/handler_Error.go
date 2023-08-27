package Test

import (
	"github.com/aadityadike/RSS-Aggregator/response"
	"net/http"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	response.RespondWithError(w, 400, "An error occurred. Please try again later.")
}
