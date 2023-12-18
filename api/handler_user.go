package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/gentcod/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func (server *Server) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := server.store.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt't create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (server *Server) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))
}