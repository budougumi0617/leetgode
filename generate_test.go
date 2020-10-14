package leetgode

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestGenerateCmd(t *testing.T) {
	t.SkipNow()
	tests := [...]struct {
		name string
		args []string
	}{
		{name: "SuccessGenerateFiles", args: []string{"1"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cmd := &GenerateCmd{}
			if err := cmd.Run(context.TODO(), ioutil.Discard, tt.args); err != nil {
				t.Errorf("GenerateCmd() error = %v", err)
			}
		})
	}
}

func TestGenerateCmd_Error(t *testing.T) {
	tests := [...]struct {
		name string
		args []string
		want error
	}{
		{name: "FailedPaidOnlyProblem", args: []string{"1602"}, want: PaidOnlyError},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cmd := &GenerateCmd{}
			if err := cmd.Run(context.TODO(), ioutil.Discard, tt.args); err != tt.want {
				t.Errorf("GenerateCmd() error = %v", err)
			}
		})
	}
}
