package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandAddFeed(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	if !strings.Contains(feedUrl, "http") {
		return fmt.Errorf("invalid url")
	}

	user, err := s.Database.User.GetUser(context.Background(), s.Config.GetCurrentUser())
	if err != nil {
		return err
	}

	feed, err := s.Database.Feed.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			Name: feedName,
			Url: feedUrl,
			UserID: user.ID,
		},
	)
	if err != nil {
		return err
	}

	if _, err := s.Database.Feed.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	); err != nil {
		if runtime.IsExistsError(err) {
			return runtime.ErrFeedFollowExists
		}
		return err
	}
	
	s.Output.FeedCreated(feed)
	return nil
}