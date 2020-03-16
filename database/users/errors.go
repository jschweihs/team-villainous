package users

import "errors"

var (
	// ErrUserNotFound is returned when a user could not be found.
	ErrUserNotFound = errors.New("User could not be found")
)
