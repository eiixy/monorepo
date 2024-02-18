package errors

import "github.com/pkg/errors"

func New(msg string) error {
	return errors.New(msg)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
