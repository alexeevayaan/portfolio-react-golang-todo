package domain

import (
	"errors"
)

var (
	ErrUUIDInvalid = errors.New("uuid is invalid")
	ErrNotFound = errors.New("not found")
)
