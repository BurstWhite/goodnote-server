package note

import "errors"

var (
	ErrNotFound       = errors.New("note not found")
	ErrInvalidTitle   = errors.New("invalid note title")
	ErrInvalidContent = errors.New("invalid note content")
)
