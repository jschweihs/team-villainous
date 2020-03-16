package roles

import "errors"

var (
	// ErrLimitInvalid is returned when the limit parameter is invalid.
	ErrLimitInvalid = errors.New("Limit parameter is invalid, must be an integer")

	// ErrLimitMax is returned when the limit parameter is greater than the
	// maximum allowable limit.
	ErrLimitMax = errors.New("Limit parameter is greater than maximum allowable limit")
)
