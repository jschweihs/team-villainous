package blog

import "errors"

var (
	// ErrEntryNotFound is returned when a blog entry could not be found.
	ErrBlogEntryNotFound = errors.New("Blog entry could not be found")
)
