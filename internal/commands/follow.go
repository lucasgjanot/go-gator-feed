package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandFollow(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]
	fmt.Println(feedUrl)
	if !strings.Contains(feedUrl, "http") {
		return fmt.Errorf("invalid url")
	}

	user, err := s.Database.User.GetUser(
		context.Background(),
		s.Config.GetCurrentUser(),
	)
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrUserNotFound
		}
		return err
	}

	feed, err := s.Database.Feed.GetFeedByUrl(
		context.Background(),
		feedUrl,
	)
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrFeedNotFound
		}
		return err
	}

	feedFollow, err := s.Database.Feed.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		if runtime.IsExistsError(err) {
			return runtime.ErrFeedFollowExists
		}
		return err
	}
	s.Output.FeedFollowCreated(feedFollow)
	return nil
}