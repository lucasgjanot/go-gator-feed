package runtime

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
	ErrNoUsers = errors.New("no users registered")
	ErrFeedNotFound = errors.New("feed not found")
	ErrFeedExists   = errors.New("feed already exists")
	ErrNoFeed = errors.New("no feeds registered")
	ErrFeedFollowExists = errors.New("user already follow this feed")
	ErrFeedFollowNotFound = errors.New("user isn't following any feed")
	ErrFeedFollowNotExist = errors.New("user don't follow this feed")
	ErrNoFeedToFetch = errors.New("No more feeds to fetch")
)

func IsExistsError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505" // unique_violation
	}
	return false
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}