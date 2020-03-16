package blog

import (
	"errors"
	dbblog "vil/database/blog"
)

var (
	// ErrTitleEmpty is returned when no title is provided
	ErrTitleEmpty = errors.New("Title parameter is empty")

	// ErrUserIDEmpty is returned when no user id is provided
	ErrUserIDEmpty = errors.New("UserID parameter is empty")

	// ErrContentEmpty is returned when no content is provided
	ErrContentEmpty = errors.New("Content parameter is empty")

	// ErrBlogEntryNotFound is returned when a blog entry could not be found
	ErrBlogEntryNotFound = dbblog.ErrBlogEntryNotFound
)
