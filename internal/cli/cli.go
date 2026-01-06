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

func (CLIOutput) PrintFeed(feed rss.RSSFeed) {
	fmt.Printf("Feed: %+v\n", feed)
}