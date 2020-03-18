package file

import "errors"

var (
	// ErrNameInvalid is returned when the multipart form was not parsed
	ErrNameInvalid = errors.New("Name for file provided is invalid")

	// ErrFolderInvalid is returned when the multipart form was not parsed
	ErrFolderInvalid = errors.New("Folder name provided is invalid")
)
