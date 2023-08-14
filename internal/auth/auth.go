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
		return "", errors.New("authentication error occurs while getting Api key")
	}

	val := strings.Split(value, " ")

	if len(val) != 2 {
		return "", errors.New("authentication values are not correct")
	}

	if val[0] != "ApiKey" {
		return "", errors.New("header is not correct please specify correct headers")
	}

	return val[1], nil

}
