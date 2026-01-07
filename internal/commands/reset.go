package commands

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)



func CommandReset(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.Database.User.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	s.Output.ResetedDatabase()
	return nil
}