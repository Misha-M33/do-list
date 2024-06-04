package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetTokens(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("wrong auth header")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("wrong first part of auth header")
	}

	return vals[1], nil
}
