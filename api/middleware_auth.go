package api

import (
	"fmt"
	"net/http"

	"github.com/gentcod/RSSAggregator/internal/auth"
	database "github.com/gentcod/RSSAggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (server *Server) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		}

		user, err := server.store.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		} 
		handler(w, r, user)
	}

}