package middleware

import (
	"context"

	"github.com/lucasgjanot/go-gator-feed/internal/commands"
	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func MiddlewareLoggedIn(
	handler func(s *runtime.State, cmd commands.Command, user database.User) error,
) func(*runtime.State, commands.Command) error {
	return func(s *runtime.State, cmd commands.Command) error {
		user, err := s.Database.User.GetUser(
			context.Background(),
			s.Config.GetCurrentUser(),
		)
		if err != nil {
			if runtime.IsNotFoundError(err) {
				return runtime.ErrUserNotFound
			}
			return err
		}

		return handler(s, cmd, user)
	}
}
