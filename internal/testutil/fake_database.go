package testutil

import (
	"context"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type FakeDatabase struct {
	Users map[string]database.User
	Feeds map[string]database.Feed

	CreateUserErr error
	GetUserErr    error
	CreateFeedErr error
}

func NewFakeDatabase() *FakeDatabase {
	return &FakeDatabase{
		Users: make(map[string]database.User),
		Feeds: make(map[string]database.Feed),
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
