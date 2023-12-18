package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/gentcod/RSSAggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (server *Server) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := server.store.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt't create feed follow: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (server *Server) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := server.store.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt't get feed follow: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (server *Server) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(FeedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	err = server.store.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}