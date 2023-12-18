package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/gentcod/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func (server *Server) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url string `jsom:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := server.store.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt't create feed: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedToFeed(feed))
}

func (server *Server) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := server.store.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt't get feed: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsToFeeds(feeds))
}