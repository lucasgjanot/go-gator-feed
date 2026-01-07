package testutil

import (
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/rss"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
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

func (f *FakeOutput) ListUsers(s *runtime.State, users []database.User) {
	for _, user := range users {
		if user.Name == s.Config.GetCurrentUser() {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
}


func (f *FakeOutput) FeedCreated(feed database.Feed) {
	fmt.Println("Feed created successfully:")
	fmt.Printf(" * ID:   %v\n", feed.ID)
	fmt.Printf(" * Name: %v\n", feed.Name)
	fmt.Printf(" * Url: %v\n", feed.Url)
	fmt.Printf(" * UserID: %v\n", feed.UserID)
}
func (f *FakeOutput) PrintFeed(feed rss.RSSFeed){
	fmt.Printf("Feed: %+v\n", feed)
}