package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandRegister(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.Database.CreateUser(context.Background(), name)
	if err != nil {
		if runtime.IsUserExistsError(err) {
			return runtime.ErrUserExists
		}
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	s.Output.UserCreated(user)
	return nil
}
