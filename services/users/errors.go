package users

import (
	"errors"

	dbusers "vil/database/users"
)

var (
	// ErrEmailEmpty is returned when the email param is empty
	ErrEmailEmpty = errors.New("Email parameter is empty")

	// ErrEmailExists is returned when the email already exists.
	ErrEmailExists = errors.New("Email already exists")

	// ErrPassword is returned when the password is in an invalid format
	ErrPassword = errors.New("Password provided does not meet requirements")

	// ErrInvalidLogin is returned when the login information provided is invalid
	ErrInvalidLogin = errors.New("Login information provided is invalid")

	// ErrUserNotFound is returned when a User could not be found
	ErrUserNotFound = dbusers.ErrUserNotFound
)
