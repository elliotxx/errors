package errors

import (
	"strings"

	errors2 "github.com/pkg/errors"
)

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(causeMsg string) DetailError {
	return &detailErr{
		nil,
		// Create an error that contains the stack trace
		errors2.New(causeMsg),
	}
}

func Errorf(format string, args ...interface{}) DetailError {
	return &detailErr{
		nil,
		// Create an error that contains the stack trace
		errors2.Errorf(format, args...),
	}
}

func WithErrorCode(errCode ErrorCoder, cause error) DetailError {
	if errCode != nil {
		if e, ok := errCode.(*errorCode); ok {
			return &detailErr{
				e,
				cause,
			}
		}
	}

	return &detailErr{nil, cause}
}

func NewErrorCode(code string, message string) ErrorCoder {
	return &errorCode{
		code:    code,
		message: message,
	}
}

func NewDetailError(code string, message string, cause error) DetailError {
	return &detailErr{
		errorCode: &errorCode{code, message},
		cause:     cause,
	}
}

func Msg(err error) string {
	if err != nil {
		if e, ok := err.(ErrorCoder); ok {
			return e.Msg()
		}

		return err.Error()
	}

	return ""
}

func Cause(err error) error {
	if err != nil {
		if e, ok := err.(Causer); ok {
			return e.Cause()
		}
	}

	return err
}

func ErrorCode(err error) ErrorCoder {
	if err != nil {
		if e, ok := err.(ErrorCoder); ok {
			return e
		}
	}

	return nil
}

func InvalidErrorCode(code string) bool {
	return strings.TrimSpace(code) == ""
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return errors2.Wrap(err, message)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return errors2.Wrapf(err, format, args...)
}

func Is(err, target error) bool { return errors2.Is(err, target) }

func As(err error, target interface{}) bool { return errors2.As(err, target) }

func Unwrap(err error) error {
	return errors2.Unwrap(err)
}
