package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKet extracts an API Key from the header of an HTTP request

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authentication header key")
	}
	return vals[1], nil
}