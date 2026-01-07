package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)


func CommandFollowing(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	followingFeeds, err := s.Database.Feed.GetFeedFollowsForUser(
		context.Background(),
		s.Config.GetCurrentUser(),
	)
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrFeedFollowNotFound
		}
	}

	if len(followingFeeds) == 0 {
		return runtime.ErrFeedFollowNotFound
	}

	s.Output.PrintFeedFollowing(followingFeeds)
	return nil

}