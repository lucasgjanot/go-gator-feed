package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lucasgjanot/go-gator-feed/internal/database"
	"github.com/lucasgjanot/go-gator-feed/internal/runtime"
)

func CommandBrowse(s *runtime.State, cmd Command, user database.User) error {
	if len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	limit := 2

	if len(cmd.Args) == 1 {
		if val, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = val
		} else {
			return fmt.Errorf("invalid limit parameter, value must be a number")
		}
	}


	posts, err := s.Database.Post.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			Name: user.Name,
			Limit: int32(limit),
		},
	)
	if err != nil {
		if runtime.IsNotFoundError(err) {
			return runtime.ErrFeedFollowNotExist
		}
		return err
	}

	s.Output.PrintPosts(posts, user)
	return nil
}