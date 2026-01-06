package cli

import (
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
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
