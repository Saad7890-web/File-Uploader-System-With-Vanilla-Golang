package file

import "errors"

var (
	ErrInvalidOwner = errors.New("invalid file owner")
	ErrInvalidName  = errors.New("invalid file name")
	ErrInvalidSize  = errors.New("invalid file size")

	ErrFileNotFound = errors.New("file not found")
	ErrAccessDenied = errors.New("access denied")
)
