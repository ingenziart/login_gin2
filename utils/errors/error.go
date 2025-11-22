package errors

import "errors"

var( 
	ErrNoFieldToUpdate = errors.New("at least one field must be provided")
	Errno
)
