package main

import (
	"fmt"
	"net/http"

	database "github.com/gentcod/RSSAggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't fetch  posts: %v", err))
	}

	respondWithJson(w, 200, databasePostsToPosts(posts))
}