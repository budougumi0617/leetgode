package leetgode

import (
	"context"
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
			if err := cmd.Run(context.TODO(), tt.args); err != nil {
				t.Errorf("GenerateCmd() error = %v", err)
			}
		})
	}
}
