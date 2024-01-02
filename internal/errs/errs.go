package errs

import "errors"

var (
	ErrNotImplemented = errors.New("not implemented")

	ErrNotFound       = errors.New("not found")
	ErrExists         = errors.New("exists")
	ErrNoUpdate       = errors.New("nothing to update")
	ErrNoPrimaryEmail = errors.New("no primary email was provided")
)
