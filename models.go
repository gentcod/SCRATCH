package main

import (
	"time"

	db "github.com/gentcod/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"crated_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
}

func databaseUserToUser(dbUser db.User) User {
	return User {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}
}