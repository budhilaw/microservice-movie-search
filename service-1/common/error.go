package common

import "errors"

var (
	ErrNotOk = errors.New("HTTP Response not 200 OK")
)
