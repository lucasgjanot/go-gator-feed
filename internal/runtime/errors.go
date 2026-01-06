package runtime

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

func IsUserExistsError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505" // unique_violation
	}
	return false
}

func IsUserNotFoundError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}