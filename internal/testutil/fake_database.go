package testutil

import (
	"context"
	"database/sql"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type FakeDatabase struct {
	Users map[string]database.User

	CreateUserErr error
	GetUserErr    error
}

func NewFakeDatabase() *FakeDatabase {
	return &FakeDatabase{
		Users: make(map[string]database.User),
	}
}

func (f *FakeDatabase) GetUser(ctx context.Context, name string) (database.User, error) {
	if f.GetUserErr != nil {
		return database.User{}, f.GetUserErr
	}

	u, ok := f.Users[name]
	if !ok {
		return database.User{}, sql.ErrNoRows
	}
	return u, nil
}

func (f *FakeDatabase) CreateUser(ctx context.Context, name string) (database.User, error) {
	if f.CreateUserErr != nil {
		return database.User{}, f.CreateUserErr
	}

	u := database.User{Name: name}
	f.Users[name] = u
	return u, nil
}

func (f *FakeDatabase) ResetUsers(ctx context.Context) error {
	f.Users = map[string]database.User{}
	return nil
}