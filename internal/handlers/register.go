package handlers

import (
	"context"
	"fmt"

	"github.com/lucasgjanot/go-gator-feed/internal/app"
	"github.com/lucasgjanot/go-gator-feed/internal/commands"
	"github.com/lucasgjanot/go-gator-feed/internal/database"
)

func HandlerRegister(s *app.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.DB.CreateUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
