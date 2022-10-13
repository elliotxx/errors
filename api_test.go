package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestUsage(t *testing.T) {
	detailErr := New("name is empty")
	_ = assert.Equal(t, "cause [name is empty]", detailErr.Error())

	ErrBlankParameter := NewErrorCode("A0401", "required parameter is blank")
	_ = assert.Equal(t, "code [A0401], msg [required parameter is blank]", ErrBlankParameter.Error())

	detailErr = NewDetailError("A0401", "required parameter is blank", fmt.Errorf("name is empty"))
	_ = assert.Equal(t, "code [A0401], msg [required parameter is blank], cause [name is empty]", detailErr.Error())

	detailErr = Errorf("%s is not exist", "John")
	_ = assert.Equal(t, "cause [John is not exist]", detailErr.Error())

	detailErr = Code("A0401").Msg("required parameter is blank").Cause(fmt.Errorf("name is empty"))
	_ = assert.Equal(t, "code [A0401], msg [required parameter is blank], cause [name is empty]", detailErr.Error())

	detailErr = Code("A0401").Msg("required parameter is blank")
	_ = assert.Equal(t, "code [A0401], msg [required parameter is blank]", detailErr.Error())

	detailErr = Code("A0401").Cause(fmt.Errorf("name is empty"))
	_ = assert.Equal(t, "code [A0401], cause [name is empty]", detailErr.Error())

	detailErr = Msg("required parameter is blank")
	_ = assert.Equal(t, "msg [required parameter is blank]", detailErr.Error())

	detailErr = ErrBlankParameter.Cause(fmt.Errorf("name is empty"))
	_ = assert.Equal(t, "code [A0401], msg [required parameter is blank], cause [name is empty]", detailErr.Error())

	_ = assert.True(t, InvalidErrorCode(""))
}
