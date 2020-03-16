package roles

import (
	"errors"

	dbroles "vil/database/roles"
)

var (

	// ErrRoleEmpty is returned when there is no role provided.
	ErrRoleEmpty = errors.New("No role was provided")

	// ErrRoleExists is returned when the role already exists.
	ErrRoleExists = errors.New("Role already exists")

	// ErrRoleNotFound is returned when a Role could not be found
	ErrRoleNotFound = dbroles.ErrRoleNotFound
)
