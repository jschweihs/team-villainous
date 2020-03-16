package roles

import "errors"

var (
	// ErrRoleNotFound is returned when a role could not be found.
	ErrRoleNotFound = errors.New("Role could not be found")
)
