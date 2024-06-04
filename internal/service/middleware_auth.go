package service

import (
	"do-list/src/auth"

	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request)

func JwtAuthMiddleware(handler authHandler, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken, err := auth.GetTokens(r.Header)
		if err != nil {
			SendResponseJSON(w, r, err.Error(), http.StatusForbidden)
			return
		}
		authorized, err := IsAuthorized(authToken, secret)
		if authorized {
			userID, err := ExtractIDFromToken(authToken, secret)
			if err != nil {
				SendResponseJSON(w, r, err.Error(), http.StatusUnauthorized)
				return
			}
			r.Header.Set("id", userID)
			handler(w, r)
			return
		}
		SendResponseJSON(w, r, err.Error(), http.StatusUnauthorized)
	}
}
