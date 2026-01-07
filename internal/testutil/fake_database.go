package testutil

import (
	"context"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type FakeDatabase struct {
	Users map[string]database.User
	Feeds map[string]database.Feed
	FeedsWithUsername map[string]database.GetFeedsWithUserNameRow

	CreateUserErr error
	GetUserErr    error
	CreateFeedErr error
}

func NewFakeDatabase() *FakeDatabase {
	return &FakeDatabase{
		Users: make(map[string]database.User),
		Feeds: make(map[string]database.Feed),
		FeedsWithUsername: make(map[string]database.GetFeedsWithUserNameRow),
	}
}

/* ================= USERS ================= */

func (f *FakeDatabase) CreateUser(
	ctx context.Context,
	name string,
) (database.User, error) {
	if f.CreateUserErr != nil {
		return database.User{}, f.CreateUserErr
	}

	user := database.User{Name: name}
	f.Users[name] = user
	return user, nil
}

func (f *FakeDatabase) GetUser(
	ctx context.Context,
	name string,
) (database.User, error) {
	if f.GetUserErr != nil {
		return database.User{}, f.GetUserErr
	}

	user, ok := f.Users[name]
	if !ok {
		return database.User{}, nil
	}

	return user, nil
}

func (f *FakeDatabase) ResetUsers(ctx context.Context) error {
	f.Users = make(map[string]database.User)
	return nil
}

func (f *FakeDatabase) GetUsers(ctx context.Context) ([]database.User, error) {
	users := make([]database.User, 0, len(f.Users))
	for _, u := range f.Users {
		users = append(users, u)
	}
	return users, nil
}

/* ================= FEEDS ================= */

func (f *FakeDatabase) CreateFeed(
	ctx context.Context,
	args database.CreateFeedParams,
) (database.Feed, error) {
	if f.CreateFeedErr != nil {
		return database.Feed{}, f.CreateFeedErr
	}

	feed := database.Feed{
		Name:   args.Name,
		Url:    args.Url,
		UserID: args.UserID,
	}

	f.Feeds[args.Url] = feed
	return feed, nil
}

func (f *FakeDatabase) GetFeeds(ctx context.Context) ([]database.Feed, error) {
	feeds := make([]database.Feed, 0, len(f.Feeds))
	for _, feed := range f.Feeds {
		feeds = append(feeds, feed)
	}
	return feeds, nil

}

func (f *FakeDatabase) GetFeedsWithUserName(ctx context.Context) ([]database.GetFeedsWithUserNameRow, error) {
	feeds := make([]database.GetFeedsWithUserNameRow, 0, len(f.FeedsWithUsername))
	for _, feed := range f.FeedsWithUsername {
		feeds = append(feeds, feed)
	}
	return feeds, nil
}