package leetgode

import (
	"context"
	"testing"
)

func TestGenerateCmd(t *testing.T) {
	tests := [...]struct {
		name    string
		id      string
		wantErr bool
	}{
		{name: "SuccessGenerateFiles", id: "two-sum"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateCmd(context.TODO(), tt.id); (err != nil) != tt.wantErr {
				t.Errorf("GenerateCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
