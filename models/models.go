package models

import (
	"time"

	"github.com/aadityadike/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

// DEFINING OWR OWN TYPES SO THAT WE CAN MODIFY ACCORDINGLY.
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

// ASSIGNING TYPE TO OUR OWN TYPES.
func DatabaseFeedFollowsToFollows(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func DatabasePostsToPosts(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		FeedID:      dbPost.FeedID,
		PublishedAt: dbPost.PublishedAt,
	}
}

// FUNCTIONS FOR QUERIES THAT ARE RETURNING MULTIPLE OBJECTS.
func RangeOfFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}

	for _, feed := range dbFeed {
		feeds = append(feeds, DatabaseFeedToFeed(feed))
	}
	return feeds
}

func RangeOfFeedFollowing(dbFeedFollowing []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, each := range dbFeedFollowing {
		feedFollows = append(feedFollows, DatabaseFeedFollowsToFollows(each))
	}
	return feedFollows
}

func RangeOfUsers(dbGetUsers []database.User) []User {
	allUsers := []User{}

	for _, each := range dbGetUsers {
		allUsers = append(allUsers, DatabaseUserToUser(each))
	}
	return allUsers
}

func RangeOfPosts(dbPost []database.Post) []Post {
	posts := []Post{}

	for _, each := range dbPost {
		posts = append(posts, DatabasePostsToPosts(each))
	}
	return posts
}
