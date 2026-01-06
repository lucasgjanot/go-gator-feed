package commands

import (
	"context"
	"fmt"


	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)



func CommandReset(s *runtime.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	err := s.Database.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	s.Output.ResetedDatabase()
	return nil
}