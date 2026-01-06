package commands

import (
	"context"

	"github.com/lucasgjanot/go-gator-feed/internal/rss"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandAgg(s *runtime.State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	s.Output.PrintFeed(*feed)
	return nil
}