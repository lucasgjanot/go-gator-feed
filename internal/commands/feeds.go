package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandFeeds(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.Database.Feed.GetFeedsWithUserName(context.Background())
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrNoFeed
		}
	}

	s.Output.PrintFeeds(feeds)
	return nil
}