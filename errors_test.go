package errors

import (
	"fmt"
	"testing"
)

func Test_errorCode_Error(t *testing.T) {
	type fields struct {
		code    string
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorCoder{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorCoder.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorCode_Code(t *testing.T) {
	type fields struct {
		code    string
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorCoder{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.GetCode(); got != tt.want {
				t.Errorf("errorCoder.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorCode_Msg(t *testing.T) {
	type fields struct {
		code    string
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorCoder{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.GetMsg(); got != tt.want {
				t.Errorf("errorCoder.Msg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detailErr_Error(t *testing.T) {
	type fields struct {
		errorCode errorCoder
		cause     error
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &detailErr{
				&tt.fields.errorCode,
				tt.fields.cause,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("detailErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detailErr_Format(t *testing.T) {
	type fields struct {
		errorCode errorCoder
		cause     error
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "full error",
			fields: fields{
				errorCode: errorCoder{
					code:    "TestCode",
					message: "test message",
				},
				cause: fmt.Errorf("test cause"),
			},
			want: `code [TestCode], msg [test message], cause [test cause]`,
		},
		{
			name: "missing code",
			fields: fields{
				errorCode: errorCoder{
					message: "test message",
				},
				cause: fmt.Errorf("test cause"),
			},
			want: `msg [test message], cause [test cause]`,
		},
		{
			name: "missing message",
			fields: fields{
				errorCode: errorCoder{
					code: "TestCode",
				},
				cause: fmt.Errorf("test cause"),
			},
			want: `code [TestCode], cause [test cause]`,
		},
		{
			name: "missing cause",
			fields: fields{
				errorCode: errorCoder{
					code:    "TestCode",
					message: "test message",
				},
				cause: nil,
			},
			want: `code [TestCode], msg [test message], cause [<nil>]`,
		},
		// {
		// 	name: "full error with stack trace",
		// 	fields: fields{
		// 		errorCoder: errorCoder{
		// 			code:    "TestCode",
		// 			message: "test message",
		// 		},
		// 		cause: errors2.Errorf("test cause"),
		// 	},
		// 	want: ``,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &detailErr{
				&tt.fields.errorCode,
				tt.fields.cause,
			}
			if got := fmt.Sprintf("%+v", e); got != tt.want {
				t.Errorf("detailErr.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detailErr_Cause(t *testing.T) {
	type fields struct {
		errorCode errorCoder
		cause     error
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &detailErr{
				&tt.fields.errorCode,
				tt.fields.cause,
			}
			if err := e.GetCause(); (err != nil) != tt.wantErr {
				t.Errorf("detailErr.Cause() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
