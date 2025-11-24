package errors

import "errors"

var (
	ErrNoFieldToUpdate     = errors.New("at least one field must be provided")
	ErrFailedToExtractInDb = errors.New("fail to extract in db ")
)
