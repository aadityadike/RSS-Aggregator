package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithJson is a utility function to send JSON responses to clients over HTTP.
// It takes a ResponseWriter, an HTTP status code, and a payload (data) to be converted to JSON.
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload (data) to JSON format.
	data, err := json.Marshal(payload)
	if err != nil {
		// If there is an error while marshaling the JSON, log the error and return a 500 Internal Server Error.
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	// Set the response header to indicate that the content is in JSON format.
	w.Header().Add("Content-type", "application/json")
	// Write the HTTP status code to the response.
	w.WriteHeader(code)
	// Write the JSON data to the response body.
	w.Write(data)
}
