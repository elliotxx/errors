package errors

import (
	"fmt"
	"io"
)

var _ ErrorCoder = (*errorCode)(nil)

type errorCode struct {
	code    string
	message string
}

func (e *errorCode) Code() string {
	if e == nil {
		return ""
	}

	return e.code
}

func (e *errorCode) Msg() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *errorCode) Error() string {
	if e == nil {
		return ""
	}

	switch {
	case e.code != "" && e.message == "":
		return fmt.Sprintf("code [%s]", e.code)
	case e.code == "" && e.message != "":
		return fmt.Sprintf("msg [%s]", e.message)
	case e.code != "" && e.message != "":
		return fmt.Sprintf("code [%s], msg [%s]", e.code, e.message)
	default:
		return ""
	}
}

func (e *errorCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s", e.Error())

			return
		}

		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

var _ DetailError = (*detailErr)(nil)

type detailErr struct {
	*errorCode
	cause error
}

func (e *detailErr) Cause() error {
	if e == nil {
		return nil
	}

	return e.cause
}

func (e *detailErr) ErrorCode() ErrorCoder {
	return e.errorCode
}

func (e *detailErr) WithCode(code string) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.code = code

	return ee
}

func (e *detailErr) WithMsg(msg string) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.message = msg

	return ee
}

func (e *detailErr) WithCause(err error) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.cause = err

	return ee
}

func (e *detailErr) WithErrorCode(errCode ErrorCoder) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.errorCode = &errorCode{
		code:    errCode.Code(),
		message: errCode.Msg(),
	}

	return ee
}

func (e *detailErr) Error() string {
	if e == nil {
		return ""
	}

	codeErr := e.errorCode.Error()

	switch codeErr {
	case "":
		return fmt.Sprintf("cause [%s]", e.cause.Error())
	default:
		return fmt.Sprintf("%s, cause [%s]", codeErr, e.cause.Error())
	}
}

func (e *detailErr) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			codeErr := fmt.Sprintf("%+v", e.errorCode)

			switch codeErr {
			case "":
				fmt.Fprintf(s, "cause [%+v]", e.cause)
			default:
				fmt.Fprintf(s, "%s, cause [%+v]", codeErr, e.cause)
			}

			return
		}

		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

func copyDetailErr(e *detailErr) *detailErr {
	return &detailErr{
		&errorCode{
			code:    e.Code(),
			message: e.Msg(),
		},
		e.cause,
	}
}
