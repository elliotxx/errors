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
			e := &errorCode{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorCode.Error() = %v, want %v", got, tt.want)
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
			e := &errorCode{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("errorCode.Code() = %v, want %v", got, tt.want)
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
			e := &errorCode{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Msg(); got != tt.want {
				t.Errorf("errorCode.Msg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detailErr_Error(t *testing.T) {
	type fields struct {
		errorCode errorCode
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
		errorCode errorCode
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
				errorCode: errorCode{
					code:    "TestCode",
					message: "test message",
				},
				cause: fmt.Errorf("test cause"),
			},
			want: `ErrorCode: TestCode
Message: test message
Cause: test cause
`,
		},
		{
			name: "missing code",
			fields: fields{
				errorCode: errorCode{
					message: "test message",
				},
				cause: fmt.Errorf("test cause"),
			},
			want: `Message: test message
Cause: test cause
`,
		},
		{
			name: "missing message",
			fields: fields{
				errorCode: errorCode{
					code: "TestCode",
				},
				cause: fmt.Errorf("test cause"),
			},
			want: `ErrorCode: TestCode
Cause: test cause
`,
		},
		{
			name: "missing cause",
			fields: fields{
				errorCode: errorCode{
					code:    "TestCode",
					message: "test message",
				},
				cause: nil,
			},
			want: `ErrorCode: TestCode
Message: test message
`,
		},
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
		errorCode errorCode
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
			if err := e.Cause(); (err != nil) != tt.wantErr {
				t.Errorf("detailErr.Cause() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
