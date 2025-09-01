package errors

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrNilValue    = errors.New("nil value not allowed")
	ErrEmptyKey    = errors.New("empty key not allowed")
)