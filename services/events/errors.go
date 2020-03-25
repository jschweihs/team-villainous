package events

import (
	"errors"
	dbevents "vil/database/events"
)

var (
	// ErrNameEmpty is returned when the name param is empty
	ErrNameEmpty = errors.New("Name parameter is empty")

	// ErrEventNotFound is returned when a Event could not be found
	ErrEventNotFound = dbevents.ErrEventNotFound
)
