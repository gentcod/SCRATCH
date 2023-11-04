package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	database "github.com/gentcod/RSSAggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// type parameters struct {
	// 	Limit int	`json:"limit"`
	// }
	// decoder := json.NewDecoder(r.Body)

	// params := parameters{}

	// err := decoder.Decode(&params)
	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	// 	return
	// }

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't fetch  posts: %v", err))
	}

	respondWithJson(w, 200, databasePostsToPosts(posts))
}