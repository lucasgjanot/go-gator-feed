package cli

import (
	"fmt"
	"time"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/rss"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

type CLIOutput struct{}

func (CLIOutput) UserCreated(user database.User) {
	fmt.Println("User created successfully:")
	fmt.Printf(" * ID:   %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
	fmt.Printf(" * Created:       %v\n", user.CreatedAt)
	fmt.Printf(" * Updated:       %v\n", user.UpdatedAt)
}


func (CLIOutput) UserLoggedIn(username string) {
	fmt.Printf("Logged in as %s\n", username)
}

func (CLIOutput) ResetedDatabase() {
	fmt.Println("Database reset successfully!")
}

func (CLIOutput) ListUsers(s * runtime.State, users []database.User) {
	for _, user := range users {
		if user.Name == s.Config.GetCurrentUser() {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
}



func (CLIOutput) FeedCreated(feed database.Feed) {
	fmt.Println("Feed created successfully:")
	fmt.Printf(" * ID:   %v\n", feed.ID)
	fmt.Printf(" * Name: %v\n", feed.Name)
	fmt.Printf(" * Url: %v\n", feed.Url)
	fmt.Printf(" * UserID: %v\n", feed.UserID)
	fmt.Printf(" * Created:       %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf(" * Last Fetched:       %v\n", feed.LastFetchedAt)
}

func (CLIOutput) PrintFeed(feed rss.RSSFeed) {
	fmt.Printf("Feed: %+v\n", feed)
}

func (CLIOutput) PrintFeedItems(feed rss.RSSFeed) {
	for _, item := range feed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	
}

func (CLIOutput) PrintFeeds(feeds []database.GetFeedsWithUserNameRow) {
	for _, feed := range feeds {
		fmt.Printf("- Name: %s, Url: %s, Username: %s\n", feed.Name, feed.Url, feed.Username)
	}
}

func (CLIOutput) FeedFollowCreated(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf(
		"User: %s is now following %s Feed\n",
		feedFollow.Username,
		feedFollow.FeedName,
	 )
}

func (CLIOutput) PrintFeedFollowing(feedFollowing []database.GetFeedFollowsForUserRow) {
	for _, item := range feedFollowing {
		fmt.Printf("- %s\n", item.FeedName)
	}
}

func (CLIOutput) FeedFollowDeleted(feed database.Feed, user database.User) {
	fmt.Printf("User: %s stopped following Feed: %s\n", user.Name, feed.Name)
}

func (CLIOutput) PrintRequestInterval(timeInterval time.Duration) {
	fmt.Printf("Collecting feeds every %s\n", timeInterval.String())
}

func(CLIOutput) Print(str string) {
	fmt.Println(str)
}

func(CLIOutput) PrintPosts(posts []database.GetPostsForUserRow, user database.User) {
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
}