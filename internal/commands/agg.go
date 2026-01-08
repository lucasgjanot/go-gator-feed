package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/rss"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandAgg(s *runtime.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}
	fmt.Println(cmd.Args)
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	s.Output.PrintRequestInterval(timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *runtime.State) error {
	nextFeed, err := s.Database.Feed.GetNextFeedToFetch(context.Background())
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrNoFeedToFetch
		}
		return err
	}
	if _, err = s.Database.Feed.MarkFeedFetched(
		context.Background(),
		nextFeed.ID,
	); err != nil {
		return err
	}

	feedData, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	if feedData == nil {
		return fmt.Errorf("rss.FetchFeed returned nil feed for URL %s", nextFeed.Url)
	}


	feedDataValue := *feedData
	for _, item := range feedDataValue.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		if _, err := s.Database.Post.CreatePost(
			context.Background(),
			database.CreatePostParams{
				Title: item.Title,
				Url: item.Link,
				Description: sql.NullString{
					String: item.Description,
					Valid:  true,
				},
				PublishedAt: publishedAt,
				FeedID: nextFeed.ID,

			},
		); err != nil {
			if runtime.IsExistsError(err) {
				continue
			}
			s.Output.Print(fmt.Sprintf("Couldn't create post: %v", err))
			return err
		}
	}
	s.Output.Print(fmt.Sprintf("Feed %s collected, %v posts found", nextFeed.Name, len(feedData.Channel.Item)))
	return nil
}