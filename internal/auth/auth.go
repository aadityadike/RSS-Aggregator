package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Header is going to look like this ->
// Authorization: "ApiKey" "The Actual Api key"

func GetApiKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")

	if value == "" {
		return "", errors.New("authorization error occurs while getting Api key")
	}

	val := strings.Split(value, " ")

	
}
