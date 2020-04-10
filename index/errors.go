package index

import "errors"

var (
	// ErrEmptyKey should be returned by backends when an operation is
	// requested with an empty key.
	ErrEmptyKey = errors.New("invalid sized key")

	// ErrImmutable should be returned by backends when write operation
	// (put/del) is attempted on a readonly.
	ErrImmutable = errors.New("operation not allowed in read-only mode")
)
