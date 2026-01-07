package cli

import (
	"fmt"

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
}

func (CLIOutput) PrintFeed(feed rss.RSSFeed) {
	fmt.Printf("Feed: %+v\n", feed)
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