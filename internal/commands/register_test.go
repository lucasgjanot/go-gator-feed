package commands

import (
	"testing"

	"github.com/lucasgjanot/go-gator-feed/internal/testutil"
)

func TestRegisterCommand(t *testing.T) {
	t.Run("register user successfully", func(t *testing.T) {
		builder := testutil.NewState()
		state := builder.Build()

		cmd := Command{
			Name: "register",
			Args: []string{"registerUser"},
		}

		err := CommandRegister(state, cmd)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verifica se o usuário atual foi atualizado
		if builder.Config.CurrentUser != "registerUser" {
			t.Errorf(
				"expected current user to be 'registerUser', got '%s'",
				builder.Config.CurrentUser,
			)
		}

		// Verifica se o output foi chamado
		if !builder.Output.UserCreatedCalled {
			t.Fatalf("expected UserCreated to be called")
		}

		// Verifica os dados passados para o output
		if builder.Output.CreatedUser.Name != "registerUser" {
			t.Errorf(
				"expected output user name to be 'registerUser', got '%s'",
				builder.Output.CreatedUser.Name,
			)
		}
	})
}
