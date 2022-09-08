package errors

type ErrorCoder interface {
	error
	Code() string
	Msg() string
}

type Causer interface {
	Cause() error
}

type DetailError interface {
	ErrorCoder
	Causer
	ErrorCode() ErrorCoder
	WithErrorCode(errorCode ErrorCoder) DetailError
	WithCode(code string) DetailError
	WithMsg(msg string) DetailError
	WithCause(err error) DetailError
}
