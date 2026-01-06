package commands

import (
	"testing"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/testutil"
)

func TestResetCommand(t *testing.T) {
	t.Run("sucessfuly reseting database", func(t *testing.T) {
		builder := testutil.NewState()
		state := builder.Build()

		builder.DB.Users["testuser"] = database.User{Name: "testuser"}

		cmd := Command{
			Name: "reset",
			Args: []string{},
		}

		err := CommandReset(state, cmd)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}


		if len(builder.DB.Users) != 0 {
			t.Fatalf("expected lenght of users table to be 0, got %v", len(builder.DB.Users))
		}
	})
}