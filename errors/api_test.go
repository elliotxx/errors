package errors

import (
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name: "t1",
			args: args{
				err:     NewDetailError("InvalidParams", "parameters must be valid", fmt.Errorf("name is empty")),
				message: "failed to wrap",
			},
			wantErr: "failed to wrap: code [InvalidParams], msg [parameters must be valid], cause [name is empty]",
		},
		{
			name: "t2",
			args: args{
				err:     Wrap(NewDetailError("InvalidParams", "parameters must be valid", fmt.Errorf("name is empty")), "failed to wrap"),
				message: "failed to wrap2",
			},
			wantErr: "failed to wrap2: failed to wrap: code [InvalidParams], msg [parameters must be valid], cause [name is empty]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err, tt.args.message); (err != nil) && err.Error() != tt.wantErr {
				t.Errorf("Wrap().Error() error = %v, wantErr %v", err.Error(), tt.wantErr)
			}
		})
	}
}
