package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/idkwhyureadthis/agg-project/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DBUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func DBFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func DBFeedsToFeeds(dbfeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbfeeds {
		feeds = append(feeds, DBFeedToFeed(dbFeed))
	}
	return feeds
}

func DBFeedFollowToFeedFollow(dbFF database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFF.ID,
		CreatedAt: dbFF.CreatedAt,
		UpdatedAt: dbFF.UpdatedAt,
		UserID:    dbFF.UserID,
		FeedID:    dbFF.FeedID,
	}
}

func DBFeedFollowsToFeedFollows(dbfeeds []database.FeedFollow) []FeedFollow {
	feeds := []FeedFollow{}
	for _, dbFeedFollow := range dbfeeds {
		feeds = append(feeds, DBFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feeds
}

func DBPostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func DBPostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, DBPostToPost(dbPost))
	}
	return posts
}
