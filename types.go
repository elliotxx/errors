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
	Causef(format string, args ...interface{}) DetailError
	Causewf(err error, format string, args ...interface{}) DetailError
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
	Causef(format string, args ...interface{}) DetailError
	Causewf(err error, format string, args ...interface{}) DetailError
	ErrorCode(errorCode ErrorCode) DetailError
}
