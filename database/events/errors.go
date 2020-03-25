package events

import "errors"

var (
	// ErrEventNotFound is returned when an event could not be found.
	ErrEventNotFound = errors.New("Event could not be found")
)
