package runtime

import (
	"context"
	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type Database interface {
	CreateUser(ctx context.Context, name string) (database.User, error)
	GetUser(ctx context.Context, name string) (database.User, error)
	ResetUsers(ctx context.Context) error
}

type UserConfig interface {
	SetUser(name string) error
	GetCurrentUser() string
}

type State struct {
	Database Database
	Config  UserConfig
	Output   Output
}