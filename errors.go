package errors

import (
	"fmt"
	"io"

	errors2 "github.com/pkg/errors"
)

var _ ErrorCode = (*errorCoder)(nil)

type errorCoder struct {
	code    string
	message string
}

func (e *errorCoder) GetCode() string {
	if e == nil {
		return ""
	}

	return e.code
}

func (e *errorCoder) GetMsg() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *errorCoder) Code(code string) ErrorCode {
	if e == nil {
		return e
	}

	ee := copyErrorCode(e)
	ee.code = code

	return ee
}

func (e *errorCoder) Msg(msg string) ErrorCode {
	if e == nil {
		return e
	}

	ee := copyErrorCode(e)
	ee.message = msg

	return ee
}

func (e *errorCoder) Cause(err error) DetailError {
	if e == nil {
		return nil
	}

	ee := copyFromErrorCode(e)
	ee.cause = err

	return ee
}

func (e *errorCoder) Causef(err error, format string, args ...interface{}) DetailError {
	return e.Cause(errors2.Wrapf(err, format, args...))
}

func (e *errorCoder) Error() string {
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

func (e *errorCoder) Format(s fmt.State, verb rune) {
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
	*errorCoder
	cause error
}

func (e *detailErr) GetCode() string {
	if e == nil || e.errorCoder == nil {
		return ""
	}

	return e.code
}

func (e *detailErr) GetMsg() string {
	if e == nil || e.errorCoder == nil {
		return ""
	}

	return e.message
}

func (e *detailErr) GetCause() error {
	if e == nil {
		return nil
	}

	return e.cause
}

func (e *detailErr) GetErrorCode() ErrorCode {
	return e.errorCoder
}

func (e *detailErr) Code(code string) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.code = code

	return ee
}

func (e *detailErr) Msg(msg string) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.message = msg

	return ee
}

func (e *detailErr) Cause(err error) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.cause = err

	return ee
}

func (e *detailErr) Causef(err error, format string, args ...interface{}) DetailError {
	return e.Cause(errors2.Wrapf(err, format, args...))
}

func (e *detailErr) ErrorCode(errCode ErrorCode) DetailError {
	if e == nil {
		return e
	}

	ee := copyDetailErr(e)
	ee.errorCoder = &errorCoder{
		code:    errCode.GetCode(),
		message: errCode.GetMsg(),
	}

	return ee
}

func (e *detailErr) Error() string {
	if Zero(e) {
		return ""
	}

	if e.errorCoder != nil && e.cause == nil {
		return e.errorCoder.Error()
	}

	if e.errorCoder == nil && e.cause != nil {
		return fmt.Sprintf("cause [%s]", e.cause.Error())
	}

	return fmt.Sprintf("%s, cause [%s]", e.errorCoder.Error(), e.cause.Error())
}

func (e *detailErr) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			codeErr := fmt.Sprintf("%+v", e.errorCoder)

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

func copyFromErrorCode(e *errorCoder) *detailErr {
	return &detailErr{
		&errorCoder{
			code:    e.GetCode(),
			message: e.GetMsg(),
		},
		nil,
	}
}

func copyDetailErr(e *detailErr) *detailErr {
	return &detailErr{
		&errorCoder{
			code:    e.GetCode(),
			message: e.GetMsg(),
		},
		e.cause,
	}
}

func copyErrorCode(c *errorCoder) *errorCoder {
	return &errorCoder{
		code:    c.GetCode(),
		message: c.GetMsg(),
	}
}
