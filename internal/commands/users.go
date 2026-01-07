package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandUsers(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.Database.User.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return runtime.ErrNoUsers
	}

	s.Output.ListUsers(s, users)
	return nil
}