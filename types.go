package errors

import "fmt"

type ErrorCode interface {
	error
	fmt.Formatter

	GetCode() string
	GetMsg() string

	Code(code string) ErrorCode
	Msg(msg string) ErrorCode

	Cause(err error) DetailError
}

type DetailError interface {
	error
	fmt.Formatter

	GetCode() string
	GetMsg() string
	GetCause() error
	GetErrorCode() ErrorCode

	Code(code string) DetailError
	Msg(msg string) DetailError
	Cause(err error) DetailError
	ErrorCode(errorCode ErrorCode) DetailError
}
