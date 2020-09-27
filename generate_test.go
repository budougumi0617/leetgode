package leetgode

import (
	"context"
	"testing"
)

func TestGenerateCmd(t *testing.T) {
	tests := [...]struct {
		name string
		id   int
	}{
		{name: "SuccessGenerateFiles", id: 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateCmd(context.TODO(), tt.id); err != nil {
				t.Errorf("GenerateCmd() error = %v", err)
			}
		})
	}
}
