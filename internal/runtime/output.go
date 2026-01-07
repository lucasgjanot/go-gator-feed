package runtime

import (
	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/rss"
)

type Output interface {
	UserCreated(user database.User)
	UserLoggedIn(username string)
	ResetedDatabase()
	ListUsers(s *State, users []database.User)

	PrintFeed(feed rss.RSSFeed)
	PrintFeeds(feeds []database.GetFeedsWithUserNameRow)
	FeedCreated(feed database.Feed)
	FeedFollowCreated(feedFollow database.CreateFeedFollowRow)
	PrintFeedFollowing(feedFollowing []database.GetFeedFollowsForUserRow)
	
}
