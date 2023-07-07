package main

import (
	"net/http"
	"strings"

	"github.com/krishnanshagarwal112/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig)middlewareAuth(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationString := r.Header.Get("Authorization")
		apiKey := strings.TrimPrefix(authorizationString, "ApiKey ")

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "no user found")
			return
		}
		next(w,r,user)
	}
}
