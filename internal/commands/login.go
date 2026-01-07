package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandLogin(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}
	name := cmd.Args[0]

	if _, err := s.Database.User.GetUser(context.Background(), name); err != nil {
		if runtime.IsUserNotFoundError(err) {
			return runtime.ErrUserNotFound
		}
		return err
	}

	if err := s.Config.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	username := s.Config.GetCurrentUser()
	s.Output.UserLoggedIn(username)
	return nil
}
