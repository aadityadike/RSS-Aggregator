package main

import "net/http"

// handlerReadiness is an HTTP request handler that responds to readiness checks.
// It returns an empty JSON response with a status code of 200 (OK).
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// The respondWithJson function is called to send an HTTP response with JSON data.
	// In this case, an empty struct{}{} is used as the JSON payload.
	// The status code 200 indicates a successful response.
	respondWithJson(w, 200, struct{}{})
}
