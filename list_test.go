package leetgode

import (
	"context"
	"testing"
)

func TestListCmd(t *testing.T) {
	cmd := &ListCmd{}
	if err := cmd.Run(context.TODO(), []string{}); err != nil {
		t.Errorf("ListCmd() error = %v", err)
	}
}
