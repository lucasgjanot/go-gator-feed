package handlers

import (
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/app"
	"github.com/lucasgjanot/go-gator-feed/internal/commands"
)

func HandlerLogin(s *app.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.Config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}