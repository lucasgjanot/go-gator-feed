package testutil

import (
	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/rss"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

type FakeOutput struct {
	// Flags de chamada
	UserCreatedCalled        bool
	UserLoggedInCalled       bool
	ResetedDatabaseCalled    bool
	ListUsersCalled          bool
	FeedCreatedCalled        bool
	PrintFeedCalled          bool
	PrintFeedsCalled         bool
	FeedFollowCreatedCalled  bool
	PrintFeedFollowingCalled bool

	// Dados capturados
	CreatedUser        database.User
	LoggedUser         string
	Users              []database.User
	ListUsersState     *runtime.State

	CreatedFeed        database.Feed
	PrintedFeed        rss.RSSFeed
	PrintedFeeds       []database.GetFeedsWithUserNameRow

	CreatedFeedFollow  database.CreateFeedFollowRow
	FeedFollowing      []database.GetFeedFollowsForUserRow
}

func (f *FakeOutput) UserCreated(user database.User) {
	f.UserCreatedCalled = true
	f.CreatedUser = user
}

func (f *FakeOutput) UserLoggedIn(username string) {
	f.UserLoggedInCalled = true
	f.LoggedUser = username
}

func (f *FakeOutput) ResetedDatabase() {
	f.ResetedDatabaseCalled = true
}

func (f *FakeOutput) ListUsers(s *runtime.State, users []database.User) {
	f.ListUsersCalled = true
	f.ListUsersState = s
	f.Users = users
}

func (f *FakeOutput) FeedCreated(feed database.Feed) {
	f.FeedCreatedCalled = true
	f.CreatedFeed = feed
}

func (f *FakeOutput) PrintFeed(feed rss.RSSFeed) {
	f.PrintFeedCalled = true
	f.PrintedFeed = feed
}

func (f *FakeOutput) PrintFeeds(feeds []database.GetFeedsWithUserNameRow) {
	f.PrintFeedsCalled = true
	f.PrintedFeeds = feeds
}

func (f *FakeOutput) FeedFollowCreated(feedFollow database.CreateFeedFollowRow) {
	f.FeedFollowCreatedCalled = true
	f.CreatedFeedFollow = feedFollow
}

func (f *FakeOutput) PrintFeedFollowing(feedFollowing []database.GetFeedFollowsForUserRow) {
	f.PrintFeedFollowingCalled = true
	f.FeedFollowing = feedFollowing
}
