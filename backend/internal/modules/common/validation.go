package common

import "errors"

type ValidationError struct {
	Message string
}

func NewValidationError(message string) error {
	return ValidationError{Message: message}
}

func (err ValidationError) Error() string {
	return err.Message
}

func IsValidationError(err error) bool {
	var validationErr ValidationError
	return errors.As(err, &validationErr)
}
