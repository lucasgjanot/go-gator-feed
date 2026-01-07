package runtime

import (
	"context"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type UsersInterface interface {
	CreateUser(ctx context.Context, name string) (database.User, error)
	GetUser(ctx context.Context, name string) (database.User, error)
	ResetUsers(ctx context.Context) error
	GetUsers(ctx context.Context) ([]database.User, error)
}


type FeedsInterface interface {
	CreateFeed(ctx context.Context, args database.CreateFeedParams) (database.Feed, error)
	GetFeeds(ctx context.Context) ([]database.Feed, error)
	GetFeedsWithUserName(ctx context.Context) ([]database.GetFeedsWithUserNameRow, error)
}

type UserConfig interface {
	SetUser(name string) error
	GetCurrentUser() string
}

type Database struct {
	User UsersInterface
	Feed FeedsInterface
}

type State struct {
	Database Database
	Config  UserConfig
	Output   Output
}