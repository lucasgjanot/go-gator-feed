package testutil

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

// FakeDatabase simula o comportamento do database/sqlc em memória.
// Não possui efeitos colaterais (SQL, IO, stdout).
type FakeDatabase struct {
	Users       map[string]database.User
	Feeds       map[uuid.UUID]database.Feed
	FeedFollows []database.FeedFollow
}

func NewFakeDatabase() *FakeDatabase {
	return &FakeDatabase{
		Users: make(map[string]database.User),
		Feeds: make(map[uuid.UUID]database.Feed),
	}
}

/*
=====================
USERS
=====================
*/

func (f *FakeDatabase) CreateUser(
	ctx context.Context,
	name string,
) (database.User, error) {

	if _, exists := f.Users[name]; exists {
		return database.User{}, errors.New("user already exists")
	}

	now := time.Now()
	user := database.User{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	f.Users[name] = user
	return user, nil
}

func (f *FakeDatabase) GetUser(
	ctx context.Context,
	name string,
) (database.User, error) {

	user, exists := f.Users[name]
	if !exists {
		return database.User{}, errors.New("user not found")
	}

	return user, nil
}

func (f *FakeDatabase) GetUsers(
	ctx context.Context,
) ([]database.User, error) {

	users := make([]database.User, 0, len(f.Users))
	for _, user := range f.Users {
		users = append(users, user)
	}

	return users, nil
}

func (f *FakeDatabase) ResetUsers(
	ctx context.Context,
) error {

	f.Users = make(map[string]database.User)
	return nil
}

/*
=====================
FEEDS
=====================
*/

func (f *FakeDatabase) CreateFeed(
	ctx context.Context,
	arg database.CreateFeedParams,
) (database.Feed, error) {

	now := time.Now()

	feed := database.Feed{
		ID:        uuid.New(),
		Name:      arg.Name,
		Url:       arg.Url,
		UserID:    arg.UserID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	f.Feeds[feed.ID] = feed
	return feed, nil
}

func (f *FakeDatabase) GetFeedByUrl(
	ctx context.Context,
	url string,
) (database.Feed, error) {

	for _, feed := range f.Feeds {
		if feed.Url == url {
			return feed, nil
		}
	}

	return database.Feed{}, errors.New("feed not found")
}

func (f *FakeDatabase) GetFeeds(
	ctx context.Context,
) ([]database.Feed, error) {

	feeds := make([]database.Feed, 0, len(f.Feeds))
	for _, feed := range f.Feeds {
		feeds = append(feeds, feed)
	}

	return feeds, nil
}

func (f *FakeDatabase) GetFeedsWithUserName(
	ctx context.Context,
) ([]database.GetFeedsWithUserNameRow, error) {

	var result []database.GetFeedsWithUserNameRow

	for _, feed := range f.Feeds {
		var username string
		for _, user := range f.Users {
			if user.ID == feed.UserID {
				username = user.Name
				break
			}
		}

		result = append(result, database.GetFeedsWithUserNameRow{
			ID:        feed.ID,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Username:  username,
		})
	}

	return result, nil
}

/*
=====================
FEED FOLLOWS
=====================
*/

func (f *FakeDatabase) CreateFeedFollow(
	ctx context.Context,
	arg database.CreateFeedFollowParams,
) (database.CreateFeedFollowRow, error) {

	var (
		username string
		userFound bool
	)

	for _, user := range f.Users {
		if user.ID == arg.UserID {
			username = user.Name
			userFound = true
			break
		}
	}

	if !userFound {
		return database.CreateFeedFollowRow{}, errors.New("user not found")
	}

	feed, exists := f.Feeds[arg.FeedID]
	if !exists {
		return database.CreateFeedFollowRow{}, errors.New("feed not found")
	}

	now := time.Now()
	row := database.CreateFeedFollowRow{
		ID:        uuid.New(),
		UserID:    arg.UserID,
		FeedID:    arg.FeedID,
		CreatedAt: now,
		UpdatedAt: now,
		Username:  username,
		FeedName:  feed.Name,
	}

	f.FeedFollows = append(f.FeedFollows, database.FeedFollow{
		ID:        row.ID,
		UserID:    row.UserID,
		FeedID:    row.FeedID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	})

	return row, nil
}

func (f *FakeDatabase) GetFeedFollowsForUser(
	ctx context.Context,
	name string,
) ([]database.GetFeedFollowsForUserRow, error) {

	user, exists := f.Users[name]
	if !exists {
		return nil, errors.New("user not found")
	}

	var result []database.GetFeedFollowsForUserRow

	for _, follow := range f.FeedFollows {
		if follow.UserID != user.ID {
			continue
		}

		feed, exists := f.Feeds[follow.FeedID]
		if !exists {
			continue
		}

		result = append(result, database.GetFeedFollowsForUserRow{
			ID:        follow.ID,
			UserID:    follow.UserID,
			FeedID:    follow.FeedID,
			CreatedAt: follow.CreatedAt,
			UpdatedAt: follow.UpdatedAt,
			FeedName:  feed.Name,
			Username:  user.Name,
		})
	}

	return result, nil
}

func (f *FakeDatabase) DeleteFeedFollow(
	ctx context.Context,
	arg database.DeleteFeedFollowParams,
) (database.FeedFollow, error) {

	for i, follow := range f.FeedFollows {
		if follow.UserID == arg.UserID && follow.FeedID == arg.FeedID {

			// remove do slice (preservando ordem)
			f.FeedFollows = append(
				f.FeedFollows[:i],
				f.FeedFollows[i+1:]...,
			)

			return follow, nil
		}
	}

	return database.FeedFollow{}, errors.New("feed follow not found")
}
