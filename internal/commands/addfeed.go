package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
	"github.com/lucasgjanot/go-gator-feed/internal/utils"
)

func CommandAddFeed(s *runtime.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	if err := utils.ValidateURL(feedURL); err != nil {
		return err
	}

	feed, err := s.Database.Feed.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			Name: feedName,
			Url: feedURL,
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