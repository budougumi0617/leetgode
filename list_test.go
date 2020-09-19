package leetgode

import (
	"context"
	"testing"
)

func TestListCmd(t *testing.T) {
	if err := ListCmd(context.TODO()); err != nil {
		t.Errorf("ListCmd() error = %v", err)
	}
}
