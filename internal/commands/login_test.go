package commands

import (
	"errors"
	"testing"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
	"github.com/lucasgjanot/go-gator-feed/internal/testutil"
)

func TestLoginCommand(t *testing.T) {
	t.Run("sucessful login", func(t *testing.T) {
		builder := testutil.NewState()

		builder.Database.Users["testuser"] = database.User{Name: "testuser"}

		state := builder.Build()
		cmd := Command{Name: "login", Args: []string{"testuser"}}

		err := CommandLogin(state, cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if builder.Config.CurrentUser != "testuser" {
			t.Errorf("expected CurrentUser 'janot', got %s", builder.Config.CurrentUser)
		}
	})

	t.Run("login user not found", func(t *testing.T) {
		state := testutil.NewState().Build()
		cmd := Command{Name: "login", Args: []string{"unknown"}}

		err := CommandLogin(state, cmd)
		if !errors.Is(err, runtime.ErrUserNotFound) {
			t.Fatalf("expected ErrUserNotFound, got %v", err)
		}
	})
}
