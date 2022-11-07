package errors

import (
	"strings"

	errors2 "github.com/pkg/errors"
)

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(causeStr string) DetailError {
	return &detailErr{
		nil,
		// Create an error that contains the stack trace
		errors2.New(causeStr),
	}
}

func NewErrorCode(code string, message string) ErrorCode {
	return &errorCoder{
		code:    code,
		message: message,
	}
}

func NewDetailError(code string, message string, cause error) DetailError {
	return &detailErr{
		errorCoder: &errorCoder{code, message},
		cause:      cause,
	}
}

func Errorf(format string, args ...interface{}) DetailError {
	return &detailErr{
		nil,
		// Create an error that contains the stack trace
		errors2.Errorf(format, args...),
	}
}

func Code(code string) DetailError {
	return &detailErr{
		errorCoder: &errorCoder{
			code: code,
		},
	}
}

func Msg(msg string) DetailError {
	return &detailErr{
		errorCoder: &errorCoder{
			message: msg,
		},
	}
}

func Cause(cause error) DetailError {
	return &detailErr{
		cause: cause,
	}
}

func Causef(format string, args ...interface{}) DetailError {
	return &detailErr{
		cause: errors2.Errorf(format, args...),
	}
}

func Causewf(err error, format string, args ...interface{}) DetailError {
	return &detailErr{
		cause: errors2.Wrapf(err, format, args...),
	}
}

func InvalidErrorCode(code string) bool {
	return strings.TrimSpace(code) == ""
}

func Zero(e DetailError) bool {
	return e == nil || e.GetCode() == "" && e.GetMsg() == "" && e.GetCause() == nil
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
