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
	ApiKey 	string `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"crated_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	URL 	string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID	`json:"id"`
	CreatedAt time.Time	`json:"crated_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	UserID    uuid.UUID	`json:"user_id"`
	FeedID    uuid.UUID	`json:"feed_id"`
}

func databaseUserToUser(dbUser db.User) User {
	return User {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed db.Feed) Feed {
	return Feed {
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		URL: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []db.Feed) []Feed {
	feeds := []Feed{}
	
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow db.FeedFollow) FeedFollow {
	return FeedFollow {
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}