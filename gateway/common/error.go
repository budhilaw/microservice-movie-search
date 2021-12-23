package common

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("The requested movie is not found")
	ErrConflict            = errors.New("The movie already exist")
)
