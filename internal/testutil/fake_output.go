package testutil

import (
	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

type FakeOutput struct {
	UserCreatedCalled bool
	UserLoggedInCalled bool
	ResetedDatabaseCalled bool

	User              database.User
	LoggedUser  string
}

func (f *FakeOutput) UserCreated(user database.User) {
	f.UserCreatedCalled = true
	f.User = user
}

func (f *FakeOutput) UserLoggedIn(username string) {
	f.UserLoggedInCalled = true
	f.LoggedUser = username
}

func (f *FakeOutput) ResetedDatabase() {
	f.ResetedDatabaseCalled = true
}