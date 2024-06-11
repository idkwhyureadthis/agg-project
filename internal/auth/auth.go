package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API key from header of request
// Example:
// Authorization: ApiKey {key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth data found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("wrong format header met")
	}
	return vals[1], nil
}
