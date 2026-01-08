package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
	"github.com/lucasgjanot/go-gator-feed/internal/utils"
)

func CommandUnfollow(s *runtime.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feedURL := cmd.Args[0]
	if err := utils.ValidateURL(feedURL); err != nil {
		return err
	}

	feed, err := s.Database.Feed.GetFeedByUrl(
		context.Background(),
		feedURL,
	)
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrFeedNotFound
		}
		return err
	}

	_, err = s.Database.Feed.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrFeedFollowNotExist
		}
		return err
	}

	s.Output.FeedFollowDeleted(feed, user)
	return nil
}