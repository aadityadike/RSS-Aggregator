package Test

import (
	"github.com/aadityadike/RSS-Aggregator/response"
	"net/http"
)

func HandlerResponse(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, 200, struct{}{})
}
